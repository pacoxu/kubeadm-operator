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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func setupUpgrade() map[string]string {
	return map[string]string{}
}

func planUpgrade(operation *operatorv1.Operation, spec *operatorv1.UpgradeOperationSpec, c client.Client) *operatorv1.RuntimeTaskGroupList {
	// TODO support upgrade to v1.n-1~v1.n of current kubernetes server version.
	// If the current kubernetes server version is v1.n-2 which is below the target version, we need to generate a further upgrade plan

	var items []operatorv1.RuntimeTaskGroup
	dryRun := operation.Spec.GetTypedOperationExecutionMode() == operatorv1.OperationExecutionModeDryRun
	serverNeedUpgrade := true

	serverVersion, err := getServerVersion()
	if err != nil {
		return nil
	}
	// TODO is it the right way to check if the `kubeadm upgrade apply` is successful?
	// check ClusterConfiguration.kubernetesVersion in kubeadm-config configmap?
	if operation.Spec.Upgrade.KubernetesVersion == serverVersion {
		//skip upgrade if the current kubernetes server version is the same as the target version
		serverNeedUpgrade = false
	}

	// nodeNeedUpgrade := true
	// nodeVersion, err := getNodeVersion(c, getNodeName())
	// if err != nil {
	// 	return nil
	// }
	// if operation.Spec.Upgrade.KubernetesVersion == nodeVersion {
	// 	//skip upgrade node if the current kubernetes server version is the same as the target version
	// 	nodeNeedUpgrade = false
	// }

	if serverNeedUpgrade {
		t1 := createUpgradeApplyTaskGroup(operation, "01", "upgrade-apply")
		setCP1Selector(&t1)
		// run `upgrade apply`` on the first node of all control plane
		t1.Spec.NodeFilter = string(operatorv1.RuntimeTaskGroupNodeFilterHead)
		t1.Spec.Template.Spec.Commands = append(t1.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				UpgradeKubeadm: &operatorv1.UpgradeKubeadmCommandSpec{
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubeadmUpgradeApply: &operatorv1.KubeadmUpgradeApplyCommandSpec{
					DryRun:            dryRun,
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
				},
			},
			operatorv1.CommandDescriptor{
				UpgradeKubeletAndKubeactl: &operatorv1.UpgradeKubeletAndKubeactlCommandSpec{
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
		)
		items = append(items, t1)
	}

	// this can be skipped if there is only one control-plane node.
	// currently it depends on the selector
	t2 := createBasicTaskGroup(operation, "02", "upgrade-cp")
	setCPSelector(&t2)
	cpNodes, err := listNodesBySelector(c, &t2.Spec.NodeSelector)
	if err != nil {
		fmt.Printf("failed to list nodes: %v", err)
		return nil
	}
	if cpNodes.Size() > 1 {

		t2.Spec.Template.Spec.Commands = append(t2.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				UpgradeKubeadm: &operatorv1.UpgradeKubeadmCommandSpec{
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubeadmUpgradeNode: &operatorv1.KubeadmUpgradeNodeCommandSpec{
					DryRun: operatorv1.OperationExecutionMode(operation.Spec.ExecutionMode) == operatorv1.OperationExecutionModeDryRun,
				},
			},
			operatorv1.CommandDescriptor{
				UpgradeKubeletAndKubeactl: &operatorv1.UpgradeKubeletAndKubeactlCommandSpec{
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
		)
		items = append(items, t2)
	}

	// this can be skipped if there are no worker nodes.
	// currently it depends on the selector
	t3 := createBasicTaskGroup(operation, "02", "upgrade-w")
	setWSelector(&t3)
	workerNodes, err := listNodesBySelector(c, &t3.Spec.NodeSelector)
	if err != nil {
		fmt.Printf("failed to list nodes: %v", err)
		return nil
	}
	if workerNodes.Size() > 0 {
		t3.Spec.Template.Spec.Commands = append(t3.Spec.Template.Spec.Commands,
			operatorv1.CommandDescriptor{
				KubectlDrain: &operatorv1.KubectlDrainCommandSpec{},
			},
			operatorv1.CommandDescriptor{
				UpgradeKubeadm: &operatorv1.UpgradeKubeadmCommandSpec{
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubeadmUpgradeNode: &operatorv1.KubeadmUpgradeNodeCommandSpec{
					DryRun: operatorv1.OperationExecutionMode(operation.Spec.ExecutionMode) == operatorv1.OperationExecutionModeDryRun,
				},
			},
			operatorv1.CommandDescriptor{
				UpgradeKubeletAndKubeactl: &operatorv1.UpgradeKubeletAndKubeactlCommandSpec{
					KubernetesVersion: operation.Spec.Upgrade.KubernetesVersion,
					Local:             operation.Spec.Upgrade.Local,
				},
			},
			operatorv1.CommandDescriptor{
				KubectlUncordon: &operatorv1.KubectlUncordonCommandSpec{},
			},
		)
		items = append(items, t3)
	}

	return &operatorv1.RuntimeTaskGroupList{
		Items: items,
	}
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
