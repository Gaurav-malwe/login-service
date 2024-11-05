package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/Gaurav-malwe/login-service/mongodb/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// returns the mongo client
func (p *MongoDB) Client() *mongo.Client {
	client, err := p.client()
	if err != nil {
		log.Println("Cannot refresh mongo client")
		return p.mdbClient
	}
	return client
}

func (p *MongoDB) client() (*mongo.Client, error) {
	p.clientMu.Lock()
	defer p.clientMu.Unlock()

	if p.connectState == MONGO_CLIENT_NEEDS_TO_DISCONNECT {
		err := p.Close()
		if err != nil {
			return p.mdbClient, fmt.Errorf("mongodb MongoDB.Client p.closer error: %w", err)
		}
	}
	if p.connectState != MONGO_CLIENT_DICONNECTED {
		return p.mdbClient, nil
	}

	var (
		client *mongo.Client
		err    error
	)

	log.Println("MongoDB client needs to be refrehed")
	client, err = newClient(p.Config)
	if err != nil {
		return p.mdbClient, fmt.Errorf("mongodb MongoDB.client newClient Error: %w", err)
	}
	log.Println("MongoDB client refreshed")

	p.mdbClient = client
	p.connectState = MONGO_CLIENT_CONNECTED
	return p.mdbClient, nil
}

func newClient(c *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	opts, err := getClientOptions(c, c.AppName)
	if err != nil {
		return nil, errors.Wrap(err, "error creating MongoDB client options")
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "mongodb newClient mongo.NewClient error")
	}
	return client, nil
}

func getClientOptions(c *config.Config, appName string) (*options.ClientOptions, error) {
	uri, err := c.URI.GetURI()
	if err != nil {
		return nil, fmt.Errorf("util GetClientOptions error: %w", err)
	}

	opts := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(c.Timeout).
		SetHeartbeatInterval(c.HeartbeatInterval).
		SetMaxPoolSize(c.MaxPoolSize).
		SetMinPoolSize(c.MinPoolSize).
		SetMaxConnIdleTime(c.MaxConnIdleTime).
		SetPoolMonitor(c.PoolMonitor).
		SetMonitor(c.CommandMonitor).
		SetAppName(appName).
		SetServerSelectionTimeout(c.ServerSelectionTimeout).
		SetMonitor(otelmongo.NewMonitor())

	return opts, nil
}
