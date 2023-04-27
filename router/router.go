package router

import (
	"fmt"
	"net/http"

	"github.com/NamanArora/mathGameServer/database"
	dbModels "github.com/NamanArora/mathGameServer/database/models"
	"github.com/NamanArora/mathGameServer/router/models"
	"github.com/gin-gonic/gin"
)

func SaveScore(c *gin.Context) {
	var req models.SaveScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.SaveScoreResponse{
			Message: err.Error(),
		})
		return
	}

	db := database.DefaultDatabase()
	err := db.SaveScore(c, dbModels.Score{
		Username: req.Username,
		Country:  req.Country,
		Score:    req.Score,
		UserID:   req.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.SaveScoreResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SaveScoreResponse{
		Message: "DONE",
	})
}

func FindNearby(c *gin.Context) {
	var req models.FindNearbyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.FindNearbyResponse{Message: err.Error()})
		return
	}
	fmt.Printf("%+v", req)
	db := database.DefaultDatabase()
	scores, err := db.FindNearby(c, req.UserID, req.Country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FindNearbyResponse{Message: err.Error()})
		return
	}

	rankList := make([]models.Rank, 0, len(scores))

	for idx, score := range scores {
		rankList = append(rankList, models.Rank{
			Rank:     idx + 1,
			Username: score.Username,
			Score:    score.Score,
			Country:  score.Country,
		})
	}

	c.JSON(http.StatusOK, models.FindNearbyResponse{
		Message: "OK",
		Ranks:   rankList,
	})
}

func CheckUsername(c *gin.Context) {
	var req models.CheckUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.CheckUsernameResponse{})
		return
	}

	db := database.DefaultDatabase()
	exists, err := db.CheckUsername(c, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.CheckUsernameResponse{
		Exists: exists,
	})
}

func GetPercentile(c *gin.Context) {
	var req models.GetPercentileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.GetPercentileResponse{})
		return
	}
	fmt.Printf("%+v", req)
	db := database.DefaultDatabase()
	p, err := db.GetPercentile(c, req.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.GetPercentileResponse{
		Percentile: p,
		Text:       motivationalText(p),
	})
}

func motivationalText(percentile float32) string {
	var motivation string

	switch {
	case percentile >= 90 && percentile <= 100:
		motivation = fmt.Sprintf("You're a math superstar! Keep up the amazing work and stay at the top.")
	case percentile >= 70 && percentile <= 89:
		motivation = fmt.Sprintf("Fantastic job! You're among the top performers. Keep practicing to reach the next level.")
	case percentile >= 40 && percentile <= 69:
		motivation = fmt.Sprintf("Great effort! You're making steady progress. Keep practicing and you'll see the improvement.")
	case percentile >= 0 && percentile <= 39:
		motivation = fmt.Sprintf("Good start! Keep practicing and you'll soon see the results.")
	default:
		motivation = fmt.Sprintf("Invalid percentile score. Please try again.")
	}

	return motivation
}
