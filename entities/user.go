package Entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;default:null" json:"email"`
	UserName string `gorm:"unique;not null;default:null" json:"username"`
	Password string `json:"password"`
	UserImage  string  `gorm:"default:null" json:"user_image"`
	Events  []Event `gorm:"many2many:user_events;joinForeignKey:UserID;joinReferences:EventID" json:"events"`
}
