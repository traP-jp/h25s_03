<script lang="ts" setup>
import { computed, ref } from 'vue'
import { onMounted } from 'vue'
import { useTemplateRef } from 'vue'
import confetti from 'canvas-confetti'
import { watch } from 'vue'

const props = defineProps<{
  users: string[]
  winner: string | undefined
  rolling?: boolean
}>()

const users = ref<string[]>(props.users)

const count = computed(() => users.value.length)
const containerElement = useTemplateRef<HTMLDivElement>('roulette-container')
const rouletteElement = useTemplateRef<HTMLDivElement>('roulette')

watch(
  () => props.users,
  (newUsers) => {
    if (!containerElement.value || !rouletteElement.value) return
    if (rouletteElement.value.classList.contains('animate')) return
    if (rouletteElement.value.classList.contains('animate-roll')) return
    if (rouletteElement.value.classList.contains('resetting')) return
    if (props.winner !== undefined) return

    containerElement.value.style.setProperty('--count', newUsers.length.toString())
    users.value = newUsers
  },
  { immediate: true },
)

onMounted(() => {
  if (!rouletteElement.value) return

  rouletteElement.value.addEventListener('animationend', (e) => {
    if (!e.animationName.startsWith('spin')) return

    const offset = parseInt(rouletteElement.value?.style.getPropertyValue('--animate-count')!)
    const selectedIndex = (count.value - (offset % count.value)) % count.value
    const selectedElement = Array.from(
      rouletteElement.value!.querySelectorAll('.roulette-item').values(),
    )[selectedIndex]
    console.log(`Selected index: ${selectedIndex}, Element:`, selectedElement)

    selectedElement.classList.add('selected')
    confetti({
      particleCount: 300,
      spread: 70,
      origin: { y: 0.6 },
      angle: 90,
    })
  })
})

const startRoulette = () => {
  if (props.rolling) {
    rouletteElement.value?.classList.add('animate-roll')
  } else {
    rouletteElement.value?.classList.add('animate')
  }
  console.log(props.winner)

  const winnerIndex = users.value.indexOf(props.winner!)
  if (winnerIndex === -1) {
    console.error('Winner not found in users list:', props.winner)
    return
  }
  const animateCount = Math.floor(Math.random() * 10) * count.value + winnerIndex
  rouletteElement.value?.style.setProperty('--animate-count', `${animateCount}`)
}

const resetRoulette = () => {
  rouletteElement.value?.classList.remove('animate')
  rouletteElement.value?.classList.remove('animate-roll')
  document.querySelector('.roulette-item.selected')?.classList.remove('selected')

  const animateCount = parseInt(rouletteElement.value?.style.getPropertyValue('--animate-count')!)
  const resetCount =
    animateCount % count.value < count.value / 2
      ? (animateCount % count.value) + count.value
      : animateCount % count.value
  rouletteElement.value?.style.setProperty('--animate-count', `${resetCount}`)
  rouletteElement.value?.classList.add('resetting')

  rouletteElement.value?.addEventListener(
    'animationend',
    () => {
      rouletteElement.value?.classList.remove('resetting')
    },
    { once: true },
  )
}

watch(
  () => props.winner,
  (newWinner) => {
    if (newWinner) {
      startRoulette()
    } else {
      resetRoulette()
    }
  },
)
</script>

<template>
  <div ref="roulette-container" class="roulette-container">
    <div class="roulette-wrapper">
      <div ref="roulette" class="roulette">
        <div v-for="(user, i) in users" class="roulette-item" :style="`--i: ${i}`">
          <img :src="`https://q.trap.jp/api/v3/public/icon/${user}`" width="60" height="60" /><span
            >@{{ user }}</span
          >
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.roulette-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  --size: 60px;
  --r: calc(var(--size) * var(--count) / 5.5);
  --tilt: max(atan2(calc(-1 * var(--size) - 10px), calc(var(--r) * 2)), -5deg);
}

.button-container {
  display: flex;
  gap: 8px;
  z-index: 10;
}

.roulette-wrapper {
  width: 640px;
  height: 640px;
  perspective: calc(var(--r) * 2);
}

.roulette {
  pointer-events: none;
  width: 100%;
  height: 100%;
  position: relative;
  transform-style: preserve-3d;
  transform: rotateX(var(--tilt));
}

.roulette.animate {
  animation: spin 5s cubic-bezier(0.19, 1, 0.22, 1) forwards;
  --animate-count: 0;
}

.roulette.animate-roll {
  animation: spin-roll 5s cubic-bezier(0.19, 1, 0.22, 1) forwards;
  --animate-count: 0;
}

.roulette.resetting {
  animation: reset-shrink-roulette 0.6s ease forwards;
}

@keyframes reset-shrink-roulette {
  from {
    transform: rotateX(var(--tilt)) rotateY(calc(360deg / var(--count) * var(--animate-count)));
  }
  to {
    transform: rotateX(var(--tilt)) rotateY(0deg);
  }
}

.roulette.resetting .roulette-item {
  animation: reset-shrink-item 0.6s ease forwards;
}

@keyframes reset-shrink-item {
  0% {
    transform: translate(-50%, -50%) rotateY(calc(360deg * var(--i) / var(--count))) rotateZ(0)
      translateZ(var(--r)) scale(1);
    opacity: 1;
  }
  50% {
    transform: translate(-50%, -50%) rotateY(calc(360deg * var(--i) / var(--count))) rotateZ(0)
      translateZ(calc(var(--r) * 0.3)) scale(0.3);
    opacity: 0.3;
  }
  100% {
    transform: translate(-50%, -50%) rotateY(calc(360deg * var(--i) / var(--count))) rotateZ(0)
      translateZ(var(--r)) scale(1);
    opacity: 1;
  }
}

.roulette-item {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%) rotateY(calc(360deg * var(--i) / var(--count)))
    translateZ(var(--r));
  transform-origin: center center;
  transform-style: preserve-3d;
  width: var(--size);
  height: var(--size);
  background: #fff;
  border: 2px solid #333;
  border-radius: 50%;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.roulette-item img {
  width: var(--size);
  height: var(--size);
  border-radius: 50%;
}

.roulette-item span {
  font-size: 8px;
  color: #555;
  font-weight: bold;
  text-align: center;
}

.roulette-item.selected {
  z-index: 1;
  transform: translate(-50%, -50%) rotateY(calc(360deg * var(--i) / var(--count)))
    translateZ(calc(var(--r) * 1.1)) scale(1.4);
  transition: transform 0.4s ease;
  animation: glow 1.2s ease-in-out infinite;
  border-color: rgb(255, 215, 0);
}

.roulette-item.selected span {
  color: #ffd900;
}

@keyframes spin {
  to {
    transform: rotateX(var(--tilt)) rotateY(calc(360deg / var(--count) * var(--animate-count)));
  }
}

@keyframes spin-roll {
  to {
    transform: rotateX(var(--tilt)) rotateY(calc(360deg / var(--count) * var(--animate-count)))
      rotateZ(calc(360deg * 4));
  }
}

@keyframes glow {
  0%,
  100% {
    box-shadow:
      0 0 10px rgba(255, 215, 0, 0.8),
      0 0 20px rgba(255, 215, 0, 0.6);
  }
  50% {
    box-shadow:
      0 0 20px rgba(255, 255, 0, 1),
      0 0 30px rgba(255, 255, 0, 0.8);
  }
}
</style>
