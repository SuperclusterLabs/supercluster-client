package supercluster

import (
	"errors"
)

var ErrNotFound = errors.New("File does not exist")
var ErrFileExists = errors.New("File already exists")
var ErrUserExists = errors.New("User already exists")
var ErrUserNotFound = errors.New("User not found")
var ErrNeedActivation = errors.New("User activation status needs to be specified")
var ErrClusterNotFound = errors.New("Cluster not found")
var ErrMissingParam = errors.New("Missing param: ")
var ErrInvalidAddrs = errors.New("Invalid multiaddrs provided")
var ErrRequestUnmarshalled = errors.New("Request could not be unmarshalled")
var ErrCannotCreate = errors.New("File could not be created")
var ErrExistingFileRead = errors.New("Could not read existing file")
