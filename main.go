package main

import (
	"log"
	"net/http"

	"github.com/fereicod/test-golang-yofio/db"
	"github.com/fereicod/test-golang-yofio/rest"
	"github.com/fereicod/test-golang-yofio/utils"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file.")
	}

	database := db.OpenConnect()
	defer database.Close()

	http.HandleFunc("/credit-assignment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			rest.ProcessInvestment(w, r, database)
			return
		}

		http.Error(w, "Method not allowed to /credit-assignment", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/statistics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rest.GetStatistics(w, database)
			return
		}

		http.Error(w, "Method not allowed to /statistics", http.StatusMethodNotAllowed)
	})

	println("Server listening on port" + utils.SERVER_PORT + " ...")
	http.ListenAndServe(utils.SERVER_PORT, nil)

}
