package user_entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Surname  string `gorm:"not null" json:"surname"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `gorm:"not null, size:255" json:"password"`
}