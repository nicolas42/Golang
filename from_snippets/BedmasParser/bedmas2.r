rebol []

split-any: func [ a b /local o p1 p2 ] [

	o: copy []
	b: charset b
	p1: a

	until [
		p2: find p1 b
		append o copy/part p1 p2
		append o p2/1
		p1: next p2

		any [ tail? p1 tail? p2 ]
	]
	o
]

a: split-any "2234*((1+2)^^2 + (3-5)*2)" "*/+-^^()e<>=" 


op: charset "*/+-^^()e<>=" 
nop: complement op 

arg: "2*((1+2)^^2 + (3-5)*2)"
arg: collect [ foreach a arg [ keep a ] ] 

while [
	find arg charset ")^^*/+-"
	foreach char ")^^*/+-" [
		pos: attempt [ find arg char ]
		if found? pos [ break/return true ]
	]
	false
][ 
	either pos/1 = #")" [
		end: next pos
		start: find/reverse pos #"("
	] [
		end: find next pos charset "()^^*/+-"
		start: find/reverse pos charset "()^^*/+-"
	]
	expression: copy/part start end
	result: do parse-math rejoin expression
	arg: attempt [ compose [ (copy/part arg start) (result) (end) ] ] ; attempt for when result is a single number
	probe rejoin arg
]


while [ attempt [
end: find/tail arg #")"
start: find/reverse end #"("
]][
expression: copy/part start end
result: do parse-math rejoin expression
arg: compose [ (copy/part arg start) (result) (end) ]
probe arg
input
]
probe do parse-math rejoin arg



parse-math: funct [
	{Turn a math expression like "4+4*4/(a+4)" into a rebol block}
	str
] [
	op: charset "*/+-^^()e<>=" 
	nop: complement op 
	out: copy [] 
	a: none 
	parse str [
		some [
			copy a some nop (append out a) 
			| copy a some op (append out a)
		] 
		end
	] 
	out: form out 
	foreach [a b] [
		"--" "+" 
		"+-" "-" 
		"*-" " * - " 
		"/-" " / - " 
		" e " "e" 
		" e- " "e-" 
		"^^" "** "
	] [
		replace/all out a b
	] 
	load/all out
]

do main

