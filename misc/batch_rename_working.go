package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"./copy"
	//"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for i, f := range files {
		// fmt.Println(f.Name())

		oldName := f.Name()
		newName := fmt.Sprintf("image_%d.jpg", i)

		fmt.Println(oldName, newName)

		copy.Copy("test", "test.copy")

		/*
			    err := os.Rename(oldName, newName)
			    if err != nil {
				    log.Fatal(err)
			    }
		*/
	}
}
