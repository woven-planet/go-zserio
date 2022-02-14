#!/bin/bash
set -euxo pipefail

readonly dir="$(realpath "$(dirname $0)")"
cd $dir

git grep -h '^\$' ./README.md |
  sed -e 's/^\$ //g' -e 's/"/\\"/g' |
  xargs -I {} bash -exc "{}"
