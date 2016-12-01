#!/usr/bin/env bash

set -e -x

apt-get update
apt-get -y install wget

cd build
wget https://s3.amazonaws.com/dev-bosh-softlayer-cpi-stemcells/bosh-stemcell-3312.3-softlayer-hvm-ubuntu-trusty-go_agent.tgz

