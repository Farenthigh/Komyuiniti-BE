package CommentUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type CommentRepository interface {
	GetAll() ([]Entities.Comment, error)
	GetByID(id *uint) (*Entities.Comment, error)
	Create(comment *Entities.Comment) error
	Update(comment *Entities.Comment) error
	DeleteByID(id *uint) error
}
