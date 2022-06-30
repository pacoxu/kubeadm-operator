# Kubeadm operator

The kubeadm-operator is an experimental project, still WIP.
Do not use in production.

See [KEP](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/2505-Kubeadm-operator) for more details.

## Quick Start

Configure kubeconfig for your cluster.

```
git clone git@github.com:pacoxu/kubeadm-operator.git
cd kubeadm-operator
make install
make deploy
```
**Notice: for pod cannot restart kubelet, [kubelet-reloader](https://github.com/pacoxu/kubelet-reloader)should run on all cluster nodes as a daemon or process.
**
### kubelet reloader installation
```
wget https://github.com/pacoxu/kubeadm-operator/releases/download/v0.1.0/kubelet-reloader-v0.2.0
chmod +x kubelet-reloader-v0.2.0
./kubelet-reloader-v0.2.0
```
https://github.com/pacoxu/kubeadm-operator/issues/85 is for this problem.

## Version Support Matrix

- ✅✅  means supported and suggested
- ✅❌ means supported but not suggested
- ❌❌ means not supported and not suggested
- Empty means no testing yet.

| initial version\ target version | v1.20 | v1.21 | v1.22 | v1.23 | v1.24 |
|---------------------------------|-------|-------|-------|-------|
| v1.20                           |       |       |       |       |       |
| v1.21                           |       |       |       |       |       |
| v1.22                           |       |       | ✅✅    | ✅✅    | ✅✅    |
| v1.23                           |       | ❌❌      | ✅❌    | ✅✅    | ✅✅    |
| v1.24                           |       | ❌❌      | ❌❌    | ✅❌    | ✅✅    |


## Dev Tips

To build your own kubeadm-operator image:
```
make docker-build
```

You can change the `IMG` in `Makefile` to push the image to your repository and install/deploy to the development cluster.
```
make release
make install
make deploy
```


## Demo 

After installation, a deploy named `operator-controller-manager` is running in namespace `operator-system`.
```
[root@daocloud ~]# kubectl get pod -n operator-system -o wide
NAME                                          READY   STATUS    RESTARTS   AGE   IP               NODE       NOMINATED NODE   READINESS GATES
operator-controller-manager-64c448f5b-p682x   2/2     Running   0          77m   172.32.230.197   daocloud   <none>           <none>
```

If you create a dry-run upgrade operation, there will be a runtimetaqskgroup with 

```
[root@paco ~]# cat up.yaml
apiVersion: operator.kubeadm.x-k8s.io/v1alpha1
kind: Operation
metadata:
  name: upgrade-v1.24.1
spec:
  executionMode: Auto
  upgrade:
    kubernetesVersion: v1.24.1
    local: false
```

After the operation is done, the operation and task group are all `Succeeded`.

```
NAME                                                  PHASE       GROUPS   SUCCEEDED   FAILED
operation.operator.kubeadm.x-k8s.io/upgrade-v1.24.1   Succeeded   4	   4

NAME                                                                                       PHASE       STARTTIME   COMMAND   COMPLETIONTIME
runtimetask.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.23.0-01-upgrade-apply-paco        Succeeded   20m         3/3       12m
runtimetask.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.23.0-04-upgrade-worker-daocloud   Succeeded   12m  	   5/5       10m
runtimetask.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.24.1-01-upgrade-apply-paco        Succeeded   10m         3/3       3m25s
runtimetask.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.24.1-04-upgrade-worker-daocloud   Succeeded   3m21s       5/5       55s

NAME                                                                                   PHASE	   NODES   SUCCEEDED   FAILED
runtimetaskgroup.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.23.0-01-upgrade-apply    Succeeded   1	   1
runtimetaskgroup.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.23.0-04-upgrade-worker   Succeeded   1	   1
runtimetaskgroup.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.24.1-01-upgrade-apply    Succeeded   1       1
runtimetaskgroup.operator.kubeadm.x-k8s.io/upgrade-v1.24.1-v1.24.1-04-upgrade-worker   Succeeded   1       1

[root@paco ~]# kubelet --version
Kubernetes v1.24.1
[root@paco ~]# kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"24", GitVersion:"v1.24.1", GitCommit:"3ddd0f45aa91e2f30c70734b175631bec5b5825a", GitTreeState:"clean", BuildDate:"2022-05-24T12:24:38Z", GoVersion:"go1.18.2", Compiler:"gc", Platform:"linux/amd64"}
[root@paco ~]# kubectl version
WARNING: This version information is deprecated and will be replaced with the output from kubectl version --short.  Use --output=yaml|json to get the full version.
Client Version: version.Info{Major:"1", Minor:"24", GitVersion:"v1.24.1", GitCommit:"3ddd0f45aa91e2f30c70734b175631bec5b5825a", GitTreeState:"clean", BuildDate:"2022-05-24T12:26:19Z", GoVersion:"go1.18.2", Compiler:"gc", Platform:"linux/amd64"}
Kustomize Version: v4.5.4
Server Version: version.Info{Major:"1", Minor:"24", GitVersion:"v1.24.1", GitCommit:"3ddd0f45aa91e2f30c70734b175631bec5b5825a", GitTreeState:"clean", BuildDate:"2022-05-24T12:18:48Z", GoVersion:"go1.18.2", Compiler:"gc", Platform:"linux/amd64"}
[root@paco ~]# kubectl get node
NAME       STATUS   ROLES           AGE     VERSION
daocloud   Ready    <none>          6h36m   v1.24.1
paco       Ready    control-plane   6h37m   v1.24.1
```
