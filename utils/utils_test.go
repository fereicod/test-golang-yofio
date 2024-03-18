package utils_test

import (
	"testing"

	"github.com/fereicod/test-golang-yofio/utils"
	"github.com/stretchr/testify/suite"
)

type CalculateCreditTypesTestSuite struct {
	suite.Suite
}

func TestCalulateCreditTypeTestSuite(t *testing.T) {
	suite.Run(t, new(CalculateCreditTypesTestSuite))
}

func (s *CalculateCreditTypesTestSuite) SetupTest() {
	utils.CleanCreditType()
}

func (s *CalculateCreditTypesTestSuite) TestCalculateCreditTypes_multiple() {
	var investment = 1500
	var default_count = int32(1)

	utils.CalculateCreditTypes(investment)
	isValidate, diff := utils.ValidateResult(investment)

	s.True(isValidate)
	s.Equal(diff, 0)
	for _, creditType := range utils.CountCreditType {
		s.Equal(creditType, default_count)
	}
}

func (s *CalculateCreditTypesTestSuite) TestCalculateCreditTypes_not_multiple() {
	var investment = 7600
	var countCreditType = map[int32]int32{}
	countCreditType[300] = int32(7)
	countCreditType[500] = int32(4)
	countCreditType[700] = int32(5)

	utils.CalculateCreditTypes(investment)
	isValidate, diff := utils.ValidateResult(investment)

	s.True(isValidate)
	s.Equal(diff, 0)
	for creditType, count := range utils.CountCreditType {
		var ct = countCreditType[creditType]
		s.Equal(count, ct)
	}
}

func (s *CalculateCreditTypesTestSuite) TestGetStatistics() {
	var inputs = [...]int{1500, 7600}
	var investments []utils.Investment
	investments = append(investments, utils.Investment{
		Investment:      inputs[0],
		Credit_type_300: 1,
		Credit_type_500: 1,
		Credit_type_700: 1,
		Processed:       true,
	})
	investments = append(investments, utils.Investment{
		Investment:      inputs[1],
		Credit_type_300: 7,
		Credit_type_500: 4,
		Credit_type_700: 5,
		Processed:       true,
	})
	investments = append(investments, utils.Investment{
		Investment:      30,
		Credit_type_300: 0,
		Credit_type_500: 0,
		Credit_type_700: 0,
		Processed:       false,
	})

	for investment := range inputs {
		utils.CalculateCreditTypes(investment)
	}

	total_of_assignemts, successful_assignments, unsuccessful_assignments, average_successful_assignments, average_unsuccessful_assignments := utils.GetStatistics(investments)
	s.Equal(total_of_assignemts, 3)
	s.Equal(successful_assignments, 2)
	s.Equal(unsuccessful_assignments, 1)
	s.Equal(average_successful_assignments, float32(4550))
	s.Equal(average_unsuccessful_assignments, float32(30))
}
