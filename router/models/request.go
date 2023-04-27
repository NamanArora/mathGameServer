package models

type SaveScoreRequest struct {
	Username string `json:"username"`
	Country  string `json:"country"`
	UserID   string `json:"user_id"`
	Score    int    `json:"score"`
}

type CheckUsernameRequest struct {
	Username string `json:"username"`
}

type FindNearbyRequest struct {
	UserID  string `json:"user_id"`
	Country string `json:"country"`
}

type GetPercentileRequest struct {
	Score int `json:"score"`
}
