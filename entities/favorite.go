package Entities

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	// Tweets  []Tweet  `gorm:"many2many:tweets;joinForeignKey:FavoriteID;joinReferences:TweetID" json:"tweets_id"`
	// Events  []Event  `gorm:"many2many:events;joinForeignKey:EventID;joinReferences:EventID" json:"events_id"`
	// Reviews []Review `gorm:"many2many:reviews;joinForeignKey:ReviewID;" json:"reviews"`
}
