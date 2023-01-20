package store

import (
	"context"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
)

type P2PStore interface {
	Create(ctx *gin.Context, name string, contents []byte) (*model.File, error)
	Modify(ctx context.Context, name, contents string) (*model.File, error)
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]model.File, error)
	GetInfo(ctx context.Context) (*util.AddrsResponse, error)
	PinFile(ctx *gin.Context, c string) error
	ConnectPeer(ctx *gin.Context, addr ...string) error
}
