package repository

import (
	"LangBot/Internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

// Interface
type StatsRepo struct {
	db *sqlx.DB
}

// Constructor
func NewStatsRepo(db *sqlx.DB) *StatsRepo {
	return &StatsRepo{db: db}
}

func (s *StatsRepo) GetSummary(userId int, from time.Time, to time.Time) (*models.Summary, error) {
	var sum models.Summary
	err := s.db.QueryRow(`SELECT 
    SUM(CASE WHEN type='income' THEN sum ELSE 0 END),
    SUM(CASE WHEN type='expense' THEN sum ELSE 0 END)
	FROM transactions 
	WHERE user_id=$1 AND created_at BETWEEN $2 AND $3`, userId, from, to).Scan(&sum.Income, &sum.Expense)
	if err != nil {
		return nil, err
	}
	sum.Balance = sum.Income - sum.Expense
	return &sum, nil

}

func (s *StatsRepo) GetCategory(userId int, from time.Time, to time.Time) ([]models.TopCategory, error) {
	res, err := s.db.Query(`
	SELECT c.name, SUM(t.sum) as total
	FROM transactions t
	JOIN categories c ON t.category_id = c.id
	WHERE t.user_id=$1 AND t.type='expense' 
	AND t.created_at BETWEEN $2 AND $3
	GROUP BY c.name
	ORDER BY total DESC
	LIMIT 3`, userId, from, to)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var categories []models.TopCategory
	for res.Next() {
		var c models.TopCategory
		err = res.Scan(&c.CategoryName, &c.Total)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}
