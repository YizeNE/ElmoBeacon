<script setup lang="ts">
import { nextTick, onMounted, onUnmounted, ref } from "vue";
import { GetLatestRelease, GetVersion, UpdateTo, CancelUpdate } from "../../wailsjs/go/handler/App";
import { EventsOn, EventsOff } from "../../wailsjs/runtime/runtime";
import { ElMessageBox } from "element-plus";
import { NotifyError, NotifySuccess } from "../utils/notify.ts";
import { useI18n } from "vue-i18n";
import { marked } from "marked";
import { BrowserOpenURL } from "../../wailsjs/runtime/runtime";
import { Loading } from '@element-plus/icons-vue'
 
const { t } = useI18n()
const version = ref('dev')
const checking = ref(false)
 
// 下载状态
const downloading = ref(false)
const downloadPercent = ref(0)
const downloadedBytes = ref(0)
const totalBytes = ref(0)
 
const formatBytes = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const units = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return (bytes / Math.pow(1024, i)).toFixed(1) + ' ' + units[i]
}
 
const checkUpdate = () => {
  if (checking.value) return
  checking.value = true
 
  GetLatestRelease().then(release => {
    checking.value = false
    if (release.tag_name != version.value) {
      const changelog = release.body
        ? `<div class="update-dialog">
              <p class="update-title">${t('version.update.notify')} ${release.tag_name}</p>
              <div style="height:10px;"></div>
              <div class="changelog">
                <p class="changelog-label">${t('version.update.changelog')}</p>
                <div class="prose prose-sm changelog-content">${marked.parse(release.body)}</div>
              </div>
          </div>`
        : ''
      ElMessageBox.confirm(
        changelog,
        {
          confirmButtonText: t('version.update.confirm'),
          cancelButtonText: t('version.update.cancel'),
          dangerouslyUseHTMLString: true,
          customClass: 'update-message-box',
        }
      ).then(() => {
        startUpdate(release.tag_name)
      })
      
      // 拦截 changelog 内的链接，用系统浏览器打开
      nextTick(() => {
        const content = document.querySelector('.changelog-content')
        if (content) {
          content.addEventListener('click', (e) => {
            const target = e.target as HTMLElement
            const anchor = target.closest('a')
            if (anchor && anchor.href) {
              e.preventDefault()
              BrowserOpenURL(anchor.href)
            }
          })
        }
      })
    } else {
      NotifySuccess(version.value, t('version.update.latest'))
    }
  }).catch(err => {
    checking.value = false
    NotifyError('Error', err)
  })
}
 
const startUpdate = (targetVersion: string) => {
  downloading.value = true
  downloadPercent.value = 0
  downloadedBytes.value = 0
  totalBytes.value = 0
 
  // 监听进度事件
  EventsOn("update:progress", (data: any) => {
    downloadPercent.value = data.percent
    downloadedBytes.value = data.downloaded
    totalBytes.value = data.total
  })
 
  UpdateTo(targetVersion).catch(err => {
    if (err !== 'update canceled by user') {
      NotifyError('Error', err)
    }
  }).finally(() => {
    downloading.value = false
    EventsOff("update:progress")
  })
}
 
const handleCancelUpdate = () => {
  CancelUpdate()
}
 
onMounted(async () => {
  await GetVersion().then(res => {
    if (res) {
      version.value = res;
    }
  })
  if (version.value != "dev") {
    checkUpdate()
  }
})
 
onUnmounted(() => {
  EventsOff("update:progress")
})
</script>
 
<style>
.update-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 16px;
}
 
.changelog {
  max-height: 260px;
  overflow-y: auto;
}
 
.changelog-label {
  font-size: 12px;
  font-weight: 600;
  margin-bottom: 8px;
}
 
.update-message-box .el-message-box__message {
  text-align: center;
  width: 100%;
}
 
.update-message-box .changelog {
  text-align: left;
}
 
.update-message-box .el-message-box__btns {
  display: flex;
  justify-content: center;
  gap: 24px;
  padding: 16px 0 8px;
}
 
.update-message-box .el-message-box__btns button {
  width: 120px;
  height: 38px;
  font-size: 14px;
}
 
.update-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}
 
.update-panel {
  background: var(--theme-from, #0c1a32);
  border-radius: 12px;
  padding: 32px 40px;
  min-width: 360px;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.1);
}
 
.update-panel-title {
  font-size: 16px;
  font-weight: 600;
  color: #e0e0e0;
  margin-bottom: 20px;
}
 
.update-panel-progress {
  margin-bottom: 8px;
}
 
.update-panel-bytes {
  font-size: 12px;
  color: #888;
  margin-bottom: 20px;
}
 
.update-panel-cancel {
  width: 100px;
}
</style>
 
<template>
  <!-- 下载进度面板 -->
  <div v-if="downloading" class="update-overlay">
    <div class="update-panel">
      <p class="update-panel-title">{{ t('version.update.downloading') }} {{ downloadPercent }}%</p>
      <el-progress
        class="update-panel-progress"
        :percentage="downloadPercent"
        :stroke-width="8"
        :show-text="false"
        color="#10b981"
      />
      <p class="update-panel-bytes">
        {{ formatBytes(downloadedBytes) }} / {{ formatBytes(totalBytes) }}
      </p>
      <el-button
        class="update-panel-cancel"
        size="small"
        @click="handleCancelUpdate"
      >
        {{ t('version.update.cancel') }}
      </el-button>
    </div>
  </div>
 
  <el-tooltip :content="t('version.checkUpdate')" placement="bottom" effect="dark">
    <span
      class="inline-flex items-center justify-center px-2.5 py-0.5 rounded-full bg-emerald-500/20 border border-emerald-400/30 text-emerald-300 text-xs font-mono cursor-pointer hover:bg-emerald-500/30 transition-colors"
      style="--wails-draggable:no-drag"
      @click="checkUpdate"
    >
      {{ version }}
 
      <el-icon v-if="checking" class="is-loading ml-1" style="font-size: 14px;">
        <Loading />
      </el-icon>
    </span>
  </el-tooltip>
</template>