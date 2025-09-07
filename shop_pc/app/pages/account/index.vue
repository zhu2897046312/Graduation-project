<script setup lang="ts">
import api from "../../../api";
import { NList, NListItem, NThing,NButton } from "naive-ui"


const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})

const { data } = await useAsyncData('address', async () => {
  const res = await api.shop.address.list({ page_no: 1, page_size: 1})
  if (res.list.length > 0) {
    return { address: res.list[0], total: res.total }
  } else {
    return { address: null, total: 0 }
  }
})

const {data: order, status: orderStatus} = await useAsyncData('orders', async () => {
  const res = await api.shop.order.list({page_no: 1, page_size: 100})
  if (res.list.length > 0) {
    return {list: res.list, total: res.total}
  }
  return {list: [], total: 0}
})

// 订单状态:1=待付款;2=待发货;3=已发货;4=已完成;5=已关闭;6=无效订单
const stateArr = ['', 'Pending Payment', 'Pending Shipment', 'Shipped', 'Completed', 'Canceled', 'Invalid Order']


</script>

<template>
  <div class="container">
    <template v-if="siteInfo != null">
      <Title>Order History - {{ siteInfo.seo_title ? siteInfo.seo_title : ' ' }}</Title>
      <Meta name="keywords" content="Order History " />
      <Meta name="description" content="Order History " />
    </template>
    <h2 class="title mt-8">My Account</h2>
    <div class="flex gap-8 mt-8">
      <div class="flex flex-col gap-4 flex-grow-[2] basis-0" v-if="orderStatus == 'success' && order">
        <span>Order History</span>
        <span v-if="order?.total == 0">You haven't placed any orders yet.</span>
        <div v-else>
          <NList bordered>
            <NListItem v-for="item in order.list" :key="item.id">
              <div class="order-products">
                <NThing class="product" v-for="p in item.items" :key="p.id">
                  <template #avatar>
                    <img class="product-thumb" :src="p.thumb" />
                  </template>
                  <template #header>
                    <span class="product-header">{{ p.title }}</span>
                  </template>
                  <template #description>
                    <span class="product-sku">{{ p.sku_title }}</span>
                    <span class="product-quantity">x{{ p.quantity }}</span>
                  </template>
                  <!-- <template #header-extra>
                    <span class="price">{{ p.pay_amount }}</span>
                  </template> -->
                </NThing>
              </div>

              <div class="order-total flex justify-between items-center pt-2 mt-2 border-t border-dashed border-gray-200">
                <div>
                  <span>Total:</span>
                  <span class="price mx-2">{{ item.pay_amount }}</span>
                </div>
                
                <div justify="end">
                  <NButton type="info" size="small" @click="navigateTo(`/order-detail/${item.visitor_query_code}`)">View Details</NButton>
                </div>
              </div>
              <div style="height: 20px;"></div>
            </NListItem>
          </NList>
        </div>
      </div>
      <div class="flex flex-col gap-4 flex-grow flex-shrink-0 basis-0">
        <span>Account Details</span>
        <div class="flex flex-col gap-1" v-if="data && data.address">
          <span>{{ data.address.first_name }}</span>
          <span>{{ data.address.detail_address }}</span>
          <span>{{ data.address.city }}</span>
          <span>{{ data.address.region }}</span>
          <span>{{ data.address.province }}</span>
        </div>
        <NuxtLink to="/account/addresses" v-if="data">View Addresses ({{ data.total }})</NuxtLink>
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
.title {
  color: #545454;
  font-size: 1.733em;
  font-style: normal;
}
a {
  color: #f4b3c2;
}
.product {
  margin-bottom: 10px;
}
.product-thumb {
  width: 48px;
  height: 48px;
}
.product-header {
  font-size: 14px;
  font-weight: normal;
}
</style>