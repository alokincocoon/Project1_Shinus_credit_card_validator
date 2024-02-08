package api

import (
	"creditcardvalidator/validator"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	CardNumber string
}

type ResponseBody struct {
	IsValid bool
}

func validateCreditCard(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	isValidCreditCard := validator.IsValidCreditCard(requestBody.CardNumber)
	response := ResponseBody{IsValid: isValidCreditCard}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func CheckValid() {

	const endpoint string = "/v1/api/validate/creditcards"
	http.HandleFunc(endpoint, validateCreditCard)

	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	fmt.Printf("go to: http://localhost:8080" + endpoint)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println(err)
	}
}
