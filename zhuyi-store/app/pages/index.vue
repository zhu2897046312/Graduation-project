<script setup lang="ts">
import api from '../api'
import type { ProductListResponse, ProductItem } from '../api/type'
import { getProductImage } from '../utils/auth'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

// 分页参数
const page = ref(1)
const pageSize = ref(20)

// 获取商品列表
const { data: productData, pending, refresh } = await useAsyncData<ProductListResponse>(
  'products',
  async () => {
    return await api.shop.product.list({
      page_no: page.value,
      page_size: pageSize.value
    })
  },
  {
    watch: [page]
  }
)

// 商品列表
const products = computed<ProductItem[]>(() => {
  return productData.value?.list || []
})

// 总商品数
const totalProducts = computed(() => {
  return productData.value?.total || 0
})

// 处理商品图片的工具函数
const getProductThumb = (picture: string) => {
  return getProductImage(picture)
}

// // 打印商品列表数据结构
// watchEffect(() => {
//   if (productData.value) {
//     console.group('🛍️ 商品列表数据')
//     console.log('完整商品数据:', JSON.stringify(productData.value, null, 2))
//     console.log('商品列表:', products.value)
//     console.log('商品数量:', products.value.length)
//     console.log('总商品数:', totalProducts.value)
//     if (products.value.length > 0) {
//       console.log('第一个商品:', products.value[0])
//       console.log('第一个商品的所有属性:', Object.keys(products.value[0] || {}))
//     }
//     console.groupEnd()
//   }
// })

// SEO
useHead({
  title: '首页 - 商品列表',
  meta: [
    { name: 'description', content: '浏览我们的精选商品，发现更多优质产品' }
  ]
})
</script>

<template>
  <!-- 页面标题 -->
  <div class="mb-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
      精选商品
    </h1>
    <p class="text-gray-600 dark:text-gray-400">
      发现更多优质产品
    </p>
  </div>

  <!-- 加载状态 -->
  <div v-if="pending" class="flex justify-center items-center py-20">
    <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
  </div>

  <!-- 商品列表 -->
  <div v-else-if="products && products.length > 0" class="space-y-8">
    <!-- 商品网格 - 一行4个 -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <ProductCard
        v-for="product in products"
        :key="product.id"
        :product-id="product.id"
        :title="product.title"
        :thumb="getProductThumb(product.picture)"
        :price="product.price"
        :original-price="product.original_price"
      />
    </div>

    <!-- 分页 -->
    <div v-if="totalProducts > pageSize" class="flex justify-center mt-8">
      <UPagination
        v-model="page"
        :total="totalProducts"
        :items-per-page="pageSize"
        :max="7"
      />
    </div>
  </div>

  <!-- 空状态 -->
  <UEmpty
    v-else
    icon="i-lucide-package"
    title="暂无商品"
    description="当前没有可用的商品，请稍后再试"
  />
</template>
