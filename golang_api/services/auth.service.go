package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/golang/nginx/models"
	"github.com/golang/nginx/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (db *DatabaseImpl) UserCollection(mongoclient *mongo.Client ,database string, collection string)  {
	mongoclient.Database(database).Collection("users")
}


func (db *DatabaseImpl) SignUpUser(user *models.SignUpInput, ctx context.Context) (*models.DBResponse, error) {
	session, err := db.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt
		user.Email = strings.ToLower(user.Email)
		user.PasswordConfirm = ""
		user.Verified = false
		user.Role = "user"

		hashedPassword, _ := utils.HashPassword(user.Password)
		user.Password = hashedPassword

		res, err := db.collection.InsertOne(sessCtx, user)
		if err != nil {
			if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
				return nil, errors.New("user with that email already exists")
			}
			return nil, err
		}
		opt := options.Index().SetUnique(true)
		index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
		if _, err := db.collection.Indexes().CreateOne(sessCtx, index); err != nil {
			return nil, errors.New("could not create index for email")
		}

		var newUser *models.DBResponse
		query := bson.M{"_id": res.InsertedID}
		err = db.collection.FindOne(ctx, query).Decode(&newUser)
		if err != nil {
			return nil, err
		}
		return newUser, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.DBResponse), nil
}
