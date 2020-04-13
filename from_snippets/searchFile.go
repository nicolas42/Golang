// https://xojoc.pw/justcode/golang-file-tree-traversal.html
// To Do:
// Make cli program that takes two arguments
// e.g. searchFile /Volumes/Elements2TB INN371

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type filename string

func main() {

	out := []filename{}

	rootDir := "/Users/nick"
	query := ".pdf"

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, query) && info.Size() > 1e6 {
			out = append(out, filename(path)) // fmt.Println(path)
		}
		//		if info.IsDir() {			fmt.Println(path)		}
		if err != nil {
			fmt.Println(err)
		}
		return err
	})

	fmt.Println(out)
}

func (fn filename) String() string {
	_, file := filepath.Split(string(fn))
	return file
}
n