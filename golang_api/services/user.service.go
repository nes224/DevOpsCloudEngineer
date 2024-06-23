package services

import (
	"context"

	"github.com/golang/nginx/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func FindUserById(id string, collection *mongo.Collection, ctx context.Context) (*models.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var user *models.DBResponse
	query := bson.M{"_id": oid}
	err := collection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}
	return user, nil
}

func FindUserByEmail(email string,collection *mongo.Collection, ctx context.Context) (*models.DBResponse, error) {
	query := bson.M{"email":email}
	var user *models.DBResponse
	err := collection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}
	return user, nil
}