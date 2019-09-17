package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", webHandler)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/image", imageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yafoo I love images"))
}

// https://stackoverflow.com/questions/26744814/serve-image-in-go-that-was-just-created
func webHandler(w http.ResponseWriter, r *http.Request) {

	img, err := os.Open("image.png")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)
}
