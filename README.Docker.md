# Docker 部署指南

## 快速开始

### 1. 启动所有服务

```bash
docker-compose up -d --build
```

### 2. 查看服务状态

```bash
docker-compose ps
```

### 3. 查看日志

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f backend
docker-compose logs -f mysql
```

## 访问地址

- **后端 API**: http://localhost:8080
- **管理后台**: http://localhost:3001
- **PC 商城**: http://localhost:3000
- **MySQL**: localhost:3306 (root/root123456)
- **Redis**: localhost:6379

## 常用命令

### 启动服务
```bash
docker-compose up -d
```

### 停止服务
```bash
docker-compose down
```

### 停止并删除数据（⚠️ 会删除数据库数据）
```bash
docker-compose down -v
```

### 重启服务
```bash
docker-compose restart backend
```

### 重新构建
```bash
docker-compose build
docker-compose up -d
```

## 配置说明

### 数据库配置

配置文件 `server/config/config.docker.yaml` 已配置好 Docker 环境：
- 数据库主机：`mysql`（Docker 服务名）
- Redis 主机：`redis`（Docker 服务名）

### 数据库初始化

将 SQL 文件放在 `server/sql/` 目录下，Docker 会自动执行初始化。

## 故障排查

### 后端无法连接数据库

1. 检查 MySQL 服务是否健康：`docker-compose ps`
2. 检查后端日志：`docker-compose logs backend`
3. 确认配置中的 host 使用服务名 `mysql` 而不是 `localhost`

### 前端无法访问后端 API

1. 检查后端服务是否正常运行
2. 确认前端配置中的 API 地址正确
3. 检查网络连接

### 数据库初始化失败

1. 检查 SQL 文件是否在 `server/sql/` 目录
2. 查看 MySQL 日志：`docker-compose logs mysql`
3. 手动导入 SQL 文件
