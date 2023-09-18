package models

import "time"

type User struct {
	ID          int                   `json:"id" gorm:"primary_key:auto_increment"`
	Fullname    string                `json:"fullname" gorm:"type: varchar(255)"`
	Email       string                `json:"email" gorm:"type: varchar(255)"`
	Password    string                `json:"password" gorm:"type: varchar(255)"`
	Status      string                `json:"status" gorm:"type:varchar(255)"`
	Phone       string                `json:"phone" gorm:"type:varchar(255)"`
	Image       string                `json:"image"  gorm:"type:varchar(255)"`
	Transaction []TransactionResponse `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Token       string                `json:"token" gorm:"type:varchar(255)"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	// Password string `json:"password" gorm:"type: varchar(255)"`
	Status string `json:"status" gorm:"type:varchar(255)"`
	// Phone  string `json:"phone" gorm:"type:varchar(255)"`
	// Transaction []TransactionResponse `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Token string `json:"token" gorm:"type:varchar(255)"`
}

// type UserProfileResponse struct{

// }

func (UserResponse) TableName() string {
	return "users"
}
