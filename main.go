package main

import (
	"log"
	"os"
	"strconv"

	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		log.Println("Please provide a command.")
		log.Println("beamer list : list all available serial ports")
		log.Println("beamer start [port] : start beamer on serial port")
		log.Println("beamer stop [port] : stop beamer on serial port")
		return
	}
	cmd := arg[1]
	switch cmd {
	case "list":
		listPorts()
		case "start":
		if len(arg) < 3 {
			log.Println("Error: No port declared")
			log.Println("beamer start [port]")
			return
		}
		startBeamer(arg[2])
	case "stop":
		if len(arg) < 3 {
			log.Println("Error: No port declared")
			log.Println("beamer start [port]")
			return
		}
		stopBeamer(arg[2])
	default:
		log.Println("Please provide a valid command.")
		log.Println("beamer list : list all available serial ports")
		log.Println("beamer start [port] : start beamer on serial port")
		log.Println("beamer stop [port] : stop beamer on serial port")
	}
}

func listPorts() {
	log.Println("List all available serial ports")
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		log.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			log.Printf("USB ID     %s:%s\n", port.VID, port.PID)
			log.Printf("USB serial %s\n", port.SerialNumber)
		}
	}
}

func startBeamer(port string) {
	writToBeamer(port, []byte("~0000 1\r"))
}

func stopBeamer(port string) {
	writToBeamer(port, []byte("~0000 0\r"))
}

func writToBeamer(port string, data []byte) {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity: serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port1, err := serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}
	b, err := port1.Write(data)
	if err != nil {
		log.Println("Could not send data to port "+port)
		log.Println(err)
	} else {
		log.Println("Send "+strconv.Itoa(b)+" bytes to port "+port)
	}
	err = port1.Close()
	if err != nil {
		log.Println("Error closing serial port")
		log.Println(err)
	}
	log.Println("Port closed")
}