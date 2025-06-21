import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { vuetify } from './plugins/vuetify'

const app = createApp(App)
const pinia = createPinia()

app.use(router)
app.use(pinia)
app.use(vuetify)

// MSW
if (import.meta.env.DEV) {
  const { worker } = await import('./api/mocks/browser')
  await worker.start()
}

app.mount('#app')
