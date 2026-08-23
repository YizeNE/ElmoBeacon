<script setup lang="ts">
import PoolMain from "./PoolMain.vue";
import PoolRadio from "./PoolRadio.vue";

const poolTypeList = [1, 3, 4, 5, 6, 7, 8, 9]
</script>

<template>
  <div class="flex-1 overflow-hidden flex flex-row px-4 py-4 gap-4">
    <div class="w-60 shrink-0 flex flex-col gap-2 overflow-y-auto h-full pr-1">
      <PoolRadio v-for="poolType in poolTypeList" :poolType="poolType"/>
    </div>
    <div class="flex-1 min-w-0 overflow-hidden">
      <PoolMain/>
    </div>
  </div>
</template>