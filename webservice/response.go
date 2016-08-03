package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpResponse struct {
	State string `json:"state"`
}

type requestEndpoint struct {
	Endpoint      string   `json:"endpoint"`
	Methods       []string `json:"methods"`
	Documentation string   `json:"documentation"`
}

func SendRootResponse(w http.ResponseWriter) {
	response := []requestEndpoint{
		requestEndpoint{
			Endpoint:      "/device",
			Methods:       []string{"GET"},
			Documentation: "Get list of connected device and their details.",
		},
		requestEndpoint{
			Endpoint:      "/device/[LogicalAddress]",
			Methods:       []string{"GET"},
			Documentation: "Get details for single device by its LogicalAddress",
		},
		requestEndpoint{
			Endpoint: "/device/[LogicalAddress]/power",
			Methods:  []string{"GET", "POST"},
			Documentation: `GET: Get power status of device.
				POST: Set power state on or off by sending post data '{\"state\": \"on\"}'`,
		},
		requestEndpoint{
			Endpoint:      "/device/[LogicalAddress]/volume",
			Methods:       []string{"POST"},
			Documentation: "Set volume level up, down, or mute by sending post data '{\"state\": \"up\"}'",
		},
	}

	SendOjectResponse(w, response)
}

func SendResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)

	response := httpResponse{
		State:  message,
	}

	outputResponse(w, response)
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	response := httpResponse{
		State:  fmt.Sprintf("Error: %s", message),
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
