package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	cnpjFirstDigitTable  = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjSecondDigitTable = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

const (
	// CNPJFormatPattern is the pattern for string replacement
	// with Regex
	CNPJFormatPattern string = `([\d]{2})([\d]{3})([\d]{3})([\d]{4})([\d]{2})`
)

// CNPJ type definition
type CNPJ string

// NewCNPJ is a helper function to convert and clean a new CNPJ
// from a string
func NewCNPJ(s string) CNPJ {
	return CNPJ(CleanCPFCNPJ(s))
}

// IsValid returns if CNPJ is a valid CNPJ document
func (c *CNPJ) IsValid() bool {
	return ValidateCNPJ(string(*c))
}

// String returns a formatted CNPJ document
// 00.000.000/0001-00
func (c *CNPJ) String() string {

	str := string(*c)

	if !c.IsValid() {
		return str
	}

	expr, err := regexp.Compile(CNPJFormatPattern)

	if err != nil {
		return str
	}

	return expr.ReplaceAllString(str, "$1.$2.$3/$4-$5")
}

// ValidateCNPJ validates a CNPJ document
// You should use without punctuation
func ValidateCNPJ(cnpj string) bool {

	if len(cnpj) != 14 {
		return false
	}

	firstPart := cnpj[:12]
	sum1 := sumDigit(firstPart, cnpjFirstDigitTable)
	rest1 := sum1 % 11
	d1 := 0

	if rest1 >= 2 {
		d1 = 11 - rest1
	}

	secondPart := fmt.Sprintf("%s%d", firstPart, d1)
	sum2 := sumDigit(secondPart, cnpjSecondDigitTable)
	rest2 := sum2 % 11
	d2 := 0

	if rest2 >= 2 {
		d2 = 11 - rest2
	}

	finalPart := fmt.Sprintf("%s%d", secondPart, d2)
	return finalPart == cnpj
}

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

const (
	// CPFFormatPattern is the pattern for string replacement
	// with Regex
	CPFFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
)

// CPF type definition
type CPF string

// NewCPF is a helper function to convert and clean a new CPF
// from a string
func NewCPF(s string) CPF {
	return CPF(CleanCPFCNPJ(s))
}

// IsValid returns if CPF is a valid CPF document
func (c *CPF) IsValid() bool {
	return ValidateCPF(string(*c))
}

// String returns a formatted CPF document
// 000.000.000-00
func (c *CPF) String() string {

	str := string(*c)

	if !c.IsValid() {
		return str
	}

	expr, err := regexp.Compile(CPFFormatPattern)

	if err != nil {
		return str
	}

	return expr.ReplaceAllString(str, "$1.$2.$3-$4")
}

// ValidateCPF validates a CPF document
// You should use without punctuation
func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}

func sumDigit(s string, table []int) int {

	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

// Clean can be used for cleaning formatted documents
func CleanCPFCNPJ(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}
