package main

import (
	"go-samples/geoMap/yandexmap"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
)

func main() {

	api := os.Getenv("YANDEX_MAP_API_KEY")
	fmt.Println("API KEY:", api)

	address := "СПб, пр. Пятилеток, 15-1"
	fmt.Println("Dial address:", address)

	pn, err := yandexmap.Geocode(address, api)
	fmt.Println(err)
	spew.Dump(pn)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
