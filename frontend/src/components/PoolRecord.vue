<script setup lang="ts">
import { GetIcon } from '../../wailsjs/go/handler/App'
import { ref, watch } from 'vue'
import { usePoolStore } from '../store/poolStore'


const poolStore = usePoolStore()

const props = defineProps({
  name: {type: String, required: true},
  icon: {type: String, required: true},
  count: {type: Number, required: true},
  timestamp: {type: Number, required: true},
  isMissing: {type: Boolean, required: true},
})

const iconSrc = ref('')
 
const loadIcon = async () => {
  iconSrc.value = ''
  if (props.icon) {
    try {
      iconSrc.value = await GetIcon(props.icon)
    } catch {
      iconSrc.value = ''
    }
  }
}
 
watch(() => props.icon, loadIcon, { immediate: true })

const getBgColor = () => {
  //80抽保底|70抽保底|50抽保底|无保底
  if (poolStore.poolType==1||poolStore.poolType==3||poolStore.poolType==6){
      if (props.count <= 58) {
        return 'bg-green-500'
      } else if (props.count < 66) {
        return 'bg-amber-500'
      } else {
        return 'bg-red-600'
      }
  }else if(poolStore.poolType==4||poolStore.poolType==5||poolStore.poolType==7){
      if (props.count <= 50) {
        return 'bg-green-500'
      } else if (props.count < 58) {
        return 'bg-amber-500'
      } else {
        return 'bg-red-600'
      }
  }else if(poolStore.poolType==5){
      if (props.count < 50) {
        return 'bg-green-500'
      } else{
        return 'bg-red-600'
      }
  }else{
    if (props.count<=100){
      return 'bg-green-500'
    }else if(props.count<=400){
      return 'bg-amber-500'
    }else{
      return 'bg-red-600'
    }
  }
}
</script>

<template>
  <div :class="['w-52 h-12 relative shadow-xl rounded-md shrink-0 select-none flex justify-center items-center text-white',getBgColor()]">
    <img v-if="iconSrc" :src="iconSrc" class="absolute left-1 h-10 w-10 rounded object-cover" />
    <el-tooltip effect="dark" :content="new Date(timestamp*1000).toLocaleString()" placement="top">
      <div class="relative truncate" :class="iconSrc ? 'pl-10' : ''">
        {{ `${name}「${count}」` }}
      </div>
    </el-tooltip>
  </div>
</template>