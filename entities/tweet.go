package Entities

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Description string `gorm:"not null;default:null" json:"description"`
	AuthorID    uint   `gorm:"not null;default:null" json:"author_id"`
	Author      User   `gorm:"foreignKey:AuthorID" json:"author"`
	// Favorite    []Favorite `gorm:"many2many:favorites;joinForeignKey:TweetID;joinReferences:FavoriteID" json:"favorites"`
	Comments    []Comment  `gorm:"foreignKey:TweetID" json:"comments"`
	Tweet_image string    `gorm:"nullable;default:null" json:"tweet_image"`
}
