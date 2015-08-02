package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [...]string{"first", "second", "third", "fourth", "fifth"}
	fmt.Println(a)
	sort.Strings(a[0:len(a)])
	fmt.Println(a)

	i := [...]int{7, 2, 10, 33, 4, -9, 0, 123, 74}
	fmt.Println(i)
	sort.Ints(i[0:len(i)])
	fmt.Println(i)
}
