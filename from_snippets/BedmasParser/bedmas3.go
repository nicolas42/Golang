// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "strings"
import "strconv"
import "math"

// Next Step

// Example of how to use []interface{}
// http://stackoverflow.com/questions/20170275/how-to-find-a-type-of-a-object-in-golang

// Example of type switch
// https://play.golang.org/p/fym_m4DqL3 

// fmt.Scanln(&var)
// http://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line


type Entry struct {
	op rune // operation, rune is like char
	num float64 // number
}
// if op is 0 then entry is a number

// Pretty printer
func (e Entry) String() string {
	if e.op != 0 {
		return fmt.Sprintf("%v", string(e.op))
	}else{
		return fmt.Sprintf("%v", e.num)
	}
}

func main() {
	
	str := "2234*((1+2)^2 + (3-5)*2)"
	chars := "*/+-^()e<>="
	arg := []Entry{}
	p1, p2 := 0, 0
	var result float64

	// debugging
	debug := ""
	// test, _ := strconv.ParseFloat(strings.TrimSpace("2 "),64)
	// fmt.Println(test)
	
	
	// Separate numbers and operations
	for ;; {
		p2 = strings.IndexAny(str[p1:], chars )
		if p2 == -1 { break }
		num := strings.TrimSpace(str[p1:p1+p2])
		float, err := strconv.ParseFloat(num,64)
		if err == nil {
			e1 := Entry {num: float}
			arg = append(arg, e1)
		}
		e2 := Entry {op: rune(str[p1+p2])}
		arg = append(arg, e2)
		p1=p1+p2+1
		// fmt.Println(arg)
		// fmt.Scanln(&debug)
	}
	fmt.Println(arg)	
	
	for ;len(arg) > 1; {
	
		// Find deepest paren if it exists, define possible range of evaluation
		start, end := 0, len(arg)

		// fmt.Println(arg, arg[end-1],arg[0])

		for i, v := range arg {
			if v.op == ')' { end = i; break }
		}

		for i := end-1;i>=0; i-=1 {
			if arg[i].op == '(' { start = i; break }
		}
	
		// fmt.Println(arg[start:end])
	
		// Find the position of the highest priority operation
		pos := 0
		for _, char := range "^*/+-" { 
			//fmt.Println(arg[start:end]) 
			for i, v := range arg[start:end] {
				//fmt.Println(arg[start+i])
				if v.op == rune(char) { pos = start+i; break }
			}
			if pos != 0 { break }
		}
		//fmt.Println(arg[pos-1:pos+2])


		// Evaluate result of operation
		a := arg[pos-1].num
		b := arg[pos+1].num
		switch arg[pos].op {
		case '^': 
			result = math.Pow(a,b)
		case '*': 
			result = a*b
		case '/': 
			result = a/b
		case '+': 
			result = a+b
		case '-': 
			result = a-b
		default:
			panic("argh...")
		}
	
		// fmt.Println(result)
		if len(arg) == 3 { break }
	
		newArg := []Entry{}
		newEntry := Entry{num: result}
	
		// fmt.Println(arg[pos-2].op =='(', arg[pos+2])
	
		// Reform argument
		if arg[pos-2].op =='(' && arg[pos+2].op == ')' {
	
			newArg = append(newArg, arg[:pos-2]...)
			newArg = append(newArg, newEntry)
			newArg = append(newArg, arg[pos+3:]...)
		} else {
			newArg = append(newArg, arg[:pos-1]...)
			newArg = append(newArg, newEntry)
			newArg = append(newArg, arg[pos+2:]...)

		}
		arg = newArg
	
		// debug
		fmt.Print(arg)
		fmt.Scanln(&debug)


	}
	fmt.Println(result)
}