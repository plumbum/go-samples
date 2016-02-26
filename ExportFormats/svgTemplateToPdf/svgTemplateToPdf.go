package main

import (
	"os/exec"
	"os"
	"text/template"
	"time"
	"bytes"
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
)


type Data struct {
	Title string
	Date string
	Number int
	String string
	Boolean bool
}

var (
	converter = flag.String("conv", "inkscape", "Use converter [inkscape | rsvg]")
	inputFile = flag.String("i", "gopher.svg", "<input filename.svg>")
	outputFile = flag.String("o", "gopher.pdf", "<output filename.pdf")
	outputPng = flag.String("png", "", "<output filename.png> (only for inkscape converter)")
)

func main() {

	flag.Parse()

	data := Data{
		"SVG -> PDF",
		time.Now().Format("_2 Jan 2006 15:04:05"),
		20150803,
		"Строка кириллицей",
		true,
	}

	var err error

	// Load template from file
	template, err := template.ParseFiles(*inputFile)
	chk(err)

	// Store template to buffer
	buf := new(bytes.Buffer)
	err = template.Execute(buf, data)
	chk(err)

	// Convert via external application
	// Install before use
	// # apt-get install librsvg2-bin
	// or
	// # apt-get install inkscape
	var cmd *exec.Cmd

	if *converter == "inkscape" {
		fmt.Println("Generate via inkscape")
		cmd = exec.Command("inkscape", "--without-gui", "/dev/stdin", "--export-pdf=/dev/stdout")
		if *outputPng != "" {
			cmd.Args = append(cmd.Args, "--export-png", *outputPng)
		}
	} else {
		fmt.Println("Generate via rsvg-convert")
		cmd = exec.Command("rsvg-convert", "-f", "pdf")
	}
	cmd.Stdin = bytes.NewReader(buf.Bytes())

	// Write pdf to file
	out, err := os.OpenFile(*outputFile, os.O_CREATE | os.O_WRONLY, 0666)
	chk(err)
	defer out.Close()
	cmd.Stdout = out

	timeStart := time.Now().UnixNano()
	err = cmd.Run() // Syncronous run external application
	chk(err)
	timeEnd := time.Now().UnixNano()

	fmt.Println("Conversion time (ms)", (timeEnd - timeStart)/1000000)

	// Open output file using the OS's default application
	open.Run(*outputFile)

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
