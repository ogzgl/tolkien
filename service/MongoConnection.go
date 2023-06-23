package service

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	mongoOnce.Do(func() {
		connectionString := os.Getenv("MONGO_CONN")
		err := godotenv.Load()
		// Set client options
		clientOptions := options.Client().ApplyURI(connectionString)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
