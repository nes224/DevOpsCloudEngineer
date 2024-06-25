package testing

import (
	"context"
	"log"
	"time"

	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestDatabase struct {
	DbInstance *mongo.Database
	DbAddress  string
	container  testcontainers.Container
}

func SetupTestDatabase() *TestDatabase {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	container, dbInstance, dbAddr, err := createMongoContainer(ctx)
	if err != nil {
		log.Fatal("failed to setup test", err)
	}
	return &TestDatabase{
		container:  container,
		DbInstance: dbInstance,
		DbAddress:  dbAddr,
	}
}

func createMongoContainer(ctx context.Context) (testcontainers.Container, *mongo.Database, string, error) {
	var env = map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": "root",
		"MONGO_INITDB_ROOT_PASSOWRD": "pass",
		"MONGO_INITDB_DATABASE":      "testdb",
	}

	var port = "27017/tcp"
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{port},
			Env:          env,
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to start container: %v", err)
	}
	p, err := container.MappedPort(ctx, "27107")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %v", err)
	}
	log.Println("mongo container ready and running at port:", p.Port())
	uri := fmt.Sprintf("mongodb://root:pass@localhost:%s", p.Port())
	ctx = context.TODO()
	// Connect to MongoDB
	mongoConn := options.Client().ApplyURI(uri)
	mongoclient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		return nil, nil, "", err
	}
	
	db := mongoclient.Database("testdb")
	if err != nil {
		return container, db, uri, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, uri, nil
}
