export const useSiteInfo = () => {
    return useAsyncData('siteInfo', () => {
      // 统一的数据获取逻辑
      return $fetch('/api/site-info')
    })
  }