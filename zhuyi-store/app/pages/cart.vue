<script setup lang="ts">
import api from '../api'
import type { CartItem } from '../api/type'
import { getProductImage } from '../utils/auth'
import { useCartShared } from '../composables/useCartShared'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

const router = useRouter()
const toast = useToast()

// 使用共享的购物车 composable
const { cartNum, useCartList, refreshAllCartData } = useCartShared()

// SEO
useHead({
  title: 'Shopping Cart',
  meta: [
    { name: 'description', content: 'View and manage your shopping cart' },
    { name: 'keywords', content: 'Shopping Cart' }
  ]
})

// 获取购物车列表
const { data: cart_list, refresh, pending } = await useCartList('cart-list')

// 锁定状态，防止重复操作
const lock = ref(false)

// 计算总金额（原价）
const total_amount = computed(() => {
  if (!cart_list.value) return 0
  return cart_list.value.reduce((total, item) => total + item.original_price * item.quantity, 0)
})

// 计算实付金额（折扣后）
const pay_amount = computed(() => {
  if (!cart_list.value) return 0
  return cart_list.value.reduce((total, item) => total + item.price * item.quantity, 0)
})

// 计算折扣金额
const discount_amount = computed(() => {
  return total_amount.value - pay_amount.value
})

// 处理数量变化
const handleChangeCartItem = async (index: number, newQuantity: number) => {
  if (!cart_list.value || lock.value) return
  
  const cartItem = cart_list.value[index]
  if (!cartItem) return
  
  const oldQuantity = cartItem.quantity
  const quantityDiff = newQuantity - oldQuantity
  
  if (quantityDiff === 0) return
  
  lock.value = true
  
  try {
    await api.shop.cart.act({
      product_id: cartItem.product_id,
      sku_id: cartItem.sku_id,
      quantity: Math.abs(quantityDiff),
      add: quantityDiff > 0
    })
    
    // 刷新所有购物车数据（包括数量和所有页面的列表）
    await refreshAllCartData()
    
    toast.add({
      title: 'Cart Updated',
      description: 'Item quantity has been updated',
      color: 'success'
    })
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : 'Failed to update cart item'
    toast.add({
      title: 'Update Failed',
      description: errorMessage,
      color: 'error'
    })
  } finally {
    lock.value = false
  }
}

// 删除购物车商品
const handleRemoveCartItem = async (index: number) => {
  if (!cart_list.value || lock.value) return
  
  const cartItem = cart_list.value[index]
  if (!cartItem) return
  
  lock.value = true
  
  try {
    await api.shop.cart.act({
      product_id: cartItem.product_id,
      sku_id: cartItem.sku_id,
      quantity: cartItem.quantity,
      add: false
    })
    
    // 刷新所有购物车数据（包括数量和所有页面的列表）
    await refreshAllCartData()
    
    toast.add({
      title: 'Item Removed',
      description: 'Item has been removed from cart',
      color: 'success'
    })
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : 'Failed to remove item from cart'
    toast.add({
      title: 'Remove Failed',
      description: errorMessage,
      color: 'error'
    })
  } finally {
    lock.value = false
  }
}

// 跳转到结算页
const handleToCheckout = () => {
  router.push('/checkout')
}

// 格式化价格
const formatPrice = (price: number) => {
  return `$${price.toFixed(2)}`
}

// 处理商品图片
const getProductThumb = (thumb: string) => {
  return getProductImage(thumb)
}

// 确认删除（在 Modal 中调用）
const confirmDelete = async (index: number) => {
  await handleRemoveCartItem(index)
}
</script>

<template>
  <div class="cart-page max-w-7xl mx-auto w-full">
    <!-- 页面标题 -->
    <div class="mb-8">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
            Shopping Cart
          </h1>
          <p class="text-gray-600 dark:text-gray-400">
            Review and manage your items
          </p>
        </div>
        <UButton
          variant="ghost"
          color="neutral"
          icon="i-lucide-arrow-left"
          @click="router.push('/')"
        >
          继续购物
        </UButton>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="pending" class="flex justify-center items-center py-20">
      <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
    </div>

    <!-- 购物车内容 -->
    <div v-else-if="!pending && cart_list && Array.isArray(cart_list) && cart_list.length > 0" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- 购物车商品列表 -->
      <div class="lg:col-span-2 space-y-4">
        <UCard>
          <template #header>
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
                购物车商品 ({{ cart_list.length }})
              </h2>
              <UBadge color="primary" variant="subtle">
                {{ cart_list.length }} 件商品
              </UBadge>
            </div>
          </template>

          <div class="space-y-4">
            <div
              v-for="(item, index) in cart_list"
              :key="item.id"
              class="cart-item flex gap-4 p-4 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600 transition-colors"
            >
              <!-- 商品图片 -->
              <div class="cart-item-thumb flex-shrink-0">
                <NuxtLink :to="`/product/${item.product_id}`" class="block">
                  <div class="w-20 h-20 rounded-lg overflow-hidden bg-gray-100 dark:bg-gray-800 flex items-center justify-center">
                    <img
                      :src="getProductThumb(item.thumb)"
                      :alt="item.title"
                      class="w-full h-full object-cover"
                      @error="(e: Event) => { 
                        const target = e.target as HTMLImageElement
                        if (target) target.src = '/placeholder-product.jpg' 
                      }"
                    />
                  </div>
                </NuxtLink>
              </div>

              <!-- 商品信息 -->
              <div class="cart-item-info flex-1 min-w-0">
                <div class="cart-item-title mb-2">
                  <NuxtLink
                    :to="`/product/${item.product_id}`"
                    class="text-base font-medium text-gray-900 dark:text-white hover:text-primary-600 dark:hover:text-primary-400 transition-colors line-clamp-2"
                  >
                    {{ item.title }}
                  </NuxtLink>
                </div>
                <div
                  v-if="item.sku_title && item.sku_title.length > 0"
                  class="cart-item-desc text-sm text-gray-500 dark:text-gray-400 mb-3"
                >
                  <UBadge color="neutral" variant="subtle" size="xs">
                    {{ item.sku_title }}
                  </UBadge>
                </div>
                
                <!-- 价格和操作区域 -->
                <div class="flex items-center justify-between">
                  <!-- 价格 -->
                  <div class="flex items-baseline gap-2">
                    <span class="text-lg font-semibold text-[#ff4d4f] dark:text-red-400">
                      {{ formatPrice(item.price) }}
                    </span>
                    <span
                      v-if="Number(item.price) < Number(item.original_price)"
                      class="text-sm text-gray-500 dark:text-gray-400 line-through"
                    >
                      {{ formatPrice(item.original_price) }}
                    </span>
                  </div>

                  <!-- 数量控制和删除按钮 -->
                  <div class="flex items-center gap-3">
                    <div class="flex items-center gap-2">
                      <span class="text-sm text-gray-600 dark:text-gray-400">数量:</span>
                      <UInputNumber
                        :model-value="item.quantity"
                        :min="1"
                        :max="50"
                        :disabled="lock"
                        size="sm"
                        class="w-24"
                        @update:model-value="(val: number | null) => {
                          if (val !== null && val >= 1) {
                            handleChangeCartItem(index, val)
                          }
                        }"
                      />
                    </div>
                    <UModal>
                      <UButton
                        variant="ghost"
                        color="error"
                        size="sm"
                        icon="i-lucide-trash-2"
                        :disabled="lock"
                        aria-label="Remove item"
                      />
                      <template #content>
                        <UCard>
                          <template #header>
                            <div class="flex items-center justify-between">
                              <div class="flex items-center gap-3">
                                <UIcon name="i-lucide-alert-triangle" class="w-6 h-6 text-error-600 dark:text-error-400" />
                                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                                  确认删除
                                </h3>
                              </div>
                            </div>
                          </template>

                          <div class="py-4">
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                              确定要从购物车中移除
                              <span class="font-medium text-gray-900 dark:text-white">
                                "{{ item.title }}"
                              </span>
                              吗？
                            </p>
                          </div>

                          <template #footer>
                            <div class="flex justify-end gap-3">
                              <UButton
                                variant="ghost"
                                color="neutral"
                              >
                                取消
                              </UButton>
                              <UButton
                                color="error"
                                @click="confirmDelete(index)"
                                :disabled="lock"
                              >
                                删除
                              </UButton>
                            </div>
                          </template>
                        </UCard>
                      </template>
                    </UModal>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </UCard>
      </div>

      <!-- 订单摘要 -->
      <div class="lg:col-span-1">
        <UCard class="sticky top-24">
          <template #header>
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
              订单摘要
            </h2>
          </template>

          <div class="space-y-4">
            <!-- 商品小计 -->
            <div class="flex justify-between items-center text-sm">
              <span class="text-gray-600 dark:text-gray-400">商品小计</span>
              <span class="text-gray-900 dark:text-white font-medium">
                {{ formatPrice(total_amount) }}
              </span>
            </div>

            <!-- 折扣 -->
            <div
              v-if="discount_amount > 0"
              class="flex justify-between items-center text-sm"
            >
              <span class="text-gray-600 dark:text-gray-400">折扣优惠</span>
              <span class="text-green-600 dark:text-green-400 font-medium">
                -{{ formatPrice(discount_amount) }}
              </span>
            </div>

            <!-- 分隔线 -->
            <div class="border-t border-gray-200 dark:border-gray-700 pt-4">
              <div class="flex justify-between items-center">
                <span class="text-lg font-semibold text-gray-900 dark:text-white">Total</span>
                <span class="text-2xl font-bold text-[#ff4d4f] dark:text-red-400">
                  {{ formatPrice(pay_amount) }}
                </span>
              </div>
            </div>

            <!-- 提示信息 -->
            <p class="text-xs text-gray-500 dark:text-gray-400 pt-2">
              税费已包含。运费将在结算时计算。
            </p>

            <!-- 结算按钮 -->
            <UButton
              color="primary"
              size="lg"
              block
              icon="i-lucide-shopping-bag"
              :disabled="lock || pending"
              :loading="lock || pending"
              @click="handleToCheckout"
            >
              前往结算
            </UButton>
          </div>
        </UCard>
      </div>
    </div>

    <!-- 空购物车状态 -->
    <UEmpty
      v-else-if="!pending && (!cart_list || !Array.isArray(cart_list) || cart_list.length === 0)"
      icon="i-lucide-shopping-cart"
      title="购物车是空的"
      description="添加一些商品到购物车继续购物"
      class="py-12"
    >
      <template #actions>
        <UButton
          color="primary"
          size="lg"
          icon="i-lucide-arrow-left"
          @click="router.push('/')"
        >
          继续购物
        </UButton>
      </template>
    </UEmpty>


  </div>
</template>


<style scoped>
.cart-page {
  width: 100%;
}

.cart-item {
  transition: all 0.2s ease;
  background-color: transparent;
}

.cart-item:hover {
  background-color: rgba(0, 0, 0, 0.02);
}

.dark .cart-item:hover {
  background-color: rgba(255, 255, 255, 0.02);
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* 响应式优化 */
@media (max-width: 1024px) {
  .cart-item {
    flex-direction: column;
    gap: 1rem;
  }
  
  .cart-item-thumb {
    align-self: flex-start;
  }
  
  .cart-item-info {
    width: 100%;
  }
}

@media (max-width: 640px) {
  .cart-item {
    padding: 1rem;
  }
}
</style>
