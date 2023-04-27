package mongo

import (
	"context"

	"github.com/NamanArora/mathGameServer/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (md *mongodb) SaveScore(ctx context.Context, score models.Score) error {
	collection := md.client.Database("mydb").Collection("scores")

	filter := bson.M{"_id": score.UserID}
	var existingScore models.Score
	err := collection.FindOne(ctx, filter).Decode(&existingScore)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if existingScore.UserID == score.UserID && existingScore.Score < score.Score {
		update := bson.M{"$set": bson.M{"score": score.Score}}
		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
	} else {
		_, err := collection.InsertOne(ctx, score)
		if err != nil {
			return err
		}
	}

	return nil
}
