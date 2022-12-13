package repository

import (
	. "github.com/samber/mo"
	"strconv"
	"sync"
	. "todo-go/pkg/domain"
	. "todo-go/pkg/repository/internal"
)

func NewRepository() TodoRepository {
	todos := make(map[uint16]Todo)

	for i := 1; i <= 10; i++ {
		todos[uint16(i)] = Todo{
			Id:          uint16(i),
			Title:       "Title " + strconv.Itoa(i),
			Description: "Description " + strconv.Itoa(i),
			Completed:   false,
		}
	}

	return &InMemoryTodoRepositoryImpl{
		Todos:  todos,
		LastId: 0,
		Rw:     &sync.RWMutex{},
	}
}

type TodoRepository interface {
	GetById(id uint16) (Option[Todo], error)
	GetCompleted() ([]Todo, error)
	GetUncompleted() ([]Todo, error)
	Add(todo AddTodo) (Todo, error)
	MarkCompleted(id uint16) (Option[Todo], error)
	MarkNotCompleted(id uint16) (Option[Todo], error)
}
