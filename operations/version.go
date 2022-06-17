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

package operations

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/version"
)

// by default, we use v1.x.0 directly, but if there is a prefered version, we use it
var preferedKubeadmVersion = map[uint]string{
	// TODO change the version to lastest version after https://github.com/kubernetes/kubeadm/issues/2426(a bug in v1.24) is fixed
	24: "v1.24.2",
}

func upgradeCheck(current, target string) (isSupported, isCrossVersion, canSkip bool) {
	currentVer, err := version.ParseSemantic(current)
	if err != nil {
		return false, false, false
	}
	targetVer, err := version.ParseSemantic(target)
	if err != nil {
		return false, false, false
	}
	return upgradeCheckVersion(currentVer, targetVer)
}

func upgradeCheckVersion(current, target *version.Version) (isSupported, isCrossVersion, canSkip bool) {
	// no need to check major as only major 1 is supported
	// if current.Major() != target.Major() {
	// 	return false
	// }
	if current.Minor() == target.Minor() {
		if current.Patch() == target.Patch() {
			// just skip as the version is the same
			return true, false, true
		}
		// upgrade to a patched version
		return true, false, false
	} else if current.Minor() < target.Minor() {
		if current.Minor()+1 == target.Minor() {
			// upgrade to a minor version
			return true, false, false
		}
		// upgrade multi-minor version, need to split into multiple upgrade tasks
		return true, true, false
	}
	// downgrade is not supported
	if current.Minor()-1 == target.Minor() {
		// TODO downgrade to a minor version
		// this is just for test purpose, need to define if it should be supported in the future
		return true, false, false
	}
	return false, false, false

}

func isSupported(ver string) bool {
	v, err := version.ParseSemantic(ver)
	if err != nil {
		return false
	}
	return isSupportedVersion(v)
}

func isSupportedVersion(ver *version.Version) bool {
	// TODO a table of supported versions needs to be created in the docs
	if ver.Major() != 1 && ver.Minor() < 17 {
		return false
	}
	return true
}

// before this, we should make sure the version is supported
func getCrossVersions(current, target string) []string {
	versions := []string{}
	cur, err := version.ParseSemantic(current)
	if err != nil {
		return versions
	}
	tar, err := version.ParseSemantic(target)
	if err != nil {
		return versions
	}
	_, isCross, _ := upgradeCheckVersion(cur, tar)
	if !isCross {
		return versions
	}
	tarMinor := tar.Minor()
	for i := cur.Minor() + 1; i < tarMinor; i++ {
		version := getPerferedVersion(i)
		versions = append(versions, version)
	}
	return versions
}

func getPerferedVersion(minor uint) string {
	if preferedKubeadmVersion[minor] != "" {
		return preferedKubeadmVersion[minor]
	}
	return fmt.Sprintf("v1.%d.0", minor)
}
