package main

import (
	"ecom-tt/internal/app"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	slog.SetDefault(logger)
	if err := app.Run(); err != nil {
		slog.Error("app run", slog.Any("error", err))
		os.Exit(1)
	}
}
