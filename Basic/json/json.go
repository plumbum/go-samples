package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	A int		`json:"id"`
	B string	`json:"name"`
}

func main() {

	ta := [...]T{
		{1, "John"},
		{2, "Bob"},
		{3, "Mary"},
		{4, "Dave"},
		{5, "Ken"},
	}

	str, err := json.Marshal(ta)
	chk(err)
	fmt.Println(string(str))

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
