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
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/util/wait"
	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

// runUpgradeKubectlAndKubelet skip download when using local.
// If not local, download kubectl and kubelet. Kubelet replacement needs to be done after kubelet stopped.
func runUpgradeKubectlAndKubelet(spec *operatorv1.UpgradeKubeletAndKubeactlCommandSpec, log logr.Logger) error {
	if spec.Local {
		return nil
	}

	err := wait.Poll(100*time.Millisecond, 300*time.Second, func() (bool, error) {
		if err := DownloadFromOfficialWebsite(spec.KubernetesVersion, "kubectl", log); err != nil {
			log.Error(err, "Failed to download kubectl and kubelet")
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return err
	}

	var cmd *cmd

	cmd = newCmd("systemctl", "stop", "kubelet")
	stop, err := cmd.RunAndCapture()
	if err != nil {
		return errors.WithStack(errors.WithMessage(err, strings.Join(stop, "\n")))
	}
	log.Info(fmt.Sprintf("%s", strings.Join(stop, "\n")))

	err = DownloadFromOfficialWebsite(spec.KubernetesVersion, "kubelet", log)
	if err != nil {
		return err
	}

	cmd = newCmd("systemctl", "start", "kubelet")
	start, err := cmd.RunAndCapture()
	if err != nil {
		return errors.WithStack(errors.WithMessage(err, strings.Join(start, "\n")))
	}
	log.Info(fmt.Sprintf("%s", strings.Join(start, "\n")))

	return nil
}
