package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	page := ReadUrl(os.Args[1])
	fmt.Println(page)
}

func ReadUrl(url string) string {
	// To Do: error handling
	fmt.Printf("Reading %v...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
