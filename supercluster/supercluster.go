package supercluster

import (
	"fmt"

	"github.com/SuperclusterLabs/supercluster-client/ui"
	"github.com/gin-gonic/gin"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	plugin "github.com/ipfs/kubo/plugin"
)

type SuperclusterPlugin struct{}

var _ plugin.PluginDaemon = (*SuperclusterPlugin)(nil)

// Name returns the plugin's name, satisfying the plugin.Plugin interface.
func (*SuperclusterPlugin) Name() string {
	return "greeter"
}

// Version returns the plugin's version, satisfying the plugin.Plugin interface.
func (*SuperclusterPlugin) Version() string {
	return "0.1.0"
}

// Init initializes plugin, satisfying the plugin.Plugin interface. Put any
// initialization logic here.
func (*SuperclusterPlugin) Init(env *plugin.Environment) error {
	fmt.Println("Hello init!")
	return nil
}

func (*SuperclusterPlugin) Start(c coreiface.CoreAPI) error {
	// make sure we hold on to coreAPI before starting server
	setCoreAPIInstance(&c)

	go func(c coreiface.CoreAPI) {

		fmt.Println("Hello start!")
		r := gin.Default()
		store, err := newIpfsStore()
		if err != nil {
			panic("Cannot create store: " + err.Error())
		}

		addRoutes(r, store)
		ui.AddRoutes(r)

		r.Run(":3000")
	}(c)
	return nil
}

func (*SuperclusterPlugin) Close() error {
	fmt.Println("Goodbye!")
	return nil
}
