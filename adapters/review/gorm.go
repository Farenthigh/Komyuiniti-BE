package ReviewAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"gorm.io/gorm"
)

type ReviewGorm struct {
	db *gorm.DB
}

func NewReviewGorm(db *gorm.DB) *ReviewGorm {
	return &ReviewGorm{
		db: db,
	}
}
func (g *ReviewGorm) GetAll() ([]Entities.Review, error) {
	var reviews []Entities.Review
	if err := g.db.Preload("Anime").Preload("Author").Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (g *ReviewGorm) GetByID(id *uint) (*Entities.Review, error) {
	var review Entities.Review
	if err := g.db.Preload("Author").Where("id = ?", id).First(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (g *ReviewGorm) Create(review *Entities.Review) error {
	if err := g.db.Create(&review).Error; err != nil {
		return err
	}
	return nil
}
func (g *ReviewGorm) Update(review *Entities.Review) error {
	if err := g.db.Save(&review).Error; err != nil {
		return err
	}
	return nil
}
func (g *ReviewGorm) DeleteByID(id *uint) error {
	if err := g.db.Delete(&Entities.Review{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *ReviewGorm) GetByAnimeID(animeID *uint) ([]Entities.Review, error) {
	var reviews []Entities.Review
	if err := g.db.Model("Review").Preload("Anime").Preload("Author").Where("anime_id = ?", animeID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}