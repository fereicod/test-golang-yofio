package utils_test

import (
	"testing"

	"github.com/fereicod/test-golang-yofio/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestIsMulti_default(t *testing.T) {
	var a int = 1500

	isMulti, mod, div := utils.IsMulti(a, nil)
	assert.Equal(t, true, isMulti)
	assert.Equal(t, mod, 0)
	assert.Equal(t, div, (a / 1500))
}

func TestIsMulti_with_two_args(t *testing.T) {
	var a int = 1000
	var b interface{} = 5

	isMulti, mod, div := utils.IsMulti(a, b.(int))
	assert.Equal(t, true, isMulti)
	assert.Equal(t, mod, 0)
	assert.Equal(t, div, (a / b.(int)))
}

type ValidateResultTestSuite struct {
	suite.Suite
}

func TestValidateResultTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateResultTestSuite))
}

func (s *ValidateResultTestSuite) SetupTest() {
	utils.CountCreditType[300] = 1
	utils.CountCreditType[500] = 1
	utils.CountCreditType[700] = 1
}

func (s *ValidateResultTestSuite) TestValidateResult_success() {
	var a int = 1500

	isValidate, diff := utils.ValidateResult(a)
	s.True(isValidate)
	s.Equal(diff, 0)
}

func (s *ValidateResultTestSuite) TestValidateResult_fail() {
	var a int = 1000

	isValidate, diff := utils.ValidateResult(a)
	s.False(isValidate)
	s.NotEqual(diff, 0)
}
