package ReviewUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type ReviewUsecase interface {
	GetAll() ([]Entities.Review, error)
	GetByID(id *uint) (*Entities.Review, error)
	Create(review *Entities.Review) error
	Update(review *Entities.Review) error
	DeleteByID(id *uint) error
	GetByAnimeID(animeID *uint) ([]Entities.Review, error)
}

type ReviewService struct {
	ReviewRepository ReviewRepository
}

func NewReviewService(repo ReviewRepository) *ReviewService {
	return &ReviewService{
		ReviewRepository: repo,
	}
}

func (s *ReviewService) GetAll() ([]Entities.Review, error) {
	reviews, err := s.ReviewRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
func (s *ReviewService) GetByID(id *uint) (*Entities.Review, error) {
	review, err := s.ReviewRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return review, nil
}
func (s *ReviewService) Create(review *Entities.Review) error {
	err := s.ReviewRepository.Create(review)
	if err != nil {
		return err
	}
	return nil
}
func (s *ReviewService) Update(review *Entities.Review) error {
	err := s.ReviewRepository.Update(review)
	if err != nil {
		return err
	}
	return nil
}
func (s *ReviewService) DeleteByID(id *uint) error {
	err := s.ReviewRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *ReviewService) GetByAnimeID(animeID *uint) ([]Entities.Review, error) {
	reviews, err := s.ReviewRepository.GetByAnimeID(animeID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}