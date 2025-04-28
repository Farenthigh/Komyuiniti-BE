package tweetAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"gorm.io/gorm"
)

type TweetGorm struct {
	db *gorm.DB
}

func NewTweetGorm(db *gorm.DB) *TweetGorm {
	return &TweetGorm{
		db: db,
	}
}
func (g *TweetGorm) GetAll() ([]Entities.Tweet, error) {
	var forums []Entities.Tweet
	if err := g.db.Preload("Author").Find(&forums).Error; err != nil {
		return nil, err
	}
	return forums, nil
}
func (g *TweetGorm) GetByID(id *uint) (*Entities.Tweet, error) {
	var forum Entities.Tweet
	if err := g.db.Preload("Author").Where("id = ?", id).First(&forum).Error; err != nil {
		return nil, err
	}
	return &forum, nil
}
func (g *TweetGorm) Create(tweet *Entities.Tweet) error {
	if err := g.db.Create(&tweet).Error; err != nil {
		return err
	}
	return nil
}
func (g *TweetGorm) Update(tweet *Entities.Tweet) error {
	if err := g.db.Save(&tweet).Error; err != nil {
		return err
	}
	return nil
}
func (g *TweetGorm) DeleteByID(id *uint) error {
	if err := g.db.Delete(&Entities.Tweet{}, id).Error; err != nil {
		return err
	}
	return nil
}
