import http from '/@/utils/http';

const base_url = '/mp/user';

const api = {
    list: (data: any) => http.post(`${base_url}/list`, data),
    del: async (id: number) => http.get(`${base_url}/delete`, { params: { id } }),
    updateState: (data: any) => http.post(`${base_url}/updateState`, data),
    delivery: (data: any) => http.post(`${base_url}/delivery`, data),
    infoByOrderCode: (data: any) => http.get(`${base_url}/infoByCode`, { params: data }),
}


export default api;
