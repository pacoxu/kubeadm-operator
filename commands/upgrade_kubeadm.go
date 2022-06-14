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
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"

	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

const kubeadmDownloadPath = "https://storage.googleapis.com/kubernetes-release/release/%s/bin/linux/%s/kubeadm"

func runUpgradeKubeadm(spec *operatorv1.UpgradeKubeadmCommandSpec, log logr.Logger) error {
	file, err := ioutil.TempFile(".", "upgrade.*.sh")
	if err != nil {
		log.Error(err, "Cannot create a temp file")
	}
	defer os.Remove(file.Name())

	path := fmt.Sprintf(kubeadmDownloadPath, spec.Version, runtime.GOARCH)
	script := "apt update && apt install wget -y \n" +
		"wget " + path + " -O /usr/bin/kubeadm-" + spec.Version + "\n" +
		"chmod +x /usr/bin/kubeadm-" + spec.Version + "\n" +
		"cp -f /usr/bin/kubeadm-" + spec.Version + " /usr/bin/kubeadm"

	_, err = file.Write([]byte(script))
	if err != nil {
		log.Error(err, "failed with creating the upgrade script")
	}
	err = wait.Poll(100*time.Millisecond, 30*time.Second, func() (bool, error) {
		cmd := exec.Command("sh", file.Name())
		_, err = cmd.Output()
		if err != nil {
			return false, errors.New("update failed with error: " + err.Error())
		}
		return true, nil
	})
	return err
}
