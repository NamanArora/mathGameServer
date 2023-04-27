package models

type GetPercentileResponse struct {
	Percentile float32 `json:"percentile"`
	Text       string  `json:"text"`
}

type Rank struct {
	Username string `json:"username"`
	Rank     int    `json:"rank"`
	Score    int    `json:"score"`
	Country  string `json:"country"`
}

type FindNearbyResponse struct {
	Ranks   []Rank `json:"ranks"`
	Message string
}

type CheckUsernameResponse struct {
	Exists bool `json:"exists"`
}

type SaveScoreResponse struct {
	Message string
}
