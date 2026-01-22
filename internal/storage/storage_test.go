package storage

import (
	"testing"

	"ecom-tt/internal/models"
)

func TestCreateTask(t *testing.T) {
	store := NewStorage()

	task := models.Todo{
		Title:       "Test task",
		Description: "Test description",
		Completed:   false,
	}

	created := store.CreateTask(task)

	if created.ID == 0 {
		t.Fatal("expected task ID to be set")
	}

	if created.Title != task.Title {
		t.Fatalf("expected title %q, got %q", task.Title, created.Title)
	}
}

func TestCreateTaskAutoIncrementID(t *testing.T) {
	store := NewStorage()

	task1 := store.CreateTask(models.Todo{Title: "Task 1"})
	task2 := store.CreateTask(models.Todo{Title: "Task 2"})

	if task2.ID != task1.ID+1 {
		t.Fatalf("expected auto increment ID, got %d and %d", task1.ID, task2.ID)
	}
}

func TestGetTaskByID(t *testing.T) {
	store := NewStorage()

	created := store.CreateTask(models.Todo{
		Title:     "Find me",
		Completed: false,
	})

	task, ok := store.GetByTaskID(created.ID)
	if !ok {
		t.Fatal("expected task to be found")
	}

	if task.Title != "Find me" {
		t.Fatalf("unexpected task title: %s", task.Title)
	}
}

func TestGetTaskByIDNotFound(t *testing.T) {
	store := NewStorage()

	_, ok := store.GetByTaskID(999)
	if ok {
		t.Fatal("expected task not to be found")
	}
}

func TestGetAllTasks(t *testing.T) {
	store := NewStorage()

	store.CreateTask(models.Todo{Title: "Task 1"})
	store.CreateTask(models.Todo{Title: "Task 2"})

	tasks := store.GetAlTasks()

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	store := NewStorage()

	task := store.CreateTask(models.Todo{Title: "Delete me"})

	ok := store.DeleteTask(task.ID)
	if !ok {
		t.Fatal("expected delete to succeed")
	}

	_, ok = store.GetByTaskID(task.ID)
	if ok {
		t.Fatal("task should be deleted")
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	store := NewStorage()

	ok := store.DeleteTask(123)
	if ok {
		t.Fatal("expected delete to fail for non-existing task")
	}
}
