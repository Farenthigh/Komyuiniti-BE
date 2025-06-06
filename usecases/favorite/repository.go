package FavoriteUsecase

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
)

type FavoriteRepository interface {
	GetAll() ([]Entities.Favorite, error)
	Create(favorite *Entities.Favorite) error
	GetByID(id *uint) (*Entities.Favorite, error)
	Update(favorite *Entities.Favorite) error
	DeleteByID(id *uint) error
	GetByUserID(userID *uint) ([]Entities.Favorite, error)
}