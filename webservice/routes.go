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
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")

	s := r.PathPrefix("/device").Subrouter()
	s.HandleFunc("/", deviceHandler).Methods("GET")
	s.HandleFunc("/{port:[0-9]+}", deviceHandler).Methods("GET")
	s.HandleFunc("/{port:[0-9]+}/power", powerHandler).Methods("GET", "POST")
	s.HandleFunc("/{port:[0-9]+}/volume", volumeHandler).Methods("POST")

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	SendResponse(w, "Hello")
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	port := vars["port"]

	if port != "" {
		port, _ := strconv.Atoi(port)
		SendOjectResponse(w, hdmiControl.GetDeviceInfo(port))
	} else {
		SendOjectResponse(w, hdmiControl.GetActiveDeviceList())
	}
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
		err := hdmiControl.Power(getRequestBody(w, r).State)
		if err != nil {
			SendError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func volumeHandler(w http.ResponseWriter, r *http.Request) {
	err := hdmiControl.SetVolume(getRequestBody(w, r).State)

	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
	}
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
