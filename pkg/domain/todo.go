package domain

type Todo struct {
	Id          uint16
	Title       string
	Description string
	Completed   bool
}

type AddTodo struct {
	Title       string
	Description string
}

func (t *Todo) GetId() uint16 {
	return t.Id
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) GetDescription() string {
	return t.Description
}
