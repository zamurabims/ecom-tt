package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ecom-tt/internal/models"
	"ecom-tt/internal/storage"
)

func TestCreateTaskSuccess(t *testing.T) {
	store := storage.NewStorage()
	handler := NewTodoHandler(store)

	body := []byte(`{
		"title": "Napisat testovoe",
		"description": "Vipolnit trebovaniya",
		"completed": false
	}`)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Tasks(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
}

func TestCreateTaskValidationError(t *testing.T) {
	store := storage.NewStorage()
	handler := NewTodoHandler(store)

	body := []byte(`{
		"title": "",
		"description": "Invalid task",
		"completed": false
	}`)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Tasks(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestCreateTaskInvalidJSON(t *testing.T) {
	store := storage.NewStorage()
	handler := NewTodoHandler(store)

	body := []byte(`{ invalid json }`)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Tasks(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestGetTasks(t *testing.T) {
	store := storage.NewStorage()
	handler := NewTodoHandler(store)

	store.CreateTask(models.Todo{
		Title:     "Task 1",
		Completed: false,
	})
	store.CreateTask(models.Todo{
		Title:     "Task 2",
		Completed: false,
	})

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rr := httptest.NewRecorder()

	handler.Tasks(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Task 1") || !strings.Contains(body, "Task 2") {
		t.Fatal("response does not contain created tasks")
	}
}
