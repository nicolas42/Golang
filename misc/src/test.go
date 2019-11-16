package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func main() {

	// download.Test()

	// download.FileAndWrite("https://upload.wikimedia.org/wikipedia/commons/6/6c/Abhar-iran.JPG")

	url := "http://127.0.0.1:8080/test"
	response, err := http.Get(URL)
	defer response.Body.Close()
	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	return data.Bytes(), nil
	defer response.Body.Close()
	var data bytes.Buffer

}

// File downloads a file
func File(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}
