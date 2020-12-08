#!/usr/bin/env bash
#
# Copyright 2019 HAProxy Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
set -eo pipefail

HAPROXY_VERSION=${HAPROXY_VERSION:-2.2}
DOCKER_BASE_IMAGE="${DOCKER_BASE_IMAGE:-haproxytech/haproxy-alpine}:${HAPROXY_VERSION}"
DOCKER_CONTAINER_NAME="dataplaneapi-e2e"
ROOT_DIR=$(git rev-parse --show-toplevel)
export E2E_PORT=${E2E_PORT:-8042}
export E2E_DIR=${ROOT_DIR}/e2e
export LOCAL_IP_ADDRESS=${LOCAL_IP_ADDRESS:-127.0.0.1}

if ! docker version > /dev/null 2>&1; then
  echo '>>> Docker is not installed: cannot proceed for e2e test suite'
fi

if ! docker inspect "${DOCKER_BASE_IMAGE}" > /dev/null 2>&1; then
  echo '>>> Downloading base Docker image'
  docker pull "${DOCKER_BASE_IMAGE}"
fi

echo '>>> Provisioning the e2e environment'
docker run \
  --rm \
  --detach \
  --name ${DOCKER_CONTAINER_NAME} \
  --publish "${E2E_PORT}":8080 \
  "${DOCKER_BASE_IMAGE}" > /dev/null 2>&1
docker cp "${ROOT_DIR}/build/dataplaneapi" ${DOCKER_CONTAINER_NAME}:/usr/local/bin/dataplaneapi
docker cp "${E2E_DIR}/fixtures/haproxy.cfg" ${DOCKER_CONTAINER_NAME}:/etc/haproxy/haproxy.cfg
docker cp "${E2E_DIR}/fixtures/userlist.cfg" ${DOCKER_CONTAINER_NAME}:/etc/haproxy/userlist.cfg
docker exec -d ${DOCKER_CONTAINER_NAME} sh -c "dataplaneapi --log-level=debug --userlist-file=/etc/haproxy/userlist.cfg --host=0.0.0.0 --port=8080 --reload-cmd='kill -SIGUSR2 1' --restart-cmd='kill -SIGUSR2 1'"

echo '>>> Waiting dataplane API to be up and running'
until nc -z "${LOCAL_IP_ADDRESS}" "${E2E_PORT}" 2>&1; do sleep 1; done

# deferring Docker container removal
# shellcheck disable=SC1090
source "${E2E_DIR}/libs/cleanup.bash"
trap 'cleanup ${DOCKER_CONTAINER_NAME}' EXIT

echo '>>> Starting test suite'
bats -t "${E2E_DIR}"/tests/*