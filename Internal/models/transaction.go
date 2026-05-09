package models

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Type       string    `json:"type"`
	CategoryId int       `json:"category_id"`
	Sum        float64   `json:"sum"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}
