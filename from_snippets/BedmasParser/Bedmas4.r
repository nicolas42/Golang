rebol []

split-any: func [ str chars2 /local out p1 p2 chars ] [

	; Split a string at any of the given chars in the second argument
	; split-any "2234*((1+2)^^2 + (3-5)*2)" "*/+-^^()e<>="
	; == ["2234" #"*" #"(" #"(" "1" #"+" "2" #")" #"^^" "2 " #"+" " " #"(" "3" #"-" "5" #")" #"*" "2" #")"]

	p1: p2: str ; position 1 and 2
	out: copy []
	chars: charset chars2

	while [ p2: find p1 chars ] [

		; probe p1 probe p2 probe out input

		if p1 <> p2 [
			append out copy/part p1 p2
		]
		append out p2/1
		p1: p2: next p2

	]
	; append the rest
	if not empty? p1 [ append out p1 ]
	out
]



do-math: func [ arg  /debug /local start end pos result] [

	{ do-math "2234 * ( ( 1 + 2 ) ^^ 2 + ( 3 - 5 ) * 2 )"
	"2234 * ( 3 ^^ 2 + ( 3 - 5 ) * 2 )"
	"2234 * ( 3 ^^ 2 + -2 * 2 )"
	"2234 * ( 9.0 + -2 * 2 )"
	"2234 * ( 9.0 + -4 )"
	"2234 * 5.0"
	"11170.0"
	== 11170.0 }
	
	; remove whitespace
	remove-each a arg [ find { ^-^/} a ] 
	
	; split string into operations and numbers
	arg: split-any arg "*/+-^^()e<>=" 

	while [ 1 <> length? arg ] [

		; range of interest is either parens (parentheses) or the whole expression
		start: arg
		end: tail arg
	
		; set range to parens, if there are any
		if pos: find arg #")" [
			start: pos
			end: find/reverse pos #"("
		]
		; find deepest paren by finding the first closing paren then searching backwards 
		; to its opening paren
	
		; find the highest priority operation - num1 op num2
		; ^ is a rebol escape char so ^^ just means ^
		foreach char "^^*/+-" [
			if pos: attempt [ find/part start char end ] [break]
		]
		pos: back pos ; go to num1
		
		
		; evalute binomial operation
		either #"^^" = pos/2 [ 
			result: power do pos/1 do pos/3
		] [
			result: do reform copy/part pos 3
		]
	
		; if the parenthesis only contains three values then remove the parens as well
		either all [ 
			#"(" = pos/-1
			#")" = pos/4
		] [
			change/part back pos result 5
		] [
			change/part pos result 3
		]
		
		if debug [
			probe reform arg
			input
		]
	]
	return arg/1
]

; do-math/debug "2234*((1+2)^^2 + (3-5)*2)" 
do-math "2234 * ( ( 1 + 2 ) ^^ 2 + ( 3 - 5 ) * 2 )"
; took 269 microseconds
; 0:00:00.000269
; roughly 10X the time of the golang version
