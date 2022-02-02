#!/bin/bash
set -euo pipefail

main() {
  cd "$(dirname "$0")/../.."

  zserio="$(bazel run --run_under=echo //zserio)"
  OUT=test/go
  rm -rf "$OUT/reference_modules"
  "$zserio" generate \
    --out ${OUT} \
    --rootpackage github.com/woven-planet/go-zserio/test/go \
    testdata/reference_modules
  gazelle
}

main "$@"
