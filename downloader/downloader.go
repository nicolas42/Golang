package main

import (
//    "fmt"
    "net/http"
    "errors"
    "io"
    "bytes"
	"io/ioutil"
	"log"
	"path/filepath"

    // readlines
    "bufio"
    "os"
)

const inputFile = "urls.txt"
const outputDir = "output"

func main(){

    // Make images dir
    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
        os.Mkdir(outputDir, 0777)
    }

    // Read the urls
    urls, err := readLines(inputFile)
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

    // Download the files
    println("Downloading...")
    downloadMultipleFiles(urls)
}

func downloadMultipleFiles(urls []string) ([][]byte, error) {
	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))
	for _, URL := range urls {
		go func(URL string) {
			b, err := downloadFileAndWrite(URL)
			if err != nil {
				errch <- err
				done <- nil
				return
			}
			done <- b
			errch <- nil
		}(URL)
	}
	bytesArray := make([][]byte, 0)
	var errStr string
	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}
	var err error
	if errStr!=""{
		err = errors.New(errStr)
	}
	return bytesArray, err
}

func downloadFileAndWrite(URL string) ([]byte, error) {
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
	file := filepath.Join(outputDir, filepath.Base(URL))
	err = ioutil.WriteFile(file[:min(100, len(file))], data.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
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


