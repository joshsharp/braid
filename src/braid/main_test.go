package main

import (
	"testing"
	"braid/ast"
)

var examples = []string{`
# test
let main = func {
	5 + 5
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

func TestEmptyModule(t *testing.T){
	m := ast.Module{Name:"Nothing",Subvalues:[]ast.Ast{}}
	env := make(ast.State)
	m2, err := ast.Infer(m, &env,[]ast.Type{})
	if err != nil {
		t.Error(err.Error())
	} else {
		if m2.GetInferredType().GetName() != ast.Unit.GetName() {
			t.Error(m2.GetInferredType())
		}
	}

}

func TestBasicFunc(t *testing.T){
	num := ast.BasicAst{ValueType:ast.INT, IntValue:5}
	//e := ast.Expr{Subvalues:[]ast.Ast{num}}
	f := ast.Func{Name:"Main", Subvalues:[]ast.Ast{num}}
	m := ast.Module{Name:"Nothing",Subvalues:[]ast.Ast{f}}
	env := make(ast.State)
	m2, err := ast.Infer(m, &env,[]ast.Type{})
	if err != nil {
		t.Error(err.Error())
	} else {
		if m2.GetInferredType().GetName() != ast.Unit.GetName() {
			t.Error(m2.GetInferredType())
		}
	}

}

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