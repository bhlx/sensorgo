package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"io/ioutil"
	"log"
)

type DHT11State struct {
	temperature float32
	humidity    float32
}

func ReadDHT11() {

	// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)
	if err != nil {
		log.Println("Encountered error reading dht11", err)
	}

	// Print temperature and humidity
	log.Printf("Read DHT11: Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)

	s := fmt.Sprintf("Temperature = %v*C, Humidity = %v%%", temperature, humidity)
	storeInFile([]byte(s))
}

func storeInFile(bytes []byte) {
	if err := ioutil.WriteFile("/tmp/dht", bytes, 0644); err != nil {
		log.Println("Could not store new data to file")
	}
}
