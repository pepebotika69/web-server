package v1

import (
	"github.com/labstack/echo/v4"
)

func CreateRouting(e *echo.Echo) {
	e.GET("/max-amount", getMaxAmount)
}
