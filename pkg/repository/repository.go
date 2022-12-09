package repository

import (
	. "github.com/samber/mo"
	"sync"
	. "todo-go/pkg/domain"
	. "todo-go/pkg/repository/internal"
)

func NewRepository() TodoRepository {
	return &InMemoryTodoRepositoryImpl{
		Todos:  make(map[uint16]Todo),
		LastId: 0,
		Rw:     &sync.RWMutex{},
	}
}

type TodoRepository interface {
	GetById(id uint16) (Option[*Todo], error)
	GetCompleted() ([]*Todo, error)
	GetUncompleted() ([]*Todo, error)
	Add(todo AddTodo) (*Todo, error)
	MarkCompleted(id uint16) (Option[*Todo], error)
	MarkNotCompleted(id uint16) (Option[*Todo], error)
}
