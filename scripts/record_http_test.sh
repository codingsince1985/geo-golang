#!/bin/bash
set -e

curdir="$( builtin cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
projectdir="${curdir}/.."

# Call tests and enable DVR HTTP recording
DVR_MODE="record" "${projectdir}/scripts/test.sh"
