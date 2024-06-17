package services

import (
	"github.com/tiagofs/bolttech-todo-list/api/models"
	"github.com/tiagofs/bolttech-todo-list/api/repository"
)

type ProjectServiceInterface interface {
	NewProject(projectName string) (*models.Project, error)
}

type ProjectService struct {
	ProjectRepository repository.ProjectRepositoryInterface
}

func NewProjectService(projectRepository repository.ProjectRepositoryInterface) *ProjectService {
	return &ProjectService{
		ProjectRepository: projectRepository,
	}
}

func (s *ProjectService) NewProject(projectName string) (*models.Project, error) {
	createdProject, err := s.ProjectRepository.NewProject(projectName)

	if err != nil {
		return nil, err
	}
	return createdProject, nil
}
