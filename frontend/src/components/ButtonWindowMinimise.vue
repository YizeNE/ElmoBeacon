<script setup lang="ts">
import {WindowMinimise} from "../../wailsjs/runtime";
</script>

<template>
  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 p-1 rounded-full hover:bg-white/10 text-white/80 hover:text-white cursor-pointer transition-colors" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" @click="WindowMinimise">
    <path stroke-linecap="round" stroke-linejoin="round" d="M5 12h14" />
  </svg>
</template>