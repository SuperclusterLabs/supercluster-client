package example

import (
	"github.com/ipfs/kubo/plugin"

	// delaystore "github.com/SuperclusterLabs/supercluster-client/delaystore"
	// greeter "github.com/SuperclusterLabs/supercluster-client/greeter"
	supercluster "github.com/SuperclusterLabs/supercluster-client/supercluster"
)

// Plugins is an exported list of plugins that will be loaded by Kubo.
var Plugins = []plugin.Plugin{
	// &delaystore.DelaystorePlugin{},
	// &greeter.GreeterPlugin{},
	&supercluster.SuperclusterPlugin{},
}
