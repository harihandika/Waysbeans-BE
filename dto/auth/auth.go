package authdto

type RegisterRequest struct {
	Name     string `json:"name" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" gorm:"typer:varchar(255)"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" gorm:"typer:varchar(255)"`
}

type LoginResponse struct {
	Name    string `json:"name" gorm:"type: varchar(255)"`
	Email   string `json:"email" gorm:"type: varchar(255)"`
	Token   string `json:"token" gorm:"type: varchar(255)"`
	Status  string `gorm:"type: varchar(50)"  json:"status"`
	Address string `json:"address" gorm:"typer:varchar(255)"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
}
