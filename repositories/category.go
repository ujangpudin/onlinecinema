package repositories

import (
	"backend_project/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategory() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category, ID int) (models.Category, error)
	DeleteCategory(category models.Category, ID int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategory() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Find(&category).Error

	return category, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

// func (r *repository) UpdateCategory(category models.Category, ID int) (models.Category, error) {
// 	err := r.db.Raw("UPDATE categories SET name=? WHERE id=?", category.Name, ID).Scan(&category).Error

// 	return category, err
// }

func (r *repository) UpdateCategory(category models.Category, ID int) (models.Category, error) {
	err := r.db.Save(&category).Error // Using Save method

	return category, err
}

// func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {
// 	err := r.db.Raw("DELETE FROM categories WHERE id=?", ID).Scan(&category).Error

// 	return category, err
// }

func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {
	err := r.db.Delete(&category).Error

	return category, err
}
