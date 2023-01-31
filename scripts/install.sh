#!/bin/bash

SUPERCLUSTER_URL=http://localhost:3000
SUPERCLUSTER_DIR="$(echo $HOME)/.supercluster"
SUPERCLUSTER_BINS="supercluster ipfs"

# unzip and set permissions
# this should contain ipfs, supercluster, ipfs-cluster bins
unzip artifacts.zip
if [ $? -ne 0 ]; then
  echo "Unzipping failed, check if you have `unzip` installed"
  exit
fi

# make bins executable
chmod +x $SUPERCLUSTER_BINS

if [[ ! -d $SUPERCLUSTER_DIR ]]; then
    mkdir $SUPERCLUSTER_DIR
fi

# move everything to their final resting place and init IPFS
mv $SUPERCLUSTER_BINS $SUPERCLUSTER_DIR
cd $SUPERCLUSTER_DIR
./ipfs init

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
