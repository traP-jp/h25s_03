<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { apiClient } from '@/api/apiClient'
import type { components } from '@/api/schema'
const ifDUplicated = ref(false)
const route = useRoute()
const eventId = route.params.eventId as string
const lotteryId = route.params.lotteryId as string
const lottery = ref<components['schemas']['Lottery'] | undefined>(undefined)
const postLottery = async () => {
  await apiClient.POST('/events/{eventID}/lotteries/{lotteryID}', {
    params: {
      path: { eventID: eventId, lotteryID: lotteryId },
      query: { ifDuplicated: ifDUplicated.value }
    },
  })
  await fetchLottery()
}

const fetchLottery = async () => {
  const response = await apiClient.GET('/events/{eventID}/lotteries/{lotteryID}', {
    params: { path: { eventID: eventId, lotteryID: lotteryId } }
  })
  lottery.value = response.data
}

onMounted(fetchLottery)
</script>
<template>
  <v-container max-width="600">
    <v-card class="d-flex flex-column align-center ga-5">
      <h1>
        {{ lottery?.title }}
      </h1>
      <v-btn @click="postLottery"> 抽選スタート！！ </v-btn>
      <v-checkbox
      v-model="ifDUplicated"
      :label="`イベント内の重複当選を許可する`"
    ></v-checkbox>
      <div class="d-flex">
       
      </div>
      <div v-if="lottery !== undefined" class="d-flex ga-2">
        <div>当選者 :</div>
        <div>
          <div v-for="winner in lottery.winners" key="winner_id">
            {{ winner }}

            <v-avatar :image="`https://q.trap.jp/api/v3/public/icon/${winner}`" size="25"></v-avatar>

          </div>
        </div>
      </div>
      <v-card-actions>
        <v-btn :to="{ name: 'EventDetail', params: { eventId: 'hoge' } }"> 戻る </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
<style scoped></style>
