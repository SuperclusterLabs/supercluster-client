#!/bin/bash

SUPERCLUSTER_URL=http://localhost:3030
SUPERCLUSTER_DIR="$(echo $HOME)/.supercluster"
SUPERCLUSTER_BINS="supercluster ipfs"

cp -r . $SUPERCLUSTER_DIR

# start backend
nohup ./supercluster >> ./logs.txt &

if [ $? -ne 0 ]; then
  echo "Installation failed :("
  echo "Please contact support at our discord: https://discord.gg/3aVrFvdW"
  exit
fi

xdg-open $SUPERCLUSTER_URL
if [ $? -ne 0 ]; then
  echo "Supercluster is up and running at" $SUPERCLUSTER_URL
  exit
fi
