package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

func main() {
	fmt.Println("Arduino RC Controller Serial Monitor")

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}

	for {

		line := ""
		buff := make([]byte, 1)
		on := true
		for on != false {
			line = ""
			for {
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
				}
				if n == 0 {
					fmt.Println("\nEOF")
					break
				}
				line = line + string(buff[:n])
				if strings.Contains(string(buff[:n]), "\n") {
					break
				}

			}

			fmt.Print("\033[u\033[K")
			fmt.Printf(" - %s\n", line)

		}
	}
}
