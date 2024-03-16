package main

import (
	"fmt"
	"net/http"

	"github.com/fereicod/test-golang-yofio/rest"
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

	http.HandleFunc("/credit-assignment", func(w http.ResponseWriter, r *http.Request) {
		rest.HandleInvestment(w, r)
	})

	println("Server listening on port" + utils.SERVER_PORT + " ...")
	http.ListenAndServe(utils.SERVER_PORT, nil)

}
