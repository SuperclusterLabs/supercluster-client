package router

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/util"
	"github.com/gin-gonic/gin"
)

func wshandler(ctx *gin.Context, _ store.P2PStore) {
	conn, err := wsupgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: ", err)
		return
	}

	go func() {
		log.Println("wss channel opening")

		for m := range wsCh {
			conn.WriteJSON(m)
		}
	}()

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func createFile(ctx *gin.Context, s store.P2PStore) {
	log.Println(ctx.Request)
	f, h, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			util.ResponseError{Error: err.Error()})
		return
	}

	// read file into bytes from request
	log.Println(h.Filename)
	defer f.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	file, err := s.Create(ctx, h.Filename, buf.Bytes())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ResponseError{
			Error: util.ErrCannotCreate.Error(),
		})
		return
	}

	// let frontend know to transmit xmtp msg
	// info, err := s.GetInfo(ctx)
	// n := make(map[string]interface{})
	// n["cid"] = *&file.ID
	// n["action"] = "pin"
	// n["ipfsAddr"] = info.ID
	// n["addrs"] = info.Addrs

	// wsCh <- n

	ctx.JSON(http.StatusOK, util.CreateResponse{
		File: *file,
	})
}

func deleteFile(ctx *gin.Context, s store.P2PStore) {
	cid := ctx.Param("fileCid")

	err := s.Delete(ctx, cid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ResponseError{
			Error: err.Error(),
		})
		return
	}
	// let frontend know to transmit xmtp msg
	// n := make(map[string]interface{})
	// n["cid"] = cid
	// n["action"] = "unpin"
	// wsCh <- n

	ctx.Status(http.StatusOK)
}

func modifyFile(ctx *gin.Context, s store.P2PStore) {
	name := ctx.Param("name")

	payload := &util.ModifyPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ResponseError{
			Error: util.ErrRequestUnmarshalled.Error(),
		})
		return
	}

	f, err := s.Modify(ctx, name, payload.Contents)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ResponseError{
			Error: util.ErrNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, util.ModifyResponse{
		File: *f,
	})
}

func listFiles(ctx *gin.Context, s store.P2PStore) {
	fs, err := s.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ResponseError{
			Error: util.ErrExistingFileRead.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, util.ListResponse{
		Files: fs,
	})
}

func createPin(ctx *gin.Context, s store.P2PStore) {
	p := &util.PinRequest{}
	if err := ctx.BindJSON(p); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ResponseError{
			Error: util.ErrRequestUnmarshalled.Error() + err.Error(),
		})
		return
	}
	err := s.PinFile(ctx, p.Cid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ResponseError{
			Error: util.ErrExistingFileRead.Error() + err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}