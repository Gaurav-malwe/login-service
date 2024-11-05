package repository

import (
	"context"

	"github.com/Gaurav-malwe/login-service/internal/model"
	log "github.com/Gaurav-malwe/login-service/utils/logging"

	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	InsertUser(ctx context.Context, user *model.User) error
}

func (r *repository) InsertUser(ctx context.Context, user *model.User) error {
	log := log.GetLogger(ctx)
	log.Debug("Repository::InsertUser")

	_, err := r.getCollection(context.Background(), UserCollection).InsertOne(context.TODO(), user)
	if err != nil {
		log.WithContext(ctx).WithFields(map[string]interface{}{
			"error":  err,
			"method": "InsertUser",
		}).Error("Repository::InsertUser::Error while inserting user")
		return err
	}
	return nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	log := log.GetLogger(ctx)
	log.Debug("Repository::GetUserByEmail")

	var user model.User
	err := r.getCollection(context.Background(), UserCollection).FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.WithContext(ctx).WithFields(map[string]interface{}{
			"error": err,
		}).Error("Repository::GetUserByEmail::Error while getting user")
		return nil, err
	}
	return &user, nil
}
