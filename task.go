package todo

// Task is the basic data structures for TODO
type Task struct {
	title string
}

// NewTask Creates a new Task
func NewTask(title string) *Task {
	return &Task{title: title}
}
