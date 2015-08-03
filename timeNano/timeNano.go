package main

import (
	"time"
	"math"
	"fmt"
)

func main() {

	startNs := time.Now().UnixNano()

	sp := 0.0
	for i:=0; i<10000000; i++ {
		sp += math.Sin(float64(i)/math.Pi)
	}

	endNs := time.Now().UnixNano()

	fmt.Println("Calculate time ", float64(endNs-startNs)/1000000.0, " ms")
	fmt.Println("Summ:", sp)

}
