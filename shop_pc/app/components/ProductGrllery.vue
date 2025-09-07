<script setup lang="ts">
const props = withDefaults(defineProps<{
  grllery: any[]
}>(), {
  grllery: () => [] 
})

const current = useState('current', () => 0)

import { ref } from 'vue'

const bigPictureRef = ref<any>(null)
const zoomRef = ref<any>(null)
const zoomShowStyle = ref<any>({})
const zoomVisible = ref(false)
const zoomStyle = ref({})

const handleMouseEnter = (event: MouseEvent) => {
  const rect = bigPictureRef.value.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  const zoomX = Math.max(0, Math.min(x - 100, rect.width - 401)) // Ensure zoom box stays within the image bounds
  const zoomY = Math.max(0, Math.min(y - 100, rect.height -401)) // Ensure zoom box stays within the image bounds


  zoomStyle.value = {
    left: `${zoomX}px`,
    top: `${zoomY}px`,
  }

  zoomShowStyle.value = {
    backgroundImage: `url(${props.grllery[current.value]})`,
    backgroundPosition: `-${zoomX}px -${-zoomY}px`
  }

  zoomVisible.value = true
}

const handleMouseMove = (event: MouseEvent) => {
  const rect = bigPictureRef.value.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  const zoomX = Math.max(0, Math.min(x - 100, rect.width - 401)) // Ensure zoom box stays within the image bounds
  const zoomY = Math.max(0, Math.min(y - 100, rect.height - 401)) // Ensure zoom box stays within the image bounds

  zoomStyle.value = {
    left: `${zoomX}px`,
    top: `${zoomY}px`,
  }
  zoomShowStyle.value = {
    backgroundImage: `url(${props.grllery[current.value]})`,
    backgroundPosition: `-${zoomX}px -${zoomY}px`
  }
}

const handleMouseLeave = () => {
  zoomVisible.value = false
}
</script>

<template>
  <div class="product-grllery">
    <div class="grllery">
      <div :class="{'grllery-item': true, 'checked': index === current}" v-for="(item, index) in props.grllery" :key="index" @click="current = index">
        <img :src="item" class="thumb" alt=""></img>
      </div>
    </div>
    <div class="big-picture">
      <img ref="bigPictureRef" @mouseenter="handleMouseEnter"
      @mousemove="handleMouseMove"
      @mouseleave="handleMouseLeave" :src="props.grllery[current]" class="thumb" alt=""></img>
      <div v-if="zoomVisible" class="zoom-box" :style="zoomStyle" ref="zoomRef"></div>
      <div class="show_zoom" v-if="zoomVisible" :style="zoomShowStyle">
        
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
.product-grllery {
  width: 100%;
  display: flex;
  gap: 5px;
}
.big-picture {
  border-radius: 8px;
  background-color: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 502px;
  height: 502px;
  position: relative; 
}
.big-picture img{
  width: 502px; 
  height: 502px;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
.grllery {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}
.grllery-item {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 94px;
  height: 94px;
  border: 1px solid #f5f5f5;
  background-color: #f5f5f5;
  border-radius: 5px;
  cursor: pointer;
  transition: all 0.3s ease;
}
.grllery-item:hover {
  border-color: var(--primary-color);
}
.grllery-item.checked {
  border-color: var(--primary-color);
}

.grllery img {
  max-width: 90%;
  max-height: 90%;
  display: block;
}

.zoom-box {
  position: absolute;
  width: 401px;
  height: 401px;
  background-repeat: no-repeat;
  border: 1px solid #ccc;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  background-color: rgba(77, 77, 77, 0.295);
}
.show_zoom {
  position: absolute;
  width: 502px;
  height: 502px;
  top: 0;
  left: 512px;
  z-index: 20;
  transition: all 0.2s;
  background-repeat: no-repeat;
  background-size: 654px 654px;
}
</style>