package main

import (
    "log"
    "net/http"
)

func main() {
    println("Serving home 2TB on 8080...")
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/home/nicolas"))))
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
