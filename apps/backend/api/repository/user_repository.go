package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tiagofs/bolttech-todo-list/api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
	Delete(id string) (bool, error)
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	if db == nil {
		return nil
	}
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddUser(user *models.User) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	query := `INSERT INTO users (email, password, first_name, last_name, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var userID string
	err = repo.db.QueryRow(context.Background(), query, user.Email, user.Password, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return repo.FindByID(userID)
}

func (repo *UserRepository) FindByID(id string) (*models.User, error) {
	user := &models.User{}
	err := repo.db.QueryRow(context.Background(), "SELECT id, email, password, first_name, last_name, created_at, updated_at FROM users WHERE id = $1  and deleted_at is null", id).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := repo.db.QueryRow(context.Background(), "SELECT id, email, password, first_name, last_name FROM users WHERE email = $1 and deleted_at is null", email).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) Delete(id string) (bool, error) {
	_, err := repo.db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
