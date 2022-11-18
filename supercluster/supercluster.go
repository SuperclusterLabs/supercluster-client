package supercluster

import (
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/SuperclusterLabs/supercluster-client/ui"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	fmt.Println("Hello init!")
	return nil
}

func (*SuperclusterPlugin) Start(c coreiface.CoreAPI) error {
	// make sure we hold on to coreAPI before starting server
	setCoreAPIInstance(&c)

	// TODO: remove firebase
	// initialize firebase
	opt := option.WithCredentialsFile("/home/gov/dev/supercluster-client/supercluster-2d071-firebase-adminsdk-8qkm4-6688c64d73.json")
	config := &firebase.Config{
		DatabaseURL: "https://supercluster-2d071-default-rtdb.firebaseio.com/",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Println("Error initializing firebase: ", err.Error())
		panic("error initializing app: " + err.Error())
	}
	db = DB{instance: app}

	/** test **/
	ctx := context.Background()
	acc := User{
		Id:       uuid.New(),
		EthAddr:  "0xE4475EF8717d14Bef6dCBAd55E41dE64a0cc8510",
		IpfsAddr: "12D3KooWCk54bkeehLMDv52vmjTEvsB7EvXyA7s3E9WsGFUYudoY",
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("test")
	if err := client.NewRef("accounts/alice").Set(ctx, acc); err != nil {
		log.Fatal(err)
	}
	/**/

	go func(c coreiface.CoreAPI) {
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
