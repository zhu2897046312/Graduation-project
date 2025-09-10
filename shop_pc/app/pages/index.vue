<script setup lang="ts">
import api from '../../api';
import { useAsyncData } from 'nuxt/app';

const { data: hotList, status: hotStatus } = await useAsyncData('hot', async () => {
  return (await api.shop.product.list({ page_no: 1, page_size: 12, hot: '1' })).list
})
const { data: newList, status: newStatus } = await useAsyncData('new', async () => {
  return (await api.shop.product.list({ page_no: 1, page_size: 20 })).list
})
const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  const res = await api.shop.market.siteInfo()
  return  {
    ...res,
    // Set default SEO values if null
    seo_title: res.seo_title || `${res.title} | Shop Name`,
    seo_keyword: res.seo_keyword || `${res.title}, ${res.tags?.map((t : any) => t.title).join(', ') || ''}`,
    seo_description: res.seo_description || `Buy ${res.title} at best price. ${res.property_list?.map((p : any) => `${p.title}: ${p.value}`).join('. ') || ''}`
  };
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
