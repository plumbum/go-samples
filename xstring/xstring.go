package main

import (
	"strings"
	"fmt"
)

type Xstring string

func (xs *Xstring) String() string {
	return string(*xs)
}

func (xs *Xstring) TrimSpaceLeft() *Xstring {
	*xs = Xstring(strings.TrimLeft(xs.String(), " \r\n\t\000"))
	return xs
}

func (xs *Xstring) TrimSpaceRight() *Xstring {
	*xs = Xstring(strings.TrimRight(xs.String(), " \r\n\t\000"))
	return xs
}

func main() {

	var s Xstring = "      Hello world!     "
	fmt.Println("[",s,"]")
	fmt.Println("[",s.TrimSpaceRight(),"]")
	fmt.Println("[",s.TrimSpaceLeft(),"]")

}