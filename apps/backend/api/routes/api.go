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
	projectRepository := repository.NewProjectRepository(dbPool)

	userService := services.NewUserService(userRepository)
	projectService := services.NewProjectService(projectRepository)

	authController := controllers.NewAuthController(userService)
	projectController := controllers.NewProjectController(projectService)

	v1.POST("/auth/login", authController.Login)
	v1.POST("/auth/register", authController.Register)

	v1.POST("/projects", projectController.NewProject)
}
