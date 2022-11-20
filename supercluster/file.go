package supercluster

import (
	"github.com/ipfs/go-cid"
)

type File struct {
	Name    string  `json"name"`
	Cid     cid.Cid `json:"cid"`
	Creator string  `json:"creator:`
	// TODO: whitelist of users that can access
}
