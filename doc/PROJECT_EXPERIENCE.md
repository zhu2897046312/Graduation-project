# 项目经历：电商购物平台 PC 端

## 一、项目基本信息

| 公司 | 广东真格软件有限公司 |
|------|----------------------|
| 岗位 | 初级前端工程师（实习） |
| 项目名称 | 电商购物平台（PC 端） |
| 项目类型 | 公司实际业务项目 / C 端电商前端 |
| 技术栈 | Nuxt 4、Vue 3、TypeScript、Naive UI、Tailwind CSS |
| 项目定位 | C 端电商前台，支持 SSR/SSG，注重 SEO 与首屏体验 |

---

## 二、项目描述

在实习期间参与公司 **电商购物平台 PC 端** 的前端开发。项目基于 **Nuxt 4** 构建，采用 Vue 3 + TypeScript 技术栈，实现从商品浏览、加购、结算到订单查看的完整购物流程。对接后端 REST API，支持用户注册登录、购物车、多地址、PayPal 支付等能力，并针对 SEO 与 SSR 做了专门优化，已用于生产环境（如 earring18.com 等站点）。

---

## 三、技术栈与职责

### 3.1 技术栈

- **框架**：Nuxt 4、Vue 3、Vue Router 4  
- **语言**：TypeScript  
- **UI**：Naive UI、@nuxt/ui、Tailwind CSS  
- **工具**：uuid（设备指纹）、@css-render/vue3-ssr（SSR 样式）  
- **部署/运行**：Nitro（Node Server）、开发/生产环境 API 切换  

### 3.2 工作职责

- 负责电商 PC 端前端页面开发与部分架构设计  
- 统一封装 API 层与 HTTP 请求（Token、设备指纹、错误处理）  
- 实现商品列表/详情、购物车、结算、订单、用户中心等核心页面与交互  
- 独立排查并解决 Nuxt SSR 下 Naive UI 样式延迟、设备指纹一致性等线上/开发环境问题  
- 编写 TypeScript 类型定义与项目文档（架构说明、技术难点记录）  

---

## 四、核心功能模块

### 4.1 商品模块

- **首页**：热门商品、新品推荐，基于 `useAsyncData` 做服务端数据拉取与 SEO  
- **商品列表**：分页、排序、筛选；支持按分类、标签查看  
- **商品详情**：图片画廊、多规格 SKU 选择（`ProductSkuBox`）、属性展示、详情/规格 Tab  
- **分类**：多级分类树、分类页、面包屑（对接 `market.breadcrumb`）  

### 4.2 购物车模块

- 加购、改数量、删除；实时总价与数量徽标  
- 使用 `useState('cartNum')` 全局同步购物车数量  
- 自定义 Hook `useCartHook` 封装加购、拉列表、登录态校验与 Naive UI 弹窗提示  

### 4.3 订单与结算

- **结算页**：收货地址选择/表单（省市区级联）、配送与支付方式  
- **订单**：创建订单、订单列表、订单详情（含 PayPal 支付流程对接）  
- 地址管理：列表、新增、编辑、删除、默认地址  

### 4.4 用户与认证

- 注册、登录、找回密码（邮箱验证）  
- Token 存 Cookie，请求头自动携带；未登录加购时弹出登录模态框  
- 设备指纹（UUID + Cookie）用于游客识别与风控  

### 4.5 其他

- **搜索**：关键词搜索与结果展示  
- **内容**：博客/文档详情、推荐位（recommend）  
- **多布局**：default（头+导航+页脚）、blank、minilayout（如结算页）  

---

## 五、技术亮点与难点

### 5.1 Nuxt 3.12+ 下 Naive UI 的 SSR 样式问题

- **问题**：刷新首屏出现 FOUC（样式延迟），影响体验与 SEO。  
- **原因**：Nuxt 3.12 移除 `ssrContext.styles`，Naive UI 的 SSR 样式无法注入 `<head>`。  
- **做法**：在 `app/plugins/naive-ui.ts` 使用 `@css-render/vue3-ssr` 的 `setup` 收集样式，通过 `nuxtApp.ssrContext.head.push` 手动注入解析后的 style 块，保证服务端直出时样式已存在。  
- **效果**：消除首屏样式闪烁，利于 SEO 与 Core Web Vitals（如 CLS）。  

### 5.2 SSR 环境下设备指纹（UUID）一致性

- **问题**：服务端与客户端各执行一次，若直接每次 `uuidv4()` 会得到不同值，导致同一用户设备识别不一致。  
- **做法**：在 `utils/auth.ts` 中通过 Cookie 持久化 `device_id`；服务端用模块级变量缓存当次请求的 UUID，客户端在无 Cookie 时再生成并写入 Cookie，保证同一次请求/同一浏览器指纹一致。  
- **效果**：购物车、风控等依赖设备 ID 的逻辑可稳定工作。  

### 5.3 分层架构与数据流

- **API 层**（`api/index.ts`）：按业务划分 shop（category、product、user、cart、address、order、market、tag）与 blogs（document、recommend），统一出口。  
- **请求层**（`server/request.ts`）：封装 `HttpRequest.exec`，统一 baseURL、Token、`X-Device-Fingerprint`、错误与业务 code 处理。  
- **页面**：文件系统路由 + `useAsyncData`/`useState`，组合式函数（如 `useSiteInfo`）与 Hook（如 `useCartHook`）复用逻辑。  

### 5.4 类型与可维护性

- 在 `api/type.ts` 中集中定义 Category、ProductItem、ProductInfo、CartItem、Address、SKU/属性相关 VO 及枚举，便于接口联调与重构。  

---

## 六、项目收获与可优化方向

### 6.1 收获

- 在实习中掌握 Nuxt 4 全栈能力：文件路由、SSR、`useAsyncData`、运行时配置、Nitro 部署。  
- 熟练使用 Vue 3 组合式 API、`<script setup>`、Composables 与 Hooks 进行模块化开发。  
- 在实际业务中独立解决 SSR 样式与设备指纹一致性问题，加深对同构渲染与生产环境排查的理解。  
- 完整参与电商前台链路开发（商品→购物车→结算→订单→支付），养成类型安全与接口规范化习惯。  

### 6.2 可优化方向（可与面试结合）

- 性能：图片懒加载、代码分割、ISR/缓存策略细化。  
- 功能：收藏、评价、优惠券、多语言。  
- 工程：引入 Pinia 做集中状态、补充单元测试与 E2E、CI/CD。  

---

## 七、简历用简短描述（可直接粘贴）

**广东真格软件有限公司 | 初级前端工程师（实习）**  
**电商购物平台 PC 端**（Nuxt 4 + Vue 3 + TypeScript）  
- 负责电商前台核心页面开发，完成商品展示、购物车、结算、订单及用户认证等完整流程，项目已上线生产环境。  
- 统一封装 API 与 HTTP 请求（Token、设备指纹），独立解决 Nuxt SSR 下 Naive UI 样式 FOUC 及服务端/客户端设备指纹不一致问题，提升首屏体验与 SEO。  
- 使用 Naive UI + Tailwind CSS 实现响应式 UI，对接 PayPal 支付与多地址管理，参与 TypeScript 类型体系与项目文档维护。  

---

*可根据实际实习时间、参与模块深度，对「工作职责」「项目时间」等做微调。*
