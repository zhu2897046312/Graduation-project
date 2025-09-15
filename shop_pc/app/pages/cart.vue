<script setup lang="ts">
import api from '../../api';
import type { CartList } from '../../api/type';
import { NTag, NPopconfirm, NInputNumber, NCard, NButton, useMessage,NEmpty } from 'naive-ui';

const message = useMessage();
const cartNum = useState('cartNum')

useSeoMeta({
  title: 'Shopping Cart',
  description: 'Shopping Cart',
  keywords: 'Shopping Cart'
})

const { data: cart_list, refresh } = await useAsyncData(async () => {
  try {
    const res: any = await api.shop.cart.list() 
    
    let count = 0
    res.list.forEach((it: any) => {
      count += it.quantity
    })
    cartNum.value = count
    return res.list as CartList;
  } catch (error) {
    console.error('Failed to fetch cart:', error)
    return [] as CartList // 提供空数组作为默认值
  }
})


const lock = useState<boolean>('lock', () => false)

const total_amount = computed(() => {
  let total = 0
  cart_list.value?.forEach((item: any) => {
    total += Number(item.original_price) * item.quantity
  })
  return total
})
const pay_amount = computed(() => {
  let total = 0
  cart_list.value?.forEach((item: any) => {
    total += Number(item.price) * item.quantity
  })
  return total
})

const handleChangeCartItem = async (e: any, value: number) => {
  if (!cart_list.value) return
  
  if (lock.value) return
  
  lock.value = true
  let messageReactive = message.loading('Updating cart', { duration: 0 })
  
  // 使用临时变量
  const currentCartList = cart_list.value
  const cartItem = currentCartList[e]
  
  try {
    let _quantity = cartItem!.quantity - value
    if (_quantity != 0) {
      await api.shop.cart.act({
        product_id: cartItem!.product_id,
        sku_id: cartItem!.sku_id,
        quantity: Math.abs(_quantity),
        add: _quantity < 0
      })
      refresh()
    }
  } finally {
    messageReactive.destroy()
    lock.value = false
  }
}

/**
 * Remove item from cart
 * @param e Item index
 */
const handleRemoveCartItem = async (e: any) => {
  if (cart_list.value) {
    if (lock.value) {
      return
    }
    lock.value = true
    let messageReactive = message.loading('updating cart', { duration: 0 })
    
    const currentCartList = cart_list.value
    const productID = currentCartList[e]!.product_id
    const skuID = currentCartList[e]!.sku_id
    const quantity = currentCartList[e]!.quantity
    try {
      await api.shop.cart.act({
        product_id: productID,
        sku_id: skuID,
        quantity: quantity,
        add: false
      })
      refresh()
    } finally {
      messageReactive.destroy()
      lock.value = false
    }

  }
}

const router = useRouter()
const handleToCheckout = () => {
  router.push('/checkout')
}

const formatPrice = (price: unknown) => {
  // 转换为数字
  const num = Number(price)
  // 检查是否为有效数字
  return !isNaN(num) ? `$${num.toFixed(2)}` : '$0.00'
}

onMounted(() => {
  console.log("cart list: ", cart_list.value)
})
</script>

<template>
  <div class="container flex gap-2 mt-5">
    <div class="cart-box">
      <NCard v-if="cart_list && cart_list.length > 0" title="Shopping Cart">
        <div class="cart-list">
          <div class="cart-item" v-for="(item, index) in cart_list">
            <div class="cart-item-thumb">
              <img :src="item.thumb" :alt="item.title"></img>
            </div>
            <div class="cart-item-info">
              <div class="cart-item-title">
                <NuxtLink class="ami_link" :href="`/product/${item.product_id}`">{{ item.title }}</NuxtLink>
              </div>
              <div class="cart-item-desc" v-if="item.sku_title && item.sku_title.length > 0">Specification: {{ item.sku_title }}</div>
            </div>
            <div class="cart-item-price">
              <div class="flex items-baseline gap-1.5">
                <div class="discounted-price text-[#FF5000] font-medium text-base leading-tight">
                  {{ formatPrice(item.price) }}
                </div>
                <div v-if="Number(item.price) < Number(item.original_price)" class="original-price text-gray-500 text-xs line-through leading-tight">
                  {{ formatPrice(item.original_price) }}
                </div>
              </div>
            </div>
            <div class="cart-item-quantity"> 
                <NInputNumber 
                  :disabled="lock" 
                  :min="1" 
                  :max="50" 
                  :value="item.quantity" 
                  @update-value="(e: any) => handleChangeCartItem(index, e)"
                >
                  <template #add-icon>
                    <svg width="16" height="16" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M256 112V400M400 256H112" stroke="currentColor" stroke-width="32" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>
                  </template>
                  <template #minus-icon>
                    <svg width="16" height="16" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <line x1="400" y1="256" x2="112" y2="256" stroke="currentColor" stroke-width="32" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>
                  </template>
                </NInputNumber>
            </div>
            <div class="cart-item-action">
              <NPopconfirm :disabled="lock" @positive-click="handleRemoveCartItem(index)" positive-text="Confirm" negative-text="Cancel">
                <template #trigger>
                  <NTag type="error">Remove</NTag>
                </template>
                Remove this item?
              </NPopconfirm>
            </div>
          </div>
        </div>
      </NCard>
      <NEmpty v-else description="Your cart is empty">
          <template #extra>
            <NButton type="primary" @click="router.push('/')">
              Continue Shopping
            </NButton>
          </template>
        </NEmpty>
    </div>

    <div class="cart-computer"  v-if="cart_list && cart_list.length > 0">
      <NCard title="Order Summary">
        <div class="cell_item">
          <div class="cell_item_label">Items Subtotal</div>
          <div class="cell_item_value">
            <div class="price">{{ total_amount.toFixed(2) }}</div>
          </div>
        </div>
        <div class="cell_item">
          <div class="cell_item_label">Discount</div>
          <div class="cell_item_value">
            <div class="price">{{ (total_amount - pay_amount).toFixed(2) }}</div>
          </div>
        </div>
        <div class="cell_item">
          <div class="cell_item_label">Total</div>
          <div class="cell_item_value">
            <div class="price">{{ pay_amount.toFixed(2) }}</div>
          </div>
        </div>
        <div class="text-gray-400 py-1">Tax included. Shipping calculated at checkout.</div>
        <div class="mt-2">
          <NButton type="info" block @click="handleToCheckout">Proceed to Checkout</NButton>
        </div>
      </NCard>
    </div>
  </div>
</template>


<style lang="css" scoped>
.cart-box {
  flex: 1;
  min-height: 450px;
}
.cart-list {
  padding: 10px 0;
}
.cart-item {
  padding-top: 10px;
  padding-bottom: 10px;
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid #f5f5f5;
}
.cart-item-thumb {
  display: flex;
  width: 50px;
  height: 50px;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  background-color: #f5f5f5;
}
.cart-item-thumb img {
  display: block;
  max-width: 90%;
  max-height: 90%;
}
.cart-item-info  {
  flex: 1;
}
.cart-item-info .cart-item-title {
  font-size: 15px;
  color: var(--font-text-color);
}
.cart-item-info  .cart-item-desc {
  font-size: 14px;
  color: var(--font-text-color);
  opacity: 0.8;
}
.cart-item-price {
  padding-right: 20px;
  color: var(--price-color);
  font-size: 16px;
  font-weight: 600;
  line-height: 20px;
}
.cart-item-quantity {
  width: 130px;
  padding-right: 20px;
}


.cart-computer {
  width: 380px;
}
.cell_item_value .price {
  font-size: 18px;
}
</style>
