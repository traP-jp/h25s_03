<script setup lang="ts">
defineProps<{
  id: string
  size?: string | number
}>()
</script>

<template>
  <v-avatar :image="`https://q.trap.jp/api/v3/public/icon/${id}`" :size="size || 'default'" />
</template>
