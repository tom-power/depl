#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

$root/build.sh

targets=(
  concher/concher
  flanalyst/flanalyst
  flanalyst-clock
  flanalyst-lib
)

for target in "${targets[@]}"
do
  cp $root/go/build/flanalystsh ~/proj/os/flanalyst.ovh/$target/sh/sh/
  cp $root/sh/flanalyst.sh ~/proj/os/flanalyst.ovh/$target/sh/sh/
done

echo "installed"