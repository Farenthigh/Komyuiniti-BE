package FavoriteAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"gorm.io/gorm"
)

type FavoriteGorm struct {
	db gorm.DB
}

func NewFavoriteGorm(db *gorm.DB) *FavoriteGorm {
	return &FavoriteGorm{
		db: *db,
	}
}
func (g *FavoriteGorm) GetAll() ([]Entities.Favorite, error) {
	var favorites []Entities.Favorite
	if err := g.db.Preload("Tweets").Preload("Events").Preload("Reviews").Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}
func (g *FavoriteGorm) GetByID(id *uint) (*Entities.Favorite, error) {
	var favorite Entities.Favorite
	if err := g.db.Preload("Tweets").Preload("Events").Preload("Reviews").Where("id = ?", id).First(&favorite).Error; err != nil {
		return nil, err
	}
	return &favorite, nil
}
func (g *FavoriteGorm) Create(favorite *Entities.Favorite) error {
	if err := g.db.Create(&favorite).Error; err != nil {
		return err
	}
	return nil
}
func (g *FavoriteGorm) Update(favorite *Entities.Favorite) error {
	if err := g.db.Save(&favorite).Error; err != nil {
		return err
	}
	return nil
}
func (g *FavoriteGorm) DeleteByID(id *uint) error {
	if err := g.db.Delete(&Entities.Favorite{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (g *FavoriteGorm) GetByUserID(userID *uint) ([]Entities.Favorite, error) {
	var favorites []Entities.Favorite
	if err := g.db.Preload("Tweets").Preload("Events").Preload("Reviews").Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}