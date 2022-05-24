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

package v1alpha1

// CommandDescriptor represents a command to be performed.
// Only one of its members may be specified.
type CommandDescriptor struct {
	// +optional
	KubeadmRenewCertificates *KubeadmRenewCertsCommandSpec `json:"kubeadmRenewCertificates,omitempty"`

	// +optional
	KubeadmUpgradeApply *KubeadmUpgradeApplyCommandSpec `json:"kubeadmUpgradeApply,omitempty"`

	// +optional
	KubeadmUpgradeNode *KubeadmUpgradeNodeCommandSpec `json:"kubeadmUpgradeNode,omitempty"`

	// +optional
	Preflight *PreflightCommandSpec `json:"preflight,omitempty"`

	// +optional
	UpgradeKubeadm *UpgradeKubeadmCommandSpec `json:"upgradeKubeadm,omitempty"`

	// +optional
	UpgradeKubeletAndKubeactl *UpgradeKubeletAndKubeactlCommandSpec `json:"upgradeKubeletAndKubeactl,omitempty"`

	// +optional
	KubectlDrain *KubectlDrainCommandSpec `json:"kubectlDrain,omitempty"`

	// +optional
	KubectlUncordon *KubectlUncordonCommandSpec `json:"kubectlUncordon,omitempty"`

	// Pass provide a dummy command for testing the kubeadm-operator workflow.
	// +optional
	Pass *PassCommandSpec `json:"pass,omitempty"`

	// Fail provide a dummy command for testing the kubeadm-operator workflow.
	// +optional
	Fail *FailCommandSpec `json:"fail,omitempty"`

	// Wait pauses the execution on the next command for a given number of seconds.
	// +optional
	Wait *WaitCommandSpec `json:"wait,omitempty"`
}

// PreflightCommandSpec provides...
type PreflightCommandSpec struct {

	// INSERT ADDITIONAL SPEC FIELDS -
	// Important: Run "make" to regenerate code after modifying this file
}

// UpgradeKubeadmCommandSpec provides...
type UpgradeKubeadmCommandSpec struct {
	// +optional
	// KubernetesVersion specifies the target kubernetes version
	// If the version is empty, we will skip this command.
	KubernetesVersion string `json:"kubernetesVersion"`

	// +optional
	// Local is by default true.
	// If true, operator will use the binary in /usr/bin
	// If not, operator will download the binary from official websites
	Local bool `json:"local,omitempty"`

	// INSERT ADDITIONAL SPEC FIELDS -
	// Important: Run "make" to regenerate code after modifying this file
}

// KubeadmUpgradeApplyCommandSpec provides...
type KubeadmUpgradeApplyCommandSpec struct {
	// +optional
	// KubernetesVersion specifies the target kubernetes version
	// If the version is empty, we will skip this command.
	KubernetesVersion string `json:"kubernetesVersion"`
	// for dry run mode
	DryRun bool `json:"dryRun,omitempty"`
}

// TODO download the specified version bin and replace it in the node
// KubeadmUpgradeNodeCommandSpec provides...
type KubeadmUpgradeNodeCommandSpec struct {
	// for dry run mode
	DryRun bool `json:"dryRun,omitempty"`
}

// KubectlDrainCommandSpec provides...
type KubectlDrainCommandSpec struct {
	// +optional
	// Nodes is a list of nodes to uncordon
	Nodes []string `json:"nodes"`

	// TODO support dry-run
	// TODO support ignore daemonsets

	// INSERT ADDITIONAL SPEC FIELDS -
	// Important: Run "make" to regenerate code after modifying this file
}

// KubectlUncordonCommandSpec provides...
type KubectlUncordonCommandSpec struct {
	// +optional
	// Nodes is a list of nodes to uncordon
	Nodes []string `json:"nodes"`

	// INSERT ADDITIONAL SPEC FIELDS -
	// Important: Run "make" to regenerate code after modifying this file
}

// TODO download the specified version bin and replace it in the node
// UpgradeKubeletAndKubeactlCommandSpec provides...
type UpgradeKubeletAndKubeactlCommandSpec struct {
	// +optional
	// KubernetesVersion specifies the target kubernetes version
	// If the version is empty, we will skip this command.
	KubernetesVersion string `json:"kubernetesVersion"`

	// +optional
	// Local is by default true.
	// If true, operator will use the binary in /usr/bin
	// If not, operator will download the binary from official websites
	Local bool `json:"local,omitempty"`

	// INSERT ADDITIONAL SPEC FIELDS -
	// Important: Run "make" to regenerate code after modifying this file
}

// KubeadmRenewCertsCommandSpec provides...
type KubeadmRenewCertsCommandSpec struct {
	// +optional
	// Commands is a list of commands to run: all, apiserver, apiserver-etcd-client, apiserver-kubelet-client,
	// controller-manager.conf, etcd-healthcheck-client, etcd-peer, etcd-server, front-proxy-client, scheduler.conf
	Commands []string `json:"commands"`

	// INSERT ADDITIONAL SPEC FIELDS -
	// Important: Run "make" to regenerate code after modifying this file
}

// PassCommandSpec provide a dummy command for testing the kubeadm-operator workflow.
type PassCommandSpec struct {
}

// FailCommandSpec provide a dummy command for testing the kubeadm-operator workflow.
type FailCommandSpec struct {
}

// WaitCommandSpec pauses the execution on the next command for a given number of seconds.
type WaitCommandSpec struct {
	// Seconds to pause before next command.
	// +optional
	Seconds int32 `json:"seconds,omitempty"`
}
