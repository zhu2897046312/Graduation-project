import api from '../api'
import type { CartItem } from '../api/type'

/**
 * 共享的购物车 composable
 * 用于在多个页面中共享购物车列表和数量
 */
export const useCartShared = () => {
  // 购物车数量状态（全局共享）
  const cartNum = useState<number>('cartNum', () => 0)

  /**
   * 获取购物车列表
   * @param key - useAsyncData 的 key，用于区分不同的调用场景
   * @param options - useAsyncData 的选项，包括 lazy、watch 等
   */
  const useCartList = (
    key: string = 'cart-list', 
    options?: { 
      lazy?: boolean
      watch?: any[]
      onError?: (error: unknown) => void
    }
  ) => {
    return useAsyncData<CartItem[]>(key, async () => {
      try {
        const res = await api.shop.cart.list()
        // API 返回的是 CartListResponse 格式 { list: CartItem[], total: number }
        const list = res.list || []
        
        // 自动更新购物车数量
        const count = list.reduce((total, item) => total + item.quantity, 0)
        cartNum.value = count
        
        return list
      } catch (error) {
        console.error('Failed to fetch cart:', error)
        // 出错时重置数量
        cartNum.value = 0
        
        // 如果提供了错误处理回调，调用它
        if (options?.onError) {
          options.onError(error)
        }
        
        return []
      }
    }, options)
  }

  /**
   * 手动更新购物车数量
   * @param list - 购物车列表
   */
  const updateCartNum = (list: CartItem[]) => {
    const count = list.reduce((total, item) => total + item.quantity, 0)
    cartNum.value = count
  }

  /**
   * 刷新所有购物车数据
   * 这个方法会刷新所有使用 useCartList 的页面
   */
  const refreshAllCartData = async () => {
    try {
      const res = await api.shop.cart.list()
      const list = res.list || []
      
      // 更新购物车数量
      const count = list.reduce((total, item) => total + item.quantity, 0)
      cartNum.value = count
      
      // 刷新所有使用 useAsyncData 的购物车列表
      await refreshNuxtData('cart-list')
      await refreshNuxtData('checkout-cart')
      
      return list
    } catch (error) {
      console.error('Failed to refresh cart:', error)
      cartNum.value = 0
      return []
    }
  }

  return {
    cartNum: readonly(cartNum),
    useCartList,
    updateCartNum,
    refreshAllCartData
  }
}
