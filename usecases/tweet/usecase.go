package tweetUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type TweetUsecase interface {
	GetAll() ([]Entities.Tweet, error)
	GetByID(id *uint) (*Entities.Tweet, error)
	Create(tweet *Entities.Tweet) error
	Update(tweet *Entities.Tweet) error
	DeleteByID(id *uint) error
}

type TweetService struct {
	tweetRepo TweetRepository
}

func NewTweetService(tweetRepo TweetRepository) TweetUsecase {
	return &TweetService{
		tweetRepo: tweetRepo,
	}
}
func (service *TweetService) GetAll() ([]Entities.Tweet, error) {
	tweets, err := service.tweetRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
func (service *TweetService) GetByID(id *uint) (*Entities.Tweet, error) {
	tweet, err := service.tweetRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
func (service *TweetService) Create(tweet *Entities.Tweet) error {
	if err := service.tweetRepo.Create(tweet); err != nil {
		return err
	}
	return nil
}
func (service *TweetService) Update(tweet *Entities.Tweet) error {
	selectedTweet, err := service.tweetRepo.GetByID(&tweet.ID)
	if err != nil {
		return err
	}
	if selectedTweet == nil {
		return nil
	}
	selectedTweet.Description = tweet.Description

	if err := service.tweetRepo.Update(selectedTweet); err != nil {
		return err
	}
	return nil
}
func (service *TweetService) DeleteByID(id *uint) error {
	if err := service.tweetRepo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
