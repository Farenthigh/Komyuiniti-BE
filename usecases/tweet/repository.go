package tweetUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type TweetRepository interface {
	GetAll() ([]Entities.Tweet, error)
	GetByID(id *uint) (*Entities.Tweet, error)
	Create(forum *Entities.Tweet) error
	Update(forum *Entities.Tweet) error
	DeleteByID(id *uint) error
}
