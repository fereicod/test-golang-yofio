package utils

import "strconv"

func ValidateInvestment(investment int) (bool, error) {
	var err GeneralError
	var is_validate bool = true

	is_multiple, _, _ := isMulti(investment, 100)
	if !is_multiple || investment < Minimum {
		is_validate = false
		if investment <= Minimum {
			err.Message = "La cantidad no puede ser menos de $" + strconv.Itoa(int(Minimum))
		} else {
			err.Message = "La cantidad no es multiplo de 100, favor de ingresar de nuevo"
		}
		return is_validate, err
	}

	return is_validate, nil
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
	return false, total_amounts - Investment
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
