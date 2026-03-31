# TiDB Dashboard 项目记忆

## 项目概述
- TiDB Dashboard 是 TiDB 集群的 Web UI 监控/诊断/管理工具
- 后端 Go (pkg/)，前端 React (ui/ v1, ui-v2/ v2)
- 使用 uber/fx 依赖注入，gin HTTP 框架，Pullstate 状态管理
- 分支 `feature/easygraph` 上进行了 NoTiDB 模式开发

## 开发环境
- Go 1.25.7, Node 25.x, pnpm 8.6.12
- 本地集群: `tiup playground --mode tikv-slim --kv 3` (PD: 2379, TiKV: 20160-20162)
- 后端启动: `bin/tidb-dashboard --debug --experimental --no-tidb --pd http://127.0.0.1:2379`
- 前端开发: `cd ui && pnpm dev:op` (port 3001)

## 关键架构
- 认证: JWT + 三种认证器 (sqlauth/ssoauth/codeauth)，NoTiDB 模式用 noauth
- 前端 v1: single-spa 微前端 + hash 路由 + Ant Design
- 前端 v2 (孵化中): TanStack Router + @tidbcloud/uikit，仅 3 个模块
- distro 机制: 可定制组件名称，默认硬编码

## NoTiDB 模式 (feature/easygraph 分支)
- `--no-tidb` 启用纯 TiKV 模式
- noauth 替代 sqlauth 认证
- 后端排除: statement, slowquery, queryeditor, diagnose, topsql, deadlock, resourcemanager
- 前端隐藏: Top SQL, Statement, Slow Query, System Report, Resource Manager, Query Editor
- 保留: Overview, Cluster Info, Key Viz, Monitoring, Search Logs, Profiling, Debug API, Configuration

最后更新: 2026-03-28
