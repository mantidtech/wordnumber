#!/bin/bash

TASK=build
if [[ "${1}" = "install" ]]; then
  TASK=install
fi

MODULE="$(grep module go.mod | cut -d ' ' -f2)"
VERSION=$(git describe --always --dirty)
BUILD=$(date -u +"%Y%m%d%H%M%S")

for p in cmd/*; do
  PROJECT="$(basename "${p}")"
  if [[ ! -d "cmd/${PROJECT}" ]] || [[ -f "cmd/${PROJECT}/.noinstall" ]]; then
    continue
  fi
  FILES="cmd/${PROJECT}/${PROJECT}.go"
  go "${TASK}" -ldflags="-X ${MODULE}.version=${VERSION} -X ${MODULE}.build=${BUILD}" "${FILES}" \
    && echo "${TASK} of ${PROJECT} ${VERSION} successful"
done
