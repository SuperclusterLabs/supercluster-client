package db

import (
	"context"

	"github.com/SuperclusterLabs/supercluster-client/model"
)

// Test DB for offline deveopment
type NullDB struct{}

var _ SuperclusterDB = (*NullDB)(nil)

/** User routes **/

func (n *NullDB) GetUserById(ctx context.Context, uId string) (*model.User, error) {
	return nil, nil
}
func (n *NullDB) GetClustersForUser(ctx context.Context, userId string) ([]*model.Cluster, error) {
	return []*model.Cluster{}, nil
}
func (n *NullDB) GetUserByEthAddr(ctx context.Context, ethAddr string) (*model.User, error) {
	return &model.User{}, nil
}
func (n *NullDB) UpdateUser(ctx context.Context, u model.User) (*model.User, error) {
	return &model.User{}, nil
}
func (n *NullDB) UpdateUserClusters(ctx context.Context, eAddr string, cs ...string) (*model.User, error) {
	return &model.User{}, nil
}

/** Cluster routes **/

func (n *NullDB) GetClusterById(ctx context.Context, cId string) (*model.Cluster, error) {
	return &model.Cluster{}, nil
}

func (n *NullDB) CreateCluster(ctx context.Context, c model.Cluster) (*model.Cluster, error) {
	return &model.Cluster{}, nil
}
