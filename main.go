package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	// Setup a mux router
	router := mux.NewRouter()

	router.HandleFunc("/api/say-hello", SayHello).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func SayHello(w http.ResponseWriter, r *http.Request) {

	msg := hello{"Hi sensorgo"}

	_ = json.NewEncoder(w).Encode(msg)
}

type hello struct {
	Message string `json:"message,omitempty"`
}
