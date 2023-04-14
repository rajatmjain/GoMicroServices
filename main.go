package main

import (
	"GoMicroServices/details"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Root Handler
	r.HandleFunc("/", rootHandler)

	// Health Handler
	r.HandleFunc("/health", healthHandler)

	// Details Handler
	r.HandleFunc("/details", detailsHandler)

	fmt.Println("Web server up!")
	log.Fatal(http.ListenAndServe(":80", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving root")
	w.WriteHeader(http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Application health")
	response := map[string]string{
		"status":    "UP",
		"timeStamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving details")
	hostName, error := details.GetHostName()
	if error != nil {
		panic(error)
	}
	fmt.Println(hostName)
	IP, error := details.GetIP()
	if error != nil {
		panic(error)
	}
	fmt.Println(IP)

	response := map[string]string{
		"hostName": hostName,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}
