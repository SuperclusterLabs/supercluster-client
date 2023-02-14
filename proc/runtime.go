package proc

import "github.com/gin-gonic/gin"

type SuperclusterRuntime struct {
	IPFSDaemon ManagedProcess
	*gin.Engine
}
