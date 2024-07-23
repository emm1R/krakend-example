#!/usr/bin/env sh
set -eux

SCRIPT=$(realpath "$0")
SCRIPT_DIR=$(dirname "$SCRIPT")
OLD_PWD=$(pwd)

. "${SCRIPT_DIR}"/plugins.sh

echo "Building plugins"

# We must be in the directory where go.mod is
# in order to build plugins with external dependencies.
for plugin in $PLUGINS; do 
    echo building $plugin
    cd "${SCRIPT_DIR}"/${plugin}
    ${GO_BIN} build -buildmode=plugin -o ../$plugin.so *.go
done
cd "${OLD_PWD}"
