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
	if err := g.db.Preload("Reviews").Find(&animes).Error; err != nil {
		return nil, err
	}
	// Calculate the average rating for each anime
	for i, anime := range animes {
		var totalRating float64
		for _, review := range anime.Reviews {
			totalRating += float64(review.Rating)
		}
		if len(anime.Reviews) > 0 {
			animes[i].Rating = totalRating / float64(len(anime.Reviews))
		} else {
			animes[i].Rating = 0
		}
	}

	return animes, nil
}
func (g *AnimeGorm) GetByID(id *uint) (*Entities.Anime, error) {
	var anime Entities.Anime
	if err := g.db.Preload("Reviews").Preload("Reviews.Author").Where("id = ?", id).First(&anime).Error; err != nil {
		return nil, err
	}
	// Calculate the average rating
	var totalRating float64
	for _, review := range anime.Reviews {
		totalRating += float64(review.Rating)
	}
	if len(anime.Reviews) > 0 {
		anime.Rating = totalRating / float64(len(anime.Reviews))
	} else {
		anime.Rating = 0
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

