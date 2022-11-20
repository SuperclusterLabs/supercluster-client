package supercluster

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
)

type DB struct {
	instance *firebase.App
}

var db DB

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

	id, _ := uuid.Parse(u["id"].(string))

	var cs []string
	if u["Clusters"] != nil {
		cs = u["clusters"].([]string)
	}

	user := User{
		Id:       id,
		EthAddr:  ethAddr,
		IpfsAddr: u["ipfsAddr"].(string),
		Clusters: cs,
	}

	return &user, nil
}

func (d *DB) createUser(ctx context.Context, u User) (*User, error) {
	u.Id = uuid.New()
	client, err := d.instance.Database(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.NewRef("users/"+u.Id.String()).Set(ctx, u); err != nil {
		return nil, err
	}
	return &u, nil
}
