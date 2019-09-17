package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"log"
	"net/http"

	"./mandelbrot"
)

func main() {

	port := ":8080"

	fmt.Printf("Serving mandelbrot on localhost%s...\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// https://stackoverflow.com/questions/26744814/serve-image-in-go-that-was-just-created

	img := mandelbrot.Draw(1000, 1000)

	var buffer bytes.Buffer
	png.Encode(&buffer, img)

	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, &buffer)
}
