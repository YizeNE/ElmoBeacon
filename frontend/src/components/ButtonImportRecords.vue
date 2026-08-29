<script setup lang="ts">
import { ElLoading } from "element-plus";
import { ImportRecords } from "../../wailsjs/go/handler/App";
import { useI18n } from "vue-i18n";
import { useUserStore } from "../store/userStore.ts";
import { usePoolStore } from "../store/poolStore.ts";
import { NotifyError, NotifySuccess } from "../utils/notify.ts";
import { ref } from "vue";
 
const { t } = useI18n()
const userStore = useUserStore()
const poolStore = usePoolStore()
 
const dialogVisible = ref(false)
const loading = ref(false)
const formData = ref({
  uid: '',
  server: 'cn',
  fileContent: ''
})
 
const serverOptions = [
  { key: 'hao-jp',   value: 'jp'},
  { key: 'hao-kr',   value: 'kr'},
  { key: 'hao-asia', value: 'tw'},
  { key: 'hao-intl', value: 'intl'},
  { key: 'dw-cn',    value: 'cn' },
  { key: 'dw-us',    value: 'us'},
]
 
const selectImportFile = async () => {
  const input = document.createElement('input')
  input.type = 'file'
  input.accept = '.json'
  input.onchange = async (e: Event) => {
    const file = (e.target as HTMLInputElement).files?.[0]
    if (!file) return
    try {
      const text = await file.text()
      JSON.parse(text)
      formData.value.fileContent = text
    } catch {
      NotifyError(t('import.dialog.importError'), t('import.dialog.invalidJson'))
    }
  }
  input.click()
}
 
const openDialog = () => {
  if (userStore.userId && userStore.userList) {
    const currentUser = userStore.userList.find(u => u.id === userStore.userId);
    if (currentUser) {
      formData.value.uid = currentUser.uid.toString();
      formData.value.server = currentUser.server;
    } else {
      formData.value.uid = '';
    }
  }
  formData.value.fileContent = ''
  dialogVisible.value = true
}
 
const handleSubmit = async () => {
  if (!formData.value.uid) {
    NotifyError(t('import.dialog.importError'), t('import.dialog.uidRequired'))
    return
  }
  if (!formData.value.fileContent) {
    NotifyError(t('import.dialog.importError'), t('import.dialog.fileRequired'))
    return
  }
 
  const elLoading = ElLoading.service({
    lock: true,
    text: t('import.loading'),
    background: 'rgba(0, 0, 0, 0.7)',
  })
  loading.value = true
 
  try {
    const importData = JSON.parse(formData.value.fileContent)
    const result = await ImportRecords(
      parseInt(formData.value.uid),
      formData.value.server,
      importData
    )
 
    NotifySuccess(
      t('import.result.success.title', { server: t(`server.${result.Server}`), uid: result.Uid }),
      result.DiffList ? result.DiffList.map(diff =>
        t('import.result.success.changed', {
          poolType: t(`gacha.type.${diff.PoolType}`),
          successCount: diff.SuccessCount,
          failCount: diff.FailCount
        })
      ).join('<br/>') : t('import.result.success.unchanged')
    )
 
    if (result.Id === userStore.userId) {
      await poolStore.updatePoolInfo()
    } else {
      await userStore.updateUserList()
      const targetUser = userStore.userList.find(u => u.uid === result.Uid)
      if (targetUser) {
        await userStore.updateUserId(targetUser.id)
      }
      await poolStore.updatePoolInfo()
    }
 
    dialogVisible.value = false
  } catch (err) {
    NotifyError(t('import.dialog.importError'), err)
  } finally {
    loading.value = false
    elLoading.close()
  }
}
</script>
 
<template>
  <div v-if="dialogVisible" class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="!loading && (dialogVisible = false)"></div>
 
    <div class="relative bg-white/10 backdrop-blur-xl rounded-xl shadow-2xl w-full max-w-md mx-4 overflow-hidden border border-white/15">
 
      <div class="px-6 py-4 transition-colors duration-500"
        :style="{ background: `linear-gradient(to right, var(--theme-from), var(--theme-to))` }">
        <h2 class="text-white text-lg font-semibold">{{ $t('import.dialog.title') }}</h2>
      </div>
 
      <div class="p-6 space-y-5">
        <div>
          <label class="block text-gray-200 text-sm font-medium mb-2">{{ $t('import.dialog.uid') }}</label>
          <input v-model="formData.uid" type="text" inputmode="numeric" pattern="[0-9]*"
            :placeholder="$t('import.dialog.uidPlaceholder')"
            class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-[var(--theme-to)] focus:border-transparent transition-colors"
            :disabled="loading" />
        </div>
 
        <div>
          <label class="block text-gray-200 text-sm font-medium mb-2">{{ $t('import.dialog.server') }}</label>
          <select v-model="formData.server"
            class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-lg text-white focus:outline-none focus:ring-2 focus:ring-[var(--theme-to)] transition-colors"
            :disabled="loading">
            <option v-for="opt in serverOptions" :key="opt.value" :value="opt.value"
              class="bg-gray-800 text-white">{{ $t(`server.${opt.value}`)}}</option>
          </select>
        </div>
 
        <div>
          <label class="block text-gray-200 text-sm font-medium mb-2">{{ $t('import.dialog.file') }}</label>
          <div class="flex gap-2">
            <input :value="formData.fileContent ? $t('import.dialog.fileSelected') : ''" type="text" readonly
              :placeholder="$t('import.dialog.filePlaceholder')"
              class="flex-1 px-4 py-3 bg-white/10 border border-white/20 rounded-lg text-white placeholder-gray-400 transition-colors cursor-pointer overflow-hidden text-ellipsis whitespace-nowrap"
              @click="!loading && selectImportFile()" />
            <button @click="selectImportFile()" :disabled="loading"
              class="px-4 py-3 bg-white/20 hover:bg-white/30 disabled:bg-white/10 disabled:cursor-not-allowed text-white rounded-lg transition-colors whitespace-nowrap">
              {{ $t('import.dialog.browse') }}
            </button>
          </div>
        </div>
 
        <div class="bg-white/10 border-l-4 border-[var(--theme-to)] p-4 rounded">
          <div class="flex items-start gap-3">
            <svg class="w-5 h-5 text-white mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                clip-rule="evenodd" />
            </svg>
            <p class="text-gray-100 text-sm leading-relaxed break-words">{{ $t('import.dialog.tips') }}</p>
          </div>
        </div>
      </div>
 
      <div class="px-6 py-4 bg-white/5 border-t border-white/10 flex justify-end gap-3">
        <button @click="dialogVisible = false" :disabled="loading"
          class="px-6 py-2.5 border border-white/20 text-gray-200 hover:text-white hover:bg-white/10 disabled:bg-white/10 disabled:text-gray-500 rounded-lg transition-colors">
          {{ $t('import.dialog.cancel') }}
        </button>
        <button @click="handleSubmit" :disabled="loading"
          class="px-6 py-2.5 text-white rounded-lg transition-all flex items-center gap-2 shadow-lg border border-white/20 disabled:opacity-50 disabled:cursor-not-allowed"
          :style="{ background: `linear-gradient(to right, var(--theme-from), var(--theme-to))` }">
          <span v-if="loading" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
          <span>{{ loading ? $t('import.dialog.importing') : $t('import.dialog.startImport') }}</span>
        </button>
      </div>
    </div>
  </div>
 
  <el-tooltip :content="$t('import.button.title')" placement="bottom">
    <button
      @click="openDialog"
      :disabled="loading"
      class="p-1.5 rounded-full text-white/80 hover:text-white hover:bg-white/10 transition-colors"
      style="--wails-draggable:no-drag"
    >
      <i-mdi-file-import class="h-6 w-6" />
    </button>
  </el-tooltip>
</template>