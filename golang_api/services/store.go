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
var (
	ctx context.Context
)

type Queries interface {
}

type Store interface {
	SignUpUser(user *models.SignUpInput) (*models.DBResponse, error)
}


func StoreDb(uri string) {
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

}
