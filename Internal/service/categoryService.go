package service

import (
	"LangBot/Internal/models"
	"LangBot/Internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepo
}

func NewCategoryService(repo *repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (t *CategoryService) CreateCategory(tx *models.Category) error {
	err := t.repo.CreateCategory(tx)
	if err != nil {
		return err
	}
	return nil
}

func (t *CategoryService) GetCategoryService(id int) (*models.Category, error) {
	tx, err := t.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *CategoryService) GetByUserIdService(id int) ([]models.Category, error) {
	tx, err := t.repo.GetByUserId(id)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *CategoryService) DeleteCategoryService(id int) error {
	err := t.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *CategoryService) DefaultCtService(id int) error {
	err := t.repo.DefaultCategories(id)
	if err != nil {
		return err
	}
	return nil
}
