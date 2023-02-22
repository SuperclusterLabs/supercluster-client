package runtime

import (
	"errors"

	"github.com/SuperclusterLabs/supercluster-client/db"
	"github.com/SuperclusterLabs/supercluster-client/proc"

	"github.com/gin-gonic/gin"
)

type SuperclusterRuntime struct {
	// IPFS daemon process
	IPFSDaemon proc.ManagedProcess

	// pool of cluster pinning services
	// TODO: process pool logic
	clusterCtlProcs map[string]*proc.IPFSClusterProcess
	// metadata DB for users and clusters
	AppDB db.SuperclusterDB

	*gin.Engine
}

func NewSuperclusterRuntime(ipfs proc.ManagedProcess, db db.SuperclusterDB) SuperclusterRuntime {
	return SuperclusterRuntime{
		IPFSDaemon:      ipfs,
		clusterCtlProcs: make(map[string]*proc.IPFSClusterProcess),
		AppDB:           db,
		Engine:          gin.Default(),
	}
}

var GlobalRuntime SuperclusterRuntime

// TODO: should this check if the process is running?
func (r *SuperclusterRuntime) AddProcess(clusterId string, p *proc.IPFSClusterProcess) error {
	// TODO: abstract
	if len(r.clusterCtlProcs) == 10 {
		return errors.New("Max clusters reached")
	}
	r.clusterCtlProcs[clusterId] = p
	return nil
}

func (r *SuperclusterRuntime) GetProcess(clusterId string) (*proc.IPFSClusterProcess, error) {
	if p, ok := r.clusterCtlProcs[clusterId]; ok {
		return p, nil
	}
	return nil, errors.New("No cluster running with that ID")
}

// TODO: should this check if the process is ended?
func (r *SuperclusterRuntime) RemoveProcess(clusterId string, p *proc.IPFSClusterProcess) {
	delete(r.clusterCtlProcs, clusterId)
}
