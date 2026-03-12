import { httpRequest } from "../utils/request";
import type {
  Category,
  ProductDetailData,
  ProductListResponse,
  CartListResponse,
  Address,
  AddressListResponse,
  ProductListParams,
  UserRegisterParams,
  UserLoginParams,
  UserSendEmailParams,
  UserResetPwdParams,
  CartActParams,
  AddressListParams,
  AddressCreateParams,
  AddressModifyParams,
  OrderCreateParams,
  OrderListParams,
  OrderGetPaymentUrlParams,
  OrderCaptureOrderParams,
  MarketBreadcrumbParams,
  BreadcrumbItem,
  MarketInfo,
  TagProductListParams,
  SpTagEntity,
  DocumentListParams,
  DocumentListResponse,
  DocumentInfoResponse,
  OrderInfoResponse,
} from "./type";

const shop = {
  category: {
    tree: async (): Promise<Category[]> => httpRequest.exec('GET', '/shop/category/tree', {}),
    info: async (id: number): Promise<Category> => httpRequest.exec('GET', '/shop/category/info', {id}),
    getInfoByCode: async (code: string): Promise<Category> => httpRequest.exec('GET', '/shop/category/getInfoByCode', {code}),
    getParents: async (code: string): Promise<Category[]> => httpRequest.exec('GET', '/shop/category/getParents', {code}),
  },
  product: {
    list: async (data: ProductListParams): Promise<ProductListResponse> => httpRequest.exec('POST', '/shop/product/list', data),
    info: async (id: number | string): Promise<ProductDetailData> => httpRequest.exec('GET', '/shop/product/info', {id: Number(id)}),
  },
  user: {
    register: async (data: UserRegisterParams): Promise<unknown> => httpRequest.exec('POST', '/shop/userAuth/register', data),
    login: async (data: UserLoginParams): Promise<unknown> => httpRequest.exec('POST', '/shop/userAuth/login', data),
    currentUser: async (): Promise<unknown> => httpRequest.exec('GET', '/shop/userAuth/currentUser', {}),
    SendEmail: async (data: UserSendEmailParams): Promise<unknown> => httpRequest.exec('POST', '/shop/userAuth/sentMailResetPassword', data),
    ResetPwd: async (data: UserResetPwdParams): Promise<unknown> => httpRequest.exec('POST', '/shop/userAuth/resetPasswd', data),
  },
  cart: {
    act: async (data: CartActParams): Promise<unknown> => httpRequest.exec('POST', '/shop/userCart/act', data),
    list: async (): Promise<CartListResponse> => httpRequest.exec('POST', '/shop/userCart/list', {}),
  },
  address: {
    list: async (data: AddressListParams): Promise<AddressListResponse> => httpRequest.exec('POST', '/shop/userAddress/list', data),
    info: async (id: number): Promise<Address> => httpRequest.exec('GET', '/shop/userAddress/info', {id}),
    create: async (data: AddressCreateParams): Promise<Address> => httpRequest.exec('POST', '/shop/userAddress/create', data),
    modify: async (data: AddressModifyParams): Promise<Address> => httpRequest.exec('POST', '/shop/userAddress/modify', data),
    del: async (id: number): Promise<unknown> => httpRequest.exec('GET', '/shop/userAddress/del', {id}),
  },
  order: {
    create: async (data: OrderCreateParams): Promise<string> => httpRequest.exec('POST', '/shop/order/create', data),
    list: async (data: OrderListParams): Promise<unknown> => httpRequest.exec('POST', '/shop/order/list', data),
    get: async (queryCode: string): Promise<OrderInfoResponse> => httpRequest.exec('GET', '/shop/order/query-code', { queryCode }),
    getPaymentUrl: async (data: OrderGetPaymentUrlParams): Promise<{ approveUrl?: string }> => httpRequest.exec('POST', `/payment/paypal/create-order`, data),
    captureOrder: async (data: OrderCaptureOrderParams): Promise<unknown> => httpRequest.exec('POST', `/payment/paypal/capture-order`, data),
  },
  market: {
    breadcrumb: async (data: MarketBreadcrumbParams): Promise<BreadcrumbItem[]> => httpRequest.exec('POST', '/shop/market/breadcrumb', data),
    siteInfo: async (): Promise<MarketInfo> => httpRequest.exec('GET', '/shop/market/siteInfo', {}),
    freight: async (): Promise<unknown> => httpRequest.exec('GET', '/shop/market/freight', {}),
  },
  tag: {
    info: async (code: string): Promise<SpTagEntity> => httpRequest.exec('GET', '/shop/tag/info', {code}),
    product_list: async (data: TagProductListParams): Promise<ProductListResponse> => httpRequest.exec('POST', '/shop/tag/list', data),
  }
}

const blogs = {
  document: {
    info: async (code: string): Promise<DocumentInfoResponse> => httpRequest.exec('GET', '/shop/document/info', {code}),
    list: async (data: DocumentListParams): Promise<DocumentListResponse> => httpRequest.exec('GET', '/shop/document/list', data)
  },
  recommend: {
    /**
     * 获取推荐位内容
     * @param data 
     * @returns 
     */
    list: async (data: {code: string, page_no: number, page_size: number}): Promise<unknown> => httpRequest.exec('POST', '/shop/recommendIndex/list', data),
  }
}

export default {
  shop,
  blogs
}
