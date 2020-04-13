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

func main() {
	Bemdas("2234*((1+2)^2 + (3-5)*2)")
}

func Bemdas( str string ) float64 {
	
	// input
	// str := "2234*((1+2)^2 + (3-5)*2)"
	
	// Entries
	en := []Entry{}
	
	// positions
	p1, p2 := 0, 0
	
	var result float64
	
	// Make Entries from the input string
	for ;; {

		// find the next operation character
		p2 = strings.IndexAny(str[p1:], "*/+-^()" )
		if p2 == -1 { break }
		p2 = p1+p2
		
		// If there's a number between p1 and p2 then add it
		float, err := strconv.ParseFloat(strings.TrimSpace(str[p1:p2]),64)
		if err == nil {
			e1 := Entry {num: float}
			en = append(en, e1)
		}
		// Append the operation
		e2 := Entry {op: rune(str[p2])}
		en = append(en, e2)

		p1=p2+1
	}
	// Add the last number if it's there
	float, err := strconv.ParseFloat(strings.TrimSpace(str[p1:]),64)
	if err == nil {
		e1 := Entry {num: float}
		en = append(en, e1)
	}
	fmt.Println(e1)
	// Initial print
	fmt.Println(en)
			
	for ;len(en) > 1; {
	
		// Find deepest paren
		start := 0
		end := len(en)
		
		for i, v := range en { // find first closing paren
			if v.op == ')' { end = i; break }
		}
		for i := end-1;i>=0; i-=1 { // then search backwards from it for opening paren
			if en[i].op == '(' { start = i; break }
		}
	
		// Find the position of the highest priority operation
		pos := 0
		
		for _, char := range "^*/+-" { 
			for i, v := range en[start:end] {
				if v.op == rune(char) { pos = start+i; break }
			}
			if pos != 0 { break }
		}

		// Find result
		a := en[pos-1].num
		b := en[pos+1].num
		
		switch en[pos].op {
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
			panic("Invalid operation?")
		}

		// Exit loop if there's only three entries
		if len(en) <= 3 { break }

		// Otherwise replace entries
		o1, o2 := -1, 2 // offsets for removal (from operation)
		
		if pos-2>=0 && en[pos-2].op =='(' && en[pos+2].op == ')' {	
			o1, o2 = -2, 3 // delete surrounding parens too
			// indexes can't be negative
		} 
		
		// Reform entries
		en2 := []Entry{}
		en2 = append(en2, en[:pos+o1]... ) 
		en2 = append(en2, Entry{num: result} ) // result entry
		en2 = append(en2, en[pos+o2:]... )
		en = en2
		
		// Middle prints
		fmt.Println(en)
		
		fmt.Scanln()
	}
	// Final print
	fmt.Println(result)
	return result
}


// Pretty printing
func (e Entry) String() string {
	if e.op != 0 {
		return fmt.Sprintf("%v", string(e.op))
	}else{
		return fmt.Sprintf("%v", e.num)
	}
}