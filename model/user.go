package model

type UserModel struct {
	ID       uint `gorm:"unique"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Status   string
}
