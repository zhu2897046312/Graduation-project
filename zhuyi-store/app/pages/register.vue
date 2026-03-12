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
  username: '',
  password: ''
})

// 表单验证规则
const emailRules = [
  (v: string) => !!v || '请输入邮箱地址',
  (v: string) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v) || '请输入有效的邮箱地址'
]

const usernameRules = [
  (v: string) => !!v || '请输入用户名',
  (v: string) => v.length >= 3 || '用户名长度至少为3位',
  (v: string) => /^[a-zA-Z0-9_]+$/.test(v) || '用户名只能包含字母、数字和下划线'
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
  username: '',
  password: ''
})

// 验证单个字段
const validateField = (field: 'email' | 'username' | 'password') => {
  const rules = field === 'email' ? emailRules : field === 'username' ? usernameRules : passwordRules
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

// 处理注册
const handleRegister = async () => {
  // 验证所有字段
  const isEmailValid = validateField('email')
  const isUsernameValid = validateField('username')
  const isPasswordValid = validateField('password')
  
  if (!isEmailValid || !isUsernameValid || !isPasswordValid) {
    return
  }
  
  loading.value = true
  
  try {
    await api.shop.user.register({
      email: form.email,
      username: form.username,
      password: form.password
    })
    
    toast.add({
      title: '注册成功',
      description: '账户创建成功，请登录',
      color: 'success'
    })
    
    // 注册成功后跳转到登录页面
    await router.push('/login')
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : '注册失败，请重试'
    toast.add({
      title: '注册失败',
      description: errorMessage,
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}

// 回车键提交
const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Enter' && !loading.value) {
    handleRegister()
  }
}

// SEO
useHead({
  title: '注册 - 用户注册',
  meta: [
    { name: 'description', content: '注册新账户以开始购物' }
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
          创建账户
        </h2>
        <p class="mt-1 text-sm text-gray-600 dark:text-gray-400">
          注册新账户以开始购物
        </p>
      </div>

      <!-- 注册表单卡片 -->
      <UCard class="shadow-xl border-0">
        <template #header>
          <div class="text-center py-1">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
              账户注册
            </h3>
          </div>
        </template>

        <form @submit.prevent="handleRegister" class="space-y-4">
          <!-- 邮箱输入 -->
          <UFormField
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
          </UFormField>

          <!-- 用户名输入 -->
          <UFormField
            label="用户名"
            :error="errors.username"
            required
          >
            <UInput
              v-model="form.username"
              type="text"
              placeholder="请输入用户名"
              size="lg"
              :disabled="loading"
              @blur="validateField('username')"
              @keydown="handleKeydown"
              icon="i-lucide-user"
              autocomplete="username"
              class="w-full"
            />
          </UFormField>

          <!-- 密码输入 -->
          <UFormField
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
              autocomplete="new-password"
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
          </UFormField>

          <!-- 注册按钮 -->
          <UButton
            type="submit"
            block
            size="lg"
            :loading="loading"
            :disabled="loading"
            class="w-full font-semibold mt-2"
          >
            {{ loading ? '注册中...' : '注册' }}
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
                已有账户？
              </span>
            </div>
          </div>
        </div>

        <!-- 登录链接 -->
        <div class="mt-4 text-center">
          <NuxtLink
            to="/login"
            class="inline-flex items-center text-sm font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400 dark:hover:text-primary-300 transition-colors"
          >
            立即登录
            <UIcon name="i-lucide-arrow-right" class="ml-1 w-4 h-4" />
          </NuxtLink>
        </div>
      </UCard>

      <!-- 底部提示 -->
      <p class="text-center text-xs text-gray-500 dark:text-gray-400 mt-2">
        注册即表示您同意我们的
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
