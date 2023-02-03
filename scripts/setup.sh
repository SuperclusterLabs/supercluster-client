#!/bin/bash

bins_dir=build

if [ ! -d "../$bins_dir" ]; then
  mkdir "../$bins_dir"
fi

cd "../$bins_dir"
wget https://dist.ipfs.tech/kubo/v0.18.1/kubo_v0.18.1_linux-amd64.tar.gz
tar -xvzf kubo_v0.18.1_linux-amd64.tar.gz
rm kubo_v0.18.1_linux-amd64.tar.gz
