package main

import (
	"log"
	"os"

	"github.com/SuperclusterLabs/supercluster-client/db"
	"github.com/SuperclusterLabs/supercluster-client/proc"
	"github.com/SuperclusterLabs/supercluster-client/router"
	"github.com/SuperclusterLabs/supercluster-client/runtime"
	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/ui"
	"github.com/SuperclusterLabs/supercluster-client/util"
)

func main() {
	confDir := util.GetConfDir()

	_, err := os.Stat(confDir)
	if err != nil {
		panic("Supercluster dir doesn't exist")
	}

	// TODO: remove firebase
	db, err := db.NewFirebaseDB()
	if err != nil {
		panic(err)
	}

	// Start IPFS
	ipfs, err := proc.NewIPFSProcess()
	if err != nil {
		panic(err)
	}
	if err = ipfs.Init(); err != nil {
		panic(err)
	}
	if err = ipfs.Start(); err != nil {
		panic(err)
	}

	runtime.GlobalRuntime = runtime.NewSuperclusterRuntime(ipfs, db)

	s, err := store.NewIPFSStore()
	if err != nil {
		panic("Cannot create store: " + err.Error())
	}

	// load folders/cluster-service configs from disk
	runtime.GlobalRuntime.Init()

	router.AddRoutes(runtime.GlobalRuntime, s)
	ui.AddRoutes(runtime.GlobalRuntime)

	// TODO: add version here
	log.Println("Supercluster started!")

	runtime.GlobalRuntime.Run(":3030")
}
