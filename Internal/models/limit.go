package models

type Limit struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	CategoryID int     `json:"category_id"`
	Limit      float64 `json:"limit"`
}
