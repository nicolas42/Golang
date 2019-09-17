package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	var x, y uint32
	var xmin, xmax, ymin, ymax float64
	var xscale, yscale, zx, zy, cx, cy, zxtemp float64
	var i uint32
	var isInSet bool

	const WIDTH = 100 // pixels
	const HEIGHT = 100

	// int data_length = 4 * WIDTH * HEIGHT;
	// char data[data_length];
	// // char* data = malloc(data_length);

	xmin = -2
	xmax = 2
	ymin = -2
	ymax = 2

	xscale = WIDTH / (xmax - xmin)
	yscale = HEIGHT / (ymax - ymin)

	const STRIDE = 3
	const IMAX = 40

	// black := [4]uint8{0, 0, 0, 255}
	// white := [4]uint8{255, 255, 255, 255}
	// color := [4]uint8{255, 255, 255, 255}

	// char* filename = "mandelbrot.ppm";

	// // Open file and write header
	// FILE *fp = fopen(filename, "wb"); /* b - binary mode */
	// (void) fprintf(fp, "P6\n%d %d\n255\n", WIDTH, HEIGHT);

	// Work out the other pixels
	for y = 0; y < HEIGHT; y += 1 {
		for x = 0; x < WIDTH; x += 1 {

			// image is flipped vertically. numbers go up, pixels go down
			cx = float64(x)/xscale - xmax
			cy = float64(y)/yscale - ymax
			zx = 0
			zy = 0
			isInSet = true

			for i = 1; i <= IMAX; i += 1 {
				zxtemp = zx*zx - zy*zy + cx
				zy = 2*zx*zy + cy
				zx = zxtemp

				if zx*zx+zy*zy > 4 {
					isInSet = false
					break
				}
			}

			if isInSet {
				// (void) fwrite(black, 1, 3, fp);
				print("*")

			} else {
				// fputc(white[0] * i / IMAX, fp);
				// fputc(white[1] * i / IMAX, fp);
				// fputc(white[2] * i / IMAX, fp);
				// fputc(white[3] * i / IMAX, fp);
				print(" ")
			}
		}
		print("\n")
	}

	return
}
