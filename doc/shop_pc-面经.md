# 电商购物平台 PC 端 - 面经（面试准备）

> 基于 Nuxt 4 + Vue 3 + TypeScript 的 C 端电商前台项目，支持 SSR/SSG、SEO，适用于前端/全栈岗位面试时介绍项目、回答技术问题与难点。

---

## 一、项目一分钟介绍（开场话术）

**版本 A（偏业务）：**

> 我参与的是公司电商购物平台 PC 端的前端开发，是一个已上线的 C 端项目。技术栈是 Nuxt 4 + Vue 3 + TypeScript，UI 用的 Naive UI 和 Tailwind。我负责整条购物链路：首页和商品列表用 useAsyncData 做服务端拉数，利于 SEO；商品详情有多规格 SKU 选择、加购；购物车用 useState 做全局数量同步，并封装了 useCartHook 统一加购、登录校验和弹窗提示；结算页有地址表单、省市区级联、PayPal 支付；还有订单列表、用户中心、地址管理等。请求层统一封装了 Token、设备指纹和错误处理。过程中我独立解决了两个线上问题：一是 Nuxt 3.12 之后 Naive UI 在 SSR 下首屏样式闪烁（FOUC），用 css-render 手动注入样式解决；二是服务端和客户端设备指纹不一致，通过 Cookie 持久化加服务端模块级缓存保证同构一致。

**版本 B（偏技术）：**

> 这是一个 Nuxt 4 的电商前台项目，Vue 3 + TypeScript，已上线生产。我重点做了几块：一是 API 和请求封装，统一 baseURL、Authorization、X-Device-Fingerprint，以及业务 code 处理；二是用 useAsyncData 做首屏数据拉取和 SEO，用 useState 做购物车数量等跨页面状态；三是封装了 useCartHook、useSiteInfo 等组合式函数，把加购、登录态校验、站点信息请求复用起来；四是解决了 SSR 下 Naive UI 样式 FOUC 和设备指纹在服务端/客户端不一致的问题。业务上覆盖首页、商品列表/详情、购物车、结算、订单、用户与地址管理，并对接了 PayPal 支付。

---

## 二、技术栈速记

| 类别     | 技术 |
|----------|------|
| 框架     | Nuxt 4、Vue 3、Vue Router 4 |
| 语言     | TypeScript |
| UI       | Naive UI、@nuxt/ui、Tailwind CSS |
| 请求     | $fetch + 自封装 HttpRequest（server/request.ts） |
| 工具     | uuid（设备指纹）、@css-render/vue3-ssr（SSR 样式） |
| 部署     | Nitro（node-server）、开发/生产 API 切换、ISR |

---

## 三、高频面试题与参考答案

### 1. 介绍一下这个电商前台项目？

按「定位 → 技术栈 → 我做的 → 成果」说：

- **定位**：公司电商购物平台 PC 端，C 端用户侧，已上线（如 earring18.com），支持从浏览、加购、结算到支付、订单的完整流程。
- **技术栈**：Nuxt 4 + Vue 3 + TypeScript，Naive UI + Tailwind，请求自封装，对接后端 REST API。
- **我做的**：统一 API 层与请求封装（Token、设备指纹）；首页/列表/详情用 useAsyncData 做 SSR 与 SEO；购物车用 useState 全局数量 + useCartHook 封装加购与登录校验；结算页地址与 PayPal；独立解决 Naive UI SSR 样式 FOUC 与设备指纹一致性问题。
- **成果**：完整购物链路、类型定义与文档、线上问题排查与修复。

---

### 2. Nuxt 下 Naive UI 的 SSR 样式闪烁（FOUC）是怎么解决的？

**现象**：首屏刷新时先出现无样式或错乱，再闪一下才正常，影响体验和 SEO（如 CLS）。

**原因**：Nuxt 3.12+ 移除了对 `ssrContext.styles` 的自动收集，Naive UI 依赖的 `@css-render/vue3-ssr` 生成的样式无法自动注入到服务端输出的 `<head>`，导致首屏没有组件样式。

**做法**（`app/plugins/naive-ui.ts`）：

- 仅在服务端（`import.meta.server`）执行。
- 使用 `@css-render/vue3-ssr` 的 `setup(nuxtApp.vueApp)` 得到样式收集函数 `collect`。
- 在合适的生命周期（如渲染完成后）调用 `collect()` 得到样式 HTML 字符串。
- 将字符串按 `</style>` 拆成多个 style 块，解析出 `cssr-id` 和 `innerHTML`，通过 `nuxtApp.ssrContext.head.push({ style: () => [...] })` 把 style 注入到服务端 HTML 的 head 中。

**结果**：服务端直出的 HTML 已包含 Naive UI 样式，首屏无闪烁，有利于 Core Web Vitals 和 SEO。

---

### 3. SSR 环境下设备指纹（UUID）如何保证服务端和客户端一致？

**需求**：购物车、风控等依赖设备 ID，若服务端渲染时生成一个 UUID、客户端 hydration 又生成一个，同一用户会被识别成两个设备。

**实现要点**（`utils/auth.ts`）：

- **持久化**：用 Nuxt 的 `useCookie('device_id', { maxAge: 1年, sameSite: 'lax' })` 存设备 ID，保证同源下前后一致。
- **服务端**：使用模块级变量 `_deviceId` 缓存。当 Cookie 里没有值时，本次请求内只生成一次 UUID 并赋给 `deviceId.value`，同一请求多次调用 `getDeviceId()` 得到同一值。
- **客户端**：无 Cookie 时再 `uuidv4()` 生成并写入 Cookie，之后都从 Cookie 读。

这样同一次请求（SSR）或同一浏览器（CSR）内设备 ID 一致，接口层在 `server/request.ts` 里把 `getDeviceId()` 放进请求头 `X-Device-Fingerprint` 即可。

---

### 4. 请求层（HttpRequest）是怎么封装的？

**位置**：`server/request.ts`（Nuxt 的 server 目录，运行在服务端；若需客户端也用，需通过 useFetch/$fetch 在应用内调用，实际请求会走 Nuxt 的上下文）。

**要点**：

- **入口**：`exec(method, url, data)`，method 为 GET/POST，GET 时 data 做 params，POST 时做 body。
- **baseURL**：从 `useRuntimeConfig().public.apiUrl` 读，开发/生产可配不同 API 地址。
- **请求头**：`Authorization: useCookie('accessToken').value`；`X-Device-Fingerprint: getDeviceId()`。
- **请求方式**：使用 `$fetch<ResultData>(url, options)` 发请求。
- **响应处理**：`data.code === 0` 时 `return Promise.resolve(data.result)`，业务层直接用 result；否则 `Promise.reject(new Error(data.message))`，由调用方捕获提示（如 Naive UI dialog）。

这样所有接口都统一带 Token、设备指纹，且业务代码只关心 result，错误信息统一从 message 取。

---

### 5. 购物车数量如何跨页面同步？useCartHook 做了什么？

**跨页面同步**：使用 Nuxt 的 `useState('cartNum', () => 0)`。`useState` 在 Nuxt 中是基于 key 的全局状态，同一 key 在所有页面共享，且支持 SSR（服务端取初值、客户端复用），所以导航栏徽标、购物车页、加购后的数量可以一致。

**useCartHook**（`hook/useCartHook.ts`）：

- 暴露 `addCart(product_id, sku_id, quantity)`：调 `api.shop.cart.act({ add: true, ... })`，成功用 `useDialog().success` 提示并可跳转 cart；失败时若为「请先登录」则把 `useState('showLoginModal')` 设为 true，弹出登录框。
- 加购成功后调用内部 `getList()` 拉购物车列表，遍历 item 累加 quantity 更新 `cartNum.value`，这样导航栏数量即时更新。
- 返回 `{ addCart }`，商品详情页等直接 `const { addCart } = useCart()` 即可，逻辑和 UI 提示都集中在一处。

---

### 6. useAsyncData 在项目里怎么用的？和 SEO 的关系？

**用法**：在页面里 `await useAsyncData(key, fetcher)`，fetcher 里调 api（如商品列表、商品详情、站点信息）。key 唯一即可，用于去重和缓存，例如 `product-${productId}`、`hot`、`siteInfo`。

**和 SEO 的关系**：

- 服务端执行时会在 SSR 阶段执行 fetcher，把结果序列化到 HTML，首屏内容已包含数据，爬虫和用户都能直接看到。
- 同时可配合 `<Title>`、`<Meta name="keywords">`、`<Meta name="description">` 使用 `useSiteInfo()` 或接口返回的 seo_title、seo_keyword、seo_description 动态设置，利于搜索引擎收录。

**useSiteInfo**：封装成 `useAsyncData('siteInfo', () => api.shop.market.siteInfo())`，首页等复用，既拿到站点信息又享受 SSR 缓存。

---

### 7. 商品详情页多规格 SKU 是怎么做的？

**数据**：商品接口返回 `sku_list` 以及属性结构（如 SpProductProdFrontVo：属性 id + 属性值列表）。有默认展示的 SKU（如 `default_show === 1`）。

**ProductSkuBox 组件**：

- Props：`list`（属性+值树）、`defaultSelected`（默认选中的 sku_code 或 value 组合）。
- 用 `ref` 维护当前选中 `current_selected: { prop_id, value_id }[]`，在 `onMounted` 里根据 `defaultSelected` 解析并 emit 一次 `change`，把选中的 sku 信息传给父组件（用于价格、库存、加购时 sku_id）。
- 用户点击规格时更新 `current_selected`，通过计算属性得到当前选中的 sku（或不可选状态），并 emit 给父组件；父组件（如 product/[id].vue）用 `useState` 存 `current_sku_id`、`price`、`original_price`，加购时传 `current_sku_id`。

**注意**：规格选中状态用 ref 且在 onMounted 里初始化，避免在 SSR 阶段依赖客户端状态，减少 hydration 报错。

---

### 8. 布局和路由是怎么组织的？

**布局**：`app/layouts/` 下 default（头+导航+页脚）、blank、minilayout（如结算页简化布局），页面通过 `setPageLayout('minilayout')` 或约定文件所在布局切换。

**路由**：Nuxt 文件系统路由。例如 `pages/index.vue` 为首页，`pages/product/[id].vue` 为商品详情，`pages/cart.vue`、`pages/checkout.vue`、`pages/order-detail/[id].vue`、`pages/account/*`、`pages/blogs/[code].vue`、`pages/collections/[code].vue`、`pages/tag/[code].vue`、`pages/search.vue` 等，对应列表、集合、标签、搜索、博客文档等。

**配置**：`nuxt.config.ts` 里可配置 `routeRules`（如 `/blogs` 重定向）、生产环境 `apiUrl`、Nitro preset（如 node-server）、开发环境 API 等。

---

### 9. TypeScript 类型是怎么维护的？

在 `api/type.ts` 中集中定义与后端对齐的接口类型，例如：Category、ProductItem、ProductList、ProductInfo、MarketInfo、CartItem、CartList、Address、订单相关、SpProductProdFrontVo（SKU/属性）、DocumentItem、推荐位等。请求层返回的 `data.result` 在业务层按这些类型做类型断言或泛型，便于联调、重构和减少运行时错误。

---

### 10. 结算页和 PayPal 支付流程大致是怎样的？

**结算页**：地址信息用 `useState('address')` 存表单（含 first_name、last_name、country、province、city、detail_address、postal_code、email、phone 等）；省市区可用本地 address.json 做级联；配送与支付方式（如 PayPal）选择；表单校验用 Naive UI NForm rules；提交时调 `api.shop.order.create(...)` 创建订单，再根据支付方式调 `api.shop.order.getPaymentUrl` 等拿到 PayPal 跳转 URL，跳转至 PayPal 完成支付后回调到订单详情或成功页，再调 capture 等接口确认订单状态。

**简历/口述**：结算页负责收集收货信息与支付方式，创建订单后对接 PayPal 获取支付链接，用户完成支付后通过回调更新订单状态。

---

## 四、可主动抛出的亮点（引导面试官追问）

| 亮点             | 一句话说明 |
|------------------|------------|
| SSR 样式问题     | Nuxt 3.12+ 下用 @css-render/vue3-ssr 手动收集并注入 Naive UI 样式到 head，解决 FOUC，利于 SEO 与 CLS。 |
| 设备指纹一致性   | Cookie 持久化 + 服务端模块级缓存，保证同构渲染下设备 ID 一致，购物车与风控逻辑稳定。 |
| 请求与 API 分层  | server/request 统一 Token、设备指纹、code 处理；api/index 按 shop/blogs 分模块，类型在 type.ts 集中维护。 |
| 状态与数据流     | useState 做购物车数量等全局状态；useAsyncData 做首屏数据与 SEO；useCartHook、useSiteInfo 做逻辑复用。 |
| 完整业务链路     | 首页→列表→详情→加购→购物车→结算→下单→PayPal→订单，参与整条 C 端链路开发与上线。 |

---

## 五、可能的追问与简答

- **为什么用 Nuxt 而不是纯 Vue SPA？**  
  需要 SEO 和首屏速度，Nuxt 提供 SSR/SSG、useAsyncData、文件路由和 Nitro 部署，和现有后端 REST API 配合即可，无需再起 Node 渲染层。

- **useState 和 Pinia 的区别？**  
  本项目用 Nuxt 自带的 useState 做少量全局状态（如 cartNum、showLoginModal），与 SSR 兼容好；若状态很多、需要 devtools 或持久化，可再引入 Pinia，两者可并存。

- **如何保证 useAsyncData 的 key 不冲突？**  
  key 按页面+数据维度设计，如 `product-${productId}`、`recommended-products-${productId}`、`siteInfo`、`hot`，不同页面、不同数据用不同 key，避免缓存串数据。

- **生产环境 API 怎么切换？**  
  在 nuxt.config 的 `$production` 里配置 `runtimeConfig.public.apiUrl` 为生产域名（如 https://www.earring18.com/api/client），Nitro 打包后使用该配置。

---

## 六、STAR 示例（讲一个难点）

**S**：项目上线后，用户反馈首屏刷新时页面会先「秃」一下再出现样式，体验差，也影响 SEO 的 CLS。  
**T**：要在不推翻技术选型的前提下，让 Naive UI 在 Nuxt SSR 下首屏就带样式。  
**A**：查文档和 issue 发现 Nuxt 3.12 移除了对 ssrContext.styles 的收集，而 Naive UI 的 SSR 样式依赖 @css-render/vue3-ssr。我在插件里仅在服务端用 setup 拿到 collect，在合适的时机把收集到的样式按块解析，通过 ssrContext.head.push 注入到当前请求的 HTML head 中，保证直出 HTML 已包含组件样式。  
**R**：首屏 FOUC 消失，Core Web Vitals 和 SEO 表现更好，也积累了同构渲染和第三方 UI 库在 Nuxt 下的排查经验。

---

## 七、简历一句话（可直接用）

> 参与电商购物平台 PC 端（Nuxt 4 + Vue 3 + TypeScript）前端开发，完成商品展示、购物车、结算、订单及用户认证等完整流程并上线；统一封装 API 与请求（Token、设备指纹），独立解决 Nuxt SSR 下 Naive UI 首屏样式 FOUC 及服务端/客户端设备指纹不一致问题，提升首屏体验与 SEO；使用 Naive UI + Tailwind 实现响应式 UI，对接 PayPal 支付与多地址管理，参与 TypeScript 类型与项目文档维护。

---

*可根据实际实习时间、参与模块深度，对「工作职责」「项目时间」等做微调；面试时可按岗位突出 SSR/SEO（前端）、请求与安全（全栈）、业务理解（业务线）。*
