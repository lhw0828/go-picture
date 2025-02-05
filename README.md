# Go-Picture 图片管理系统

## 项目简介
这是一个基于 Go 语言开发的分布式图片管理系统，采用微服务架构设计，使用 go-zero 框架实现。系统支持多空间管理、团队协作、权限控制等功能。

## 系统架构

### 已实现服务
1. 用户服务 (User Service)
   - [x] 用户注册登录
   - [x] JWT 认证授权
   - [x] 用户信息管理
   - [x] 角色权限控制

2. 空间服务 (Space Service)
   - [x] 空间创建和管理
   - [x] 空间成员权限控制
   - [x] 空间容量管理
   - [x] 空间使用统计
   - [x] 团队协作支持

### 技术栈
- 框架：go-zero (微服务框架)
- 通信：gRPC (服务间通信)
- API：REST API (对外接口)
- 认证：JWT (用户认证)
- 数据库：MySQL
- 缓存：Redis
- 日志：Zap
- 监控：Prometheus

## 项目特性

### 错误处理机制
- 统一错误码管理
- 错误信息国际化
- 自定义错误类型
- 错误链路追踪

### 日志系统
- 分级日志记录
- 链路追踪集成
- 自动日志轮转
- 日志压缩存储

### 缓存设计
- Redis缓存集成
- 序列化处理
- 过期策略
- 批量操作支持

### 中间件
- 认证中间件 (JWT验证)
- 错误处理中间件
- 日志中间件
- 监控中间件

## 项目结构

├── api                     # API服务层
│   ├── picture-api        # 图片服务API
│   │   ├── etc           # 配置文件
│   │   └── internal      # 内部实现
│   │       ├── config    # 配置结构
│   │       ├── handler   # 请求处理器
│   │       ├── logic     # 业务逻辑
│   │       ├── svc       # 服务上下文
│   │       └── types     # 类型定义
│   ├── space-api         # 空间服务API
│   │   ├── etc          # 配置文件
│   │   └── internal     # 内部实现
│   │       ├── config   # 配置结构
│   │       ├── handler  # 请求处理器
│   │       ├── logic    # 业务逻辑
│   │       ├── svc      # 服务上下文
│   │       └── types    # 类型定义
│   └── user-api          # 用户服务API
│       ├── etc          # 配置文件
│       └── internal     # 内部实现
│           ├── config   # 配置结构
│           ├── handler  # 请求处理器
│           ├── logic    # 业务逻辑
│           ├── svc      # 服务上下文
│           └── types    # 类型定义
├── common                 # 公共模块
│   ├── constants         # 常量定义
│   ├── errorx           # 错误处理
│   ├── middleware       # 中间件
│   ├── response         # 响应处理
│   ├── sql             # SQL脚本
│   ├── types           # 公共类型
│   └── utils           # 工具函数
├── configs               # 配置文件目录
│   └── dev             # 开发环境配置
├── pkg                  # 基础组件包
│   ├── cache           # 缓存组件
│   │   └── redis.go   # Redis实现
│   ├── logger          # 日志组件
│   │   └── logger.go  # 日志实现
│   ├── metrics         # 监控指标
│   └── trace          # 链路追踪
└── rpc                  # RPC服务层
├── picture-rpc      # 图片服务RPC
│   ├── etc         # 配置文件
│   └── internal    # 内部实现
│       ├── config  # 配置结构
│       ├── dao     # 数据访问
│       ├── logic   # 业务逻辑
│       ├── server  # 服务实现
│       └── svc     # 服务上下文
├── space-rpc        # 空间服务RPC
│   ├── etc         # 配置文件
│   └── internal    # 内部实现
│       ├── config  # 配置结构
│       ├── dao     # 数据访问
│       ├── logic   # 业务逻辑
│       ├── server  # 服务实现
│       └── svc     # 服务上下文
└── user-rpc         # 用户服务RPC
├── etc         # 配置文件
└── internal    # 内部实现
├── config  # 配置结构
├── dao     # 数据访问
├── logic   # 业务逻辑
├── server  # 服务实现
└── svc     # 服务上下文


## 服务通信
1. API层通过gRPC客户端调用RPC服务
2. 服务间通过gRPC进行通信
3. 对外提供REST API接口

## 错误处理
系统实现了统一的错误处理机制：
- 定义了统一的错误码
- 实现了错误处理中间件
- 提供了基础错误类型
