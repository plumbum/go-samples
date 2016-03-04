package main

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/charmap"
	"strings"
	"io/ioutil"
	"fmt"
	"bytes"
	"github.com/fiam/gounidecode/unidecode"
)

// Encoding UTF-8 to Cp1251 and back
func main() {
	var err error

	encoder := charmap.Windows1251.NewEncoder()
	decoder := charmap.Windows1251.NewDecoder()

	inUtf8 := "Ёжики пушистые 好 ἱερογλυφικὰ γράμματ"
	// inUtf8 := "Ёжики пушистые"

	sr := strings.NewReader(inUtf8)
	tr := transform.NewReader(sr, encoder)
	inCp1251, err := ioutil.ReadAll(tr)
	if err != nil {
		fmt.Println("Encoding error: ", err)
	}

	srBack := bytes.NewReader(inCp1251)
	trBack := transform.NewReader(srBack, decoder)
	outUtf8, err := ioutil.ReadAll(trBack)
	if err != nil {
		fmt.Println("Decoding error: ", err)
	}

	fmt.Println("Source UTF8:", inUtf8)
	fmt.Println("CP1251:", inCp1251, string(inCp1251))
	fmt.Println("Result UTF8:", string(outUtf8))

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Test https://github.com/fiam/gounidecode")
	fmt.Println("Original: ", inUtf8)
	fmt.Println("Translit: ", unidecode.Unidecode(inUtf8))
}

