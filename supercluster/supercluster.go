package supercluster

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/SuperclusterLabs/supercluster-client/ui"
	"github.com/gin-gonic/gin"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	plugin "github.com/ipfs/kubo/plugin"
	"google.golang.org/api/option"
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
	return nil
}

func (*SuperclusterPlugin) Start(c coreiface.CoreAPI) error {
	// make sure we hold on to coreAPI before starting server
	setCoreAPIInstance(&c)

	// TODO: remove firebase
	// initialize firebase
	d, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	opt := option.WithCredentialsFile(d + "/.ipfs/keystore/supercluster-2d071-firebase-adminsdk-8qkm4-6688c64d73.json")
	config := &firebase.Config{
		DatabaseURL: "https://supercluster-2d071-default-rtdb.firebaseio.com/",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Println("Error initializing firebase: ", err.Error())
		panic("error initializing app: " + err.Error())
	}
	db = DB{instance: app}

	go func(c coreiface.CoreAPI) {
		defer close(wsCh)

		r := gin.Default()
		store, err := newIpfsStore()
		if err != nil {
			panic("Cannot create store: " + err.Error())
		}

		addRoutes(r, store)
		ui.AddRoutes(r)

		// TODO: add version dynamically
		log.Println("Supercluster started!")

		r.Run(":3000")
	}(c)
	return nil
}

func (*SuperclusterPlugin) Close() error {
	return nil
}
