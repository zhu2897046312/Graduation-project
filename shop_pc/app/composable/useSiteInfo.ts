// composables/useSiteInfo.ts
import api from '../../api';
import  type { MarketInfo } from '../../api/type';

const DEFAULT_MARKET_INFO: MarketInfo = {
  exchange: "1.0",
  freight: "0.0",
  original: "USD",
  seo_title: "Default Title",
  seo_keyword: "default,keywords",
  seo_description: "Default description for the site."
};

export const useSiteInfo = async () => {
  return await useAsyncData('siteInfo', async () => {
      try {
          const res = await api.shop.market.siteInfo() as MarketInfo;
          // 如果请求成功但返回的数据为空，使用默认值
          if (!res) {
              return DEFAULT_MARKET_INFO;
          }
          return {
              ...res,
          };
      } catch (error) {
          console.error('Failed to fetch site info:', error);
          // 请求失败时返回默认值
          return DEFAULT_MARKET_INFO;
      }
  });
};