# Step 3: 屏蔽需要 TiDB 的后端 API 模块

## 目标
在 NoTiDB 模式下，不注册依赖 TiDB SQL 连接的后端 API 模块。

## 实施方案

### `NoTiDBModules` 中排除的模块

以下模块因依赖 TiDB SQL 连接（通过 `MWConnectTiDB` 中间件或直接使用 `tidb.Client.OpenSQLConn`）被从 NoTiDB 模式中排除：

| 模块 | 排除原因 |
|------|---------|
| `statement` | 查询 `INFORMATION_SCHEMA.CLUSTER_STATEMENTS_SUMMARY` |
| `slowquery` | 查询 `INFORMATION_SCHEMA.CLUSTER_SLOW_QUERY` |
| `queryeditor` | 直接执行 SQL 语句 |
| `diagnose` | 生成诊断报告需要 SQL 连接 |
| `topsql` | 依赖 TiDB 状态 API |
| `deadlock` | 查询 `INFORMATION_SCHEMA.CLUSTER_DEADLOCKS` |
| `resourcemanager` | 管理 TiDB 资源组 |

### `NoTiDBModules` 中保留的模块

| 模块 | 保留原因 |
|------|---------|
| `info` | 基本信息（但 /databases、/tables 路由被跳过）|
| `clusterinfo` | 集群拓扑信息（PD/TiKV/TiFlash）|
| `keyvisual` | Key Visualizer（依赖 PD region 数据）|
| `metrics` | Prometheus 监控（不依赖 TiDB）|
| `logsearch` | 日志搜索（PD/TiKV/TiFlash）|
| `profiling` | 性能分析（PD/TiKV/TiFlash）|
| `conprof` | 持续分析（依赖 ngm）|
| `debugapi` | 调试 API（支持 PD/TiKV 等）|
| `configuration` | 配置管理 |

### 不注册的路由

在 NoTiDB 模式的 `fx.Invoke` 中排除了以下路由注册：
- `diagnose.RegisterRouter`
- `queryeditor.RegisterRouter`

## 完成情况
- ✅ 实施完成（已集成在 Step 2 的 apiserver.go 修改中）
