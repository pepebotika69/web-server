package v1

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMaxAmount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/max-amount/")

	expectedResponse := `{"data":null,"error":null,"status":"ok"}` + "\n"
	// Assertions
	if assert.NoError(t, getMaxAmount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResponse, rec.Body.String())
	}

}

func TestGetMaxAmountCustom(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/max-amount-custom/")

	expectedResponse := `{"data":[{"currency_code":"USD","amount":0}],"error":null,"status":"ok"}`
	// Assertions
	if assert.NoError(t, getMaxAmountCustom(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResponse, rec.Body.String())
	}

}

type TestInput struct {
	data []MaxAmountData
	err  error
}
type TestCases struct {
	input    TestInput
	expected string
}

func GetTestCases() []*TestCases {
	maxAmountData := []MaxAmountData{
		{CurrencyCode: "USD", Amount: 123.33},
	}

	tests := []*TestCases{
		{
			input:    TestInput{data: maxAmountData, err: nil},
			expected: `{"data":[{"currency_code":"USD","amount":123.33}],"error":null,"status":"ok"}`,
		},
		{
			input:    TestInput{data: nil, err: nil},
			expected: `{"data":null,"error":null,"status":"ok"}`,
		},
		{
			input:    TestInput{data: nil, err: errors.New("some_error_slug")},
			expected: `{"data":null,"error":"some_error_slug","status":"error"}`,
		},
	}

	return tests

}
func TestGetMaxAmountFunction(t *testing.T) {
	tests := GetTestCases()

	for _, test := range tests {
		testName := fmt.Sprintf("%+v", test.input)
		t.Run(testName, func(t *testing.T) {
			maxAmount := NewMaxAmount(test.input.data, test.input.err)
			response := GetMaxAmount(maxAmount)
			responseForCompare := string(response)

			if responseForCompare != test.expected {
				t.Errorf("got %s\n, want %s", responseForCompare, test.expected)
			}
		})
	}
}
