<script setup lang="ts">
import { NForm, NFormItem, NInput, NSelect, NCard, NButton, useMessage, useDialog } from 'naive-ui';
import type { FormRules } from 'naive-ui'
import addressJson from '../../data/address.json'
import api from '../../api';
setPageLayout('minilayout')
useSeoMeta({
  title: 'checkout',
  description: 'checkout',
  keywords: 'checkout'
})

const selectedPayType = ref(1) 
const message = useMessage()
const formRef = ref<any>(null)
const address = useState<any>('address', () => {
  return {
    first_name: '',
    last_name: '',
    email: '',
    country: null,
    phone: '',
    province: null,
    city: '',
    region: '',
    detail_address: '',
    postal_code:''
  }
})
const addressRef = useState<any>('addressRef', () => null)
const checkoutRules = ref<FormRules>({
    email: [
      {
        required: true,
        message: 'Email is required'
      },
      {
        pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
        message: 'Please enter a valid email address'
      }
    ],
    first_name: [{
      required: true,
      message: 'First name is required'
    }],
    last_name: [{
      required: true,
      message: 'Last name is required'
    }],
    phone: [{
      required: true,
      message: 'Phone is required'
    }],
    country: [{
      required: true,
      message: 'Country is required'
    }],
    province: [{
      required: true,
      message: 'region is required'
    }],
    postal_code: [{
      required: true,
      message: 'Postal code is required'
    }],
    detail_address: [{
      required: true,
      message: 'Address is required'
    }],
    city: [{
      required: true,
      message: 'City is required'
    }],
})
const payTypeOptions = [
  { value: 1, label: 'PayPal' },
  // { value: 2, label: 'Alipay (支付宝)' },
  // { value: 3, label: 'WeChat Pay (微信支付)' },
  // { value: 4, label: 'Credit Card (银行卡)' },
  // { value: 5, label: 'Cash on Delivery (货到付款)' }
]

const countryOptions = computed(() => {
  return addressJson.countries.map(c => ({
    label: c.full_name,  // 显示用
    value: c.name,       // 提交用
    raw: c              // 保留完整数据
  }))
})
const provinceOptions = computed(() => {
  if (!address.value.country) return []
  
  const selectedCountry = addressJson.countries.find(
    c => c.name === address.value.country
  )
  
  return selectedCountry?.provinces?.map(p => ({
    label: p,
    value: p
  })) || []
})


const { data: products, status } = await useAsyncData(async() =>{
  try {
    const res = await api.shop.cart.list()
    console.log("checkout page res : ",res)
    return res.list
  } catch (err) {
    console.error("购物车数据加载失败:", err)
    throw createError({ 
      statusCode: 400, 
      message: "无法加载购物车数据" 
    })
  }
})

const { data: freight } = await useAsyncData('freight', async() =>{
  return await api.shop.market.freight()
})

const total_amount = computed(() => {
  let total = 0
  if (products.value) {
    products.value.forEach((item: any) => {
      total += Number(item.pay_amount)
    })
  }
  return total
})

const formatPrice = (price: unknown) => {
  // 转换为数字
  const num = Number(price)
  // 检查是否为有效数字
  return !isNaN(num) ? `$${num.toFixed(2)}` : '$0.00'
}
const handleValidate = async () => {
  console.log(formRef.value)
  formRef.value.validate((errors: any) => {
    console.log(errors)
    if (!errors) {
      handleCheckout()
    } else {
      window.scrollTo(0, 70);
    }
  })
}
const handleCheckout = async () => {
  const payload = {
    product_items: products.value.map((item: any) => {
       return {
          product_id: item.product_id,
          quantity: item.quantity,
          sku_id: item.sku_id,
       }
    }),
    pay_type: selectedPayType.value,
    first_name: address.value.first_name,
    last_name: address.value.last_name,
    email: address.value.email,
    phone: address.value.phone,
    province: address.value.province,
    country: address.value.country,
    city: address.value.city,
    // region: address.value.region,
    postal_code: address.value.postal_code,
    detail_address: address.value.detail_address,
  }
  const ins = message.loading('Processing...', { duration: 0 })

  try {
    const res = await api.shop.order.create(payload)

    const paymentRes = await api.shop.order.getPaymentUrl({
      orderId: res, 
      payType: selectedPayType.value.toString()
    })
    
    window.location.href = paymentRes.approveUrl;
  } catch (error: any) {
    message.error(error.message)
  } finally {
    ins.destroy()
  }
}

watch(() => address.value.country, (newCountry) => {
  address.value.province = '' // 切换国家时清空省份选择
})

</script>

<template>
  <div>
    <div class="checkout">
      <div class="checkout-box">
        <ClientOnly>
        <NCard style="width: 723px;" :bordered="false" >
          <h1 class="text-xl font-bold mb-4">Shipping Information</h1>
          <NForm size="large" ref="formRef" :rules="checkoutRules" :model="address">
            <NFormItem label="Email" path="email">
              <div class="flex flex-col gap-2 w-full">
                <NInput v-model:value="address.email" placeholder="Enter email address" />
                <span class="text-sm text-gray-500 pt-2 px-1">Email me with news and offers</span>
              </div>
            </NFormItem>
            <div class="flex items-center gap-2 w-full">

              <NFormItem class="flex-1" label="Last Name" path="last_name">
                <NInput v-model:value="address.first_name" placeholder="Enter last name"/>
              </NFormItem>
              <NFormItem class="flex-1" label="First Name" path="first_name">
                <NInput v-model:value="address.last_name" placeholder="Enter first name" />
              </NFormItem>
            </div>
            <NFormItem label="Phone Number" path="phone">
              <NInput v-model:value="address.phone" placeholder="Enter phone number"/>
            </NFormItem>
            
            <NFormItem label="Country" path="country">
              <NSelect 
                v-model:value="address.country" 
                placeholder="Select Country" 
                :options="countryOptions"
                label-field="label"
                value-field="value"
                clearable
              />
            </NFormItem>
            <NFormItem label="Region" path="province">
              <NSelect 
              v-model:value="address.province" 
              placeholder="Select region/province" 
              :options="provinceOptions"
              :disabled="!address.country"
              clearable
              />
            </NFormItem>
            <NFormItem label="City" path="city">
              <NInput v-model:value="address.city" placeholder="Enter city"/>
            </NFormItem>
            <NFormItem label="Address" path="detail_address">
              <NInput v-model:value="address.detail_address" placeholder="Enter address"/>
            </NFormItem>
            <NFormItem label="ZIP Code" path="postal_code" >
              <NInput v-model:value="address.postal_code" placeholder="Enter postal_code"/>
            </NFormItem>
          </NForm>
        </NCard>
      </ClientOnly> 
      <ClientOnly>
        <NCard style="width: 723px;" :bordered="false" >
          <h1 class="text-xl font-bold mb-4">Payment</h1>
          <div class="pay-type-item flex items-center justify-center">
            <div class="space-y-2 w-full">
              <label 
                v-for="option in payTypeOptions" 
                :key="option.value"
                class="flex items-center p-3 border rounded-lg cursor-pointer"
                :class="{
                  'border-blue-500 bg-blue-50': selectedPayType === option.value,
                  'border-gray-200': selectedPayType !== option.value
                }"
              >
                <input
                  type="radio"
                  v-model="selectedPayType"
                  :value="option.value"
                  class="hidden"
                >
                <div class="flex items-center">
                  <div class="w-5 h-5 rounded-full border-2 mr-3 relative"
                    :class="{
                      'border-blue-500': selectedPayType === option.value,
                      'border-gray-300': selectedPayType !== option.value
                    }"
                  >
                    <div v-if="selectedPayType === option.value" 
                      class="absolute inset-1 bg-blue-500 rounded-full">
                    </div>
                  </div>
                  <span>{{ option.label }}</span>
                </div>
                <img v-if="option.value === 1" class="h-8  ml-auto" src="~/assets/paypal.png">
              </label>
            </div>
          </div>
        </NCard>
      </ClientOnly> 
        <div style="width: 700px;padding-right: 10px;" class="mt-2 flex justify-center mb-5">
          <NButton block size="large" type="info" @click.prevent="handleValidate">Place your order now</NButton>
        </div>
      </div>
      <div class="checkout-info">
        <div class="checkout-summary">
          <NCard style="width: 485px;" >
            <h1 class="text-xl font-bold mb-4">Order Summary</h1>
            <div class="product-list">
              <div class="product-item" v-for="item in products">
                <div class="product-item-thumb">
                  <img :src="item.thumb" />
                  <span>{{ item.quantity }}</span>
                </div>
                <div class="product-item-info">
                  <div class="product-item-title">{{ item.title }}</div>
                  <div class="product-item-sku" v-if="item.sku_title && item.sku_title.length > 0">{{ item.sku_title }}</div>
                </div>
                <div class="product-item-amount">
                  <div class="flex  items-baseline gap-1.5">
                      <div class="discounted-price text-[#FF5000] font-medium text-lg leading-tight">
                        {{ formatPrice(item.pay_amount) }}
                      </div>
                      <div v-if="Number(item.pay_amount) < Number(item.total_amount)" class="original-price text-gray-500 text-base line-through leading-tight">
                        {{ formatPrice(item.total_amount) }}
                      </div>
                    </div>
                </div>
              </div>
            </div>

            <div class="cell_item">
              <div class="cell_item_label">Items Subtotal</div>
              <div class="cell_item_value">
                <div class="price">{{ total_amount.toFixed(2) }}</div>
              </div>
            </div>
            <div class="cell_item">
              <div class="cell_item_label">Freight</div>
              <div class="cell_item_value">
                <div class="price">{{ freight }}</div>
              </div>
            </div>
            <div class="cell_item">
              <div class="cell_item_label">Total</div>
              <div class="cell_item_value">
                <div class="price">{{ (total_amount + Number(freight)).toFixed(2) }}</div>
              </div>
            </div>
          </NCard>
        </div>
      </div>
    </div>
    <div v-if="status == 'success'">
      <ClientOnly>
        <UserAddressModal ref="addressRef" @selected="(e : any) => address = e" />
      </ClientOnly> 
    </div>
  </div>
</template>


<style lang="css" scoped>
.checkout {
  display: grid;
  view-transition-name: shell-content;
  grid-area: shell-content;
  grid-template-areas: "main order-summary";
  grid-template-columns: minmax(min-content, calc(50% + calc( calc( 66rem - 52rem ) / 2 ))) 1fr;
  border-top: 1px solid #dedede;
}
.checkout-box {
  grid-area: main;
  display: flex;
  flex-direction: column;
  align-items: self-end;
  height: 100%;
  border-right: 1px solid #dedede;
}
.checkout-info {
  background-color: #f5f5f5;
  display: block;
  grid-area: order-summary;
}
.checkout-summary {
  position: sticky;
  padding: 10px;
  width: 100%;
  right: auto;
  left: auto;
  top: 0;
  bottom: 0;
}

.product-list .product-item {
  padding: 10px 0;
  display: flex;
  align-items: center;
  gap: 5px;
}
.product-item-thumb {
  width: 74px;
  height: 74px;
  background-color: #fff;
  align-items: center;
  justify-content: center;
  border-radius: 5px;
  border: 1px solid #f5f5f5;
  position: relative;
}
.product-item-thumb img {
  max-width: 100%;
  max-height: 100%;
  display: block;
}
.product-item-thumb span {
  position: absolute;
  right: -10px;
  top: -10px;
  display: block;
  border-radius: 50%;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  background-color: #f50;
  width: 20px;
  height: 20px;
}
.product-item-info {
  padding: 0 10px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.product-item-title {
  font-size: 16px;
  color: var(--font-text-color);
}
.product-item-sku {
  font-size: 14px;
  color: var(--font-text-color);
  opacity: 0.8;
}
.product-item-amount .price {
  font-size: 20px;
}
.pay-type-item {
  padding: 10px;
  border: 1px solid #f5f5f5;
  border-radius: 5px;;
}
</style>