package main

import (
    "log"
    "net/http"
)

func main() {
    println("Serving home 2TB on 55555...")
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/home/nicolas"))))
    if err := http.ListenAndServe(":55555", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


