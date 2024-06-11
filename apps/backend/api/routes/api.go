package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tiagofs/bolttech-todo-list/api/controllers"
	"github.com/tiagofs/bolttech-todo-list/api/repository"
	"github.com/tiagofs/bolttech-todo-list/api/services"
)

func SetupApiRoutes(app *gin.Engine, dbPool *pgxpool.Pool) {
	api := app.RouterGroup.Group("/api")
	v1 := api.Group("/v1")

	userRepository := repository.NewUserRepository(dbPool)

	userService := services.NewUserService(userRepository)

	authController := controllers.NewAuthController(userService)

	//TODO: AuthRequired middleware
	// v1.POST("/api/shortUrl", func(c *gin.Context) {
	// 	shortUrlController.CreateShortUrl(c)
	// })

	v1.POST("/auth/login", authController.Login)
	v1.POST("/auth/register", authController.Register)
}
