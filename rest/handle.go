package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fereicod/test-golang-yofio/utils"
)

func HandleInvestment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		processInvestment(w, r)
		return
	}

	http.Error(w, "Method not allowed to /credit-assignment", http.StatusMethodNotAllowed)
}

func processInvestment(w http.ResponseWriter, r *http.Request) bool {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body of request.", http.StatusBadRequest)
		return false
	}
	defer r.Body.Close()

	investment := BodyToInvestment(body)
	if investment == nil {
		http.Error(w, "Error wrapping the body to investment.", http.StatusBadRequest)
		return false
	}

	utils.Investment = investment.Amount
	_, error := utils.ValidateInvestment(utils.Investment)

	if error != nil {
		fmt.Println(error.Error())
		http.Error(w, error.Error(), http.StatusBadRequest)
		return false
	}

	utils.PrintDivider()
	investment_req := utils.InvestmentReq{}
	amount_300, amount_500, amount_700, err := investment_req.CreditAssigner(int32(utils.Investment.(int)))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "No es posible hacer la operacion", http.StatusBadRequest)
		return false
	}

	resData := &utils.InvestmentJSON{
		Amount_300: int(amount_300),
		Amount_500: int(amount_500),
		Amount_700: int(amount_700),
	}

	resJSON := GetResponseDataJSON(*resData)

	if resJSON == nil {
		http.Error(w, "Error converting the response data to JSON. ", http.StatusInternalServerError)
		return false
	}
	utils.CleanProcess()

	w.WriteHeader(http.StatusOK)
	w.Write(*resJSON)

	return true
}
