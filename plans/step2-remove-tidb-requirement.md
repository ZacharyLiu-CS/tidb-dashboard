# Step 2: 移除 TiDB 实例的强制要求

## 目标
实现 `--no-tidb` 模式，使 tidb-dashboard 可以在没有 TiDB 实例的集群上运行。

## 核心问题
TiDB Dashboard 的认证依赖于 TiDB SQL 连接（通过 MySQL 协议连接 TiDB 验证用户名密码和权限）。在没有 TiDB 实例时，SQL Auth 完全无法工作。

## 解决方案

### 修改的文件

1. **`pkg/config/config.go`**
   - 在 `Config` 结构体中添加 `NoTiDB bool` 字段

2. **`cmd/tidb-dashboard/main.go`**
   - 添加 `--no-tidb` CLI 参数

3. **`pkg/apiserver/user/noauth/noauth.go`** (新建)
   - 创建 NoAuth 认证器，auth type=0（替代 sqlauth）
   - 不需要 TiDB SQL 连接，直接接受登录
   - 设置 `HasTiDBAuth=false` 以跳过 TiDB 连接中间件

4. **`pkg/apiserver/apiserver.go`**
   - 添加 `NoTiDBModules` 模块集合
   - 修改 `Start()` 方法根据 `NoTiDB` 配置动态选择模块和路由
   - NoTiDB 模式下排除: diagnose, queryeditor, statement, slowquery, deadlock, resourcemanager, topsql

5. **`pkg/apiserver/utils/tidb_conn.go`**
   - 修改 `MWConnectTiDB` 中间件：当 `HasTiDBAuth=false` 时，不拒绝请求，而是设置 nil 连接并继续

6. **`pkg/apiserver/info/info.go`**
   - 在 NoTiDB 模式下不注册 `/databases` 和 `/tables` 路由

## 完成情况
- ✅ 编译成功
- ✅ API 启动正常 (info/info 正确返回)
- ✅ 登录认证成功 (root 用户获得 JWT token)
- ✅ NoTiDB 模式下排除了依赖 TiDB 的 API 模块
