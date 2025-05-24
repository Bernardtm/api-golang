package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

// GenerateOTP generates an OTP
func GenerateOTP(length int, isAlphanumeric bool) (string, error) {
	if isAlphanumeric {
		return generateAlphanumericOTP(length)
	}
	return generateNumericOTP(length)
}

// GenerateNumericOTP generates an OTP consisting of only numeric digits
func generateNumericOTP(length int) (string, error) {
	otp := ""
	for i := 0; i < length; i++ {
		// Generate a random number between 0 and 9
		num, err := rand.Int(rand.Reader, big.NewInt(10)) // numbers from 0 to 9
		if err != nil {
			return "", fmt.Errorf("error generating random number: %w", err)
		}
		otp += strconv.Itoa(int(num.Int64()))
	}
	return otp, nil
}

// GenerateAlphanumericOTP generates an OTP with numbers, letters, and special symbols
func generateAlphanumericOTP(length int) (string, error) {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%&*"
	otp := ""
	for i := 0; i < length; i++ {
		// Generate a random index within the charset
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error generating random character: %w", err)
		}
		otp += string(charset[index.Int64()])
	}
	return otp, nil
}
