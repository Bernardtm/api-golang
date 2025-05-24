package utils_test

import (
	"bernardtm/backend/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBool(t *testing.T) {
	generatedBool := utils.GenerateBool()
	assert.True(t, generatedBool == true || generatedBool == false)
}

func TestGenerateInt(t *testing.T) {
	generatedInt := utils.GenerateInt(5, 9)
	assert.True(t, generatedInt >= 5 && generatedInt <= 9)
}

func TestGenerateFloat64(t *testing.T) {
	assert.True(t, utils.GenerateFloat64(1.0, 2.0) >= 1.0)
	assert.True(t, utils.GenerateFloat64(1599.0, 1599.99) >= 1598.99)
}

func TestGenerateDate(t *testing.T) {
	minDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	generatedDate := utils.GenerateDate(minDate, maxDate)
	assert.True(t, generatedDate.After(minDate))
	assert.True(t, generatedDate.Before(maxDate))
}

func TestGenerateRandomNumbersString(t *testing.T) {
	length := 10
	generatedNumberString := utils.GenerateRandomNumbersString(length)
	assert.Len(t, generatedNumberString, length)
}

func TestFundNameGenerator(t *testing.T) {
	fundName := utils.FundNameGenerator()
	assert.NotEmpty(t, fundName)
}

func TestSettlementNameGenerator(t *testing.T) {
	settlementName := utils.SettlementNameGenerator()
	assert.NotEmpty(t, settlementName)
}
