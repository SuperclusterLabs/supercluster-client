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
		fmt.Println("Failed to set websocket upgrade: %+v", err)
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

func createTodo(ctx *gin.Context, s Store) {
	payload := &CreatePayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, CreateResponse{
		Todo: s.Create(payload.Description),
	})
}

func deleteTodo(ctx *gin.Context, s Store) {
	id := ctx.Param("id")
	s.Delete(id)
}

func checkTodo(ctx *gin.Context, s Store) {
	id := ctx.Param("id")

	payload := &CheckPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	t, err := s.Check(id, payload.Completed)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, CheckResponse{
		Todo: t,
	})
}

func listTodos(ctx *gin.Context, s Store) {
	ctx.JSON(http.StatusOK, ListResponse{
		Todos: s.List(),
	})
}
