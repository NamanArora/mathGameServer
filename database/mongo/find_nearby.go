package mongo

import (
	"context"

	"github.com/NamanArora/mathGameServer/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (md *mongodb) FindNearby(ctx context.Context, userID, country string) ([]models.Score, error) {
	// Get a handle to the scores collection
	scoresColl := md.client.Database("mydb").Collection("scores")

	// Build the query to get nearby scores
	filter := bson.M{}
	if country != "" {
		filter["country"] = country
	}
	sort := bson.M{"score": -1}
	options := options.Find().SetSort(sort).SetLimit(50)

	// Execute the query to get the nearby scores
	cursor, err := scoresColl.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Build the response with the nearby scores
	var scores []models.Score
	for cursor.Next(ctx) {
		var score models.Score
		if err := cursor.Decode(&score); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Calculate the rank of the user
	// userRank := 0
	// for i, score := range scores {
	// 	if score.UserID == userID {
	// 		userRank = i + 1
	// 		break
	// 	}
	// }

	return scores, nil
}
