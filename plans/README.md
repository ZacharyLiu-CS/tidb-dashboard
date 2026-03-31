# TiDB Dashboard NoTiDB 模式 — 总体规划与实施记录

## 项目背景
TiDB Dashboard 原本要求集群中必须有 TiDB 实例才能正常工作。在 `tiup playground --mode tikv-slim --kv 3` 模式下启动的集群只有 TiKV，没有 TiDB 实例，导致 Dashboard 无法登录。

## 目标
使 TiDB Dashboard 能在没有 TiDB 实例的纯 TiKV 集群中正常运行，同时屏蔽需要 TiDB SQL 连接的功能。

## 分支
所有修改在 `feature/easygraph` 分支上完成。

## 修改文件清单

### 后端 (Go)
| 文件 | 修改类型 | 说明 |
|------|---------|------|
| `pkg/config/config.go` | 修改 | 添加 `NoTiDB` 配置字段 |
| `cmd/tidb-dashboard/main.go` | 修改 | 添加 `--no-tidb` CLI 参数 |
| `pkg/apiserver/user/noauth/noauth.go` | **新建** | NoAuth 认证器，无需 TiDB SQL 验证 |
| `pkg/apiserver/apiserver.go` | 修改 | 添加 `NoTiDBModules`，动态选择模块和路由 |
| `pkg/apiserver/utils/tidb_conn.go` | 修改 | NoTiDB 模式下跳过 TiDB 连接（不阻断请求）|
| `pkg/apiserver/info/info.go` | 修改 | 添加 `NoTiDB` 响应字段；条件注册 TiDB 路由 |

### 前端 (TypeScript)
| 文件 | 修改类型 | 说明 |
|------|---------|------|
| `.../tidb-dashboard-client/.../info-info-response.ts` | 修改 | 添加 `no_tidb` 字段 |
| `.../tidb-dashboard-lib/src/client/models.ts` | 修改 | 添加 `no_tidb` 字段 |
| `.../tidb-dashboard-lib/src/utils/store.ts` | 修改 | 添加 `useIsNoTiDB()` hook |
| `.../tidb-dashboard-for-op/.../Sider/index.tsx` | 修改 | 根据 NoTiDB 状态隐藏菜单项 |

## 使用方式
```bash
# 编译
go build -o bin/tidb-dashboard cmd/tidb-dashboard/main.go

# 以 no-tidb 模式启动
bin/tidb-dashboard --debug --experimental --no-tidb --pd http://127.0.0.1:2379

# 前端开发
cd ui && pnpm dev:op
```

## 详细步骤文档
- [Step 1: 构建后端并连接集群](step1-build-and-connect.md)
- [Step 2: 移除 TiDB 实例强制要求](step2-remove-tidb-requirement.md)
- [Step 3: 屏蔽 TiDB 后端模块](step3-disable-tidb-backend-modules.md)
- [Step 4: 屏蔽 TiDB 前端模块](step4-disable-tidb-frontend-modules.md)
- [Step 5: 验证](step5-verification.md)
