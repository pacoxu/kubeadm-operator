#!/usr/bin/env bash
# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# shellcheck source=/dev/null
source "$(dirname "$0")/utils.sh"
# cd to the root path
cd_root_path

# create a temporary directory
TMP_DIR=$(mktemp -d)

# cleanup
exitHandler() (
  echo "Cleaning up..."
  rm -rf "${TMP_DIR}"
)
trap exitHandler EXIT

# pull goimports
export GO111MODULE=on
URL="https://github.com/golang/tools.git"
git clone --quiet --depth=1 "${URL}" "${TMP_DIR}"
pushd "${TMP_DIR}" > /dev/null
popd > /dev/null

# build goimports
BIN_PATH="${TMP_DIR}/cmd/goimports"
pushd "${BIN_PATH}" > /dev/null
echo "Building goimports..."
go build > /dev/null
popd > /dev/null

# check for goimports diffs
diff=$(git ls-files | grep "\.go$" | grep -v -e "zz_generated" | xargs "${BIN_PATH}/goimports" -local k8s.io/kubeadm/operator -d  2>&1)
if [[ -n "${diff}" ]]; then
  echo "${diff}"
  echo
  echo "Check failed. Please run hack/update-goimports.sh"
  exit 1
fi
