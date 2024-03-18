package db

const INSERT_ASSIGNER_STATEMENT = "INSERT INTO assigner (investment, credit_type_300, credit_type_500, credit_type_700, processed) VALUES (?, ?, ?, ?, ?)"

const GET_ASSIGNER_STATEMENT = "SELECT a.investment, a.credit_type_300, a.credit_type_500, a.credit_type_700, a.processed FROM assigner AS a"
