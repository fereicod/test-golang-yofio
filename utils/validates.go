package utils

import "strconv"

func ValidateInvestment(investment interface{}) (bool, error) {
	var err GeneralError
	var isValidate bool = true

	amount, valueAmount, error := isNumeric(investment)
	if error != nil {
		return valueAmount, error
	}

	is_multiple, _, _ := isMulti(amount, 100)
	if !is_multiple || amount < Minimum {
		isValidate = false
		if amount <= Minimum {
			err.Message = "La cantidad no puede ser menos de $" + strconv.Itoa(int(Minimum))
		} else {
			err.Message = "La cantidad no es multiplo de 100, favor de ingresar de nuevo"
		}
		return isValidate, err
	}

	return isValidate, nil
}

func isNumeric(amount interface{}) (int, bool, error) {
	var err GeneralError
	value, boolValue := amount.(int)
	if !boolValue {
		err.Message = "No es un dato numerico"
		return int(value), boolValue, err
	}
	return int(value), boolValue, nil
}

func ValidateResult() (bool, int) {
	var total_amounts int
	for key, value := range CountAmount {
		total := (key * value)
		total_amounts += int(total)
	}
	if total_amounts == Investment {
		return true, 0
	}
	return false, total_amounts - Investment.(int)
}

func isMulti(amount int, multiple interface{}) (bool, int, int) {
	if multiple == nil {
		multiple = Minimum
	}
	mod := amount % multiple.(int)
	div := amount / multiple.(int)
	return (mod == 0), mod, div
}

func existsInProportional(searchAmount int) bool {
	for _, amount := range Porportional {
		if amount == searchAmount {
			return true
		}
	}
	return false
}
