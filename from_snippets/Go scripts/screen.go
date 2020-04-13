package screen

// screen.Size(13,16,10)

import "fmt"
import "math"

type screen struct {
	diag    float64
	xaspect float64
	yaspect float64
	x       float64
	y       float64
}

func Size(diag, xaspect, yaspect float64) screen {
	// transforms diagonal length and an aspect ratio to a width and height and returns a screen object with those contents
	scaler := math.Sqrt(diag * diag / (xaspect*xaspect + yaspect*yaspect))
	return screen{x: scaler * xaspect, y: scaler * yaspect, diag: diag, xaspect: xaspect, yaspect: yaspect}
}

func (s screen) String() string {
	return fmt.Sprintf("%v inch %v:%v width: %.1f height: %.1f", s.diag, s.xaspect, s.yaspect, s.x, s.y)
}



