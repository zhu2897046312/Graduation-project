
export interface Category {
    id: number;
    pid: number;
    title: string;
    code: string;
    state: number;
    icon: string;
    picture: string;
    description: string;
    sort_num: number;
    seo_title: string;
    seo_keyword: string;
    seo_description: string;
    created_time: string;
    updated_time: string;
    deleted_time: string | null;
}

export interface ProductItem {
    id: number;
    category_id: number;
    title: string;
    state: number;
    price: number;
    original_price: number;
    cost_price: number;
    stock: number;
    picture: string;
    sold_num: number;
    sort_num: number;
    putaway_time: string;
    open_sku: number; // 根据数据是 number 类型
    created_time: string;
    updated_time: string;
    category: Category | null;
}

export type ProductList = ProductItem[];


// type.ts
export interface MarketInfo {
    exchange: string; // 汇率值
    freight: string;   // 运费
    original: string;  // 原始货币值或标识
    seo_title : string;
    seo_keyword : string;
    seo_description : string;
}

// type.ts
export interface CartItem {
    id: number;
    user_id: number;
    product_id: number;
    sku_id: number;
    fingerprint: string;
    title: string;
    sku_title: string;
    sku_code: string;
    thumb: string;
    total_amount: number;
    pay_amount: number;
    quantity: number;
    price: number;
    original_price: number;
    created_time: string;
    updated_time: string;
    deleted_time: string | null;
  }
  
  // 购物车商品列表类型
  export type CartList = CartItem[];

  // type.ts
export interface DocumentItem {
    id: number;
    title: string;
    category_id: number;
    associated_placeId: number;
    user_id: number;
    admin_id: number;
    cont_tpl: string;
    keyword: string;
    description: string;
    thumb: string;
    state: number;
    link_type: number;
    link: string;
    document_type: number;
    video_duration: number;
    send_time: string | null;
    author: string;
    source: string;
    read_num: number;
    like_num: number;
    sort_num: number;
    created_time: string;
    updated_time: string;
    deleted_time: string | null;
    code: string;
  }
  
  // 文档列表类型
  export type DocumentList = DocumentItem[];

  export interface SkuConfigItem {
    id: number;
    product_id: number;
    sku_id: number;
    prod_attributes_id: number;
    prod_attributes_value_id: number;
    created_time: string; // ISO 8601 格式的时间字符串
    updated_time: string; // ISO 8601 格式的时间字符串
  }

  export type SkuConfigList = SkuConfigItem[];

  // 商品状态枚举
export enum SpProductStatus {
    ONLINE = 1, // 上线
    OFFLINE = 2, // 下线
    PENDING_REVIEW = 3, // 待审核
    REVIEW_FAILED = 4, // 审核不通过
  }
  
  // SKU 是否开启枚举
  export enum SpProductOpenSku {
    CLOSED = 0, // 未开启（根据常见约定，0通常表示关闭）
    OPENED = 1, // 已开启
  }
  
  // SKU 默认显示枚举
  export enum SpProductSkuDefaultShow {
    NOT_DEFAULT = 0, // 非默认显示
    DEFAULT = 1, // 默认显示
  }
  
  // SKU 状态枚举
  export enum SpProductSkuStatus {
    ENABLED = 1, // 启用
    DISABLED = 2, // 停用
  }
  
  // 标签状态枚举
  export enum SpTagStatus {
    PUBLISHED = 1, // 已发布
    UNPUBLISHED = 2, // 未发布
  }

  // 商品属性值前端VO
export interface SpProductProdValueFrontVo {
    id: number;
    title: string;
    sort_num: number;
  }
  
  // 商品属性前端VO
  export interface SpProductProdFrontVo {
    id: number;
    title: string;
    sort_num: number;
    value: SpProductProdValueFrontVo[];
  }
  
  // 商品SKU前端信息VO
  export interface SpProductSkuFrontInfoVo {
    id: number;
    sku_code: string;
    title: string;
    price: number;
    original_price: number;
    stock: number;
    default_show: SpProductSkuDefaultShow;
    state: SpProductSkuStatus;
  }
  
  // 商品属性实体
  export interface SpProductPropertyEntity {
    id: number;
    product_id: number;
    title: string;
    value: string;
    sort_num: number;
    created_time: string;
    updated_time: string;
  }
  
  // 标签实体
  export interface SpTagEntity {
    id: number;
    title: string;
    code: string;
    thumb: string;
    state: SpTagStatus;
    read_num: number;
    sort_num: number;
    created_time: string;
    updated_time: string;
    deleted_time: string | null;
    match_word: string;
  }

  // 商品前端信息VO（主类型）
export interface ProductInfo{
    id: number;
    category_id: number;
    title: string;
    state: SpProductStatus;
    price: number;
    original_price: number;
    stock: number;
    picture: string;
    picture_gallery: string[];
    description: string;
    sold_num: number;
    open_sku: SpProductOpenSku;
    sort_num: number;
    putaway_time: string;
    content: string;
    seo_title: string;
    seo_keyword: string;
    seo_description: string;
    property_list: SpProductPropertyEntity[];
    sku_list: SpProductSkuFrontInfoVo[];
    sku_config: SpProductProdFrontVo[];
    tags: SpTagEntity[];
  }

