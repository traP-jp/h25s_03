<script setup lang="ts">
import UserAvatar from '@/components/UserAvatar.vue'
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { apiClient } from '@/api/apiClient'
import type { components } from '@/api/schema'

const route = useRoute()
const eventId = route.params.eventId as string

const event = ref<components['schemas']['Event']>()
const lotteries = ref<components['schemas']['Lottery'][]>()
const newLotteryTitle = ref<string>('')
const isDialogActive = ref(false)

const isDense = computed(() => {
  return event.value && event.value.attendees.length > 40
})

onMounted(async () => {
  await fetchEvent()
  await fetchLotteries()
})

const fetchEvent = async () => {
  event.value = (
    await apiClient.GET('/events/{eventID}', {
      params: { path: { eventID: eventId } },
    })
  ).data
}
const fetchLotteries = async () => {
  lotteries.value = (
    await apiClient.GET('/events/{eventID}/lotteries', {
      params: { path: { eventID: eventId }, query: { ifDeleted: false } },
    })
  ).data
}
const postLottery = async () => {
  if (!newLotteryTitle.value) return
  await apiClient.POST('/events/{eventID}/lotteries', {
    params: { path: { eventID: eventId } },
    body: { title: newLotteryTitle.value },
  })
  await fetchLotteries()
  newLotteryTitle.value = ''
  isDialogActive.value = false
}
</script>

<template>
  <v-container max-width="600">
    <v-sheet v-if="event && lotteries" elevation="2" class="px-8 py-4">
      <div class="d-flex justify-space-between align-end">
        <h1>{{ event.title }}</h1>
        <div class="d-flex ga-1">
          <span>by</span>
          <user-avatar v-for="admin in event.admins" :key="admin" :id="admin" size="x-small" />
        </div>
      </div>
      <span class="text-body-2 text-medium-emphasis">{{ event.date }}</span>
      <p class="text-body-1">{{ event.description }}</p>
      <div class="py-2">
        <div class="d-flex align-center ga-1">
          <h2 class="text-overline">attendees</h2>
          <span class="text-caption text-medium-emphasis">({{ event.attendees.length }})</span>
        </div>
        <div class="d-flex align-end flex-wrap mx-2" :class="{ 'ga-1': !isDense }">
          <user-avatar
            v-for="attendee in event.attendees"
            :key="attendee"
            :id="attendee"
            :size="isDense ? 'x-small' : 'small'"
          />
        </div>
      </div>
      <div class="py-2">
        <h2 class="text-overline">lottery</h2>
        <div class="">
          <v-list class="mx-0">
            <v-list-item v-for="lottery in lotteries" :key="lottery.lottery_id">
              <div class="d-flex align-end ga-4">
                <router-link
                  :to="{
                    name: 'Lottery',
                    params: { eventId: eventId, lotteryId: lottery.lottery_id },
                  }"
                >
                  <v-list-item-title class="">{{ lottery.title }}</v-list-item-title>
                </router-link>
                <div class="d-flex ga-1">
                  <user-avatar
                    v-for="winner in lottery.winners"
                    :key="winner"
                    :id="winner"
                    size="x-small"
                  />
                </div>
              </div>
            </v-list-item>
          </v-list>
          <v-dialog v-model="isDialogActive" max-width="500">
            <template v-slot:activator="{ props: activatorProps }">
              <v-btn
                v-bind="activatorProps"
                color="primary"
                density="comfortable"
                text="新規作成"
              />
            </template>
            <template v-slot:default>
              <v-card class="pa-2">
                <v-card-title>抽選の作成</v-card-title>
                <v-text-field
                  v-model="newLotteryTitle"
                  label="タイトル"
                  density="comfortable"
                  outlined
                  hide-details
                  class="px-4"
                />
                <v-card-actions>
                  <v-btn
                    density="comfortable"
                    text="作成"
                    :disabled="!newLotteryTitle"
                    @click="postLottery"
                  />
                </v-card-actions>
              </v-card>
            </template>
          </v-dialog>
        </div>
      </div>
    </v-sheet>
  </v-container>
</template>
