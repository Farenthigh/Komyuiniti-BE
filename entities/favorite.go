package Entities

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	// Tweets  []Tweet  `gorm:"many2many:tweets;joinForeignKey:TweetID;joinReferences:TweetID"`
	// Events  []Event  `gorm:"many2many:events;joinForeignKey:EventID;joinReferences:EventID"`
	// Reviews []Review `gorm:many2many:reviews;joinForeignKey:ReviewID;joinReferences:ReviewID`
}
