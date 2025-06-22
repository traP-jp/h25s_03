<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { apiClient } from '@/api/apiClient'
import type { components } from '@/api/schema'
import LotteryRoulette from '@/components/LotteryRoulette.vue'

const ifDuplicated = ref(false)
const route = useRoute()
const eventId = route.params.eventId as string
const lotteryId = route.params.lotteryId as string
const lottery = ref<components['schemas']['Lottery'] | undefined>(undefined)
const event = ref<components['schemas']['Event'] | undefined>(undefined)
const winner = ref<string | undefined>(undefined)

const reset = () => {
  winner.value = undefined
}

const postLottery = async () => {
  await fetchLottery()
  await fetchEvent()

  const res = await apiClient.POST('/events/{eventID}/lotteries/{lotteryID}', {
    params: {
      path: { eventID: eventId, lotteryID: lotteryId },
      query: { ifDuplicated: ifDuplicated.value },
    },
  })
  if (res.data === undefined) {
    alert('抽選に失敗しました。')
    return
  }
  winner.value = res.data.winner
}
const fetchLottery = async () => {
  const response = await apiClient.GET('/events/{eventID}/lotteries/{lotteryID}', {
    params: { path: { eventID: eventId, lotteryID: lotteryId } },
  })
  if (response.data === undefined) return
  lottery.value = response.data
}
const fetchEvent = async () => {
  const response = await apiClient.GET('/events/{eventID}', {
    params: { path: { eventID: eventId } },
  })
  if (response.data === undefined) return
  event.value = response.data
}

onMounted(() => {
  fetchLottery()
  fetchEvent()

  const interval = setInterval(() => {
    fetchEvent()
  }, 1000)

  return () => {
    clearInterval(interval)
  }
})
</script>

<template>
  <v-container>
    <v-card class="d-flex flex-column align-center ga-5">
      <h1>
        {{ lottery?.title }}
      </h1>
      <div class="d-flex ga-4">
        <v-btn @click="postLottery" :disabled="winner !== undefined"> 抽選スタート！！ </v-btn>
        <v-btn @click="reset" :disabled="winner === undefined">リセット</v-btn>
      </div>
      <LotteryRoulette v-if="lottery" :users="event?.attendees ?? []" :winner="winner" />
      <v-checkbox v-model="ifDuplicated" label="イベント内の重複当選を許可する"></v-checkbox>
      <div v-if="lottery" class="d-flex ga-2">
        <span>当選者 :</span>
        <div>
          <div v-for="winner in lottery.winners" :key="winner">
            {{ winner }}

            <v-avatar
              :image="`https://q.trap.jp/api/v3/public/icon/${winner}`"
              size="25"
            ></v-avatar>
          </div>
        </div>
      </div>
      <v-card-actions>
        <v-btn :to="{ name: 'EventDetail', params: { eventId: eventId } }"> 戻る </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
