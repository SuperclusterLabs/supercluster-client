package db

import (
	"context"
	"log"
	"os"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/util"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type FirebaseDB struct {
	Instance *firebase.App
}

var _ SuperclusterDB = (*FirebaseDB)(nil)

/** User routes **/
func NewFirebaseDB() (SuperclusterDB, error) {
	// initialize firebase
	d, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	opt := option.WithCredentialsFile(d + "/.ipfs/keystore/supercluster-2d071-firebase-adminsdk-8qkm4-6688c64d73.json")
	config := &firebase.Config{
		DatabaseURL: "https://supercluster-2d071-default-rtdb.firebaseio.com/",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Println("Error initializing firebase: ", err.Error())
		return nil, err
	}
	s := &FirebaseDB{Instance: app}

	return s, nil
}

func (d *FirebaseDB) GetUserById(ctx context.Context, uId string) (*model.User, error) {
	client, err := d.Instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	var u model.User
	if err := client.NewRef("users/"+uId).Get(ctx, &u); err != nil {
		return nil, err
	}
	if u.Id == uuid.Nil {
		return nil, util.ErrUserNotFound
	}
	return &u, nil
}

// checks if user has been added to a cluster by eth address
// then, checks if the NFTs a user has makes them eligible to join a cluster
// and adds them to those clusters
func (d *FirebaseDB) GetClustersForUser(ctx context.Context, userId string, nftList []string) ([]*model.Cluster, error) {
	client, err := d.Instance.Database(ctx)
	if err != nil {
		return nil, err
	}

	// Get the user from the User ID
	var u model.User
	if err := client.NewRef("users/"+userId).Get(ctx, &u); err != nil {
		return nil, err
	}
	if u.Id == uuid.Nil {
		return nil, util.ErrUserNotFound
	}

	var cs []string
	if u.Clusters != nil {
		is := u.Clusters
		for _, i := range is {
			cs = append(cs, i)
		}
	}

	// Get cluster information for each cluster
	var uClusters []*model.Cluster
	var NFTClusterIDs []string

	for _, cId := range cs {
		var c model.Cluster
		if err := client.NewRef("clusters/"+cId).Get(ctx, &c); err != nil {
			return nil, err
		}
		uClusters = append(uClusters, &c)
	}

	// check for clusters matching nfts
	ref := client.NewRef("clusters")
	// Construct a query for each value in ns
	queries := make([]*db.Query, len(nftList))
	for i, n := range nftList {
		queries[i] = ref.OrderByChild("nftAddr").EqualTo(n)
	}

	// execute all queries in parallel
	results := make(chan db.QueryNode)
	for _, query := range queries {
		go func(q *db.Query) {
			snapshots, err := q.GetOrdered(ctx)
			if err != nil {
				// Handle error
				log.Println("error when executing NFT query: " + err.Error() +
					", for user " + u.Id.String())
				return
			}
			for _, snapshot := range snapshots {
				results <- snapshot
			}
			close(results)
		}(query)
	}

	// Collect the results from all queries
	for r := range results {
		var c model.Cluster
		err := r.Unmarshal(&c)
		if err != nil {
			return nil, err
		}
		c.Members = append(c.Members, u.Id.String())
		_, err = d.CreateCluster(ctx, c)
		if err != nil {
			log.Println("Inconsistent DB state adding nft holder to cluster! " + err.Error())
		}
		uClusters = append(uClusters, &c)
		NFTClusterIDs = append(NFTClusterIDs, c.Id.String())
	}
	_, err = d.UpdateUserClusters(ctx, u.EthAddr, NFTClusterIDs...)
	if err != nil {
		log.Println("Inconsistent DB state adding cluster to user's list! " + err.Error())
	}

	return uClusters, nil
}

func (d *FirebaseDB) GetUserByEthAddr(ctx context.Context, ethAddr string) (*model.User, error) {
	client, err := d.Instance.Database(ctx)
	if err != nil {
		return nil, err
	}

	ref := client.NewRef("users")
	q := ref.OrderByChild("ethAddr").EqualTo(ethAddr)

	// smartypants at google decided that firebase queries
	// should be returned as {{item1}, {item2}, ...} and
	// not [{item1}, {item2}, ...]
	var result map[string]interface{}
	if err := q.Get(ctx, &result); err != nil {
		return nil, err
	} else {
		if len(result) == 0 {
			return nil, util.ErrUserNotFound
		}
	}

	// TODO: clean this up
	// manually wrangle types from json for now.
	// will these type assertions ever fail if
	// we got to this point?
	var u map[string]interface{}
	for k := range result {
		u, _ = result[k].(map[string]interface{})
	}
	// for k, v := range u {
	// if k == "clusters" {
	// log.Println(v.([]interface{}))
	// }
	// }
	id, _ := uuid.Parse(u["id"].(string))

	var cs []string
	if u["clusters"] != nil {
		is := u["clusters"].([]interface{})
		for _, i := range is {
			cs = append(cs, i.(string))
		}
	}

	user := model.User{
		Id:        id,
		EthAddr:   ethAddr,
		IpfsAddr:  u["ipfsAddr"].(string),
		Clusters:  cs,
		Activated: u["activated"].(string),
	}

	return &user, nil
}

func (d *FirebaseDB) UpdateUser(ctx context.Context, u model.User) (*model.User, error) {
	if u.Id == uuid.Nil {
		u.Id = uuid.New()
	}
	client, err := d.Instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.NewRef("users/"+u.Id.String()).Set(ctx, u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *FirebaseDB) UpdateUserClusters(ctx context.Context, eAddr string, cs ...string) (*model.User, error) {
	u, err := d.GetUserByEthAddr(ctx, eAddr)
	if err != nil {
		return nil, err
	}

	u.Clusters = append(u.Clusters, cs...)

	d.UpdateUser(ctx, *u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

/** Cluster routes **/

func (d *FirebaseDB) GetClusterById(ctx context.Context, cId string) (*model.Cluster, error) {
	client, err := d.Instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	var c model.Cluster
	if err := client.NewRef("clusters/"+cId).Get(ctx, &c); err != nil {
		return nil, err
	}
	if c.Id == uuid.Nil {
		return nil, util.ErrClusterNotFound
	}
	return &c, nil
}

func (d *FirebaseDB) CreateCluster(ctx context.Context, c model.Cluster) (*model.Cluster, error) {
	if c.Id == uuid.Nil {
		c.Id = uuid.New()
	}
	client, err := d.Instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.NewRef("clusters/"+c.Id.String()).Set(ctx, c); err != nil {
		return nil, err
	}
	return &c, nil
}
