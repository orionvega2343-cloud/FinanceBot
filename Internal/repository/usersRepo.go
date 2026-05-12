package repository

import (
	"LangBot/Internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(user *models.User) error {
	_, err := u.db.Exec(`INSERT INTO users (tg_id,username) VALUES ($1,$2)`, user.TgId, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetByTgId(tgId int64) (*models.User, error) {
	var user models.User
	err := u.db.QueryRow(`SELECT id, tg_id, username FROM users WHERE tg_id=$1`, tgId).Scan(&user.Id, &user.TgId, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
