package api

import (
	"fmt"
	"net/http"

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

func wshandler(ctx *gin.Context, _ Store) {
	conn, err := wsupgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func createFile(ctx *gin.Context, s Store) {
	payload := &CreatePayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := s.Create(payload.Name, payload.Contents)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, CreateResponse{
		File: *file,
	})
}

func deleteFile(ctx *gin.Context, s Store) {
	name := ctx.Param("name")

	err := s.Delete(name)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.Status(http.StatusOK)
}

func modifyFile(ctx *gin.Context, s Store) {
	name := ctx.Param("name")

	payload := &ModifyPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f, err := s.Modify(name, payload.Contents)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, ModifyResponse{
		File: *f,
	})
}

func listFiles(ctx *gin.Context, s Store) {
	ctx.JSON(http.StatusOK, ListResponse{
		Files: s.List(),
	})
}
