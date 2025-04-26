package UserUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type UserRepository interface {
	Register(*Entities.User) error
	GetAll() ([]Entities.User, error)
	GetByEmail(email string) (*Entities.User, error)
	GetByID(id *uint) (*Entities.User, error)
}
