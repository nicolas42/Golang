// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "strings"

func main() {

	arg := "2*((1+2)^2 + (3-5)*2)"
	end := 1+strings.Index(arg, ")")
	start := strings.LastIndex(arg[:end], "(")
	expr := arg[start+1:end-1] // remove parens
	
	// Debugging
	fmt.Printf("Separate Brackets\n%v\n%v\n%v\n\n", arg[:start], arg[start:end], arg[end:])	
	op := strings.IndexAny(expr, "^*/+-")	
	fmt.Printf("Separate Ops: %v %v %v %v\n", expr, expr[:op], string(expr[op]), expr[op+1:]) // expr[op] returns uint	
	evalExpression("8*4+2/3", "*")

	
//	fmt.Println(start,end)
//	se := strings.SplitAfter(expr, "+") // split expression
//	fmt.Println(se)
//	result: do parse-math expression
//	arg: rejoin [ copy/part arg start result end ]

}

func evalExpression ( expr string, arg string ) {
	
	// find operation
	pos := strings.Index(expr, arg) // operation
	
	op := string(expr[pos]) // expr[i] returns uint8
	a, b := expr[:pos], expr[pos+1:]
	fmt.Println(a,op,b)
	start := strings.LastIndexAny(expr[:pos], "^*/+-")
	end := len(expr[:pos])+len(op)+strings.IndexAny(expr[pos+1:], "^*/+-")
	fmt.Println(start,end)
	
	if start != -1 { 
		a = expr[start+1:pos]
	}
	if end != -1 { 
		b = expr[pos+1:end]
	}
	
	fmt.Println(a,op,b)
}
	
	
	
	