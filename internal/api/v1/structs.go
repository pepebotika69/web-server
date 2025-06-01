package v1

type MaxAmountData struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float32 `json:"amount"`
}

type MaxAmount struct {
	Data   []MaxAmountData `json:"data"`  //pointer so if empty marshals to null
	Error  *string         `json:"error"` //pointer so if empty marshals to null
	Status string          `json:"status"`
}

func nullString(s string) *string {
	return &s
}

func NewMaxAmount(data []MaxAmountData, err error) *MaxAmount {
	var e *string
	status := StatusOk

	if err != nil {
		e = nullString(err.Error())
		status = StatusError
	}

	return &MaxAmount{
		Data:   data,
		Error:  e,
		Status: status,
	}
}
