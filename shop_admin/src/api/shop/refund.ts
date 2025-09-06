import http from '/@/utils/http';

const base_url = '/shop/refund';

const api = {
    list: (data: any) => http.post(`${base_url}/list`, data),
    info: async (id: number) => http.get(`${base_url}/info`, { params: { id } }),
    create: async (data: any) => http.post(`/payment/paypal/refund`, data),
}

export default api;