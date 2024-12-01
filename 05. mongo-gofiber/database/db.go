package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = "testdb"

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MG MongoInstance

func getenv() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
		return "", err
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}
	return uri, nil
}

func Connect() error {
	uri, err := getenv()
	if err != nil {
		log.Fatal("Error getting environment variable", err)
	}

	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}
	opt := options.Client().
		ApplyURI(uri).
		SetBSONOptions(bsonOpts)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opt)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	MG = MongoInstance{
		DB:     db,
		Client: client,
	}

	return nil
}

/*
docker run -d \
--name mongodb-container \
-p 27017:27017 \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=adminpassword \
mongo
*/

// mongodb://admin:adminpassword@localhost:27017

