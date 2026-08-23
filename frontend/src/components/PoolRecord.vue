<script setup lang="ts">
import { usePoolStore } from '../store/poolStore'


const poolStore = usePoolStore()

//1x1透明图片
const TRANSPARENT = 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg=='

const props = defineProps({
  name: { type: String, required: true },
  iconSrc: { type: String, required: true },
  count: { type: Number, required: true },
  timestamp: { type: Number, required: true },
  isMissing: { type: Boolean, required: true },
  rank: { type: Number, required: true }
})


const getBgColor = () => {
  //常驻池|角色池  武器池  新手池  神秘箱  皮肤池
  if (poolStore.poolType == 1 || poolStore.poolType == 3 || poolStore.poolType == 6) {
    if (props.count <= 58) {
      return 'bg-gradient-to-br from-emerald-500 to-green-700 border border-emerald-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else if (props.count < 66) {
      return 'bg-gradient-to-br from-amber-500 to-orange-700 border border-amber-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else {
      return 'bg-gradient-to-br from-rose-500 to-red-700 border border-rose-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    }
  } else if (poolStore.poolType == 4 || poolStore.poolType == 7) {
    if (props.count <= 50) {
      return 'bg-gradient-to-br from-emerald-500 to-green-700 border border-emerald-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else if (props.count < 58) {
      return 'bg-gradient-to-br from-amber-500 to-orange-700 border border-amber-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else {
      return 'bg-gradient-to-br from-rose-500 to-red-700 border border-rose-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    }
  } else if (poolStore.poolType == 5) {
    if (props.count < 50) {
      return 'bg-gradient-to-br from-emerald-500 to-green-700 border border-emerald-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else {
      return 'bg-gradient-to-br from-rose-500 to-red-700 border border-rose-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    }
  } else if (poolStore.poolType == 8) {
    if (props.count <= 100) {
      return 'bg-gradient-to-br from-emerald-500 to-green-700 border border-emerald-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else if (props.count <= 400) {
      return 'bg-gradient-to-br from-amber-500 to-orange-700 border border-amber-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    } else {
      return 'bg-gradient-to-br from-rose-500 to-red-700 border border-rose-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
    }
  } else if (poolStore.poolType == 9) {
    if (props.rank == 5) {
      if (props.count <= 50) {
        return 'bg-gradient-to-br from-emerald-500 to-green-700 border border-emerald-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
      } else if (props.count < 90) {
        return 'bg-gradient-to-br from-amber-500 to-orange-700 border border-amber-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
      } else {
        return 'bg-gradient-to-br from-rose-500 to-red-700 border border-rose-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
      }
    } else if (props.rank == 4) {
      if (props.count <= 30) {
        return 'bg-gradient-to-br from-emerald-500 to-green-700 border border-emerald-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
      } else if (props.count < 65) {
        return 'bg-gradient-to-br from-amber-500 to-orange-700 border border-amber-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
      } else {
        return 'bg-gradient-to-br from-rose-500 to-red-700 border border-rose-300/50 shadow-[inset_0_1px_0_rgba(255,255,255,0.3)]'
      }
    }
  }
}
</script>

<template>
  <div
    :class="['w-[202px] h-12 relative shadow-xl rounded-md shrink-0 select-none flex justify-center items-center text-white', getBgColor()]">
    <img :src="iconSrc || TRANSPARENT"
      class="absolute left-1 h-10 w-10 rounded object-cover drop-shadow-[0_2px_4px_rgba(0,0,0,0.8)]" />
    <el-tooltip effect="dark" :content="new Date(timestamp * 1000).toLocaleString()" placement="top">
      <div class="relative truncate pl-10">
        {{ `${name}「${count}」` }}
      </div>
    </el-tooltip>
  </div>
</template>