package main

import (
	"github.com/ventu-io/go-shortid"
	"fmt"
)

func main() {
	for i := range [100]struct{}{} {
		id, err := shortid.MustGenerate()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%3d: %s\n", i, id)
	}
}
