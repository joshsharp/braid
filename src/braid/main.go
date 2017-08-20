package main

import (
	"fmt"
	"strings"
	"braid/ast"
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
			let c = 5 + 5
			# mm
			let a = Person{name:"Josh", age: 32}
			let b = OK("yes")
			let c = Error("failed to do thing")
			let d = Some("braid")
			let e = None()
			# hi
		}

		`}

	input := examples[1]

	lines := strings.Split(input, "\n")

	// first we make a reader from the input, which is a string
	r := strings.NewReader(input)
	// then we parse the input into ast
	result, err := ast.ParseReader("", r)

	if err != nil {

		fmt.Println("ERROR:")

		list := err.(ast.ErrorLister).Errors()
		for _, err := range list {
			// for each error, get the internal error
			pe := err.(ast.ParserError)
			printError(pe, lines)
		}
	} else {
		// print the input
		for i, el := range lines {
			fmt.Printf("%03d|%s\n", i+1, el)
		}

		// print the ast
		a := result.(ast.Ast)
		//fmt.Println("=", a.Print(0))

		env := make(ast.State)

		// infer types for the ast
		_, err := ast.Infer(a.(ast.Module), &env, nil)
		if err != nil {
			return
		}

		// print the compiled Go
		fmt.Println(a.Compile(env))
	}

}

func printError(pe ast.ParserError, lines []string) {

	// determine how many past lines to render
	start := pe.Pos()[0] - 1
	if pe.Pos()[0] >= 5 {
		start = pe.Pos()[0] - 5
	}

	// print those past lines up until the line of the error
	for i, el := range lines[start:pe.Pos()[0]] {
		offset := start
		fmt.Printf("%03d|%s\n", i+1+offset, el)
		//i += 1
	}

	// print the caret pointing to the position
	line := lines[pe.Pos()[0]-1]
	fmt.Printf("    ")
	for _, el := range line[:pe.Pos()[1]-1] {
		if el == '\t' {
			fmt.Printf("----")
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Printf("^\n\n")

	// print the actual error
	fmt.Printf("Line %d, character %d: ", pe.Pos()[0], pe.Pos()[1])
	fmt.Println(pe.InnerError())
}
