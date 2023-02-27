package store

import (
	"github.com/SuperclusterLabs/supercluster-client/model"

	"github.com/gin-gonic/gin"
)

type P2PStore interface {
	Get(ctx *gin.Context, cid string) ([]byte, *model.File, error)
	Create(ctx *gin.Context, name string, contents []byte) (*model.File, error)
	Modify(ctx *gin.Context, name, contents string) (*model.File, error)
	Delete(ctx *gin.Context, name string) error
	DeleteAll(ctx *gin.Context) error
	List(ctx *gin.Context) ([]model.File, error)
	GetInfo(ctx *gin.Context) (*P2PNodeInfo, error)
	PinFile(ctx *gin.Context, c string) error
	ConnectPeer(ctx *gin.Context, addr ...string) error
}

type P2PNodeInfo struct {
	ID              string
	PublicKey       string
	Addresses       []string
	AgentVersion    string
	ProtocolVersion string
}
