#!/bin/bash

# build frontend
cd ui
yarn build

# build backend
cd ..
go build

# start app
rm -rf store/
if [[ -n "$BROWSER" ]]; then
    # TODO: extract port as commandline param
    $BROWSER http://localhost:4000
fi
./supercluster-client
