# Supercluster Plugin

> Supercluster plugin for Kubo

This repository contains the source for the Supercluster IPFS client.

## Quickstart
If you're running Ubuntu 22.04, you can get started immediately with our pre-generated package from release. To do this:
- Download the latest release binaries, shipped as `supercluster.zip`
- Unzip the file
- Run the install script after making it executable using `chmod +x ./install.sh && ./install.sh`

## Building and Installing
- First you should build the UI:
  - clone the [supercluster UI](https://github.com/SuperclusterLabs/supercluster-ui-svelte.git) and build it with `yarn build`
  - copy the generated UI build folder into the `ui` folder. This is required for embedding the ui into the exe
-  You can run `make build` if you just want to build the binary. This will set up the dev environment by fetching kubo and putting it into the `build` folder
- Run `make install` to run build (if it wasn't already), and install it to the system. This moves the supercluster binary to `/usr/local/bin`
- You can now run the backend by running the `supercluster` command

## License

MIT
