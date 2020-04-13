package main

import "fmt"
import "math"

type screen struct {
	diag    float64
	xaspect float64
	yaspect float64
	x       float64
	y       float64
}

func main() {
	// fmt.Printf("%g\n", 1.000)
	// %g removes trailing zeroes, very handy %+v shows struct field names, %9.2 width 9, 2 decimal points

	fmt.Print(
		"Macbook Pro ", screenSize(13, 16, 10), "\n",
		"Lenovo X250 ", screenSize(12.5, 16, 9), "\n",
		"iPad ", screenSize(10, 3, 4), "\n",
	)

	// print(math.Sqrt(7.3*7.3+9.7*9.7))

}

/*
func screenSz(a string) string {
	screenSz("13 inch 16:10") >> screenSize(13,16,10)
}*/

func screenSize(diag, xaspect, yaspect float64) screen {
	scaler := math.Sqrt(diag * diag / (xaspect*xaspect + yaspect*yaspect))
	return screen{x: scaler * xaspect, y: scaler * yaspect, diag: diag, xaspect: xaspect, yaspect: yaspect}
}

func (s screen) String() string {
	return fmt.Sprintf("%v inch %v:%v width: %.1f height: %.1f", s.diag, s.xaspect, s.yaspect, s.x, s.y)
}
