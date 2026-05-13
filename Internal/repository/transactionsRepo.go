package repository

import (
	"LangBot/Internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (t *TransactionRepo) Create(tx *models.Transaction) error {
	_, err := t.db.Exec(`INSERT INTO transactions (user_id,type,category_id,sum,comment) VALUES ($1,$2,$3,$4,$5)`, tx.UserId, tx.Type, tx.CategoryId, tx.Sum, tx.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepo) GetByUserId(id int64, limit int, offset int) ([]models.Transaction, error) {
	res, err := t.db.Query(`SELECT user_id,type,category_id,sum,comment,created_at FROM transactions WHERE user_id=$1 LIMIT $2 OFFSET $3`, id, limit, offset)
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction //Create new slice
	defer res.Close()                     //Defer closing table
	for res.Next() {                      //iterate over a SQL call
		var tx models.Transaction                                                                 //create new model
		err = res.Scan(&tx.UserId, &tx.Type, &tx.CategoryId, &tx.Sum, &tx.Comment, &tx.CreatedAt) //Scan params
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx) //Add model to slice
	}
	return transactions, nil

}

func (t *TransactionRepo) GetById(id int) (*models.Transaction, error) {
	var tx models.Transaction
	err := t.db.QueryRow(`SELECT id,user_id,type,category_id,sum,comment FROM transactions WHERE id=$1`, id).Scan(&tx.Id, &tx.UserId, &tx.Type, &tx.CategoryId, &tx.Sum, &tx.Comment)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (t *TransactionRepo) Delete(id int) error {
	_, err := t.db.Exec(`DELETE FROM transactions WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (l *TransactionRepo) GetByCategoryAndMonth(categoryID int, from time.Time, to time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	rows, err := l.db.Query(`SELECT  id,user_id,category_id,sum FROM transactions WHERE category_id=$1 AND type='expense' AND created_at BETWEEN $2 AND $3
`, categoryID, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tx models.Transaction
		err = rows.Scan(&tx.Id, &tx.UserId, &tx.CategoryId, &tx.Sum)
		transactions = append(transactions, tx)
		if err != nil {
			return nil, err
		}
	}
	return transactions, nil
}
