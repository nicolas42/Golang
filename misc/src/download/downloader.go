// cd $(dirname $0)
// go run downloader.go

package download

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// var inputFile = "urls"
// var outputDir = "output"

// // MultipleFiles has a comment because the ide wanted one
// func MultipleFiles(urls []string) ([][]byte, error) {
// 	done := make(chan []byte, len(urls))
// 	errch := make(chan error, len(urls))
// 	for _, URL := range urls {
// 		go func(URL string) {
// 			b, err := FileAndWrite(URL)
// 			if err != nil {
// 				errch <- err
// 				done <- nil
// 				return
// 			}
// 			done <- b
// 			errch <- nil
// 		}(URL)
// 	}
// 	bytesArray := make([][]byte, 0)
// 	var errStr string
// 	for i := 0; i < len(urls); i++ {
// 		bytesArray = append(bytesArray, <-done)
// 		if err := <-errch; err != nil {
// 			errStr = errStr + " " + err.Error()
// 		}
// 	}
// 	var err error
// 	if errStr != "" {
// 		err = errors.New(errStr)
// 	}
// 	return bytesArray, err
// }

// FileAndWrite :)
func FileAndWrite(URL string) ([]byte, error) {
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
	//    fmt.Println(data)
	filename := filepath.Join(outputDir, filepath.Base(URL))
	err = ioutil.WriteFile(filename[0:min(100, len(filename))], data.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return data.Bytes(), nil
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

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
