package main

import (
	"log"
	"reflect"

	"go.bug.st/serial"
)

// HaloUSBID stores the USB ID string for Halo Scanner.
const HaloUSBID string = "04d8:00df"

type port struct {
	port serial.Port
}

func main() {
	mode := &serial.Mode{BaudRate: 115200}
	port, err := serial.Open("/dev/ttyACM0", mode)
	if err != nil {
		log.Fatalf("Failed to open port: %v\n", err)
	}

	msg := []byte("PA\r")
	n, err := port.Write(msg)
	if err != nil {
		log.Fatalf("Error writing to port: %v\n", err)
	}
	log.Printf("Wrote %v bytes: %v\n", n, msg)

	var got []byte
	buff := make([]byte, 4)
	for {
		n, err = port.Read(buff)
		if err != nil {
			log.Printf("Error reading from port: %v", err)
			break
		}
		if n == 0 {
			log.Printf("EOF\n")
			break
		}
		if buff[0] == 0 {
			break
		}
		got = append(got, buff[:n]...)
	}
	log.Printf("Received: %v\n", got)

	want := []byte("Halo")
	if !reflect.DeepEqual(want, got) {
		log.Println("Device is not a Halo Scanner!")
		log.Fatalf("Expected: %v, Received: %v\n", want, got)
	}
	log.Println("Device is Halo Scanner")

	// log.Println("Requesting microchip IDs")
	// msg = []byte("SC\r")
	// n, err = port.Write(msg)
	// if err != nil {
	// 	log.Fatalf("Failed writing to port: %v\n", err)
	// }
	// log.Printf("Wrote %v bytes: %v\n", n, msg)

	// buff = make([]byte, 6)
	// for {
	// 	n, err = port.Read(buff)
	// 	if err != nil {
	// 		log.Printf("Error reading from port: %v", err)
	// 		break
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// }
}
