package api

import (
	validators "mrlparker/high-score/validators"
	"net/http"
)

func GetScore(responseWriter http.ResponseWriter, request *http.Request) {
	valid, err := validators.ValidateGetRequest("", "")

	if !valid && err != nil {
		// Get request is invalid, we should let the user know.
		return
	}

	// Request is valid, we should call the data access layer.
}

func PostScore(responseWriter http.ResponseWriter, request *http.Request) {
	valid, err := validators.ValidatePostRequest("", "", -5)

	if !valid && err != nil {
		// Post request is invalid, we should let the user know.
		return
	}

	// Request is valid, we should call the data access layer.
}
