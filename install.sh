#!/bin/bash

SUPERCLUSTER_PATH=http://localhost:3000

# unzip and set permissions
unzip artifacts.zip
if [ $? -ne 0 ]; then
  echo "Unzipping failed, check if you have `unzip` installed"
  exit
fi

# make bins executable
chmod +x ipfs supercluster-plugin.so

# init IPFS and move everything to their final resting places
ipfs init
mv supercluster-plugin.so ~/.ipfs/plugins/
mv ipfs /usr/local/bin

# check if everything works
/usr/local/bin/ipfs version
if [ $? -ne 0 ]; then
  echo "Installation failed :("
  echo "Please contact support at our discord: https://discord.gg/3aVrFvdW"
  exit
fi

xdg-open $SUPERCLUSTER_PATH
if [ $? -ne 0 ]; then
  echo "Supercluster is up and running at" $SUPERCLUSTER_PATH
  exit
fi
