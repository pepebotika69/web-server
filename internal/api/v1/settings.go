package v1

import (
	"log/slog"
	"os"
)

func GetLLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}
