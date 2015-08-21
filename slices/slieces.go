package main

import (
	"fmt"
)

const InitialCapacity = 8

func main() {

	// Создадим пустой слайс но с начальной ёмкостью.
	// Можно поигратсься с ёмкость, что бы посмотреть на реакцию слайса
	slice := make([]int, 0, InitialCapacity)
	fmt.Printf("SOURCE: cap=%d; len=%d; %v\r\n", cap(slice), len(slice), slice)

	// Начинаем заполнять слайс
	for i:=0; i<32; i++ {
		// append при недостатке ёмкости увеличивает её вдвое
		slice = append(slice, i)
		fmt.Printf("%5d: cap=%d; len=%d; %v\r\n", i, cap(slice), len(slice), slice)
	}

	// Append element to sliece
	slice = append(slice, 99)
	fmt.Printf("Append element: cap=%d; len=%d; %v\r\n", cap(slice), len(slice), slice)
	// Append sliece to sliece
	aslice := []int{201,202,203,204}
	slice = append(slice, aslice...)
	fmt.Printf("Append slice: cap=%d; len=%d; %v\r\n", cap(slice), len(slice), slice)

	// String sliced
	str := "Hello world!"
	fmt.Println(str[6:11])
}

