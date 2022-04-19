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

set -o nounset
set -o pipefail

# shellcheck source=/dev/null
source "$(dirname "$0")/utils.sh"

# cd to the root path
cd_root_path

failure() {
    if [[ "${1}" != 0 ]]; then
        res=1
        failed+=("${2}")
        outputs+=("${3}")
    fi
}

# exit code, if a script fails we'll set this to 1
res=0
failed=()
outputs=()

# run all verify scripts, optionally skipping any of them

if [[ "${VERIFY_WHITESPACE:-true}" == "true" ]]; then
  echo "[*] Verifying whitespace..."
  out=$(hack/verify-whitespace.sh 2>&1)
  failure $? "verify-whitespace.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_SPELLING:-true}" == "true" ]]; then
  echo "[*] Verifying spelling..."
  out=$(hack/verify-spelling.sh 2>&1)
  failure $? "verify-spelling.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_BOILERPLATE:-true}" == "true" ]]; then
  echo "[*] Verifying boilerplate..."
  out=$(hack/verify-boilerplate.sh 2>&1)
  failure $? "verify-boilerplate.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_GOFMT:-true}" == "true" ]]; then
  echo "[*] Verifying gofmt..."
  out=$(hack/verify-gofmt.sh 2>&1)
  failure $? "verify-gofmt.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_GOIMPORTS:-true}" == "true" ]]; then
  echo "[*] Verifying goimports..."
  out=$(hack/verify-goimports.sh 2>&1)
  failure $? "verify-goimports.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_GOLINT:-true}" == "true" ]]; then
  echo "[*] Verifying golint..."
  out=$(hack/verify-golint.sh 2>&1)
  failure $? "verify-golint.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_GOVET:-true}" == "true" ]]; then
  echo "[*] Verifying govet..."
  out=$(hack/verify-govet.sh 2>&1)
  failure $? "verify-govet.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_DEPS:-true}" == "true" ]]; then
  echo "[*] Verifying deps..."
  out=$(hack/verify-deps.sh 2>&1)
  failure $? "verify-deps.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_GOTEST:-true}" == "true" ]]; then
  echo "[*] Verifying gotest..."
  out=$(hack/verify-gotest.sh 2>&1)
  failure $? "verify-gotest.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_BUILD:-true}" == "true" ]]; then
  echo "[*] Verifying build..."
  out=$(hack/verify-build.sh 2>&1)
  failure $? "verify-build.sh" "${out}"
  cd_root_path
fi

if [[ "${VERIFY_DOCKER_BUILD:-true}" == "true" ]]; then
 echo "[*] Verifying manager docker image build..."
 out=$(hack/verify-docker-build.sh 2>&1)
 failure $? "verify-docker-build.sh" "${out}"
 cd_root_path
fi

# exit based on verify scripts
if [[ "${res}" = 0 ]]; then
  echo ""
  echo "All verify checks passed, congrats!"
else
  echo ""
  echo "Some of the verify scripts failed:"
  for i in "${!failed[@]}"; do
      echo "- ${failed[$i]}:"
      echo "${outputs[$i]}"
      echo
  done
fi
exit "${res}"
