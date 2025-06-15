package interest

func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default: // balance >= 5000
		return 2.475
	}
}

func Interest(balance float64) float64 {
	rate := float64(InterestRate(balance))
	return balance * rate / 100
}

func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	currentBalance := balance

	for currentBalance < targetBalance {
		currentBalance = AnnualBalanceUpdate(currentBalance)
		years++
	}

	return years
}