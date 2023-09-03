package business

import (
	"reflect"
	"testing"
)

func Test_getCardType(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   string
	}{
		{"378282246310005", AMEX},
		{"6011000990139424", DISCOVER},
		{"9937373", ""},
		{"4222222222222", VISA},
		{"371449635398431", AMEX},
		{"30569309025904", DINERS},
		{"4012888888881881", VISA},
	}

	for _, test := range tests {
		t.Run(test.cardNumber, func(t *testing.T) {
			actual, _ := getCardType(test.cardNumber)
			if actual != test.expected {
				t.Errorf("Card type don't match: expected %v but got %v\n", test.expected, actual)
			}
		})
	}
}
func Test_isCardNumberValid(t *testing.T) {
	//{card number, is valid or not}
	expectedResult := map[string]bool{
		"4624748233249780": true,
		"378282246310005":  true,
		"6011000990139424": true,
		"6011257926988965": true,
	}
	actualResult := make(map[string]bool)
	for card := range expectedResult {
		actualResult[card] = isCardNumberValid(getIntArrayFromString(card))
	}
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Failed!\nActual result is %v\n", actualResult)
	}
}
