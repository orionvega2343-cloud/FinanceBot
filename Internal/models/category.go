package models

type Category struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	UserId    int    `json:"user_id"`
	IsDefault bool   `json:"-"`
}
