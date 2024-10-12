package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfig_ValidEnvVars tests loading configuration with valid environment variables
func TestLoadConfig_ValidEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("PORT", "8080")
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_MIN_POOL", "5")
	os.Setenv("MONGO_MAX_POOL", "15")
	os.Setenv("JWT_SECRET", "secret")
	os.Setenv("ENDERECO_API", "http://api.example.com")

	defer func() {
		// Clean up environment variables after test
		os.Unsetenv("PORT")
		os.Unsetenv("MONGO_URI")
		os.Unsetenv("MONGO_MIN_POOL")
		os.Unsetenv("MONGO_MAX_POOL")
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("ENDERECO_API")
	}()

	config, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "mongodb://localhost:27017", config.MongoURI)
	assert.Equal(t, uint64(5), config.MongoMinPool)
	assert.Equal(t, uint64(15), config.MongoMaxPool)
	assert.Equal(t, "secret", config.JWTSecret)
	assert.Equal(t, "http://api.example.com", config.EnderecoAPI)
}

// TestLoadConfig_NoEnvFile tests loading configuration without a .env file
func TestLoadConfig_NoEnvFile(t *testing.T) {
	// Clean up environment variables before the test
	os.Unsetenv("PORT")
	os.Unsetenv("MONGO_URI")
	os.Unsetenv("MONGO_MIN_POOL")
	os.Unsetenv("MONGO_MAX_POOL")
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("ENDERECO_API")

	// Load configuration without a .env file
	config, err := LoadConfig()
	assert.NoError(t, err)

	// Assert default values
	assert.Equal(t, "", config.Port)
	assert.Equal(t, "", config.MongoURI)
	assert.Equal(t, uint64(1), config.MongoMinPool)  // Default fallback
	assert.Equal(t, uint64(10), config.MongoMaxPool) // Default fallback
	assert.Equal(t, "", config.JWTSecret)
	assert.Equal(t, "", config.EnderecoAPI)
}

// TestParseUint tests the parseUint helper function
func TestParseUint(t *testing.T) {
	tests := []struct {
		input     string
		fallback  uint64
		expected  uint64
		expectErr bool
	}{
		{"", 1, 1, false},       // Fallback case
		{"5", 1, 5, false},      // Valid number
		{"invalid", 1, 0, true}, // Invalid number
		{"10", 5, 10, false},    // Another valid number
	}

	for _, test := range tests {
		result, err := parseUint(test.input, test.fallback)
		if test.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

// TestLoadConfig_InvalidMinPool tests loading configuration with invalid MIN_POOL value
func TestLoadConfig_InvalidMinPool(t *testing.T) {
	os.Setenv("MONGO_MIN_POOL", "invalid")

	defer os.Unsetenv("MONGO_MIN_POOL")

	_, err := LoadConfig()
	assert.Error(t, err)
}

// TestLoadConfig_InvalidMaxPool tests loading configuration with invalid MAX_POOL value
func TestLoadConfig_InvalidMaxPool(t *testing.T) {
	os.Setenv("MONGO_MAX_POOL", "invalid")

	defer os.Unsetenv("MONGO_MAX_POOL")

	_, err := LoadConfig()
	assert.Error(t, err)
}
