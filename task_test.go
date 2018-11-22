package todo

import "testing"

func TestNewTask(t *testing.T) {
	var testEqual = func(got, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' \n", got, want)
		}
	}

	// testing
	t.Run("Create a new task", func(t *testing.T) {
		got := NewTask("learn go")
		want := &Task{title: "learn go"}
		testEqual(got.title, want.title, t)
	})

}
