import { createRouter, createWebHistory } from 'vue-router'
import EventsListView from '@/views/EventsListView.vue'
import EventDetail from '@/views/EventDetailView.vue'
import LotteryView from '@/views/LotteryView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'EventsList',
      component: EventsListView,
    },
    {
      path: '/:eventId',
      name: 'EventDetail',
      component: EventDetail,
    },
    {
      path: '/:eventId/:lotteryId',
      name: 'Lottery',
      component: LotteryView,
    },
  ],
})

export default router
