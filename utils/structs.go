package utils

import (
	"strconv"
)

type InvestmentReq struct {
	Amount int `json:"investment"`
}

func (i InvestmentReq) CreditAssigner(investment int32) (int32, int32, int32, error) {
	var err GeneralError

	CalculateCreditTypes(int(investment))

	success, diffAmount := ValidateResult(int(investment))
	if !success {
		err.Message = "Error " + strconv.Itoa(int(diffAmount)) + " to be distributed."
		return 0, 0, 0, err
	}
	return CountCreditType[300], CountCreditType[500], CountCreditType[700], nil

}

type InvestmentJSON struct {
	Credit_type_300 int `json:"credit_type_300"`
	Credit_type_500 int `json:"credit_type_500"`
	Credit_type_700 int `json:"credit_type_700"`
}

type Investment struct {
	Investment      int
	Credit_type_300 int
	Credit_type_500 int
	Credit_type_700 int
	Processed       bool
}

type Statistics struct {
	Total_of_assignments             int     `json:"total_of_assignemts"`
	Successful_assignments           int     `json:"successful_assignments"`
	Unsuccessful_assignments         int     `json:"unsuccessful_assignments"`
	Average_successful_assignments   float32 `json:"average_successful_assignments"`
	Average_unsuccessful_assignments float32 `json:"average_unsuccessful_assignments"`
}
