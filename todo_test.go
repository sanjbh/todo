package todo_test

import (
	"os"
	"testing"

	"github.com/sanjbh/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}

}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

	if l[0].Done == true {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)

	if l[0].Done != true {
		t.Errorf("New task should be completed")
	}

}

func TestDelete(t *testing.T) {
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	l := todo.List{}

	for _, t := range tasks {
		l.Add(t)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "test-case-*")
	if err != nil {
		t.Errorf("Error creating temp file: %v", err)
	}
	defer os.Remove(tf.Name())

	if err = l1.Save(tf.Name()); err != nil {
		t.Errorf("Error saving list to file %s", tf.Name())
	}

	if err = l2.Get(tf.Name()); err != nil {
		t.Errorf("Error getting list from file %s", tf.Name())
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
