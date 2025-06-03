package Entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	// TweetID  uint   `gorm:"not null;default:null" json:"tweet_id"`
	// Comment  string `gorm:"not null;default:null" json:"comment"`
	// AuthorID uint   `gorm:"not null;default:null" json:"author_id"`
	// Author   User   `gorm:"foreignKey:AuthorID" json:"author"`
}
