package service

import (
	"LangBot/Internal/models"
	"LangBot/Internal/repository"
	"time"
)

type StatsService struct {
	repo *repository.StatsRepo
}

func NewStatsService(repo *repository.StatsRepo) *StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetSummaryService(userId int, from time.Time, to time.Time) (*models.Summary, error) {
	sum, err := s.repo.GetSummary(userId, from, to)
	if err != nil {
		return nil, err
	}
	return sum, nil
}

func (s *StatsService) GetCategoryService(userId int, from time.Time, to time.Time) ([]models.TopCategory, error) {
	ct, err := s.repo.GetCategory(userId, from, to)
	if err != nil {
		return nil, err
	}
	return ct, nil
}
