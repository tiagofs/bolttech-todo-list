package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tiagofs/bolttech-todo-list/api/models"
	"github.com/tiagofs/bolttech-todo-list/api/services"
)

type AuthController struct {
	userService services.UserServiceInterface
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func NewAuthController(userService services.UserServiceInterface) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, errs)
			return
		}
		c.JSON(http.StatusUnprocessableEntity, "Error while validating the request data")
		return
	}

	loginResponse, err := ac.userService.Login(request.Email, request.Password)

	if err != nil {
		slog.Error(err.Error())

		if err.Error() == "user not found" || err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, "Invalid login")
			return
		}
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: loginResponse.Token,
		TokenType:   loginResponse.Type,
		ExpiresIn:   loginResponse.ExpiresIn,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	var request RegisterRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, errs)
			return
		}
		c.JSON(http.StatusUnprocessableEntity, "Error while validating the request data")
		return
	}

	newUser := &models.User{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = ac.userService.AddUser(newUser)

	if err != nil {
		slog.Error(err.Error())

		if err.Error() == "user not found" || err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, "Invalid login")
			return
		}
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, "Successfully registered the new user")
}
