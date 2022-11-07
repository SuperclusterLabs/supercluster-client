package supercluster

import (
	"errors"
)

var ErrNotFound = errors.New("File does not exist")
var ErrFileExists = errors.New("File already exists")
var ErrRequestUnmarshalled = errors.New("Request could not be unmarshalled")
var ErrCannotCreate = errors.New("File could not be created")
var ErrExistingFileRead = errors.New("Could not read existing file:")
