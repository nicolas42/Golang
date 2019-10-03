	// Demo a+bx least squares fitting as per
// http://mathworld.wolfram.com/LeastSquaresFitting.html


package main

import (
	"fmt"
)
	
func main() {
	fmt.Println("Hello, playground")

	x := []float64{1,2,3,4,5,6,7,8,9}
	y := []float64{1,2,3,4,5,6,7,8,9}
	var n,sx,sxx,sy,sxy float64
	var a,b,c,d float64
	var a0,a1 float64
	
	n = float64(len(x))
	
	for i := 0; i < len(x); i++ {
		sx += x[i]
		sxx += x[i]*x[i]
		sy += y[i]
		sxy += x[i]*y[i]
	}
	// Want a0 and a1 for line of best fit: y ~ a0 + a1*x
	// [n sx; sx sxx] [a0; a1] = [sy; sxy]  (source: http://mathworld.wolfram.com/LeastSquaresFitting.html)
	// rename matrix and solve coefficients 
  	// [a b; c d] [a0; a1] = [sy; sxy]
	// [a0; a1] = 1/(ad-bc) * [d -b; -c a] * [sy; sxy]
	
	a=n; b=sx; c=sx; d=sxx;
	a0 = 1/(a*d-b*c) * (d*sy - b*sxy)
	a1 = 1/(a*d-b*c) * (-c * sy + a*sxy)
	
	fmt.Println(a0,a1)
}

	
	
