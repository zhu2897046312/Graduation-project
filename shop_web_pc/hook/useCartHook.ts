import { useDialog } from "naive-ui"
import api from "~/api"

export default function useCart() {
  const dialog = useDialog()
  const cartNum = useState('cartNum', () => 0)
  const showLoginModal = useState('showLoginModal', () => false)

  const router = useRouter()

  const addCart = async (product_id: number, sku_id: number, quantity: number) => {
    try {
      const res = await api.shop.cart.act({
        product_id,
        sku_id,
        quantity,
        add: true
      })
      console.debug('Add to cart success', res)
      dialog.success({
        title: 'Added to Cart',
        content: 'Proceed to checkout now?',
        positiveText: 'Checkout',
        negativeText: 'Continue Shopping',
        draggable: true,
        onPositiveClick: () => {
          router.push('/cart')
        },
      })
      return await getList()
    } catch (error: any) {
      dialog.warning({
        title: 'Add to Cart Failed',
        content: error.message,
        positiveText: 'OK',
        draggable: true,
        onPositiveClick: () => {
          if (error.message === 'Please login first') {
            showLoginModal.value = true
          }
        }
      })
      return Promise.reject(error)
    }
  }

  const getList = async () => {
    try {
      const res = await api.shop.cart.list()
      console.debug('Get cart list success', res)
      let count = 0
      res.forEach((it: any) => {
        count += it.quantity
      })
      cartNum.value = count
      return res
    } catch (error: any) {
      return []
    }
  }

  return {
    addCart,
  }

}