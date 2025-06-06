package ReviewUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type ReviewRepository interface {
	GetAll() ([]Entities.Review, error)
	GetByID(id *uint) (*Entities.Review, error)
	Create(review *Entities.Review) error
	Update(review *Entities.Review) error
	DeleteByID(id *uint) error
	GetByAnimeID(animeID *uint) ([]Entities.Review, error)
}
