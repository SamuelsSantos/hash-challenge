package config

import (
	"os"
	"strings"
	"time"
)

// LayoutISO "YYYY-MM-DD"
const LayoutISO = "2006-01-02"

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

	date, err := time.Parse(LayoutISO, value)
	if err != nil {
		date, _ := time.Parse(LayoutISO, standard)
		return date
	}

	return date
}
