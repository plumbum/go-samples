package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {

	// Creating UUID Version 1
	// UUIDv1: 81c7d93d-f581-11e5-ba55-8c89a5c1a947
	u1 := uuid.NewV1()
	fmt.Printf("UUIDv1: %s\n", u1)

	// Creating UUID Version 2
	// UUIDv2: 000003e8-f581-21e5-ba00-8c89a5c1a947
	u2 := uuid.NewV2(0)
	fmt.Printf("UUIDv2: %s\n", u2)

	// Creating UUID Version 3
	// UUIDv3: a00172db-706e-3ac5-97cf-50d25b514754
	u3 := uuid.NewV3(u2, "MD5")
	fmt.Printf("UUIDv3: %s\n", u3)

	// Creating UUID Version 4
	// UUIDv4: a0ef3801-651a-44d7-a6f5-65c26bf77d1d
	u4 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u4)

	// Creating UUID Version 5
	// UUIDv5: bf50b865-0487-5022-ba8b-6fb7272ff487
	u5 := uuid.NewV5(u4, "SHA-1")
	fmt.Printf("UUIDv5: %s\n", u5)

	// Parsing UUID from string input
	uX, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something gone wrong: %s\n", err)
	}
	fmt.Printf("Successfully parsed: %s\n", uX)
}
