package repository

import (
	"ginessential/dao"
	"ginessential/models"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: dao.GetDB()}
}

func (c CategoryRepository) Create(name string) (*models.Category, error) {
	category := models.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		panic(err)
		//return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Update(category models.Category, name string) (*models.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) SelectById(id int) (*models.Category, error) {
	var category models.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
