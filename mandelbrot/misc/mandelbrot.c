// gcc mandelbrot.c && ./a.out
// push(){git add . && git commit -m $1 && git push}

#include <stdio.h>
// #include <stdlib.h>
// #include <string.h>


int main(int argc, char** argv){

    static unsigned char black[4] = {0, 0, 0, 255};
    static unsigned char white[4] = {255, 255, 255, 255};
    char color[4] = {255, 255, 255, 255};

    int x;
    int y;
    int WIDTH = 100; // pixels
    int HEIGHT = 100;

	// int data_length = 4 * WIDTH * HEIGHT;
	// char data[data_length];
	// // char* data = malloc(data_length);
    

    double xmin = -2;
    double xmax = 2;
    double ymin = -2;
    double ymax = 2;

    double xscale = WIDTH / (xmax - xmin);
    double yscale = HEIGHT / (ymax - ymin);

    double zx;
    double zxtemp;
    double zy;
    double cx;
    double cy;

    int isInSet;
    int pos;
    int STRIDE = 3;
    int i;
    int IMAX = 40;

	// char* filename = "mandelbrot.ppm";


    // // Open file and write header
	// FILE *fp = fopen(filename, "wb"); /* b - binary mode */
	// (void) fprintf(fp, "P6\n%d %d\n255\n", WIDTH, HEIGHT);

    // Work out the other pixels
    for (y = 0; y < HEIGHT; y += 1) {
        for (x = 0; x < WIDTH; x += 1) {

            // image is flipped vertically. numbers go up, pixels go down
            cx = x / xscale - xmax;
            cy = y / yscale - ymax;
            zx = 0;
            zy = 0;
            isInSet = 1;

            for(i = 1; i <= IMAX; i += 1){
                zxtemp = zx * zx - zy * zy + cx;
                zy = 2 * zx * zy + cy;
                zx = zxtemp;
                if(zx * zx + zy * zy > 4){
                    isInSet = 0;
                    break;
                }
            }
            
            if (isInSet) {
				// (void) fwrite(black, 1, 3, fp);
				printf("*");
            } else {
				// fputc(white[0] * i / IMAX, fp);
				// fputc(white[1] * i / IMAX, fp);
				// fputc(white[2] * i / IMAX, fp);
				// fputc(white[3] * i / IMAX, fp);
				printf(" ");
			}
        }
		printf("\n");
    }

//   (void) fclose(fp);
  return 0;
}


	


