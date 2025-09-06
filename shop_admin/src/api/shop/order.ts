import http from '/@/utils/http';

const base_url = '/shop/order';

const api = {
    list: (data: any) => http.post(`${base_url}/list`, data),
    info: async (id: number) => http.get(`${base_url}/info`, { params: { id } }),
    updateState: (data: any) => http.post(`${base_url}/updateState`, data),
    delivery: (data: any) => http.post(`${base_url}/delivery`, data),
    infoByOrderCode: (data: any) => http.get(`${base_url}/infoByCode`, { params: data }),
}


export default api;
