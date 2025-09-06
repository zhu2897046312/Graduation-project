import { v4 as uuidv4 } from 'uuid';

let _deviceId: string | null = null;


export const getDeviceId = () => {
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