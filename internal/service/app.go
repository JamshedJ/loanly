package service

import (
	"context"

	"github.com/JamshedJ/REST-api/internal/models"
	"github.com/JamshedJ/REST-api/internal/repository"
)

type Authorization interface {
}

type Task interface {
	CreateTask(cxt context.Context, userID int, params models.TaskParams) (task *models.Task, err error)
	GetTaskByID(userID int, taskID int) (task models.Task, err error)
	GetTasks(userID int) (tasks []models.Task, err error)
	UpdateTask(userID int, params models.TaskParams) (task models.Task, err error)
	DeleteTask(userID int, params models.TaskParams) (id int, err error)
	MarkTask(userID int) (task models.Task, err error)
}

type Service struct {
	Authorization
	Task
	Repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Repo: repo,
		Task: NewTaskService(repo),
	}
}
