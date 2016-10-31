#!/usr/bin/env bash

set -e -x

[ -f published-stemcell/version ] || exit 1

published_version=3263.8.1

# check for minor (only supports x and x.x)
echo "${published_version}" > version/semver