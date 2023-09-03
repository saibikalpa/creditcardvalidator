package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/saibikalpa/creditcardvalidator/pkg/model"
)

func TestCheckCreditCardNumber(t *testing.T) {

	tests := []struct {
		cardNumber string
		expected   int //status code
	}{
		{"12345", 400},
		{"4222222222222", 200},
		{"5105105105105100", 200},
		{"378734493671000", 200},
		{"99222222", 400},
	}

	for _, test := range tests {
		t.Run(test.cardNumber, func(t *testing.T) {
			card := model.CreditCard{Number: test.cardNumber}
			body, _ := json.Marshal(card)
			t.Logf("body is %v", string(body))
			r := httptest.NewRequest("POST", "/validate", bytes.NewReader(body))
			w := httptest.NewRecorder()
			CheckCreditCardNumber(w, r)
			if w.Result().StatusCode != test.expected {
				t.Errorf("Error occured while checking credit card: expected status code %v but received %v", test.expected, w.Result().StatusCode)
			}
		})
	}
}
