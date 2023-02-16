package proc

import "github.com/gin-gonic/gin"

type SuperclusterRuntime struct {
	IPFSDaemon ManagedProcess
	// TODO: this should be abstracted to a pool
	ClusterCtlProc []*IPFSClusterProcess

	*gin.Engine
}
