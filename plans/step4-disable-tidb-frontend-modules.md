# Step 4: 屏蔽前端中依赖 TiDB 的功能模块

## 目标
在 NoTiDB 模式下，前端侧边栏隐藏与 TiDB 相关的功能入口。

## 实施方案

### 后端 API 修改
1. **`pkg/apiserver/info/info.go`** — `InfoResponse` 添加 `NoTiDB bool` 字段
2. 通过 `/info/info` API 返回 `no_tidb: true` 给前端

### 前端类型修改
1. **`ui/packages/tidb-dashboard-client/src/client/api/models/info-info-response.ts`** — 添加 `no_tidb?: boolean`
2. **`ui/packages/tidb-dashboard-lib/src/client/models.ts`** — 同上

### 前端 Store 修改
1. **`ui/packages/tidb-dashboard-lib/src/utils/store.ts`** — 添加 `useIsNoTiDB()` hook

### 前端侧边栏修改
1. **`ui/packages/tidb-dashboard-for-op/src/dashboardApp/layout/main/Sider/index.tsx`**
   - 导入 `useIsNoTiDB`
   - 在 NoTiDB 模式下隐藏以下菜单项：
     - Top SQL
     - Statement
     - Slow Query  
     - System Report (Diagnose)
     - Resource Manager
     - Query Editor (实验性功能)

### NoTiDB 模式下保留的功能
| 功能 | 说明 |
|------|------|
| Overview | 集群概览（TiKV/PD 实例信息）|
| Cluster Info | 集群拓扑信息 |
| Key Visualizer | 数据热点可视化 |
| Monitoring | Prometheus 监控图表 |
| Search Logs | PD/TiKV 日志搜索 |
| Instance Profiling | 性能分析 |
| Continuous Profiling | 持续分析 |
| Debug API | 调试 API |
| Configuration | 配置管理 |
| User Profile | 用户设置 |

## 完成情况
- ✅ 后端 API 添加 `no_tidb` 字段
- ✅ 前端类型定义更新
- ✅ Store hook 添加
- ✅ 侧边栏菜单条件隐藏
