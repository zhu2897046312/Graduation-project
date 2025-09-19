import { v4 as uuidv4 } from 'uuid';
import { useCookie } from 'nuxt/app';

let _deviceId: string | null = null;


export  const getDeviceId = () => {
  const deviceId = useCookie('device_id', {
    maxAge: 60 * 60 * 24 * 365, // 1年有效期
    sameSite: 'lax',
  });
  
  if (!deviceId.value ) {
    if(import.meta.server){
      if(!_deviceId){
        deviceId.value = uuidv4(); // 生成新的UUID
        _deviceId = deviceId.value
      }
    }
    if(import.meta.client){
      deviceId.value = uuidv4(); // 生成新的UUID
    }
  }
  return deviceId.value;
}

// export const getDeviceId = () => {
//   const deviceId = useCookie('X-Device-Fingerprint', {
//     maxAge: 60 * 60 * 24 * 365, // 1年有效期
//     sameSite: 'lax',
//   });
  
//   if (!deviceId.value) {
//     deviceId.value = uuidv4(); // 生成新的UUID
//   }
//   return deviceId.value;
// }

// 处理商品图片
export const getProductImage = (thumb: string) => {
  try {
    // 解析 JSON 字符串数组
    const images = JSON.parse(thumb)
    // 返回第一个图片 URL
    return images && images.length > 0 ? images[0] : '/placeholder-product.jpg'
  } catch (error) {
    console.error('Failed to parse thumb:', error)
    // 如果解析失败，尝试直接使用（可能是普通字符串）
    return thumb && thumb !== '' ? thumb : '/placeholder-product.jpg'
  }
}