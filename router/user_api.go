package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/runtime"
	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
)

func createUser(ctx *gin.Context) {
	u := &model.User{}
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error(),
		})
		return
	}

	uDb, err := runtime.GlobalRuntime.AppDB.GetUserByEthAddr(ctx, u.EthAddr)
	log.Println(uDb)

	if err != nil && err != util.ErrUserNotFound {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	} else if err == util.ErrUserNotFound {
		u.Activated = "true"
		u, err = runtime.GlobalRuntime.AppDB.UpdateUser(ctx, *u)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})

			return
		}
	} else if uDb.Activated == "false" {
		uDb.Activated = "true"
		_, _ = runtime.GlobalRuntime.AppDB.UpdateUser(ctx, *uDb)
		u = uDb
	} else {
		// ctx.JSON(http.StatusBadRequest, ResponseError{
		// Error: ErrUserExists.Error(),
		// })

		u = uDb
	}

	ctx.JSON(http.StatusOK, u)
}

func modifyUser(ctx *gin.Context) {
	u := &model.User{}
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error(),
		})
		return
	}
	if u.Activated != "true" && u.Activated != "false" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrNeedActivation.Error(),
		})
		return
	}
	uDB, err := runtime.GlobalRuntime.AppDB.GetUserByEthAddr(ctx, u.EthAddr)
	if err != nil {
		if err == util.ErrUserNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: util.ErrUserNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})
		}
		return
	}
	// trust that any change to the struct is intentional
	// TODO: as there are more fields, will this assumption still
	// hold? Else we'll need to prepopulate the user struct
	// with a seemingly unnecessary extra db call so that
	// this func doesn't scrub data
	u.Id = uDB.Id
	u, err = runtime.GlobalRuntime.AppDB.UpdateUser(ctx, *u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, u)
}

func getUser(ctx *gin.Context) {
	ethAddr := ctx.Query("ethAddr")
	if ethAddr == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "ethAddr",
		})
		return
	}

	u, err := runtime.GlobalRuntime.AppDB.GetUserByEthAddr(ctx, ethAddr)
	if err != nil {
		if err == util.ErrUserNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: util.ErrUserNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func getUserClusters(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "ethAddr",
		})
		return
	}

	uClusters, err := runtime.GlobalRuntime.AppDB.GetClustersForUser(ctx, userId)

	if err != nil {
		if err == util.ErrUserNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: util.ErrUserNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, uClusters)
}

func connectPeer(ctx *gin.Context, s store.P2PStore) {
	a := &store.P2PNodeInfo{}
	if err := ctx.BindJSON(a); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error() + err.Error(),
		})
		return
	}

	err := s.ConnectPeer(ctx, a.Addresses...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

func getAddrs(ctx *gin.Context, s store.P2PStore) {
	resp, err := http.Post("http://localhost:5001/api/v0/id", "application/json", nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	var ar store.P2PNodeInfo
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ar)
}
