import { curdApi } from '../curd';
import type { ResultType } from '../curd';
const base_url = '/core/admin';

export function coreAdminCurdApi(type: ResultType, data: any) {
  return curdApi<any>(base_url, type, data);
}