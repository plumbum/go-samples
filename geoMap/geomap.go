package main

import (
	"go-samples/geoMap/yandexmap"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
)

func main() {

	api := ""
	cfg, err := os.Open("api.cfg")
	if err == nil {
		defer cfg.Close()
		fmt.Fscanln(cfg, &api)
	}

	address := "СПб, пр. Пятилеток, 15-1"

	pn, err := yandexmap.Geocode(address, api)
	fmt.Println(err)
	spew.Dump(pn)

	/*
	var gc geo.GoogleGeocoder

	pt, err := gc.Geocode(address)
	chk(err)

	spew.Dump(pt)
	*/


	/*
	var rmap map[string]*json.RawMessage
	json.Unmarshal(data, &rmap)
	spew.Dump(rmap)
	*/

	/*
	spew.Dump(resp.StatusCode)
	spew.Dump(resp.Status)
	spew.Dump(resp.ContentLength)
	spew.Dump(resp.Proto)
	spew.Dump(resp.Header)
	spew.Dump(resp.Cookies())
	spew.Dump(resp.Close)
	*/


	/*
	fmt.Println(string(data))
	*/

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
