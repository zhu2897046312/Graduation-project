import http from '/@/utils/http';

export type ResultType = 'LIST' | 'INFO' | 'CREATE' | 'UPDATE' | 'DELETE' | 'MODIFY';

export function curdApi<T>(base_url: string, result_type: ResultType, data: any): Promise<T> {
  let fn: Promise<T>;
  switch (result_type) {
    case 'LIST':
      fn = http.post(`${base_url}/list`, data);
      break;
    case 'INFO':
      fn = http.get(`${base_url}/info`, { params: { id: data } });
      break;
    case 'CREATE':
      fn = http.post(`${base_url}/create`, data);
      break;
    case 'UPDATE':
      fn = http.post(`${base_url}/update`, data);
      break;
    case 'MODIFY':
      fn = http.post(`${base_url}/modify`, data);
      break;
    case 'DELETE':
      fn = http.get(`${base_url}/delete`, { params: { id: data } });
      break;
    default:
      fn = new Promise<T>((_, reject) => {
        reject('不存在通用curl接口');
      });
  }
  return fn;
}
