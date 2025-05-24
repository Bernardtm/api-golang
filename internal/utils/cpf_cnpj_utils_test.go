package utils_test

import (
	"bernardtm/backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCPF(t *testing.T) {
	cpf := "417.206.492-27"

	assert.True(t, utils.ValidateCPF(cpf))
}

func TestValidCNPJ(t *testing.T) {
	cnpj := "70.866.482/0001-42"

	assert.True(t, utils.ValidateCNPJ(cnpj))
}

func TestNotValidCPF(t *testing.T) {
	cpf := "417.206.492-28"

	assert.False(t, utils.ValidateCPF(cpf))
}

func TestNotValidCNPJ(t *testing.T) {
	cnpj := "12.345.678/0001-98"

	assert.False(t, utils.ValidateCNPJ(cnpj))
}
