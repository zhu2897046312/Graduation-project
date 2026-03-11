<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-card-header">
        <h1 class="login-title">虾皮商城管理</h1>
        <p class="login-subtitle"></p>
      </div>

      <form @submit.prevent="handleLogin" class="login-form">
        <div class="login-field">
          <Input
            v-model:value="loginForm.account"
            placeholder="请输入账号"
            class="login-input"
          >
            <template #prefix>
              <UserOutlined class="login-input-icon" />
            </template>
          </Input>
        </div>

        <div class="login-field">
          <Input.Password
            v-model:value="loginForm.pwd"
            placeholder="请输入密码"
            class="login-input"
          >
            <template #prefix>
              <LockOutlined class="login-input-icon" />
            </template>
          </Input.Password>
        </div>

        <Button
          type="primary"
          html-type="submit"
          :loading="loading"
          class="login-btn"
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
/* 登录页容器：铺满视口、居中 */
.login-page {
  min-height: 100vh;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(to bottom right, #eff6ff, #fff);
  position: relative;
  overflow: hidden;
}

/* 登录卡片 */
.login-card {
  width: 420px;
  max-width: 420px;
  min-width: 420px;
  box-sizing: border-box;
  position: relative;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.08);
}

.login-card-header {
  text-align: center;
  margin-bottom: 2rem;
}

.login-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #1f2937;
  margin: 0 0 0.5rem 0;
}

.login-subtitle {
  font-size: 0.875rem;
  color: #6b7280;
  margin: 0;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.login-field {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.login-input {
  width: 100%;
}

.login-input-icon {
  color: #9ca3af;
}

.login-btn {
  width: 100%;
  height: 40px;
  white-space: nowrap;
  font-size: 16px;
  font-weight: 500;
  border: none;
  background: linear-gradient(to right, #3b82f6, #2563eb);
  border-radius: 6px;
}

.login-btn:hover {
  background: linear-gradient(to right, #2563eb, #1d4ed8);
}

:deep(.ant-input-affix-wrapper) {
  padding: 8px 11px;
  border-radius: 6px;
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