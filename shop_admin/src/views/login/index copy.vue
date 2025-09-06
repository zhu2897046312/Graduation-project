<template>
  <div class="login_page">
    <div class="login_box">
      <div class="login_form">
        <div class="login_form_title">
          <img src="/bg/small_logo.png"/>
          <p>任务平台</p>
        </div>
        <Form ref="loginRef" :model="loginForm" name="loginForm" @finish="handleLogin">
          <Form.Item name="account" :rules="[{ required: true, message: '请输入登陆账户', trigger: 'blur' }]">
            <Input v-model:value="loginForm.account" size="large" placeholder="请输入账户">
              <template #prefix>
                <UserOutlined />
              </template>
            </Input>
          </Form.Item>
          <Form.Item name="pwd" :rules="[{ required: true, message: '请输入登陆密码', trigger: 'blur' }]">
            <Input.Password autocomplete="off" v-model:value="loginForm.pwd" size="large" placeholder="请输入密码">
              <template #prefix>
                <LockOutlined />
              </template>
            </Input.Password>
          </Form.Item>
          <Form.Item style="text-align: center">
            <Button size="large" type="primary" style="width: 168px;"  html-type="submit">登录</Button>
          </Form.Item>
        </Form>
      </div>
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
  import { Form, Input, Button, message } from 'ant-design-vue';
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

  const loginRef = ref<any>(null);

  const handleLogin = async (values: LoginFromModel) => {
    console.log(values);
    const hide = message.loading('登陆中...', 0);
    try {
      const token: any = await apiAuthLogin(values);
      authStore.setToken(token);
      const info = await apiGetAuthInfo();
      authStore.setUser(info);
      
      const dicts: any = await apiGetEnumDict()
      enumDict.initDict(dicts);
      router.push({ name: 'Base' })
    } catch (e: any) {
      console.warn(e);
    } finally {
      hide();
    }
  }
</script>

<style scoped>
  .login_page {
    position: relative;
    overflow: hidden;
    width: 100vw;
    height: 100vh;
    background-color: rgb(204, 204, 204);
  }
  .login_box {
    position: absolute;
    top: 3vh;
    left: 2vw;
    width: 96vw;
    height: 94vh;
    background-color: rgba(255,255,255,.8);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    background-image: url('/bg/login_bg.png');
  }
  .login_form {
    width: 420px;
    background-color: #fff;
    padding: 40px;
    border-radius: 10px;
    box-shadow: 2px 3px 7px rgba(0,0,0,.2);
  }
  .login_form_title {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-items: center;
    justify-content: center;
    padding-bottom: 20px;
    color: var(--primary-color);
  }
  .login_form_title img {
    display: block;
    width: 80px;
    height: 80px;
    border-radius: 80px;
  }
  .login_form_title p {
    font-size: 40px;
    padding-left: 20px;
  }
</style>