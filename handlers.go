package todo

import "fmt"

// NewTask Creates a new Task
func NewTask(title string) (*Task, error) {

	if len(title) == 0 {
		return nil, fmt.Errorf("Empty Title")
	}

	return &Task{title: title, done: false}, nil
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Save(task *Task) {
	copy := *task
	m.tasks = append(m.tasks, &copy)
}

func (m *TaskManager) All() []*Task {
	return m.tasks
}
