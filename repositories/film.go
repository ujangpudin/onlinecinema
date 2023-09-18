package repositories

import (
	"backend_project/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film, ID int) (models.Film, error)
	DeleteFilm(film models.Film, ID int) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilm() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error

	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category", ID).First(&film, ID).Error
	return film, err
}

// Write this code
func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error

	return film, err
}

func (r *repository) UpdateFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Save(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Delete(&film).Error // Using Delete method

	return film, err
}
