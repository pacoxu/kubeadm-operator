/*
Copyright 2022 The Kubernetes Authors.

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

package commands

import (
	"github.com/go-logr/logr"

	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

// TODO it depends on https://github.com/kubernetes/kubeadm/issues/1318
// Currently, kubeadm upgrade apply doesn't support phase
func runUpgradeKubeProxy(spec *operatorv1.KubeadmUpgradeKubeProxySpec, log logr.Logger) error {
	return nil
}
