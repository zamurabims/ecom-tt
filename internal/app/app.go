package app

import (
	"ecom-tt/internal/handlers"
	"ecom-tt/internal/storage"
	"log/slog"
	"net/http"
)

func Run() error {

	storageMap := storage.NewStorage()
	todoHandler := handlers.NewTodoHandler(storageMap)
	router := getRouter(todoHandler)
	slog.Info("started server on port", slog.String("addr", ":8080"))
	return http.ListenAndServe(":8080", router)

}
