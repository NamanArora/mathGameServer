package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (md *mongodb) CheckUsername(ctx context.Context, username string) (bool, error) {
	// Get the collection
	collection := md.client.Database("mydb").Collection("scores")

	// Check if the username exists in the collection
	filter := bson.M{"username": username}
	fmt.Printf("checking for username %s\n", username)
	count, err := collection.CountDocuments(ctx, filter, nil)
	if err != nil {
		return false, err
	}

	// Return a JSON response indicating if the username is unique or not
	return count > 0, nil
}
