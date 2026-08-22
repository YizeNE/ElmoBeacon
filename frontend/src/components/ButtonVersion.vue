<script setup lang="ts">
import { onMounted, ref } from "vue";
import { GetLatestRelease, GetVersion, UpdateTo } from "../../wailsjs/go/handler/App";
import { ElLoading, ElMessageBox } from "element-plus";
import { NotifyError, NotifySuccess } from "../utils/notify.ts";
import { useI18n } from "vue-i18n";
import { marked } from "marked";

const { t } = useI18n()
const version = ref('dev')

const checkUpdate = () => {
  GetLatestRelease().then(release => {
    if (release.tag_name != version.value) {
      const changelog = release.body
        ? `<div class="update-dialog">
              <p class="update-title">${t('version.update.notify')} ${release.tag_name}</p>
              <div style="height:10px;"></div>
              <div class="changelog">
                <p class="changelog-label">${t('version.update.changelog')}</p>
                <div class="prose prose-sm">${marked.parse(release.body)}</div>
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
        const loading = ElLoading.service({
          lock: true,
          text: `Update to ${release.tag_name}...`,
          background: 'rgba(0, 0, 0, 0.7)'
        })
        UpdateTo(release.tag_name).catch(err => {
          NotifyError('Error', err)
        }).finally(() => {
          loading.close()
        })
      })
    } else {
      NotifySuccess(release.tag_name, t('version.update.latest'))
    }
  })
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

/* 通过 customClass控制弹窗的 footer */
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

</style>

<template>
  <el-tag class="ml-2 cursor-pointer" size="small" type="success" effect="light" @click="checkUpdate">
    {{ version }}
  </el-tag>
</template>