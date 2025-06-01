package v1

import (
	"encoding/json"
	"log/slog"
)

func GetMaxAmount(maxAmount *MaxAmount) []byte {
	b, err := json.Marshal(maxAmount)
	logger := GetLLogger()
	if err != nil {
		logger.Info("Error in GetMaxAmount:", slog.String("error", err.Error()))
	}

	return b

}
