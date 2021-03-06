#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[@]}")/..
CODEGEN_PKG="${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}"

echo "Calling ${CODEGEN_PKG}/generate-groups.sh"
"${CODEGEN_PKG}"/generate-groups.sh all \
    riasc.eu/wice/pkg/signaling/k8s/client \
    riasc.eu/wice/pkg/signaling/k8s/apis \
    wice:v1 \
    --go-header-file="${CODEGEN_PKG}"/hack/boilerplate.go.txt \
    --trim-path-prefix riasc.eu/wice
