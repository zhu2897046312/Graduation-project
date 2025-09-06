import http from '/@/utils/http';

const base_url = '/core/permission';

export function apiGetPermissionList(): Promise<any[]> {
  return http.get(`${base_url}/list`);
}

export function apiGetPermissionTopList(): Promise<any[]> {
  return http.get(`${base_url}/topList`);
}

export function apiGetPermissionUrls(): Promise<any[]> {
  return http.get(`${base_url}/urls`);
}

export function apiCreatePermission(data: any) {
  return http.post(`${base_url}/create`, data);
}

export function apiUpdatePermission(data: any) {
  return http.post(`${base_url}/update`, data);
}

export function apiGetPermissionInfo(id: number): Promise<any> {
  return http.get(`${base_url}/info`, { params: { id } });
}

export async function apiGetPermissionListTree() {
  const res = await apiGetPermissionList();
  const top = res.filter((it) => {
    return it.pid === 0;
  });
  const sec = res.filter((it) => {
    return it.pid > 0;
  });
  return top.map((it) => {
    const child = sec.filter((ic) => {
      return ic.pid === it.id;
    });
    console.log(child);
    if (child && child.length > 0) {
      it.children = child;
    }
    return it;
  });
}
