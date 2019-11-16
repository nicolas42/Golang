package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {

	containsAll := true
	for i, _ := range searchTerms {
		if !strings.Contains(f.Name(), searchTerms[i]) {
			containsAll = false
			break
		}
	}
	if containsAll {
		fmt.Println(path)
	}

	return nil
}

var root = "."
var searchTerms []string

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("usage: %s [dir] [search terms]\r\n", os.Args[0])
	}

	root = os.Args[1]
	for i := 2; i < len(os.Args); i++ {
		searchTerms = append(searchTerms, os.Args[i])
	}

	filepath.Walk(root, visit)

}
