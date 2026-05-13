package repository

import (
	"LangBot/Internal/models"

	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (u *CategoryRepo) CreateCategory(c *models.Category) error {
	_, err := u.db.Exec(`INSERT INTO categories (name,user_id,is_default) VALUES ($1,$2,$3)`, c.Name, c.UserId, c.IsDefault)
	if err != nil {
		return err
	}
	return nil
}

func (u *CategoryRepo) GetById(id int) (*models.Category, error) {
	var c models.Category
	err := u.db.QueryRow(`SELECT id,name,user_id,is_default FROM categories WHERE id=$1`, id).Scan(&c.Id, &c.Name, &c.UserId, &c.IsDefault)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (u *CategoryRepo) GetByUserId(id int) ([]models.Category, error) {
	var categories []models.Category
	res, err := u.db.Query("SELECT id,name,user_id,is_default 	FROM categories WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var c models.Category
		err = res.Scan(&c.Id, &c.Name, &c.UserId, &c.IsDefault)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (u *CategoryRepo) Delete(id int) error {
	_, err := u.db.Exec(`DELETE FROM categories WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *CategoryRepo) DefaultCategories(id int) error {
	categories := []string{"Еда", "Транспорт", "Развлечения", "Здоровье", "Другое"}
	for _, name := range categories {
		var c models.Category
		c.Name = name
		c.UserId = id
		c.IsDefault = true
		err := u.CreateCategory(&c)
		if err != nil {
			return err
		}
	}
	return nil
}
