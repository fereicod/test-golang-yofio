package rest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fereicod/test-golang-yofio/db"
	"github.com/fereicod/test-golang-yofio/utils"
)

func ProcessInvestment(w http.ResponseWriter, r *http.Request, database *sql.DB) bool {

	body, errRead := io.ReadAll(r.Body)
	if errRead != nil {
		http.Error(w, "Error reading the body of request.", http.StatusBadRequest)
		return false
	}
	defer r.Body.Close()

	investment := BodyToInvestment(body)
	if investment == nil {
		http.Error(w, "Error wrapping the body to investment.", http.StatusBadRequest)
		return false
	}

	assigner := db.AssignerDB{
		Investment:      0,
		Credit_type_300: 0,
		Credit_type_500: 0,
		Credit_type_700: 0,
		Processed:       false,
	}

	errorAssigner := utils.GeneralError{}
	if investment.Amount > 0 {
		utils.PrintDivider()
		investment_req := utils.InvestmentReq{}
		credit_type_300, credit_type_500, credit_type_700, errorAssigner := investment_req.CreditAssigner(int32(investment.Amount))

		assigner.Investment = int32(investment.Amount)
		assigner.Credit_type_300 = credit_type_300
		assigner.Credit_type_500 = credit_type_500
		assigner.Credit_type_700 = credit_type_700

		if errorAssigner == nil {
			assigner.Processed = true
		}

		errorDB := db.InsertInvestment(assigner, database)
		if errorDB != nil {
			utils.CleanCreditType()
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error creating the assigner in the database. " + errorDB.Error()))
			return false
		}

	} else {
		errorAssigner.Message = "Error in distributing 0 investment."
	}

	if !assigner.Processed {
		utils.CleanCreditType()
		fmt.Println(errorAssigner.Error())
		http.Error(w, "Error in distributing the money. "+errorAssigner.Error(), http.StatusBadRequest)
		return false
	}

	resData := &utils.InvestmentJSON{
		Credit_type_300: int(assigner.Credit_type_300),
		Credit_type_500: int(assigner.Credit_type_500),
		Credit_type_700: int(assigner.Credit_type_700),
	}

	resJSON := GetResponseDataJSON(*resData)
	if resJSON == nil {
		utils.CleanCreditType()
		http.Error(w, "Error converting the response data to JSON. ", http.StatusInternalServerError)
		return false
	}

	utils.CleanCreditType()

	w.WriteHeader(http.StatusOK)
	w.Write(*resJSON)

	return true
}

func GetStatistics(w http.ResponseWriter, database *sql.DB) {
	rows, errorDB := db.GetInvestment(database)
	if errorDB != nil {
		panic(errorDB.Error())
	}
	defer rows.Close()

	var investments []utils.Investment
	for rows.Next() {
		var investment utils.Investment
		err := rows.Scan(&investment.Investment, &investment.Credit_type_300, &investment.Credit_type_500, &investment.Credit_type_700, &investment.Processed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		investments = append(investments, investment)
	}

	total_of_assignemts, successful_assignments, unsuccessful_assignments, average_successful_assignments, average_unsuccessful_assignments := utils.GetStatistics(investments)
	statistics := utils.Statistics{
		Total_of_assignments:             total_of_assignemts,
		Successful_assignments:           successful_assignments,
		Unsuccessful_assignments:         unsuccessful_assignments,
		Average_successful_assignments:   average_successful_assignments,
		Average_unsuccessful_assignments: average_unsuccessful_assignments,
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statistics)

	//return users
}
