package models

type User struct {
	Base
	Name     string     `json:"name"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Task     []Task     `gorm:"foreignKey:UserId"`
	Category []Category `gorm:"foreignKey:UserId"`
}
