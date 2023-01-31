package router

import (
	"net/http"

	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/util"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// FIXME: This is for development ONLY! We need
	// to set this for local development and not
	// all reqs!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsCh chan map[string]interface{} = make(chan map[string]interface{})

func AddRoutes(r util.SuperclusterRuntime, s store.P2PStore) {
	/** middleware/config **/
	// cors allow all
	// TODO: should we be doing this?
	r.Use(cors.Default())
	// set a lower memory limit for multipart forms (default is 32 MiB)
	// see https://gin-gonic.com/docs/examples/upload-file/single-file/
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	api := r.Group("/api")
	api.GET("/files", func(ctx *gin.Context) { listFiles(ctx, s) })
	api.POST("/files", func(ctx *gin.Context) { createFile(ctx, s) })
	api.DELETE("/files/:name", func(ctx *gin.Context) { deleteFile(ctx, s) })
	api.PUT("/files/:name", func(ctx *gin.Context) { modifyFile(ctx, s) })
	api.GET("/ws", func(ctx *gin.Context) { wshandler(ctx, s) })

	// user API
	api.GET("/user", func(ctx *gin.Context) { getUser(ctx) })
	api.POST("/user", func(ctx *gin.Context) { createUser(ctx) })
	api.PUT("/user", func(ctx *gin.Context) { modifyUser(ctx) })
	api.POST("/user/connectPeer", func(ctx *gin.Context) { connectPeer(ctx, s) })
	api.GET("/user/myAddr", func(ctx *gin.Context) { getAddrs(ctx, s) })
	api.GET("/user/clusters", func(ctx *gin.Context) { getUserClusters(ctx) })
	api.POST("/user/pinFile", func(ctx *gin.Context) { createPin(ctx, s) })

	// cluster API
	api.POST("/cluster", func(ctx *gin.Context) { createCluster(ctx) })
	api.GET("/cluster/files", func(ctx *gin.Context) { listPinnedFiles(ctx, s) })

	api.GET("/cluster/:clusterId", func(ctx *gin.Context) { getCluster(ctx) })
	api.PUT("/cluster/:clusterId", func(ctx *gin.Context) { modifyCluster(ctx) })

	// file API
	api.POST("/cluster/:clusterId", func(ctx *gin.Context) { createFile(ctx, s) })
	api.DELETE("/cluster/:clusterId/:fileCid", func(ctx *gin.Context) { deleteFile(ctx, s) })
}
