package repository

import (
	"LangBot/Internal/models"

	"github.com/jmoiron/sqlx"
)

type LimitRepo struct {
	db *sqlx.DB
}

func NewLimitRepo(db *sqlx.DB) *LimitRepo {
	return &LimitRepo{db: db}
}

func (l *LimitRepo) CreateLimit(limit *models.Limit) error {
	_, err := l.db.Exec(`INSERT INTO limits (user_id,category_id,amount) VALUES($1,$2,$3)`, limit.UserId, limit.CategoryID, limit.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (l *LimitRepo) GetByUserId(id int) ([]models.Limit, error) {
	var limits []models.Limit
	rows, err := l.db.Query(`SELECT id,user_id,category_id,amount FROM limits WHERE user_id = $1`, id)
	if err != nil {
		return limits, err
	}
	defer rows.Close()
	for rows.Next() {
		var lim models.Limit
		err = rows.Scan(&lim.Id, &lim.UserId, &lim.CategoryID, &lim.Amount)
		if err != nil {
			return limits, err
		}
		limits = append(limits, lim)
	}
	return limits, nil
}

func (l *LimitRepo) GetByCategoryID(categoryId int) (*models.Limit, error) {
	var limit models.Limit
	err := l.db.QueryRow(`SELECT id,user_id,category_id,amount FROM limits WHERE category_id=$1`, categoryId).Scan(&limit.Id, &limit.UserId, &limit.CategoryID, &limit.Amount)
	if err != nil {
		return &limit, err
	}
	return &limit, nil
}
