import http from '/@/utils/http';

const base_url = '/shop/prodAttributesValue';

const api = {
    list: (data: any) => http.post(`${base_url}/list`, data),
    create: async (data: any) => http.post(`${base_url}/create`, data),
    modify: async (data: any) => http.post(`${base_url}/modify`, data),
    info: async (id: number) => http.get(`${base_url}/info`, { params: { id } }),
    del: async (id: number) => http.get(`${base_url}/del`, { params: { id } }),
    getProdValues:  async (id: number) => {
        const res: any = await api.list({page_no: 1, page_size: 300, prod_attributes_id: id})
        return res.list
    },
}


export default api;
