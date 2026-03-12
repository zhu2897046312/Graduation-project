<script setup lang="ts">
import api from '../api'
import { useCookie } from 'nuxt/app'

// 禁用默认布局，使用独立页面
definePageMeta({
  layout: false
})

const router = useRouter()
const toast = useToast()

// 表单数据
const form = reactive({
  email: '',
  password: ''
})

// 表单验证规则
const emailRules = [
  (v: string) => !!v || '请输入邮箱地址',
  (v: string) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v) || '请输入有效的邮箱地址'
]

const passwordRules = [
  (v: string) => !!v || '请输入密码',
  (v: string) => v.length >= 6 || '密码长度至少为6位'
]

// 表单状态
const loading = ref(false)
const showPassword = ref(false)
const errors = reactive({
  email: '',
  password: ''
})

// 验证单个字段
const validateField = (field: 'email' | 'password') => {
  const rules = field === 'email' ? emailRules : passwordRules
  const value = form[field]
  
  for (const rule of rules) {
    const error = rule(value)
    if (error !== true) {
      errors[field] = error
      return false
    }
  }
  
  errors[field] = ''
  return true
}

// 处理登录
const handleLogin = async () => {
  // 验证所有字段
  const isEmailValid = validateField('email')
  const isPasswordValid = validateField('password')
  
  if (!isEmailValid || !isPasswordValid) {
    return
  }
  
  loading.value = true
  
  try {
    const result = await api.shop.user.login({
      email: form.email,
      password: form.password
    })
    
    // 登录成功，保存 token（根据实际 API 响应调整）
    // 支持多种可能的响应格式：result.token, result.accessToken, 或 result 本身就是 token
    const token = result?.token || result?.accessToken || (typeof result === 'string' ? result : null)
    
    if (token) {
      const accessToken = useCookie('accessToken', {
        maxAge: 60 * 60 * 24 * 7, // 7天
        sameSite: 'lax',
        secure: true
      })
      accessToken.value = token
      
      toast.add({
        title: '登录成功',
        description: '欢迎回来！',
        color: 'success'
      })
      
      // 跳转到首页或之前访问的页面
      const redirect = useRoute().query.redirect as string || '/'
      await router.push(redirect)
    } else {
      throw new Error('登录响应格式异常，未找到 token')
    }
  } catch (error: any) {
    toast.add({
      title: '登录失败',
      description: error.message || '邮箱或密码错误，请重试',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}

// 回车键提交
const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Enter' && !loading.value) {
    handleLogin()
  }
}

// SEO
useHead({
  title: '登录 - 用户登录',
  meta: [
    { name: 'description', content: '登录您的账户以继续购物' }
  ]
})
</script>

<template>
  <div class="h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900 px-4 sm:px-6 lg:px-8 overflow-hidden">
    <div class="max-w-md w-full space-y-4">
      <!-- Logo 和标题区域 -->
      <div class="text-center">
        <NuxtLink to="/" class="inline-block mb-4">
          <AppLogo class="w-auto h-8 mx-auto" />
        </NuxtLink>
        <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
          欢迎回来
        </h2>
        <p class="mt-1 text-sm text-gray-600 dark:text-gray-400">
          登录您的账户以继续购物
        </p>
      </div>

      <!-- 登录表单卡片 -->
      <UCard class="shadow-xl border-0">
        <template #header>
          <div class="text-center py-1">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
              账户登录
            </h3>
          </div>
        </template>

        <form @submit.prevent="handleLogin" class="space-y-4">
          <!-- 邮箱输入 -->
          <div>
            <UFormGroup
              label="邮箱地址"
              :error="errors.email"
              required
            >
              <UInput
                v-model="form.email"
                type="email"
                placeholder="your.email@example.com"
                size="lg"
                :disabled="loading"
                @blur="validateField('email')"
                @keydown="handleKeydown"
                icon="i-lucide-mail"
                autocomplete="email"
                class="w-full"
              />
            </UFormGroup>
          </div>

          <!-- 密码输入 -->
          <div>
            <UFormGroup
              label="密码"
              :error="errors.password"
              required
            >
              <UInput
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                placeholder="请输入密码"
                size="lg"
                :disabled="loading"
                @blur="validateField('password')"
                @keydown="handleKeydown"
                icon="i-lucide-lock"
                autocomplete="current-password"
                class="w-full"
              >
                <template #trailing>
                  <UButton
                    :icon="showPassword ? 'i-lucide-eye-off' : 'i-lucide-eye'"
                    color="neutral"
                    variant="link"
                    :padded="false"
                    @click="showPassword = !showPassword"
                  />
                </template>
              </UInput>
            </UFormGroup>
          </div>

          <!-- 忘记密码链接 -->
          <div class="flex items-center justify-end">
            <NuxtLink
              to="/forgot-password"
              class="text-sm font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400 dark:hover:text-primary-300 transition-colors"
            >
              忘记密码？
            </NuxtLink>
          </div>

          <!-- 登录按钮 -->
          <UButton
            type="submit"
            block
            size="lg"
            :loading="loading"
            :disabled="loading"
            class="w-full font-semibold"
          >
            {{ loading ? '登录中...' : '登录' }}
          </UButton>
        </form>

        <!-- 分隔线 -->
        <div class="mt-4">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300 dark:border-gray-700" />
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-2 bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-400">
                还没有账户？
              </span>
            </div>
          </div>
        </div>

        <!-- 注册链接 -->
        <div class="mt-4 text-center">
          <NuxtLink
            to="/register"
            class="inline-flex items-center text-sm font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400 dark:hover:text-primary-300 transition-colors"
          >
            立即注册
            <UIcon name="i-lucide-arrow-right" class="ml-1 w-4 h-4" />
          </NuxtLink>
        </div>
      </UCard>

      <!-- 底部提示 -->
      <p class="text-center text-xs text-gray-500 dark:text-gray-400 mt-2">
        登录即表示您同意我们的
        <NuxtLink to="/terms" class="underline hover:text-gray-700 dark:hover:text-gray-300 transition-colors">
          服务条款
        </NuxtLink>
        和
        <NuxtLink to="/privacy" class="underline hover:text-gray-700 dark:hover:text-gray-300 transition-colors">
          隐私政策
        </NuxtLink>
      </p>
    </div>
  </div>
</template>

<style scoped>
/* 确保表单在深色模式下正常显示 */
:deep(.dark) {
  color-scheme: dark;
}

/* 优化卡片样式 */
:deep(.card) {
  background: white;
}

:deep(.dark .card) {
  background: rgb(17 24 39);
}

/* 确保输入框宽度一致 */
:deep(.form-group) {
  width: 100%;
}

:deep(.form-group input) {
  width: 100%;
}
</style>
