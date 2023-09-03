package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/saibikalpa/creditcardvalidator/pkg/api"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/validate/", api.CheckCreditCardNumber).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
}
