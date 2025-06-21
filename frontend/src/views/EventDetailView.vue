<script setup lang="ts">
import UserAvatar from '@/components/UserAvatar.vue'
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { apiClient } from '@/api/apiClient'
import type { components } from '@/api/schema'

const route = useRoute()
const eventId = route.params.eventId as string

const event = ref<components['schemas']['Event'] | undefined>()

const avatarSize = computed(() =>
  event.value && event.value.attendees.length >= 50 ? 'x-small' : 'small',
)

onMounted(async () => {
  await getEvent()
})

const getEvent = async () => {
  event.value = (
    await apiClient.GET('/events/{eventID}', {
      params: { path: { eventID: eventId } },
    })
  ).data
}
</script>

<template>
  <v-container max-width="600">
    <v-card v-if="event" class="pa-4">
      <div class="d-flex justify-space-between align-end">
        <h1>{{ event.title }}</h1>
        <div>
          <span>by</span>
          <user-avatar v-for="admin in event.admins" :key="admin" :id="admin" size="x-small" />
        </div>
      </div>
      <span>{{ event.date }}</span>
      <p>{{ event.description }}</p>
      <div>
        <h2>参加者リスト</h2>
        <div class="d-flex flex-wrap">
          <v-avatar
            v-for="attendee in event.attendees"
            :key="attendee"
            :image="`https://q.trap.jp/api/v3/public/icon/${attendee}`"
            :size="avatarSize"
          />
        </div>
      </div>
      <div>
        <h2>抽選リスト</h2>
        <v-list></v-list>
        <v-btn color="primary">抽選の作成</v-btn>
      </div>
    </v-card>
  </v-container>
</template>
