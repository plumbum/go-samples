package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	counter := decimal.New(0, 0)
	inc := decimal.New(1, -2)
	fmt.Println("Increment: ", inc)
	for i:=0; i<1000; i++ {
		counter = counter.Add(inc)
	}
	fmt.Println("Total: ", counter)
	counter = counter.Add(decimal.New(1, 16))
	fmt.Println("Total: ", counter)
	counter = counter.Add(decimal.New(1, -10))
	fmt.Println("Total: ", counter)
}
