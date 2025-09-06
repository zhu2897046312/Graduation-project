import { curdApi } from '../curd';
import type { ResultType } from '../curd';
const base_url = '/shop';

export const document = {
  curdApi: (type: ResultType, data: any) => {
    return curdApi<any>(`${base_url}/document`, type, data)
  },
}
