package main

import (
	"encoding/json"
	"fmt"
	"github.com/plumbum/go-samples/Basic/json/data"
)

func main() {

	ta := &data.Users{}
	ta.Users = []data.Item{
		{1, "John", "123"},
		{2, "Bob", "123"},
		{3, "Mary", "123"},
		{4, "Dave", "123"},
		{5, "Ken", "123"},
	}

	str, err := json.Marshal(ta)
	chk(err)
	fmt.Println(string(str))

	str2, err := ta.MarshalJSON()
	chk(err)
	fmt.Println(string(str2))

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
