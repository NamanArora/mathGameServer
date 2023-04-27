package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientOnce sync.Once
	client     *mongo.Client
)

type mongodb struct {
	client *mongo.Client
}

// GetMongoClient returns a singleton instance of the MongoDB client
func GetMongoClient() *mongodb {
	var err error
	clientOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@mongodb_container:27017")

		// Connect to MongoDB
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Check the connection
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	return &mongodb{client: client}
}
