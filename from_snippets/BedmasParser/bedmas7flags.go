// You can edit this code!
// Click here and start typing.
// Includes the shell option -debug
// go run bedmas7.go -debug
// more info here: https://gobyexample.com/command-line-flags

package main

import "fmt"
import "strings"
import "strconv"
import "math"
import "time"
import "flag"

type Entry struct {
	op  rune    // operation, rune is a utf8 symbol, like a char
	num float64 // number
}


var debug bool // print intermediate stages if true

func main() {

    boolPtr := flag.Bool("debug", false, "a bool")
	flag.Parse()

	debug = *boolPtr

	start := time.Now()
	evalBemdas("2234*((1+2)^2 + (3-5)*2)")
	elapsed := time.Since(start)
	fmt.Printf("%s elapsed\n", elapsed)
	fmt.Printf("%v function calls possible per second\n", int( time.Second / elapsed ) )
	
}

func evalBemdas(str string) float64 {

	// input
	// str := "2234*((1+2)^2 + (3-5)*2)"

	// Entries
	en := []Entry{}

	// positions
	p1, p2 := 0, 0

	var result float64

	// Make Entries from the input string
	for {

		// find the next operation character
		p2 = strings.IndexAny(str[p1:], "*/+-^()")
		if p2 == -1 {
			break
		}
		p2 = p1 + p2

		// If there's a number between p1 and p2 then append it
		float, err := strconv.ParseFloat(strings.TrimSpace(str[p1:p2]), 64)
		if err == nil {
			e1 := Entry{num: float}
			en = append(en, e1)
		}
		// Append the operation
		e2 := Entry{op: rune(str[p2])}
		en = append(en, e2)

		p1 = p2 + 1
	}
	// Add the last number if it's there
	float, err := strconv.ParseFloat(strings.TrimSpace(str[p1:]), 64)
	if err == nil {
		e1 := Entry{num: float}
		en = append(en, e1)
	}

	// Initial print
	if debug { 
			fmt.Println(en) 
			// fmt.Scanln() // wait for input between loops
	}

	for len(en) > 1 {

		// Find deepest paren
		start := 0
		end := len(en)

		// by finding first closing paren
		for i, v := range en { 
			if v.op == ')' {
				end = i
				break
			}
		}
		// then searching backwards for opening paren
		for i := end - 1; i >= 0; i -= 1 { 
			if en[i].op == '(' {
				start = i
				break
			}
		}

		// Find the position of the highest priority operation
		op := 0

		for _, query := range "^*/+-" {
			for i, v := range en[start:end] {
				if v.op == query {
					op = start + i
					break
				}
			}
			if op != 0 {
				break
			}
		}

		// Calculate result
		a := en[op-1].num
		b := en[op+1].num

		switch en[op].op {
		case '^':
			result = math.Pow(a, b)
		case '*':
			result = a * b
		case '/':
			result = a / b
		case '+':
			result = a + b
		case '-':
			result = a - b
		default:
			panic("Invalid operation?")
		}

		// Exit loop if there's only three entries
		if len(en) <= 3 {
			break
		}

		// Otherwise reform entries
		o1, o2 := -1, 2 // offsets for removal

		if op-2 >= 0 && en[op-2].op == '(' && en[op+2].op == ')' {
			o1, o2 = -2, 3 // delete surrounding parens too
			// indexes can't be negative
		}

		en2 := []Entry{}
		en2 = append(en2, en[:op+o1]...)
		en2 = append(en2, Entry{num: result}) // result entry
		en2 = append(en2, en[op+o2:]...)
		en = en2

		// Middle prints
		if debug { 
			fmt.Println(en) 
			// fmt.Scanln() // wait for input between loops
		}
	}
	// Final print
	if debug { fmt.Println(result) }
	return result
}

// Pretty printing
func (e Entry) String() string {
	if e.op != 0 {
		return fmt.Sprintf("%v", string(e.op))
	} else {
		return fmt.Sprintf("%v", e.num)
	}
}
