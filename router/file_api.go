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

func getFile(ctx *gin.Context, s store.P2PStore) {
	cid := ctx.Param("fileCid")
	if cid == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "cid",
		})
		return
	}

	bs, _, err := s.Get(ctx, cid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	ctx.Data(http.StatusOK, "application/octet-stream", bs)
}

func createFile(ctx *gin.Context, s store.P2PStore) {
	clusterId := ctx.Param("clusterId")
	if clusterId == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "clusterId",
		})
		return
	}

	f, h, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			ResponseError{Error: err.Error()})
		return
	}

	// read file into bytes from request
	defer f.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	file, err := s.Create(ctx, h.Filename, buf.Bytes())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: util.ErrCannotCreate.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, CreateResponse{
		File: *file,
	})
}

func deleteFile(ctx *gin.Context, s store.P2PStore) {
	cid := ctx.Param("fileCid")

	err := s.Delete(ctx, cid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
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

func deleteAll(ctx *gin.Context, s store.P2PStore) {
	err := s.DeleteAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

func modifyFile(ctx *gin.Context, s store.P2PStore) {
	name := ctx.Param("name")

	payload := &ModifyPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error(),
		})
		return
	}

	f, err := s.Modify(ctx, name, payload.Contents)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ModifyResponse{
		File: *f,
	})
}

func listFiles(ctx *gin.Context, s store.P2PStore) {
	fs, err := s.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: util.ErrExistingFileRead.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, ListResponse{
		Files: fs,
	})
}

func createPin(ctx *gin.Context, s store.P2PStore) {
	p := &PinRequest{}
	if err := ctx.BindJSON(p); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error() + err.Error(),
		})
		return
	}
	err := s.PinFile(ctx, p.Cid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: util.ErrExistingFileRead.Error() + err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
