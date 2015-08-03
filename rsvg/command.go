package main

import (
	"os/exec"
	"os"
	"text/template"
	"time"
	"bytes"
)


type Data struct {
	Title string
	Date string
	Number int
	String string
	Boolean bool
}

func main() {

	data := Data{
		"SVG -> PDF",
		time.Now().Format("_2 Jan 2006 15:04:05"),
		20150803,
		"Строка кириллицей",
		true,
	}

	var err error

	// Load template from file
	template, err := template.ParseFiles("gopher.svg")
	chk(err)

	// Store template to buffer
	buf := new(bytes.Buffer)
	err = template.Execute(buf, data)
	chk(err)

	// Write pdf to file
	out, err := os.OpenFile("gopher.pdf", os.O_CREATE | os.O_WRONLY, 0666)
	chk(err)
	defer out.Close()

	// Convert via external application
	// Install before use
	// # apt-get install librsvg2-bin
	cmd := exec.Command("rsvg-convert", "-f", "pdf")
	cmd.Stdin = bytes.NewReader(buf.Bytes())
	cmd.Stdout = out
	err = cmd.Run() // Syncronous run external application
	chk(err)

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
