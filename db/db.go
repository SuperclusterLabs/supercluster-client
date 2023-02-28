package db

import (
	"context"

	"github.com/SuperclusterLabs/supercluster-client/model"
)

type SuperclusterDB interface {

	/** User routes **/

	GetUserById(ctx context.Context, uId string) (*model.User, error)
	GetClustersForUser(ctx context.Context, userId string, nftList []string) ([]*model.Cluster, error)
	GetUserByEthAddr(ctx context.Context, ethAddr string) (*model.User, error)
	UpdateUser(ctx context.Context, u model.User) (*model.User, error)
	UpdateUserClusters(ctx context.Context, eAddr string, cs ...string) (*model.User, error)

	/** Cluster routes **/

	GetClusterById(ctx context.Context, cId string) (*model.Cluster, error)
	CreateCluster(ctx context.Context, c model.Cluster) (*model.Cluster, error)
}
