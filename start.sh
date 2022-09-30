#!/bin/bash

# build frontend
cd ui
yarn build

# build backend
cd ..
go build

# start app
./supercluster-client
if [[ -n "$BROWSER" ]]; then
    $BROWSER http://localhost:4000
fi
