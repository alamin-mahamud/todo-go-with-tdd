package todo

import "testing"

func TestNewTask(t *testing.T) {
	var testEqualString = func(got, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' \n", got, want)
		}
	}

	var testEqualBool = func(got, want bool, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got '%t' want '%t' \n", got, want)
		}
	}

	t.Run("Create a new task", func(t *testing.T) {
		got, _ := NewTask("learn go")
		want := &Task{title: "learn go"}
		testEqualString(got.title, want.title, t)
	})

	t.Run("new tasks should not be done.", func(t *testing.T) {
		got, _ := NewTask("learn go")
		want := &Task{title: "learn go", done: false}
		testEqualBool(got.done, want.done, t)
	})

	t.Run("Empty Title are not ACCEPTABLE", func(t *testing.T) {
		_, err := NewTask("")
		if err == nil {
			t.Errorf("task with empty title created.")
		}
	})

	t.Run("Need a Task Manager to Save and Retrieve Task", func(t *testing.T) {
		task, err := NewTask("New Task")

		if err != nil {
			t.Errorf("Could not create new task %v", err)
		}

		taskManager := NewTaskManager()
		taskManager.Save(task)

		all := taskManager.All()

		if len(all) != 1 {
			t.Errorf("expected 1 task, got %v", len(all))
		}

		if *all[0] != *task {
			t.Errorf("expected %v, got %v", task, all)
		}

		newTask, err := NewTask("New tasks")

		if err != nil {
			t.Errorf("Could not create new task -> %v", err)
		}

		taskManager.Save(newTask)
		all = taskManager.All()

		if len(all) != 2 {
			t.Errorf("expected 2 task, got %v", len(all))
		}

		if *all[1] != *newTask {
			t.Errorf("expected %v, got %v", newTask, all)
		}
	})

	t.Run("Modifying after saving doesn't change saved copy", func(t *testing.T) {
		task, _ := NewTask("New Task For modifying")

		taskManager := NewTaskManager()
		taskManager.Save(task)

		all := taskManager.All()
		task2 := all[0]
		task.title = "Changed"
		if task2.title == task.title {
			t.Errorf("modification after save won't have any effect")
		}
	})

	t.Run("test save twice should not duplicate data", func(t *testing.T) {
		task, _ := NewTask("New Task")
		m := NewTaskManager()
		m.Save(task)
		m.Save(task)
		all := m.All()
		if len(all) != 1 {
			t.Errorf("Save one task only once")
		}
	})
}
