package api

import (
	"github.com/gin-gonic/gin"

	"github.com/SuperclusterLabs/supercluster-client/ui"
)

func Start() {
	// TODO: figure out how IPFS context should work here
	store, _, err := newIpfsStore()
	if err != nil {
		panic("Cannot create store: " + err.Error())
	}
	router := gin.Default()

	addRoutes(router, store)
	ui.AddRoutes(router)

	router.Run(":4000")
}
