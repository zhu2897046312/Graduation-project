import { useCookie, useFetch } from "nuxt/app";
import { getDeviceId } from './auth'
// import { Ref } from "vue";

interface ResultData<T = unknown> {
  code: number;
  result: T;
  message: string;
  error?: unknown;
  time?: number;
}


class HttpRequest {
  async exec<T = unknown> (method: 'GET' | 'POST', url: string, data: unknown) : Promise<T> {
    const config = useRuntimeConfig()
    const accessToken = useCookie('accessToken')
    const deviceId = await getDeviceId(); // 获取设备指纹ID（支持游客购买）
    
    // 服务端 SSR 用 runtimeConfig.apiUrl（Docker 中为 http://backend:8080），浏览器用 public.apiUrl（localhost:8080）
    const baseURL = (config.apiUrl || config.public.apiUrl) || 'http://localhost:8080/api/client'
    if (import.meta.server) {
      console.log('[zhuyi-store] API baseURL:', baseURL)
    }
    let fullUrl = `${baseURL}${url}`
    
    // 构建请求选项
    const fetchOptions: {
      method: 'GET' | 'POST';
      headers: Record<string, string>;
      body?: string;
    } = {
      method: method,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': accessToken.value || '',
        'X-Device-Fingerprint': deviceId // 添加浏览器指纹到请求头，用于游客身份识别
      }
    }
    
    // 根据请求方法设置参数或请求体
    if (method === 'POST') {
      fetchOptions.body = JSON.stringify(data)
    } else if (data && typeof data === 'object' && data !== null) {
      // GET 请求使用查询参数
      const params = new URLSearchParams()
      Object.entries(data).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          params.append(key, String(value))
        }
      })
      const queryString = params.toString()
      if (queryString) {
        fullUrl = `${fullUrl}?${queryString}`
      }
    }

    try {
      console.log(`${method} 请求:`, fullUrl, method === 'POST' ? fetchOptions.body : '')
      const response = await $fetch<ResultData<T>>(fullUrl, fetchOptions)
      
      if (response == null) {
        return Promise.reject(new Error('请求失败'))
      }
      
      if (response.code === 0) {
        return Promise.resolve(response.result)
      }
      return Promise.reject(new Error(response.message || '请求失败'))
    } catch (e: unknown) {
      console.error('API 请求错误:', e)
      const errorMessage = e instanceof Error ? e.message : '请求失败'
      return Promise.reject(new Error(errorMessage))
    }
  }
}

export const httpRequest = new HttpRequest()

