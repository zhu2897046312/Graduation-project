import axios from 'axios';
import type { InternalAxiosRequestConfig, AxiosResponse } from 'axios';
import { useAuthStore } from '/@/store/authStore';
import { message } from 'ant-design-vue';

interface YexResponse {
  code: number;
  message: string;
  result?: any;
  time: number;
  error?: {
    message: string;
  };
}

const service = axios.create({
  baseURL: `http://localhost:8080/api/manage`,
  timeout: 30000,
});


service.interceptors.request.use((config: InternalAxiosRequestConfig<any>) => {
  const authStore = useAuthStore();
  if (authStore.currentToken) {
    config.headers.Authorization = authStore.currentToken;
  }
  return config;
}, (err: any) => {
  console.warn(err);
  return Promise.reject(err);
});

service.interceptors.response.use((res: AxiosResponse<YexResponse>) => {
  if (res.status != 200) {
    message.warn('网络异常');
    return Promise.reject('网络异常');
  }
  const { code, result } = res.data;
  if (code === 18000) {
    const authStore = useAuthStore();
    authStore.loginOut();
    window.location.href = '/#/login';
    return Promise.reject('登陆已超时，请重新登录');
  }
  if (code != 0) {
    message.warn(res.data.message);
    return Promise.reject(res.data.message);
  }
  return result;
}, (err: any) => {
  console.warn(err);
  return Promise.reject(err);
});

export default service;
