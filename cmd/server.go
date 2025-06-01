package main

import (
	"github.com/labstack/echo/v4"
	"web-server/internal/api/v1"
)

func main() {
	e := echo.New()
	v1.CreateRouting(e)

	e.Logger.Fatal(e.Start(":1323"))
}
