package entities

type Todo struct {
	Id        string
	Label     string
	Completed bool
}

func NewTodo(label string) *Todo {
	return &Todo{Label: label}
}
