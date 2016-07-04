package main

import (
	"net/http"

	"github.com/bah2830/hdmi-cec-rest/webservice"
)

func main() {
	router := webservice.GetRouter()
    http.ListenAndServe(":5000", router)
}
