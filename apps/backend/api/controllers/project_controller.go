package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tiagofs/bolttech-todo-list/api/services"
)

type ProjectController struct {
	projectService services.ProjectServiceInterface
}

type NewProjectRequest struct {
	ProjectName string `json:"project_name"`
}

func NewProjectController(projectService services.ProjectServiceInterface) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (pc *ProjectController) NewProject(c *gin.Context) {
	// Service -> repo
	var request NewProjectRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request")
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

	newProjectResponse, err := pc.projectService.NewProject(request.ProjectName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, newProjectResponse)
}
