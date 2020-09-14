package domain

import (
	"time"
)

const layoutISO = "2006-01-02"

type dateRule func(date time.Time) bool

type blackFriday struct {
	Pct     float64
	IsValid dateRule
}

type birthDay struct {
	Pct     float64
	IsValid dateRule
}

func isBlackFriday(data time.Time) bool {
	if data.IsZero() {
		return false
	}

	today := time.Now()
	return data.Format(layoutISO) == today.Format(layoutISO)
}

func isBirthDay(data time.Time) bool {
	if data.IsZero() {
		return false
	}
	today := time.Now()
	return data.Month() == today.Month() && data.Day() == today.Day()
}

// DiscountStrategies is a collection of strategies to calculate discount
type DiscountStrategies struct {
	BlackFriday *blackFriday
	BirthDay    *birthDay
}

// NewDiscountStrategies is instance of the DiscountStrategies
func NewDiscountStrategies() *DiscountStrategies {
	return &DiscountStrategies{
		BlackFriday: &blackFriday{Pct: 0.1, IsValid: isBlackFriday},
		BirthDay:    &birthDay{Pct: 0.05, IsValid: isBirthDay},
	}
}

// StrategyCalculator is an interface that defines a contract to apply a discount in price
type StrategyCalculator interface {
	ApplyDiscount(priceInCents int64) (pct float64, valueInCents int64)
}

// Calculate discount
func Calculate(pct float64, valueInCents int64) int64 {

	value := float64(valueInCents) * pct

	return int64(value)
}

// GetDiscountPercentual -> If it's the user’s birthday, the product has 5% discount.
// If it is black friday (for this test you can assume BlackFriday is November 25th), the product has 10% discount
// No product discount can be bigger than 10%
func GetDiscountPercentual(blackFriday, birthDay time.Time) float64 {

	strategy := NewDiscountStrategies()
	if strategy.BlackFriday.IsValid(blackFriday) {
		return strategy.BlackFriday.Pct
	} else if strategy.BirthDay.IsValid(birthDay) {
		return strategy.BirthDay.Pct
	}

	return 0.0
}
