<script setup lang="ts">
import { usePoolStore } from "../store/poolStore.ts";
import PoolStatistic from "./PoolStatistic.vue";
import PoolRecord from "./PoolRecord.vue";
import { ref, watch } from "vue";
import { GetIcon } from "../../wailsjs/go/handler/App";

const poolStore = usePoolStore()
const iconMap = ref<Record<string, string>>({})
const iconLoading = ref(false)
const ICON_LOAD_TIMEOUT = 15000  // 15秒兜底

watch(() => poolStore.poolInfo, async (info) => {
  if (!info?.recordList) {
    iconMap.value = {}
    return
  }
  iconLoading.value = true
  const names = [...new Set(info.recordList.map(r => r.Icon).filter(Boolean))]

  const timeout = new Promise<void>((_, reject) =>
    setTimeout(() => reject(new Error('icon load timeout')), ICON_LOAD_TIMEOUT)
  )

  try {
    const load = Promise.all(
      names.map(name =>
        GetIcon(name)
          .then(src => [name, src] as const)
          .catch(() => [name, ''] as const)
      )
    ).then(results => {
      iconMap.value = Object.fromEntries(results)
    })

    await Promise.race([load, timeout])  // 兜底时间结束后即使图片没有全部加载完也强制显示
  } catch(err) {
    console.warn('icon load timeout:', err)
  } finally {
    iconLoading.value = false            // 无论如何都结束 loading
  }
}, { immediate: true })

</script>

<template>
  <div class="w-full h-full px-2 flex flex-col gap-4">
    <template v-if="poolStore.loading || iconLoading">
      <div class="w-full flex-1 flex flex-col items-center justify-center gap-4">
        <div class="w-12 h-12 border-4 border-white/30 border-t-blue-400 rounded-full animate-spin"></div>
        <span class="text-white/60 text-lg">{{ $t('gacha.loading') }}</span>
      </div>
    </template>

    <template v-else-if="poolStore.poolInfo">
      <div class="w-full rounded-2xl p-4 flex flex-col gap-4 shadow-2xl backdrop-blur-xl border border-white/10 h-full overflow-hidden"
       :style="{ background: 'rgba(255, 255, 255, 0.05)' }">

        <!-- 统计区 -->
        <div class="flex flex-wrap  gap-3 shrink-0">
          <PoolStatistic class="text-gray-200" :title="$t('gacha.statistic.totalCount')"
            :value="poolStore.poolInfo.totalCount" />
          <PoolStatistic v-if="poolStore.poolType != 9" class="text-orange-300" :title="$t('gacha.statistic.pityCount')"
            :value="poolStore.poolInfo.storedCount" />
          <PoolStatistic class="text-yellow-200" :title="$t('gacha.statistic.rank5Data')"
            :value="poolStore.poolInfo.rank5Count" :note="(poolStore.poolInfo.rank5Rate * 100).toFixed(2) + '%'" />
          <PoolStatistic class="text-purple-300" :title="$t('gacha.statistic.rank4Data')"
            :value="poolStore.poolInfo.rank4Count" :note="(poolStore.poolInfo.rank4Rate * 100).toFixed(2) + '%'" />
          <PoolStatistic v-if="poolStore.poolType != 9" class="text-blue-300" :title="$t('gacha.statistic.rank3Data')"
            :value="poolStore.poolInfo.rank3Count" :note="(poolStore.poolInfo.rank3Rate * 100).toFixed(2) + '%'" />
          <PoolStatistic class="text-green-300" :title="$t('gacha.statistic.rank5Avg')"
            :value="poolStore.poolInfo.rank5Avg" />
          <PoolStatistic v-if="poolStore.poolType == 3 || poolStore.poolType == 4" class="text-pink-300"
            :title="$t('gacha.statistic.upRank5Avg')" :value="poolStore.poolInfo.rank5UpAvg" />
          <PoolStatistic v-if="poolStore.poolType == 3 || poolStore.poolType == 4" class="text-red-300"
            :title="$t('gacha.statistic.nonUpRate')"
            :note="Math.round(poolStore.poolInfo.missingRate * 1000) / 10 + '%'" />
        </div>

         <div class="h-12 w-full shadow-md select-none text-white text-center rounded-md flex flex-row gap-2 justify-center items-center backdrop-blur-sm shrink-0 border border-white/10"
         :style="{ background: 'rgba(255, 255, 255, 0.08)' }">
          <span>{{ $t('gacha.records.title') }}</span>
          <el-tooltip effect="light" :content="$t('gacha.records.tip')" placement="top">
            <i-mdi-help class="text-sm text-slate-400 hover:text-white" />
            <template #content>
              <div class="text-lg max-w-96">{{ $t('gacha.records.tip') }}</div>
            </template>
          </el-tooltip>
        </div>

        <!-- 记录区 -->
        <el-scrollbar class="flex-1 min-h-0">
          <div class="w-full flex gap-x-2 gap-y-2 flex-wrap pb-2">
            <PoolRecord v-for="record in poolStore.poolInfo.recordList" :name="record.Name"
              :icon-src="iconMap[record.Icon] || ''" :count="record.Count" :timestamp="record.Timestamp"
              :is-missing="record.IsMissing" :rank="record.Rank" />
          </div>
        </el-scrollbar>

      </div>
    </template>
  </div>
</template>