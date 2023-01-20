package main

import (
	"log"

	"github.com/SuperclusterLabs/supercluster-client/db"
	"github.com/SuperclusterLabs/supercluster-client/router"
	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/ui"

	"github.com/gin-gonic/gin"
)

func main() {

	// TODO: remove firebase
	var err error
	db.AppDB, err = db.NewFirebaseDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	s := store.NewIpfsStore()
	if err != nil {
		panic("Cannot create store: " + err.Error())
	}

	router.AddRoutes(r, s)
	ui.AddRoutes(r)

	// TODO: add version here
	log.Println("Supercluster started!")

	r.Run(":3000")

}
