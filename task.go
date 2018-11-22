package todo

// Task is the basic data structures for TODO
type Task struct {
	title string
	done  bool
}

type TaskManager struct {
	tasks []*Task
}
