package main

import "fmt"
import "math"

func main() {
	// radius, phi1, theta1, phi2, theta2
	dist := DistanceOnSphereDegrees(1, 180, 30, 50, math.Pi/2) 
	fmt.Printf("%vpi", dist/math.Pi)
}

func DistanceOnSphere(r, phi1, theta1, phi2, theta2 float64) float64 {

	v1 := []float64{r * math.Sin(phi1) * math.Cos(theta1), r * math.Sin(phi1) * math.Cos(math.Pi/2-theta1), r * math.Cos(phi1)}
	v2 := []float64{r * math.Sin(phi2) * math.Cos(theta2), r * math.Sin(phi2) * math.Cos(math.Pi/2-theta2), r * math.Cos(phi2)}
	angle := math.Acos(DotProduct(v1, v2) / math.Pow(r, 2))
	d := angle * r
	return d
}

func DistanceOnSphereDegrees(r, phi1, theta1, phi2, theta2 float64) float64 {

	phi1 = phi1 / 180 * math.Pi
	theta1 = theta1 / 180 * math.Pi
	phi2 = phi2 / 180 * math.Pi
	theta2 = theta2 / 180 * math.Pi

	return DistanceOnSphere(r, phi1, theta1, phi2, theta2)
}

//dot product of two slices
func DotProduct(a []float64, b []float64) float64 {
	c := float64(0)
	for i, _ := range a {
		c = c + (a[i] * b[i])
	}
	return c
}




