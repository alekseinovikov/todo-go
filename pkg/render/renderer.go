package render

import (
	. "todo-go/pkg/render/internal"
	"todo-go/pkg/service"
)

func NewRenderer(service service.TodoService) Renderer {
	return &PTermRenderer{Area: nil, Service: service, Selected: nil}
}

type Renderer interface {
	PrintLogo()
	PrepareDisplay()
}
