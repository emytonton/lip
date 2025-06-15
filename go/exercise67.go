package phonenumber

import (
	"errors"
	"unicode"
)


func Number(phoneNumber string) (string, error) {

	digits := ""
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			digits += string(r)
		}
	}

	
	if len(digits) == 11 {
		if digits[0] == '1' {
			digits = digits[1:]
		} else {
			return "", errors.New("11-digit number must start with '1'")
		}
	}

	
	if len(digits) != 10 {
		return "", errors.New("incorrect number of digits")
	}


	if digits[0] < '2' || digits[3] < '2' {
		return "", errors.New("invalid area code or exchange code")
	}

	return digits, nil
}


func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}


func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	area := number[:3]
	exchange := number[3:6]
	subscriber := number[6:]
	return "(" + area + ") " + exchange + "-" + subscriber, nil
}
