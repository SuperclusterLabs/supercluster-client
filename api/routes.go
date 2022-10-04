package api

import (
	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine, store Store) {
	api := r.Group("/api")
	api.GET("/files", func(ctx *gin.Context) { listFiles(ctx, store) })
	api.POST("/files", func(ctx *gin.Context) { createFile(ctx, store) })
	api.DELETE("/files/:name", func(ctx *gin.Context) { deleteFile(ctx, store) })
	api.PUT("/files/:name", func(ctx *gin.Context) { modifyFile(ctx, store) })
	api.GET("/ws", func(ctx *gin.Context) { wshandler(ctx, store) })
}
