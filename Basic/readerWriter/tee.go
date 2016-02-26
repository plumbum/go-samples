package main

import (
	"os"
	"time"
	"strings"
	"io"
	"io/ioutil"
	"fmt"
)

func main() {

	s := time.Now().String()
	in := strings.NewReader(s)

	// Write to file
	out, err := os.OpenFile("output.txt", os.O_CREATE | os.O_WRONLY, 0666)
	chk(err)
	defer out.Close()

	tee := io.TeeReader(in, out)

	data, err := ioutil.ReadAll(tee)
	chk(err)
	fmt.Println(string(data))
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
