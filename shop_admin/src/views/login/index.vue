<template>
  <div class="min-h-screen w-full flex items-center justify-center bg-gradient-to-br from-blue-50 to-white relative overflow-hidden">
  

    <!-- 登录卡片 -->
    <div class="relative w-[420px] bg-white/80 backdrop-blur-sm p-8 rounded-2xl shadow-lg">
      <div class="text-center mb-8">
        <h1 class="text-2xl font-bold text-gray-800 mb-2">虾皮商城管理</h1>
        <p class="text-gray-500 text-sm"></p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div class="space-y-2">
          <Input
            v-model:value="loginForm.account"
            placeholder="请输入账号"
            :class="['!rounded-button']"
            class="w-full"
          >
            <template #prefix>
              <UserOutlined class="text-gray-400" />
            </template>
          </Input>
        </div>

        <div class="space-y-2">
          <Input.Password
            v-model:value="loginForm.pwd"
            placeholder="请输入密码"
            :class="['!rounded-button']"
            class="w-full"
          >
            <template #prefix>
              <LockOutlined class="text-gray-400" />
            </template>
          </Input.Password>
        </div>

        <Button
          type="primary" 
          html-type="submit"
          :loading="loading"
          class="w-full h-10 !rounded-button whitespace-nowrap bg-gradient-to-r from-blue-500 to-blue-600 border-none hover:from-blue-600 hover:to-blue-700"
        >
          登 录
        </Button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
interface LoginFromModel {
  account: string;
  pwd: string;
}
</script>

<script lang="ts" setup>
import { ref } from 'vue';
import { message, Input, Button } from 'ant-design-vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';

import { apiAuthLogin, apiGetAuthInfo, apiGetEnumDict } from '/@/api/auth';
import { useAuthStore } from '/@/store/authStore';
import { useRouter } from 'vue-router';
import enumDict from '/@/utils/enum-dict';

const authStore = useAuthStore();
const router = useRouter();

const loginForm = ref<LoginFromModel>({
  account: '',
  pwd: '',
});
const loading = ref(false);

const handleLogin = async () => {
  const values = {...loginForm.value}
  const hide = message.loading('登陆中...', 0);
  loading.value = true;
  try {
    const token: any = await apiAuthLogin(values);
    authStore.setToken(token);
    const info = await apiGetAuthInfo();
    authStore.setUser(info);

    const dicts: any = await apiGetEnumDict()
    enumDict.initDict(dicts);
    router.push({ name: 'Base' })
  } catch (e: any) {
    console.error(e);
  } finally {
    hide();
    loading.value = false;
  }
}
</script>

<style scoped>
:deep(.ant-input-affix-wrapper) {
  padding: 8px 11px;
}

:deep(.ant-input-affix-wrapper .anticon) {
  font-size: 16px;
}

:deep(.ant-input) {
  font-size: 14px;
}

:deep(.ant-btn-primary) {
  font-size: 16px;
  font-weight: 500;
}

</style>