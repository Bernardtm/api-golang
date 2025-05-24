package utils

import (
	"fmt"
	"time"
)

const (
	DateLayout        = "2006-01-02"
	DateTimeLayout    = "2006-01-02 15:04:05"
	DateInvitesLayout = "2006-01-02 15:04"
)

func ParseDateISO8601(dateString *string) (string, error) {
	if dateString == nil || *dateString == "" {
		return "", nil
	}
	parsedDate, err := time.Parse(time.RFC3339, *dateString)
	if err != nil {
		return "", err
	}

	return parsedDate.Format("02/01/2006"), nil
}

func ParseDateInvites(dateString *string) (string, error) {
	parsedDate, err := time.Parse(time.RFC3339, *dateString)
	if err != nil {
		return "", err
	}

	return parsedDate.Format("02/01/2006 15:04"), nil
}

func ParseDateTime(value string) (string, error) {
	inputFormat := DateTimeLayout
	outputFormat := "02/01/2006 15:04:05"
	parsedTime, err := time.Parse(inputFormat, value)

	if err != nil {
		return "", err
	}

	return parsedTime.Format(outputFormat), nil
}

func ParseDate(dateString *string) (string, error) {
	if dateString == nil || *dateString == "" {
		return "", nil
	}

	// Attempt to parse the date using the proper format
	parsedDate, err := time.Parse(DateLayout, *dateString) // Use custom format "YYYY-MM-DD"
	if err != nil {
		return "", err
	}

	return parsedDate.Format("02/01/2006"), nil
}

func ParseStringToDate(dateString *string) (time.Time, error) {

	// Parse the string into a time.Time object
	parsedTime, err := time.Parse(DateLayout, *dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %e", err)
	}

	return parsedTime, nil
}
