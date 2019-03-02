package main

import (
	"encoding/json"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
	"time"
)

func main() {

	// Setup a mux router
	// router := mux.NewRouter()

	// router.HandleFunc("/api/say-hello", SayHello).Methods("GET")

	// log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {

	msg := hello{"Hi sensorgo"}

	_ = json.NewEncoder(w).Encode(msg)
}

type hello struct {
	Message string `json:"message,omitempty"`
}
