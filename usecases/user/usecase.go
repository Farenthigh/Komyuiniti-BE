package UserUsecase

import (
	"errors"
	"fmt"

	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	UserModels "github.com/Farenthigh/Fitbuddy-BE/model/user"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(*UserModels.RegisterInput) (string, error)
	GetAll() ([]Entities.User, error)
	Login(*UserModels.LoginInput) (string, error)
	Me(userID *uint) (*Entities.User, error)
	Update(userID *uint, user *Entities.User) (*Entities.User, error)
	GetMyTweets(userID *uint) ([]Entities.Tweet, error)
	GetMyEvents(userID *uint) ([]Entities.Event, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserUsecase {
	return &UserService{
		userRepo: userRepo,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (service *UserService) Register(user *UserModels.RegisterInput) (string, error) {
	if user.Password != user.ConfirmPassword {
		return "Password and Confirm Password must be the same", errors.New("password and confirm password must be the same")
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return "Internal server error", err
	}
	var userEntity Entities.User
	userEntity.Email = user.Email
	userEntity.UserName = user.UserName
	userEntity.Password = hashedPassword
	if err := service.userRepo.Register(&userEntity); err != nil {
		return "Internal server error", err
	}
	return "User created", nil
}

func (service *UserService) GetAll() ([]Entities.User, error) {
	users, err := service.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) Login(user *UserModels.LoginInput) (string, error) {
	var userEntity Entities.User
	userEntity.Email = user.Email
	userEntity.Password = user.Password

	// Check if the user exists
	existingUser, err := service.userRepo.GetByEmail(userEntity.Email)
	if err != nil {
		return "Email not found", err
	}

	// Compare the password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userEntity.Password))
	if err != nil {
		return "Invalid password", err
	}
	token, err := utils.CreateToken(existingUser.ID, existingUser.Email, existingUser.UserName)
	if err != nil {
		return "Failed to generate token", err
	}

	return token, nil
}

func (service *UserService) Me(userID *uint) (*Entities.User, error) {
	user, err := service.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
func (service *UserService) Update(userID *uint, user *Entities.User) (*Entities.User, error) {
	fmt.Println(user.UserName)
	userEntity, err := service.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if userEntity == nil {
		return nil, errors.New("user not found")
	}
	userEntity.UserName = user.UserName
	userEntity.Email = user.Email
	userEntity.UserImage = user.UserImage
	return service.userRepo.Update(userID, userEntity)
}

func (service *UserService) GetMyTweets(userID *uint) ([]Entities.Tweet, error) {
	tweets, err := service.userRepo.GetMyTweets(userID)
	if err != nil {
		return nil, errors.New("failed to get tweets")
	}
	return tweets, nil
}

func (service *UserService) GetMyEvents(userID *uint) ([]Entities.Event, error) {
	events, err := service.userRepo.GetMyEvents(userID)
	if err != nil {
		return nil, errors.New("failed to get events")
	}
	return events, nil
}