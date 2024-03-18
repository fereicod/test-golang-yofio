package utils

import "fmt"

func CalculateCreditTypes(investment int) {

	isMultiple, _, n := isMulti(investment, nil)
	if isMultiple {
		addCreditTypes(n, nil)
	} else {
		convertToMultiple(investment)
	}
}

func convertToMultiple(investment int) {
	for _, creditType := range CreditType {
		success := subCalculateCreditTypes(investment-creditType, creditType)
		if success {
			break
		}
	}
}

func subCalculateCreditTypes(investment int, fixedACreditType int) bool {
	if investment >= 0 {
		_, rest, n := isMulti(investment, nil)
		for _, creditType := range CreditType {
			isMultiple, _, nCT := isMulti(rest, creditType)
			if isMultiple {
				incrCreditType(fixedACreditType)
				addCreditTypes(n, nil)
				addCreditTypes(nCT, creditType)
				return true
			}
		}
	}
	return false
}

func addCreditTypes(n int, fixedACreditType interface{}) {
	for n > 0 {
		if fixedACreditType != nil {
			incrCreditType(fixedACreditType.(int))
		} else {
			for _, creditType := range CreditType {
				incrCreditType(creditType)
			}
		}
		n--
	}
}

func incrCreditType(creditType int) {
	CountCreditType[int32(creditType)] += 1
}

func CleanCreditType() {
	CountCreditType = map[int32]int32{}
}

func GetStatistics(investments []Investment) (total_of_assignemts int, successful_assignments int, unsuccessful_assignments int, average_successful_assignments float32, average_unsuccessful_assignments float32) {
	total_of_assignemts = len(investments)

	var total_amount_successful, total_amount_unsuccessful int
	for _, investment := range investments {
		if investment.Processed {
			successful_assignments += 1
			total_amount_successful += investment.Investment
		} else {
			unsuccessful_assignments += 1
			total_amount_unsuccessful += investment.Investment
		}
	}

	average_successful_assignments = float32(total_amount_successful) / float32(successful_assignments)
	average_unsuccessful_assignments = float32(total_amount_unsuccessful) / float32(unsuccessful_assignments)
	return
}

func PrintDivider() {
	fmt.Println("-----------------------------------")
}

func PrintResult(investment int) {
	var total_amounts int
	fmt.Println("El monto de inversion fue: $", investment)
	for key, value := range CountCreditType {
		total := (key * value)
		total_amounts += int(total)
		fmt.Println("Monto de $", key, " x ", value, " = $", total)
	}
	PrintDivider()
	fmt.Println("Un total de --------> = $", total_amounts)
}
