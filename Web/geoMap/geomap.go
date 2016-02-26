package main

import (
	"github.com/plumbum/go-samples/Web/geoMap/yandexmap"
	"fmt"
	"os"
	"github.com/k0kubun/pp"
	// "github.com/kr/pretty"
	// "github.com/davecgh/go-spew/spew"
)

func main() {

	api := os.Getenv("YANDEX_MAP_API_KEY")
	fmt.Println("API KEY:", api)

	address := "СПб, пр. Пятилеток, 15-1"
	fmt.Println("Dial address:", address)

	position, err := yandexmap.Geocode(address, api)
	if err != nil {
		fmt.Println("Request error:", err)
	}

	pp.Println(position)
}

