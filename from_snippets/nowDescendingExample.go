package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("20060102-150405 MST"))
	// time.Format() takes a reference time
	// Since MST is GMT-0700, the reference time can be thought of as
	// 01/02 03:04:05PM '06 -0700
	// or 20060102-150405.000000 -07
	// https://golang.org/pkg/time/#pkg-constants

	x := big.NewInt(1)
	one := big.NewInt(1)
	y := big.NewInt(1)
	for i := 1; i <= 1000; i += 1 {
		x.Mul(x, y)
		y.Add(y, one)
	}
	fmt.Println(x.String())
}
