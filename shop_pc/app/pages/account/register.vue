<script setup lang="ts">
import api from '../../../api';
import { useMessage ,NButton,NForm , NInput,NFormItem} from 'naive-ui';
import type { FormRules } from 'naive-ui'
const accessToken = useCookie('accessToken');
const message = useMessage();
const router = useRouter();
const formRef = ref<any>(null)
const form = ref({
  nickname: '',
  email: '',
  password: ''
})
const registerRules = ref<FormRules>({
  nickname: [{
      required: true,
      message: 'Name is required'
    }
  ],
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
const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})
const handleValidate = async () => {
  formRef.value.validate((errors: any) => {
    if (!errors) {
      handleSubmit()
    }
  })
}
const handleSubmit = async () => {
  try {
    const res: any = await api.shop.user.register(form.value);
      accessToken.value = res;
      message.success('register success');
      router.push('/');
  } catch (error: any) {
    message.error(error.message);
  }
};
const toggleForm = () => {
  form.value = {
    nickname: '',
    email: '',
    password: ''
  };
  router.push('/account/login');
};
</script>

<template>
  <div class="container mt-8">
    <template v-if="siteInfo != null">
      <Title>Sign Up{{ siteInfo ? ` - ${siteInfo.seo_title}` : '' }}</Title>
      <Meta name="keywords" :content="siteInfo ? siteInfo.seo_keyword : 'register, sign up, create account'" />
      <Meta name="description" :content="siteInfo ? siteInfo.seo_description : 'Create an account to start shopping with us'" />
    </template>
    
    <h1>Sign Up</h1>

    <NForm ref="formRef" size="large" style="width: 30%;" :rules=" registerRules" :model="form">
      <NFormItem label="Email" path="email">
        <NInput placeholder="Please enter your email" v-model:value="form.email" />
      </NFormItem>
      <NFormItem label="Name" path="nickname">
        <NInput placeholder="Please enter your name" v-model:value="form.nickname" />
      </NFormItem>
      <NFormItem label="Password" path="password">
        <NInput placeholder="Please enter your password" v-model:value="form.password" type="password" />
      </NFormItem>
      <div class="flex flex-col items-center gap-2 mt-3 w-full">
        <div class="w-full">
          <NButton type="info" @click.prevent="handleValidate" style="width: 100%;" size="large">
            Sign Up
          </NButton>
        </div>
        <div class="auth-switch">
          <span>Already have an account? </span>
          <a href="#" @click.prevent="toggleForm" class="switch-link">
            Sign In
          </a>
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
.switch-link {
  color: #FB7F86;
  text-decoration: none;
  font-weight: 500;
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

/* 抖动动画 */
.shake-animation {
  animation: shake 0.5s cubic-bezier(.36,.07,.19,.97) both;
  transform: translate3d(0, 0, 0);
}

@keyframes shake {
  10%, 90% {
    transform: translate3d(-1px, 0, 0);
  }
  20%, 80% {
    transform: translate3d(2px, 0, 0);
  }
  30%, 50%, 70% {
    transform: translate3d(-4px, 0, 0);
  }
  40%, 60% {
    transform: translate3d(4px, 0, 0);
  }
}

/* 链接样式 */
a {
  color: #f4b3c2;
  text-decoration: none;
}
a:hover {
  text-decoration: underline;
}
</style>
