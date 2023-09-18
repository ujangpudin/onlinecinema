package repositories

import (
	"backend_project/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	UpdatesTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	GetFilmTransaction(ID int) (models.Film, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetFilmTransaction(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category", ID).First(&film, ID).Error
	return film, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Film").Preload("Film.Category").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").Preload("Film").Preload("Film.Category").First(&transaction, ID).Error
	// .Preload("Transaction.User")

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, ID)
	if status != transaction.Status && status == "success" {
		var user models.User
		r.db.First(&user, transaction.UserID)
		// user.Subscribe = "Active"
		r.db.Save(&user)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error
	return err
}

func (r *repository) UpdatesTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
