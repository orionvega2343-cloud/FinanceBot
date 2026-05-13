package service

import (
	"LangBot/Internal/models"
	"LangBot/Internal/repository"
	"time"
)

type LimitsService struct {
	repo  *repository.LimitRepo
	tRepo *repository.TransactionRepo
}

func NewLimitService(repo *repository.LimitRepo, tRepo *repository.TransactionRepo) *LimitsService {
	return &LimitsService{repo: repo, tRepo: tRepo}
}

func (l *LimitsService) CreateLimitsService(lim *models.Limit) error {
	err := l.repo.CreateLimitRepo(lim)
	if err != nil {
		return err
	}
	return nil
}

func (l *LimitsService) GetLimitByUserIdService(id int) ([]models.Limit, error) {
	lim, err := l.repo.GetByUserId(id)
	if err != nil {
		return nil, err
	}
	return lim, nil
}

func (l *LimitsService) CheckLimit(CategoryId int, from time.Time, to time.Time) (bool, error) {
	var total float64
	lim, err := l.repo.GetByCategoryID(CategoryId)
	if err != nil {
		return false, err
	}
	exp, err := l.tRepo.GetByCategoryAndMonth(CategoryId, from, to)
	if err != nil {
		return false, err
	}

	for _, tx := range exp {
		total += tx.Sum
	}
	if total > lim.Amount {
		return true, nil
	}
	return false, nil
}
