<script setup lang="ts">
import {ref, nextTick, onMounted, onBeforeUnmount} from "vue";
import {useUserStore} from "../store/userStore.ts";
import {useI18n} from "vue-i18n";
import {ElMessageBox} from "element-plus";
 
const userStore = useUserStore()
const {t} = useI18n()
const menuVisible = ref(false)
const containerRef = ref<HTMLElement>()
const hoveredUserId = ref<number | null>(null)
 
const onClickOutside = (e: MouseEvent) => {
  if (containerRef.value && !containerRef.value.contains(e.target as Node)) {
    menuVisible.value = false
  }
}
 
onMounted(() => document.addEventListener('click', onClickOutside))
onBeforeUnmount(() => document.removeEventListener('click', onClickOutside))
 
const confirmDelete = (user: any) => {
  menuVisible.value = false
  nextTick(() => {
    ElMessageBox.confirm(
      t('user.deleteHint'),
      t('user.deleteTitle'),
      {
        confirmButtonText: t('user.confirmBtn'),
        cancelButtonText: t('user.cancelBtn'),
        type: 'warning',
      }
    ).then(() => {
      userStore.deleteUser(user.id, t('user.deleteSuccess'))
    }).catch(() => {})
  })
}
</script>
 
<template>
  <div ref="containerRef" class="relative inline-flex" v-if="userStore.userId && userStore.userList">
    <el-tooltip :content="$t('user.change')" placement="bottom" :disabled="menuVisible">
      <button
        class="p-1.5 rounded-full hover:bg-white/10 transition-colors flex items-center"
        @click="menuVisible = !menuVisible"
      >
        <i-mdi-person class="h-6 w-6 text-white/80 hover:text-white"/>
      </button>
    </el-tooltip>
    <div
      v-show="menuVisible"
      class="user-menu absolute left-1/2 -translate-x-1/2 top-full mt-1 py-1 bg-white rounded-lg shadow-md border border-gray-200 z-50"
    >
      <div
        v-for="user in userStore.userList"
        :key="user.id"
        class="flex items-center px-4 py-1.5 cursor-pointer text-sm whitespace-nowrap group"
        :class="[
          user.id == userStore.userId ? 'text-red-500' : 'text-gray-700',
          hoveredUserId === user.id ? 'bg-gray-100' : ''
        ]"
        @click="userStore.updateUserId(user.id); menuVisible = false"
        @mouseenter="hoveredUserId = user.id"
        @mouseleave="hoveredUserId = null"
      >
        <span>{{ $t(`server.${user.server}`) + ' ' + user.uid }}</span>
        <i-mdi-delete-outline
          class="h-3.5 w-3.5 ml-1.5 shrink-0 transition-opacity duration-150"
          :class="hoveredUserId === user.id
            ? 'opacity-100 pointer-events-auto text-gray-400 hover:text-red-500'
            : 'opacity-0 pointer-events-none'"
          @click.stop="confirmDelete(user)"
        />
      </div>
    </div>
  </div>
</template>

<style>
.user-menu::before {
  content: '';
  position: absolute;
  top: -5px;
  left: 50%;
  transform: translateX(-50%);
  border-left: 5px solid transparent;
  border-right: 5px solid transparent;
  border-bottom: 5px solid white;
}
.user-menu::after {
  content: '';
  position: absolute;
  top: -6px;
  left: 50%;
  transform: translateX(-50%);
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-bottom: 6px solid #e5e7eb;
}
</style>