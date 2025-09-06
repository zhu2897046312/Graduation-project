<script setup lang="ts">
import api from '~/api';
import { NPagination } from 'naive-ui'

const router = useRouter()
const route = useRoute()
const code = route.params.code as string
const page = useState('page', () => Number.parseInt(route.query.page as string || '1'))
const page_count = useState('page_count', () => 1)
const page_size = 20

const { data: info } = await useAsyncData(`tag-${code}`, async () => {
  page.value = Number.parseInt(route.query.page as string || '1')
  return await api.shop.tag.info(code)
})
const { data, refresh, status } = await useAsyncData('tag-products-${code}', async () => {
  const out: any = {}
  const res: any = await api.shop.tag.product_list({code: code, page_no: page.value, page_size: page_size})

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
    
    <Title>{{ info.seo_title ? info.seo_title : info.title }}</Title>
    <Meta name="keywords" :content="info.seo_keyword ? info.seo_keyword : info.title" />
    <Meta name="description" :content="info.seo_description ? info.seo_description : info.title" />

    <div class="list_header">
      <h2 class="list_header_title">{{ info.title }}</h2>
    </div>
    <!-- 商品列表 -->
    <div class="product_list" v-if="data != null">
        <ProductCard
          v-for="item in data.products"
          :key="item.id"
          :product-id="item.id"
          :title="item.title"
          :thumb="item.picture"
          :price="item.price"
          :original-price="item.original_price" 
          @add-cart="onAddCart"/>
    </div>

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
