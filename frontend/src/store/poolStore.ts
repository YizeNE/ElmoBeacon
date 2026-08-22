import {defineStore} from "pinia";
import {ref} from "vue";
import {handler} from "../../wailsjs/go/models.ts";
import {GetPoolInfo} from "../../wailsjs/go/handler/App";
import {useUserStore} from "./userStore.ts";
import {NotifyError} from "../utils/notify.ts";
import PoolInfo = handler.PoolInfo;

export const usePoolStore = defineStore('pool', () => {
    const poolType = ref(1)
    const poolInfo = ref<PoolInfo>()
    const loading = ref(false)         
    let fetchController: AbortController | null = null 

    const userStore = useUserStore()

    const updatePoolInfo = async () => {
        if (!userStore.userId) return
        // 取消上一次未完成的请求
        fetchController?.abort()
        fetchController = new AbortController()
 
        loading.value = true    
        try {
            const res = await GetPoolInfo(userStore.userId, poolType.value)
            // 确保不是被 abort 掉的旧请求
            if (!fetchController.signal.aborted) {
                poolInfo.value = res
            }
        } catch (err) {
            NotifyError('Error', err)
        } finally {
            loading.value = false    
        }
    }

    const init = async () => {
        await updatePoolInfo()
    }

    return {poolType, poolInfo, loading, updatePoolInfo, init}
})