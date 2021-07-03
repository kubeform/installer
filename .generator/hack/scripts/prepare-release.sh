#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# set -eou pipefail

SCRIPT_ROOT=$(realpath "$(dirname "${BASH_SOURCE[0]}")/../..")
SCRIPT_NAME=$(basename "${BASH_SOURCE[0]}")
pushd "$SCRIPT_ROOT"

input_dir=$(realpath "${SCRIPT_ROOT}/..")
echo "using input_dir $input_dir"

providers=({{ join " " .Providers }})

for p in "${providers[@]}"; do
    # ref: https://stackoverflow.com/a/55331060/244009
    # ref: https://linuxhint.com/bash_lowercase_uppercase_strings/
    ctrl_tag="KUBEFORM_PROVIDER_${p^^}_CONTROLLER_TAG"
    if [ -z "${!ctrl_tag:-}" ]; then
        echo "${ctrl_tag} env var is not set"
        exit 1
    fi

    echo
    echo "checking out https://github.com/kubeform/provider-${p}-api.git ${!ctrl_tag}"
    cd $input_dir || exit
    git clone --no-tags --no-recurse-submodules --depth=1 "https://github.com/kubeform/provider-${p}-api.git"
    cd "provider-${p}-api" || exit
    git checkout "${!ctrl_tag}"
done

cd "$SCRIPT_ROOT" || exit

echo
cmd="go run ./hack/generate/main.go --input-dir=$input_dir"
echo "$cmd"
$cmd

if [ -z "${CHART_REGISTRY}" ]; then
    echo "CHART_REGISTRY env var is not set"
    exit 1
fi
if [ -z "${CHART_REGISTRY_URL}" ]; then
    echo "CHART_REGISTRY_URL env var is not set"
    exit 1
fi
if [ -z "${TAG}" ]; then
    echo "TAG env var is not set"
    exit 1
fi

echo
cmd="make chart-kubeform-provider"
echo "$cmd"
$cmd

for p in "${providers[@]}"; do
    # ref: https://stackoverflow.com/a/55331060/244009
    # ref: https://linuxhint.com/bash_lowercase_uppercase_strings/
    ctrl_tag="KUBEFORM_PROVIDER_${p^^}_CONTROLLER_TAG"
    if [ -z "${!ctrl_tag:-}" ]; then
        echo "${ctrl_tag} env var is not set"
        exit 1
    fi
    export CHART_VERSION=${TAG}
    export APP_VERSION="${!ctrl_tag}"

    echo "make chart-kubeform-provider-${p}"
    for crd in $(find "$input_dir/provider-${p}-api/apis" -maxdepth 1 -mindepth 1 -type d -printf '%f '); do
        cmd="make chart-kubeform-provider-${p}-${crd}-crds"
        echo "$cmd"
        $cmd
    done
done

echo
cmd="make clientset gen-crds gen-values-schema fmt"
echo "$cmd"
$cmd
