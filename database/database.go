package database

import (
	"context"

	"github.com/NamanArora/mathGameServer/database/models"
	"github.com/NamanArora/mathGameServer/database/mongo"
)

type Operations interface {
	CheckUsername(context.Context, string) (bool, error)
	FindNearby(context.Context, string, string) ([]models.Score, error)
	SaveScore(context.Context, models.Score) error
	GetPercentile(ctx context.Context, score int) (float32, error)
}

func DefaultDatabase() Operations {
	return mongo.GetMongoClient()
}
