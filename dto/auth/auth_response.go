package authdto

type AuthResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}

type LoginResponse struct {
	// ID     int         `json:"id"`
	// Status string      `gorm:"type: varchar(255)" json:"status"`
	// UserID int         `gorm:"type:varchar(255)" json:"user_id"`
	// User   models.User `gorm:"type: varchar(255)" json:"user"`
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}

type RegisterResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}
