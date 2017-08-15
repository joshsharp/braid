package types

import (
    "braid/ast"
)

var (
	nextId int  = 0
	nextVarName = 'a'
	Integer = TypeOperator{"int",[]ast.Ast{}}
	Boolean = TypeOperator{"bool",[]ast.Ast{}}
	Float = TypeOperator{"int",[]ast.Ast{}}
	String = TypeOperator{"string",[]ast.Ast{}}
	Char = TypeOperator{"char",[]ast.Ast{}}
	List = TypeOperator{"list",[]ast.Ast{}}
	Unit = TypeOperator{"()",[]ast.Ast{}}

)

type TypeVariable struct {
	Name     string
	Id       int
	Instance ast.Ast // TODO: needs type
}

type TypeOperator struct {
	Name string
	Types []ast.Ast // TODO: needs type
}

type Function struct {
	TypeOperator

}

type Type interface {

}


func Infer(node ast.Ast, env map[string]interface{}) Type {

	switch node.(type){
	case ast.Module:
		return Unit
	case ast.BasicAst:
		switch node.(ast.BasicAst).Type {
		case ast.CHAR:
			return Char
		case ast.INT:
			return Integer
		case ast.FLOAT:
			return Float
		case ast.BOOL:
			return Boolean
		}
	case ast.Call:
		return TypeVariable{}
	case ast.If:
		return TypeVariable{}
	case ast.Container:
		return List
	case ast.Func:
		return TypeVariable{}
	default:
		panic("Don't know this type")
	}

	return nil
}
