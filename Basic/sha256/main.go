package main

import (
	"fmt"
	"io/ioutil"
	"crypto/sha256"

	b58 "github.com/jbenet/go-base58"
	"github.com/jbenet/go-multihash"
	"crypto/sha512"
	"encoding/base64"
	"os"
)

func main() {

	dat, err := ioutil.ReadFile(os.Args[0])
	check(err)

	// Calculate sums
	s256 := sha256.Sum256(dat)
	s512 := sha512.Sum512(dat)

	// Calculate multihash
	h, err := multihash.Encode(s256[:], multihash.SHA2_256)
	check(err)
	multiSha256 := b58.Encode(h)

	fmt.Println("Multihash256/58: ", len(multiSha256), string(multiSha256))

	// Base 58 encoding
	s256b58 := b58.Encode(s256[:])
	s512b58 := b58.Encode(s512[:])

	// Base 64 encoding
	enc := base64.RawURLEncoding

	s256b64 := enc.EncodeToString(s256[:])
	s512b64 := enc.EncodeToString(s512[:])

	fmt.Println("SHA256/base58  : ", len(s256b58), string(s256b58))
	fmt.Println("SHA256/base64  : ", len(s256b64), string(s256b64))
	fmt.Println("SHA512/base58  : ", len(s512b58), string(s512b58))
	fmt.Println("SHA512/base64  : ", len(s512b64), string(s512b64))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
