package main

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"web-server/internal/api/v1"
)

func main() {
	port := ":1323"

	logger := v1.GetLLogger()
	logger.Info("Running Selections app:", slog.String("port", port))

	e := echo.New()
	v1.CreateRouting(e)

	err := e.Start(port)
	//todo why on exit do not print it?
	if err != nil {
		logger.Error("Selections app exited with error: ",
			slog.String("error", err.Error()),
		)
	} else {
		logger.Error("Selections app stoped")
	}

}
