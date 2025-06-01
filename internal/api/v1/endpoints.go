package v1

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"slices"
	"strings"
)

func getDbData() *[]MaxAmountData {
	dbData := &[]MaxAmountData{
		{CurrencyCode: "USD", Amount: 111.11},
		{CurrencyCode: "RUB", Amount: 222.22},
		{CurrencyCode: "EUR", Amount: 333.33},
		{CurrencyCode: "YEN", Amount: 444.44},
	}

	return dbData
}

func getMaxAmount(c echo.Context) error {
	currencyCodes := c.QueryParam("currency_codes")
	baseMaxAmount := c.QueryParam("base_max_amount")
	baseMaxAmountCurrencyCode := c.QueryParam("base_max_amount_currency_code")

	logger := GetLLogger()
	logger.Info(
		"Running Selections app:",
		slog.String("currency_codes", string(currencyCodes)),
		slog.String("base_max_amount", baseMaxAmount),
		slog.String("base_max_amount_currency_code", baseMaxAmountCurrencyCode),
	)

	dbData := getDbData()
	var maxAmountData []MaxAmountData
	parsedCurrencyCodes := strings.Split(currencyCodes, ",")
	for _, data := range *dbData {
		if slices.Contains(parsedCurrencyCodes, data.CurrencyCode) {
			maxAmountData = append(maxAmountData, data)
		}

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
