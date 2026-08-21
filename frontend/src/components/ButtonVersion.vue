<script setup lang="ts">
import {onMounted, ref} from "vue";
import {GetLatestRelease, GetVersion, UpdateTo} from "../../wailsjs/go/handler/App";
import {ElLoading, ElMessageBox} from "element-plus";
import {NotifyError, NotifySuccess} from "../utils/notify.ts";
import {useI18n} from "vue-i18n";

const {t} = useI18n()
const version = ref('dev')

const checkUpdate = () => {
  GetLatestRelease().then(release => {
    if (release.tag_name != version.value) {
      const changelog = release.body ? `<br>${release.body}` : ''
      ElMessageBox.confirm(
        t('version.update.notify') + changelog,
        release.tag_name,
        {
          confirmButtonText: t('version.update.confirm'),
          cancelButtonText: t('version.update.cancel'),
          type: 'info',
          dangerouslyUseHTMLString: true, // 如需渲染 HTML
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

<template>
  <el-tag class="ml-2 cursor-pointer" size="small" type="success" effect="light" @click="checkUpdate">
    {{ version }}
  </el-tag>
</template>