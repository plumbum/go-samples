package main

import (
	"github.com/goburrow/modbus"
	"time"
	"log"
	"fmt"
)

func main() {

	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0") // I use FT232 based handmade converter
	handler.BaudRate = 19200
	handler.DataBits = 8
	handler.Parity = "O" // Odd parity
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 2 * time.Second

	err := handler.Connect()
	chk(err)
	defer handler.Close()

	client := modbus.NewClient(handler)
	for i := uint16(0); i<7; i++ {
		results, err := client.WriteSingleCoil(64512+i, 0xFF00)
		chk(err)
		fmt.Println(results)
		results1, err := client.WriteSingleCoil(64512+i, 0)
		chk(err)
		fmt.Println(results1)
	}
}

func chk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Read about protocol: https://ru.wikipedia.org/wiki/Modbus
