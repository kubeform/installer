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

set -e

crd_dir=${1:-}

api_repo_url=${API_REPO_URL} # eg: https://github.com/stashed/apimachinery.git
api_repo_tag=${API_REPO_TAG:-master}

if [ "$#" -ne 1 ]; then
    if [ "${api_repo_tag}" == "master" ]; then
        echo "Error: missing path_to_input_crds_directory"
        echo "Usage: import-crds.sh <path_to_input_crds_directory>"
        exit 1
    fi

    tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
    # always cleanup temp dir
    # ref: https://opensource.com/article/20/6/bash-trap
    trap \
        "{ rm -rf "${tmp_dir}"; }" \
        SIGINT SIGTERM ERR EXIT

    mkdir -p ${tmp_dir}
    pushd $tmp_dir
    git clone $api_repo_url
    repo_dir=$(ls -b1)
    cd $repo_dir
    git checkout $api_repo_tag
    popd
    crd_dir=${tmp_dir}/${repo_dir}/crds
fi

crd-importer \
    --input=${crd_dir} \
    --out=./charts/stash-crds/crds \
    --gk=BackupConfiguration.stash.appscode.com \
    --gk=BackupSession.stash.appscode.com \
    --gk=Function.stash.appscode.com \
    --gk=Repository.stash.appscode.com \
    --gk=RestoreSession.stash.appscode.com \
    --gk=Task.stash.appscode.com

crd-importer \
    --input=${crd_dir} \
    --out=. --output-yaml=stash-crds.yaml \
    --gk=BackupConfiguration.stash.appscode.com \
    --gk=BackupSession.stash.appscode.com \
    --gk=Function.stash.appscode.com \
    --gk=Repository.stash.appscode.com \
    --gk=RestoreSession.stash.appscode.com \
    --gk=Task.stash.appscode.com

crd-importer --v=v1beta1 \
    --input=${crd_dir} \
    --out=./charts/stash-catalog/crds \
    --gk=Function.stash.appscode.com \
    --gk=Task.stash.appscode.com

crd-importer --v=v1beta1 \
    --input=${crd_dir} \
    --out=. --output-yaml=stash-catalog-crds.yaml \
    --gk=Function.stash.appscode.com \
    --gk=Task.stash.appscode.com

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/kubernetes-1.21.1/crds/metrics.appscode.com_metricsconfigurations.v1.yaml \
    --out=./charts/stash-metrics/crds