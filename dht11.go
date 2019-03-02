package main

import (
	"github.com/d2r2/go-dht"
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

	CurrentDHT11State = DHT11State{
		temperature: temperature,
		humidity:    humidity,
	}

}
