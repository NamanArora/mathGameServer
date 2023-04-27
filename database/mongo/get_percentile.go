package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (md *mongodb) GetPercentile(ctx context.Context, score int) (float32, error) {
	collection := md.client.Database("mydb").Collection("scores")
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	// Get the count of scores that are less than or equal to the given score
	rank, err := collection.CountDocuments(ctx, bson.M{"score": bson.M{"$lte": score}})
	if err != nil {
		return 0, err
	}

	percentile := float32(rank) / float32(count) * 100
	return percentile, nil
}
