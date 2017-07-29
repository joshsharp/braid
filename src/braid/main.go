package main

import (
	"fmt"
	"strings"
)

func main() {
	examples := []string{`
# test
let cheesy = func item item2 {
	let _, b = 5.0 + 6.5
	let c = [5, 6, 7]
    # more test
    item ++ " and " ++ item2 ++ " with cheese"
    
}
let tester = func a {
	let nothing = a + 1
	if a > 100 {
		a + 1
	} else if a > 5 {
		a + 50
	} else {
		a + 100
	}
	
}
let main = func {
	let something = func {
		4 + 9
	}
	let result = 5 * (4 + 6) * 2
	let yumPizza = cheesy("pineapple", ("bbq" ++ "sauce"))
	# hoo boy this is a good'un
	let five = 1 / 1 + 3 * (55 - 2)
	# let mmm = 1 + 1
}
`, `
let adder = func a b {
    # whoop
    let b = 4 + 5
    # hi
	Mod.f(4, 5)
    # yes
	let b = f()

}

let main = func {
	# one
	let a = 2
	# two
	let b = 3 + -2
	let _ = List.add(1, 2, [3])
	adder(4, 5)
	Mod.f()
}
`,
`
type Person = { name: string, age: int }

type IntList = list int

type Result 'a 'b =
	| OK 'a
	| Error 'b

type Option 'a =
	| Some 'a
	| None

let main = func {
	# thing
	let a = 3
	let b = 45
	let c = 5
	# no
	let d = [5, 6]
	let e = b
	test((5 + 6), Person{name: "no", age: -1})
}

let test = func p {
	let a = Person{name:"Josh", age: 32}
	let b = OK("yes")
	let c = Error("failed to do thing")
	let d = Some("braid")
	let e = None()
}

`}

	input := examples[1]

	lines := strings.Split(input, "\n")
	
	
	//fmt.Println(input)
	r := strings.NewReader(input)
	result, err := ParseReader("", r) //FailureTracking(true)

	if err != nil {
		
		fmt.Println("ERROR:")
		list := err.(errList)
		for _, err := range list {
			
			pe := err.(*parserError)

			
			//for (i < pe.pos.line){
			for i, el := range(lines[pe.pos.line-1:pe.pos.line]){
				offset := pe.pos.line-1
				fmt.Printf("%03d|%s\n", i + 1 + offset, el)
				//i += 1
			}
			
			
			line := lines[pe.pos.line-1]
			fmt.Printf("    ")
			for _, el := range(line[:pe.pos.col-1]){
				if el == '\t'{
					fmt.Printf("----")
				} else {
					fmt.Printf("-")
				}
			}
			fmt.Printf("^\n\n")
			fmt.Printf("Line %d, character %d: ", pe.pos.line, pe.pos.col)
			fmt.Println(pe.Inner)
		}
	} else {
		for i, el := range(lines){
			fmt.Printf("%03d|%s\n", i + 1, el)
		}
		
		ast := result.(Ast)
		fmt.Println("=", ast.Print(0))

        types := infer(ast.(Module))

		fmt.Println(ast.Compile(types))
	}

}
