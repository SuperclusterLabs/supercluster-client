# Supercluster Plugin

> Supercluster plugin for Kubo

This repository contains the source for the Supercluster IPFS plugin. It was adopted from the Protocol Labs [example plugin repo](https://github.com/ipfs/go-ipfs-example-plugin).

Packages:

* supercluster: daemon plugin that serves the supercluster UI and handles DB ops using the ipfs coreAPI
* ui: React frontend for designs

**NOTE 1:** Plugins only work on Linux and MacOS at the moment. You can track the progress of this issue here: https://github.com/golang/go/issues/19282

## Building and Installing
### Tl;dr
- First install frontend dependencies:

``` sh
cd ui
yarn # or npm install
```

- clone [kubo](https://github.com/ipfs/kubo)
- install it by running the following command in the kubo dir:
go install "-trimpath" -ldflags="-X "github.com/ipfs/kubo".CurrentCommit=38117db6f-dirty" -asmflags=all=-trimpath="" -gcflags=all=-trimpath="" ./cmd/ipfs
- point the plugin to kubo's dir by setting the IPFS_VERSION env variable, for e.g.: export IPFS_VERSION=~/dev/kubo
- from the plugin dir, run make install
- If everything goes well IPFS will have some extra info:

``` bash
$ ipfs version
Hello init!
ipfs version 0.16.0
```

- You can now run the backend by running `ipfs daemon` (it'll spin up our server on port 3000)

### More details

You can build the supercluster plugin by running `make build`. You can then install it into your local IPFS repo by running `make install`.

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
