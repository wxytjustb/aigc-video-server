# AIGC Video Server

基于 Gin + Vue + Element Plus 的 AIGC 视频服务器全栈开发基础平台。

## 项目简介

本项目是一个现代化的全栈 Web 应用，后端采用 Go 语言的 Gin 框架，前端使用 Vue3 + Element Plus，提供完整的 RBAC 权限管理、代码生成器、表单生成器等功能。

## 技术栈

### 后端技术栈
- **核心框架**: Gin (Go Web Framework)
- **ORM**: GORM
- **数据库支持**: MySQL, PostgreSQL, SQLServer, SQLite, Oracle
- **缓存**: Redis
- **认证**: JWT
- **权限**: Casbin (RBAC)
- **文档**: Swagger
- **日志**: Zap
- **配置**: Viper
- **对象存储**: 支持本地存储、阿里云OSS、腾讯云COS、七牛云、MinIO、AWS S3、华为云OBS等
- **其他**: MongoDB、定时任务(Cron)、文件上传下载、Excel导入导出

### 前端技术栈
- **核心框架**: Vue 3.5+
- **构建工具**: Vite 6+
- **UI组件库**: Element Plus 2.10+
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **HTTP客户端**: Axios
- **富文本编辑器**: WangEditor
- **图表**: ECharts
- **代码编辑器**: Ace Editor
- **CSS框架**: UnoCSS
- **其他**: Vue Draggable、Vue Cropper、Marked等

## 项目结构

```
.
├── deploy/              # 部署目录
│   ├── docker/         # Docker部署配置
│   │   ├── Dockerfile
│   │   └── entrypoint.sh
│   ├── docker-compose/ # Docker Compose配置
│   │   └── docker-compose.yaml
│   └── kubernetes/     # Kubernetes部署配置
│       ├── server/     # 后端K8s配置
│       └── web/        # 前端K8s配置
│
├── dev/                # 开发环境
│   └── docker-compose.yaml  # 本地开发环境配置
│
├── server/             # 后端代码
│   ├── api/           # API接口层
│   ├── config/        # 配置定义
│   ├── core/          # 核心组件
│   ├── docs/          # Swagger文档
│   ├── global/        # 全局变量
│   ├── initialize/    # 初始化模块
│   ├── middleware/    # 中间件
│   ├── model/         # 数据模型
│   ├── plugin/        # 插件系统
│   ├── resource/      # 静态资源
│   ├── router/        # 路由定义
│   ├── service/       # 业务逻辑层
│   ├── source/        # 数据源
│   ├── task/          # 定时任务
│   ├── utils/         # 工具函数
│   ├── main.go        # 程序入口
│   ├── config.yaml    # 配置文件
│   ├── go.mod         # Go依赖管理
│   └── go.sum
│
├── tools/             # 工具页面
│
└── web/               # 前端代码
    ├── src/
    │   ├── api/       # API接口定义
    │   ├── components/# Vue组件
    │   ├── core/      # 核心配置
    │   ├── directive/ # 自定义指令
    │   ├── hooks/     # 组合式函数
    │   ├── pinia/     # 状态管理
    │   ├── plugin/    # 插件
    │   ├── router/    # 路由配置
    │   ├── style/     # 样式文件
    │   ├── utils/     # 工具函数
    │   ├── view/      # 页面视图
    │   ├── App.vue    # 根组件
    │   └── main.js    # 入口文件
    ├── vitePlugin/    # Vite插件
    ├── package.json   # npm依赖管理
    ├── vite.config.js # Vite配置
    └── index.html     # HTML模板
```

## 快速开始

### 环境要求

- **后端**
  - Go 1.24.0+
  - MySQL 5.7+ / PostgreSQL / SQLServer / SQLite
  - Redis (可选)

- **前端**
  - Node.js 16+
  - npm/yarn/pnpm

### 后端启动

1. 进入后端目录
```bash
cd server
```

2. 安装依赖
```bash
go mod download
```

3. 配置数据库
编辑 `config.yaml` 文件，配置数据库连接信息

4. 启动服务
```bash
go run main.go
```

后端服务默认运行在 `http://localhost:8888`

### 前端启动

1. 进入前端目录
```bash
cd web
```

2. 安装依赖
```bash
npm install
```

3. 启动开发服务器
```bash
npm run dev
```

前端服务默认运行在 `http://localhost:5173`

4. 构建生产版本
```bash
npm run build
```

## 部署方式

### Docker 部署

使用 Docker Compose 快速部署：

```bash
cd deploy/docker-compose
docker-compose up -d
```

### Kubernetes 部署

1. 部署后端服务：
```bash
kubectl apply -f deploy/kubernetes/server/
```

2. 部署前端服务：
```bash
kubectl apply -f deploy/kubernetes/web/
```

### 本地开发环境

使用开发环境配置：
```bash
cd dev
docker-compose up -d
```

## 核心功能

- ✅ 用户管理：用户的增删改查、角色分配
- ✅ 角色管理：角色的增删改查、权限分配
- ✅ 菜单管理：菜单的增删改查、权限配置
- ✅ API管理：API的增删改查、权限配置
- ✅ 权限管理：基于Casbin的RBAC权限控制
- ✅ 代码生成器：一键生成CRUD代码
- ✅ 表单生成器：可视化表单设计
- ✅ 文件上传：支持多种对象存储
- ✅ 断点续传：大文件上传支持
- ✅ Excel导入导出：数据导入导出功能
- ✅ 操作日志：系统操作日志记录
- ✅ 定时任务：基于Cron的任务调度
- ✅ 插件系统：灵活的插件扩展机制
- ✅ MongoDB支持：NoSQL数据库集成

## API 文档

启动后端服务后，访问 Swagger 文档：
```
http://localhost:8888/swagger/index.html
```

## 配置说明

### 后端配置

主要配置文件：`server/config.yaml`

- 系统配置：端口、环境、日志级别等
- 数据库配置：主库及多数据库配置
- Redis配置：缓存配置
- JWT配置：Token配置
- 对象存储配置：各类OSS配置
- 邮件配置：邮件发送配置

### 前端配置

- 开发环境：`.env.development`
- 生产环境：`.env.production`
- Vite配置：`vite.config.js`

## 常见问题

### 后端相关

**Q: 如何切换数据库？**  
A: 在 `config.yaml` 中修改 `system.db-type` 字段，支持 mysql/pgsql/sqlite/sqlserver/oracle

**Q: 如何配置对象存储？**  
A: 在 `config.yaml` 中配置对应的OSS配置项，并设置 `local.type` 字段

### 前端相关

**Q: 如何修改API地址？**  
A: 在对应环境的 `.env` 文件中修改 `VITE_BASE_API` 变量

**Q: 如何打包生产版本？**  
A: 运行 `npm run build`，生成的文件在 `dist` 目录

## 开发指南

详细的开发文档请参考：
- 后端开发文档：`server/README.md`
- 前端开发文档：`web/README.md`

## 版本信息

- 后端版本：基于 Go 1.24.0
- 前端版本：v2.8.8

## 许可证

本项目基于开源许可证发布，具体许可证信息请查看项目根目录的 LICENSE 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 联系方式

如有问题或建议，请通过 Issue 与我们联系。
