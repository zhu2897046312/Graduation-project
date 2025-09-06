<script lang="ts" setup>
import { NInput, NButton, useMessage } from 'naive-ui'
import api from '~/api'
definePageMeta({
  layout: 'blank'
})
const route = useRoute()
const message = useMessage()
const token = route.params.code as string
const password = ref('')
const confirm_password = ref('')
const loading = ref(false)

const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})

const handleSubmit = async () => {
  if (!password.value) {
    message.error('Please enter a new password')
    return
  }
  
  if (password.value.length < 6) {
    message.error('Password must be at least 6 characters')
    return
  }
  
  if (password.value !== confirm_password.value) {
    message.error('Passwords do not match')
    return
  }

  try {
    loading.value = true
    await api.shop.user.ResetPwd({
      code: token,
      password: password.value,
      confirm_password: confirm_password.value
    })
    
    message.success('Password reset successfully!')
    password.value = ''
    confirm_password.value = ''
  } catch (error: any) {
    message.error(error.message || 'Password reset failed')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center p-4">
    <div class="w-full max-w-md bg-white rounded-xl shadow-sm p-8 space-y-6">
      <template v-if="siteInfo != null">
        <Title>Reset Password </Title>
        <Meta name="keywords" content="Reset Password" />
        <Meta name="description" content="Reset Password" />
      </template>
      
      <div class="text-center space-y-2">
        <h1 class="text-2xl font-bold text-gray-800">Reset Password</h1>
        <p class="text-gray-500">Please enter your new password below</p>
      </div>
      
      <div class="space-y-4">
        <NInput 
          v-model:value="password"
          type="password"
          placeholder="New password (min 8 characters)"
          clearable
          class="w-full"
          @keydown.enter="handleSubmit"
          :input-props="{ class: 'py-2' }"
        />
        
        <NInput 
          v-model:value="confirm_password"
          type="password"
          placeholder="Confirm new password"
          clearable
          class="w-full"
          @keydown.enter="handleSubmit"
          :input-props="{ class: 'py-2' }"
        />
        
        <NButton 
          type="primary"
          size="medium" 
          @click="handleSubmit"
          :loading="loading"
          class="w-full bg-[#f4b3c2] hover:bg-[#e8a0b0] text-white font-medium py-2"
        >
          {{ loading ? 'Processing...' : 'Reset Password' }}
        </NButton>
      </div>
      
      <div class="text-center text-sm text-gray-500">
        Remember your password? 
        <NuxtLink 
          to="/account/login" 
          class="text-[#f4b3c2] hover:text-[#e8a0b0] font-medium transition-colors"
        >
          Sign in
        </NuxtLink>
      </div>
    </div>
  </div>
</template>