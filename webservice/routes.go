package webservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bah2830/hdmi-cec-rest/hdmiControl"
	"github.com/gorilla/mux"
)

type Request struct {
    State string `json:"state"`
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/device", deviceHandler).Methods("GET")
	router.HandleFunc("/device/{port}/power", powerHandler).Methods("GET", "POST")
	router.HandleFunc("/device/{port}/volume", volumeHandler).Methods("POST")

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	SendResponse(w, "Hello")
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    port, _ := strconv.Atoi(vars["port"])

    hdmiControl.SetPort(port)
    SendOjectResponse(w, hdmiControl.GetActiveDeviceList())
}

func powerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    port, _ := strconv.Atoi(vars["port"])

    hdmiControl.SetPort(port)

	switch r.Method {
		case "GET":
			status := hdmiControl.GetPowerStatus()

			SendResponse(w, status)
		case "POST":
			hdmiControl.Power(getRequestBody(w, r).State)
	}
}

func volumeHandler(w http.ResponseWriter, r *http.Request) {
	hdmiControl.SetVolume(getRequestBody(w, r).State)
}

func getRequestBody(w http.ResponseWriter, r *http.Request) Request {
	var request Request

    decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
	}

	return request
}