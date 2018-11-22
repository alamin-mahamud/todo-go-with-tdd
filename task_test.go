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
		got := NewTask("learn go")
		want := &Task{title: "learn go"}
		testEqualString(got.title, want.title, t)
	})

	t.Run("new tasks should not be done.", func(t *testing.T) {
		got := NewTask("learn go")
		want := &Task{title: "learn go", done: false}
		testEqualBool(got.done, want.done, t)
	})
}
