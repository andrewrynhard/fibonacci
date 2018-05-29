#!/bin/bash

set -e

CGO_ENABLED=1
GOPACKAGES=$(go list ./... | grep -v /vendor/ | grep -v /generated/)
GOFILES=$(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./pkg/generated/*")

lint_packages() {
  echo "Linting packages"
  gometalinter --exclude=pkg/generated/* --aggregate --sort=path --sort=severity --sort=line --vendor --enable-all  --disable=test --disable=gas --disable=gotype --disable=lll --disable=safesql --deadline=600s ./...
}

perform_unit_tests() {
  echo "Performing short tests"
  go test -v -short ./...
}

perform_integration_tests() {
  echo "Performing tests"
  go test -v ./...
}

perform_coverage_tests() {
  echo "Performing tests with coverage"
  local coverage_report="coverage.txt"
  local profile="profile.out"
  if [[ -f ${coverage_report} ]]; then
    rm ${coverage_report}
  fi
  touch ${coverage_report}
  for package in ${GOPACKAGES[@]}; do
    go test -v -race -coverprofile=${profile} -covermode=atomic $package
    if [ -f ${profile} ]; then
      cat ${profile} >> ${coverage_report}
      rm ${profile}
    fi
  done
}

case $1 in
  --lint)
  lint_packages
  ;;
  --unit)
  perform_unit_tests
  ;;
  --integration)
  perform_integration_tests
  ;;
  --coverage)
  perform_coverage_tests
  ;;
  *)
  exit 1
  ;;
esac

exit 0
