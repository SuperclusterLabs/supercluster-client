package supercluster

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func wshandler(ctx *gin.Context, _ ipfsStore) {
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

func createFile(ctx *gin.Context, s ipfsStore) {
	log.Println(ctx.Request)
	f, h, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			ResponseError{Error: err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: ErrCannotCreate.Error(),
		})
		return
	}

	// let frontend know to transmit xmtp msg
	info, err := s.GetInfo(ctx)
	n := make(map[string]interface{})
	n["cid"] = *&file.ID
	n["action"] = "pin"
	n["ipfsAddr"] = info.ID
	n["addrs"] = info.Addrs

	wsCh <- n

	ctx.JSON(http.StatusOK, CreateResponse{
		File: *file,
	})
}

func deleteFile(ctx *gin.Context, s ipfsStore) {
	cid := ctx.Param("fileCid")

	err := s.Delete(ctx, cid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
		return
	}
	// let frontend know to transmit xmtp msg
	n := make(map[string]interface{})
	n["cid"] = cid
	n["action"] = "unpin"
	wsCh <- n
	ctx.Status(http.StatusOK)
}

func modifyFile(ctx *gin.Context, s ipfsStore) {
	name := ctx.Param("name")

	payload := &ModifyPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}

	f, err := s.Modify(ctx, name, payload.Contents)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ModifyResponse{
		File: *f,
	})
}

func listFiles(ctx *gin.Context, s ipfsStore) {
	fs, err := s.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: ErrExistingFileRead.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, ListResponse{
		Files: fs,
	})
}

func createPin(ctx *gin.Context, s ipfsStore) {
	c := ctx.GetString("cid")
	if c == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrMissingParam.Error() + "cid",
		})
	}
	err := s.pinFile(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: ErrExistingFileRead.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
