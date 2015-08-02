package main

import (
	"reflect"
	"fmt"
	"strings"
)

type T struct {
	A int		`json:"id"`
	B string	`json:"name"`
}

func typeOfVar(v interface{}) {
	fmt.Println("Type of variable", v, "is", reflect.TypeOf(v))
}

func main() {

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Variables")

	x := 3.1415926
	typeOfVar(x)
	y := 3
	typeOfVar(y)

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Do structure reflect")

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v (Tag: %s)\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface(),
				typeOfT.Field(i).Tag)
	}

}

// Look other sample https://gist.github.com/drewolson/4771479
