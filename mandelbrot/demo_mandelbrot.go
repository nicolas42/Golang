package main

import (
	"bytes"
	"image/png"
	"log"
	"os"

	"./mandelbrot"
)

func main() {

	img := mandelbrot.Draw(1000, 1000)

	var buffer bytes.Buffer
	png.Encode(&buffer, img)

	f, err := os.Create("image.png")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	f.Write(buffer.Bytes())
	f.Close()
	println("Wrote mandelbrot image to image.png");

}

// func TestDraw(t *testing.T) {
// 	var v float64
// 	v = Average([]float64{1, 2})
// 	if v != 1.5 {
// 		t.Error("Expected 1.5, got ", v)
// 	}
// }
