package supercluster

import "github.com/google/uuid"

type Cluster struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json"description"`
	NftAddr     string        `json:"nftAddr"`
	Files       []interface{} `json:"files"`
	Admins      []User        `json:"admins"`
	Creator     string        `json:"creator"`
	Members     []User        `json:"members"`
}
