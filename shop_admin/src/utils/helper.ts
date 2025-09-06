import { message } from 'ant-design-vue';

export interface LoadingTaskOptions {
  msg?: string;
  fail?: (e: any) => void;
}

export async function loadingTask(fn: () => Promise<any>, opt?: LoadingTaskOptions) {
  opt = opt || {};
  const msg = opt.msg || '加载中...';
  // const hide = message.loading(msg, 0);
  try {
    await fn();
  } catch (e) {
    console.warn(e);
    opt.fail && opt.fail(e);
  } finally {
    // hide();
  }
}