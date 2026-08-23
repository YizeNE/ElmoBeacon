<script setup lang="ts">
import { WindowIsMaximised, WindowMaximise, WindowUnmaximise } from "../../wailsjs/runtime";
import ButtonVersion from "./ButtonVersion.vue";
import ButtonWindowMinimise from "./ButtonWindowMinimise.vue";
import ButtonWindowMaxmise from "./ButtonWindowMaxmise.vue";
import ButtonWindowClose from "./ButtonWindowClose.vue";
import SelectorLanguage from "./SelectorLanguage.vue";
import SelectorUser from "./SelectorUser.vue";
import ButtonColorPalette from "./ButtonColorPalette.vue"; // 新增导入
import ButtonSyncRecords from "./ButtonSyncRecords.vue";
import ButtonGithub from "./ButtonGithub.vue";

const handleWindowMaximise = async () => {
  await WindowIsMaximised().then(res => res ? WindowUnmaximise() : WindowMaximise())
}
</script>

<template>
  <div style="--wails-draggable:drag"
    class="h-16 shrink-0 px-8 text-lg text-stone-400 flex flex-row justify-between items-center gap-2"
    @dblclick.self="handleWindowMaximise">
    <!-- GitHub图标 -->
    <div class="flex flex-row items-center gap-2 mr-auto">
      <span style="--wails-draggable:no-drag" class="flex items-center">
        <ButtonGithub />
      </span>
      
      <!-- 标题文字 -->
      <span class="text-2xl font-bold select-none bg-gradient-to-l from-orange-500 to-slate-50 bg-clip-text text-transparent">
        {{ $t('window.title') }}
      </span>
      
      <!-- 检查更新按钮 -->
      <span style="--wails-draggable:no-drag" class="flex items-center ml-1">
        <ButtonVersion />
      </span>
    </div>

    <div style="--wails-draggable:no-drag" class="flex flex-row items-center gap-1">
      <SelectorUser />
      <ButtonSyncRecords />
      <ButtonColorPalette />
      <SelectorLanguage />

      <!-- 一条细微的分割线 -->
      <div class="w-px h-5 bg-white/20 mx-2"></div>

      <ButtonWindowMinimise />
      <ButtonWindowMaxmise />
      <ButtonWindowClose />
    </div>
  </div>
</template>