package transactiondto

import (
	"backend_project/models"
	"time"
)

type TransactionResponse struct {
	ID            int                 `json:"id"`
	Stardate      time.Time           `json:"startdate" form:"startdate"`
	Duedate       time.Time           `json:"duedate" form:"duedate"`
	UserID        int                 `json:"user_id" form:"user_id"`
	User          models.UserResponse `json:"user"`
	FilmID        int                 `json:"film_id"`
	Film          models.FilmResponse `json:"film"`
	Attache       string              `json:"attache" form:"attache"`
	Status        string              `json:"status" form:"status"`
	Email         string              `json:"email" form:"email"`
	AccountNumber int                 `json:"account_number"`
	TransferProof string              `json:"transfer_proof"`
	OrderDate     time.Time           `json:"-"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
}
