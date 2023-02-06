package router

import "github.com/SuperclusterLabs/supercluster-client/model"

type ListResponse struct {
	Files []model.File `json:"files"`
}

type CreatePayload struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

type CreateResponse struct {
	File model.File `json:"file"`
}

type ModifyPayload struct {
	Contents string `json:"contents"`
}

type ModifyResponse struct {
	File model.File `json:"file"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type AddrsResponse struct {
	ID    string   `json:"id"`
	Addrs []string `json:"addrs"`
}

type PinRequest struct {
	Cid string `json:"cid"`
}
