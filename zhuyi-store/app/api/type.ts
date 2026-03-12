
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

// 商品列表响应类型
export interface ProductListResponse {
  list: ProductItem[];
  total: number;
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

  // 购物车列表响应类型
  export interface CartListResponse {
    list: CartItem[];
    total: number;
  }

  // export interface SkuConfigItem {
  //   id: number;
  //   product_id: number;
  //   sku_id: number;
  //   prod_attributes_id: number;
  //   prod_attributes_value_id: number;
  //   created_time: string; // ISO 8601 格式的时间字符串
  //   updated_time: string; // ISO 8601 格式的时间字符串
  // }

  // export type SkuConfigList = SkuConfigItem[];

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


// 单个地址的数据类型
export interface Address {
  id: number;
  title: string;
  default_status: number;
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  province: string;
  city: string;
  region: string;
  detail_address: string;
  created_time: string; // 或者可以使用 Date 类型，但需要转换
  updated_time: string; // 或者可以使用 Date 类型，但需要转换
  country: string;
  postal_code: string;
};

export type AddressList = Address[];

// 地址列表响应的数据类型
export interface AddressListResponse {
  list: Address[];
  total: number;
};

// ========== API 请求参数类型 ==========

// 分类相关请求参数
export interface CategoryGetInfoByCodeParams {
  code: string;
}

export interface CategoryGetParentsParams {
  code: string;
}

// 商品相关请求参数
export interface ProductListParams {
  page_no?: number;
  page_size?: number;
  category_id?: number;
  keyword?: string;
  [key: string]: unknown; // 允许其他可选参数
}

// 用户认证相关请求参数
export interface UserRegisterParams {
  email: string;
  username: string;
  password: string;
}

export interface UserLoginParams {
  email: string;
  password: string;
}

export interface UserSendEmailParams {
  email: string;
}

export interface UserResetPwdParams {
  email: string;
  code: string;
  password: string;
}

// 购物车相关请求参数
export interface CartActParams {
  product_id: number;
  sku_id: number;
  quantity: number;
  add: boolean;
}

// 地址相关请求参数
export interface AddressListParams {
  page_no?: number;
  page_size?: number;
  [key: string]: unknown;
}

export interface AddressCreateParams {
  title: string;
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  province: string;
  city: string;
  region?: string;
  detail_address: string;
  country: string;
  postal_code: string;
  default_status?: number;
}

export interface AddressModifyParams extends AddressCreateParams {
  id: number;
}

// 订单相关请求参数
export interface OrderProductItem {
  product_id: number;
  quantity: number;
  sku_id: number;
}

export interface OrderCreateParams {
  product_items: OrderProductItem[];
  pay_type: string;
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  province: string;
  country: string;
  city: string;
  postal_code: string;
  detail_address: string;
}

export interface OrderListParams {
  page_no?: number;
  page_size?: number;
  status?: number;
  [key: string]: unknown;
}

export interface OrderGetPaymentUrlParams {
  order_id: string;
  pay_type: number;
}

export interface OrderCaptureOrderParams {
  orderId: string;
  [key: string]: unknown;
}

// 市场相关请求参数
export interface MarketBreadcrumbParams {
  model?: number;
  mode?: number | string;
  category_code?: string;
  product_id?: number | string;
  [key: string]: unknown;
}

// 标签相关请求参数
export interface TagProductListParams {
  page_no?: number;
  page_size?: number;
  code: string;
  [key: string]: unknown;
}

// 文档相关请求参数
export interface DocumentListParams {
  page_no?: number;
  page_size?: number;
  category_id?: number;
  keyword?: string;
  [key: string]: unknown;
}

// 通用类型（根据项目中 common 包的习惯推测）
type MyID = number;           // 可能是 string化的雪花ID / uint64 / int64
type MyState = number;                 // 通常 0 = 禁用/关闭, 1 = 启用/开启
type MyNumber = number;                // 用于 stock、销量、阅读数等
type MySortNum = number;
type MyType   = number;            // 用于 link_type, document_type 等枚举值

// ────────────────────────────────────────────────
// 商品主信息
export interface Product {
  id: MyID;
  category_id: MyID;
  title: string;
  state: MyState;
  price: number;                  // decimal(10,2)
  original_price: number;
  cost_price: number;
  stock: MyNumber;
  open_sku: MyState;              // 是否开启规格 0/1
  picture: string;                // 主图 URL
  picture_gallery: string[];      // 详情图列表（后端 json.RawMessage 通常转为 string[]）
  description: string;            // 简短描述（200字符限制）
  sold_num: MyNumber;
  version: number;
  sort_num: MySortNum;
  hot: number;                    // uint8，通常 0 或 1
  putaway_time: string | null;    // ISO8601 或 null
  created_time: string;
  updated_time: string;
  deleted_time: string | null;
  detail_url: string;             // 商品详情页链接（H5/PC）
  price_locked: boolean;
}

// ────────────────────────────────────────────────
// 商品详情富文本内容（独立表）
export interface ProductContent {
  id: MyID;
  product_id: MyID;
  content: string;                // 富文本 HTML / Markdown
  seo_title: string;
  seo_keyword: string;
  seo_description: string;
  created_time: string;
  updated_time: string;
}

// ────────────────────────────────────────────────
// 商品属性（规格参数 / 关键卖点）
export interface ProductProperty {
  id: MyID;
  product_id: MyID;
  title: string;                  // 如：材质、重量、产地
  value: string;                  // 如：棉、500g、中国
  sort_num: MySortNum;
  created_time: string;
  updated_time: string;
}

// ────────────────────────────────────────────────
// SKU（单品规格）
export interface Sku {
  id: MyID;
  product_id: MyID;
  sku_code: string;               // 商家编码
  title: string;                  // 如：白色-36码
  price: number;
  original_price: number;
  cost_price: number;
  stock: MyNumber;
  default_show: MyState;          // 是否默认展示该 SKU（0/1）
  state: MyState;
  version: number;
  created_time: string;
  updated_time: string;
  deleted_time: string | null;
}

// SKU 的属性组合索引（用于前端规格选择器）
export interface SkuConfigItem {
  id: MyID;
  product_id: MyID;
  sku_id: MyID;
  prod_attributes_id: MyID;         // 属性组ID（如颜色、尺码）
  prod_attributes_value_id: MyID;   // 属性值ID（如红色、XL）
  created_time: string;
  updated_time: string;
}

// 注意：这里 sku_config_list 通常是一个数组，里面的每一项代表一条 规格-值 的对应关系
// 前端常用来构建规格选择矩阵

// ────────────────────────────────────────────────
// 店铺/商品标签（活动标签、推荐标签等）
export interface ShopTag {
  id: MyID;
  title: string;
  code: string;                   // 标签唯一码
  thumb: string;                  // 标签小图标
  state: MyState;
  read_num: MyNumber;
  sort_num: MySortNum;
  created_time: string;
  updated_time: string;
  deleted_time: string | null;
  match_word: string;             // 匹配关键词（可能用于搜索）
}

// SKU 配置值类型
export interface SkuConfigValue {
  id: MyID;
  title: string;
  sort_num: MySortNum;
}

// SKU 配置类型（用于前端规格选择器）
export interface SkuConfig {
  id: MyID;
  title: string;
  sort_num: MySortNum;
  value: SkuConfigValue[];
}

/**
 * 商品详情页接口返回的 result 数据结构（扁平化结构）
 */
export interface ProductDetailData {
  /** 商品ID */
  id: MyID;
  
  /** 分类ID */
  category_id: MyID;
  
  /** 商品标题 */
  title: string;
  
  /** 商品状态 */
  state: MyState;
  
  /** 商品价格 */
  price: number;
  
  /** 商品原价 */
  original_price: number;
  
  /** 成本价 */
  cost_price?: number;
  
  /** 库存 */
  stock: MyNumber;
  
  /** 是否开启SKU */
  open_sku: MyState;
  
  /** 主图URL */
  picture: string;
  
  /** 详情图列表 */
  picture_gallery: string[];
  
  /** 简短描述 */
  description: string;
  
  /** 销量 */
  sold_num: MyNumber;
  
  /** 排序号 */
  sort_num: MySortNum;
  
  /** 上架时间 */
  putaway_time: string;
  
  /** 商品详情富文本内容（HTML/Markdown） */
  content: string;
  
  /** SEO标题 */
  seo_title: string;
  
  /** SEO关键词 */
  seo_keyword: string;
  
  /** SEO描述 */
  seo_description: string;
  
  /** 商品属性参数列表（如：产地、材质、规格参数等） */
  property_list: ProductProperty[];
  
  /** 该商品的所有 SKU */
  sku_list: Sku[];
  
  /** SKU 配置（用于前端规格选择器构建） */
  sku_config: SkuConfig[];
  
  /** 商品关联的标签（促销标签、推荐标签等），可能为 null */
  tags: ShopTag[] | null;
}

export interface BreadcrumbItem {
 title: string;
 link: string;
}

/**
 * CMS 文档/文章 基本结构（列表 & 详情通用）
 */
export interface CmsDocument {
  /** 文档ID */
  id: MyID;

  /** 标题 */
  title: string;

  /** 分类ID */
  category_Id: MyID;

  /** 关联地点/场所ID（视业务可能是门店、景区等） */
  associated_PlaceId: MyID;

  /** 发布/创建用户ID */
  user_Id: MyID;

  /** 管理员ID（审核或后台操作者） */
  admin_Id: MyID;

  /** 内容模板标识（如：news, article, video 等） */
  cont_Tpl: string;

  /** 关键词（SEO 用，逗号分隔或单值） */
  keyword: string;

  /** 描述/摘要（SEO meta description + 列表页展示） */
  description: string;

  /** 封面图 / 缩略图 URL */
  thumb: string;

  /** 状态（0=草稿, 1=已发布, -1=删除 等） */
  state: MyState;

  /** 链接类型（内部链接/外链/富媒体等） */
  link_type: MyType;

  /** 文档类型（文章/视频/图集/公告 等） */
  document_type: MyType;

  /** 视频时长（秒），非视频为 0 */
  video_duration: number;

  /** 发布时间（定时发布用，可能为空） */
  send_time: string | null;

  /** 作者名（显示用，可与 user_Id 对应） */
  author: string;

  /** 来源（原创/转载/媒体名称） */
  source: string;

  /** 阅读数量 */
  read_num: MyNumber;

  /** 点赞数量 */
  like_num: MyNumber;

  /** 排序值（手动排序用） */
  sort_num: MySortNum;

  /** 创建时间 */
  created_time: string;

  /** 最后更新时间 */
  updated_time: string;

  /** 软删除时间 */
  deleted_time: string | null;

  /** 唯一编码（短链接/邀请码/标识等，可选） */
  code?: string;
}

// ────────────────────────────────────────────────

/**
 * 文档列表接口返回结构（最符合你提供的 GetAll handler）
 */
export interface DocumentListResponse {
  list: CmsDocument[];
  total: number;
}

/**
 * CMS 文档内容存档 / 扩展信息
 * （通常与 CmsDocument 主表一对一关联，存放正文、附件、SEO 等）
 */
export interface CmsDocumentArchive {
  /** 关联的主文档ID（主键） */
  document_id: MyID;

  /** 正文内容（富文本 HTML / Markdown，通常较长） */
  cont: string;

  /**
   * 可下载文件列表
   * 常见格式示例：
   *   string[]                  → ["https://...", "https://..."]
   *   {name: string, url: string}[]  → 最推荐的结构
   *   {title: string, path: string, size?: number, type?: string}[]
   */
  download_files: string[] | DownloadFileItem[];

  /** SEO 标题（覆盖主表的 title，用于搜索引擎） */
  seo_title: string;

  /** SEO 关键词（逗号分隔或空格分隔） */
  seo_keyword: string;

  /** SEO 描述（meta description） */
  seo_description: string;

  /** 软删除时间 */
  deletedTime: string | null;
}

/**
 * 最常见的 download_files 条目结构（推荐使用）
 * 如果后端实际返回的是对象数组，请优先使用这个类型
 */
export interface DownloadFileItem {
  /** 文件显示名称 */
  name?: string;

  /** 文件下载地址 */
  url: string;

  /** 文件大小（字节）- 可选 */
  size?: number;

  /** 文件类型 / MIME - 可选 */
  type?: string;

  /** 上传时间或创建时间 - 可选 */
  created_at?: string;

  // 其他可能的扩展字段
  // ext?: string;         // 文件后缀
  // md5?: string;         // 文件校验
}

export interface DocumentInfoResponse {
  document: CmsDocument;
  archive: CmsDocumentArchive;
}

type OrderState = number;              // 订单状态枚举值（常见：待支付、已支付、已发货等）
type PayType = number;              // 支付方式枚举值（常见：支付宝、微信、银行卡等）

/**
 * 前端订单详情 / 订单查询 VO（单条订单视图对象）
 * 对应 Go 中的 SpOrderFrontQueryVo
 */
export interface OrderFrontQueryVo {
  /** 订单主键ID */
  id: MyID;

  /** 订单编号（显示给用户看的编号，通常更友好） */
  code: string;

  /** 下单用户ID（登录用户） */
  user_id: MyID;

  /** 游客查询码（未登录/游客订单使用） */
  visitor_query_code: string;

  /** 下单人昵称（冗余或游客填写） */
  nickname: string;

  /** 联系邮箱（收货/通知用） */
  email: string;

  /** 订单商品总金额（未优惠前） */
  total_amount: number;

  /** 实际支付金额（优惠/抵扣后） */
  pay_amount: number;

  /** 支付方式（uint8 映射为 number） */
  pay_type: PayType;

  /** 订单状态 */
  state: OrderState;

  /** 支付完成时间（ISO 字符串或 null） */
  payment_time: string | null;

  /** 发货时间（物流开始时间） */
  delivery_time: string | null;

  /** 确认收货时间 */
  receive_time: string | null;

  /** 物流公司名称/代码 */
  delivery_company: string;

  /** 物流单号 */
  delivery_sn: string;

  /** 运费金额 */
  freight: number;
}

/**
 * 前端订单商品项视图对象
 * 对应 Go 中的 SpOrderItemFrontVo
 */
export interface OrderItemFrontVo {
  /** 订单项主键ID */
  id: MyID;

  /** 商品标题（下单时的快照） */
  title: string;

  /** SKU 规格描述（如：红色 / 128GB / M码） */
  sku_title: string;

  /** 商品主图 / 缩略图 URL（下单快照） */
  thumb: string;

  /** 所属订单ID */
  order_id: MyID;

  /** 商品ID */
  product_id: MyID;

  /** 选择的 SKU ID（无规格商品可能为 0 或空） */
  sku_id: MyID;

  /** 该商品项总金额（price × quantity，通常包含优惠前金额） */
  total_amount: number;

  /** 该商品项实际支付金额（优惠/折扣后） */
  pay_amount: number;

  /** 购买数量 */
  quantity: number;

  /** 下单时的单价（快照价格） */
  price: number;

  /** 下单时的划线价 / 原价（快照） */
  original_price: number;
}

// ────────────────────────────────────────────────

/**
 * 订单收货地址信息
 * 对应 Go 中的 SpOrderReceiveAddress
 */
export interface OrderReceiveAddress {
  /** 地址记录ID */
  id: MyID;

  /** 所属订单ID */
  order_id: MyID;

  /** 名字（名） */
  first_name: string;

  /** 姓氏（Last Name） */
  last_name: string;

  /** 联系邮箱 */
  email: string;

  /** 联系电话 */
  phone: string;

  /** 省 / 州 */
  province: string;

  /** 城市 */
  city: string;

  /** 区 / 县 / 郡 */
  region: string;

  /** 详细地址（街道、门牌号等） */
  detail_address: string;

  /** 国家（英文代码或中文名称，视业务） */
  country: string;

  /** 邮政编码 */
  postal_code: string;

  /** 创建时间（ISO 字符串） */
  created_time: string;

  /** 最后更新时间（ISO 字符串） */
  updated_time: string;
}

export interface OrderInfoResponse {
  order: OrderFrontQueryVo;
  items: OrderItemFrontVo[];
  address: OrderReceiveAddress;
}

export enum OrderStatus {
  PendingPayment   = 1,
  PendingShipment  = 2,
  Shipped          = 3,
  Completed        = 4,
  Closed           = 5,
  Invalid          = 6,
}

// Optional: label map
export const OrderStatusLabel: Record<OrderStatus, string> = {
  [OrderStatus.PendingPayment]:   "Pending Payment",
  [OrderStatus.PendingShipment]:  "Pending Shipment",
  [OrderStatus.Shipped]:          "Shipped",
  [OrderStatus.Completed]:        "Completed",
  [OrderStatus.Closed]:           "Closed",
  [OrderStatus.Invalid]:          "Invalid Order",
};