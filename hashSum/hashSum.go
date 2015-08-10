package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
)

func main() {
	s := "Hello world!"
	b := []byte(s)

	md5Sum := md5.Sum(b)
	fmt.Println("MD5   :", myHex(md5Sum[:]))
	fmt.Println("      :", myBase64(md5Sum[:]))

	sha1Sum := sha1.Sum(b)
	fmt.Println("SHA1  :", myHex(sha1Sum[:]))
	fmt.Println("      :", myBase64(sha1Sum[:]))

	sha224Sum := sha256.Sum224(b)
	fmt.Println("SHA224:", myHex(sha224Sum[:]))
	fmt.Println("      :", myBase64(sha224Sum[:]))

	sha256Sum := sha256.Sum256(b)
	fmt.Println("SHA256:", myHex(sha256Sum[:]))
	fmt.Println("      :", myBase64(sha256Sum[:]))

}

func myHex(b []byte) string {
	return hex.EncodeToString(b)
}

func myBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
