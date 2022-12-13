package supercluster

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func createUser(ctx *gin.Context) {
	u := &User{}
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}

	uDb, err := db.getUserByEthAddr(ctx, u.EthAddr)
	log.Println(uDb)

	if err != nil && err != ErrUserNotFound {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	} else if err == ErrUserNotFound {
		u.Activated = "true"
		u, err = db.updateUser(ctx, *u)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})

			return
		}
	} else if uDb.Activated == "false" {
		uDb.Activated = "true"
		_, _ = db.updateUser(ctx, *uDb)
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
	u := &User{}
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}
	if u.Activated != "true" && u.Activated != "false" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrNeedActivation.Error(),
		})
		return
	}
	uDB, err := db.getUserByEthAddr(ctx, u.EthAddr)
	if err != nil {
		if err == ErrUserNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: ErrUserNotFound.Error(),
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
	u, err = db.updateUser(ctx, *u)
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
			Error: ErrMissingParam.Error() + "ethAddr",
		})
		return
	}

	u, err := db.getUserByEthAddr(ctx, ethAddr)
	if err != nil {
		if err == ErrUserNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: ErrUserNotFound.Error(),
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
			Error: ErrMissingParam.Error() + "ethAddr",
		})
		return
	}

	uClusters, err := db.getClustersForUser(ctx, userId)

	if err != nil {
		if err == ErrUserNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: ErrUserNotFound.Error(),
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

func connectPeer(ctx *gin.Context) {
	a := &AddrsResponse{}
	if err := ctx.BindJSON(a); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error() + err.Error(),
		})
		return
	}

	var mas []multiaddr.Multiaddr
	for _, addr := range a.Addrs {
		ma, err := multiaddr.NewMultiaddr(addr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: ErrInvalidAddrs.Error() + err.Error() + " : " + addr,
			})
		}
		mas = append(mas, ma)
	}

	ipfs := *getCoreAPIInstance()
	err := ipfs.Swarm().Connect(ctx, peer.AddrInfo{
		ID:    peer.ID(a.ID),
		Addrs: mas,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

func getAddrs(ctx *gin.Context) {
	var s ipfsStore
	info, err := s.GetInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, info)
}
