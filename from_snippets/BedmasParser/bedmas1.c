#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>


struct Entry {
	c  char
	d double
}

func main() {

	int debug = 0;

//	arg := "2234*((1+2)^2 + (3-5)*2)"
	arg := "1*(2+3)^4";

	double result = 0.0;
	result = evalBemdas(arg);
	printf("%s\n=%f\n", arg, result);


//	start = time.Now()
//	elapsed := time.Since(start)	
//	fmt.Printf("%s elapsed\n", elapsed)
//	fmt.Printf("%v function calls possible per second\n", int( time.Second / elapsed ) )
//	fmt.Println()
	
}

double evalBemdas(str string) {

	double result; result float64
	
	// === Make Entries from the input string ===

	entries = Entry[];
	int p1 = 0;
	int p2 = 0; // positions
	
	for (;;) {

		if (debug) { 
			printf("%f %c", entries[0].d, entries[0].c);
			// scanln();
		}

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
			entries = append(entries, e1)
		}
		// Append the operation
		e2 := Entry{op: rune(str[p2])}
		entries = append(entries, e2)

		p1 = p2 + 1
	}
	// Add the last number if it's there
	float, err := strconv.ParseFloat(strings.TrimSpace(str[p1:]), 64)
	if err == nil {
		e1 := Entry{num: float}
		entries = append(entries, e1)
	}

	// === Main loop ===
	for len(entries) > 1 {

		// 1. Find deepest parenthesis
		// by finding the first closing paren
		// then searching backwards for its opening paren

		start := 0
		end := len(entries)

		for i, v := range entries { 
			if v.op == ')' {
				end = i
				break
			}
		}
		for i := end - 1; i >= 0; i -= 1 { 
			if entries[i].op == '(' {
				start = i
				break
			}
		}

		// 2. Find the position of the highest priority operation
		// within the deepest paren
		op := 0

		for _, query := range "^*/+-" {
			for i, entry := range entries[start:end] {
				if entry.op == query {
					op = start + i
					break
				}
			}
			if op != 0 {
				break
			}
		}

		// 3. Calculate result
		a := entries[op-1].num
		b := entries[op+1].num

		switch entries[op].op {
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

		// Exit loop if there's only three entries left
		if len(entries) <= 3 {
			break
		}

		// 4. Otherwise reform entries, rinse, repeat
		entries2 := []Entry{}
		o1, o2 := -1, 2 // offsets that define span to replace

		// If paren only contains 3 elements then delete parens too
		if op-2 >= 0 && entries[op-2].op == '(' && entries[op+2].op == ')' {
			o1, o2 = -2, 3
		}
		entries2 = append(entries[:op+o1], Entry{num: result}) // insert result
		entries = append(entries2, entries[op+o2:]...)
		// slice tricks
		// http://stackoverflow.com/questions/21326109/why-are-lists-used-infrequently-in-go

	}
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
