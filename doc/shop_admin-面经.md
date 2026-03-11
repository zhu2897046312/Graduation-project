# 电商 ERP 管理后台 - 面经（面试准备）

> 基于 Vue 3 + TypeScript + Ant Design Vue 的 B 端管理后台项目，适用于前端/全栈岗位面试时介绍项目、回答技术问题与难点。

---

## 一、项目一分钟介绍（开场话术）

**版本 A（偏业务）：**

> 我负责的是电商独立站的管理后台前端，是一个实际落地的 B 端 ERP 系统。技术栈是 Vue 3 + TypeScript + Vite + Ant Design Vue，对接 Go 后端的 RESTful API。我主要做了几块：一是搭好工程和通用能力，比如 Axios 统一拦截、Token 注入、登录态失效跳转；二是封装了 API 驱动的表格、图片裁剪上传、富文本、ECharts 图表等通用组件，在商品、订单、内容等 10+ 列表页复用；三是做了 RBAC 权限，用自定义指令做按钮级权限控制；四是完成了仪表盘、商品多规格 SKU、订单发货退款、内容与站点配置等业务模块。整体大约 30+ 页面、60+ 接口，注重类型安全和代码复用。

**版本 B（偏技术）：**

> 这是一个电商 ERP 管理后台的前端项目，Vue 3 Composition API + TypeScript + Vite。我重点做了三方面：工程上统一了请求封装和路由守卫；组件上封装了 ApiBasicTable、图片裁剪上传、字典组件等，实现 API 驱动列表和表单的复用；权限上做了 RBAC + v-permission 指令做按钮级控制。业务上覆盖商品（含多规格 SKU）、订单、内容、站点、系统管理，仪表盘用 ECharts 做趋势图。可以重点展开讲表格组件设计、权限指令实现和登录态处理。

---

## 二、技术栈速记

| 类别     | 技术 |
|----------|------|
| 框架     | Vue 3 (Composition API) |
| 语言     | TypeScript |
| 构建     | Vite 5 |
| UI       | Ant Design Vue 4.x |
| 状态     | Pinia |
| 路由     | Vue Router 4 (Hash) |
| HTTP     | Axios |
| 图表     | ECharts 5 |
| 富文本   | WangEditor 5 |
| 其他     | dayjs、xlsx、cropperjs、Tailwind CSS |

---

## 三、高频面试题与参考答案

### 1. 介绍一下你这个管理后台项目？

按「业务定位 → 技术栈 → 我做的 → 成果」说：

- **定位**：电商独立站 ERP 管理后台，B 端 SaaS，对接后端 RESTful API，支撑商品、订单、内容、站点、权限等日常运营。
- **技术栈**：Vue 3 + TypeScript + Vite + Ant Design Vue + Pinia，图表 ECharts，富文本 WangEditor。
- **我做的**：工程搭建与请求/路由封装；通用组件（ApiBasicTable、图片裁剪、富文本、字典）；RBAC + v-permission；仪表盘与各业务模块开发。
- **成果**：约 30+ 页面、15+ 封装组件、60+ 接口对接，TypeScript 与 ESLint 规范。

---

### 2. ApiBasicTable 是怎么设计的？为什么能 API 驱动？

**思路**：把「谁请求、请求参数、怎么展示」都通过 props 和插槽交给父组件，表格内部只负责分页、请求、Loading 和表格渲染。

**实现要点**：

- **Props**：`api` 为请求函数，签名为 `(params) => Promise<{ list, total }>`；`searchParam` 为搜索/筛选参数对象；`columns` 为列配置；还有 `showPage`、`rowKey`、`actionWidth` 等。
- **内部逻辑**：`handleLoadData` 里把 `searchParam` 和当前分页（`page_no`、`page_size`）合并后调用 `prop.api(req)`，拿到 `list`、`total` 更新表格数据和分页；分页变化时再次调用同一套逻辑。
- **插槽**：通过 `#bodyCell` 暴露列和行数据，父组件可自定义列内容（如操作按钮、标签、图片）。
- **暴露方法**：`defineExpose` 暴露 `useReload(reload)`，父组件搜索时调用 `tableRef.useReload(true)` 从第一页重新请求。

这样每个列表页只需传 `api`、`searchParam`、`columns` 和必要的插槽，就能复用分页、加载、高度计算等逻辑。

---

### 3. 权限指令 v-permission 是怎么实现的？

**需求**：按后端下发的权限码控制按钮/区域显示，无权限则不渲染该 DOM。

**实现**（`directive/permission.ts`）：

- 在 `mounted` 里拿到指令绑定值 `binding.value`，约定为权限码字符串，支持逗号分隔多码（满足其一即显示）。
- 从 Pinia 的 `useAuthStore()` 取当前用户 `user.permission` 数组（登录后由后端接口写入）。
- 若 `binding.value` 中任意一个权限码在 `user.permission` 里，则不做处理；否则用 `el.parentNode.removeChild(el)` 把该节点从 DOM 移除。

**注意**：是移除 DOM 而不是 v-if，这样无权限时按钮完全不渲染，避免仅靠样式隐藏带来的安全隐患。

---

### 4. 登录态失效（如 Token 过期）是怎么处理的？

**流程**：后端在 Token 无效时返回固定业务码（如 `code === 18000`），前端在 Axios 响应拦截器里统一处理。

**实现**（`utils/http.ts`）：

- 在 `response interceptor` 里判断 `res.data.code === 18000`。
- 调用 `useAuthStore().loginOut()`：清空 Pinia 里的 user、token，并移除 localStorage 的 `YEX_AUTH`。
- 然后 `window.location.href = '/#/login'` 跳转登录页。

这样所有接口的 18000 都会自动触发登出和跳转，无需在每个请求里单独写。

---

### 5. 路由守卫里做了哪些事？

**时机**：`router.beforeEach`，在进入非登录页时执行。

**逻辑**：

1. 若目标不是 Login：
   - 没有 `currentToken`（内存或 localStorage 的 `YEX_AUTH`）→ 直接 `next({ name: 'Login' })`。
   - 有 Token 但没有 user 信息 → 调 `apiGetAuthInfo()` 拉用户信息，并拉字典 `apiGetEnumDict()` 做全局字典初始化；成功则 `authStore.setUser(user)` 再 `next()`，失败则提示并 `next({ name: 'Login' })`。
2. 最后 `next()` 放行。

这样既保证未登录进不去后台，又保证进入后台时一定有用户信息和字典，供权限指令和菜单使用。

---

### 6. 多规格 SKU 表单是怎么联动的？（商品录入）

**场景**：商品有「开启 SKU / 关闭 SKU」。开启时用规格表（如颜色、尺码）生成多条 SKU，每条有价格、库存等；需要把「默认规格」的价格/库存等同步到主表单展示或提交。

**实现要点**（以 add.vue 为例）：

- 用 `watch` 监听 `sku_list`（或 `data.sku_list`），当列表变化且存在默认 SKU（如 `default_show === 1`）或至少有一条时，取默认 SKU 或第一条的 price、stock 等，写回主表单的 `data.price`、`data.stock` 等，用于展示或提交时的主信息。
- 关闭 SKU 时，主表单直接展示单品的价格、库存输入；开启 SKU 时隐藏主表单的单价/库存输入，改为展示规格配置组件（如 SkuForm），提交时校验 `sku_list` 非空且存在默认规格。

**可补充**：规格配置用属性/属性值（如颜色、尺码）组合成 SKU 行，用 `sku_config` 等结构维护属性与值的映射，便于和后台接口对齐。

---

### 7. Axios 请求封装做了哪些事？

- **创建实例**：`axios.create` 配置 `baseURL`（如 `/api/manage`）、`timeout`。
- **请求拦截**：从 `useAuthStore().currentToken` 取 Token，写入 `config.headers.Authorization`，保证每次请求都带认证。
- **响应拦截**：
  - 非 200 提示「网络异常」并 reject。
  - `code === 18000` 走登录态失效逻辑（见上）。
  - 其他非 0 的 code 用 `message.warn` 提示 `res.data.message` 并 reject。
  - `code === 0` 时 `return result`，这样业务代码直接拿到的就是 `result`，不用再解一层。

---

### 8. Pinia 里登录态是怎么存、怎么用的？

- **state**：`user`（含 nickname、id、avatar、permission 等）、`token`。
- **getter**：`currentToken` 优先用 state 的 token，为空则从 `localStorage.getItem('YEX_AUTH')` 取并回写 state，保证刷新后仍能拿到 Token。
- **actions**：`setToken` 同时写 state 和 localStorage；`setUser` 存用户信息；`loginOut` 清空 user、token 并移除 localStorage。
- **使用**：路由守卫、Axios 拦截器、权限指令、菜单展示等统一用 `useAuthStore()` 取 Token 和 user，保证登录态一致。

---

### 9. 为什么用 Vite 而不是 Vue CLI？

- **开发体验**：Vite 基于 ESM，开发时按需编译，冷启动快；Vue CLI 基于 Webpack 全量打包，项目大时启动和 HMR 会慢一些。
- **构建**：Vite 生产用 Rollup，Tree-shaking 和 chunk 分割更细，有利于体积和缓存。
- **配置**：Vite 配置更简洁，和 Vue 3、TypeScript 结合顺畅，适合新项目。

---

### 10. 仪表盘图表（ECharts）是怎么封装的？

- 使用 ECharts 做「用户增长趋势」「订单增长趋势」等折线图。
- 封装成通用组件（如 GrowthChart），通过 props 传入：`dataList`、`dateField`（日期字段名）、`mode`（累计 cumulative / 每日 daily）、`chartColor`、`title` 等。
- 组件内部根据 `mode` 做数据聚合（按日汇总或累计），再调用 ECharts 的 setOption 渲染，便于在仪表盘多个图表间复用。

---

## 四、可主动抛出的亮点（引导面试官追问）

| 亮点           | 一句话说明 |
|----------------|------------|
| 组件复用       | ApiBasicTable、YexUpload、YexDict 等在 10+ 业务页复用，减少重复代码。 |
| TypeScript     | 接口、Props、Store 都有类型定义，降低运行时错误。 |
| 权限与安全     | RBAC + v-permission 做按钮级控制；登录态失效统一在响应拦截器处理。 |
| 工程化         | Vite、路径别名、环境变量、开发代理，请求与路由统一封装。 |
| 复杂表单       | 商品多规格 SKU 与主表单联动、校验和提交逻辑清晰。 |

---

## 五、可能的追问与简答

- **表格数据很多时怎么优化？**  
  当前表格通过 `scroll.y` 做纵向滚动，分页加载；若再优化可考虑虚拟滚动或后端游标分页。

- **如何保证请求不重复发送？**  
  列表页通过 ApiBasicTable 内部 loading 和分页/搜索参数控制，同一时刻只有一次 list 请求；必要时可加防抖或请求 cancel。

- **权限数据从哪里来？**  
  登录后调用户信息接口，后端在 user 里带 permission 数组（权限码），前端存到 Pinia，路由和 v-permission 都基于这份数据。

- **有没有做过单元测试？**  
  可按实际情况回答：若没有，可说「关键工具函数和 Store 适合单测，目前以联调与 E2E 为主」；若有，可说用了 Vitest 等。

---

## 六、STAR 示例（讲一个难点或亮点）

**S**：列表页很多，每个都要分页、搜索、Loading、高度自适应，重复代码多且容易不一致。  
**T**：需要一套通用表格方案，让业务页只关心「请求谁、查什么、列怎么展示」。  
**A**：封装了 ApiBasicTable，通过 props 传入 api 函数和 searchParam，内部统一处理分页、合并参数、请求、list/total 更新和 Loading；用 bodyCell 插槽支持自定义列；暴露 useReload 给搜索按钮用。在商品、订单、内容等 10+ 列表页接入，列表逻辑大幅收敛。  
**R**：列表类需求开发更快，行为一致，后续改分页或错误处理只需改一处。

---

## 七、简历一句话（可直接用）

> 负责电商 ERP 管理后台前端开发（Vue 3 + TypeScript + Ant Design Vue），封装 ApiBasicTable、图片裁剪上传、ECharts 图表等通用组件，实现 RBAC 与 Axios 统一拦截，对接 60+ 接口，完成商品、订单、内容、站点、系统管理等 6 大模块共 30+ 页面。

---

*根据面试岗位可适当突出：组件化与工程化（前端岗）、与后端协作与接口设计（全栈岗）、权限与安全（安全/中后台岗）。*
