package Entities

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Description string `gorm:"not null;default:null" json:"description"`
	AuthorID    uint   `gorm:"not null;default:null" json:"author_id"`
	Author      User   `gorm:"foreignKey:AuthorID" json:"author"`
}
