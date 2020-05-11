// Usage:
// finder some words here
// case insensitive

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

//	"github.com/atotto/clipboard"
)

func visit(path string, f os.FileInfo, err error) error {

	// print(".")
	containsAll := true
	for i := range searchTerms {
		if !strings.Contains(strings.ToLower(f.Name()), strings.ToLower(searchTerms[i])) {
			containsAll = false
			break
		}
	}
	if containsAll {
		fmt.Println(path)
		// fmt.Println(f.Name())		
//		clipboard.WriteAll(path)
	}

	return nil
}

var root = "."
var searchTerms []string

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [search terms]\r\n", os.Args[0])
		return
	}

	root,_ = os.Getwd(); // os.Args[1]

	for i := 1; i < len(os.Args); i++ {
		searchTerms = append(searchTerms, os.Args[i])
	}

	filepath.Walk(root, visit)

}

