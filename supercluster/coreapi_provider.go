package supercluster

import (
	"sync"

	coreiface "github.com/ipfs/interface-go-ipfs-core"
)

var once sync.Once
var instance *coreiface.CoreAPI

func setCoreAPIInstance(c *coreiface.CoreAPI) {
	// TODO: is `once` really necessary?
	once.Do(func() {
		instance = c
	})
}

func getCoreAPIInstance() *coreiface.CoreAPI {
	return instance
}
