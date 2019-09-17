
// math number to pixel number
// -2..2 to 0..1000
// add 2, multiply by 250
// pos = y * width * 4 + x * 4;
// data[pos] is red pixel
// use ctx.putImageData(imgData, 0, 0); afterwards

// multiplying complex
// (a+bi)(c+di) = ac +(bc+ad)i -bd
// = ac-bd + (ad+bc)i
// squaring
// (a+bi)(a+bi) = aa-bb + (ab+ab)i

// if number of iterations is low then we want it to be black
// then colorful
// then black again for the convergent region
// The convergent region is not actually written to

// === Colors ===
// var white = [255, 255, 255, 255];
// var black = [0, 0, 0, 255];
// var red = [255,0,0,255];
// var yellow = [255,240,120,255];
// var teal = [0,255,255,255];

// // logistic function
// function logistic(x,imax){
//     return 1 / (1 + Math.exp(-(x-(imax-6))))
// }

// Zn+1 = Zn^2 + c
// Z0 = 0
// c is pixel value

function drawMandelbrot(){   

    var black = [0, 0, 0, 255];
    var white = [255, 255, 255, 255];
    var x;
    var y;
    var width = 1000; // pixels
    var height = 1000;
    var canvas = document.querySelector("#canvas");
    canvas.height = height;
    canvas.width = width;

    // get the image from the context - image: {width, height, data}
    var ctx = canvas.getContext("2d");
    var imgData = ctx.getImageData(0, 0, width, height);
    var data = imgData.data;

    var xmin = -2;
    var xmax = 2;
    var ymin = -2;
    var ymax = 2;

    var xscale = width / (xmax - xmin);
    var yscale = height / (ymax - ymin);

    var zx;
    var zxtemp;
    var zy;
    var cx;
    var cy;

    var isInSet;
    var pos;
    var stride = 4; // 4 bytes per pixel
    var i;
    var imax = 40;

    // Make Black Image
    for (i = 0; i < data.length; i += stride) {
    data[i + 0] = black[0];
    data[i + 1] = black[1];
    data[i + 2] = black[2];
    data[i + 3] = black[3];
    }

    // Work out the other pixels
    for (y = 0; y < height; y += 1) {
        for (x = 0; x < width; x += 1) {

            // image is flipped vertically. numbers go up, pixels go down
            cx = x / xscale - xmax;
            cy = y / yscale - ymax;
            zx = 0;
            zy = 0;

            isInSet = true;
            for(i = 1; i <= imax; i += 1){
                zxtemp = zx * zx - zy * zy + cx;
                zy = 2 * zx * zy + cy;
                zx = zxtemp;
                if(zx * zx + zy * zy > 4){
                    isInSet = false;
                    break;
                }
            }
            
            if (isInSet === false) {
                pos = y * width * stride + x * stride;
                data[pos + 0] = white[0] * i / imax;
                data[pos + 1] = white[1] * i / imax;
                data[pos + 2] = white[2] * i / imax;
                // data[pos + 3] = white[3] * i / imax; // don't mess with opacity
            }
        }
    }

     ctx.putImageData(imgData, 0, 0); // Put data in canvas' 2D context

}