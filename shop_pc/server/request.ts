import { useCookie, useFetch } from "nuxt/app";
import { getDeviceId } from '../utils/auth'
// import { Ref } from "vue";

interface ResultData {
  code: number;
  result: any;
  message: string;
}


class HttpRequest {
  async exec (method: 'GET' | 'POST', url: string, data: any) : Promise<any> {
    const config = useRuntimeConfig()
    const accessToken = useCookie('accessToken')
    const deviceId = getDeviceId(); // 获取设备ID
    // console.log(config)
    // console.log(deviceId)
    const options: any = {
      baseURL: config.public.apiUrl,
      method: method,
      headers: {
        'Authorization': accessToken.value || '',
        'X-Device-Fingerprint': deviceId // 添加指纹到请求头
      }
    };
    if (method === 'POST') {
      options.body = data
    } else {
      options.params = data
    }

    try {
      
      // const {data} = await useFetch<ResultData>(url, options)
      const data = await $fetch<ResultData>(url, options)
      if (data == null) {
        return Promise.reject(new Error('请求失败'))
      }
      // console.log(options, url, data)
      if (data.code === 0) {
        return Promise.resolve(data.result)
      }
      return Promise.reject(new Error(data.message))
    } catch (e: any) {
      return Promise.reject(e)
    }
  }
}

export const httpRequest = new HttpRequest()

