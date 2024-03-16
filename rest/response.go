package rest

import (
	"encoding/json"

	"github.com/fereicod/test-golang-yofio/utils"
)

func GetResponseDataJSON(res utils.InvestmentJSON) *[]byte {
	resJSON, err := json.Marshal(res)
	if err != nil {
		return nil
	}

	return &resJSON
}
