<script setup lang="ts">
import api from '../api'
import type { CartItem, OrderCreateParams } from '../api/type'
import { getProductImage } from '../utils/auth'
import { useCartShared } from '../composables/useCartShared'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

const router = useRouter()
const toast = useToast()

// SEO
useHead({
  title: 'Checkout',
  meta: [
    { name: 'description', content: 'Complete your order checkout' }
  ]
})

// 支付方式
const selectedPayType = ref(1)
const payTypeOptions = [
  { value: 1, label: 'PayPal' }
]

// 表单数据
const address = reactive({
  first_name: '',
  last_name: '',
  email: '',
  country: undefined as string | undefined,
  phone: '',
  province: undefined as string | undefined,
  city: '',
  region: '',
  detail_address: '',
  postal_code: ''
})

// 表单验证错误
const errors = reactive({
  email: '',
  first_name: '',
  last_name: '',
  phone: '',
  country: '',
  province: '',
  city: '',
  detail_address: '',
  postal_code: ''
})

// 验证规则
const emailRules = [
  (v: string) => !!v || 'Email is required',
  (v: string) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v) || 'Please enter a valid email address'
]

const requiredRules = (fieldName: string) => [
  (v: any) => !!v || `${fieldName} is required`
]

// 验证单个字段
const validateField = (field: keyof typeof errors) => {
  let rules: ((v: any) => boolean | string)[] = []
  
  switch (field) {
    case 'email':
      rules = emailRules
      break
    case 'first_name':
    case 'last_name':
    case 'phone':
    case 'country':
    case 'province':
    case 'city':
    case 'detail_address':
    case 'postal_code':
      rules = requiredRules(field.replace('_', ' '))
      break
  }
  
  const value = (address as any)[field]
  for (const rule of rules) {
    const error = rule(value)
    if (error !== true) {
      errors[field] = typeof error === 'string' ? error : String(error)
      return false
    }
  }
  
  errors[field] = ''
  return true
}

// 检查字段是否验证通过（用于显示成功状态）
const isFieldValid = (field: keyof typeof errors) => {
  const value = (address as any)[field]
  return value && value.toString().trim() !== '' && !errors[field]
}

// 检查字段是否有错误（用于显示错误状态）
const isFieldError = (field: keyof typeof errors) => {
  return errors[field] && errors[field].length > 0
}

// 验证所有字段
const validateAll = () => {
  const fields: (keyof typeof errors)[] = [
    'email', 'first_name', 'last_name', 'phone', 'country', 
    'province', 'city', 'detail_address', 'postal_code'
  ]
  
  let isValid = true
  fields.forEach(field => {
    if (!validateField(field)) {
      isValid = false
    }
  })
  
  return isValid
}

// 地址数据类型
interface AddressData {
  countries: Array<{
    name: string
    full_name: string
    provinces: string[]
  }>
}

// 加载地址数据 - 直接导入本地 JSON 文件
// @ts-ignore
import addressJsonData from '../../data/address.json'

// 确保数据正确加载
const addressJson = ref<AddressData>(addressJsonData as AddressData)

// 立即检查数据
if (addressJson.value && addressJson.value.countries) {
  console.log('✅ addressJson imported successfully')
  console.log('📊 Countries loaded:', addressJson.value.countries.length)
  console.log('📋 First country:', addressJson.value.countries[0])
} else {
  console.error('❌ Failed to load addressJson')
  console.error('addressJsonData:', addressJsonData)
}

// 调试：检查数据是否加载成功
onMounted(() => {
  console.log('📦 addressJson loaded:', addressJson.value)
  console.log('🌍 Countries count:', addressJson.value?.countries?.length || 0)
  if (addressJson.value?.countries?.length > 0) {
    console.log('✅ First country:', addressJson.value.countries[0])
    console.log('✅ countryOptions:', countryOptions.value)
    console.log('✅ countryOptions length:', countryOptions.value.length)
  } else {
    console.warn('⚠️ No countries loaded from address.json')
  }
})

// 国家选项
const countryOptions = computed(() => {
  if (!addressJson.value) {
    console.warn('⚠️ addressJson.value is null')
    return []
  }
  
  const countries = addressJson.value.countries || []
  console.log('🔄 countryOptions computed, countries:', countries.length)
  
  if (countries.length === 0) {
    console.warn('⚠️ No countries in addressJson')
    return []
  }
  
  const options = countries.map((c) => {
    const option = {
      label: c.full_name,
      value: c.name
    }
    return option
  })
  
  console.log('✅ countryOptions created:', options.length, 'options')
  console.log('📋 First option:', options[0])
  console.log('📋 All options:', options)
  
  return options
})

// 省份选项
const provinceOptions = computed(() => {
  if (!address.country || !addressJson.value) return []
  
  const countries = addressJson.value.countries || []
  const selectedCountry = countries.find(
    (c) => c.name === address.country
  )
  
  return selectedCountry?.provinces?.map((p: string) => ({
    label: p,
    value: p
  })) || []
})

// 使用共享的购物车 composable
const { useCartList } = useCartShared()

// 获取购物车列表
const { data: products, status, pending: productsPending } = await useCartList('checkout-cart')

// 获取运费
const { data: freight, pending: freightPending } = await useAsyncData('freight', async () => {
  try {
    const res = await api.shop.market.freight()
    return Number(res) || 0
  } catch (error) {
    console.error('运费加载失败:', error)
    return 0
  }
})

// 计算总金额
const total_amount = computed(() => {
  if (!products.value) return 0
  return products.value.reduce((total, item) => total + item.price * item.quantity, 0)
})

// 格式化价格
const formatPrice = (price: unknown) => {
  const num = Number(price)
  return !isNaN(num) ? `$${num.toFixed(2)}` : '$0.00'
}

// 处理商品图片
const getProductThumb = (thumb: string) => {
  return getProductImage(thumb)
}

// 提交订单
const loading = ref(false)

const handleCheckout = async () => {
  if (!validateAll()) {
    toast.add({
      title: 'Validation Failed',
      description: 'Please fill in all required fields',
      color: 'error'
    })
    window.scrollTo({ top: 0, behavior: 'smooth' })
    return
  }
  
  if (!products.value || products.value.length === 0) {
    toast.add({
      title: 'Cart Empty',
      description: 'Your cart is empty',
      color: 'error'
    })
    return
  }
  
  loading.value = true
  
  try {
    if (!address.province) {
      throw new Error('Province is required')
    }
    if (!address.country) {
      throw new Error('Country is required')
    }
    
    const payload: OrderCreateParams = {
      product_items: products.value.map((item) => ({
        product_id: item.product_id,
        quantity: item.quantity,
        sku_id: item.sku_id
      })),
      pay_type: selectedPayType.value.toString(),
      first_name: address.first_name,
      last_name: address.last_name,
      email: address.email,
      phone: address.phone,
      province: address.province,
      country: address.country,
      city: address.city,
      postal_code: address.postal_code,
      detail_address: address.detail_address
    }
    
    const orderId = await api.shop.order.create(payload)
    console.log('checkout result', orderId)
    
    const data = {
      order_id: orderId,
      pay_type: selectedPayType.value
    }
    
    const paymentRes = await api.shop.order.getPaymentUrl(data)
    console.log('payment response', paymentRes)
    
    if (paymentRes?.approveUrl) {
      window.location.href = paymentRes.approveUrl
    } else {
      throw new Error('Payment URL not found')
    }
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : 'Failed to process checkout'
    toast.add({
      title: 'Checkout Failed',
      description: errorMessage,
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}

// 监听国家变化，清空省份
watch(() => address.country, () => {
  address.province = undefined
})
</script>

<template>
  <div class="checkout-page max-w-7xl mx-auto w-full">
    <!-- 页面标题 -->
    <div class="mb-8">
      <div class="flex items-center gap-3 mb-3">
        <div class="p-2 rounded-lg bg-primary-100 dark:bg-primary-900/30">
          <UIcon name="i-lucide-shopping-bag" class="w-6 h-6 text-primary-600 dark:text-primary-400" />
        </div>
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
          Checkout
        </h1>
      </div>
      <p class="text-gray-600 dark:text-gray-400 ml-14">
        Complete your order information
      </p>
    </div>

    <!-- 加载状态 -->
    <div v-if="productsPending || freightPending" class="flex justify-center items-center py-20">
      <UCard class="p-8">
        <div class="flex flex-col items-center gap-4">
          <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
          <p class="text-sm text-gray-600 dark:text-gray-400">Loading checkout information...</p>
        </div>
      </UCard>
    </div>

    <!-- 主要内容 -->
    <div v-else-if="status === 'success' && products && products.length > 0" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- 左侧：表单区域 -->
      <div class="lg:col-span-2 space-y-6">
        <!-- 配送信息 -->
        <UCard>
          <template #header>
            <div class="flex items-center gap-2">
              <UIcon name="i-lucide-truck" class="w-5 h-5 text-primary-600" />
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
                Shipping Information
              </h2>
            </div>
          </template>

          <UForm :state="address" @submit="handleCheckout" class="space-y-6">
            <!-- 姓名分组 -->
            <div class="space-y-5">
              <h3 class="text-sm font-semibold text-gray-700 dark:text-gray-300 uppercase tracking-wide">
                Name
              </h3>
              
              <!-- 姓名 -->
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <UFormField label="First Name" :error="errors.first_name" required>
                  <UInput
                    v-model="address.first_name"
                    placeholder="Enter first name"
                    size="lg"
                    :color="isFieldError('first_name') ? 'error' : isFieldValid('first_name') ? 'success' : undefined"
                    @blur="validateField('first_name')"
                  >
                    <template v-if="isFieldValid('first_name')" #trailing>
                      <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                    </template>
                  </UInput>
                </UFormField>

                <UFormField label="Last Name" :error="errors.last_name" required>
                  <UInput
                    v-model="address.last_name"
                    placeholder="Enter last name"
                    size="lg"
                    :color="isFieldError('last_name') ? 'error' : isFieldValid('last_name') ? 'success' : undefined"
                    @blur="validateField('last_name')"
                  >
                    <template v-if="isFieldValid('last_name')" #trailing>
                      <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                    </template>
                  </UInput>
                </UFormField>
              </div>
            </div>

            <!-- 联系方式分组 -->
            <div class="space-y-5 pt-2 border-t border-gray-200 dark:border-gray-700">
              <h3 class="text-sm font-semibold text-gray-700 dark:text-gray-300 uppercase tracking-wide">
                Contact Information
              </h3>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <!-- Email -->
                <UFormField label="Email Address" :error="errors.email" required>
                  <UInput
                    v-model="address.email"
                    type="email"
                    placeholder="Enter email address"
                    size="lg"
                    icon="i-lucide-mail"
                    :color="isFieldError('email') ? 'error' : isFieldValid('email') ? 'success' : undefined"
                    @blur="validateField('email')"
                  >
                    <template v-if="isFieldValid('email')" #trailing>
                      <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                    </template>
                  </UInput>
                  <template #description>
                    <span class="text-sm text-gray-500 dark:text-gray-400">Email me with news and offers</span>
                  </template>
                </UFormField>

                <!-- 电话 -->
                <UFormField label="Phone Number" :error="errors.phone" required>
                  <UInput
                    v-model="address.phone"
                    type="tel"
                    placeholder="Enter phone number"
                    size="lg"
                    icon="i-lucide-phone"
                    :color="isFieldError('phone') ? 'error' : isFieldValid('phone') ? 'success' : undefined"
                    @blur="validateField('phone')"
                  >
                    <template v-if="isFieldValid('phone')" #trailing>
                      <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                    </template>
                  </UInput>
                </UFormField>
              </div>
              
            </div>

            <!-- 地址信息分组 -->
            <div class="space-y-5 pt-2 border-t border-gray-200 dark:border-gray-700">
              <h3 class="text-sm font-semibold text-gray-700 dark:text-gray-300 uppercase tracking-wide">
                Shipping Address
              </h3>

              <!-- 国家和省份 -->
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <UFormField label="Country" :error="errors.country" required>
                  <div class="relative">
                    <USelect
                      v-model="address.country"
                      :items="countryOptions"
                      placeholder="Select Country"
                      size="lg"
                      :color="isFieldError('country') ? 'error' : isFieldValid('country') ? 'success' : undefined"
                      @update:model-value="validateField('country')"
                    />
                    <UIcon
                      v-if="isFieldValid('country')"
                      name="i-lucide-check-circle-2"
                      class="absolute right-10 top-1/2 -translate-y-1/2 w-5 h-5 text-green-500 pointer-events-none"
                    />
                  </div>
                </UFormField>

                <UFormField label="Region / Province" :error="errors.province" required>
                  <div class="relative">
                    <USelect
                      v-model="address.province"
                      :items="provinceOptions"
                      placeholder="Select region/province"
                      size="lg"
                      :disabled="!address.country"
                      :color="isFieldError('province') ? 'error' : isFieldValid('province') ? 'success' : undefined"
                      @update:model-value="validateField('province')"
                    />
                    <UIcon
                      v-if="isFieldValid('province')"
                      name="i-lucide-check-circle-2"
                      class="absolute right-10 top-1/2 -translate-y-1/2 w-5 h-5 text-green-500 pointer-events-none"
                    />
                  </div>
                </UFormField>
              </div>

              <!-- 城市和邮编 -->
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <UFormField label="City" :error="errors.city" required>
                  <UInput
                    v-model="address.city"
                    placeholder="Enter city"
                    size="lg"
                    :color="isFieldError('city') ? 'error' : isFieldValid('city') ? 'success' : undefined"
                    @blur="validateField('city')"
                  >
                    <template v-if="isFieldValid('city')" #trailing>
                      <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                    </template>
                  </UInput>
                </UFormField>

                <UFormField label="ZIP / Postal Code" :error="errors.postal_code" required>
                  <UInput
                    v-model="address.postal_code"
                    placeholder="Enter postal code"
                    size="lg"
                    :color="isFieldError('postal_code') ? 'error' : isFieldValid('postal_code') ? 'success' : undefined"
                    @blur="validateField('postal_code')"
                  >
                    <template v-if="isFieldValid('postal_code')" #trailing>
                      <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                    </template>
                  </UInput>
                </UFormField>
              </div>

              <!-- 详细地址 -->
              <UFormField label="Street Address" :error="errors.detail_address" required>
                <UInput
                  v-model="address.detail_address"
                  placeholder="Enter street address"
                  size="lg"
                  icon="i-lucide-map-pin"
                  :color="isFieldError('detail_address') ? 'error' : isFieldValid('detail_address') ? 'success' : undefined"
                  @blur="validateField('detail_address')"
                >
                  <template v-if="isFieldValid('detail_address')" #trailing>
                    <UIcon name="i-lucide-check-circle-2" class="w-5 h-5 text-green-500" />
                  </template>
                </UInput>
              </UFormField>
            </div>
          </UForm>
        </UCard>

        <!-- 支付方式 -->
        <UCard>
          <template #header>
            <div class="flex items-center gap-2">
              <UIcon name="i-lucide-credit-card" class="w-5 h-5 text-primary-600" />
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
                Payment Method
              </h2>
            </div>
          </template>

          <div class="space-y-3">
            <UCard
              v-for="option in payTypeOptions"
              :key="option.value"
              :class="[
                'cursor-pointer transition-all hover:shadow-md',
                selectedPayType === option.value
                  ? 'ring-2 ring-primary-500 bg-primary-50/50 dark:bg-primary-900/20'
                  : 'hover:border-gray-300 dark:hover:border-gray-600'
              ]"
              @click="selectedPayType = option.value"
            >
              <div class="flex items-center justify-between p-4">
                <div class="flex items-center gap-4">
                  <div class="relative">
                    <div
                      class="w-5 h-5 rounded-full border-2 flex items-center justify-center transition-all"
                      :class="{
                        'border-primary-500 bg-primary-500': selectedPayType === option.value,
                        'border-gray-300 dark:border-gray-600': selectedPayType !== option.value
                      }"
                    >
                      <UIcon
                        v-if="selectedPayType === option.value"
                        name="i-lucide-check"
                        class="w-3 h-3 text-white"
                      />
                    </div>
                  </div>
                  <span class="text-base font-semibold text-gray-900 dark:text-white">
                    {{ option.label }}
                  </span>
                </div>
                <div v-if="option.value === 1">
                  <UIcon name="i-simple-icons-paypal" class="w-10 h-10 text-[#0070ba]" />
                </div>
              </div>
            </UCard>
          </div>
        </UCard>

        <!-- 提交按钮 -->
        <UCard class="bg-gradient-to-r from-primary-50 to-primary-100 dark:from-primary-900/20 dark:to-primary-800/20 border-primary-200 dark:border-primary-800">
          <UButton
            color="primary"
            size="xl"
            block
            :loading="loading"
            :disabled="loading"
            icon="i-lucide-lock"
            class="font-bold shadow-lg hover:shadow-xl transition-shadow"
            @click="handleCheckout"
          >
            {{ loading ? 'Processing...' : 'Place Your Order Now' }}
          </UButton>
        </UCard>
      </div>

      <!-- 右侧：订单摘要 -->
      <div class="lg:col-span-1">
        <UCard class="sticky top-24 shadow-lg">
          <template #header>
            <div class="flex items-center gap-2">
              <UIcon name="i-lucide-receipt" class="w-5 h-5 text-primary-600" />
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
                Order Summary
              </h2>
            </div>
          </template>

          <div class="space-y-4">
            <!-- 商品列表 -->
            <div class="space-y-4">
              <div
                v-for="item in products"
                :key="item.id"
                class="flex gap-3 pb-4 border-b border-gray-200 dark:border-gray-700 last:border-0"
              >
                <!-- 商品图片 -->
                <div class="flex-shrink-0 relative">
                  <div class="w-16 h-16 rounded-lg overflow-hidden bg-gray-100 dark:bg-gray-800">
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
                  <UBadge
                    color="error"
                    variant="solid"
                    class="absolute -top-2 -right-2 min-w-5 h-5 px-1 text-xs flex items-center justify-center"
                  >
                    {{ item.quantity }}
                  </UBadge>
                </div>

                <!-- 商品信息 -->
                <div class="flex-1 min-w-0">
                  <div class="text-sm font-medium text-gray-900 dark:text-white line-clamp-2 mb-1">
                    {{ item.title }}
                  </div>
                  <div
                    v-if="item.sku_title && item.sku_title.length > 0"
                    class="text-xs text-gray-500 dark:text-gray-400 mb-2"
                  >
                    {{ item.sku_title }}
                  </div>
                  <div class="flex items-baseline gap-2">
                    <span class="text-base font-semibold text-[#ff4d4f] dark:text-red-400">
                      {{ formatPrice(item.price) }}
                    </span>
                    <span
                      v-if="Number(item.price) < Number(item.original_price)"
                      class="text-xs text-gray-500 dark:text-gray-400 line-through"
                    >
                      {{ formatPrice(item.original_price) }}
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- 分隔线 -->
            <div class="border-t border-gray-200 dark:border-gray-700 pt-4 space-y-3">
              <!-- 商品小计 -->
              <div class="flex justify-between items-center text-sm">
                <span class="text-gray-600 dark:text-gray-400">Items Subtotal</span>
                <span class="text-gray-900 dark:text-white font-medium">
                  {{ formatPrice(total_amount) }}
                </span>
              </div>

              <!-- 运费 -->
              <div class="flex justify-between items-center text-sm">
                <span class="text-gray-600 dark:text-gray-400">Freight</span>
                <span class="text-gray-900 dark:text-white font-medium">
                  {{ formatPrice(freight) }}
                </span>
              </div>

              <!-- 总计 -->
              <div class="border-t border-gray-200 dark:border-gray-700 pt-3">
                <div class="flex justify-between items-center">
                  <span class="text-lg font-semibold text-gray-900 dark:text-white">Total</span>
                  <span class="text-2xl font-bold text-[#ff4d4f] dark:text-red-400">
                    {{ formatPrice(total_amount + Number(freight)) }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </UCard>
      </div>
    </div>

    <!-- 空购物车状态 -->
    <UEmpty
      v-else-if="status === 'success' && (!products || products.length === 0)"
      icon="i-lucide-shopping-cart"
      title="Your cart is empty"
      description="Add some items to your cart before checkout"
    >
      <template #actions>
        <UButton
          color="primary"
          size="lg"
          icon="i-lucide-arrow-left"
          @click="router.push('/cart')"
        >
          Go to Cart
        </UButton>
      </template>
    </UEmpty>
  </div>
</template>

<style scoped>
.checkout-page {
  width: 100%;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
