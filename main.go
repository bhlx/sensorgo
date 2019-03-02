package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var OpQueue chan Operation

var OpCounter int

func main() {

	OpQueue = make(chan Operation, 10)

	// Setup a mux router
	router := mux.NewRouter()

	router.HandleFunc("/api/gpio/push-operation", QueueOperation).Methods("PUT")

	// Start the worker
	go opWorker()

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

// OP Worker

func opWorker() {

	for op := range OpQueue {
		processOperation(op)
	}
}
