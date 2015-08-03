package main

import (
    "fmt"
    "github.com/satori/go.uuid"
)

func main() {
    // Creating UUID Version 1
    u1 := uuid.NewV1()
    fmt.Printf("UUIDv1: %s\n", u1)

    // Creating UUID Version 2
    u2 := uuid.NewV2(0)
    fmt.Printf("UUIDv2: %s\n", u2)

    // Creating UUID Version 3
    u3 := uuid.NewV3(u2, "MD5")
    fmt.Printf("UUIDv3: %s\n", u3)

    // Creating UUID Version 4
    u4 := uuid.NewV4()
    fmt.Printf("UUIDv4: %s\n", u4)

    // Creating UUID Version 5
    u5 := uuid.NewV5(u4, "SHA-1")
    fmt.Printf("UUIDv5: %s\n", u5)

    // Parsing UUID from string input
    uX, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    if err != nil {
        fmt.Printf("Something gone wrong: %s\n", err)
    }
    fmt.Printf("Successfully parsed: %s\n", uX)
}

