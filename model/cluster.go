package model

import "github.com/google/uuid"

type Cluster struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json"description"`
	NftAddr     string    `json:"nftAddr"`
	Files       []File    `json:"files"`
	Admins      []string  `json:"admins"`
	Creator     string    `json:"creator" binding:"required"`
	Members     []string  `json:"members"`
}
