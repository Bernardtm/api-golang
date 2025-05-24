package utils

import "testing"

func TestValidatePassword_Success(t *testing.T) {
	err := ValidatePassword("12345678")
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
}
