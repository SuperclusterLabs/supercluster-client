package supercluster

import (
	"context"
	"log"

	cors "github.com/gin-contrib/cors"
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

	/** middleware/config **/
	// cors allow all
	// TODO: should we be doing this?
	r.Use(cors.Default())
	// set a lower memory limit for multipart forms (default is 32 MiB)
	// see https://gin-gonic.com/docs/examples/upload-file/single-file/
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	api := r.Group("/api")
	api.GET("/files", func(ctx *gin.Context) { listFiles(ctx, store) })
	api.POST("/files", func(ctx *gin.Context) { createFile(ctx, store) })
	api.DELETE("/files/:name", func(ctx *gin.Context) { deleteFile(ctx, store) })
	api.PUT("/files/:name", func(ctx *gin.Context) { modifyFile(ctx, store) })
	api.GET("/ws", func(ctx *gin.Context) { wshandler(ctx, store) })

	// user API
	api.GET("/user", func(ctx *gin.Context) { getUser(ctx) })
	api.POST("/user", func(ctx *gin.Context) { createUser(ctx) })
	api.PUT("/user", func(ctx *gin.Context) { modifyUser(ctx) })

	// cluster API
	api.POST("/cluster", func(ctx *gin.Context) { createCluster(ctx) })
	api.GET("/cluster/files", func(ctx *gin.Context) { listPinnedFiles(ctx) })

	api.GET("/cluster/:clusterId", func(ctx *gin.Context) { getCluster(ctx) })
	api.PUT("/cluster/:clusterId", func(ctx *gin.Context) { modifyCluster(ctx) })

	// file API
	api.POST("/cluster/:clusterId", func(ctx *gin.Context) { createFile(ctx, store) })
	api.DELETE("/cluster/:clusterId/:fileCid", func(ctx *gin.Context) { deleteFile(ctx, store) })
}
