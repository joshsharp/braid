package main

import (
	"testing"
)

var examples = []string{`
# test
let cheesy = func item item2 {
	let b = 5.0 + 6.5
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

let add = func a b {
	a + b
}

let main = func {
	# one
	let a = 2
	# two
	let b = 3 + -2
	let c = a + b
	let d = [5, 6]
	let e = b
	let _ = add(4, 5)
	#Mod.f()
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

func TestExample0(t *testing.T){

	_, err := Compile(examples[0])
	if err != nil {
		t.Error(err.Error())
	}
	//println(result)

}

func TestExample1(t *testing.T){

	_, err := Compile(examples[1])
	if err != nil {
		t.Error(err.Error())
	}
	//println(result)

}

//func TestExample2(t *testing.T){
//
//	result, err := Compile(examples[2])
//	if err != nil {
//		t.Error(err.Error())
//	}
//	println(result)
//
//}