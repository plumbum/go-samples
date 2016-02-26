// JavaScript interpreter
// https://github.com/robertkrimen/otto

package main

import (
	"github.com/robertkrimen/otto"
	"fmt"
	"time"
)

func main() {

	vm := otto.New()

	vm.Set("dt", time.Now().String())

	fmt.Println(vm.Run(`
		abc = 2 + 2
		console.log("Value = "+abc)
		console.log("Date = "+dt)
		"return value"
	`))

	fmt.Println(vm.Get("abc"))
}
