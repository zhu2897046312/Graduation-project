<script setup lang="ts">
import api from '../../../api';
import { NPagination,NEmpty } from 'naive-ui'

const router = useRouter()
const route = useRoute()
const code = route.params.code as string
const page = useState('page', () => Number.parseInt(route.query.page as string || '1'))
const page_count = useState('page_count', () => 1)
const page_size = 20

const { data: info } = await useAsyncData(`category-${code}`, async () => {
  const out: any = {}
  const res: any = await api.shop.category.getInfoByCode(code)
  out.category = {
    ...res,
    seo_title: res.seo_title || `${res.title} Collection `,
    seo_keyword: res.seo_keyword || `${res.title}, collection, products`,
    seo_description: res.seo_description || `Browse our ${res.title} collection. Find the best ${res.title} products at great prices.`
  };
  
  return out;
  
})
const { data, refresh, status } = await useAsyncData('products', async () => {
  const out: any = {}
  const res: any = await api.shop.product.list({category_id: info.value.category.id, page_no: page.value, page_size: page_size})


  out.products = res.list;
  out.total = res.total;
  
  page_count.value = Math.ceil(res.total / page_size)

  return out
})

if (!info.value) {
  throw createError({ statusCode: 404, message: 'Page not found' })
}

const onUpdatePage = (page: number) => {
  console.log('page', page)
  router.push('?page=' + page)
  refresh()
}

const onAddCart = (event: any) => {
  console.log('event', event)
}


</script>

<template>
  <div class="container" v-if="info != null">
    
    <Title>{{ info.category.seo_title }}</Title>
    <Meta name="keywords" :content="info.category.seo_keyword" />
    <Meta name="description" :content="info.category.seo_description" />

    <div class="list_header">
      <h2 class="list_header_title">{{ info.category.title }}</h2>
    </div>
    <!-- 商品列表 -->
    <div class="product_list" v-if="data.products.length > 0">
      <!-- <div v-for="item in data.products"> -->
        <ProductCard
          v-for="item in data.products"
          :key="item.id"
          :product-id="item.id"
          :title="item.title"
          :thumb="item.picture"
          :price="item.price"
          :original-price="item.original_price" 
          @add-cart="onAddCart"/>
      <!-- </div> -->
    </div>
    <NEmpty v-else description="No items found" class="py-6">
      <template #icon>
        <svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" viewBox="0 0 24 24" fill="none" stroke="#f4b3c2" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="10" cy="7" r="4"></circle>
          <path d="M15.5 21v-2a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2"></path>
          <line x1="3" y1="3" x2="21" y2="21"></line>
        </svg>
      </template>
    </NEmpty>
    <div class="pagination" v-if="page_count > 1">
      <ClientOnly>
        <NPagination v-model:page="page" :page-count="page_count" @update:page="onUpdatePage" />
      </ClientOnly>
    </div>

  </div>
</template>

<style lang="css" scoped>
.pagination {
  margin-top: 32px;
  display: flex;
  justify-content: center;
}
</style>
