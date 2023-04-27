package models

type Score struct {
	Username string `json:"username" bson:"username"`
	Score    int    `json:"score" bson:"score"`
	Country  string `json:"country" bson:"country"`
	UserID   string `json:"userid" bson:"_id"`
}