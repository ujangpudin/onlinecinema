package repositories

import (
	"backend_project/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	// CreateUser(user models.User) (models.User, error)
	// CreateFotoProfil(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Transaction").Preload("Transaction.Film.Category").Preload("Transaction.User").Find(&users).Error // Using Find method

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction").Preload("Transaction.Film.Category").Preload("Transaction.User").First(&user, ID).Error // Using First method

	return user, err
}

// if use gorm sintax
// func (r *repository) CreateUser(user models.User) (models.User, error) {
// 	err := r.db.Exec("INSERT INTO users(fullname,email,password,gender,phone,address) VALUES (?,?,?,?,?,?)", user.Fullname, user.Email, user.Password, user.Gender, user.Phone, user.Address).Error

// 	return user, err
// }

// func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
// 	err := r.db.Raw("UPDATE users SET fullname=?, email=?, password=?, gender=?, phone=?, address=? WHERE id=?", user.Fullname, user.Email, user.Password, user.Gender, user.Phone, user.Address, ID).Scan(&user).Error

// 	return user, err
// }

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Save(&user).Error // Using Save method

	return user, err
}

// func (r *repository) CreateUser(user usersdto.CreateUserRequest) (usersdto.CreateUserRequest, error) {
// 	err := r.db.Create(&user).Error // Using Create method

// 	return user, err
// }

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}

// func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
// 	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

// 	return user, err
// }

// func (r *repository) CreateFotoProfil(user models.User) (models.User, error) {
// 	err := r.db.Create(&user).Error

// 	return user, err
// }
