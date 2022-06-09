package utils

import "strconv"

func HasSufficientBalanceChecker(availableBalance, inputAmount string) bool {
	totalAvailableBalance, _ := strconv.ParseFloat(availableBalance, 64)
	totalInputAmount, _ := strconv.ParseFloat(inputAmount,64)
	if totalAvailableBalance >= totalInputAmount {
		return true
	}
	return false
}
