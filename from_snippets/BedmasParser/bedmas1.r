rebol []

main: [
	arg: "2*((1+2)^^2 + (3-5)*2)"
	probe arg

	while [ attempt [ 
		end: find/tail arg ")"
		start: find/reverse end "("
		expression: copy/part start end
	] ] [
		result: do parse-math expression
		arg: rejoin [ copy/part arg start result end ]
		probe arg 
		input
	]
	do parse-math arg
]

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

