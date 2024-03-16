package utils

import "strconv"

type InvestmentReq struct {
	Amount int `json:"investment"`
}

func (i InvestmentReq) CreditAssigner(investment int32) (int32, int32, int32, error) {
	var err GeneralError

	CalculateAmounts(int(investment))

	success, diffAmount := ValidateResult()
	if success {
		PrintResult()
	} else {
		err.Message = "Hubo una diferencia de $" + strconv.Itoa(int(diffAmount))
		return 0, 0, 0, err
	}
	return CountAmount[300], CountAmount[500], CountAmount[700], nil

}

type InvestmentJSON struct {
	Amount_300 int `json:"credit_type_300"`
	Amount_500 int `json:"credit_type_500"`
	Amount_700 int `json:"credit_type_700"`
}
