package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("Please provide a command.")
		fmt.Println("beamer list : list all available serial ports")
		fmt.Println("beamer start [port] : start beamer on serial port")
		fmt.Println("beamer stop [port] : stop beamer on serial port")
		return
	}
	cmd := arg[1]
	switch cmd {
	case "list":
		listPorts()
	case "power":
		if len(arg) < 3 {
			fmt.Println("Error: No port declared")
			fmt.Println("beamer start [port]")
			return
		}
		readPowerState(arg[2])
		case "start":

		if len(arg) < 3 {
			fmt.Println("Error: No port declared")
			fmt.Println("beamer start [port]")
			return
		}
		startBeamer(arg[2])
	case "stop":
		if len(arg) < 3 {
			fmt.Println("Error: No port declared")
			fmt.Println("beamer start [port]")
			return
		}
		stopBeamer(arg[2])
	default:
		fmt.Println("Please provide a valid command.")
		fmt.Println("beamer list : list all available serial ports")
		fmt.Println("beamer start [port] : start beamer on serial port")
		fmt.Println("beamer stop [port] : stop beamer on serial port")
	}
}

func listPorts() {
	fmt.Println("List all available serial ports")
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("USB serial %s\n", port.SerialNumber)
		}
	}
}


func startBeamer(port string) {
	writToBeamer(port, []byte("~XX00 1"))
}

func stopBeamer(port string) {
	data, _ := hex.DecodeString("7E303030302032")
	writToBeamer(port, data)
}

func readPowerState(port string)  {
	data, _ := hex.DecodeString("7E30303132342031")

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
	b, err := port1.Read(data)
	if err != nil {
		fmt.Println("Erro reading power state")
		fmt.Println(err)
	}
	fmt.Println(string(b))
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
	status, err := port1.GetModemStatusBits()
	if err != nil {
		fmt.Println("Error getting status bytes")
		return
	}
	fmt.Println(status)
	b, err := port1.Write(data)
	if err != nil {
		fmt.Println("Could not send data to port "+port)
		fmt.Println(err)
	} else {
		fmt.Println("Send "+strconv.Itoa(b)+" bytes to port "+port)
	}
	err = port1.Close()
	if err != nil {
		fmt.Println("Error closing serial port")
		fmt.Println(err)
	}
	fmt.Println("Write to beamer complete. Port closed")
}