<script setup lang="ts">
import api from '~/api';
import { NPagination, NList, NListItem, NThing, NAlert, NInput, NButton, NEmpty } from 'naive-ui'
import { getDeviceId } from '~/utils/auth'

const route = useRoute()
const router = useRouter()
const accessToken = useCookie('accessToken')
const orderStatusMap = {
  1: 'Pending Payment',
  2: 'Pending Shipment',
  3: 'Shipped',
  4: 'Completed',
  5: 'Closed',
  6: 'Invalid Order'
} as const
// 订单状态映射
const stateArr = ['', 'Pending Payment', 'Pending Shipment', 'Shipped', 'Completed', 'Canceled', 'Invalid Order']

// 获取订单详情
const { data: res, pending} = await useAsyncData(`order-${route.params.id}`, async () => {
  try {
    const res = await api.shop.order.get(route.params.id as string)
    return res
  } catch (error) {
    console.error('Failed to fetch order:', error)
    return null
  }
})

// 处理返回订单列表
const handleBack = () => {
  if (accessToken.value) {
    router.push('/account/orders')
  } else {
    router.push('/')
  }
}

console.log(res)
</script>

<template>
  <div class="container mx-auto px-4 py-8 min-h-screen">
    <Title v-if="res">Order #{{ res.order.code }} - Details</Title>
    
    <!-- 返回按钮 -->
    <NButton 
      type="info"
      @click="handleBack"
      class="mb-6 bg-[#f4b3c2] hover:bg-[#e8a0b0] text-white"
    >
      ← Back to Orders
    </NButton>

    <div class="flex flex-col justify-between items-center mb-6">
      <div class="w-full">
        <!-- 游客提示 -->
        <NAlert 
          v-if="!accessToken"
          type="info"
          class="mb-6 border border-[#f4b3c2] bg-[#fdf2f5]"
        >
          <template #header>
            <span class="text-[#f4b3c2] font-semibold">Guest Order</span>
          </template>
          <p class="text-gray-700">This order is associated with your current device.</p>
          <NuxtLink 
            to="/account/register" 
            class="text-[#f4b3c2] font-medium hover:underline"
          >
            Register an account
          </NuxtLink> to access your orders from any device.
        </NAlert>
      </div>

      <!-- 加载状态 -->
      <div v-if="pending" class="text-center py-8 w-full">
        <div class="animate-pulse space-y-4">
          <div class="h-8 bg-[#fdf2f5] rounded w-1/3 mx-auto"></div>
          <div class="h-4 bg-[#fdf2f5] rounded w-1/2 mx-auto"></div>
        </div>
      </div>

      <!-- 订单详情 -->
      <div v-else-if="res" class="bg-white rounded-lg shadow-lg p-6 w-full border border-[#f4b3c2]/20">
        <!-- 订单概览 -->
        <div class="border-b border-[#f4b3c2]/30 pb-4 mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Order #{{ res.order.code }}</h2>
          <div class="flex justify-between mt-2">
            <span class="text-gray-600">Placed on {{ new Date(res.order.create_time).toLocaleDateString() }}</span>
            <span :class="{
              'text-yellow-500': res.order.state === 1 , 
              'text-green-400': res.order.state === 2,
              'text-green-500': res.order.state === 3 || res.order.state === 4,
              'text-red-500': res.order.state === 5 || res.order.state === 6,
              'text-[#f4b3c2]': res.order.state === 0
            }" class="font-medium">
              {{ stateArr[res.order.state] || 'Unknown Status' }}
            </span>
          </div>
        </div>

        <!-- 商品列表 -->
        <NList bordered class="border-[#f4b3c2]/30">
          <template v-if="res.items && res.items.length > 0">
            <NListItem v-for="item in res.items" :key="item.id" class="hover:bg-[#fdf2f5]">
              <NThing>
                <template #avatar>
                  <img :src="item.thumb || '/placeholder-product.jpg'" class="w-16 h-16 object-cover rounded border border-[#f4b3c2]/20" />
                </template>
                <template #header>
                  <span class="text-gray-800 font-medium">{{ item.title || 'No Product Name' }}</span>
                </template>
                <template #description>
                  <span class="text-gray-500">{{ item.sku_title || 'No SKU information' }}</span>
                </template>
                <template #header-extra>
                  <div class="flex items-center gap-4">
                    <span class="text-gray-800">x{{ item.quantity || 0 }}</span>
                    <span class="font-semibold text-[#f4b3c2]">${{ item.pay_amount || '0.00' }}</span>
                  </div>
                </template>
              </NThing>
            </NListItem>
          </template>
          <template v-else>
            <NListItem class="hover:bg-[#fdf2f5]">
              <NEmpty description="No items found in this order" class="py-8">
                <template #icon>
                  <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="#f4b3c2" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="10" cy="7" r="4"></circle>
                    <path d="M15.5 21v-2a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2"></path>
                    <line x1="3" y1="3" x2="21" y2="21"></line>
                  </svg>
                </template>
              </NEmpty>
            </NListItem>
          </template>
        </NList>

        <!-- 订单汇总 -->
        <div class="mt-6 border-t border-[#f4b3c2]/30 pt-6">
          <div class="flex justify-between mb-2">
            <span class="text-gray-600">Total:</span>
            <span class="text-gray-800">${{ res.order.pay_amount - res.order.freight  || '0.00' }}</span>
          </div>
          <div class="flex justify-between mb-2">
            <span class="text-gray-600">Freight:</span>
            <span class="text-gray-800">${{ res.order.freight || '0.00' }}</span>
          </div>
          <div class="flex justify-between font-bold text-lg mt-4 pt-4 border-t border-[#f4b3c2]/30">
            <span class="text-gray-800">Pay Amount:</span>
            <span class="text-[#f4b3c2]">${{ res.order.pay_amount  || '0.00' }}</span>
          </div>
        </div>

        <!-- 收货地址 -->
        <div class="mt-8 border-t border-[#f4b3c2]/30 pt-6">
          <h3 class="text-lg font-semibold mb-4 text-gray-800">Shipping Address</h3>
          <div v-if="res.address" class="grid grid-cols-1 md:grid-cols-2 gap-4 bg-[#fdf2f5] p-4 rounded-lg">
            <div>
              <p class="font-medium text-gray-800">{{ res.address.first_name || '' }} {{ res.address.last_name || '' }}</p>
              <p class="text-gray-700">{{ res.address.detail_address || 'No address details' }}</p>
              <p class="text-gray-700">
                <template v-if="res.address.city || res.address.region || res.address.postal_code">
                  {{ res.address.city }}{{ res.address.city && res.address.region ? ',' : '' }} 
                  {{ res.address.region }} {{ res.address.postal_code }}
                </template>
                <template v-else>
                  No city/region information
                </template>
              </p>
              <p class="text-gray-700">
                <template v-if="res.address.province || res.address.country">
                  {{ res.address.province }}{{ res.address.province && res.address.country ? ',' : '' }} 
                  {{ res.address.country }}
                </template>
                <template v-else>
                  No province/country information
                </template>
              </p>
            </div>
            <div>
              <p class="font-medium text-gray-800">Contact Information</p>
              <p class="text-gray-700">Phone: {{ res.address.phone || 'Not provided' }}</p>
              <p class="text-gray-700">Email: {{ res.address.email || 'Not provided' }}</p>
            </div>
          </div>
          <NEmpty v-else description="No shipping address information available" class="py-8">
            <template #icon>
              <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="#f4b3c2" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
                <circle cx="12" cy="10" r="3"></circle>
              </svg>
            </template>
          </NEmpty>
        </div>
      </div>

      <!-- 订单不存在 -->
      <div v-else class="text-center py-8 w-full">
        <NEmpty description="Order not found" class="py-8">
          <template #icon>
            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="#f4b3c2" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line>
            </svg>
          </template>
          <template #extra>
            <p class="text-gray-600 mb-4">We couldn't find this order. Please check the order number.</p>
            <div class="flex gap-4 justify-center">
              <NuxtLink to="/">
                <NButton type="primary" class="bg-[#f4b3c2] hover:bg-[#e8a0b0] text-white">Return to Home</NButton>
              </NuxtLink>
              <NButton @click="handleBack" class="border border-[#f4b3c2] text-[#f4b3c2] hover:bg-[#fdf2f5]">Back to Orders</NButton>
            </div>
          </template>
        </NEmpty>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
}

/* 自定义滚动条样式 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: #fdf2f5;
}

::-webkit-scrollbar-thumb {
  background: #f4b3c2;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #e8a0b0;
}
</style>