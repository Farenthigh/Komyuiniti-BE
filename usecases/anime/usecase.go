package AnimeUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type AnimeUsecase interface {
	GetAll() ([]Entities.Anime, error)
	GetByID(id *uint) (*Entities.Anime, error)
	Create(anime *Entities.Anime) error
	Update(anime *Entities.Anime) error
	DeleteByID(id *uint) error
	GetByUserID(userID *uint) ([]Entities.Anime, error)
}
type AnimeService struct {
	animeRepo AnimeRepositorty
}

func NewAnimeService(animeRepo AnimeRepositorty) AnimeUsecase {
	return &AnimeService{
		animeRepo: animeRepo,
	}
}
func (service *AnimeService) GetAll() ([]Entities.Anime, error) {
	animes, err := service.animeRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return animes, nil
}
func (service *AnimeService) GetByID(id *uint) (*Entities.Anime, error) {
	anime, err := service.animeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return anime, nil
}
func (service *AnimeService) Create(anime *Entities.Anime) error {
	if err := service.animeRepo.Create(anime); err != nil {
		return err
	}
	return nil
}
func (service *AnimeService) Update(anime *Entities.Anime) error {
	selectedAnime, err := service.animeRepo.GetByID(&anime.ID)
	if err != nil {
		return err
	}
	if selectedAnime == nil {
		return nil
	}
	selectedAnime.Title = anime.Title
	selectedAnime.Description = anime.Description

	if err := service.animeRepo.Update(selectedAnime); err != nil {
		return err
	}
	return nil
}
func (service *AnimeService) DeleteByID(id *uint) error {
	if err := service.animeRepo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
func (service *AnimeService) GetByUserID(userID *uint) ([]Entities.Anime, error) {
	animes, err := service.animeRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return animes, nil
}
