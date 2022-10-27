# topology-server

部署在 k8s 环境中，自动读取当前 pod 所在地域与可用区信息，通过接口返回。

## 部署

```bash
kubectl apply -f topology-server.yaml
```

**注意：** 如果部署到非 default 命名空间，确保修改 yaml 中的 `ClusterRoleBinding`，修改引用的 ServiceAccount 所在 namespace。