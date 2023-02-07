package proc

import "github.com/gin-gonic/gin"

type SuperclusterRuntime struct {
	IpfsDaemon *ProcessManager
	*gin.Engine
}
