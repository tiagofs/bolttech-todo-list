package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tiagofs/bolttech-todo-list/api/config"
	"github.com/tiagofs/bolttech-todo-list/api/models"
	"github.com/tiagofs/bolttech-todo-list/api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	FindByID(id string) (*models.User, error)
	Login(email, password string) (*JwtToken, error)
	AddUser(user *models.User) (*models.User, error)
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

type JwtToken struct {
	Token     string
	Type      string
	ExpiresIn int64
}

type Claims struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) Login(email, password string) (*JwtToken, error) {

	config := config.GetConfig()

	var jwtKey = []byte(config.Jwt.JwtSecret)

	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Compare password and hash password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	expirationTime := time.Now().Add(time.Duration(config.Jwt.ExpiresIn) * time.Minute)

	claims := &Claims{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	return &JwtToken{
		Token:     tokenString,
		Type:      "bearer",
		ExpiresIn: expirationTime.Unix() - time.Now().Unix(),
	}, nil
}

func (s *UserService) FindByID(id string) (*models.User, error) {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil

}

func (s *UserService) AddUser(user *models.User) (*models.User, error) {
	// TODO: validate data, check e-mail already register
	createdUser, err := s.UserRepository.AddUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
