package Entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;default:null" json:"email"`
	UserName string `gorm:"unique;not null;default:null" json:"username"`
	Password string `json:"password"`
}
