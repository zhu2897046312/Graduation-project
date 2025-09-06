import { curdApi } from '../curd';
import http from '/@/utils/http';
import type { ResultType } from '../curd';
const base_url = '/shop';

export const document = {
  curdApi: (type: ResultType, data: any) => {
    return curdApi<any>(`${base_url}/document`, type, data)
  },
}


export default {
  recommend: {
    curdApi: (type: ResultType, data: any) => {
      return curdApi<any>(`${base_url}/recommend`, type, data)
    },
  },
  recommendIndex: {
    curdApi: (type: ResultType, data: any) => {
      return curdApi<any>(`${base_url}/recommendIndex`, type, data)
    },
  },
  tag: {
    curdApi: (type: ResultType, data: any) => {
      return curdApi<any>(`${base_url}/tag`, type, data)
    },
  }
}