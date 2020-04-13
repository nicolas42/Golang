package main

import "fmt"
import "math"

type screen struct {
	x float64
	y float64
	diag float64
	xaspect float64
	yaspect float64
}

func main() {
	fmt.Printf("%g\n", 1.000) 
	// %g removes trailing zeroes, very handy %+v shows struct field names, %9.2 width 9, 2 decimal points
	
	fmt.Printf("%+v\n%+v\n%+v\n",screenSize(13,16,10),screenSize(12.5,16,9),screenSize(10,4,3))
	fmt.Println(screenSize(13,16,10),screenSize(12.5,16,9),screenSize(10,4,3))

	fmt.Println("So, the macbook pro provides ~7inches of browsing height, lenovo x250 ~6 inches, ipad ~8 inches")
	
	fmt.Println("Additionally, an inch or so can be removed for both top and bottom menu bars")

}

func (s screen) String() string {
	return fmt.Sprintf("%+v",s)
}

func screenSize(diag, xaspect, yaspect float64) screen {
	x := math.Sqrt(diag*diag/(xaspect*xaspect + yaspect+yaspect))
	return screen{diag: diag, x: xaspect*x, y: yaspect*x, xaspect: xaspect, yaspect: yaspect}
}
