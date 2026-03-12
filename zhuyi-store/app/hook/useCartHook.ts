import api from "../api"
import { useCartShared } from "../composables/useCartShared"

export default function useCart() {
  const toast = useToast()
  const router = useRouter()
  const showLoginModal = useState('showLoginModal', () => false)
  
  // 使用共享的购物车 composable
  const { refreshAllCartData } = useCartShared()

  const addCart = async (product_id: number, sku_id: number, quantity: number) => {
    try {
      const res = await api.shop.cart.act({
        product_id,
        sku_id,
        quantity,
        add: true
      })
      
      // 显示成功通知
      toast.add({
        title: 'Added to Cart',
        description: 'Item has been added to your cart successfully',
        color: 'success'
      })
      
      // 刷新所有购物车数据（包括数量和列表）
      await refreshAllCartData()
      
      return res
    } catch (error: unknown) {
      const errorMessage = error instanceof Error ? error.message : 'Failed to add item to cart. Please try again.'
      
      // 显示错误通知
      toast.add({
        title: 'Add to Cart Failed',
        description: errorMessage,
        color: 'error'
      })
      
      // 如果是需要登录的错误，跳转到登录页
      if (errorMessage === 'Please login first') {
        showLoginModal.value = true
        // 延迟跳转，让用户先看到错误提示
        setTimeout(() => {
          router.push('/login')
        }, 1500)
      }
      
      return Promise.reject(error)
    }
  }

  return {
    addCart
  }
}