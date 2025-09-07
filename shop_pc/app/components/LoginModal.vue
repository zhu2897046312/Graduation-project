<script setup lang="ts">
import { NModal, NCard, NForm, NFormItem, NInput, NButton, useMessage } from 'naive-ui';
import api from '../../api';
const show = ref(false)
const showRegister = ref(false)

const message = useMessage()

const info = ref({
  nickname: '',
  email: '',
  password: ''
})

const emits = defineEmits(['loginSuccess'])

// const accessToken = useCookie('accessToken')

const handleSubmit = async () => {
  try {
    const res: any = await api.shop.user.login(info.value)
    // accessToken.value = res
    nextTick(() => {
      emits('loginSuccess', res)
    })
    show.value = false
  } catch (error: any) {
    console.warn(error)
    message.error(error.message)
  }
  return false;
}

const handleSubmitRegister = async () => {
  try {
    const res: any = await api.shop.user.register(info.value)
    // accessToken.value = res
    nextTick(() => {
      emits('loginSuccess', res)
    })
    showRegister.value = false
  } catch (error: any) {
    console.warn(error)
    message.error(error.message)
  }
  return false;
}

const handleBackLogin = () => {
  showRegister.value = false
  show.value = true
}

const handleShowRegister = () => {
  show.value = false
  showRegister.value = true
}

defineExpose({
  useOpen: async () => {
    show.value = true
  }
})

</script>

<template>
  <NModal v-model:show="show" style="width: 500px;">
    <NCard title="Sign In">
      <NForm size="large">
        <NFormItem label="Email">
          <NInput placeholder="Please enter your email" v-model:value="info.email" />
        </NFormItem>
        <NFormItem label="Password">
          <NInput placeholder="Please enter your passwor" v-model:value="info.password" type="password"/>
        </NFormItem>
      </NForm>
      <template #footer>
        <NButton block size="large" type="info" @click="handleSubmit">Sign In</NButton>
        <NButton block text size="large" @click="handleShowRegister" style="margin-top: 20px;color: var(--font-primary-color);">Sign Up</NButton>
      </template>
    </NCard>
  </NModal>

  <NModal v-model:show="showRegister" style="width: 500px;">
    <NCard title="Sign Up">
      <NForm size="large">
        <NFormItem label="Name">
          <NInput placeholder="Please enter your name" v-model:value="info.nickname" />
        </NFormItem>
        <NFormItem label="Email">
          <NInput placeholder="Please enter your email" v-model:value="info.email" />
        </NFormItem>
        <NFormItem label="Password">
          <NInput placeholder="Please enter your password" v-model:value="info.password" type="password"/>
        </NFormItem>
      </NForm>
      <template #footer>
        <NButton block size="large" type="info" @click="handleSubmitRegister">Finish</NButton>
        <NButton block text size="large" @click="handleBackLogin" style="margin-top: 20px;color: var(--font-primary-color);">Back to Sign In</NButton>
      </template>
    </NCard>
  </NModal>
</template>