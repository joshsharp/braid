package main

import (
	"fmt"
	"strings"
)

func main(){
    examples := []string { `
# test
let _, b = 5.0 + 6.5;
let cheesy = func item item2 {
    item ++ " and " ++ item2 ++ " with cheese"; # more test
}
let tester = func a {
    let result = if a > 100 {
        a + 1;
    } else if a > 50 {
        a + 20;
    } else {
        a + 2;
    }
    result;
}
let result = 5 * (4 + 6) * 2;
let yumPizza = cheesy "pineapple" "bbq sauce";
# hoo boy this is a good'un
let five = 1 / 1 + 3 * (55 - 2);
# let mmm = 1 + 1
`,
`
# one
let a = 2;
# two
let b = 2 + 2;
let _ = List.add 1 2 3;
`,
`
let 😐 = 45;
let a = 3;
`};
	
	input := examples[1]
	
    fmt.Println(input)
    r := strings.NewReader(input)
    result, err := ParseReader("", r)
    
	if err != nil {
		fmt.Println("ERROR:")
		list := err.(errList)
		for _, err := range list {
			fmt.Println(err)
		}
	} else {
		ast := result.(Ast)
		fmt.Println("=", ast.Print(0))
	}
 
}