package main

import (
	"fmt"
    "github.com/chbmuc/cec"
	"net/http"
)

func poweron(w http.ResponseWriter, r *http.Request) {
    c, err := cec.Open("", "cec.go")
    if err != nil {
        fmt.Println(err)
    }

    c.PowerOn(0)
}

func poweroff(w http.ResponseWriter, r *http.Request) {
    c, err := cec.Open("", "cec.go")
    if err != nil {
        fmt.Println(err)
    }

    c.Standby(0)
}

func main() {
	http.HandleFunc("/power/on", poweron)
    http.HandleFunc("/power/off", poweroff)
	http.ListenAndServe(":5000", nil)
}
