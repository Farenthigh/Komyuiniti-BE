package Entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	AnimeID  uint   `gorm:"not null;default:null" json:"anime_id"`
	Anime    Anime  `gorm:"foreignKey:AnimeID" json:"anime"`
	Review   string `gorm:"not null;default:null" json:"review"`
	AuthorID uint   `gorm:"not null;default:null" json:"author_id"`
	Author   User   `gorm:"foreignKey:AuthorID" json:"author"`
}
