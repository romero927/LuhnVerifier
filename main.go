package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//https://zerotomastery.io/blog/golang-practice-projects/
//Example Number = 5333-6195-0371-5702

//SUMMARY: Credit cards will often use the Luhn algorithm to confirm the validity of a credit card number.
//First, implement the algorithm as a microservice and then expose the functionality with a JSON API.
//This project is a web-enabled micro service.
//It accepts a credit card number in an HTTP request before returning a response.
//The response indicates whether the number is valid according to the Luhn algorithm.

//Implementing this project requires a series of steps that looks something like this:
//1. Implement the Luhn algorithm
//2. Create an HTTP server
//3. Configure the server to respond to GET requests having a JSON payload
//4. Accept valid JSON requests and proceed to step 5, whilst rejecting invalid requests using an HTTP 400 status code
//5. Extract the credit card number from the JSON payload
//6. Run the Luhn algorithm on the number
//7. Wrap the result into a JSON response payload
//8. Return the payload back to the user through the HTTP server
//9. You can grab both of the packages for this project below:

// Struct to restore CC Details received on Verify Endpoint
type CreditCard struct {
	CCN string
}

// Struct to store response sent after running Luhn Algo
type Response struct {
	Valid bool `json:"valid"`
}

// Main Entry point of app
func main() {
	HTTPServer()
}

// Start-up HTTP server. Currently only implementing Verify endpoint.
func HTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/Verify", validateCreditCardNumber)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Process CCN that was received on Verify endpoint
func validateCreditCardNumber(w http.ResponseWriter, r *http.Request) {
	var c CreditCard
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		//JSON received wasn't valid, return 400
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	//Run Luhn algorithm against received CCN and build response object
	response := Response{Valid: luhnAlgorithm(c.CCN)}

	//Marshal the response struct into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		//Error during JSON marshal, throw 500
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	//Set header to JSON type
	w.Header().Set("Content-Type", "application/json")
	//Return IsValid JSON response
	w.Write(jsonResponse)
}

// Luhn implementation taken from https://github.com/durango/go-credit-card/blob/master/creditcard.go
// More details here: https://en.wikipedia.org/wiki/Luhn_algorithm
/*
function isValid(cardNumber[1..length])
    sum := 0
    parity := length mod 2
    for i from 1 to length do
        if i mod 2 != parity then
            sum := sum + cardNumber[i]
        elseif cardNumber[i] > 4 then
            sum := sum + 2 * cardNumber[i] - 9
        else
            sum := sum + 2 * cardNumber[i]
        end if
    end for
    return cardNumber[length] == (10 - (sum mod 10))
end function
*/
func luhnAlgorithm(cardNumber string) bool {
	//Strip - if included in CCN
	cardNumber = strings.Replace(cardNumber, "-", "", -1)

	var sum int
	var alternate bool

	numberLen := len(cardNumber)

	if numberLen < 13 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(cardNumber[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0
}
