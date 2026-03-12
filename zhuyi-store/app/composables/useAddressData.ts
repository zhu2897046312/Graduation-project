/**
 * 共享的地址数据 composable
 * 用于加载和管理地址 JSON 数据（国家、省份等）
 */

// 地址数据类型
export interface AddressData {
  countries: Array<{
    name: string
    full_name: string
    provinces: string[]
  }>
}

/**
 * 使用地址数据
 * @param key - useAsyncData 的 key，用于区分不同的调用场景
 */
export const useAddressData = (key: string = 'address-json') => {
  // 加载地址数据 - 使用 $fetch 加载 JSON 文件
  const { data: addressJson } = useAsyncData<AddressData | null>(key, async () => {
    try {
      return await $fetch<AddressData>('/address.json')
    } catch (error) {
      console.error('Failed to load address.json:', error)
      return null
    }
  }, {
    lazy: true,
    default: () => null
  })

  /**
   * 获取国家选项
   * @param selectedCountry - 可选，用于过滤特定国家
   */
  const countryOptions = computed(() => {
    if (!addressJson.value?.countries) {
      return []
    }
    
    return addressJson.value.countries.map((c) => ({
      label: c.full_name,
      value: c.name
    }))
  })

  /**
   * 获取省份选项
   * @param countryName - 国家名称
   */
  const getProvinceOptions = (countryName: string | undefined) => {
    if (!countryName || !addressJson.value?.countries) {
      return []
    }
    
    const selectedCountry = addressJson.value.countries.find(
      (c) => c.name === countryName
    )
    
    return selectedCountry?.provinces?.map((p: string) => ({
      label: p,
      value: p
    })) || []
  }

  return {
    addressJson,
    countryOptions,
    getProvinceOptions
  }
}
