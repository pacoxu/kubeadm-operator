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
	"net/http"
	"runtime"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

// sudo curl -L --remote-name-all https://storage.googleapis.com/kubernetes-release/release/${RELEASE}/bin/linux/${ARCH}/{kubeadm,kubelet,kubectl}
const DownloadURLTemplate = "https://storage.googleapis.com/kubernetes-release/release/%s/bin/linux/%s/%s"

// download url tempalte for servers in China that cannot access googleapis.com
const BackupDownloadURLTemplate = "http://dao-get.daocloud.io/kubernetes-release/release/$s/bin/linux/$s/%s"

func GetDownloadURLTemplate() string {
	if canAccessGoogleapis() {
		return DownloadURLTemplate
	}
	return BackupDownloadURLTemplate
}

func canAccessGoogleapis() bool {
	// check a url that can be accessed by google
	_, err := http.Get("https://storage.googleapis.com/kubernetes-release/release/v1.24.0/bin/linux/amd64/kubectl")
	if err != nil {
		print(err.Error())
		return false
	} else {

		return true
	}
}

// runUpgradeKubeadm will try to download the binary from official websites;
func runUpgradeKubeadm(spec *operatorv1.UpgradeKubeadmCommandSpec, log logr.Logger) error {
	if spec.Local {
		return nil
	}

	return DownloadFromOfficialWebsite(spec.KubernetesVersion, "kubeadm", log)
}

func DownloadFromOfficialWebsite(version, bin string, log logr.Logger) error {
	var cmd *cmd

	cmd = newCmd("curl", "-L", "--remote-name-all", fmt.Sprintf(GetDownloadURLTemplate(), version, runtime.GOARCH, bin), "-o", "/usr/bin/"+bin)
	log.Info("download", "command", cmd.command, "args", strings.Join(cmd.args, " "))
	donwlod, err := cmd.RunAndCapture()
	if err != nil {
		return errors.WithStack(errors.WithMessage(err, strings.Join(donwlod, "\n")))
	}
	log.Info(fmt.Sprintf("%s", strings.Join(donwlod, "\n")))

	cmd = newCmd("chmod", "+x", "/usr/bin/"+bin)
	lines, err := cmd.RunAndCapture()
	if err != nil {
		return errors.WithStack(errors.WithMessage(err, strings.Join(lines, "\n")))
	}
	log.Info(fmt.Sprintf("%s", strings.Join(lines, "\n")))
	return nil
}
