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

## Demo 

After installation, a deploy named `operator-controller-manager` is running in namespace `operator-system`.
```
[root@daocloud ~]# kubectl get pod -n operator-system -o wide
NAME                                          READY   STATUS    RESTARTS   AGE   IP               NODE       NOMINATED NODE   READINESS GATES
operator-controller-manager-64c448f5b-p682x   2/2     Running   0          77m   172.32.230.197   daocloud   <none>           <none>
```

If you create a dry-run upgrade operation, there will be a runtimetaqskgroup with 

```
[root@daocloud ~]# cat up.yaml
apiVersion: operator.kubeadm.x-k8s.io/v1alpha1
kind: Operation
metadata:
  name: upgrade-1
spec:
  executionMode: DryRun
  upgrade:
    kubernetesVersion: v1.23.4
    
[root@daocloud ~]# kubectl get runtimetaskgroup -w
NAME                        PHASE     NODES   SUCCEEDED   FAILED
upgrade-1-01-upgrade-cp-1   Running   1
upgrade-1-01-upgrade-cp-1   Succeeded   1       1
upgrade-1-02-upgrade-cp-n
upgrade-1-02-upgrade-cp-n   Running
upgrade-1-02-upgrade-cp-n   Succeeded
upgrade-1-02-upgrade-w
upgrade-1-02-upgrade-w      Running
upgrade-1-02-upgrade-w      Succeeded
upgrade-1-02-upgrade-w      Succeeded    
```

After the operation is done, the operation and task group are all `Succeeded`.

```
[root@daocloud ~]# kubectl get operations
NAME        PHASE       GROUPS   SUCCEEDED   FAILED
upgrade-1   Succeeded   3        3
[root@daocloud ~]# kubectl get runtimetaskgroup
NAME                        PHASE       NODES   SUCCEEDED   FAILED
upgrade-1-01-upgrade-cp-1   Succeeded   1       1
upgrade-1-02-upgrade-cp-n   Succeeded
upgrade-1-02-upgrade-w      Succeeded
[root@daocloud ~]# kubectl get runtimetask
NAME                                 PHASE       STARTTIME   COMMAND   COMPLETIONTIME
upgrade-1-01-upgrade-cp-1-daocloud   Succeeded   75m         3/3       75m


```
