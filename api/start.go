package api

import (
	"github.com/gin-gonic/gin"

	"github.com/SuperclusterLabs/supercluster-client/ui"
)

func Start() {
	store := newStore()
	router := gin.Default()

	addRoutes(router, store)
	ui.AddRoutes(router)

	router.Run(":4000")
}
