import http from '/@/utils/http';

const base_url = '/shop/product';

const api = {
    list: (data: any) => http.post(`${base_url}/list`, data),
    create: async (data: any) => http.post(`${base_url}/create`, data),
    modify: async (data: any) => http.post(`${base_url}/modify`, data),
    info: async (id: number) => http.get(`${base_url}/info`, { params: { id } }),
    del: async (id: number) => http.get(`${base_url}/del`, { params: { id } }),
}


export default api;
