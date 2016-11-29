package main

import (
	"encoding/json"
	"fmt"
	"github.com/plumbum/go-samples/Basic/json/data"
	"log"
	"github.com/k0kubun/pp"
)

type Item []int

func main() {
	var err error

	ta := &data.Users{}
	ta.Users = []data.Item{
		{1, "John", "123"},
		{2, "Bob", "123"},
		{3, "Mary", "123"},
		{4, "Dave", "123"},
		{5, "Ken", "123"},
	}

	log.Println("Marshal default")
	str, err := json.Marshal(ta)
	chk(err)
	fmt.Println(string(str))

	log.Println("Marshal ffjson")
	str2, err := ta.MarshalJSON()
	chk(err)
	fmt.Println(string(str2))

	log.Println("Unmarshal array")
	json_str := `[[1, 2, 3], [4, 5, 6], [7, 8, 9]]`
	a := make([]Item, 0)

	err = json.Unmarshal([]byte(json_str), &a)
	chk(err)
	pp.Println(a)

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
