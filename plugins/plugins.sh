#!/usr/bin/env sh
set -uex

# This file contains shared functionality for `build.sh` and `lint.sh`.
# Basically, it contains the list of plugins

PLUGINS='
allas-header-client
profile-client
'

GO_BIN=${GO_BIN:-$(which go)}
