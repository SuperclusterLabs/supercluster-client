package proc

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SuperclusterRuntime struct {
	IPFSDaemon ManagedProcess
	// TODO: this should be abstracted to a pool
	clusterCtlProcs map[uuid.UUID]*IPFSClusterProcess

	*gin.Engine
}

func NewSuperclusterRuntime(ipfs ManagedProcess) SuperclusterRuntime {
	return SuperclusterRuntime{
		IPFSDaemon:      ipfs,
		Engine:          gin.Default(),
		clusterCtlProcs: make(map[uuid.UUID]*IPFSClusterProcess),
	}
}

var GlobalRuntime SuperclusterRuntime

// TODO: should this check if the process is running?
func (r *SuperclusterRuntime) AddProcess(id uuid.UUID, p *IPFSClusterProcess) error {
	// TODO: abstract
	if len(r.clusterCtlProcs) == 10 {
		return errors.New("Max clusters reached")
	}
	r.clusterCtlProcs[id] = p
	return nil
}

// TODO: should this check if the process is ended?
func (r *SuperclusterRuntime) RemoveProcess(id uuid.UUID, p *IPFSClusterProcess) {
	delete(r.clusterCtlProcs, id)
}
