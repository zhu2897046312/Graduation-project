<script setup lang="ts">
const props = withDefaults(defineProps<{
  grllery: string[]
}>(), {
  grllery: () => [] 
})

const current = ref(0)

const bigPictureRef = ref<HTMLImageElement | null>(null)
const zoomShowStyle = ref<Record<string, string>>({})
const zoomVisible = ref(false)
const zoomStyle = ref<Record<string, string>>({})

// 当前显示的图片
const currentImage = computed(() => {
  if (!props.grllery || props.grllery.length === 0) return ''
  return props.grllery[current.value] || props.grllery[0] || ''
})

// 处理鼠标进入
const handleMouseEnter = (event: MouseEvent) => {
  if (!bigPictureRef.value || !currentImage.value) return
  
  const rect = bigPictureRef.value.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  const zoomSize = 200
  const zoomX = Math.max(0, Math.min(x - zoomSize / 2, rect.width - zoomSize))
  const zoomY = Math.max(0, Math.min(y - zoomSize / 2, rect.height - zoomSize))

  zoomStyle.value = {
    left: `${zoomX}px`,
    top: `${zoomY}px`,
  }

  zoomShowStyle.value = {
    backgroundImage: `url(${currentImage.value})`,
    backgroundPosition: `-${zoomX * 2}px -${zoomY * 2}px`
  }

  zoomVisible.value = true
}

// 处理鼠标移动
const handleMouseMove = (event: MouseEvent) => {
  if (!bigPictureRef.value || !zoomVisible.value || !currentImage.value) return
  
  const rect = bigPictureRef.value.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  const zoomSize = 200
  const zoomX = Math.max(0, Math.min(x - zoomSize / 2, rect.width - zoomSize))
  const zoomY = Math.max(0, Math.min(y - zoomSize / 2, rect.height - zoomSize))

  zoomStyle.value = {
    left: `${zoomX}px`,
    top: `${zoomY}px`,
  }
  
  zoomShowStyle.value = {
    backgroundImage: `url(${currentImage.value})`,
    backgroundPosition: `-${zoomX * 2}px -${zoomY * 2}px`
  }
}

// 处理鼠标离开
const handleMouseLeave = () => {
  zoomVisible.value = false
}

// 切换图片
const selectImage = (index: number) => {
  if (index >= 0 && index < props.grllery.length) {
    current.value = index
    zoomVisible.value = false
  }
}
</script>

<template>
  <div v-if="props.grllery && props.grllery.length > 0" class="product-gallery flex gap-4">
    <!-- 缩略图列表 -->
    <div class="flex flex-col gap-2 flex-shrink-0">
      <button
        v-for="(item, index) in props.grllery"
        :key="index"
        :class="[
          'relative w-20 h-20 rounded-lg overflow-hidden border-2 transition-all duration-200',
          index === current
            ? 'border-primary-500 dark:border-primary-400 ring-2 ring-primary-200 dark:ring-primary-800'
            : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
        ]"
        @click="selectImage(index)"
      >
        <img
          :src="item"
          :alt="`Thumbnail ${index + 1}`"
          class="w-full h-full object-cover"
          loading="lazy"
        />
        <div
          v-if="index === current"
          class="absolute inset-0 bg-primary-500/10 dark:bg-primary-400/10"
        />
      </button>
    </div>

    <!-- 主图区域 -->
    <div class="flex-1 relative">
      <div class="relative w-full aspect-square bg-gray-100 dark:bg-gray-800 rounded-lg overflow-hidden">
        <img
          ref="bigPictureRef"
          :src="currentImage"
          :alt="`Product image ${current + 1}`"
          class="w-full h-full object-contain cursor-zoom-in"
          @mouseenter="handleMouseEnter"
          @mousemove="handleMouseMove"
          @mouseleave="handleMouseLeave"
          @error="(e: any) => { e.target.src = '/placeholder-product.jpg' }"
        />
        
        <!-- 放大镜框 -->
        <div
          v-if="zoomVisible"
          class="absolute w-48 h-48 border-2 border-primary-500 dark:border-primary-400 bg-primary-500/20 dark:bg-primary-400/20 pointer-events-none rounded-lg"
          :style="zoomStyle"
        />
      </div>

      <!-- 放大预览 -->
      <div
        v-if="zoomVisible"
        class="absolute left-full ml-4 top-0 w-full aspect-square bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700 rounded-lg overflow-hidden shadow-xl z-50 hidden lg:block"
        :style="zoomShowStyle"
      />
    </div>
  </div>
  
  <!-- 空状态 -->
  <div v-else class="flex items-center justify-center w-full aspect-square bg-gray-100 dark:bg-gray-800 rounded-lg">
    <UEmpty
      icon="i-lucide-image"
      title="暂无图片"
      description="该商品暂时没有图片"
    />
  </div>
</template>

<style scoped>
.product-gallery {
  width: 100%;
}

@media (max-width: 768px) {
  .product-gallery {
    flex-direction: column;
  }
}
</style>