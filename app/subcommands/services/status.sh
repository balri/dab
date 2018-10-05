#!/bin/sh
# Description: Display the status of the given service, or all services if none is given
# Usage: [<TOOL_NAME>]
# vim: ft=sh ts=4 sw=4 sts=4 noet
set -euf

# shellcheck disable=SC1091
. ./lib/compose.sh

servicepose ps "$@"
