package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sort strings")
	a := [...]string{"first", "second", "third", "fourth", "fifth"}
	fmt.Println("  Unsorted:", a)
	sort.Strings(a[:])
	fmt.Println("    Sorted:", a)

	fmt.Println("Sort strings")
	nums := []string{"2", "7", "10", "33", "101"}
	fmt.Println("  Unsorted:", nums)
	sort.Strings(nums)
	fmt.Println("    Sorted:", nums)

	fmt.Println("Sort integers")
	ints := []int{7, 2, 10, 33, 4, -9, 0, 123, 74}
	fmt.Println("  Unsorted:", ints)
	sort.Ints(ints)
	fmt.Println("    Sorted:", ints)
}
