interface DictItem {
  label: string;
  value: string | number;
  color?: string;
}

import { apiGetEnumDict } from '/@/api/auth';

const dict: Map<String, DictItem[]> = new Map();

function initDict(source: any[]) {
  for (const it of source) {
    dict.set(it.code, it.items);
  }
  console.debug('dict', dict);
}

function getDictInfo(dict_name: string, value: string | number): DictItem | null {
  if (dict.size === 0) {
    apiGetEnumDict().then((res: any) => {
      initDict(res)
    })
  }
  const items: DictItem[] | undefined = dict.get(dict_name);
  if (items == undefined) {
    return null;
  }
  for (const it of items) {
    if (it.value === value) {
      return it;
    }
  }
  return null;
}

function getDictList(dict_name: string){
  const items: DictItem[] | undefined = dict.get(dict_name);
  return items
}

export default {
  initDict,
  getDictInfo,
  getDictList,
};
