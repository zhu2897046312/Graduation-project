<script setup lang="ts">
const props = defineProps<{
  productId: number,
  title: string,
  thumb: string,
  price: string,
  originalPrice: string
}>()
</script>

<template>
  <div class="group block transition-opacity duration-300 hover:opacity-80">
    <NuxtLink 
      :to="`/product/${props.productId}`" 
      :title="props.title"
      class="block"
    >
      <!-- 图片容器 - 强制正方形比例 -->
      <div class="relative w-full aspect-square overflow-hidden">
        <img 
          class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105" 
          :src="props.thumb" 
          :alt="props.title"
          loading="lazy"
        />
      </div>
    </NuxtLink>

    <!-- 商品信息 -->
    <div class="mt-3 space-y-1">
      <NuxtLink 
        :to="`/product/${props.productId}`" 
        :title="props.title"
        class="block text-gray-800 text-base leading-snug line-clamp-2 h-[3.2em] hover:text-[#f4b3c2] transition-colors"
      >
        {{ props.title }}
      </NuxtLink>
      
      <div class="flex items-center gap-2">
        <span class="text-xl font-medium text-[#ff4d4f]">
          ${{ props.price }}
        </span>
        <span class="text-sm text-gray-500 line-through">
          ${{ props.originalPrice }}
        </span>
      </div>
    </div>
  </div>
</template>
<style scoped>
/* 确保这里没有语法错误 */
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