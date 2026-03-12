<script setup lang="ts">
const props = defineProps<{
  productId: any
  title: any
  thumb: any
  price: any
  originalPrice: any
}>()

// 图片加载状态
const imageError = ref(false)
const imageSrc = ref(props.thumb || '')

// 处理图片加载失败
const handleImageError = () => {
  imageError.value = true
  imageSrc.value = '/placeholder-product.jpg'
}

// 监听 thumb 变化
watch(() => props.thumb, (newThumb) => {
  if (newThumb) {
    imageError.value = false
    imageSrc.value = newThumb
  } else {
    imageError.value = true
    imageSrc.value = '/placeholder-product.jpg'
  }
}, { immediate: true })
</script>

<template>
  <div class="group block transition-opacity duration-300 hover:opacity-80">
    <NuxtLink 
      :to="`/product/${props.productId}`" 
      :title="props.title"
      class="block"
    >
      <UCard class="h-full hover:shadow-lg transition-shadow duration-300">
        <!-- 图片容器 - 强制正方形比例 -->
        <div class="relative w-full aspect-square overflow-hidden bg-gray-100 dark:bg-gray-800 rounded-t-lg flex items-center justify-center">
          <img 
            v-if="!imageError && imageSrc"
            class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105" 
            :src="imageSrc" 
            :alt="props.title"
            loading="lazy"
            @error="handleImageError"
          />
          <!-- 默认占位图标 -->
          <div v-else class="w-full h-full flex items-center justify-center bg-gray-200 dark:bg-gray-700">
            <UIcon 
              name="i-lucide-image" 
              class="w-16 h-16 text-gray-400 dark:text-gray-500"
            />
          </div>
        </div>

        <!-- 商品信息 -->
        <div class="mt-3 space-y-1 p-4">
          <NuxtLink 
            :to="`/product/${props.productId}`" 
            :title="props.title"
            class="block text-gray-800 dark:text-gray-200 text-base leading-snug line-clamp-2 h-[3.2em] hover:text-primary-600 dark:hover:text-primary-400 transition-colors"
          >
            {{ props.title }}
          </NuxtLink>
          
          <div class="flex items-center gap-2">
            <span class="text-xl font-medium text-[#ff4d4f] dark:text-red-400">
              ${{ props.price }}
            </span>
            <span 
              v-if="props.originalPrice && props.originalPrice > props.price"
              class="text-sm text-gray-500 dark:text-gray-400 line-through"
            >
              ${{ props.originalPrice }}
            </span>
          </div>
        </div>
      </UCard>
    </NuxtLink>
  </div>
</template>

<style scoped>
.group {
  display: block;
}

.aspect-square {
  aspect-ratio: 1/1;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.transition-opacity {
  transition-property: opacity;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 300ms;
}

.hover\:opacity-80:hover {
  opacity: 0.8;
}
</style>
