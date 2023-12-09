package repository

import (
	"github.com/JamshedJ/REST-api/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateTask(userID int, params models.TaskParams) (id int, err error) {
	return 
}

func (r *Repository) GetTaskByID(userID int, taskID int) (task models.Task, err error) {
	return 
}

func (r *Repository) GetTasks(userID int) (tasks []models.Task, err error) {
	return 
}

func (r *Repository) UpdateTask(userID int, params models.TaskParams) (task models.Task, err error) {
	return 
}

func (r *Repository) DeleteTask(userID int, params models.TaskParams) (id int, err error) {
	return 
}

func (r *Repository) MarkTask(userID int) (task models.Task, err error) {
	return 
}
