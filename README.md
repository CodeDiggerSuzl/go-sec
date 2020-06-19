# Go 秒杀系统
知识储备
- RabbitMQ 入门
- Iris 入门

基本功能
- 后端商品管理功能开发
- 后端订单管理功能开发
- 前台用户鞥路功能
- 商品详情功能的开发

性能调优
- 架构调优
- 前端优化
- 服务端优化
- 安全优化

业务流程:

![](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfwwo4ldzzj31gm0u0qdr.jpg)


## 2. 需求管理和系统设计
### 2.1 需求分析

主要的功能点:
1. 前台用户登录,商品展示,商品抢购
2. 后台订单管理

### 2.2 原型设计

主要设计页面,使用磨刀原型设计工具
- 前台用户登录页面,商品展示页面,抢购结果页面
- 订单管理页面

### 2.3 系统架构设计

#### 系统需求分析

- 前端页面需要承载大流量
- 在大并发的状态下解决超卖的问题
- 后端接口需要满足横向扩展

#### 系统架构设计


## 3. 环境搭建
- MySQL
- go
- RabbitMQ

### 3.1 RabbitMQ 介绍
定义和特征:

1. RabbitMQ 是面向消息的中间件,用于组件之间的解耦,主要体现在消息的发送和消费者之间的无依赖关系.
2. RabbitMQ 特点: 高可用,扩展性,多语言客户端,管理界面.
3. 主要的使用场景: 流量削峰,异步处理,应用解耦.

#### 常用命令
常用管理命令

- 启动/停止: `systemctl start/stop rabbitmq-server`
- 插件管理命令:
    1. 查看插件: `rabbitmq-plugins list`
    2. 启动禁用插件:`rabbit-plugins enable/disable plugin-name`
    
#### 管理界面的使用





#### rabbitMQ 基础知识
- 登录用户名和账号都是 guest
- 端口: 15672

### 3.2 核心概念

- VirtualHost 数据隔离的作用
- Connection
- Exchange 交换机
- Channel 
- Queue 队列 绑定交换机 或者临时存储消息
- Binding

#### 3.3 快速入门
##### 6 种工作模式
1. Simple 模式
    最简单的工作模式 p->[==]->c
---

- [ ] CDN
- [ ] 静态回源
- [ ] SLB alibaba

