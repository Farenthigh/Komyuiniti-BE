package UserAdapter

import (
	"errors"
	"fmt"

	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	UserUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/user"
	"gorm.io/gorm"
)

type UserGorm struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) UserUsecase.UserRepository {
	return &UserGorm{
		db: db,
	}
}

func (g *UserGorm) Register(user *Entities.User) error {
	if user.Email == "" {
		return errors.New("email cannot be null")
	}
	if user.UserName == "" {
		return errors.New("UserName cannot be null")
	}
	if user.Password == "" {
		return errors.New("password cannot be null")
	}

	if err := g.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("email %s already exists", user.Email)
		}
		return err
	}
	return nil
}

func (g *UserGorm) GetAll() ([]Entities.User, error) {
	var users []Entities.User
	if err := g.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (g *UserGorm) GetByEmail(email string) (*Entities.User, error) {
	var user Entities.User
	if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}
	return &user, nil
}

func (g *UserGorm) GetByID(id *uint) (*Entities.User, error) {
	var user Entities.User
	if err := g.db.Omit("password").Where("id = ?", &id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}
	return &user, nil
}
func (g *UserGorm) Update(id *uint, user *Entities.User) (*Entities.User, error) {
	if err := g.db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}
	return user, nil
}

func (g *UserGorm) GetMyTweets(userID *uint) ([]Entities.Tweet, error) {
	var tweets []Entities.Tweet
	if err := g.db.Where("author_id = ?", userID).Preload("Author").Find(&tweets).Error; err != nil {
		return nil, err
	}
	return tweets, nil
}

func (g *UserGorm) GetMyEvents(userID *uint) ([]Entities.Event, error) {
	var events []Entities.Event
	if err := g.db.Where("author_id = ?", userID).Preload("Author").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}