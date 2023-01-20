package model

type File struct {
	Name      string `json:"name"`
	Cid       string `json:"cid"`
	Creator   string `json:"creator"`
	CreatedAt int64  `json:"createdAt"`
	// TODO: whitelist of users that can access
}
