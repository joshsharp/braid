package main

import (
	"braid/ast"
	"testing"
)

var examples = []string{`
module Main

// test
let main = {
	5 + 5
}
`, `
module Main

let add = (a, b) {
	a + b
}

let main = {
	// one
	let a = 2
	// two
	let b = 3 + -2
	let c = a + b
	let d = [5, 6]
	let e = b
	let _ = add(4, 5)
}
`,
	`
module Main

type Person = { name: string, age: int64 }

type Result ('a, 'b) =
	| OK 'a
	| Error 'b

type Option ('a) =
	| Some 'a
	| None

let test = (p: Person) -> () {
	// comment
	let a = p
	let b = OK("yes")
	let c = Error("failed to do thing")
	let d = Some("braid")
	let e = None()
	let f = a.name
	let g = 5 + 5
	// hi
}

let main = {
	// thing
	let a = 3
	let b = 45
	let c = 5 + b
	// another comment
	let d = [5, 6]
	let e = b
	test(Person{name: "no", age: -1})
}


`}

func TestEmptyModule(t *testing.T) {
	m := ast.Module{Name: "Nothing", Subvalues: []ast.Ast{}}
	env := ast.State{Env: make(map[string]ast.Type), UsedVariables: make(map[string]bool),
		Module: &m,
	}
	m2, err := m.Infer(&env, []ast.Type{})
	if err != nil {
		t.Error(err.Error())
	} else {
		if m2.GetInferredType().GetName() != ast.Unit.GetName() {
			t.Error(m2.GetInferredType())
		}
	}

}

func TestBasicFunc(t *testing.T) {
	num := ast.BasicAst{ValueType: ast.INT, IntValue: 5}
	//e := ast.Expr{Subvalues:[]ast.Ast{num}}
	f := ast.Func{Name: "Main", Subvalues: []ast.Ast{num}}
	m := ast.Module{Name: "Nothing", Subvalues: []ast.Ast{f}}
	env := ast.State{Env: make(map[string]ast.Type), UsedVariables: make(map[string]bool),
		Module: &m,
	}
	m2, err := m.Infer(&env, []ast.Type{})
	if err != nil {
		t.Error(err.Error())
	} else {
		if m2.GetInferredType().GetName() != ast.Unit.GetName() {
			t.Error(m2.GetInferredType())
		}
	}

}

func TestExample0(t *testing.T) {

	_, err := Compile(examples[0], false)
	if err != nil {
		t.Error(err.Error())
	}
	//println(result)

}

func TestExample1(t *testing.T) {

	_, err := Compile(examples[1], false)
	if err != nil {
		t.Error(err.Error())
	}
	//println(result)

}

func TestExample2(t *testing.T) {

	_, err := Compile(examples[2], false)
	if err != nil {
		t.Error(err.Error())
	}
	//println(result)

}
