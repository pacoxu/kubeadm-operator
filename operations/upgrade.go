/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package operations

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func setupUpgrade() map[string]string {
	return map[string]string{}
}

func planUpgrade(operation *operatorv1.Operation, spec *operatorv1.UpgradeOperationSpec, c client.Client) (*operatorv1.RuntimeTaskGroupList, error) {
	log := ctrl.Log.WithName("operations").WithName("Upgrade").WithValues("task", operation.Name)
	var items []operatorv1.RuntimeTaskGroup

	serverVersion, err := getServerVersion()
	if err != nil {
		log.Error(err, "get server version failed")
		return nil, errors.Wrap(err, "Failed to get server version")
	}
	isServerSupported, isServerCrossVersion, isServerCanSkip := upgradeCheck(serverVersion, spec.KubernetesVersion)
	if !isServerSupported {
		log.Info("Upgrade is not supported", "serverVersion", serverVersion, "kubernetesVersion", spec.KubernetesVersion)
		// TODO current not supported operation will succeed immeditely.
		return nil, errors.Wrap(err, "Upgrade is not supported")
	}
	log.Info("Upgrade is supported", "serverVersion", serverVersion, "kubernetesVersion", spec.KubernetesVersion)
	nodes, err := listNodesBySelector(c, getAllSelector())
	if err != nil {
		log.Error(err, "list nodes failed")
		return nil, errors.Wrap(err, "Failed to list nodes")
	}
	var isClientSupported, isClientCrossVersion, isClientCanSkip bool = true, false, true
	var clientServerMatch bool = true
	log.Info("nodes list", "nodes", len(nodes.Items))
	for _, n := range nodes.Items {
		supported, cross, skip := upgradeCheck(n.Status.NodeInfo.KubeletVersion, spec.KubernetesVersion)
		isClientSupported = isClientSupported && supported
		isClientCrossVersion = isClientCrossVersion || cross
		isClientCanSkip = isClientCanSkip && skip
		clientServerMatch = clientServerMatch && isClientServerMatch(n.Status.NodeInfo.KubeletVersion, serverVersion)
		if n.Status.NodeInfo.KubeletVersion != serverVersion {
			log.Info("node is not match server version", "node", n.Name, "serverVersion", serverVersion, "kubeletVersion", n.Status.NodeInfo.KubeletVersion)
		}
	}
	if !isClientSupported {
		log.Info("Upgrade is not supported", "clientVersion", spec.KubernetesVersion)
		return nil, errors.Wrap(err, "Upgrade is not supported for some nodes' kubelet version")
	}
	log.Info("show all client version check results", "isClientSupported", isClientSupported, "isClientCrossVersion", isClientCrossVersion, "isClientCanSkip", isClientCanSkip, "clientServerMatch", clientServerMatch)
	log.Info("show all server version check results", "isServerSupported", isServerSupported, "isServerCrossVersion", isServerCrossVersion, "isServerCanSkip", isServerCanSkip)
	if isClientCanSkip && isServerCanSkip {
		// skip upgrade directly
		return &operatorv1.RuntimeTaskGroupList{
			Items: items,
		}, nil
	} else if isClientCrossVersion || isServerCrossVersion {
		// support upgrade to v1.n-1~v1.n of current kubernetes server version.
		// If the current kubernetes server version is v1.n-2 which is below the target version, we need to generate a further upgrade plan
		log.Info("Upgrade is not supported, need cross version for client or server", "targetVersion", spec.KubernetesVersion, "serverVersion", serverVersion)
		if !clientServerMatch {
			// upgrade nodes to the target version
			log.Info("[cross-upgrade] add items to make server client match", "serverVersion", serverVersion)
			items = append(items, planNextUpgrade(operation, serverVersion, c, true)...)
		}
		crossVersions := getCrossVersions(serverVersion, spec.KubernetesVersion)
		for _, v := range crossVersions {
			log.Info("[cross-upgrade] add items to upgrade to a middle version", "version", v)
			items = append(items, planNextUpgrade(operation, v, c, false)...)
		}
		log.Info("[cross-upgrade] add items to upgrade to the target version", "version", spec.KubernetesVersion)
		items = append(items, planNextUpgrade(operation, operation.Spec.Upgrade.KubernetesVersion, c, false)...)

	} else {
		log.Info("add items to upgrade to the target version", "version", spec.KubernetesVersion)
		items = append(items, planNextUpgrade(operation, operation.Spec.Upgrade.KubernetesVersion, c, isServerCanSkip)...)
	}

	return &operatorv1.RuntimeTaskGroupList{
		Items: items,
	}, nil
}

// the version may not be operation.Spec.Upgrade.KubernetesVersion for cross upgrade
func planNextUpgrade(operation *operatorv1.Operation, version string, c client.Client, isServerCanSkip bool) []operatorv1.RuntimeTaskGroup {
	log := ctrl.Log.WithName("operations").WithName("Upgrade").WithValues("task", operation.Name)
	log.Info("add task for upgrading", "version", version, "isServerCanSkip", isServerCanSkip)

	var items []operatorv1.RuntimeTaskGroup
	dryRun := operation.Spec.GetTypedOperationExecutionMode() == operatorv1.OperationExecutionModeDryRun

	if !isServerCanSkip {
		t1 := createUpgradeApplyTaskGroup(operation, fmt.Sprintf("%s-01", version), "upgrade-apply")
		setCP1Selector(&t1)
		// run `upgrade apply`` on the first node of all control plane
		t1.Spec.NodeFilter = string(operatorv1.RuntimeTaskGroupNodeFilterHead)
		t1.Spec.Template.Spec.Commands = append(t1.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				UpgradeKubeadm: &operatorv1.UpgradeKubeadmCommandSpec{
					KubernetesVersion: version,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubeadmUpgradeApply: &operatorv1.KubeadmUpgradeApplyCommandSpec{
					DryRun:            dryRun,
					KubernetesVersion: version,
					SkipKubeProxy:     operation.Spec.Upgrade.UpgradeKubeProxyAtLast,
				},
			},
			// as it depends on kubelet-reloader, we need to run it after upgrade-kubeadm apply
			operatorv1.CommandDescriptor{
				UpgradeKubeletAndKubeactl: &operatorv1.UpgradeKubeletAndKubeactlCommandSpec{
					KubernetesVersion: version,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
		)
		log.Info("add upgrade-apply task group", "task", t1.Name)
		items = append(items, t1)
	}

	// this can be skipped if there is only one control-plane node.
	// currently it depends on the selector
	t2 := createBasicTaskGroup(operation, fmt.Sprintf("%s-02", version), "upgrade-cp")
	setCPSelector(&t2)
	cpNodes, err := listNodesBySelector(c, &t2.Spec.NodeSelector)
	if err != nil {
		log.Info("failed to list nodes.", "error", err)
		return items
	}
	// Workaround: if isServerCanSkip, the only server's kubelet may need to upgrade.
	// TODO to make more accurate task, we may cut tasks into small ones, for instance, upgrade kubelet only.
	if len(cpNodes.Items) > 1 || isServerCanSkip {

		t2.Spec.Template.Spec.Commands = append(t2.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				UpgradeKubeadm: &operatorv1.UpgradeKubeadmCommandSpec{
					KubernetesVersion: version,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubeadmUpgradeNode: &operatorv1.KubeadmUpgradeNodeCommandSpec{
					DryRun: operatorv1.OperationExecutionMode(operation.Spec.ExecutionMode) == operatorv1.OperationExecutionModeDryRun,
				},
			},
			// as it depends on kubelet-reloader, we need to run it after upgrade-kubeadm
			operatorv1.CommandDescriptor{
				UpgradeKubeletAndKubeactl: &operatorv1.UpgradeKubeletAndKubeactlCommandSpec{
					KubernetesVersion: version,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
		)
		log.Info("add upgrade-cp task group", "task", t2.Name)
		items = append(items, t2)
	}

	if operation.Spec.Upgrade.UpgradeKubeProxyAtLast {
		t3 := createBasicTaskGroup(operation, fmt.Sprintf("%s-03", version), "upgrade-kube-proxy")
		t3.Spec.Template.Spec.Commands = append(t3.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				KubeadmUpgradeKubeProxy: &operatorv1.KubeadmUpgradeKubeProxySpec{
					KubernetesVersion: version,
				},
			},
		)
		log.Info("add upgrade-kube-proxy task group", "task", t3.Name)
		items = append(items, t3)
	}

	// this can be skipped if there are no worker nodes.
	// currently it depends on the selector
	t4 := createBasicTaskGroup(operation, fmt.Sprintf("%s-04", version), "upgrade-worker")
	setWSelector(&t4)
	workerNodes, err := listNodesBySelector(c, &t4.Spec.NodeSelector)
	if err != nil {
		fmt.Printf("failed to list nodes: %v", err)
		return items
	}
	log.Info("workerNodes check", "workerNum", len(workerNodes.Items))
	if len(workerNodes.Items) > 0 {
		t4.Spec.Template.Spec.Commands = append(t4.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				KubectlDrain: &operatorv1.KubectlDrainCommandSpec{},
			},
			operatorv1.CommandDescriptor{
				UpgradeKubeadm: &operatorv1.UpgradeKubeadmCommandSpec{
					KubernetesVersion: version,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				UpgradeKubeletAndKubeactl: &operatorv1.UpgradeKubeletAndKubeactlCommandSpec{
					KubernetesVersion: version,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubeadmUpgradeNode: &operatorv1.KubeadmUpgradeNodeCommandSpec{
					DryRun: operatorv1.OperationExecutionMode(operation.Spec.ExecutionMode) == operatorv1.OperationExecutionModeDryRun,
				},
			},
			operatorv1.CommandDescriptor{
				KubectlUncordon: &operatorv1.KubectlUncordonCommandSpec{},
			},
		)
		log.Info("add upgrade-w task group", "task", t4.Name)
		items = append(items, t4)
	}
	return items
}

// check the current kubernetes server version
func getServerVersion() (string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", err
	}
	clusterclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("failed to create a cluster client: %w", err)
	}
	clusterversion, err := clusterclient.Discovery().ServerVersion()
	return clusterversion.String(), nil
}

func getNodeVersion(c client.Client, nodeName string) (string, error) {
	node := &corev1.Node{}
	err := c.Get(context.TODO(), types.NamespacedName{Name: nodeName}, node)
	if err != nil {
		return "", err
	}
	return node.Status.NodeInfo.KubeletVersion, nil
}

func listNodesBySelector(c client.Client, selector *metav1.LabelSelector) (*corev1.NodeList, error) {
	s, err := metav1.LabelSelectorAsSelector(selector)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert TaskGroup.Spec.NodeSelector to a selector")
	}

	nodes := &corev1.NodeList{}
	if err := c.List(
		context.Background(), nodes,
		client.MatchingLabelsSelector{Selector: s},
	); err != nil {
		return nil, err
	}

	return nodes, nil
}

func getNodeName() string {
	return os.Getenv("MY_NODE_NAME")
}
