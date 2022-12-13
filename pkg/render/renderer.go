package render

import (
	. "todo-go/pkg/render/internal"
	"todo-go/pkg/service"
)

func NewRenderer(service service.TodoService) Renderer {
	return &TeaRenderer{Service: service, CursorPointer: 0}
}

type Renderer interface {
	Start()
}
