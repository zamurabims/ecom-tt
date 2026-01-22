package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"ecom-tt/internal/models"
	"ecom-tt/internal/storage"
)

type TodoHandler struct {
	storage *storage.TodoStorageStruct
}

func NewTodoHandler(storage *storage.TodoStorageStruct) *TodoHandler {
	return &TodoHandler{storage: storage}
}

func (h *TodoHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		todos := h.storage.GetAlTasks()
		completedParam := r.URL.Query().Get("completed")
		if completedParam != "" {
			completed, err := strconv.ParseBool(completedParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			filtered := make([]models.Todo, 0)
			for _, todo := range todos {
				if todo.Completed == completed {
					filtered = append(filtered, todo)
				}
			}
			todos = filtered
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()
		var todo models.Todo

		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			slog.Warn("failed to create task", slog.Any("error", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(todo.Title) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				slog.Warn("request timeout", slog.Any("error", ctx.Err()))
				w.WriteHeader(http.StatusRequestTimeout)
			case context.Canceled:
				slog.Info("user cancelled a request")

			}
			return
		default:

		}
		created := h.storage.CreateTask(todo)

		slog.Info("task created", slog.Int("id", created.ID))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) TasksByID(w http.ResponseWriter, r *http.Request) {
	const prefix = "/todos/"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	idStr = strings.Trim(idStr, "/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		todo, ok := h.storage.GetByTaskID(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)

	case http.MethodPut:
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()
		var todo models.Todo

		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			slog.Warn("failed to decode todo", slog.Any("error", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(todo.Title) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				slog.Warn("request timeout", slog.Any("error", ctx.Err()))
				w.WriteHeader(http.StatusRequestTimeout)
			case context.Canceled:
				slog.Info("user cancelled a request")

			}
			return
		default:

		}

		updated, ok := h.storage.UpdateTask(id, todo)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		slog.Info("task updated", slog.Int("id", id))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				slog.Warn("request timeout", slog.Any("error", ctx.Err()))
				w.WriteHeader(http.StatusRequestTimeout)
			case context.Canceled:
				slog.Info("user cancelled a request")

			}
			return
		default:

		}
		if !h.storage.DeleteTask(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		slog.Info("task deleted", slog.Int("id", id))
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
