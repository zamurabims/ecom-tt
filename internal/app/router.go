package app

import (
	"ecom-tt/internal/handlers"
	"net/http"
)

func getRouter(h *handlers.TodoHandler) *http.ServeMux {
	app := http.NewServeMux()

	app.HandleFunc("/todos", h.Tasks)
	app.HandleFunc("/todos/", h.TasksByID)

	return app
}
