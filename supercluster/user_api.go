package supercluster

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(ctx *gin.Context) {
	u := &User{}
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}
	_, err := db.getUserByEthAddr(ctx, u.EthAddr)
	if err != nil && err != ErrUserNotFound {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	} else if err == ErrUserNotFound {
		log.Print("creating user: ", u.Id.String())

		u, err = db.createUser(ctx, *u)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})

			return
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: ErrUserExists.Error(),
		})

		return
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
	u, err = db.createUser(ctx, *u)
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
