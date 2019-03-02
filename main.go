package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var OpQueue chan Operation

var OpCounter int

var CurrentDHT11State DHT11State

func main() {

	OpQueue = make(chan Operation, 10)

	// Setup a mux router
	router := mux.NewRouter()

	router.HandleFunc("/api/gpio/push-operation", QueueOperation).Methods("PUT")
	router.HandleFunc("/api/gpio/dht11", GetLastDHT11State).Methods("GET")

	// Start the operation worker
	go opWorker()

	// Start a worker, that reads the dht11 in intervall
	ticker := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				ReadDHT11()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}

// == Handlers for router

type RequestOperation struct {
	Operation string `json:"operation,omitempty"`
	Duration  int    `json:"duration,omitempty"`
}

func QueueOperation(w http.ResponseWriter, req *http.Request) {

	requestOperation := RequestOperation{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestOperation); err != nil {
		http.Error(w, "could not decode request body", 400)
		return
	}

	if requestOperation.Duration == 0 || requestOperation.Operation == "" {
		http.Error(w, "field duration and operation required", 400)
		return
	}

	opName := OperationNameFromString(requestOperation.Operation)
	if opName == nil {
		http.Error(w, "unknown operation", 400)
		return
	}

	op := Operation{
		Duration:      requestOperation.Duration,
		OperationName: *opName,
	}

	// Enqueue the op
	OpQueue <- op
}

func GetLastDHT11State(w http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(w).Encode(CurrentDHT11State); err != nil {
		log.Println("Failed to encode current dht11state", err)
	}
}

// OP Worker

func opWorker() {

	for op := range OpQueue {
		processOperation(op)
	}
}
