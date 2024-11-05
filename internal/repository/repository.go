package repository

import (
	"context"

	"github.com/Gaurav-malwe/login-service/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UserCollection = "user"
)

type (
	Repository interface {
		IUserRepository
	}

	repository struct {
		mongodb.IMongoRepository
	}
)

func New(mongodbProvider *mongodb.MongoDB) Repository {
	repo := mongodb.NewMongoRepository(mongodbProvider)
	return &repository{
		IMongoRepository: repo,
	}
}

func (r *repository) getCollection(ctx context.Context, name string) *mongo.Collection {
	return r.MongoDatabase(ctx).Collection(name)
}
