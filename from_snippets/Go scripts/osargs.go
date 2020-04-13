package main

import (
    "fmt"
    "os"
	"strconv"
)

func main() {
	fmt.Println(len(os.Args), os.Args)

	a := []int{}
	for i:=1;i<len(os.Args);i++ {
		v,_ := strconv.Atoi(os.Args[i])
		a = append(a,v)
	}
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(a)

}


