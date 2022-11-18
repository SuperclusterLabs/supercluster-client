package supercluster

import (
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `json:"id"`
	// TODO: retrieve NFTs on session start? interval?
	// nfts     []interface{} `json:"nfts"`
	Clusters []string `json:"clusters"`
	EthAddr  string   `json:"ethAddr"`
	IpfsAddr string   `json:"ipfsAddr"`
}
