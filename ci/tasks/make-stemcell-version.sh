#!/usr/bin/env bash

set -e -x

[ -f published-stemcell/version ] || exit 1

echo "3262.15.0" > version/semver
