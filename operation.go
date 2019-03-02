package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"strings"
	"time"
)

type Operation struct {
	OperationName OperationName
	Duration      int
}

type OperationName int

const (
	On = iota
	Off
	Blink
)

func (operationName OperationName) String() string {
	switch operationName {
	case On:
		return "ON"
	case Off:
		return "OFF"
	case Blink:
		return "BLINK"
	}
	return ""
}

func OperationNameFromString(s string) *OperationName {

	var res OperationName
	switch strings.ToUpper(s) {
	case "ON":
		res = On
		return &res
	case "OFF":
		res = Off
		return &res
	case "BLINK":
		res = Blink
		return &res
	}
	return nil
}

func processOperation(operation Operation) {

	log.Println("Starting Operation (", OpCounter, ")", operation.OperationName.String(), "about to begin for", operation.Duration, "seconds!")
	switch operation.OperationName {
	case Off:
		operationOff(operation.Duration)
	case On:
		operationOn(operation.Duration)
	case Blink:
		operationBlink(operation.Duration)
	}

	log.Println("Finished Operation (", OpCounter, ")", operation.OperationName.String(), "!")
	OpCounter++
}

func operationOn(duration int) {
	log.Println("Opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()

	pin.High()
	time.Sleep(time.Duration(int64(duration)) * time.Second)
}

func operationOff(duration int) {
	log.Println("Opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()

	pin.Low()
	time.Sleep(time.Duration(int64(duration)) * time.Second)
}

func operationBlink(duration int) {

	log.Println("Opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()

	for x := 0; x < duration*5; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
