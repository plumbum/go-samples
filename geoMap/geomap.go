package main

import (
	"github.com/plumbum/go-samples/geoMap/yandexmap"
	"fmt"
	"os"
	"strings"
	"github.com/kr/pretty"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	api := os.Getenv("YANDEX_MAP_API_KEY")
	fmt.Println("API KEY:", api)

	address := "СПб, пр. Пятилеток, 15-1"
	fmt.Println("Dial address:", address)

	pn, err := yandexmap.Geocode(address, api)
	if err != nil {
		fmt.Println("Request error:", err)
	}

	printDelim("github.com/kr/pretty")
	pretty.Println(pn)

	printDelim("github.com/davecgh/go-spew/spew")
	spew.Dump(pn)

}

func printDelim(s string) {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println(s)
	fmt.Println(strings.Repeat("=", 80))
}
