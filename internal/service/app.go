package service

import (
	"github.com/JamshedJ/REST-api/internal/repository"
)

type Authorization interface {
}

type Task interface {
}

type Service struct {
	Authorization
	Task
	Repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}
