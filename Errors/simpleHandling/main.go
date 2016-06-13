package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

var cnt int

func subErr(err error) error {
	err = errors.Wrapf(err, "Suberror #%d", cnt)
	cnt++
	return err
}

func getError() error {
	err := errors.New("my error")
	err = subErr(err)
	return errors.Wrap(err, "exit")
}

func main() {
	fmt.Println("Hello world!")
	err := getError()
	err = errors.Wrap(err, "open failed")
	err = subErr(err)
	err = errors.Wrap(err, "read config failed")

	pp.Println("Cause: ", errors.Cause(err))
	err = errors.Wrap(err, "New message")
	pp.Println("Error: ", err)
	fmt.Printf("[%+v]\n", err)
	fmt.Printf("{%+v}\n", errors.Cause(err))
}
