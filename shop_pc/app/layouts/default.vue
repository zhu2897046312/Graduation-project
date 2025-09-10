<script setup lang="ts">
import api from '../../api';
import {  NBreadcrumb, NBreadcrumbItem, useDialog, useMessage} from 'naive-ui'
const route = useRoute();
const router = useRouter();
const accessToken = useCookie('accessToken');
const dialog = useDialog();
const message = useMessage();
const cartNum = useState('cartNum', () => 0)
const showLoginModal = useState('showLoginModal', () => false)
const loginModalRef = ref()
const search = useState('search', () => route.query.q as string || '')
// const showCart = useState('showCart', () => false)
// const { data: footer_list } = await useAsyncData('footer_list', async () => {
//   const res = await api.blogs.recommend.list({code: 'FOOTER_LIST', page_no: 1, page_size: 50})
//   return res
// })
const{data: footer_list} = await useAsyncData('footer_list_',async ()=> {
  const res = await api.blogs.document.list({});
  
  return res.map((item:any) => ({
    ...item,
    link: `/blogs/${item.code}`  // 根据实际路由结构修改
  }))
})
const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})
const { data: breadcrumb, refresh: breadcrumbRefresh } = await useAsyncData('breadcrumb', async () => {
  if (route.path === '/') {
    return []
  }
  console.log('router', route.name)
  let out: any[] = [];
  if (route.name === 'collections-code') {
    out = await api.shop.market.breadcrumb({ mode: '1', category_code: route.params.code as string })
  } else if (route.name === 'product-id') {
    out = await api.shop.market.breadcrumb({ mode: '2', product_id: route.params.id as string })
  }
  return out
})
const {data: user, refresh} = await useAsyncData('user-info', async () => {
  if (accessToken.value) {
    const res = await api.shop.user.currentUser()
    return res
  }
  return null
})

const handleLogout = () => {
  dialog.warning({
    title: 'Confirm Logout',
    content: 'Are you sure you want to log out?',
    positiveText: 'Logout',
    negativeText: 'Cancel',
    onPositiveClick: () => {
      accessToken.value = null
      cartNum.value = 0
      refresh()
      message.success('Logged out successfully');
      router.push("/")
    }
  });
}

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Enter' && search.value.length > 0) {
    navigateTo({path: '/search', query: {q: search.value}})
  }
}

const handleAddCart = () => {
  navigateTo('/cart')
}

const handleLogin = async (token: string) => {
  console.log('login event received');
  accessToken.value = token
  setTimeout(async () => {
    refresh()
    const res = await api.shop.cart.list() as []
    let count = 0
    res.forEach((it: any) => {
      count += it.quantity
    })
    cartNum.value = count
  }, 300);
}

callOnce(async () => {
  if (accessToken.value) {
    const res = await api.shop.cart.list() 
    console.log("cart list", res)
    let count = 0
    res.list.forEach((it: any) => {
      count += it.quantity
    })
    cartNum.value = count
  }
})

watch(showLoginModal, (newVal) => {
  if (newVal) {
    loginModalRef.value.useOpen()
  }
})

watch(() => route.fullPath, () => {
  breadcrumbRefresh()
}, { deep: true })


</script>

<template>
  <div class="h-full">
    <header class="bg-[#FB7F86] py-2 ">
      <div class="container  flex items-center justify-between ">
        <div class="logo flex items-end w-1/3">
          <NuxtLink to="/">
            <img 
              src="~/assets/logo.png" 
              class="h-16 object-contain"
            />
          </NuxtLink>
        </div>

        <div class="flex flex-col w-2/3 mt-4">
          <div class="flex items-center justify-end gap-1 w-full">
            <div class="relative w-1/3">
              <input
                v-model="search"
                @keydown="handleKeydown"
                placeholder="Search"
                class="w-full px-4 py-2 rounded-md border border-gray-300 focus:border-white focus:ring-1 focus:ring-[#f4b3c2] focus:outline-none"
              />
              <span class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="11" cy="11" r="8"></circle>
                  <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
                </svg>
              </span>
            </div>

            
            <!-- 购物车按钮（原生实现） -->
            <div class="relative ml-1">
              <button
                @click="handleAddCart"
                class="h-12 w-12 flex items-center justify-center
                      bg-transparent hover:bg-white/10 active:bg-white/20
                      transition-colors duration-200 rounded-md"
              >
                <img src="~/assets/cart.png" alt="cart" class="w-8 h-8 object-contain"/>
              </button>
              
              <!-- 购物车徽标 -->
              <span 
                v-if="cartNum > 0"
                class="absolute -top-2 -right-2 min-w-[20px] h-5 px-1 rounded-full
                      bg-red-500 text-white text-xs font-bold flex items-center justify-center
                      border-2 border-white"
              >
                {{ cartNum > 99 ? '99+' : cartNum }}
              </span>
            </div>
          </div>

          <!-- BaseNav below -->
          <div class="flex items-center justify-between mt-8 w-full">
            <div class="flex-1 min-w-0">
              <BaseNav />
            </div>
            
            <div v-if="!accessToken" class="flex items-center ml-4 space-x-2 flex-shrink-0 text-white relative z-20">
              <!-- 登录按钮 -->
              <button
                @click="navigateTo('/account/login')"
                class="ml-1 h-12 w-24 px-3 flex items-center justify-center
                      bg-transparent hover:bg-white/10 active:bg-white/20
                      transition-colors duration-200 rounded-md
                      border-none focus:outline-none
                      text-white hover:text-white active:text-white"
              >
                Sign In
              </button>
              
              <!-- 注册按钮 -->
              <button
                @click="navigateTo('/account/register')"
                class="ml-1 h-12 w-24 px-3 flex items-center justify-center
                      bg-transparent hover:bg-white/10 active:bg-white/20
                      transition-colors duration-200 rounded-md
                      border-none focus:outline-none
                      text-white hover:text-white active:text-white"
              >
                Sign Up
              </button>
            </div>

            <div v-else class="flex items-center ml-4 space-x-2 flex-shrink-0 text-white relative z-20">
              <!-- 账户按钮 -->
              <button
                @click="navigateTo('/account/')"
                class="ml-1 h-12 w-24 px-3 flex items-center justify-center
                      bg-transparent !hover:bg-white/10 active:bg-white/20
                      transition-colors duration-200 rounded-md
                      border-none focus:outline-none
                      text-white hover:text-white active:text-white"
              >
                Account
              </button>
              
              <!-- 登出按钮 -->
              <button
                @click="handleLogout"
                class="ml-1 h-12 w-24 px-3 flex items-center justify-center
                      bg-transparent hover:bg-white/10 active:bg-white/20
                      transition-colors duration-200 rounded-md
                      border-none focus:outline-none
                      text-white hover:text-white active:text-white"
              >
                <span>LogOut</span>
              </button>
            </div>


          </div>
        </div>
      </div>
    </header>
     
    <div v-if="breadcrumb && breadcrumb.length > 0" class="container my-3 py-4 bg-[#fff] text-[#878787]">
     <NBreadcrumb>
      <NBreadcrumbItem>
        <NuxtLink to="/" class="text-[#878787]">Home</NuxtLink>
      </NBreadcrumbItem>
      <NBreadcrumbItem v-for="item in breadcrumb">
        <NuxtLink :to="item.link" class="text-[#878787]">{{ item.title }}</NuxtLink>
      </NBreadcrumbItem>
     </NBreadcrumb>
     </div>
    <slot></slot>

    <!-- 页脚 -->
    <footer class=" flex flex-col justify-center items-center h-48 bg-[#f6f6f8] text-[#878787] pd-4 mt-10 text-lg">
        <!-- <div class="footer-links ">
          <NuxtLink class="footer-link" href="#">About Us</NuxtLink>
          <NuxtLink class="footer-link" href="#">Terms of Service</NuxtLink>
          <NuxtLink class="footer-link" href="#">Privacy Policy</NuxtLink>
          <NuxtLink class="footer-link" href="#">Shopping Guide</NuxtLink>
          <NuxtLink class="footer-link" href="#">Shipping Info</NuxtLink>
          <NuxtLink class="footer-link" href="#">Returns & Exchanges</NuxtLink>
        </div> -->
        <div v-if="footer_list && footer_list.length > 0" class="footer-links grid grid-cols-2 gap-1 w-full max-w-7xl mx-auto">
        <NuxtLink 
        v-for="item in footer_list"
        :to="item.link"
        class="footer-link text-center px-4 py-1 text-sm text-[#878787] hover:text-[#f4b3c2] transition-colors">
          {{ item.title }}
        </NuxtLink>
      </div>
        <div class="payment-icons flex flex-wrap gap-2 mt-4">
          <img class="lazyload ls-is-cached" 
              src="~/assets/images/01.svg" 
              
              alt="Visa">
          <img class="lazyload ls-is-cached" 
              src="~/assets/images/02.svg"
              alt="Mastercard">
          <img class="lazyload ls-is-cached" 
              src="~/assets/images/03.svg"
                alt="American Express">
          <img class="lazyload ls-is-cached" 
              src="~/assets/images/04.svg"
                alt="PayPal">
          <img class="lazyload ls-is-cached" 
              src="~/assets/images/05.svg"
              alt="Apple Pay">
        </div>       
        <div class="footer-copyright text-center mt-4 text-sm text-[#878787]">
          <p>Copyright 2025 m.earring18.com All Rights Reserved</p>
        </div>
    </footer>
    <!-- /页脚 -->

    <LoginModal ref="loginModalRef" @login-success="handleLogin"/>
  </div>
</template>

<style lang="css" scoped>
.logo {
  display: flex;
  margin-top: 80px;
}
.logo .search-box {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-end;
  width: 500px;
  margin-left: auto;
  color: #878787;
}

.site-header--spacer {
  font-weight: 300;
  margin: 0 10px;
  color: #878787;
  font-size: 16px;
}

.form {
  display: flex;
  column-gap: 20px;
}
footer {
  margin-top: 30px;
  background: #f6f6f8;
  color: #fff;
  padding: 40px 0;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 30px;
  padding: 20px 0;
  border-top: 1px solid rgba(255,255,255,0.2);
  border-bottom: 1px solid rgba(255,255,255,0.2);
}

.footer-link {
  color: #878787;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.3s ease;
  position: relative;
  padding: 5px 0;
}

.footer-link:hover {
  color: #f4b3c2;
}

.footer-link::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 1px;
  background-color: #fff;
  transition: width 0.3s ease;
}

.footer-link:hover::after {
  width: 100%;
}
</style>