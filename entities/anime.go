package Entities

import "gorm.io/gorm"

type Anime struct {
	gorm.Model
	Title       string `gorm:"not null;default:null" json:"title"`
	Description string `gorm:"not null;default:null" json:"description"`
	// Reviews     []Review `gorm:"foreignKey:AnimeID" json:"reviews"`
}
