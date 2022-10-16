package api

import (
	"errors"
	"context"
)

var ErrNotFound = errors.New("File does not exist")
var ErrFileExists = errors.New("File already exists")
var ErrRequestUnmarshalled = errors.New("Request could not be unmarshalled")
var ErrCannotCreate = errors.New("File could not be created")
var ErrExistingFileRead = errors.New("Could not read existing file:")

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
