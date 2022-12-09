package main

import (
	"todo-go/pkg/render"
	"todo-go/pkg/repository"
	"todo-go/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	svs := service.NewService(repo)
	renderer := render.NewRenderer(svs)

	renderer.PrintLogo()
	renderer.PrepareDisplay()
}
