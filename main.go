package main

import (
	"log"
	"os"

	"github.com/SuperclusterLabs/supercluster-client/db"
	"github.com/SuperclusterLabs/supercluster-client/router"
	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/ui"
	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
)

func init() {
	// ensure requisite bins are found
	dirName := ".supercluster"

	hDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	confDir := hDir + "/" + dirName
	_, err = os.Stat(confDir)
	if err != nil {
		panic("Supercluster dir doesn't exist")
	}
}

func main() {
	// TODO: remove firebase
	var err error
	db.AppDB, err = db.NewFirebaseDB()
	if err != nil {
		panic(err)
	}

	ipfs := util.NewProcessManager("./ipfs", []string{"daemon"})
	if err = ipfs.Start(); err != nil {
		panic(err)
	}

	r := util.SuperclusterRuntime{
		IpfsDaemon: ipfs,
		Engine:     gin.Default(),
	}
	s := store.NewIpfsStore()
	if err != nil {
		panic("Cannot create store: " + err.Error())
	}

	router.AddRoutes(r, s)
	ui.AddRoutes(r)

	// TODO: add version here
	log.Println("Supercluster started!")

	r.Run(":3030")
}
