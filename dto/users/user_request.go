package usersdto

type CreateUserRequest struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	// Email    string `json:"email"`
	// Password string `json:"password"`
	Phone string `json:"phone"`
	Image string `json:"image"`
}
