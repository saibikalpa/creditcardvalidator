package api

import (
	"encoding/json"
	"net/http"

	"github.com/saibikalpa/creditcardvalidator/pkg/business"
	"github.com/saibikalpa/creditcardvalidator/pkg/model"
)

func CheckCreditCardNumber(w http.ResponseWriter, r *http.Request) {

	var card model.CreditCard
	err := json.NewDecoder(r.Body).Decode(&card)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	isValid, cardType := business.IsCreditCardValid(card.Number)

	resp := make(map[string]string)
	w.Header().Add("Content-Type", "application/json")
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Card validation failed"
	} else {

		resp["message"] = "Card validation passed."
		resp["cardType"] = cardType
		w.WriteHeader(http.StatusOK)
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
