import http from '/@/utils/http';

const base_url = '/shop/marketSetting';

const api = {
    save: async (data: any) => http.post(`${base_url}/saveMarketSetting`, data),
    info: () => http.post(`${base_url}/info`, {}),
    saveSiteInfo: async (data: any) => http.post(`${base_url}/saveSiteInfo`, data),
    siteInfo: () => http.post(`${base_url}/siteInfo`, {}),
}


export default api;
