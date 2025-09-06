import { curdApi } from '../curd';
import type { ResultType } from '../curd';
const base_url = '/core/dict';

export function coreDictCurdApi(type: ResultType, data: any) {
  return curdApi<any>(base_url, type, data);
}