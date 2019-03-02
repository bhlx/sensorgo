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
	ReverseBlink
)

func (operationName OperationName) String() string {
	switch operationName {
	case On:
		return "ON"
	case Off:
		return "OFF"
	case Blink:
		return "BLINK"
	case ReverseBlink:
		return "REVERSEBLINK"
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
	case "REVERSEBLINK":
		res = ReverseBlink
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
	case ReverseBlink:
		operationReversedBlink(operation.Duration)
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

	pin0 := rpio.Pin(18)
	pin0.Output()

	pin1 := rpio.Pin(23)
	pin1.Output()

	pin0.Low()
	pin1.Low()
	time.Sleep(time.Duration(int64(duration)) * time.Second)
}

func operationOff(duration int) {
	log.Println("Opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin0 := rpio.Pin(18)
	pin0.Output()

	pin1 := rpio.Pin(23)
	pin1.Output()

	pin0.High()
	pin1.High()
	time.Sleep(time.Duration(int64(duration)) * time.Second)
}

func operationBlink(duration int) {

	log.Println("Opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin0 := rpio.Pin(18)
	pin0.Output()
	pin0.High()

	pin1 := rpio.Pin(23)
	pin1.Output()
	pin1.High()

	for x := 0; x < duration*5; x++ {
		pin0.Toggle()
		pin1.Toggle()
		time.Sleep(time.Second / 5)
	}
}

func operationReversedBlink(duration int) {

	log.Println("Opening gpio")

	if err := rpio.Open(); err != nil {
		log.Fatal("unable to open gpio", err)
	}

	defer rpio.Close()

	pin0 := rpio.Pin(18)
	pin0.Output()
	pin0.Low()

	pin1 := rpio.Pin(23)
	pin1.Output()
	pin1.High()

	for x := 0; x < duration*5; x++ {
		pin0.Toggle()
		pin1.Toggle()
		time.Sleep(time.Second / 5)
	}
}
