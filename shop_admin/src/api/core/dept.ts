import http from '/@/utils/http';

const base_url = '/core/dept';

export function apiGetDeptList(): Promise<any[]> {
  return http.get(`${base_url}/tree`);
}

export function apiCreateDept(data: any): Promise<any> {
  return http.post(`${base_url}/create`, data);
}

export function apiUpdateDept(data: any): Promise<any> {
  return http.post(`${base_url}/update`, data);
}

export function apiGetDeptInfo(id: number): Promise<any> {
  return http.get(`${base_url}/info`, { params: { id } });
}

export async function apiFormatApiTree(top_name: string | any) {
  const res: any = await apiGetDeptList();
  if (top_name && top_name.length > 0) {
    return [{ label: top_name, value: 0, children: [...res] }];
  }
  return [...res];
}
