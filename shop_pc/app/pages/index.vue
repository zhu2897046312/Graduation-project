<script setup lang="ts">
import api from '../../api';
import type { ProductList } from '../../api/type'
import { useAsyncData } from 'nuxt/app';
import { useSiteInfo } from '../composable/useSiteInfo';

const { data: siteInfo } = await useSiteInfo()

const { data: hotList, status: hotStatus } = await useAsyncData('hot', async () => {
return (await api.shop.product.list({ page_no: 1, page_size: 12, hot: '1' })).list as ProductList
})

const { data: newList, status: newStatus } = await useAsyncData('new', async () => {
  return (await api.shop.product.list({ page_no: 1, page_size: 20 })).list as ProductList
})

onMounted(() => {
  console.log("index product list",hotList.value)
  console.log("index siteInfo ", siteInfo.value)
})
</script>

<template>
  <div class="container">
    <template v-if="siteInfo != null">
      <Title>{{ siteInfo.seo_title }}</Title>
      <Meta name="keywords" :content="siteInfo.seo_keyword" />
      <Meta name="description" :content="siteInfo.seo_description" />
    </template>

    <div class="text-center py-8">
        <h3 class="text-3xl font-bold text-gray-500 mb-2">Hot Products</h3>
        <p class="text-gray-400 max-w-2xl mx-auto">Discover our most popular items loved by customers</p>
      </div>
    <div v-if="hotStatus == 'success'" class="grid grid-cols-4 gap-6">
      <ProductCard
        v-for="item in hotList" :key="item.id"
        :title="item.title"
        :product-id="item.id"
        :original-price="item.original_price"
        :price="item.price"
        :thumb="item.picture" />
    </div>

    <div class="text-center py-8">
        <h3 class="text-3xl font-bold text-gray-500 mb-2">New Arrivals</h3>
        <p class="text-gray-400 max-w-2xl mx-auto">Explore our latest products just added to the collection</p>
      </div>
    <div v-if="newStatus == 'success'" class="grid grid-cols-4 gap-6">
      <ProductCard
        v-for="item in newList" :key="item.id"
        :title="item.title"
        :product-id="item.id"
        :original-price="item.original_price"
        :price="item.price"
        :thumb="item.picture" />
    </div>
  </div>
</template>
