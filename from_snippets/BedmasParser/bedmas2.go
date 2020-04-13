// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "strings"
import "strconv"
import "math"


func main() {
	
	str := "2234*((1+2)^2 + (3-5)*2)"
	chars := "*/+-^^()e<>="
	arg := []string{}
	p1, p2 := 0, 0
	
	// Separate numbers and operations
	for ;; {
		p2 = strings.IndexAny(str[p1:], chars )
		if p2 == -1 { break }
		arg = append(arg, string(str[p1:p1+p2]), string(str[p1+p2]))
		p1=p1+p2+1
	}
	fmt.Println(arg)	
	
	// Find deepest paren if it exists, define possible range of evaluation
	start, end := 0, len(arg)
	for i, v := range arg {
		if v == ")" { end = i; break }
	}	
	for i := end;; i-=1 {
		if string(arg[i]) == "(" { start = i; break }
		
	}
	fmt.Println(arg[start:end])
	
	// Find the highest priority operation
	pos := []string{}
	for _, char := range "^*/+-" { 
		// fmt.Println(string(char)) 
		for i, v := range arg[start:end] {
			if v == string(char) { pos = arg[start+i-1:]; break }
		}	
	}
	fmt.Println(pos[0], pos[1], pos[2])
	// Evaluate result of operation


	a, _ := strconv.ParseFloat(pos[0],64)
	b, _ := strconv.ParseFloat(pos[2],64)
	result := 0.0
	switch rune(pos[1][0]) {
	case '^': math.Pow(a,b)
	case '*': 
		result = a*b
	case '/': 
		result = a/b
	case '+': 
		result = a+b
	case '-': 
		result = a-b
	}
	fmt.Println(result)
	
}