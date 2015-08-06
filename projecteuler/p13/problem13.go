// Large sum
// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
// look file: numbers.txt
// https://projecteuler.net/problem=13

package main

import (
	"math/big"
	"os"
	"bufio"
	"fmt"
)

func main() {

	// Open file for read
	fd, err := os.Open("numbers.txt")
	chk(err)
	defer fd.Close()

	// Line by line reader
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	// Use standart library - big Integers
	bigSum := new(big.Int)

	for scanner.Scan() {
		ns := scanner.Text()
		bigInt := new(big.Int)
		bigInt.SetString(ns, 10) // Convert readed decimal string to big.Int
		bigSum.Add(bigSum, bigInt)
	}

	answerString := bigSum.String()
	fmt.Println("Result:", answerString, len(answerString))
	fmt.Println("Answer:", answerString[0:10])
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
