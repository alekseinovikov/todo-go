package internal

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"strings"
	. "todo-go/pkg/domain"
	"todo-go/pkg/service"
)

type PTermRenderer struct {
	Area     *pterm.AreaPrinter
	Service  service.TodoService
	Selected *Todo
}

func (P *PTermRenderer) PrintLogo() {
	s, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("T", pterm.NewRGB(255, 215, 0)),
		putils.LettersFromStringWithStyle("O", pterm.NewStyle(pterm.FgWhite)),
		putils.LettersFromStringWithRGB("D", pterm.NewRGB(255, 215, 0)),
		putils.LettersFromStringWithStyle("O", pterm.NewStyle(pterm.FgWhite)),
	).Srender()
	pterm.DefaultCenter.Println(s) // Print BigLetters with the default CenterPrinter
}

func (P *PTermRenderer) PrepareDisplay() {
	P.Area, _ = pterm.DefaultArea.WithCenter().Start()

	notCompleted, _ := P.Service.GetUncompleted()

	listPanel := P.sprintListPanel(notCompleted)
	detailsPanel := P.sprintDetailsPanel(P.Selected)
	mainPanel := P.sprintMainPanel(listPanel, detailsPanel)

	P.Area.Update(mainPanel)
}

func (P *PTermRenderer) updateArea(mainPanel string) {
	P.Area.Update(mainPanel)
}

func (P *PTermRenderer) sprintMainPanel(listPanel, detailsPanel string) string {
	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: listPanel}, {Data: detailsPanel}},
	}).Srender()

	return pterm.DefaultBox.WithTitle("Todos").
		WithTitleBottomRight().
		WithRightPadding(0).
		WithBottomPadding(0).
		Sprintln(panels)
}

func (P *PTermRenderer) sprintListPanel(todos []*Todo) string {
	strTodos := strings.Builder{}
	for _, todo := range todos {
		if P.Selected != nil && P.Selected.GetId() == todo.GetId() {
			strTodos.WriteString("> ")
		} else {
			strTodos.WriteString("- ")
		}
		strTodos.WriteString(todo.GetTitle())
		strTodos.WriteString("\n")
	}

	return pterm.DefaultBox.WithTitle("Not Completed").Sprint(strTodos.String())
}

func (P *PTermRenderer) sprintDetailsPanel(todo *Todo) string {
	if nil == todo {
		return pterm.DefaultBox.WithTitle("<None>").Sprint("No todo selected")
	}

	return pterm.DefaultBox.WithTitle(todo.GetTitle()).Sprint(todo.GetDescription())
}
