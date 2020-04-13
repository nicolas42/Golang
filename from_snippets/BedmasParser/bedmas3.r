rebol []

split-any: func [ a b /local o p1 p2 span ] [

	o: copy []
	b: charset b

	p1: p2: a
	while [ p2: find p1 b ] [
		; probe p1 probe p2 probe o
		span: copy/part p1 p2
		if not empty? span [ append o span ]
		append o p2/1
		p1: p2: next p2
	]
	if not empty? p1 [ append o p1 ]
	o
]

do-math: func [ arg ] [

	; arg: "2234*((1+2)^^2 + (3-5)*2)" ; debugging artefact
	
	; remove whitespace
	remove-each a arg [ find { ^-^/} a ] 
	arg: split-any arg "*/+-^^()e<>=" 

	while [ 1 <> length? arg ] [
		; range of interest. either inbetween parens or whole expression
		start: arg
		end: tail arg
	
		; find range of parentheses, if there are any
		if pos: find arg #")" [
			start: pos
			end: find/reverse pos #"("
		]
	
		; find the position 3 entry operation - num1 op num2
		; ^^ counts as a single ^ character since ^ is the rebol escape character
		foreach char "^^*/+-" [
			if pos: attempt [ find/part start char end ] [break]
		]
		pos: back pos
		
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
		probe reform arg
		input
	]
	return arg/1
]

do-math "2234*((1+2)^^2 + (3-5)*2)" 
