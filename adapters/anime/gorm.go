package AnimeAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"gorm.io/gorm"
)

type AnimeGorm struct {
	db *gorm.DB
}

func NewAnimeGorm(db *gorm.DB) *AnimeGorm {
	return &AnimeGorm{
		db: db,
	}
}
func (g *AnimeGorm) GetAll() ([]Entities.Anime, error) {
	var animes []Entities.Anime
	if err := g.db.Find(&animes).Error; err != nil {
		return nil, err
	}
	return animes, nil
}
func (g *AnimeGorm) GetByID(id *uint) (*Entities.Anime, error) {
	var anime Entities.Anime
	if err := g.db.Where("id = ?", id).First(&anime).Error; err != nil {
		return nil, err
	}
	return &anime, nil
}
func (g *AnimeGorm) Create(anime *Entities.Anime) error {
	if err := g.db.Create(&anime).Error; err != nil {
		return err
	}
	return nil
}
func (g *AnimeGorm) Update(anime *Entities.Anime) error {
	if err := g.db.Save(&anime).Error; err != nil {
		return err
	}
	return nil
}
func (g *AnimeGorm) DeleteByID(id *uint) error {
	if err := g.db.Delete(&Entities.Anime{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (g *AnimeGorm) GetByUserID(userID *uint) ([]Entities.Anime, error) {
	var animes []Entities.Anime
	if err := g.db.Where("user_id = ?", userID).Find(&animes).Error; err != nil {
		return nil, err
	}
	return animes, nil
}
