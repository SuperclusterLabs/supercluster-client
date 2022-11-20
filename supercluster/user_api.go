package supercluster

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createUser(ctx *gin.Context) {
	user := &User{}
	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}
	client, err := db.instance.Database(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	ref := client.NewRef("users")

	// create if user doesn't already exist
	q := ref.OrderByChild("ethAddr").EqualTo(user.EthAddr)

	// smartypants at google decided that firebase queries
	// should be returned as {{item1}, {item2}, ...} and
	// not [{item1}, {item2}, ...]
	var result map[string]interface{}
	if err := q.Get(ctx, &result); err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	} else {
		if len(result) != 0 {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: ErrUserExists.Error(),
			})

			return
		}
		user.Id = uuid.New()
		log.Print("creating user: ", user.Id.String())
		if err := client.NewRef("users/"+user.Id.String()).Set(ctx, user); err != nil {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})

			return
		}
	}

	ctx.JSON(http.StatusOK, user)
}

func modifyUser(ctx *gin.Context) {

	payload := &ModifyPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ModifyResponse{})
}

func getUser(ctx *gin.Context) {
	user := &User{}
	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}
	client, err := db.instance.Database(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}
	ref := client.NewRef("users")

	q := ref.OrderByChild("ethAddr").EqualTo(user.EthAddr)
	var result User
	if err := q.Get(ctx, &result); err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	} else {
		if result.EthAddr == "" {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: ErrUserNotFound.Error(),
			})

			return
		}
	}

	ctx.JSON(http.StatusOK, result)
}
