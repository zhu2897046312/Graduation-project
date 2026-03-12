import { useCookie } from 'nuxt/app'

// 简单的 UUID v4 生成函数（避免依赖）
function uuidv4(): string {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
    const r = (Math.random() * 16) | 0
    const v = c === 'x' ? r : (r & 0x3) | 0x8
    return v.toString(16)
  })
}

let _deviceId: string | null = null
let _browserFingerprint: string | null = null

/**
 * 浏览器指纹数据接口
 */
export interface BrowserFingerprintData {
  userAgent: string
  language: string
  languages: string
  platform: string
  screenWidth: number
  screenHeight: number
  colorDepth: number
  timezone: string
  timezoneOffset: number
  hardwareConcurrency: number
  deviceMemory: number | null
  cookieEnabled: boolean
  doNotTrack: string | null
  canvas?: string // Canvas指纹（客户端生成）
  webgl?: string // WebGL指纹（客户端生成）
}

/**
 * 在客户端收集浏览器指纹信息
 */
export const collectBrowserFingerprint = async (): Promise<BrowserFingerprintData | null> => {
  if (import.meta.server) {
    return null
  }

  try {
    const fingerprint: BrowserFingerprintData = {
      userAgent: navigator.userAgent || '',
      language: navigator.language || '',
      languages: navigator.languages?.join(',') || '',
      platform: navigator.platform || '',
      screenWidth: screen.width || 0,
      screenHeight: screen.height || 0,
      colorDepth: screen.colorDepth || 0,
      timezone: Intl.DateTimeFormat().resolvedOptions().timeZone || '',
      timezoneOffset: new Date().getTimezoneOffset(),
      hardwareConcurrency: navigator.hardwareConcurrency || 0,
      deviceMemory: (navigator as any).deviceMemory || null,
      cookieEnabled: navigator.cookieEnabled,
      doNotTrack: navigator.doNotTrack || null
    }

    // 生成 Canvas 指纹
    try {
      const canvas = document.createElement('canvas')
      const ctx = canvas.getContext('2d')
      if (ctx) {
        ctx.textBaseline = 'top'
        ctx.font = '14px Arial'
        ctx.fillText('Browser fingerprint 🎯', 2, 2)
        fingerprint.canvas = canvas.toDataURL()
      }
    } catch (e) {
      // Canvas 指纹生成失败，忽略
    }

    // 生成 WebGL 指纹
    try {
      const canvas = document.createElement('canvas')
      const gl = canvas.getContext('webgl') || canvas.getContext('experimental-webgl') as WebGLRenderingContext | null
      if (gl) {
        const debugInfo = gl.getExtension('WEBGL_debug_renderer_info')
        if (debugInfo) {
          const vendor = gl.getParameter(debugInfo.UNMASKED_VENDOR_WEBGL)
          const renderer = gl.getParameter(debugInfo.UNMASKED_RENDERER_WEBGL)
          fingerprint.webgl = `${vendor}|${renderer}`
        }
      }
    } catch (e) {
      // WebGL 指纹生成失败，忽略
    }

    return fingerprint
  } catch (error) {
    console.error('Failed to collect browser fingerprint:', error)
    return null
  }
}

/**
 * 生成浏览器指纹哈希
 */
export const generateFingerprintHash = (fingerprint: BrowserFingerprintData | null): string => {
  if (!fingerprint) {
    return uuidv4()
  }

  // 将指纹数据转换为字符串
  const fingerprintString = JSON.stringify({
    ua: fingerprint.userAgent,
    lang: fingerprint.language,
    langs: fingerprint.languages,
    plat: fingerprint.platform,
    sw: fingerprint.screenWidth,
    sh: fingerprint.screenHeight,
    cd: fingerprint.colorDepth,
    tz: fingerprint.timezone,
    tzo: fingerprint.timezoneOffset,
    hc: fingerprint.hardwareConcurrency,
    dm: fingerprint.deviceMemory,
    ce: fingerprint.cookieEnabled,
    dnt: fingerprint.doNotTrack,
    canvas: fingerprint.canvas?.substring(0, 100), // 只取前100个字符
    webgl: fingerprint.webgl
  })

  // 在服务端和客户端都使用简单的哈希算法
  let hash = 0
  for (let i = 0; i < fingerprintString.length; i++) {
    const char = fingerprintString.charCodeAt(i)
    hash = ((hash << 5) - hash) + char
    hash = hash & hash
  }
  
  const hashStr = Math.abs(hash).toString(36)
  return import.meta.server ? hashStr.padStart(12, '0') : hashStr.padStart(8, '0')
}

/**
 * 获取或生成设备指纹ID（浏览器指纹 + UUID）
 * 用于游客购买时的身份识别
 * 
 * @returns Promise<string> 设备指纹ID
 */
export const getDeviceId = async (): Promise<string> => {
  const deviceId = useCookie('device_id', {
    maxAge: 60 * 60 * 24 * 365, // 1年有效期
    sameSite: 'lax',
    secure: true,
    httpOnly: false // 需要客户端访问
  })

  // 如果已有设备ID，直接返回
  if (deviceId.value) {
    return deviceId.value
  }

  // 服务端：生成临时ID，等待客户端生成完整指纹
  if (import.meta.server) {
    if (!_deviceId) {
      _deviceId = uuidv4()
    }
    deviceId.value = _deviceId
    return _deviceId
  }

  // 客户端：生成完整的浏览器指纹
  try {
    // 收集浏览器指纹
    const fingerprint = await collectBrowserFingerprint()
    
    // 生成指纹哈希
    const fingerprintHash = generateFingerprintHash(fingerprint)
    
    // 结合UUID生成最终设备ID（指纹哈希 + UUID，确保唯一性）
    const uniqueId = `${fingerprintHash}-${uuidv4()}`
    
    // 保存到Cookie
    if (deviceId.value !== uniqueId) {
      deviceId.value = uniqueId
    }
    _browserFingerprint = uniqueId
    
    return uniqueId
  } catch (error) {
    console.error('Failed to generate device ID:', error)
    // 降级方案：使用UUID
    const fallbackId = uuidv4()
    deviceId.value = fallbackId
    return fallbackId
  }
}

/**
 * 同步版本（用于需要立即返回的场景）
 * 如果还没有设备ID，会生成一个临时ID，异步会更新为完整指纹
 * 
 * @returns string 设备指纹ID
 */
export const getDeviceIdSync = (): string => {
  const deviceId = useCookie('device_id', {
    maxAge: 60 * 60 * 24 * 365,
    sameSite: 'lax',
    secure: true,
    httpOnly: false
  })

  if (deviceId.value) {
    return deviceId.value
  }

  // 服务端
  if (import.meta.server) {
    if (!_deviceId) {
      _deviceId = uuidv4()
    }
    deviceId.value = _deviceId
    return _deviceId
  }

  // 客户端：如果还没有，生成临时ID（异步会更新为完整指纹）
  if (!_browserFingerprint || !deviceId.value) {
    const tempId = uuidv4()
    if (!deviceId.value) {
      deviceId.value = tempId
    }
    _browserFingerprint = deviceId.value || tempId
    
    // 异步生成完整指纹并更新
    collectBrowserFingerprint().then(fingerprint => {
      const fingerprintHash = generateFingerprintHash(fingerprint)
      const uniqueId = `${fingerprintHash}-${uuidv4()}`
      if (deviceId.value !== uniqueId) {
        deviceId.value = uniqueId
      }
      _browserFingerprint = uniqueId
    }).catch(() => {
      // 失败时保持临时ID
    })
    
    return _browserFingerprint
  }

  return _browserFingerprint
}

// 处理商品图片
export const getProductImage = (thumb: string) => {
  if (!thumb || thumb === '') {
    return '/placeholder-product.jpg'
  }
  
  try {
    // 尝试解析 JSON 字符串数组（如果图片是 JSON 格式）
    const images = JSON.parse(thumb)
    if (Array.isArray(images) && images.length > 0) {
      return images[0]
    }
  } catch (error) {
    // 如果不是 JSON 格式，直接使用原字符串
  }
  
  // 如果是普通字符串，直接返回
  // 如果图片路径是相对路径，确保以 / 开头
  if (thumb.startsWith('/') || thumb.startsWith('http://') || thumb.startsWith('https://')) {
    return thumb
  }
  
  // 如果不是以 / 开头，添加 /
  return '/' + thumb
}
