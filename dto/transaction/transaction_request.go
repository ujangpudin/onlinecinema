package transactiondto

import (
	"backend_project/models"
	"time"
)

type TransactionRequest struct {
	ID            int                 `json:"id"`
	StartDate     string              `json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	DueDate       string              `json:"dueDate" form:"dueDate" gorm:"type:varchar(255)"`
	UserID        int                 `json:"user_id" form:"user_id" gorm:"-"`
	User          models.UserResponse `json:"user"`
	FilmID        int                 `json:"film_id"`
	Film          models.FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status        string              `json:"status"`
	AccountNumber int                 `json:"account_number"`
	TransferProof string              `json:"transfer_proof"`
	OrderDate     time.Time           `json:"-"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
}

type TransactionUpdateRequest struct {
	// ID				int							`json:"id" gorm:"primary_key:auto_increment"`
	StartDate     time.Time           `json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	DueDate       time.Time           `json:"dueDate" form:"dueDate" gorm:"type:varchar(255)"`
	UserID        int                 `json:"user_id" form:"user_id" gorm:"-"`
	User          models.UserResponse `json:"user"`
	Attache       string              `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	FilmID        int                 `json:"film_id"`
	Film          models.FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status        string              `json:"status"`
	AccountNumber int                 `json:"account_number"`
	TransferProof string              `json:"transfer_proof"`
	OrderDate     time.Time           `json:"-"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
}

type TransactionUpdateResponse struct {
	ID            int                 `json:"id"`
	StartDate     time.Time           `json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	DueDate       time.Time           `json:"dueDate" form:"dueDate" gorm:"type:varchar(255)"`
	UserID        int                 `json:"user_id" form:"user_id" gorm:"-"`
	User          models.UserResponse `json:"user"`
	FilmID        int                 `json:"film_id"`
	Film          models.FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status        string              `json:"status"`
	AccountNumber int                 `json:"account_number"`
	Attache       string              `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	TransferProof string              `json:"transfer_proof"`
	OrderDate     time.Time           `json:"-"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
}

type TransactionDeleteResponse struct {
	ID int `json:"id"`
}
