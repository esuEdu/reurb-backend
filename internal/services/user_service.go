package services

import (
	"errors"

	"github.com/esuEdu/reurb-backend/internal/models"
	"github.com/esuEdu/reurb-backend/internal/repositories"
	"github.com/esuEdu/reurb-backend/internal/util"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserService interface {
	RegisterUser(name, email, password string) (*models.User, error)
	AuthenticateUser(email, password string) (string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (service *userService) RegisterUser(name, email, password string) (*models.User, error) {
	existingUser, _ := service.repo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	stringHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: stringHash,
	}

	createdUser, err := service.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil

}

func (service *userService) AuthenticateUser(email, password string) (string, error) {
	user, err := service.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = checkPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("wrong credential")
	}

	token, err := util.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		return "", errors.New("failed creating token")
	}

	return token, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
