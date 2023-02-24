package router

import (
	"log"
	"net/http"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/proc"
	"github.com/SuperclusterLabs/supercluster-client/runtime"
	"github.com/SuperclusterLabs/supercluster-client/store"
	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
)

// expects user ethAddr to be passed into `creator`
func createCluster(ctx *gin.Context) {
	c := &model.Cluster{}
	if err := ctx.BindJSON(c); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error(),
		})

		return
	}

	_, err := runtime.GlobalRuntime.AppDB.GetUserByEthAddr(ctx, c.Creator)
	if err == util.ErrUserNotFound {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrUserNotFound.Error(),
		})

		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	c, err = runtime.GlobalRuntime.AppDB.CreateCluster(ctx, *c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	// spin up a pinning service on ipfs cluster
	icp, err := proc.NewHostIPFSClusterProcess(c.Id)
	err = icp.Init()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}
	err = icp.Start()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	// add cluster ptr to process pool in Runtime
	runtime.GlobalRuntime.AddProcess(c.Id.String(), icp)

	// TODO: use xmtp (via websocket?) to "mail" the config to other members

	// add cluster to creator's list of clusters
	_, err = runtime.GlobalRuntime.AppDB.UpdateUserClusters(ctx, c.Creator, c.Id.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, c)
}

func modifyCluster(ctx *gin.Context) {
	c := &model.Cluster{}
	if err := ctx.BindJSON(c); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrRequestUnmarshalled.Error(),
		})
		return
	}

	// start by making sure new users have this cluster registered
	// TODO: more complex rules, the following would break if
	// a member is also an admin due to double-counting
	cDb, err := runtime.GlobalRuntime.AppDB.GetClusterById(ctx, c.Id.String())
	oldUsers := append(cDb.Admins, cDb.Members...)
	newUsers := append(c.Admins, c.Members...)
	var updateUs []*model.User
	var createUs []string

	for _, nUsr := range newUsers {
		in := false
		for _, oUsr := range oldUsers {
			if nUsr == oUsr {
				in = true
			}
		}
		if !in {
			u, err := runtime.GlobalRuntime.AppDB.GetUserByEthAddr(ctx, nUsr)
			if err != nil {
				if err == util.ErrUserNotFound {
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
		_, err = runtime.GlobalRuntime.AppDB.UpdateUser(ctx, *u)
	}

	// create new unactivated users
	for _, addr := range createUs {
		u := model.User{
			EthAddr:   addr,
			Activated: "false",
			Clusters:  []string{c.Id.String()},
		}
		_, err = runtime.GlobalRuntime.AppDB.UpdateUser(ctx, u)
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
	c, err = runtime.GlobalRuntime.AppDB.CreateCluster(ctx, *c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, c)
}

func getCluster(ctx *gin.Context) {
	clusterId := ctx.Param("clusterId")
	if clusterId == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "clusterId",
		})
		return
	}

	c, err := runtime.GlobalRuntime.AppDB.GetClusterById(ctx, clusterId)
	if err != nil {
		if err == util.ErrClusterNotFound {
			ctx.JSON(http.StatusBadRequest, ResponseError{
				Error: util.ErrClusterNotFound.Error(),
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

func listPinnedFiles(ctx *gin.Context, s store.P2PStore) {
	clusterId := ctx.Param("clusterId")
	if clusterId == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "clusterId",
		})
		return
	}
	ps, err := s.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, ps)
}

func listClusterAddrs(ctx *gin.Context, s store.P2PStore) {
	clusterId := ctx.Param("clusterId")
	if clusterId == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: util.ErrMissingParam.Error() + "clusterId",
		})
		return
	}
	info, err := s.GetInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, info)
}
