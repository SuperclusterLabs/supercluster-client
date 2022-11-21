package supercluster

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
)

type DB struct {
	instance *firebase.App
}

var db DB

/** User routes **/

// TODO: figure out consistent way of taking/returning pointers

func (d *DB) getUserById(ctx context.Context, uId string) (*User, error) {
	client, err := d.instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	var u User
	if err := client.NewRef("users/"+uId).Get(ctx, &u); err != nil {
		return nil, err
	}
	if u.Id == uuid.Nil {
		return nil, ErrUserNotFound
	}
	return &u, nil
}

func (d *DB) getUserByEthAddr(ctx context.Context, ethAddr string) (*User, error) {
	client, err := d.instance.Database(ctx)
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
			return nil, ErrUserNotFound
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

	log.Println(cs)
	log.Println(u)
	user := User{
		Id:        id,
		EthAddr:   ethAddr,
		IpfsAddr:  u["ipfsAddr"].(string),
		Clusters:  cs,
		Activated: u["activated"].(string),
	}

	return &user, nil
}

func (d *DB) updateUser(ctx context.Context, u User) (*User, error) {
	if u.Id == uuid.Nil {
		u.Id = uuid.New()
	}
	client, err := d.instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.NewRef("users/"+u.Id.String()).Set(ctx, u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *DB) updateUserClusters(ctx context.Context, eAddr string, cs ...string) (*User, error) {
	u, err := d.getUserByEthAddr(ctx, eAddr)
	if err != nil {
		return nil, err
	}

	u.Clusters = append(u.Clusters, cs...)

	d.updateUser(ctx, *u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

/** Cluster routes **/

func (d *DB) getClusterById(ctx context.Context, cId string) (*Cluster, error) {
	client, err := d.instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	var c Cluster
	if err := client.NewRef("clusters/"+cId).Get(ctx, &c); err != nil {
		return nil, err
	}
	if c.Id == uuid.Nil {
		return nil, ErrClusterNotFound
	}
	return &c, nil
}

func (d *DB) createCluster(ctx context.Context, c Cluster) (*Cluster, error) {
	if c.Id == uuid.Nil {
		c.Id = uuid.New()
	}
	client, err := d.instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.NewRef("clusters/"+c.Id.String()).Set(ctx, c); err != nil {
		return nil, err
	}
	return &c, nil
}
