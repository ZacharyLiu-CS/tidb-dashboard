# Step 1: 构建后端并连接本地 tikv-slim 集群

## 目标
在本地编译 tidb-dashboard 后端，以调试模式启动并连接到本地 tikv-slim 集群（PD: 127.0.0.1:2379）。

## 集群状态
- PD: 127.0.0.1:2379 (v8.5.5)
- TiKV: 3 个实例 (127.0.0.1:20160, 20161, 20162)
- TiDB: **无实例** (tikv-slim 模式)

## 操作步骤
1. ✅ 确认在 `feature/easygraph` 分支
2. ✅ 验证集群正常运行（PD 和 TiKV 均在线）
3. ✅ 编译后端: `go build -o bin/tidb-dashboard cmd/tidb-dashboard/main.go`
4. ✅ 创建 `bin/distro-res/strings.json`（空 JSON）
5. 以调试模式启动: `bin/tidb-dashboard --debug --experimental --pd http://127.0.0.1:2379`

## 预期问题
- Dashboard 启动后，登录页面会因为没有 TiDB 实例而无法认证（SQL Auth 需要连接 TiDB）
- 这将在 Step 2 中通过绕过认证来解决

## 完成情况
- ✅ 后端编译成功
- ⏳ 后端启动和登录验证（将在后续步骤中配合代码修改一起完成）
