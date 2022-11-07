# Supercluster Plugin

> Supercluster plugin for Kubo

This repository contains the source for the Supercluster IPFS plugin. It was adopted from the Protocol Labs [example plugin repo](https://github.com/ipfs/go-ipfs-example-plugin).

Packages:

* supercluster: daemon plugin that serves the supercluster UI and handles DB ops using the ipfs coreAPI
* ui: React frontend for designs

**NOTE 1:** Plugins only work on Linux and MacOS at the moment. You can track the progress of this issue here: https://github.com/golang/go/issues/19282

## Building and Installing
First install frontend dependencies:

``` sh
cd ui
yarn # or npm install
```


You can *build* the supercluster plugin by running `make build`. You can then install it into your local IPFS repo by running `make install`.

Plugins need to be built against the correct version of Kubo. This package generally tracks the latest Kubo release but if you need to build against a different version, please set the `IPFS_VERSION` environment variable.


You can set `IPFS_VERSION` to:

* `vX.Y.Z` to build against that version of IPFS.
* `$commit` or `$branch` to build against a specific Kubo commit or branch.
   * Note: if building against a commit or branch make sure to build that commit/branch using the -trimpath flag. For example getting the binary via `go get -trimpath github.com/ipfs/kubo/cmd/ipfs@COMMIT`
* `/absolute/path/to/source` to build against a specific Kubo checkout.

To update the Kubo version, run:

```bash
> make go.mod IPFS_VERSION=version
```

## License

MIT
