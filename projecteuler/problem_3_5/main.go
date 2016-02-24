package main

import (
	"fmt"
	"math/rand"
	"time"
)

func find(num int, chain string) (rchain string, ok bool) {
	if num == 1 {
		return chain, true
	}
	if num%3 == 0 {
		if rchain, ok = find(num/3, "*3"+chain); ok {
			return
		}
	}
	if num-5 >= 1 {
		if rchain, ok = find(num-5, "+5"+chain); ok {
			return
		}
	}
	return chain, false
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := range [100]struct{}{} {
		// income := rand.Int() % 10000 + 1
		fmt.Println("Income: ", i+1)
		chain, ok := find(i+1, "")
		if ok {
			fmt.Println("Chain: 1", chain)
		} else {
			fmt.Println("Chain not found.")
		}
	}
}
