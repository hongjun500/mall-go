# 声明
**这个项目是按照 [大佬 macrozheng ](https://github.com/macrozheng) 开源项目 [mall](https://github.com/macrozheng/mall) 使用 Go 来实现的**

任何事物 **从零到一** 才是最难能可贵的，如果你某天有幸看到这个项目并得到了一些帮助，请给 原作者 [macrozheng](https://github.com/macrozheng) 点个 star 吧

## 项目布局

```
mall-go
├───cmd                   // 项目启动入口
│   ├───admin
│   ├───portal
│   └───search
├───configs               // 配置文件
├───docs                  // swag 文档
│   ├───mall-search
│   ├───mall_admin
│   └───mall_search
├───internal             // 内部模块
│   ├───conf             // 配置项
│   ├───database         // 数据库连接
│   ├───es_index         // es 索引
│   ├───gin_common       // gin 相关 包含通用返回,错误处理
│   │   ├───mid          // gin 中间件 
│   │   └───security     // gin 安全相关, jwt, cors, casbin
│   ├───initialize       // 初始化
│   ├───models           // 数据库模型
│   ├───request          // 请求参数
│   │   ├───base_dto
│   │   ├───ums_admin_dto
│   │   └───ums_member_dto
│   ├───routers            // 路由
│   │   ├───r_mall_admin   // admin 服务路由
│   │   └───r_mall_search  // search 服务路由
│   └───services           // 服务
│       ├───s_mall_admin   // admin 服务
│       └───s_mall_search  // search 服务
├───pkg                // 公共模块
│   ├───constants      // 常量
│   ├───convert        // 类型转换 json
│   ├───redis          // redis 操作
│   └───security       // jwt,casbin 相关
├───scripts
│   ├───swag           // swag 文档生成
│   └───sql-script     // sql 脚本
│       └───insert
└───tests              // 测试
    ├───common
    ├───database
    ├───models
    └───services
```
