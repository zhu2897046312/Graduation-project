<script setup lang="ts">
import api from '../../api'
import type { ProductDetailData } from '../../api/type'
import useCart from '../../hook/useCartHook'
import { getProductImage } from '../../utils/auth'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

const route = useRoute()
const productId = computed(() => {
  const id = route.params.id
  return typeof id === 'string' ? Number(id) : Number(id) || 0
})
const activeTab = ref('details')

// Tab 配置
const tabItems = computed(() => {
  const items = [
    {
      label: 'Product Details',
      value: 'details'
    }
  ]
  
  if (info.value?.property_list && info.value.property_list.length > 0) {
    items.push({
      label: 'Specifications',
      value: 'specs'
    })
  }
  
  return items
}) 

const price = useState<number>('price', () => 0)
const original_price = useState<number>('original_price', () => 0)
const current_sku_id = useState<number>('current_sku_id', () => 0)
const defaultSkuCode = useState<string>('defaultSkuCode', () => '') // 新增：存储默认SKU code
// 商品的数量
const quantity = useState<number>('quantity', () => 1)
  
const { data: info, status } = await useAsyncData<ProductDetailData>(`product-${productId.value}`, async () => {
  try {
    if (!productId.value || productId.value === 0) {
      throw createError({ statusCode: 404, message: 'Invalid product ID' })
    }
    
    const res = await api.shop.product.info(productId.value)
    
    // 设置初始价格
    price.value = res.price
    original_price.value = res.original_price
    quantity.value = 1
    current_sku_id.value = 0
    
    // 查找默认 SKU
    const defaultSku = res.sku_list?.find((sku) => sku.default_show === 1)
    if (defaultSku) {
      defaultSkuCode.value = defaultSku.sku_code
      price.value = defaultSku.price
      original_price.value = defaultSku.original_price
      current_sku_id.value = defaultSku.id
    }

    return res
  } catch (error) {
    console.error('Failed to fetch product:', error)
    throw createError({ statusCode: 404, message: 'Product not found' })
  }
})

// 推荐产品获取 - 如果没有标签则不获取
const { data: recommendedProducts } = await useAsyncData(
  `recommended-products-${productId.value}`,
  async () => {
    try {
      // 等待主商品数据加载完成
      if (!info.value) {
        return []
      }
      
      // 获取当前商品的所有标签ID，tags 可能为 null
      const tags = info.value.tags || []
      if (tags.length === 0) {
        // 没有标签，直接返回空数组
        return []
      }
      
      const tagIds = tags.map((tag) => tag.id).filter((id) => id != null)
      if (tagIds.length === 0) {
        return []
      }
      
      // 随机选择最多4个标签（如果标签数超过4个）
      const randomTagIds = [...tagIds]
        .sort(() => 0.5 - Math.random())
        .slice(0, 4)
      
      // 获取这些标签下的热门商品
      const res = await api.shop.product.list({
        tag_ids: randomTagIds.join(','),
        page_size: 4,
        sort_by: 'sales',
        sort_order: 'desc'
      })
      
      // 直接返回列表，API 已返回标准化的 ProductListResponse
      return res.list
    } catch (error) {
      console.error('Failed to fetch recommended products:', error)
      return []
    }
  },
  {
    // 延迟执行，等待主商品加载完成
    lazy: true,
    // 监听主商品数据变化
    watch: [() => info.value]
  }
)

// 处理商品图片
const getProductThumb = (picture: string) => {
  return getProductImage(picture)
}

if (!info.value) {
  throw createError({ statusCode: 404, message: 'Page not found' })
}


const { addCart } = useCart()

/**
 * 添加购物车
 */
const handleAddCart = async () => {
  try {
    const productID = info!.value?.id as number
    await addCart(productID, current_sku_id.value , quantity.value)    
  } catch (_) {
    // console.error(error)
  }
}

/**
 * 购物车发生变化
 * @param sku_code 
 */
const handleOnSkuChange = (sku_code: string) => {
  const skuList = info.value?.sku_list || []
  const selectedSku = skuList.find((it) => it.sku_code === sku_code)
  console.log(sku_code, selectedSku)
  if (selectedSku) {
    price.value = selectedSku.price
    original_price.value = selectedSku.original_price
    current_sku_id.value = selectedSku.id
  } else {
    console.log('没有找到sku')
  }
}


// SEO
useHead({
  title: info.value?.seo_title || info.value?.title || 'Product Details',
  meta: [
    { name: 'keywords', content: info.value?.seo_keyword || '' },
    { name: 'description', content: info.value?.seo_description || '' }
  ]
})
</script>

<template>
  <div v-if="status == 'success'">
    <div class="product-header">
      <div class="product-image">
        <ProductGrllery
          :grllery="info?.picture_gallery && info.picture_gallery.length > 0 ? info.picture_gallery : (info?.picture ? [info.picture] : [])"
        />
      </div>
      <div class="product-info">
        <h1 class="product-title">{{ info?.title }}</h1>
        <div v-if="info?.tags && info.tags.length > 0" class="tags-container">
          <NuxtLink 
            v-for="item in info.tags" 
            :key="item.id"
            :to="`/tag/${item.code}`"
            class="tag-item"
          >
            <UBadge :color="'primary'" variant="subtle" size="sm">
              {{ item.title }}
            </UBadge>
          </NuxtLink>
        </div>
        <div class="product-price-container">
          <span class="price">${{ price }}</span>
          <span v-if="original_price > price" class="original-price">${{ original_price }}</span>
        </div>
        <div class="shipping-info">Tax included. Shipping cost calculated at checkout.</div>
        <div v-if="info?.property_list && info.property_list.length > 0" class="property-container">
          <div v-for="item in info?.property_list" class="property-item">
            <span class="property-item-label">{{ item.title }}:</span>
            <span class="property-item-value">{{ item.value }}</span>
          </div>
        </div>
        <div v-if="info?.sku_config && info.sku_config.length > 0" class="sku-container">
          <ProductSkuBox 
            :list="info.sku_config" 
            :defaultSelected="defaultSkuCode"
            @change="handleOnSkuChange" 
          />
        </div>
        <div class="cart-action-container">
          <div class="quantity-control-wrapper">
            <UInputNumber
              v-model="quantity"
              :min="1"
              :max="9"
              size="lg"
            />
          </div>
          <UButton
            color="primary"
            size="lg"
            block
            class="add-to-cart-btn"
            icon="i-lucide-shopping-cart"
            @click="handleAddCart"
          >
            ADD TO CART
          </UButton>
        </div>
      </div>
    </div>

      <div class="flex items-center justify-center gap-2 mt-8 mb-4">
        <div class="w-full">
          <UTabs v-model="activeTab" :items="tabItems" class="w-full" />
        </div>
      </div>
    
      <!-- 根据选项卡显示不同内容 -->
      <UCard v-if="activeTab === 'details'" class="product-body">
        <template #header>
          <h2>Product Details</h2>
        </template>
        <div class="blogs_box" v-html="info?.content"></div>
      </UCard>

      <div v-if="activeTab === 'specs'" class="property my-2">
        <UEmpty 
          v-if="!info?.property_list || info.property_list.length === 0"
          icon="i-lucide-file-question"
          title="No specifications available"
          description="This product does not have specifications yet"
          class="py-8"
        />
        
        <UCard v-else class="property-item">
          <template #header>
            <h3 class="text-lg font-semibold">Specifications</h3>
          </template>
          <div class="space-y-2">
            <div 
              v-for="item in info.property_list" 
              :key="item.title"
              class="flex gap-4 py-2 border-b border-gray-200 dark:border-gray-700 last:border-0"
            >
              <span class="property-name font-medium text-gray-600 dark:text-gray-400 min-w-[120px]">
                {{ item.title }}:
              </span>
              <span class="property-value text-gray-900 dark:text-gray-100">
                {{ item.value }}
              </span>
            </div>
          </div>
        </UCard>
      </div>

      <div class="may-like product-body mt-8">
        <h1 class="section-title">You May Also Like</h1>
        <div v-if="recommendedProducts && recommendedProducts.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mt-6">
          <ProductCard
            v-for="product in recommendedProducts"
            :key="product.id"
            :product-id="product.id"
            :title="product.title"
            :thumb="getProductThumb(product.picture)"
            :price="product.price"
            :original-price="product.original_price"
          />
        </div>
        <UEmpty
          v-else
          icon="i-lucide-package"
          title="No recommended products"
          description="There are no recommended products available at the moment"
          class="py-8"
        />
      </div>
  </div>
</template>

<style lang="css" scoped>
.product-header {
  display: flex;
  gap: 40px;
  padding: 30px;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  margin-top: 20px;
}

.dark .product-header {
  background-color: rgb(31 41 55);
}

@media (max-width: 768px) {
  .product-header {
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }
}

.product-image {
  flex: 1;
  min-width: 0;
}

.product-info {
  flex: 1;
  min-width: 0;
  padding: 0;
}

@media (max-width: 768px) {
  .product-info {
    padding: 0;
  }
}

.product-title {
  font-size: 1.6em;
  color: #333;
  padding-bottom: 12px;
  font-weight: 600;
  border-bottom: 1px solid #f0f0f0;
  margin-bottom: 15px;
}

.dark .product-title {
  color: white;
  border-bottom-color: rgb(55 65 81);
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

.add-to-cart-btn {
  flex: 1;
}

.product-body {
  margin-top: 40px;
  margin-bottom: 40px;
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

.dark .product-body h2 {
  color: white;
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
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  line-height: 1.6;
}

.dark .blogs_box {
  background-color: rgb(31 41 55);
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

.dark .section-title {
  color: white;
}
</style>