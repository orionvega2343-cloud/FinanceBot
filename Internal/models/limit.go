package models

type Limit struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	CategoryID int     `json:"category_id"`
	Amount     float64 `json:"limit"`
}
