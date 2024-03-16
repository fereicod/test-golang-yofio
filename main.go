package main

import (
	"fmt"

	"github.com/fereicod/test-golang-yofio/utils"
)

func init() {
	for _, amount := range utils.Porportional {
		utils.Minimum += amount
		utils.CountAmount[int32(amount)] = 0
	}
	fmt.Println("Minimo invertir: $", utils.Minimum)
	utils.PrintDivider()
}

func main() {

	for {
		fmt.Print("Ingrese su inversion: ")
		fmt.Scan(&utils.Investment)

		_, error := utils.ValidateInvestment(utils.Investment)

		if error != nil {
			fmt.Println(error.Error())
		} else {
			break
		}

	}
	utils.PrintDivider()
	assigner := utils.Assigner{}
	amount_300, amount_500, amount_700, err := assigner.CreditAssigner(int32(utils.Investment))
	fmt.Println(amount_300, amount_500, amount_700, err)

}
