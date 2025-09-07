<script lang="ts" setup>
import { NInput, NButton, useMessage, useDialog } from 'naive-ui'
import api from '../../../api'

const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})

const message = useMessage()
const dialog = useDialog()
const router = useRouter()
const email = ref('')
const loading = ref(false)

const handleSubmit = async () => {
  // Email validation
  if (!email.value) {
    message.error('Please enter your email address')
    return
  }
  
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    message.error('Please enter a valid email address')
    return
  }

  try {
    loading.value = true
    await api.shop.user.SendEmail({
      email: email.value
    })
    
    dialog.success({
      title: 'Success',
      content: 'Password reset email sent successfully! Please check your inbox.',
      positiveText: 'OK',
      onPositiveClick: () => {
        email.value = ''
        router.push('/account/login')
      }
    })
  } catch (error: any) {
    message.error(error.message || 'Failed to send reset email')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="container mx-auto mt-12 flex flex-col items-center gap-6 max-w-md px-4">
    <template v-if="siteInfo != null">
      <Title>Reset Password - {{ siteInfo.seo_title ? siteInfo.seo_title : '' }}</Title>
      <Meta name="keywords" content="Reset Password" />
      <Meta name="description" content="Reset Password" />
    </template>
    
    <div class="text-center">
      <h1 class="text-2xl font-bold text-gray-800 mb-2">Reset Your Password</h1>
      <p class="text-gray-600">Enter your email to receive a password reset link</p>
    </div>

    <div class="w-full space-y-4">
      <NInput
        v-model:value="email"
        type="text"
        placeholder="<EMAIL>"
        size="large"
        clearable
        :input-props="{ autocapitalize: 'off', autocorrect: 'off' }"
        @keydown.enter="handleSubmit"
      />

      <NButton
        type="primary"
        size="large"
        block
        :loading="loading"
        @click="handleSubmit"
        class="mt-2"
      >
        Send Reset Link
      </NButton>
    </div>

    <div class="text-center text-sm text-gray-500">
      Remember your password? 
      <NuxtLink 
        to="/account/login" 
        class="text-blue-600 hover:underline font-medium"
      >
        Sign in here
      </NuxtLink>
    </div>
  </div>
</template>

<style scoped>
.container {
  min-height: calc(100vh - 6rem);
}
</style>
