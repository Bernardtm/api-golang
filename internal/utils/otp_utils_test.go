package utils

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNumericOTP(t *testing.T) {

	// generate 8 digits alphanumeric otp
	size := 8
	isAlphanumeric := false
	otp, err := GenerateOTP(size, isAlphanumeric)
	if err != nil {
		t.Errorf("Error generating OTP: %v", err)
	}
	fmt.Println("OTP generated:", otp)
	assert.Equal(t, len(otp), size, "OTP length should be "+strconv.Itoa(size))
}

func TestGenerateAlphanumericOTP(t *testing.T) {

	// generate 8 digits alphanumeric otp
	size := 8
	isAlphanumeric := true
	otp, err := GenerateOTP(size, isAlphanumeric)
	if err != nil {
		t.Errorf("Error generating OTP: %v", err)
	}
	fmt.Println("OTP generated:", otp)
	assert.Equal(t, len(otp), size, "OTP length should be "+strconv.Itoa(size))
}
