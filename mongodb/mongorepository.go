package mongodb

import (
	"context"
	"log"

	uuid "github.com/gofrs/uuid/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// IMongoRepository
type IMongoRepository interface {
	MongoClient() *mongo.Client
	MongoDatabase(ctx context.Context) *mongo.Database
	RunTransaction(ctx context.Context, txnFunc func(sessionContext mongo.SessionContext) error) error
}

// MongoRepository is an abstract of mongodb provider which integrates
// a MongoDB instance and AbstractProvider.
type MongoRepository struct {
	mongoProvider *MongoDB
}

// NewMongoRepositoryWithDatabaseNameGetter returns a initialized IMongoRepository with DatabaseNameGetter.
func NewMongoRepositoryWithDatabaseNameGetter(mongoProvider *MongoDB) IMongoRepository {
	return &MongoRepository{
		mongoProvider: mongoProvider,
	}
}

// NewMongoRepository returns a initialized IMongoRepository.
func NewMongoRepository(mongoProvider *MongoDB) IMongoRepository {
	return &MongoRepository{
		mongoProvider: mongoProvider,
	}
}

// MongoClient returns a mongoDB client.
func (m MongoRepository) MongoClient() *mongo.Client {
	return m.mongoProvider.Client()
}

// MongoDatabase returns a mongonDB database.
func (m MongoRepository) MongoDatabase(ctx context.Context) *mongo.Database {
	return m.mongoProvider.Client().Database(m.mongoProvider.Config.Database)
}

// RunTransaction runs a mongodb transaction.
func (m MongoRepository) RunTransaction(ctx context.Context, txnFn func(mongo.SessionContext) error) error {
	var err error
	for i := 0; i < 2; i++ {
		err = m.mongoProvider.Client().UseSessionWithOptions(ctx,
			options.Session().SetDefaultReadPreference(readpref.Primary()), func(sessionContext mongo.SessionContext) error {
				return m.runTransactionWithRetry(sessionContext, txnFn)
			})

		return err
	}

	return err
}

func (m MongoRepository) runTransactionWithRetry(sctx mongo.SessionContext, txnFn func(mongo.SessionContext) error) error {
	txIdUUID, _ := uuid.NewV4()
	txID := txIdUUID.String()
	log.Println("Begin transaction: ", txID)
	for {
		// start transaction
		err := sctx.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.Majority()),
		)
		if err != nil {
			log.Printf("Start transaction: %s. Error: %s\n", txID, err)
			return err
		}

		err = txnFn(sctx)
		if err != nil {
			log.Printf("Abort transaction: %s. Error: %s\n", txID, err)
			if abortErr := sctx.AbortTransaction(sctx); abortErr != nil {
				log.Println("Failed to abort transaction: ", txID)
			}
			return err
		}

		err = m.commitWithRetry(sctx)
		switch e := err.(type) {
		case nil:
			log.Printf("End transaction: %s, successful. \n", txID)
			return nil
		case mongo.CommandError:
			// If transient error, retry the whole transaction
			if e.HasErrorLabel("TransientTransactionError") {
				log.Printf("TransientTransactionError: %s, retrying transaction...\n", txID)
				continue
			}
			return e
		default:
			log.Printf("End transaction: %s. Error: %s\n", txID, err)
			return e
		}
	}
}

func (m MongoRepository) commitWithRetry(sctx mongo.SessionContext) error {
	for {
		err := sctx.CommitTransaction(sctx)
		switch e := err.(type) {
		case nil:
			return nil
		case mongo.CommandError:
			// Can retry commit
			if e.HasErrorLabel("UnknownTransactionCommitResult") {
				continue
			}
			return e
		default:
			return e
		}
	}
}
