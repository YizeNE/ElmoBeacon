<script setup lang="ts">
import {ElLoading} from "element-plus";
import {SyncRecords} from "../../wailsjs/go/handler/App";
import {useI18n} from "vue-i18n";
import {useUserStore} from "../store/userStore.ts";
import {usePoolStore} from "../store/poolStore.ts";
import {service} from "../../wailsjs/go/models.ts";
import {NotifyError, NotifySuccess} from "../utils/notify.ts";
import {ref} from "vue";
import {SelectFilePath} from "../../wailsjs/go/handler/App";
 
const {t} = useI18n()
const userStore = useUserStore()
const poolStore = usePoolStore()
 
// 弹窗相关状态
const dialogVisible = ref(false)
const loading = ref(false)
const formData = ref({
  uid: '',
  gameDataDir: ''
})
 
// 同步处理函数
const handleSyncResult = (res: service.SyncResult) => {
  NotifySuccess(
    t('sync.result.success.title', {server: t(`server.${res.Server}`), uid: res.Uid}),
    res.DiffList ? res.DiffList.map(diff => {
      return t('sync.result.success.changed', {poolType: t(`gacha.type.${diff.PoolType}`), count: diff.Count})
    }).join("<br/>") : t('sync.result.success.unchanged')
  )
}
 
// 选择文件路径
const selectGameDataDir = async () => {
  try {
    const result = await SelectFilePath()
    if (result) {
      formData.value.gameDataDir = result
    }
  } catch (err) {
    NotifyError(t('sync.dialog.syncError'), t('sync.dialog.selectDirError'))
  }
}
 
// 同步记录
const syncRecords = async () => {
  // 设置默认UID：如果有当前用户则显示其UID，否则为空
  if (userStore.userId && userStore.userList){
    const currentUser = userStore.userList.find(u => u.id === userStore.userId);
    if (currentUser) {
        formData.value.uid = currentUser.uid.toString();
      }else {
      formData.value.uid = '';
    }
  }
  formData.value.gameDataDir = ''
  dialogVisible.value = true
}
 
// 提交同步
const handleSubmit = async () => {
  if (!formData.value.uid) {
    NotifyError(t('sync.dialog.syncError'), t('sync.dialog.uidRequired'))
    return
  }
  if (!formData.value.gameDataDir) {
    NotifyError(t('sync.dialog.syncError'), t('sync.dialog.dirRequired'))
    return
  }
 
  const elLoading = ElLoading.service({
    lock: true,
    text: t('sync.loading'),
    background: 'rgba(0, 0, 0, 0.7)',
  })
 
  loading.value = true
  
  try {
    const syncResult = await SyncRecords(parseInt(formData.value.uid), formData.value.gameDataDir)
    handleSyncResult(syncResult)
    
    // 处理用户切换逻辑
    if (syncResult.Uid === userStore.userId) {
      // 相同用户，只刷新池子信息
      await poolStore.updatePoolInfo()
    } else {
      // 不同用户，切换用户并刷新用户列表
      await userStore.updateUserList()
      // 找到对应用户并切换
      const targetUser = userStore.userList.find(u => u.uid === syncResult.Uid)
      if (targetUser) {
        await userStore.updateUserId(targetUser.id)
      }
    }
    
    dialogVisible.value = false
  } catch (err) {
    NotifyError(t('sync.dialog.syncError'), err)
  } finally {
    loading.value = false
    elLoading.close()
  }
}
</script>
 
<template>
  <div v-if="dialogVisible" class="fixed inset-0 z-50 flex items-center justify-center">
    <!-- 遮罩层 -->
    <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="!loading && (dialogVisible = false)"></div>
    
    <!-- 弹窗主体 -->
    <div class="relative bg-gray-800 rounded-xl shadow-2xl w-full max-w-md mx-4 overflow-hidden">
      <!-- 头部 -->
      <div class="bg-gradient-to-r from-blue-600 to-purple-600 px-6 py-4">
        <h2 class="text-white text-lg font-semibold">{{ $t('sync.dialog.title') }}</h2>
      </div>
      
      <!-- 内容区 -->
      <div class="p-6 space-y-5">
        <!-- UID输入 -->
        <div>
          <label class="block text-gray-300 text-sm font-medium mb-2">{{ $t('sync.dialog.uid') }}</label>
          <input 
            v-model="formData.uid" 
            type="text"
            inputmode="numeric"
            pattern="[0-9]*"
            :placeholder="$t('sync.dialog.uidPlaceholder')"
            class="w-full px-4 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
            :disabled="loading"
          />
        </div>
        
        <!-- 文件路径选择 -->
        <div>
          <label class="block text-gray-300 text-sm font-medium mb-2">{{ $t('sync.dialog.gameDataDir') }}</label>
          <div class="flex gap-2">
            <input 
              v-model="formData.gameDataDir" 
              type="text"
              readonly
              :placeholder="$t('sync.dialog.gameDataDirPlaceholder')"
              class="flex-1 px-4 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none transition-colors cursor-pointer overflow-hidden text-ellipsis whitespace-nowrap"
              @click="!loading && selectGameDataDir()"
            />
            <button 
              @click="selectGameDataDir()"
              :disabled="loading"
              class="px-4 py-3 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-600 disabled:cursor-not-allowed text-white rounded-lg transition-colors whitespace-nowrap"
            >
              {{ $t('sync.dialog.browse') }}
            </button>
          </div>
        </div>
        
        <!-- 提示信息 -->
        <div class="bg-blue-900/30 border-l-4 border-blue-500 p-4 rounded">
          <div class="flex items-start gap-3">
            <svg class="w-5 h-5 text-blue-400 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"/>
            </svg>
            <p class="text-blue-200 text-sm leading-relaxed break-words">{{ $t('sync.dialog.tips') }}</p>
          </div>
        </div>
      </div>
      
      <!-- 底部按钮 -->
      <div class="px-6 py-4 bg-gray-900/50 border-t border-gray-700 flex justify-end gap-3">
        <button 
          @click="dialogVisible = false"
          :disabled="loading"
          class="px-6 py-2.5 border border-gray-600 text-gray-300 hover:text-white hover:bg-gray-700 disabled:bg-gray-700 disabled:text-gray-500 rounded-lg transition-colors"
        >
          {{ $t('sync.dialog.cancel') }}
        </button>
        <button 
          @click="handleSubmit"
          :disabled="loading"
          class="px-6 py-2.5 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 disabled:from-gray-600 disabled:to-gray-700 text-white rounded-lg transition-all flex items-center gap-2"
        >
          <span v-if="loading" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
          <span>{{ loading ? $t('sync.dialog.syncing') : $t('sync.dialog.startSync') }}</span>
        </button>
      </div>
    </div>
  </div>
 
  <el-tooltip :content="$t('sync.button.tip')" placement="top">
    <el-button 
      type="primary" 
      @click="syncRecords"
      :disabled="loading"
    >
      {{ $t('sync.button.title') }}
    </el-button>
  </el-tooltip>
</template>
 
<style scoped>
/* 隐藏数字输入框的加减按钮 */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  display: none;
}
 
/* 确保文本正确换行 */
.break-words {
  word-break: break-word;
  overflow-wrap: break-word;
  hyphens: auto;
}
</style>