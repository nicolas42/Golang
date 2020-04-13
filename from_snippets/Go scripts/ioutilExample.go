package main

import (
  "io/ioutil"
  "fmt"
)

func main(){

	data,_ := ioutil.ReadFile("test.txt")
	fmt.Println(string(data))
}

