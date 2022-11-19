package plugin

import (
	"github.com/ipfs/kubo/plugin"

	supercluster "github.com/SuperclusterLabs/supercluster-client/supercluster"
)

// Plugins is an exported list of plugins that will be loaded by Kubo.
var Plugins = []plugin.Plugin{
	&supercluster.SuperclusterPlugin{},
}
