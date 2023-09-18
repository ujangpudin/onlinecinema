package usersdto

import "time"

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	// Password string `json:"password"`
	Status string `json:"status" gorm:"type:varchar(255)"`
	Phone  string `json:"phone" `
	Image  string `json:"image" `
	// Transaction []TransactionResponse `json:"transaction_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Token     string    `json:"token" gorm:"type:varchar(255)"`
}
