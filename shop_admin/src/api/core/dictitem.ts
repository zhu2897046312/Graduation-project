import { curdApi } from '../curd';
import type { ResultType } from '../curd';
const base_url = '/core/dictItem';
import http from '/@/utils/http';

export function coreDictItemCurdApi(type: ResultType, data: any) {
  return curdApi<any>(base_url, type, data);
}

export function apiGetList(dictId: number): Promise<any> {
  return http.get(`${base_url}/list`, { params: { dictId } });
}

export function apiGetListByCode(code: string): Promise<any> {
  return http.get(`${base_url}/listByCode`, { params: { dict_code: code } });
}