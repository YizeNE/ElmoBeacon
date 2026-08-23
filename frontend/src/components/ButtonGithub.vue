<script setup lang="ts">
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime';
import { NotifyError } from '../utils/notify';
import { useI18n } from "vue-i18n";

const { t } = useI18n()
const PROJECT_URL = 'https://github.com/YizeNE/ElmoBeacon';

const openProjectLink = async () => {
  try {
    // 使用 Wails 提供的标准 API，会在系统默认浏览器打开，而不是在软件内打开
    await BrowserOpenURL(PROJECT_URL);
  } catch (err) {
    NotifyError(t('github.error'), err)
  }
};
</script>

<template>
  <el-tooltip :content="$t('github.title')" placement="bottom" effect="dark">
    <button
      @click="openProjectLink"
      class="p-1.5 rounded-full hover:bg-white/10 text-white/80 hover:text-white transition-colors flex items-center"
      style="--wails-draggable: no-drag"
    >
      <i-mdi-github class="text-3xl" />
    </button>
  </el-tooltip>
</template>
