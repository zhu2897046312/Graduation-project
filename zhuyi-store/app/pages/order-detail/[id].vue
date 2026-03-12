<script setup lang="ts">
import api from '../../api'
import type { OrderInfoResponse } from '../../api/type'
import { OrderStatus, OrderStatusLabel } from '../../api/type'
import { getProductImage } from '../../utils/auth'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

const route = useRoute()
const router = useRouter()
const toast = useToast()
const accessToken = useCookie('accessToken')

// 获取订单状态颜色
const getStatusColor = (state: OrderStatus) => {
  switch (state) {
    case OrderStatus.PendingPayment:
      return 'warning'
    case OrderStatus.PendingShipment:
      return 'info'
    case OrderStatus.Shipped:
    case OrderStatus.Completed:
      return 'success'
    case OrderStatus.Closed:
    case OrderStatus.Invalid:
      return 'error'
    default:
      return 'neutral'
  }
}

// 获取订单状态文本
const getStatusLabel = (state: OrderStatus) => {
  return OrderStatusLabel[state] || 'Unknown Status'
}

// 获取订单详情
const { data: res, pending } = await useAsyncData<OrderInfoResponse>(`order-${route.params.id}`, async () => {
  try {
    const result = await api.shop.order.get(route.params.id as string)
    return result
  } catch (error) {
    console.error('Failed to fetch order:', error)
    throw createError({
      statusCode: 404,
      message: 'Order not found'
    })
  }
})

// SEO
useHead({
  title: res.value ? `Order #${res.value.order.code} - Details` : 'Order Details',
  meta: [
    { name: 'description', content: 'View your order details' }
  ]
})

// 处理返回订单列表
const handleBack = () => {
  if (accessToken.value) {
    router.push('/account/orders')
  } else {
    router.push('/')
  }
}

// 格式化价格
const formatPrice = (price: unknown) => {
  const num = Number(price)
  return !isNaN(num) ? `$${num.toFixed(2)}` : '$0.00'
}

// 格式化日期
const formatDate = (dateString: string | null) => {
  if (!dateString) return 'N/A'
  try {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch (error) {
    return dateString
  }
}
</script>

<template>
  <div class="order-detail-page max-w-7xl mx-auto w-full px-4 sm:px-6 lg:px-8 py-8">
    <!-- 返回按钮 -->
    <UButton
      variant="ghost"
      color="neutral"
      icon="i-lucide-arrow-left"
      class="mb-6"
      @click="handleBack"
    >
      Back to Orders
    </UButton>

    <!-- 游客提示 -->
    <UAlert
      v-if="!accessToken && res"
      color="info"
      variant="soft"
      icon="i-lucide-info"
      class="mb-6"
    >
      <template #title>
        Guest Order
      </template>
      <template #description>
        <p class="text-sm">
          This order is associated with your current device.
          <NuxtLink
            to="/account/register"
            class="text-primary-600 dark:text-primary-400 hover:underline font-medium"
          >
            Register an account
          </NuxtLink>
          to access your orders from any device.
        </p>
      </template>
    </UAlert>

    <!-- 加载状态 -->
    <div v-if="pending" class="flex justify-center items-center py-20">
      <UCard class="p-8">
        <div class="flex flex-col items-center gap-4">
          <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
          <p class="text-sm text-gray-600 dark:text-gray-400">Loading order details...</p>
        </div>
      </UCard>
    </div>

    <!-- 订单详情 -->
    <div v-else-if="res" class="space-y-6">
      <!-- 订单概览卡片 -->
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="p-2 rounded-lg bg-primary-100 dark:bg-primary-900/30">
                <UIcon name="i-lucide-receipt" class="w-5 h-5 text-primary-600 dark:text-primary-400" />
              </div>
              <div>
                <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
                  Order #{{ res.order.code }}
                </h1>
                <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">
                  <span v-if="res.order.payment_time">
                    Paid on {{ formatDate(res.order.payment_time) }}
                  </span>
                  <span v-else>
                    Order placed
                  </span>
                </p>
              </div>
            </div>
            <UBadge
              :color="getStatusColor(res.order.state as OrderStatus)"
              variant="subtle"
              size="lg"
            >
              {{ getStatusLabel(res.order.state as OrderStatus) }}
            </UBadge>
          </div>
        </template>

        <!-- 商品列表 -->
        <div class="space-y-4">
          <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
            Order Items
          </h2>
          
          <div v-if="res.items && res.items.length > 0" class="space-y-3">
            <UCard
              v-for="item in res.items"
              :key="item.id"
              class="hover:shadow-md transition-shadow"
            >
              <div class="flex gap-4">
                <!-- 商品图片 -->
                <div class="flex-shrink-0">
                  <div class="w-20 h-20 rounded-lg overflow-hidden bg-gray-100 dark:bg-gray-800">
                    <img
                      :src="getProductImage(item.thumb)"
                      :alt="item.title"
                      class="w-full h-full object-cover"
                      @error="(e: Event) => { 
                        const target = e.target as HTMLImageElement
                        if (target) target.src = '/placeholder-product.jpg' 
                      }"
                    />
                  </div>
                </div>

                <!-- 商品信息 -->
                <div class="flex-1 min-w-0">
                  <h3 class="text-base font-medium text-gray-900 dark:text-white mb-1">
                    {{ item.title || 'No Product Name' }}
                  </h3>
                  <p
                    v-if="item.sku_title"
                    class="text-sm text-gray-500 dark:text-gray-400 mb-2"
                  >
                    {{ item.sku_title }}
                  </p>
                  <div class="flex items-center gap-4">
                    <UBadge color="neutral" variant="subtle">
                      Qty: {{ item.quantity || 0 }}
                    </UBadge>
                    <span class="text-lg font-semibold text-primary-600 dark:text-primary-400">
                      {{ formatPrice(item.pay_amount) }}
                    </span>
                  </div>
                </div>
              </div>
            </UCard>
          </div>

          <UEmpty
            v-else
            icon="i-lucide-package-x"
            title="No items found"
            description="No items found in this order"
            class="py-8"
          />
        </div>

        <!-- 订单汇总 -->
        <div class="mt-6 pt-6 border-t border-gray-200 dark:border-gray-700">
          <div class="space-y-3">
            <div class="flex justify-between items-center text-sm">
              <span class="text-gray-600 dark:text-gray-400">Items Subtotal</span>
              <span class="text-gray-900 dark:text-white font-medium">
                {{ formatPrice((res.order.pay_amount || 0) - (res.order.freight || 0)) }}
              </span>
            </div>
            <div v-if="res.order.freight > 0" class="flex justify-between items-center text-sm">
              <span class="text-gray-600 dark:text-gray-400">Freight</span>
              <span class="text-gray-900 dark:text-white font-medium">
                {{ formatPrice(res.order.freight) }}
              </span>
            </div>
            <div class="flex justify-between items-center pt-3 border-t border-gray-200 dark:border-gray-700">
              <span class="text-lg font-semibold text-gray-900 dark:text-white">Total Amount</span>
              <span class="text-2xl font-bold text-primary-600 dark:text-primary-400">
                {{ formatPrice(res.order.pay_amount) }}
              </span>
            </div>
          </div>
        </div>
      </UCard>

      <!-- 收货地址卡片 -->
      <UCard>
        <template #header>
          <div class="flex items-center gap-2">
            <UIcon name="i-lucide-map-pin" class="w-5 h-5 text-primary-600" />
            <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
              Shipping Address
            </h2>
          </div>
        </template>

        <div v-if="res.address" class="space-y-4">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- 地址信息 -->
            <div class="space-y-3">
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Recipient</p>
                <p class="text-base font-semibold text-gray-900 dark:text-white">
                  {{ res.address.first_name || '' }} {{ res.address.last_name || '' }}
                </p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Street Address</p>
                <p class="text-base text-gray-900 dark:text-white">
                  {{ res.address.detail_address || 'No address details' }}
                </p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">City & Region</p>
                <p class="text-base text-gray-900 dark:text-white">
                  <template v-if="res.address.city || res.address.region || res.address.postal_code">
                    {{ res.address.city }}{{ res.address.city && res.address.region ? ',' : '' }}
                    {{ res.address.region }} {{ res.address.postal_code }}
                  </template>
                  <template v-else>
                    No city/region information
                  </template>
                </p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Province & Country</p>
                <p class="text-base text-gray-900 dark:text-white">
                  <template v-if="res.address.province || res.address.country">
                    {{ res.address.province }}{{ res.address.province && res.address.country ? ',' : '' }}
                    {{ res.address.country }}
                  </template>
                  <template v-else>
                    No province/country information
                  </template>
                </p>
              </div>
            </div>

            <!-- 联系信息 -->
            <div class="space-y-3">
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Contact Information</p>
              </div>
              <div class="flex items-center gap-2">
                <UIcon name="i-lucide-phone" class="w-4 h-4 text-gray-400" />
                <p class="text-base text-gray-900 dark:text-white">
                  {{ res.address.phone || 'Not provided' }}
                </p>
              </div>
              <div class="flex items-center gap-2">
                <UIcon name="i-lucide-mail" class="w-4 h-4 text-gray-400" />
                <p class="text-base text-gray-900 dark:text-white">
                  {{ res.address.email || 'Not provided' }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <UEmpty
          v-else
          icon="i-lucide-map-pin-off"
          title="No shipping address"
          description="No shipping address information available"
          class="py-8"
        />
      </UCard>
    </div>

    <!-- 订单不存在 -->
    <UEmpty
      v-else
      icon="i-lucide-file-question"
      title="Order not found"
      description="We couldn't find this order. Please check the order number."
      class="py-12"
    >
      <template #actions>
        <div class="flex gap-3 justify-center">
          <UButton
            color="primary"
            icon="i-lucide-home"
            @click="router.push('/')"
          >
            Return to Home
          </UButton>
          <UButton
            variant="outline"
            color="neutral"
            icon="i-lucide-arrow-left"
            @click="handleBack"
          >
            Back to Orders
          </UButton>
        </div>
      </template>
    </UEmpty>
  </div>
</template>

<style scoped>
.order-detail-page {
  width: 100%;
}
</style>
