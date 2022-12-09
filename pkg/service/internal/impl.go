package internal

import (
	. "github.com/samber/mo"
	. "todo-go/pkg/domain"
	"todo-go/pkg/repository"
)

type TodoServiceImpl struct {
	Repo repository.TodoRepository
}

func (t TodoServiceImpl) GetById(id uint16) (Option[*Todo], error) {
	return t.Repo.GetById(id)
}

func (t TodoServiceImpl) GetCompleted() ([]*Todo, error) {
	return t.Repo.GetCompleted()
}

func (t TodoServiceImpl) GetUncompleted() ([]*Todo, error) {
	return t.Repo.GetUncompleted()
}

func (t TodoServiceImpl) Add(todo AddTodo) (*Todo, error) {
	return t.Repo.Add(todo)
}

func (t TodoServiceImpl) MarkCompleted(id uint16) (Option[*Todo], error) {
	return t.Repo.MarkCompleted(id)
}

func (t TodoServiceImpl) MarkNotCompleted(id uint16) (Option[*Todo], error) {
	return t.Repo.MarkNotCompleted(id)
}
