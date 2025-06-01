package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getMaxAmount(c echo.Context) error {
	maxAmountData := []MaxAmountData{
		{CurrencyCode: "USD"},
	}

	maxAmount := NewMaxAmount(maxAmountData, nil)
	// not need to marshal by our self
	return c.JSON(http.StatusOK, maxAmount)
}

func getMaxAmountCustom(c echo.Context) error {
	maxAmountData := []MaxAmountData{
		{CurrencyCode: "USD"},
	}

	maxAmount := NewMaxAmount(maxAmountData, nil)
	return c.String(http.StatusOK, string(GetMaxAmount(maxAmount)))
}
