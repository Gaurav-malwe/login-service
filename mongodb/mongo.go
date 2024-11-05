package mongodb

import (
	"context"
	"log"
	"sync"

	"github.com/Gaurav-malwe/login-service/mongodb/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type clientConnectState uint8

const (
	MONGO_CLIENT_CONNECTED           clientConnectState = 0
	MONGO_CLIENT_NEEDS_TO_DISCONNECT clientConnectState = 1
	MONGO_CLIENT_DICONNECTED         clientConnectState = 2
)

type MongoDB struct {
	Config       *config.Config
	mdbClient    *mongo.Client
	connectState clientConnectState
	clientMu     sync.Mutex
}

func NewConfigFromEnv() *config.Config {
	return config.NewConfigFromEnv()
}

func New(c *config.Config) *MongoDB {
	if c == nil {
		c = config.NewConfigFromEnv()
	}
	return &MongoDB{
		Config: c,
	}
}

func (p *MongoDB) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), p.Config.Timeout)
	defer cancel()

	log.Println("db", p.Config.Database, "Connecting to MongoDb Server...")

	var client *mongo.Client
	var cerr error

	client, cerr = newClient(p.Config)
	if cerr != nil {
		log.Println("error creating MongoDB client")
		return errors.Wrap(cerr, "error creating MongoDB client")
	}

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Println("Failed to launch MongoDB Provider")
		return errors.Wrap(err, "mongodb MongoDb.Init client.Ping error")
	}

	p.mdbClient = client
	p.connectState = MONGO_CLIENT_CONNECTED

	return nil
}

func (p *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), p.Config.Timeout)
	defer cancel()

	err := p.mdbClient.Disconnect(ctx)
	if err != nil {
		log.Println("MongoDB disconnecting failed")
		return err
	}

	p.connectState = MONGO_CLIENT_DICONNECTED
	return nil
}
