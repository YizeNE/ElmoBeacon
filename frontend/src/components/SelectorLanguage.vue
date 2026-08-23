<script setup lang="ts">
import {useLangStore} from "../store/langStore.ts";

const langOptions = [
  {label: '简体中文', value: 'zh-CN'},
  {label: '繁體中文', value: 'zh-TW'},
  {label: 'English', value: 'en'},
  {label: '日本語', value: 'ja'},
  {label: '한국어', value: 'kr'},
]
const langStore = useLangStore()
</script>

<template>
  <el-dropdown>
    <button class="p-1.5 rounded-full hover:bg-white/10 transition-colors flex items-center">
      <i-mdi-translate class="h-6 w-6 text-white/80 hover:text-white" />
    </button>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-for="lang in langOptions" @click="langStore.updateLang(lang.value)">
          <span :class="lang.value==langStore.lang?'text-red-400':''">{{ lang.label }}</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>