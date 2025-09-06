<script setup lang="ts">

import api from '~/api';
import { NPagination } from 'naive-ui'

const router = useRouter()
const route = useRoute()

const page = useState('page', () => Number.parseInt(route.query.page as string || '1'))
const title = useState('title', () => route.query.q as string || '')
const page_count = useState('page_count', () => 1)
const page_size = 20

const { data, refresh } = await useAsyncData(`search-product${title}`, async () => {
  const res = await api.shop.product.list({page_no: page.value, page_szie: page_size, title: title.value})
  console.log(res, 'search-product-res')
  page_count.value = Math.ceil(res.total / page_size)
  return {products: res.list}
})

const onUpdatePage = (page: number) => {
  console.log('page', page)
  router.push('?page=' + page)
  refresh()
}

watch(route, (newVal) => {
  console.log(newVal.query)
  title.value = newVal.query.q as string
  refresh()
})
</script>

<template>
  
  <Title>{{ title }} - Search</Title>
  <Meta name="keywords" :content="title" />
  <Meta name="description" :content="title" />
  <div class="container">
    <div class="m-8 flex flex-col items-center" v-if="data && data.products.length > 0" style="font-size: 1.333em;color: #545454;">
      Search results for "{{ title }}":
    </div>
    <div class="m-8 flex flex-col items-center" v-else style="font-size: 1.333em;color: #545454;">
      No results found for "{{ title }}".
    </div>
    <!-- 商品列表 -->
    <div class="product_list" v-if="data">
      <div v-for="item in data.products">
        <ProductCard
          :product-id="item.id"
          :title="item.title"
          :thumb="item.picture"
          :price="item.price"
          :original-price="item.original_price" />
      </div>
    </div>

    <div class="pagination" v-if="page_count > 1">
      <ClientOnly>
        <NPagination v-model:page="page" :page-count="page_count" @update:page="onUpdatePage" />
      </ClientOnly>
    </div>
  </div>
</template>

