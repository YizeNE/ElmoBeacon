<script setup lang="ts">
import {usePoolStore} from "../store/poolStore.ts";

const props = defineProps({poolType: {type: Number, required: true}})
const poolStore = usePoolStore()

const handleClick = () => {
  poolStore.poolType = props.poolType
}
</script>

<template>
  <div 
    :class="[
      'h-24 p-4 rounded-xl text-lg select-none cursor-pointer transition-all duration-300 flex flex-row items-center border backdrop-blur-md',
      poolStore.poolType === poolType 
        ? 'bg-white/20 text-white border-l-4 shadow-lg' 
        : 'bg-white/5 text-gray-300 hover:bg-white/10 hover:text-white border-l-4 border-transparent'
    ]" 
    :style="poolStore.poolType === poolType ? { borderColor: `var(--theme-to)`, boxShadow: `0 0 10px var(--theme-to)` } : {}"
    @click="handleClick">
    {{ $t(`gacha.type.${poolType}`) }}
  </div>
</template>