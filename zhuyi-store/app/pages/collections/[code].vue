<script setup lang="ts">
import api from '../../api'
import type { Category, ProductListResponse, ProductItem } from '../../api/type'
import { getProductImage } from '../../utils/auth'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

const router = useRouter()
const route = useRoute()
const code = route.params.code as string

// 分页参数
const page = ref(Number.parseInt((route.query.page as string) || '1'))
const pageSize = ref(20)

// 获取分类信息
const { data: info, pending: categoryPending } = await useAsyncData<Category>(`category-${code}`, async () => {
  const res = await api.shop.category.getInfoByCode(code)
  // 如果 SEO 字段为空，使用默认值
  return {
    ...res,
    seo_title: res.seo_title || `${res.title} Collection`,
    seo_keyword: res.seo_keyword || `${res.title}, collection, products`,
    seo_description: res.seo_description || `Browse our ${res.title} collection. Find the best ${res.title} products at great prices.`
  }
})

// 404 处理
if (!info.value) {
  throw createError({ statusCode: 404, message: 'Category not found' })
}

// 获取商品列表
const { data: productData, pending: productsPending, refresh } = await useAsyncData<ProductListResponse>(
  `products-${code}`,
  async () => {
    if (!info.value) {
      return {
        list: [],
        total: 0
      }
    }
    
    return await api.shop.product.list({
      category_id: info.value.id,
      page_no: page.value,
      page_size: pageSize.value
    })
  },
  {
    watch: [page]
  }
)

// 计算总页数
const totalPages = computed(() => {
  if (!productData.value) return 1
  const total = productData.value.total || 0
  return Math.ceil(total / pageSize.value)
})

// 商品列表
const products = computed<ProductItem[]>(() => {
  return productData.value?.list || []
})

// 处理商品图片
const getProductThumb = (picture: string) => {
  return getProductImage(picture)
}

// 分页更新
const onUpdatePage = (newPage: number) => {
  page.value = newPage
  router.push({ query: { page: newPage.toString() } })
}

// SEO
useHead({
  title: info.value?.seo_title || `${info.value?.title} Collection`,
  meta: [
    { name: 'keywords', content: info.value?.seo_keyword || '' },
    { name: 'description', content: info.value?.seo_description || '' }
  ]
})
</script>

<template>
  <!-- 页面标题 -->
  <div v-if="info" class="mb-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
      {{ info.title }}
    </h1>
    <p v-if="info.seo_description" class="text-gray-600 dark:text-gray-400">
      {{ info.seo_description }}
    </p>
  </div>

  <!-- 加载状态 -->
  <div v-if="categoryPending || productsPending" class="flex justify-center items-center py-20">
    <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
  </div>

  <!-- 商品列表 -->
  <div v-else-if="products && products.length > 0" class="space-y-8">
    <!-- 商品网格 - 一行4个 -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <ProductCard
        v-for="item in products"
        :key="item.id"
        :product-id="item.id"
        :title="item.title"
        :thumb="getProductThumb(item.picture)"
        :price="item.price"
        :original-price="item.original_price"
      />
    </div>

    <!-- 分页 -->
    <div v-if="totalPages > 1" class="flex justify-center mt-8">
      <UPagination
        v-model="page"
        :total="productData?.total || 0"
        :items-per-page="pageSize"
        :max="7"
        @update:model-value="onUpdatePage"
      />
    </div>
  </div>

  <!-- 空状态 -->
  <UEmpty
    v-else
    icon="i-lucide-package"
    title="暂无商品"
    description="该分类下暂时没有可用的商品，请稍后再试"
  />
</template>
