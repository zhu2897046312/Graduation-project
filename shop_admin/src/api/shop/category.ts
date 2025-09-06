import http from '/@/utils/http';

const base_url = '/shop/category';

const api = {
    tree: () => http.get(`${base_url}/tree`),
    create: async (data: any) => http.post(`${base_url}/create`, data),
    modify: async (data: any) => http.post(`${base_url}/modify`, data),
    info: async (id: number) => http.get(`${base_url}/info`, { params: { id } }),
    del: async (id: number) => http.get(`${base_url}/del`, { params: { id } }),
    formatTree: async (top_name: string | any) => {
      const res: any = await api.tree();
      if (top_name && top_name.length > 0) {
        return [{ label: top_name, value: 0, children: [...res] }];
      }
      return [...res];
    }
}


export default api;
