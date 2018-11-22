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

func iterateOverTasks(allTasks []*Task, task *Task) bool {
	condition := false
	length := len(allTasks)

	for i := 0; i < length; i++ {
		t := allTasks[i]
		if (*t).title == (*task).title {
			condition = true
		}
	}

	return condition
}

func (m *TaskManager) Save(task *Task) {
	condition := iterateOverTasks(m.All(), task)
	if !condition {
		copy := *task
		m.tasks = append(m.tasks, &copy)
	}
}

func (m *TaskManager) All() []*Task {
	return m.tasks
}
