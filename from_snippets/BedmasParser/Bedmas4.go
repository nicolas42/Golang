// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "strings"
import "strconv"
import "math"

type Entry struct {
	op rune // operation, rune is like char
	num float64 // number
}

// Pretty printing
func (e Entry) String() string {
	if e.op != 0 {
		return fmt.Sprintf("%v", string(e.op))
	}else{
		return fmt.Sprintf("%v", e.num)
	}
}

func main() {
	
	str := "2234*((1+2)^2 + (3-5)*2)"
	arg := []Entry{} // slice (resizeable array) of Entries
	p1, p2 := 0, 0 // position 1 and 2
	var result float64
	
	// Make Entries from the input string
	for ;; {

		// From p1, find the next operation character, set that position to p2
		p2 = strings.IndexAny(str[p1:], "*/+-^()" )
		if p2 == -1 { break }

		// If there's a number p1 and p2 then append it
		float, err := strconv.ParseFloat(strings.TrimSpace(str[p1:p1+p2]),64)
		if err == nil {
			e1 := Entry {num: float}
			arg = append(arg, e1)
		}
		// Append the operation
		e2 := Entry {op: rune(str[p1+p2])}
		arg = append(arg, e2)

		p1=p1+p2+1
	}
	
	// Print Entries
	fmt.Println(arg)
			
	for ;len(arg) > 1; {
	
		// Find deepest paren
		start := 0
		end := len(arg)
		
		for i, v := range arg {
			if v.op == ')' { end = i; break }
		}
		for i := end-1;i>=0; i-=1 {
			if arg[i].op == '(' { start = i; break }
		}
	
		// Find the position of the highest priority operation
		pos := 0
		
		for _, char := range "^*/+-" { 
			for i, v := range arg[start:end] {
				if v.op == rune(char) { pos = start+i; break }
			}
			if pos != 0 { break }
		}

		// Find result
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

		// Exit if there's only three entries at this stage
		if len(arg) <= 3 { break }
	

		// Otherwise reform argument
		newArg := []Entry{}
		newEntry := Entry{num: result}
		
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
		
		// Debugging
		fmt.Println(arg)
	}
	// debugging
	fmt.Println(result)
}