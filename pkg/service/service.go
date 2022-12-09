package service

import (
	. "todo-go/pkg/repository"
	. "todo-go/pkg/service/internal"
)

func NewService(repo TodoRepository) TodoService {
	return &TodoServiceImpl{
		Repo: repo,
	}
}

type TodoService interface {
	TodoRepository
}
