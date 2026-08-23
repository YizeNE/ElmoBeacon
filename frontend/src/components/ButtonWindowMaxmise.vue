<script setup lang="ts">
import {WindowIsMaximised, WindowMaximise, WindowUnmaximise} from "../../wailsjs/runtime";

const handleWindowMaximise = async () => {
  await WindowIsMaximised().then(res => res ? WindowUnmaximise() : WindowMaximise())
}
</script>

<template>
  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 p-1 rounded-full hover:bg-white/10 text-white/80 hover:text-white cursor-pointer transition-colors" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" @click="handleWindowMaximise">
    <path stroke-linecap="round" stroke-linejoin="round" d="M5 5h14v14H5z" />
  </svg>
</template>