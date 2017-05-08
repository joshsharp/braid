package main

import (
	"fmt"
	"strings"
)

func main(){
    input := `
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
# let thing = cheesy "pineapple" "bbq sauce";
let result = 5 * (4 + 6) * 2;
let five = 1 / 1 + 3 * (55 - 2);
`
	
    fmt.Println(input)
    r := strings.NewReader(input)
    result, err := ParseReader("", r)
    ast := result.(Ast)
    fmt.Println("=", ast.Print(0))
    fmt.Println(err)
}