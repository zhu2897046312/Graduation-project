import ResetPwd from "~/pages/account/reset-pwd.vue";
import { httpRequest } from "~/server/request";



const shop = {
  category: {
    tree: async () => httpRequest.exec('GET', '/shop/category/tree', {}),
    info: async (id: number) => httpRequest.exec('GET', '/shop/category/info', {id}),
    getInfoByCode: async (code: any) => httpRequest.exec('GET', '/shop/category/getInfoByCode', {code}),
    getParents: async (code: any) => httpRequest.exec('GET', '/shop/category/getParents', {code}),
  },
  product: {
    list: async (data: any ) => httpRequest.exec('POST', '/shop/product/list', data),
    info: async (id: number ) => httpRequest.exec('GET', '/shop/product/info', {id}),
  },
  user: {
    register: async (data: any ) => httpRequest.exec('POST', '/shop/userAuth/register', data),
    login: async (data: any ) => httpRequest.exec('POST', '/shop/userAuth/login', data),
    currentUser: async () => httpRequest.exec('GET', '/shop/userAuth/currentUser', {}),
    SendEmail: async (data: any ) => httpRequest.exec('POST', '/shop/userAuth/sentMailResetPassword', data),
    ResetPwd: async (data: any ) => httpRequest.exec('POST', '/shop/userAuth/resetPasswd', data),
  },
  cart: {
    act: async (data: any) => httpRequest.exec('POST', '/shop/userCart/act', data),
    list: async () => httpRequest.exec('POST', '/shop/userCart/list', {}),
  },
  address: {
    list: async (data: any ) => httpRequest.exec('POST', '/shop/userAddress/list', data),
    info: async (id: number ) => httpRequest.exec('GET', '/shop/userAddress/info', {id}),
    create: async (data: any ) => httpRequest.exec('POST', '/shop/userAddress/create', data),
    modify: async (data: any ) => httpRequest.exec('POST', '/shop/userAddress/modify', data),
    del: async (id: number ) => httpRequest.exec('GET', '/shop/userAddress/del', {id}),
  },
  order: {
    create: async (data: any ) => httpRequest.exec('POST', '/shop/order/create', data),
    list: async (data: any ) => httpRequest.exec('POST', '/shop/order/list', data),
    get: async (queryCode : string) => httpRequest.exec('GET', '/shop/order/query-code', { queryCode }),
    getPaymentUrl: async (params: { orderId: string; payType: string }) => httpRequest.exec('POST', `/payment/paypal/create-order?${new URLSearchParams(params)}`, null),
  },
  market: {
    breadcrumb: async (data: any ) => httpRequest.exec('POST', '/shop/market/breadcrumb', data),
    siteInfo: async () => httpRequest.exec('GET', '/shop/market/siteInfo', {}),
    freight: async () => httpRequest.exec('GET', '/shop/market/freight', {}),
  },
  tag: {
    info: async (code: string ) => httpRequest.exec('GET', '/shop/tag/info', {code}),
    product_list: async (data: any ) => httpRequest.exec('POST', '/shop/tag/list', data),
  }
}

const blogs = {
  document: {
    info: async (code: string ) => httpRequest.exec('GET', '/shop/document/info', {code}),
    list: async (data: any) => httpRequest.exec('GET', '/shop/document/list', data)
  },
  recommend: {
    /**
     * 获取推荐位内容
     * @param data 
     * @returns 
     */
    list: async (data: {code: string, page_no: number, page_size: number}) => httpRequest.exec('POST', '/shop/recommendIndex/list', data),
  }
}

export default {
  shop,
  blogs
}
