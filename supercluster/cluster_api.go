package supercluster

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// expects user ethAddr to be passed into `creator`
func createCluster(ctx *gin.Context) {
	c := &Cluster{}
	if err := ctx.BindJSON(c); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})

		return
	}

	u, err := db.getUserByEthAddr(ctx, c.Creator)
	log.Println(u)
	if err == ErrUserNotFound {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrUserNotFound.Error(),
		})

		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	c, err = db.createCluster(ctx, *c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	// add cluster to creator's list of clusters
	_, err = db.updateUserClusters(ctx, c.Creator, c.Id.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, c)
}

func modifyCluster(ctx *gin.Context) {
	c := &Cluster{}
	if err := ctx.BindJSON(c); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrRequestUnmarshalled.Error(),
		})
		return
	}

	// start by making sure new users have this cluster registered
	// TODO: more complex rules, the following would break if
	// a member is also an admin due to double-counting
	cDb, err := db.getClusterById(ctx, c.Id.String())
	oldUsers := append(cDb.Admins, cDb.Members...)
	newUsers := append(c.Admins, c.Members...)
	var updateUs []*User
	var createUs []string

	for _, nUsr := range newUsers {
		in := false
		for _, oUsr := range oldUsers {
			if nUsr == oUsr {
				in = true
			}
		}
		if !in {
			u, err := db.getUserByEthAddr(ctx, nUsr)
			if err != nil {
				if err == ErrUserNotFound {
					// add unregistered user to create list
					createUs = append(createUs, nUsr)
				} else {
					// on failure discard updates
					ctx.JSON(http.StatusInternalServerError, ResponseError{
						Error: err.Error(),
					})
					return
				}
			} else {
				// add registered user to update list
				u.Clusters = append(u.Clusters, c.Id.String())
				updateUs = append(updateUs, u)
			}
		}
	}

	// update existing users
	for _, u := range updateUs {
		_, err = db.updateUser(ctx, *u)
	}

	// create new unactivated users
	for _, addr := range createUs {
		u := User{
			EthAddr:   addr,
			Activated: "false",
			Clusters:  []string{c.Id.String()},
		}
		_, err = db.updateUser(ctx, u)
	}

	if err != nil {
		// FIXME: when can this happen? how to avoid?
		log.Println("User update failed while modifying cluster: ", c.Id.String())
		log.Println("Cluster: ")
		log.Println(c)
		log.Println("Error: ")
		log.Println(err.Error())
	}

	// finally, update cluster
	c, err = db.createCluster(ctx, *c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, c)
}

func getCluster(ctx *gin.Context) {
	log.Println("get cluster")
	clusterId := ctx.Param("clusterId")
	if clusterId == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: ErrMissingParam.Error() + "ethAddr",
		})
		return
	}

	c, err := db.getClusterById(ctx, clusterId)
	if err != nil {
		if err == ErrClusterNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: ErrClusterNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ResponseError{
				Error: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func listPinnedFiles(ctx *gin.Context) {
	ipfs := *getCoreAPIInstance()
	var ps []string
	pch, err := ipfs.Pin().Ls(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}
	for p := range pch {
		ps = append(ps, p.Path().Cid().String())
	}
	ctx.JSON(http.StatusOK, ps)
}
