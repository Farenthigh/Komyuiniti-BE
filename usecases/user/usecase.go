package UserUsecase

import (
	"errors"

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
	return "User created", errors.New("")
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
