package FavoriteUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type FavoriteUsecase interface {
	GetAll() ([]Entities.Favorite, error)
	Create(favorite *Entities.Favorite) error
	GetByID(id *uint) (*Entities.Favorite, error)
	Update(favorite *Entities.Favorite) error
	DeleteByID(id *uint) error
	GetByUserID(userID *uint) ([]Entities.Favorite, error)
}
type FavoriteService struct {
	FavoriteRepository FavoriteRepository
}
func NewFavoriteService(favoriteRepository FavoriteRepository) *FavoriteService {
	return &FavoriteService{
		FavoriteRepository: favoriteRepository,
	}
}
func (s *FavoriteService) GetAll() ([]Entities.Favorite, error) {
	favorites, err := s.FavoriteRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
func (s *FavoriteService) Create(favorite *Entities.Favorite) error {
	err := s.FavoriteRepository.Create(favorite)
	if err != nil {
		return err
	}
	return nil
}
func (s *FavoriteService) GetByID(id *uint) (*Entities.Favorite, error) {
	favorite, err := s.FavoriteRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}
func (s *FavoriteService) Update(favorite *Entities.Favorite) error {
	err := s.FavoriteRepository.Update(favorite)
	if err != nil {
		return err
	}
	return nil
}
func (s *FavoriteService) DeleteByID(id *uint) error {
	err := s.FavoriteRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *FavoriteService) GetByUserID(userID *uint) ([]Entities.Favorite, error) {
	favorites, err := s.FavoriteRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}