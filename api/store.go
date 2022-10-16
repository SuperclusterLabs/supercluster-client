package api

import (
	"context"
)

type Store interface {
	Create(ctx context.Context, name, contents string) (*file, error)
	Modify(ctx context.Context, name, contents string) (*file, error)
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]file, error)
}

type file struct {
	// Name of the file
	Name string `json:"name"`
	// The contents of the file
	Contents string `json:"contents"`
	// Unix timestamp of creation
	CreatedAt int64 `json:"createdAt"`
}
