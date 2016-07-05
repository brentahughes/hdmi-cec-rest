package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)

	response := httpResponse{
		Status:  "success",
		Message: message,
	}

	outputResponse(w, response)
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	response := httpResponse{
		Status:  "error",
		Message: message,
	}

	outputResponse(w, response)
}

func SendOjectResponse(w http.ResponseWriter, response interface{}) {
	outputResponse(w, response)
}

func outputResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	outgoingJSON, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(outgoingJSON))
}
