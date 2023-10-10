#!/usr/bin/env bash
set -euo pipefail

install() {
  echo Hello World
}

upgrade() {
  echo World 2.0
}

uninstall() {
  echo Goodbye World
}

write () {
  echo $1
}

# Call the requested function and pass the arguments as-is
"$@"
