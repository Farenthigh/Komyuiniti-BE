package AnimeUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type AnimeRepositorty interface {
	GetAll() ([]Entities.Anime, error)
	GetByID(id *uint) (*Entities.Anime, error)
	Create(anime *Entities.Anime) error
	Update(anime *Entities.Anime) error
	DeleteByID(id *uint) error
	GetByUserID(userID *uint) ([]Entities.Anime, error)
}
