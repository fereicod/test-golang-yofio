package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fereicod/test-golang-yofio/utils"
)

func HandleInvestment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading the body of request.", http.StatusBadRequest)
		}
		defer r.Body.Close()

		investment := BodyToInvestment(body)
		if investment == nil {
			http.Error(w, "Error wrapping the body to investment.", http.StatusBadRequest)
		}

		resData := process(w, r, investment.Amount)
		resJSON := GetResponseDataJSON(*resData)

		if resJSON == nil {
			http.Error(w, "Error converting the response data to JSON. ", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(*resJSON)

		return
	}

	http.Error(w, "Method not allowed to /credit-assignment", http.StatusMethodNotAllowed)
}

func process(w http.ResponseWriter, r *http.Request, investment int) *utils.InvestmentJSON {

	utils.Investment = investment
	_, error := utils.ValidateInvestment(utils.Investment)

	if error != nil {
		fmt.Println(error.Error())
		http.Error(w, error.Error(), http.StatusBadRequest)
	}

	utils.PrintDivider()
	investment_req := utils.InvestmentReq{}
	amount_300, amount_500, amount_700, err := investment_req.CreditAssigner(int32(utils.Investment))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return &utils.InvestmentJSON{
		Amount_300: int(amount_300),
		Amount_500: int(amount_500),
		Amount_700: int(amount_700),
	}
}
