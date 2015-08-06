// https://github.com/cznic/mathutil

package main

import (
	"github.com/cznic/mathutil"
	"fmt"
)

func minmax(a, b int) {
	fmt.Println("Min:", mathutil.Min(a, b))
	fmt.Println("Max:", mathutil.Max(a, b))
}

func main() {

	minmax(100, 77)
	fmt.Println("ISqrt 65535 = ", mathutil.ISqrt(65535))
	fmt.Println("Log2 70000 = ", mathutil.Log2Uint32(70000))
	fmt.Println("BitLen 70 =", mathutil.BitLen(70))
	fmt.Println("BitLen 700 =", mathutil.BitLen(700))
	fmt.Println("BitLen 7000 =", mathutil.BitLen(7000))
	fmt.Println("BitLen 70000 =", mathutil.BitLen(70000))
	fmt.Println("PopCount 70000 = ", mathutil.PopCount(70000))
	fmt.Println("PopCount 70001 = ", mathutil.PopCount(70001))

}
