package model

import (
	"github.com/google/uuid"
)

type User struct {
	// TODO: retrieve NFTs on session start? interval?
	// nfts     []interface{} `json:"nfts"`
	Id       uuid.UUID `json:"id"`
	Clusters []string  `json:"clusters"`
	EthAddr  string    `json:"ethAddr" binding:"required"`
	IpfsAddr string    `json:"ipfsAddr"`
	// this would be a bool but gin doesn't unmarshal it properly
	Activated string `json:"activated"`
}
