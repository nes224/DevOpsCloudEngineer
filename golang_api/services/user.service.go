package services

import (
	"context"

	"github.com/golang/nginx/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func FindUserById(id string, ctx context.Context) (*models.DBResponse, error) {
	db := client.Database("golang_app").Collection("users")

	oid, _ := primitive.ObjectIDFromHex(id)
	var user *models.DBResponse
	query := bson.M{"_id": oid}
	err := db.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}
	return user, nil
}

func FindUserByEmail(email string, ctx context.Context) (*models.DBResponse, error) {
	db := client.Database("golang_app").Collection("users")
	query := bson.M{"email": email}
	var user *models.DBResponse
	err := db.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}
	return user, nil
}
