package webservice

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/chbmuc/cec"
    "github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    router.HandleFunc("/power", powerHandler).Methods("GET", "POST")

    return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    SendResponse(w, "Hello")
}


func powerHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    c, err := cec.Open("", "cec.go")
    if err != nil {
        fmt.Println(err)
    }

    switch r.Method {
    case "GET":
        status := c.GetDevicePowerStatus(0)

        SendResponse(w, status)
    case "POST":
        decoder := json.NewDecoder(r.Body)
        var request PowerRequest

        err := decoder.Decode(&request)
        if err != nil {
            fmt.Println(err)
        }

        if request.State == "on" {
            c.PowerOn(0)

            SendResponse(w, "on")
        } else {
            c.Standby(0)

            SendResponse(w, "off")
        }
    }
}
