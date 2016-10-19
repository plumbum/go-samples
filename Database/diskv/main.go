package main

import (
	"fmt"
	"github.com/peterbourgon/diskv"
	"strings"
)

func main() {
	// Simplest transform function: put all the data files into the base dir.
	flatTransform := func(s string) []string {
		ss := strings.Split(s, ".")
		return ss[0:len(ss)-1]
	}

	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	d := diskv.New(diskv.Options{
		BasePath:     "my-data-dir",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
		PathPerm: 0750,
		FilePerm: 0640,
	})

	// Write three bytes to the key "alpha".
	key := "alpha"
	d.Write(key, []byte{'1', '2', '3'})
	d.Write("beta", []byte{'4', '5', '6'})
	d.Write("sub.alfa", []byte("Hello alpa date"))
	d.Write("sub.omega", []byte("Good buy omega"))

	// Read the value back out of the store.
	value, _ := d.Read(key)
	fmt.Printf("%v\n", value)

	// Erase the key+value from the store (and the disk).
	d.Erase(key)
}
