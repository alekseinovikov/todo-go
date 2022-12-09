package internal

import (
	. "github.com/samber/mo"
	"sync"
	. "todo-go/pkg/domain"
)

type InMemoryTodoRepositoryImpl struct {
	Rw     *sync.RWMutex
	LastId uint16
	Todos  map[uint16]Todo
}

func (t *InMemoryTodoRepositoryImpl) GetById(id uint16) (Option[*Todo], error) {
	t.Rw.RLock()
	defer t.Rw.RUnlock()

	if todo, ok := t.Todos[id]; ok {
		return Some(&todo), nil
	}

	return None[*Todo](), nil
}

func (t *InMemoryTodoRepositoryImpl) GetCompleted() ([]*Todo, error) {
	t.Rw.RLock()
	defer t.Rw.RUnlock()

	result := make([]*Todo, 0)
	for _, v := range t.Todos {
		if v.Completed {
			result = append(result, &v)
		}
	}
	return result, nil
}

func (t *InMemoryTodoRepositoryImpl) GetUncompleted() ([]*Todo, error) {
	t.Rw.RLock()
	defer t.Rw.RUnlock()

	result := make([]*Todo, 0)
	for _, v := range t.Todos {
		if !v.Completed {
			result = append(result, &v)
		}
	}
	return result, nil
}

func (t *InMemoryTodoRepositoryImpl) Add(todo AddTodo) (*Todo, error) {
	t.Rw.Lock()
	defer t.Rw.Unlock()

	t.LastId++
	newTodo := Todo{
		Id:          t.LastId,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
	}

	t.Todos[t.LastId] = newTodo
	return &newTodo, nil
}

func (t *InMemoryTodoRepositoryImpl) MarkCompleted(id uint16) (Option[*Todo], error) {
	t.Rw.Lock()
	defer t.Rw.Unlock()

	if todo, ok := t.Todos[id]; ok {
		todo.Completed = true
		return Some(&todo), nil
	}

	return None[*Todo](), nil
}

func (t *InMemoryTodoRepositoryImpl) MarkNotCompleted(id uint16) (Option[*Todo], error) {
	t.Rw.Lock()
	defer t.Rw.Unlock()

	if todo, ok := t.Todos[id]; ok {
		todo.Completed = false
		return Some(&todo), nil
	}

	return None[*Todo](), nil
}
