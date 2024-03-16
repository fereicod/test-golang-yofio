package utils

import "fmt"

func CalculateAmounts(investment int) {

	is_multiple, rest, n := isMulti(investment, nil)
	if !is_multiple {
		exists := existsInProportional(rest)
		if exists {
			incrProportional(rest)
			allocateAmounts(n, nil)
		} else {
			toMultiple(investment)
		}
	} else {
		allocateAmounts(n, nil)
	}
}

func toMultiple(amount int) {
	for _, value := range Porportional {
		success := subCalculateAmounts(amount-value, value)
		if success {
			break
		}
	}
}

func subCalculateAmounts(amount int, value int) bool {

	_, rest_origin, n_origin := isMulti(amount, nil)
	for _, val := range Porportional {
		is_multiple, _, n := isMulti(rest_origin, val)
		if is_multiple {
			incrProportional(value)
			allocateAmounts(n_origin, nil)
			allocateAmounts(n, val)
			return true
		}
	}
	return false
}

func allocateAmounts(n int, fixedAmount interface{}) {
	for n > 0 {
		if fixedAmount != nil {
			incrProportional(fixedAmount.(int))
		} else {
			for _, amount := range Porportional {
				incrProportional(amount)
			}
		}
		n--
	}
}

func incrProportional(amount int) {
	CountAmount[int32(amount)] += 1
}

func PrintDivider() {
	fmt.Println("-----------------------------------")
}

func PrintResult() {
	var total_amounts int
	fmt.Println("El monto de inversion fue: $", Investment)
	for key, value := range CountAmount {
		total := (key * value)
		total_amounts += int(total)
		fmt.Println("Monto de $", key, " x ", value, " = $", total)
	}
	PrintDivider()
	fmt.Println("Un total de --------> = $", total_amounts)
}
