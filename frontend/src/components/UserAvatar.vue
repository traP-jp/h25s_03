<script setup lang="ts">
defineProps<{
  id: string
  size?: string | number
}>()
</script>

<template>
  <v-tooltip location="top">
    <template v-slot:activator="{ props }">
      <v-avatar
        v-bind="props"
        :image="`https://q.trap.jp/api/v3/public/icon/${id}`"
        :size="size || 'default'"
      />
    </template>
    <span class="text-caption">{{ id }}</span>
  </v-tooltip>
</template>
