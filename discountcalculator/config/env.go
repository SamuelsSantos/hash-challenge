package config

import (
	"os"
	"strings"
	"time"
)

// GetenvString get environment value or default value
func GetenvString(key, standard string) string {
	value := os.Getenv(key)
	if len(strings.TrimSpace(value)) == 0 {
		return standard
	}
	return value
}

// GetenvDate get environment value or default value
func GetenvDate(key, standard string) time.Time {
	value := os.Getenv(key)

	date, err := time.Parse("2016-01-01", value)
	if err != nil {
		date, _ := time.Parse("2016-01-01", standard)
		return date
	}

	return date
}
