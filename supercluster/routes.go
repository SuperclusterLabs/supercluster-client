package supercluster

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine, store Store) {
	c := *getCoreAPIInstance()
	var ctx context.Context
	n, err := c.Key().Self(ctx)
	log.Println("Still alive? ", n.ID().String())
	if err != nil {
		panic(err)
	}
	api := r.Group("/api")
	api.GET("/files", func(ctx *gin.Context) { listFiles(ctx, store) })
	api.POST("/files", func(ctx *gin.Context) { createFile(ctx, store) })
	api.DELETE("/files/:name", func(ctx *gin.Context) { deleteFile(ctx, store) })
	api.PUT("/files/:name", func(ctx *gin.Context) { modifyFile(ctx, store) })
	api.GET("/ws", func(ctx *gin.Context) { wshandler(ctx, store) })
}
