#!/bin/bash

CUR_DIR=$(pwd)
SUPERCLUSTER_DIR="$(echo $HOME)/.supercluster"
KUBO_DIR="$(echo $HOME)/.supercluster/kubo"
IPFS_CLUSTER_SERVICE_DIR="$(echo $HOME)/.supercluster/ipfs-cluster"
KUBO_TAR=""
IPFS_CLUSTER_SERVICE_TAR=""

if [ ! -d $SUPERCLUSTER_DIR ]; then
  mkdir $SUPERCLUSTER_DIR
fi

cd $SUPERCLUSTER_DIR
mkdir clusters logs

case $(uname) in
  "Linux")
      KUBO_TAR="kubo_v0.18.1_linux-amd64.tar.gz"
      IPFS_CLUSTER_SERVICE_TAR="ipfs-cluster-service_v1.0.5_linux-amd64.tar.gz"
      ;;
  "Darwin")
      # default to M1
      KUBO_TAR="kubo_v0.18.1_darwin-arm64.tar.gz"
      IPFS_CLUSTER_SERVICE_TAR="ipfs-cluster-service_v1.0.5_darwin-arm64.tar.gz"
      ;;
  *)
      echo "Unsupported OS"
      exit
      ;;
esac

wget https://dist.ipfs.tech/kubo/v0.18.1/$KUBO_TAR
if [ $? -ne 0 ]; then
    exit
fi
wget http://dist-ipfs-tech.ipns.localhost:48084/ipfs-cluster-service/v1.0.5/$IPFS_CLUSTER_SERVICE_TAR
if [ $? -ne 0 ]; then
    exit
fi

tar -xvzf $KUBO_TAR
tar -xvzf $IPFS_CLUSTER_SERVICE_TAR

rm $KUBO_TAR $IPFS_CLUSTER_SERVICE_TAR

cd $CUR_DIR
cp -r $SUPERCLUSTER_DIR ./build
cp ./scripts/install.sh ./build
