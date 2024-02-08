package validator

import (
	"strconv"
	"strings"
)

func IsValidCreditCard(cardNo string) bool {
	cardNo = strings.ReplaceAll(cardNo, " ", "")
	digits := make([]int, len(cardNo))
	for i, num := range cardNo {
		digit, err := strconv.Atoi(string(num))
		if err != nil {
			return false
		}
		digits[i] = digit
	}
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] *= 2
		if digits[i] > 9 {
			digits[i] -= 9
		}
	}
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	return sum%10 == 0
}
