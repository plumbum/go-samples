package main

// https://github.com/comail/colog

import (
	"comail.io/go/colog"
	"log"
	"strings"
	"fmt"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	colog.Register()
	log.Print("register colog")

	colog.SetDefaultLevel(colog.LWarning)
	log.Print("set default level LWarning")

	colog.SetFlags(log.Ldate)
	log.Print("Only date")
	printLogs()

	colog.SetDefaultLevel(colog.LInfo)
	log.Print("set default level LInfo")
	colog.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Print("Date, time, filename")
	printLogs()

	colog.ParseFields(true)
	log.Print("Parse fileds like key=value")
	printLogs()

	colog.SetMinLevel(colog.LWarning)
	log.Print("warn: set minimal level LWarning")
	printLogs()

}

func printLogs() {
	fmt.Println(strings.Repeat("=", 80))
	log.Print("Default level")
	log.Print("debug: debug message")
	log.Print("info: Hello world!")
	log.Print("warn: has some fields key=value; one=two")
	log.Print("error: Some error")
	log.Print("alert: Good bye")
	fmt.Println(strings.Repeat("-", 80))
}
