// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
// https://projecteuler.net/problem=7

package main

import (
	"fmt"
	"time"
	"github.com/plumbum/intmath/u64"
)

const targetNumbers = 10001


func isPrime(n uint64, primes []uint64) bool {
	lim := u64.Sqrt(n) // Ignore all dividers greather than square root of number
	for _, p := range primes {
		if p > lim { // Okay, is prime
			return true
		}
		if n % p == 0 {
			return false
		}
	}
	return true

}

func main() {


	startTime := time.Now().UnixNano()

	cur := uint64(3)
	primes := make([]uint64, 1, targetNumbers)
	primes[0] = 2

	for {
		if isPrime(cur, primes) {
			primes = append(primes, cur)
			if len(primes) >= targetNumbers {
				break
			}
		}
		cur += 2 // Skip even numbers
	}

	endTime := time.Now().UnixNano()

	fmt.Println(primes)
	fmt.Println("Done", targetNumbers, "in", float64(endTime-startTime)/1000000, "ms")
	fmt.Println("Answer", primes[targetNumbers-1])

}
