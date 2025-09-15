<script setup lang="ts">
import api from '../../../api';
import type { ProductInfo } from '../../../api/type';
import useCart from '../../../hook/useCartHook';
import { NButton, NInputNumber, NTag, NIcon, NEmpty} from 'naive-ui';
// import { CartOutline } from '@vicons/ionicons5'

const route = useRoute()
const productId = route.params.id as any
const activeTab = ref('details') 

const price = useState<number>('price', () => 0)
const original_price = useState<number>('original_price', () => 0)
const current_sku_id = useState<number>('current_sku_id', () => 0)
const defaultSkuCode = useState<string>('defaultSkuCode', () => '') // 新增：存储默认SKU code
// 商品的数量
const quantity = useState<number>('quantity', () => 1)
  
const { data: info, status } = await useAsyncData(`product-${productId}`,async() => {

  try {
    console.log('Fetching product info for ID:', productId)
    const res = await api.shop.product.info(productId) 
    console.log('API response:', res) 
    price.value = res.price
    original_price.value = res.original_price
    quantity.value = 1
    current_sku_id.value = 0
    const defaultSku = res.sku_list.find((sku: any) => sku.default_show === 1)
    if (defaultSku) {
      defaultSkuCode.value = defaultSku.sku_code
      price.value = defaultSku.price
      original_price.value = defaultSku.original_price
      current_sku_id.value = defaultSku.id
    }

    return {
      ...res,
      // Set default SEO values if null
      seo_title: res.seo_title || `${res.title} | Shop Name`,
      seo_keyword: res.seo_keyword || `${res.title}, ${res.tags?.map((t : any) => t.title).join(', ') || ''}`,
      seo_description: res.seo_description || `Buy ${res.title} at best price. ${res.property_list?.map((p : any) => `${p.title}: ${p.value}`).join('. ') || ''}`
    };
  } catch (_) {
    return false
  }
})

const { data: recommendedProducts } = await useAsyncData(
  `recommended-products-${productId}`,
  async () => {
    try {
      // 获取当前商品的所有标签ID
      console.log('Fetching product recommendedInfo for ID:',productId)
      const tagIds = info.value.tags?.map((tag: any) => tag.id) || []
      console.log('Fetching tags list',tagIds)
      // 随机选择最多4个标签（如果标签数超过4个）
      const randomTagIds = tagIds
        .sort(() => 0.5 - Math.random())
        .slice(0, 4)
      
      // 获取这些标签下的热门商品 --- tags 为空时 会获取 整个list列表内容 进行展示前面4个
      const res: any = await api.shop.product.list({
        tag_ids: randomTagIds.join(','),
        page_size: 4,
        sort_by: 'sales',
        sort_order: 'desc'
      })
      console.log(res)
      return res.list || []
    } catch (error) {
      console.error('Failed to fetch recommended products:', error)
      return []
    }
  }
)

if (!info.value) {
  throw createError({ statusCode: 404, message: 'Page not found' })
}


const { addCart } = useCart()

/**
 * 添加购物车
 */
const handleAddCart = async () => {
  try {
    await addCart(info.value.id, current_sku_id.value , quantity.value)    
  } catch (_) {
    // console.error(error)
  }
}

/**
 * 购物车发生变化
 * @param sku_code 
 */
const handleOnSkuChange = (sku_code: string) => {
  const hit = info.value.sku_list.filter((it: any) => it.sku_code === sku_code)
  console.log(sku_code, hit)
  if (hit.length > 0) {
    price.value = hit[0].price
    original_price.value = hit[0].original_price
    current_sku_id.value = hit[0].id
  } else {
    console.log('没有找到sku')
  }
}


console.log(info.value)
</script>

<template>
  <div v-if="status == 'success'">
    <Title>{{ info.seo_title }}</Title>
    <Meta name="keywords" :content="info.seo_keyword" />
    <Meta name="description" :content="info.seo_description" />

    <div class="container product-header">
      <div class="product-image">
        <ProductGrllery
          :grllery="info.picture_gallery && info.picture_gallery.length > 0 ? info.picture_gallery : [info.picture]"
        />
      </div>
      <div class="product-info">
        <h1 class="product-title">{{ info.title }}</h1>
        <div class="tags-container">
          <NuxtLink class="tag-item" v-for="item in info.tags" :to="`/tag/${item.code}`">
            <NTag :bordered="false" type="info">{{ item.title }}</NTag>
          </NuxtLink>
        </div>
        <div class="product-price-container">
          <span class="price">{{ price }}</span>
          <span class="original-price">${{ original_price }}</span>
        </div>
        <div class="shipping-info">Tax included. Shipping cost calculated at checkout.</div>
        <div v-if="info.property_list.length > 0" class="property-container">
          <div v-for="item in info.property_list" class="property-item">
            <span class="property-item-label">{{ item.title }}:</span>
            <span class="property-item-value">{{ item.value }}</span>
          </div>
        </div>
        <div class="sku-container">
          <ProductSkuBox 
          :list="info.sku_config" 
          :defaultSelected="defaultSkuCode"
          @change="handleOnSkuChange" 
          />
        </div>
        <div class="cart-action-container">
          <div class="quantity-control-wrapper">
            <NInputNumber v-model:value="quantity" :min="1" :max="9" />
          </div>
          <NButton type="primary" class="add-to-cart-btn" @click.passive="handleAddCart">
            <!-- <template #icon>
              <NIcon :component="CartOutline" />
            </template> -->
            ADD TO CART
          </NButton>
        </div>
      </div>
    </div>

    <div class="container flex items-center justify-center gap-2 px-3 pd-4 mt-4 mb-2">
      <div class="w-full">
        <div class="flex border-b border-gray-200">
          <div 
            @click="activeTab = 'details'"
            :class="[
              'px-4 py-2 text-sm font-medium cursor-pointer',
              activeTab === 'details' 
                ? 'border-b-2 border-[#fb7f86] text-[#fb7f86]' 
                : 'text-gray-500 hover:text-gray-700'
            ]"
          >
            Product Details
          </div>
          <div 
            v-if="info.property_list && info.property_list.length > 0"
            @click="activeTab = 'specs'"
            :class="[
              'px-4 py-2 text-sm font-medium cursor-pointer',
              activeTab === 'specs' 
                ? 'border-b-2 border-[#fb7f86] text-[#fb7f86]' 
                : 'text-gray-500 hover:text-gray-700'
            ]"
          >
            Specifications
          </div>
        </div>
      </div>
    </div>
    
    <!-- 根据选项卡显示不同内容 -->
    <div v-if="activeTab === 'details'" class="container product-body">
      <h2>Product Details</h2>
      <div class="blogs_box" v-html="info.content"></div>
    </div>

    <div v-if="activeTab === 'specs'" class="property my-2 px-3">
      <NEmpty 
        v-if="!info.property_list || info.property_list.length === 0"
        description="No specifications available"
        class="py-8"
      >
        <template #icon>
          <NIcon size="48">
            <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32">
              <path d="M30 3.414L28.586 2 2 28.586L3.414 30l2-2H26a2.003 2.003 0 0 0 2-2V5.414zM26 26H7.414l7.793-7.793l2.379 2.379a2 2 0 0 0 2.828 0L22 19l4 3.997zm0-5.832l-2.586-2.586a2 2 0 0 0-2.828 0L19 19.168l-2.377-2.377L26 7.414zM6 22v-3l5-4.997l1.373 1.374l1.416-1.416l-1.375-1.375a2 2 0 0 0-2.828 0L6 16.172V6h16V4H6a2.002 2.002 0 0 0-2 2v16z" fill="currentColor"></path>
            </svg>
          </NIcon>
        </template>
      </NEmpty>
      
      <div v-else class="container property-item px-4">
        <table class="property-table">
          <tr v-for="item in info.property_list">
            <td class="property-name">{{ item.title }}</td>
            <td class="property-value">{{ item.value }}</td>
          </tr>
        </table>
      </div>
    </div>

    <div class="may-like container product-body">
      <h1 class="section-title">You May Also Like</h1>
      <div class="product-grid product_list">
        <ProductCard
          v-for="product in recommendedProducts"
          :key="product.id"
          :product-id="product.id"
          :title="product.title"
          :thumb="product.picture"
          :price="product.price"
          :original-price="product.original_price"
          class="product-card"
        />
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
.product-header {
  display: flex;
  gap: 40px;
  padding: 30px 20px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  margin-top: 20px;
}

.product-image {
  flex: 1;
  min-width: 0;
}

.product-info {
  flex: 1;
  min-width: 0;
  padding: 0 30px;
}

.product-title {
  font-size: 1.6em;
  color: #333;
  padding-bottom: 12px;
  font-weight: 600;
  border-bottom: 1px solid #f0f0f0;
  margin-bottom: 15px;
}

.tags-container {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 8px;
  margin-bottom: 20px;
}

.tag-item:hover {
  transform: translateY(-2px);
  transition: all 0.3s ease;
}

.product-price-container {
  display: flex;
  align-items: baseline;
  gap: 12px;
  margin-bottom: 15px;
  padding: 10px 15px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.product-price-container .price {
  font-size: 26px;
  font-weight: bold;
  color: #ff4d4f;
}

.product-price-container .original-price {
  font-size: 16px;
  color: #999;
  text-decoration: line-through;
}

.shipping-info {
  font-size: 14px;
  color: #666;
  margin-bottom: 20px;
  padding-left: 2px;
}

.property-container {
  margin-bottom: 20px;
  padding: 15px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.property-item {
  display: flex;
  gap: 10px;
  padding-bottom: 8px;
}

.property-item-label {
  font-size: 14px;
  color: #666;
  min-width: 60px;
}

.property-item-value {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.sku-container {
  margin-bottom: 15px;
}

.cart-action-container {
  display: flex;
  align-items: center;
  gap: 15px;
  margin: 15px 0;
}

.quantity-control-wrapper {
  width: 120px;
}

.quantity-control-wrapper :deep(.n-input-number) {
  width: 100%;
}

.quantity-control-wrapper :deep(.n-input-number-base) {
  border: 1px solid #e0e0e0;
  border-radius: 0;
}

.quantity-control-wrapper :deep(.n-input__input) {
  text-align: center;
  font-size: 16px;
  line-height: 44px;
  padding-top: 0;
  padding-bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.quantity-control-wrapper :deep(.n-input-wrapper) {
  display: flex;
  align-items: center;
  border: none;
  border-radius: 0;
}

.quantity-control-wrapper :deep(.n-input) {
  height: 44px;
  border-radius: 0;
  display: flex;
  align-items: center;
}

.quantity-control-wrapper :deep(.n-button) {
  border-radius: 0;
  border: none;
  width: 32px;
  padding: 0;
  font-size: 18px;
}

.add-to-cart-btn {
  flex: 1;
  height: 44px;
  background-color: #ffb6c1 !important;
  border: none !important;
  border-radius: 0 !important;
  font-size: 16px;
  font-weight: 500;
  animation: shake 6s ease-in-out infinite;
}

.add-to-cart-btn:hover {
  background-color: #ff9daa !important;
}

@keyframes shake {
  0%, 83%, 100% { transform: translateX(0); }
  85% { transform: translateX(-5px); }
  87% { transform: translateX(5px); }
  89% { transform: translateX(-5px); }
  91% { transform: translateX(5px); }
  93% { transform: translateX(-5px); }
  95% { transform: translateX(5px); }
  97% { transform: translateX(-5px); }
  99% { transform: translateX(0); }
}

.add-to-cart-btn:focus {
  outline: none;
  box-shadow: 0 0 0 2px #ffb6c1;
}

.product-body {
  margin-top: 40px;
  margin-bottom: 40px;
  padding: 40px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.product-body h2 {
  font-size: 24px;
  font-weight: 600;
  text-align: center;
  color: #333;
  position: relative;
  margin-bottom: 30px;
  padding-bottom: 15px;
}

.product-body h2::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 80px;
  height: 3px;
  background-color: #ff4d4f;
  opacity: 0.7;
}

.blogs_box {
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
  line-height: 1.6;
}
.section-title {
  font-size: 24px;
  font-weight: 600;
  text-align: center;
  color: #333;
  position: relative;
  margin-bottom: 30px;
  padding-bottom: 15px;
}
</style>
