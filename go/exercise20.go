package thefarm

import (
	"errors"
	"fmt"
)

var (
	ErrScaleNegative  = errors.New("fodder amount cannot be negative")
	ErrDivisionByZero = errors.New("division by zero")
)

func DivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error) {
	if numCows == 0 {
		return 0, ErrDivisionByZero
	}

	fodderAmount, err := fodderCalculator.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	if fodderAmount < 0 {
		return 0, ErrScaleNegative
	}

	amountPerCow := (fodderAmount / float64(numCows)) * factor

	return amountPerCow, nil
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			cows:    numCows,
			message: "there are no negative cows",
		}
	}
	if numCows == 0 {
		return &InvalidCowsError{
			cows:    numCows,
			message: "no cows don't need food",
		}
	}
	return nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalculator, numCows)
}
