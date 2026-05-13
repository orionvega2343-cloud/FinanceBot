package service

import (
	"LangBot/Internal/models"
	"LangBot/Internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepo
}

func NewTransactionService(repo *repository.TransactionRepo) *TransactionService {
	return &TransactionService{repo: repo}
}

func (t *TransactionService) CreateTransaction(tx *models.Transaction) error {
	err := t.repo.Create(tx)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionService) GetTransactionById(id int) (*models.Transaction, error) {
	tx, err := t.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *TransactionService) GetTransactionByUserId(userId int64, limit int, offset int) ([]models.Transaction, error) {
	tx, err := t.repo.GetByUserId(userId, limit, offset)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *TransactionService) DeleteTransactionById(id int) error {
	err := t.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
