package Entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string    `gorm:"not null;default:null" json:"title"`
	Description string    `gorm:"not null;default:null" json:"description"`
	DateTime    time.Time `gorm:"not null;default:null" json:"date_time"`
	Location    string    `gorm:"not null;default:null" json:"location"`
	AuthorID    uint      `gorm:"not null;default:null" json:"author_id"`
	Author      User      `gorm:"foreignKey:AuthorID" json:"author"`
	// Favorite    []Favorite `gorm:"many2many:favorites;joinForeignKey:EventID;joinReferences:FavoriteID" json:"favorites"`
}
