# 项目架构文档

## 项目概述

这是一个基于 Go 语言开发的电商后端系统，采用分层架构设计，支持管理员端和客户端双端 API。系统集成了商品管理、订单处理、内容管理、支付等功能模块。

## 技术栈

### 核心框架与库
- **Web框架**: Gin v1.10.1
- **ORM**: GORM v1.30.1
- **数据库**: MySQL (驱动: gorm.io/driver/mysql)
- **缓存**: Redis (github.com/go-redis/redis/v8)
- **配置管理**: Viper v1.20.1
- **JWT认证**: github.com/dgrijalva/jwt-go
- **支付集成**: PayPal (github.com/plutov/paypal/v4)
- **邮件服务**: gopkg.in/gomail.v2
- **工具库**: 
  - github.com/google/uuid (UUID生成)
  - github.com/jinzhu/copier (对象拷贝)

### Go 版本
- Go 1.24.4 (工具链: go1.24.7)

## 项目结构

```
server/
├── cmd/                    # 应用程序入口
│   ├── main.go            # 主程序入口
│   ├── oss/               # 对象存储文件目录
│   └── uploads/           # 上传文件目录
├── config/                 # 配置管理
│   ├── config.go          # 配置结构体与初始化
│   └── config.yaml        # 配置文件
├── handlers/               # 处理器层（Controller）
│   ├── admin/             # 管理员端处理器
│   ├── client/            # 客户端处理器
│   └── constant/          # 常量定义
├── middleware/             # 中间件
│   ├── auth.go            # 认证中间件
│   └── cors.go            # 跨域中间件
├── models/                 # 数据模型层
│   ├── cms/               # 内容管理系统模型
│   ├── common/            # 通用模型
│   ├── core/              # 核心系统模型
│   ├── mp/                # 小程序用户模型
│   ├── mypaypal/          # PayPal支付模型
│   ├── shop/              # 商店标签模型
│   └── sp/                # 商城产品模型
├── repository/             # 数据访问层（Repository）
│   ├── base.go            # 基础Repository接口
│   ├── factory.go         # Repository工厂
│   ├── cms/               # CMS相关Repository
│   ├── core/              # 核心系统Repository
│   ├── mp/                # 小程序用户Repository
│   ├── paypal/            # PayPal支付Repository
│   ├── shop/              # 商店标签Repository
│   └── sp/                # 商城产品Repository
├── router/                 # 路由配置
│   └── router.go          # 路由设置
├── service/                # 业务逻辑层（Service）
│   ├── base.go            # 基础Service
│   ├── factory.go         # Service工厂
│   ├── cms/               # CMS相关Service
│   ├── core/              # 核心系统Service
│   ├── mp/                # 小程序用户Service
│   ├── paypal/            # PayPal支付Service
│   ├── shop/              # 商店标签Service
│   └── sp/                # 商城产品Service
├── utils/                  # 工具函数
│   ├── core_auth_session_util.go  # 认证会话工具
│   ├── dict_utils.go      # 字典工具
│   ├── email.go           # 邮件工具
│   ├── jwt.go             # JWT工具
│   ├── order.go           # 订单工具
│   ├── other_type_to_uint.go  # 类型转换工具
│   ├── password_utils.go  # 密码工具
│   └── response.go        # 响应工具
├── sql/                    # SQL脚本
├── go.mod                  # Go模块依赖
└── go.sum                  # 依赖校验和
```

## 架构设计

### 分层架构

系统采用经典的分层架构模式，从上到下分为：

```
┌─────────────────┐
│   Handlers      │  控制器层：处理HTTP请求和响应
│   (Controller)  │
├─────────────────┤
│   Service       │  业务逻辑层：处理业务规则和流程
│   (Business)    │
├─────────────────┤
│   Repository    │  数据访问层：封装数据库操作
│   (Data Access) │
├─────────────────┤
│   Models        │  数据模型层：定义数据结构
│   (Entity)      │
└─────────────────┘
```

### 设计模式

#### 1. 工厂模式（Factory Pattern）
- **RepositoryFactory**: 统一管理所有Repository实例，采用懒加载方式创建
- **ServiceFactory**: 统一管理所有Service实例，依赖RepositoryFactory

#### 2. 依赖注入（Dependency Injection）
- 通过工厂模式实现依赖注入
- Service依赖Repository，Handler依赖Service

#### 3. 中间件模式（Middleware Pattern）
- 认证中间件：`AuthMiddleware`、`OptionalClientAuthMiddleware`
- 跨域中间件：`Cors`
- 设备指纹中间件：`DeviceFingerprintMiddleware`

## 核心模块

### 1. Core 核心系统模块

**功能**: 系统管理、权限控制、组织架构

**主要实体**:
- `CoreAdmin`: 管理员
- `CoreRole`: 角色
- `CorePermission`: 权限
- `CoreDept`: 部门
- `CoreConfig`: 系统配置
- `CoreAdminRoleIndex`: 管理员角色关联

**主要功能**:
- 管理员登录认证（基于Redis Session）
- 角色权限管理（RBAC）
- 部门树形结构管理
- 系统配置管理

### 2. SP 商城产品模块

**功能**: 电商核心业务

**主要实体**:
- `SpProduct`: 商品
- `SpCategory`: 商品分类（树形结构）
- `SpSku`: SKU（库存单位）
- `SpSkuIndex`: SKU索引
- `SpProductContent`: 商品详情内容
- `SpProductProperty`: 商品属性
- `SpProdAttributes`: 商品属性定义
- `SpProdAttributesValue`: 商品属性值
- `SpOrder`: 订单
- `SpOrderItem`: 订单项
- `SpOrderReceiveAddress`: 订单收货地址
- `SpOrderRefund`: 订单退款
- `SpUserCart`: 用户购物车
- `SpUserAddress`: 用户地址

**主要功能**:
- 商品管理（CRUD、上下架、库存管理）
- 分类管理（树形结构）
- SKU管理（多规格商品）
- 订单管理（创建、支付、发货、退款）
- 购物车管理
- 用户地址管理

### 3. CMS 内容管理模块

**功能**: 内容发布与管理

**主要实体**:
- `CmsDocument`: 文档/文章
- `CmsDocumentArchive`: 文档归档
- `CmsRecommend`: 推荐内容
- `CmsRecommendIndex`: 推荐索引
- `CmsCategory`: 内容分类

**主要功能**:
- 文档/文章管理
- 内容推荐管理
- 文档归档

### 4. Shop 商店标签模块

**功能**: 商品标签管理

**主要实体**:
- `ShopTag`: 标签
- `ShopTagIndex`: 标签索引
- `ShopTagMate`: 标签关联

**主要功能**:
- 标签管理
- 标签与商品关联

### 5. MP 小程序用户模块

**功能**: 小程序端用户管理

**主要实体**:
- `MpUser`: 小程序用户
- `MpUserToken`: 用户Token（JWT）

**主要功能**:
- 用户注册/登录（JWT认证）
- 用户信息管理

### 6. PayPal 支付模块

**功能**: PayPal支付集成

**主要实体**:
- `PaypalOrderLogs`: PayPal订单日志
- `PaypalWebhookLogs`: PayPal Webhook日志

**主要功能**:
- PayPal订单创建
- 支付回调处理
- 支付日志记录

## API 路由设计

### 管理员端 API (`/api/manage`)

所有管理员端API需要经过 `AuthMiddleware` 认证中间件。

#### 核心系统
- `/api/manage/core/auth/login` - 管理员登录
- `/api/manage/core/auth/info` - 获取管理员信息
- `/api/manage/core/admin/*` - 管理员管理
- `/api/manage/core/role/*` - 角色管理
- `/api/manage/core/permission/*` - 权限管理
- `/api/manage/core/dept/*` - 部门管理

#### 商品管理
- `/api/manage/shop/product/*` - 商品管理
- `/api/manage/shop/category/*` - 分类管理
- `/api/manage/shop/prodAttributes/*` - 商品属性管理
- `/api/manage/shop/prodAttributesValue/*` - 商品属性值管理
- `/api/manage/shop/tag/*` - 标签管理

#### 订单管理
- `/api/manage/shop/order/*` - 订单管理
- `/api/manage/shop/refund/*` - 退款管理
- `/api/manage/payment/paypal/refund` - PayPal退款

#### 内容管理
- `/api/manage/shop/document/*` - 文档管理
- `/api/manage/shop/recommend/*` - 推荐管理
- `/api/manage/shop/recommendIndex/*` - 推荐索引管理

#### 系统配置
- `/api/manage/shop/marketSetting/*` - 市场设置

#### 文件管理
- `/api/manage/core/oss/*` - OSS文件上传/删除

### 客户端 API (`/api/client`)

客户端API使用 `OptionalClientAuthMiddleware` 和 `DeviceFingerprintMiddleware` 中间件。

#### 商品相关
- `/api/client/shop/product/*` - 商品查询
- `/api/client/shop/category/*` - 分类查询
- `/api/client/shop/tag/*` - 标签查询
- `/api/client/shop/market/*` - 市场信息

#### 用户相关
- `/api/client/shop/userAuth/*` - 用户认证（注册/登录）
- `/api/client/shop/userAddress/*` - 用户地址管理
- `/api/client/shop/userCart/*` - 购物车管理

#### 订单相关
- `/api/client/shop/order/*` - 订单创建/查询

#### 支付相关
- `/api/client/payment/paypal/*` - PayPal支付

#### 内容相关
- `/api/client/shop/document/*` - 文档查询
- `/api/client/shop/recommendIndex/*` - 推荐内容查询

## 认证与授权

### 管理员认证
- **方式**: Redis Session
- **中间件**: `AuthMiddleware`
- **Token来源**: 
  - Authorization Header (支持Bearer前缀)
  - Query参数 `token`
  - Cookie `YEX_AUTH` 或 `auth_token`
- **Session存储**: Redis

### 客户端认证
- **方式**: JWT Token
- **中间件**: `OptionalClientAuthMiddleware` (可选认证)
- **Token来源**: 同管理员认证
- **Token解析**: 使用 `utils.ParseToken()` 解析JWT

### 设备指纹
- **中间件**: `DeviceFingerprintMiddleware`
- **来源**: 
  - Header `X-Device-Fingerprint`
  - Query参数 `device_fingerprint`
  - Cookie `device_fingerprint`

## 数据库设计

### 数据库连接
- **驱动**: MySQL
- **ORM**: GORM
- **连接池配置**:
  - 最大空闲连接: 10
  - 最大打开连接: 100
  - 连接最大存活时间: 1小时

### 软删除
- 使用 GORM 的 `gorm.DeletedAt` 实现软删除
- 部分表使用 `DeletedTime *time.Time` 自定义软删除字段

### 通用字段
- `ID`: 主键（`common.MyID` 类型，uint32）
- `CreatedTime`: 创建时间
- `UpdatedTime`: 更新时间
- `DeletedTime`: 删除时间（软删除）

## 配置管理

### 配置文件结构 (`config/config.yaml`)

```yaml
server:
  port: 8080        # 服务器端口
  mode: debug       # 运行模式

database:
  driver: mysql
  host: "172.25.13.23"
  port: 3306
  username: "root"
  password: "123"
  dbname: "zg_shop"
  charset: "utf8mb4"

redis:
  host: "172.25.13.23"
  port: 6379
  password: "123"
  db: 0

email:
  host: "smtp.qq.com"
  port: 587
  username: "xxx@qq.com"
  password: "授权码"
  from: "xxx@qq.com"

payment:
  wechat:          # 微信支付配置
    app_id: "..."
    mch_id: "..."
    api_key: "..."
    notify_url: "..."
  alipay:          # 支付宝配置
    app_id: "..."
    private_key: "..."
    public_key: "..."
    notify_url: "..."
  paypal:          # PayPal配置
    paypal_client_id: "..."
    paypal_secret: "..."
    paypal_api_base: "https://api.sandbox.paypal.com"

frontend:
  url: "http://localhost:3000/order-detail/"
```

## 工具函数

### 认证相关
- `utils.GetSession()`: 从Redis获取Session
- `utils.ParseToken()`: 解析JWT Token
- `utils.GenerateToken()`: 生成JWT Token

### 密码相关
- `utils.HashPassword()`: 密码加密
- `utils.VerifyPassword()`: 密码验证

### 响应相关
- `utils.SuccessResponse()`: 成功响应
- `utils.ErrorResponse()`: 错误响应

### 其他工具
- `utils.SendEmail()`: 发送邮件
- `utils.OtherTypeToUint()`: 类型转换

## 启动流程

1. **初始化配置**: 读取 `config.yaml` 配置文件
2. **初始化数据库**: 连接MySQL数据库，配置连接池
3. **初始化Redis**: 连接Redis服务器
4. **创建工厂**:
   - 创建 `RepositoryFactory`
   - 创建 `ServiceFactory`
5. **设置路由**: 配置所有API路由
6. **启动服务器**: 监听配置的端口

## 开发规范

### 命名规范
- **包名**: 小写，单数形式
- **结构体**: 大驼峰命名
- **方法**: 大驼峰命名
- **变量**: 小驼峰命名
- **常量**: 大写下划线分隔

### 文件组织
- 每个模块按功能划分到对应目录
- Repository、Service、Handler 一一对应
- 模型文件放在 `models` 目录下对应子目录

### 错误处理
- 使用统一的响应格式
- 错误信息通过 `utils.ErrorResponse()` 返回

### 日志
- 使用标准库 `log` 包
- 关键操作记录日志

## 扩展点

### 添加新模块步骤

1. **创建模型** (`models/xxx/`)
   - 定义数据结构
   - 实现 `TableName()` 方法

2. **创建Repository** (`repository/xxx/`)
   - 实现数据访问逻辑
   - 在 `RepositoryFactory` 中添加获取方法

3. **创建Service** (`service/xxx/`)
   - 实现业务逻辑
   - 在 `ServiceFactory` 中添加获取方法

4. **创建Handler** (`handlers/admin/` 或 `handlers/client/`)
   - 处理HTTP请求
   - 调用Service层方法

5. **配置路由** (`router/router.go`)
   - 添加路由规则
   - 配置中间件

## 注意事项

1. **安全性**
   - 生产环境需要修改默认配置
   - 密码使用加密存储
   - API需要认证保护

2. **性能**
   - 数据库连接池已配置
   - Redis用于缓存和Session存储
   - 考虑添加查询缓存

3. **可维护性**
   - 遵循分层架构
   - 使用工厂模式管理依赖
   - 保持代码结构清晰

4. **扩展性**
   - 模块化设计便于扩展
   - 工厂模式便于添加新功能
   - 中间件机制便于添加横切关注点

## 待优化项

1. **代码注释**: 部分代码缺少详细注释
2. **单元测试**: 缺少单元测试覆盖
3. **API文档**: 建议使用Swagger生成API文档
4. **错误处理**: 可以统一错误码定义
5. **日志系统**: 建议使用结构化日志库（如zap）
6. **配置验证**: 启动时验证配置完整性
7. **数据库迁移**: 使用GORM Migrate管理数据库版本

## 总结

本项目采用清晰的分层架构，使用工厂模式管理依赖，支持管理员端和客户端双端API。系统模块化程度高，便于维护和扩展。主要业务模块包括核心系统、商城产品、内容管理、支付等，基本覆盖了电商系统的核心功能。
