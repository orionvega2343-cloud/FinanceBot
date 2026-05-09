package models

import "time"

type User struct {
	Id        int       `json:"id"`
	TgId      int64     `json:"tg_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
