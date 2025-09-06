<script setup lang="ts">
import api from '~/api';
import { NButton, useMessage, NForm , NInput,NFormItem } from 'naive-ui';
import type { FormRules } from 'naive-ui'
const accessToken = useCookie('accessToken')
const message = useMessage();
const router = useRouter();
const loginRules = ref<FormRules>({
  email: [{
      required: true,
      message: 'Email is required'
    },
    {
      pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
      message: 'Please enter a valid email address'
    }
  ],
  password: [{
    required: true,
    message: 'Password is required'
  }]
})
const form = ref({
  email: '',
  password: ''
})
const formRef = ref<any>(null)
  const handleValidate = async () => {
  formRef.value.validate((errors: any) => {
    if (!errors) {
      handleSubmit()
    }
  })
}
const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})

const handleSubmit = async () => {
  try {
    const res: any = await api.shop.user.login({
        email: form.value.email,
        password: form.value.password
      });
      accessToken.value = res;
      message.success('login success');
      router.push('/account');
  } catch (error: any) {
    message.error(error.message);
  }
};
const toggleForm = () => {
  // 重置表单
  form.value = {
    email: '',
    password: ''
  };
  router.push('/account/register');
};
</script>

<template>
  <div class="container mt-8" style="color: #545454;">
    <template v-if="siteInfo != null">
      <Title>Sign In - {{ siteInfo.seo_title ? siteInfo.seo_title : '' }}</Title>
      <Meta name="keywords" content="Sign In" />
      <Meta name="description" content="Sign In" />
    </template>

    <h1>Sign In</h1>

    <NForm ref="formRef" size="large" style="width: 30%;" :rules="loginRules" :model="form">
      <NFormItem label="Email" path="email">
        <NInput placeholder="Please enter your email" v-model:value="form.email" />
      </NFormItem>
      <NFormItem label="Password" path="password">
        <NInput placeholder="Please enter your password" v-model:value="form.password" type="password" />
      </NFormItem>
      <div class="w-full flex items-end justify-end">
        <NuxtLink to="/account/reset-pwd">Forget Password ?</NuxtLink>
      </div>
      <div class="flex flex-col items-center gap-2 mt-3 w-full">
        <div class="w-full">
          <NButton type="info" @click.prevent="handleValidate" style="width: 100%;" size="large">
            Sign In
          </NButton>
        </div>
        <div class="auth-switch">
          <span>Don't have an account? </span>
          <a href="#" @click.prevent="toggleForm" class="switch-link ">
            Sign Up
          </a>
        </div>
        
        <div class="flex justify-center gap-4">
          <NuxtLink to="/">Back to Home</NuxtLink>
        </div>
      </div>
    </NForm>
  </div>
</template>

<style lang="css" scoped>
h1 {
  font-size: 20px;
  text-align: center;
  font-weight: bold;
  padding: 8px;
}
form {
  width: 400px;
  margin: 10px auto;
}
.from-item {
  display: flex;
  flex-direction: column;
  gap: 5px;
  padding-bottom: 10px;
  margin-top: 12px;
}
.from-item label {
  font-size: 14px;
  color: #333;
  opacity: 0.8;
}
.from-item-input {
  padding: 8px;
  border: 1px solid #c2c2c2;
  border-radius: 6px;
  color: #333;
  transition: all 0.3s ease;
}
.from-item-input.error-border {
  border-color: #ff4757;
  background-color: #fff9f9;
}
.from-item-input:focus.error-border {
  box-shadow: 0 0 0 2px rgba(255, 71, 87, 0.2);
}
a {
  color: #f4b3c2;
}

</style>
