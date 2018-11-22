package todo

// NewTask Creates a new Task
func NewTask(title string) *Task {
	return &Task{title: title, done: false}
}
