package rest

import (
	"encoding/json"

	"github.com/fereicod/test-golang-yofio/utils"
)

func BodyToInvestment(body []byte) *utils.InvestmentReq {
	if len(body) == 0 {
		return nil
	}

	var assigner utils.InvestmentReq
	err := json.Unmarshal(body, &assigner)
	if err != nil {
		return nil
	}

	return &assigner
}
