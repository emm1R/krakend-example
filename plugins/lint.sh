#!/usr/bin/env sh
set -uex

SCRIPT=$(realpath "$0")
SCRIPT_DIR=$(dirname "$SCRIPT")
OLD_PWD=$(pwd)

. "${SCRIPT_DIR}"/plugins.sh

echo "Linting plugins"

# lint plugins in their respective directories
for plugin in $PLUGINS; do
    echo "linting $plugin"
    cd "${SCRIPT_DIR}"/${plugin}
    golangci-lint run ./... --timeout 2m -E bodyclose,gocritic,gofmt,gosec,govet,nestif,nlreturn,revive,rowserrcheck --exclude dot-imports
done
cd "${OLD_PWD}"
