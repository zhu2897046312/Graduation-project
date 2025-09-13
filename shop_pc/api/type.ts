
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

