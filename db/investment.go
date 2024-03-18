package db

import (
	"database/sql"
)

func InsertInvestment(assigner AssignerDB, database *sql.DB) error {
	_, err := database.Exec(
		INSERT_ASSIGNER_STATEMENT,
		assigner.Investment,
		assigner.Credit_type_300,
		assigner.Credit_type_500,
		assigner.Credit_type_700,
		assigner.Processed,
	)
	return err
}

func GetInvestment(database *sql.DB) (*sql.Rows, error) {
	rows, err := database.Query(GET_ASSIGNER_STATEMENT)
	return rows, err
}
