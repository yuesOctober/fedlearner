# Fedlearner Operator

## Prerequisite

Kubernetes 集群中的节点必须安装 `nfs-common` 依赖，例如：

```sh
apt install nfs-common
```

## Installation

### 安装 Fedlearner 基础设施

```sh
helm install fedlearner-stack ./deploy/charts/fedlearner-stack
```

### 安装 Fedlearner Controller 和 CRD

```sh
helm template ./deploy/charts/fedlearner --namespace leader | kubectl apply -f
helm template ./deploy/charts/fedlearner --namespace follower | kubectl apply -f
kubectl apply -f ./deploy/charts/manifests/
```
