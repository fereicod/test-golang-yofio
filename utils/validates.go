package utils

func isMulti(a int, b interface{}) (bool, int, int) {
	if b == nil {
		var min int
		for _, creditType := range CreditType {
			min += creditType
		}
		b = min
	}
	mod := a % b.(int)
	div := a / b.(int)
	return (mod == 0), mod, div
}

func ValidateResult(investment int) (bool, int) {
	var total_amounts int
	for key, value := range CountCreditType {
		total := (key * value)
		total_amounts += int(total)
	}
	if total_amounts == investment {
		return true, 0
	}
	return false, -(total_amounts - investment)
}
