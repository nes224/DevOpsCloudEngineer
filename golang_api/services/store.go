package services

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/nginx/config"
	"github.com/golang/nginx/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client = Store()

var (
	ctx context.Context
)

type ServiceStore interface {
	SignUpUser(user *models.SignUpInput, ctx context.Context) (*models.DBResponse, error)
	SignInUser(user *models.SignInInput, ctx context.Context) (*models.DBResponse, error)
	FindUserById(id string, ctx context.Context) (*models.DBResponse, error)
	FindUserByEmail(id string, ctx context.Context) (*models.DBResponse, error)
}

func Transaction(callback func(ctx mongo.SessionContext) (any, error)) (any, error) {
	session, err := client.StartSession()
	if err != nil {
		return nil, fmt.Errorf("failed creating session | %s", err.Error())
	}
	defer session.EndSession(context.TODO())
	res, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, fmt.Errorf("failed executing transaction | %s", err.Error())
	}
	return res, nil
}

func Store() *mongo.Client {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx = context.TODO()
	// Connect to MongoDB
	mongoConn := options.Client().ApplyURI(config.MONGOURL) // ใช้ชื่อบริการ
	mongoclient, err := mongo.Connect(ctx, mongoConn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")
	return mongoclient
}
