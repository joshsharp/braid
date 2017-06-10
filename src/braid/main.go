package main

import (
	"fmt"
	"strings"
)

func main() {
	examples := []string{`
# test
let cheesy = func item item2 {
	let _, b = 5.0 + 6.5;
    item ++ " and " ++ item2 ++ " with cheese"; # more test
}
let tester = func a {
    
    
    if a > 100 {
        a + 1;
    } else if a > 50 {
        a + 20;
    } else {
        a + 2;
    }
    let result = a;
    a;
}
let main = func {
	let something = func {
		4 + 9;
	}
	let result = 5 * (4 + 6) * 2;
	let yumPizza = cheesy "pineapple" "bbq sauce";
	# hoo boy this is a good'un
	let five = 1 / 1 + 3 * (55 - 2);
	# let mmm = 1 + 1
}
`,
		`
let adder = func a b {
	a + b;
}
let main = func {
	# one
	let a = 2;
	# two
	let b = 3 + -2;
	let _ = List.add 1 2 3;
	let _ = adder 4 5;
}
`,
		`
let main = func {
	let a = 3;
	let b = 45;
	# thing
}
`}

	input := examples[0]

	fmt.Println(input)
	r := strings.NewReader(input)
	result, err := ParseReader("", r) //FailureTracking(true)

	if err != nil {
		fmt.Println("ERROR:")
		list := err.(errList)
		for _, err := range list {

			pe := err.(*parserError)
			fmt.Println(pe)
		}
	} else {
		ast := result.(Ast)
		fmt.Println("=", ast.Print(0))

		s := make(map[string]interface{})

		fmt.Println(ast.Compile(s))
	}

}
