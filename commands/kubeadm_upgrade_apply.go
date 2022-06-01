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

package commands

import (
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

// runKubeadmUpgradeApply runs the kubeadm upgrade apply command
func runKubeadmUpgradeApply(spec *operatorv1.KubeadmUpgradeApplyCommandSpec, log logr.Logger) error {
	var cmd *cmd
	// TODO: add real dry run support
	cmd = newCmd("kubeadm", "upgrade", "apply", spec.KubernetesVersion, "--yes", "--v=4")
	if spec.DryRun {
		cmd = newCmd("kubeadm", "upgrade", "apply", spec.KubernetesVersion, "--yes", "--dry-run", "--v=5")
	}

	lines, err := cmd.RunAndCapture()
	if err != nil {
		log.Error(err, "kubeadm upgrade apply failed", strings.Join(lines, "\n"))
		return errors.WithStack(errors.WithMessage(err, strings.Join(lines, "\n")))
	}

	log.Info(fmt.Sprintf("%s", strings.Join(lines, "\n")))

	return nil
}
