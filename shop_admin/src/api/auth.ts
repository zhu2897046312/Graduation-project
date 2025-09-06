import http from '/@/utils/http';

const base_url = '/core/auth';

export function apiAuthLogin(data: any) {
  return http.post<string>(`${base_url}/login`, data);
}

export function apiGetAuthInfo() {
  return http.get<any>(`${base_url}/info`);
}

export function apiGetEnumDict() : Promise<any> {
  return http.get<any>(`${base_url}/enumDict`);
}

export function apiChangePwd(data: any) {
  return http.post<any>( `${base_url}/changePwd`, data);
}
