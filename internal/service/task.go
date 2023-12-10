package service

import (
	"context"

	"github.com/JamshedJ/REST-api/internal/models"
	"github.com/JamshedJ/REST-api/internal/repository"
)

type TaskService struct {
	repo repository.Repository
}

func NewTaskService(repo repository.Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, userID int, params models.TaskParams) (task *models.Task, err error) {
	return s.repo.CreateTask(ctx, userID, params)
}

func (s *TaskService) GetTaskByID(userID int, taskID int) (task models.Task, err error) {
	return s.repo.GetTaskByID(userID, taskID)
}

func (s *TaskService) GetTasks(userID int) (tasks []models.Task, err error) {
	return s.repo.GetTasks(userID)
}

func (s *TaskService) UpdateTask(userID int, params models.TaskParams) (task models.Task, err error) {
	return s.repo.UpdateTask(userID, params)
}

func (s *TaskService) DeleteTask(userID int, params models.TaskParams) (id int, err error) {
	return s.repo.DeleteTask(userID, params)
}

func (s *TaskService) MarkTask(userID int) (task models.Task, err error) {
	return s.repo.MarkTask(userID)
}