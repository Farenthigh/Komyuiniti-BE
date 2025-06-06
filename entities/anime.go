package Entities

import "gorm.io/gorm"

type Anime struct {
	gorm.Model
	Title       string   `gorm:"not null;default:null" json:"title"`
	Description string   `gorm:"not null;default:null" json:"description"`
	Reviews     []Review `gorm:"foreignKey:AnimeID" json:"reviews"`
	Anime_image string   `gorm:"not null;default:null" json:"anime_image"`
	Rating	  float64  `gorm:"not null;default:0" json:"rating"`
	Genres      string `gorm:"not null;default:null" json:"genres"`
	Status 	string   `gorm:"not null;default:null" json:"status"`
	Studio  string   `gorm:"not null;default:null" json:"studio"`
	Episodes int      `gorm:"not null;default:0" json:"episodes"`
}
