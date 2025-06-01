package v1

import (
	"github.com/labstack/echo/v4"
)

func CreateRouting(e *echo.Echo) {
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
}
