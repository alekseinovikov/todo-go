package internal

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	. "todo-go/pkg/domain"
	"todo-go/pkg/service"
)

type TeaRenderer struct {
	Service       service.TodoService
	CursorPointer int
	Selected      *Todo
	NotCompleted  []Todo
}

func (R *TeaRenderer) Start() {
	p := tea.NewProgram(R)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (R *TeaRenderer) Init() tea.Cmd {
	R.loadNotCompleted()
	R.updateSelected()

	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (R *TeaRenderer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return R, tea.Quit

		case "up", "k":
			if R.CursorPointer > 0 {
				R.CursorPointer--
			}

		case "down", "j":
			if R.CursorPointer < len(R.NotCompleted)-1 {
				R.CursorPointer++
			}

		case "enter", " ":
			R.Selected = &R.NotCompleted[R.CursorPointer]
		}
	}

	return R, nil
}

func (R *TeaRenderer) View() string {
	s := "Your todos:\n\n"

	for i, choice := range R.NotCompleted {

		cursor := " "
		if R.CursorPointer == i {
			cursor = ">"
		}

		if R.Selected != nil && R.Selected.Id == choice.Id {
			cursor = "X"
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func (R *TeaRenderer) loadNotCompleted() {
	notCompleted, _ := R.Service.GetUncompleted()
	R.NotCompleted = notCompleted
}

func (R *TeaRenderer) updateSelected() {
	if R.CursorPointer == 0 && len(R.NotCompleted) <= 0 {
		R.CursorPointer = -1
	}
}
