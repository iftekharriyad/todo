package todo_test

import (
	"os"
	"testing"

	"github.com/iftekharriyad/todo"
)

// TestAdd tests Add method of the List Type
func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expeced %s, got %s", taskName, l[0].Task)
	}
}

// TestAdd tests Add method of the List Type
func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expeced %s, got %s", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("%s should not be completed", taskName)
	}
	l[0].Done = true
	if !l[0].Done {
		t.Errorf("%s should be completed", taskName)
	}
}

// TestDelete tests Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
		"Task 4",
		"Task 5",
	}
	for _, t := range tasks {
		l.Add(t)
	}
	if len(l) != len(tasks) {
		t.Errorf("Expected %d tasks, got %d tasks.", len(tasks), len(l))
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %s, got %s .", tasks[0], l[0].Task)
	}
	l.Delete(3)
	if len(l) != len(tasks)-1 {
		t.Errorf("Expected %d tasks, got %d tasks.", len(tasks)-1, len(l))
	}
	if l[2].Task == tasks[2] {
		t.Errorf("Expected %s to be deleted", l[2].Task)
	}
}

// TestSaveGet tests Save and Get method of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "Task 1"
	l1.Add(taskName)
	if l1[0].Task != taskName {
		t.Errorf("Expeced %s, got %s", taskName, l1[0].Task)
	}
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tempFile.Name())
	if err := l1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving list to the file: %s", err)
	}
	if err := l2.Get(tempFile.Name()); err != nil {
		t.Fatalf("Error getting the list from the file: %s", err)
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %s should match task %s.", l1[0].Task, l2[0].Task)
	}
}
