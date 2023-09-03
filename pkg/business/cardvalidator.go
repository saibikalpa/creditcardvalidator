package business

import (
	"errors"
	"regexp"
)

/*
Card constants
*/
const (
	AMEX        = "AMEX"
	MASTERCARD  = "MASTERCARD"
	VISA        = "VISA"
	DINERS      = "DINERS"
	DISCOVER    = "DISCOVER"
	JCB         = "JCB"
	UNSUPPORTED = "UNSUPPORTED"
)

/*
This function returns regex patterns of supported card network
*/
func getSupportedNetworks() map[string]string {
	return map[string]string{
		AMEX:       "^3[4 7]\\d{13}$",
		MASTERCARD: "^5[1 2 3 4 5]\\d{14}$",
		VISA:       "^4(?:\\d{12}|\\d{15})$",
		DINERS:     "^3[6 8]\\d{12}|30[0-5]\\d{11}$",
		DISCOVER:   "^6011\\d{12}$",
		JCB:        "^(?:(?:2123|1800)\\d{11})|3\\d{15}$",
	}
}

/*
This function returns (true, card type) on success
else it returns (false, "")
*/
func IsCreditCardValid(cardNumber string) (bool, string) {

	cardType, err := getCardType(cardNumber)
	if err != nil {
		return false, ""
	} else {
		return isCardNumberValid(getIntArrayFromString(cardNumber)), cardType
	}
}

/*
This function returns a card type on success and returns an error
on failure
*/
func getCardType(cardNumber string) (string, error) {
	networkName, matched := UNSUPPORTED, false
	for k, v := range getSupportedNetworks() {
		expr := regexp.MustCompile(v)
		if expr.MatchString(cardNumber) {
			networkName, matched = k, true
			break
		}
	}
	if matched {
		return networkName, nil
	} else {
		return "", errors.New(UNSUPPORTED)
	}
}

/*
This function validates the credit card number with Luhn's algorithm.
*/
func isCardNumberValid(cardNumber []int) bool {

	var sum int
	length := len(cardNumber)
	for i, j := length-1, length-2; j >= 0; i, j = i-2, j-2 {
		sum += cardNumber[i] + 2*cardNumber[j]%10 + 2*cardNumber[j]/10
	}
	if length%2 != 0 {
		sum += cardNumber[0]
	}
	return sum%10 == 0
}

func getIntArrayFromString(str string) []int {
	res := make([]int, len(str))
	for i, c := range str {
		res[i] = int(c - '0')
	}
	return res
}
