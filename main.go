package main

import (
	"log"
	"os"

	"github.com/SuperclusterLabs/supercluster-client/db"
	"github.com/SuperclusterLabs/supercluster-client/proc"
	"github.com/SuperclusterLabs/supercluster-client/router"
	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/ui"
	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
)

func main() {
	confDir := util.GetConfDir()

	_, err := os.Stat(confDir)
	if err != nil {
		panic("Supercluster dir doesn't exist")
	}

	// TODO: remove firebase
	db.AppDB, err = db.NewFirebaseDB()
	if err != nil {
		panic(err)
	}

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

	r := proc.SuperclusterRuntime{
		IPFSDaemon: ipfs,
		Engine:     gin.Default(),
	}
	s, err := store.NewIPFSStore()
	if err != nil {
		panic("Cannot create store: " + err.Error())
	}

	router.AddRoutes(r, s)
	ui.AddRoutes(r)

	// TODO: add version here
	log.Println("Supercluster started!")

	r.Run(":3030")
}
