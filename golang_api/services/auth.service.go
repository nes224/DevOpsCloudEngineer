package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang/nginx/models"
	"github.com/golang/nginx/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func TxSignUpUser(user *models.SignUpInput, ctx context.Context) (*models.DBResponse, error) {
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}

	db := client.Database("golang_app").Collection("users")
	defer session.EndSession(ctx)
	callback := func(ctx mongo.SessionContext) (any, error) {
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt
		user.Email = strings.ToLower(user.Email)
		user.PasswordConfirm = ""
		user.Verified = false
		user.Role = "user"
		hashedPassword, _ := utils.HashPassword(user.Password)
		user.Password = hashedPassword
		res, err := db.InsertOne(ctx, user)
		if err != nil {
			if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
				return nil, errors.New("user with that email already exists")
			}
			return nil, err
		}
		opt := options.Index().SetUnique(true)
		index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
		if _, err := db.Indexes().CreateOne(ctx, index); err != nil {
			return nil, errors.New("could not create index for email")
		}

		var newUser *models.DBResponse
		query := bson.M{"_id": res.InsertedID}
		err = db.FindOne(ctx, query).Decode(&newUser)
		if err != nil {
			return nil, err
		}
		return newUser, nil
	}
	fmt.Println("callback::", callback)
	result, err := Transaction(callback)
	if err != nil {
		return &models.DBResponse{}, fmt.Errorf("failed executing transaction | %s", err.Error())
	}
	return result.(*models.DBResponse), nil
}

func SignUpUser(user *models.SignUpInput, ctx context.Context) (*models.DBResponse, error) {
	db := client.Database("golang_app").Collection("users")
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	res, err := db.InsertOne(ctx, user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exists")
		}
		return nil, err
	}
	opt := options.Index().SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	if _, err := db.Indexes().CreateOne(ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}
	err = db.FindOne(ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func SignInUser(user *models.SignInInput) (*models.DBResponse, error) {
	return nil, nil
}
