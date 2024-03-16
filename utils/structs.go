package utils

import "strconv"

type Assigner struct {
	Investment int
}

func (a Assigner) CreditAssigner(investment int32) (int32, int32, int32, error) {
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
