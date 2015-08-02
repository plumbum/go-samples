package main

import (
	"log"
	"net"
	"fmt"
)

func main() {
	ln, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go connectionHandler(conn)
	}
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("Message received: "))
	conn.Write(buf)
	fmt.Println("Read: ", string(buf))
}