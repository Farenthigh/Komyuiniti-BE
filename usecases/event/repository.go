package EventUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type Eventrepository interface {
	GetAll() ([]Entities.Event, error)
	GetByID(id *uint) (*Entities.Event, error)
	Create(event *Entities.Event) error
	Update(event *Entities.Event) error
	DeleteByID(id *uint) error
	GetByUserID(userID *uint) ([]Entities.Event, error)
	JoinEvent(eventID *uint, UserID *uint) error
}
