package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tiagofs/bolttech-todo-list/api/models"
)

type ProjectRepositoryInterface interface {
	NewProject(projectName string) (*models.Project, error)
}

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	if db == nil {
		return nil
	}

	return &ProjectRepository{db: db}
}

func (repo *ProjectRepository) NewProject(projectName string) (*models.Project, error) {

	query := `INSERT INTO projects (project_name) VALUES ($1)`

	repo.db.QueryRow(context.Background(), query, projectName)
	// if row != nil {
	// slog.Error(err.Error())
	// return nil, err
	// }

	project := &models.Project{
		ProjectName: projectName,
	}

	return project, nil
}
