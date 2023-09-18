package models

import "time"

type Transaction struct {
	ID            int          `json:"id" gorm:"primary_key:auto_increment"`
	StartDate     time.Time    `json:"startDate"`
	DueDate       time.Time    `json:"dueDate"`
	UserID        int          `json:"user_id"`
	User          UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price         int          `json:"price"`
	FilmID        int          `json:"film_id"`
	Film          FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status        string       `json:"status"`
	AccountNumber int          `json:"account_number"`
	TransferProof string       `json:"transfer_proof"`
	Attache       string       `json:"attache"`
	OrderDate     time.Time    `json:"-"`
	CreatedAt     time.Time    `json:"-"`
	UpdatedAt     time.Time    `json:"-"`
}

type TransactionResponse struct {
	ID            int          `json:"id"`
	StartDate     time.Time    `json:"startDate"`
	DueDate       time.Time    `json:"dueDate"`
	UserID        int          `json:"user_id"`
	User          UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FilmID        int          `json:"film_id"`
	Film          FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status        string       `json:"status"`
	AccountNumber int          `json:"account_number"`
	TransferProof string       `json:"transfer_proof"`
	Attache       string       `json:"attache"`
	OrderDate     time.Time    `json:"-"`
	CreatedAt     time.Time    `json:"-"`
	UpdatedAt     time.Time    `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
