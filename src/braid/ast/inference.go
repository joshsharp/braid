package ast

import (
	"fmt"
)

var (
	nextId int  = 0
	nextVarName = "a"
	Integer = TypeOperator{"int",[]Type{}}
	Boolean = TypeOperator{"bool",[]Type{}}
	Float = TypeOperator{"int",[]Type{}}
	String = TypeOperator{"string",[]Type{}}
	Char = TypeOperator{"char",[]Type{}}
	List = TypeOperator{"list",[]Type{}}
	Unit = TypeOperator{"()",[]Type{}}

)

type InferenceError struct {
	Message string
}

func (e InferenceError) Error() string {
	return e.Message
}

type TypeVariable struct {
	Name     string
	Id       int
	Instance Type
}

type TypeOperator struct {
	Name string
	Types []Type
}

type Function struct {
	TypeOperator
}

type Type interface {
	GetName() string
}

func (t TypeVariable) GetName() string {
	return t.Name
}

func (t TypeOperator) GetName() string {
	return t.Name
}

func (t Function) GetName() string {
	return t.Name
}

func NewTypeVariable() TypeVariable {
	t := TypeVariable{}
	t.Id = nextId
	nextId += 1
	t.Name = nextVarName
	nextVarName = string(rune(int(nextVarName[0])+1))
	return t
}

type State map[string]Type

func Infer(node Ast, env *State, nonGeneric []Type) (Type, error) {
	/*
	Computes the type of the expression given by node.

	The type of the node is computed in the context of the context of the
	supplied type environment env. Data types can be introduced into the
	language simply by having a predefined set of identifiers in the initial
	environment; this way there is no need to change the syntax or, more
	importantly, the type-checking program when extending the language.

	Args:
		node: The root of the abstract syntax tree.
		env: The type environment is a mapping of expression identifier names
			to type assignments.
			to type assignments.
		non_generic: A set of non-generic variables, or None

	Returns:
		The computed type of the expression.
	*/
	switch node.(type){

	case Module:
		statements := node.(Module).Subvalues
		for _, s := range statements {
			t, err := Infer(s, env, nonGeneric)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			} else {
				fmt.Printf("Infer %s: %s\n", s.Print(0), t.GetName())
			}
		}
		return Unit, nil
	case BasicAst:
		switch node.(BasicAst).ValueType {
		case CHAR:
			return Char, nil
		case INT:
			return Integer, nil
		case FLOAT:
			return Float, nil
		case BOOL:
			return Boolean, nil
		case STRING:
			return String, nil
		}
	case Comment:
		return Unit, nil
	case Assignment:
		rightSide, err := Infer(node.(Assignment).Right, env, nonGeneric)
		if err != nil {
			return nil, err
		}
		(*env)[node.(Assignment).Left.(BasicAst).StringValue] = rightSide
		return rightSide, nil

	case Call:
		return NewTypeVariable(), nil
	case If:
		return NewTypeVariable(), nil
	case Container:
		switch node.(Container).Type {
		case "Array":
			// TODO: Unify these types, must all be the same
			var lastType Type
			for _, s := range node.(Container).Subvalues {
				t, err := Infer(s, env, nonGeneric)
				lastType = t

				if err != nil {
					fmt.Println(err.Error())
					return nil, err
				} else {
					fmt.Printf("Infer %s: %s\n", s.Print(0), lastType.GetName())
				}
			}
			return lastType, nil

		case "CompoundExpr":
			// TODO: Unify these types, must all be the same
			var lastType Type
			for _, s := range node.(Container).Subvalues {
				t, err := Infer(s, env, nonGeneric)
				lastType = t

				if err != nil {
					fmt.Println(err.Error())
					return nil, err
				} else {
					fmt.Printf("Infer %s: %s\n", s.Print(0), lastType.GetName())
				}
			}
			return lastType, nil

		}
		return List, nil
	case Func:
		statements := node.(Func).Subvalues
		var lastType Type
		for _, s := range statements {
			t, err := Infer(s, env, nonGeneric)
			lastType = t

			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			} else {
				fmt.Printf("Infer %s: %s\n", s.Print(0), lastType.GetName())
			}
		}
		return lastType, nil
	default:
		panic("Don't know this type: " + node.Print(0))
	}

	return nil, InferenceError{"Don't know this type: " + node.Print(0)}
}


func GetType(name string, env State, nonGeneric []Type) (Type, error) {
	if _, ok := env[name]; ok {
		return Fresh(env[name].(Type), nonGeneric), nil
	}
	return nil, InferenceError{"Undefined symbol " + name}
}

func Unify(t1 *Type, t2 *Type) error {
	/* Unify the two types t1 and t2.

	Makes the types t1 and t2 the same.

	Args:
		t1: The first type to be made equivalent
		t2: The second type to be be equivalent

	Returns:
		An error if the types cannot be unified.
    */
	a := Prune(*t1)
	b := Prune(*t2)
	switch a.(type){
	case TypeVariable:
		if a.(TypeVariable).GetName() != b.GetName(){
			if OccursInType(a.(TypeVariable), b){
				return InferenceError{"Recursive unification"}
			}
			//a.(TypeVariable).Instance = b
		}
	case TypeOperator:
		switch b.(type) {
		case TypeVariable:
			return Unify(&b, &a)
		case TypeOperator:
			aTypeLen := len(a.(TypeOperator).Types)
			if a.GetName() != b.GetName() ||
				aTypeLen != len(b.(TypeOperator).Types){
				if len(b.(TypeOperator).Types) > 0 {
					return Unify(&b, &a)
				} else if a.(TypeOperator).Types[aTypeLen - 1].GetName() ==
					b.(TypeOperator).GetName(){
					return nil
				}
				return InferenceError{fmt.Sprintf("Type mismatch: %s != %s", a.GetName(), b.GetName())}
			}
			// we know that the types must match because they didn't pass into that last condition
			for i, el := range a.(TypeOperator).Types {
				Unify(&el, &b.(TypeOperator).Types[i])
			}
		}
	}
	return InferenceError{fmt.Sprintf("Types not unified: %s and %s", a.GetName(), b.GetName())}
}

func Prune(t Type) Type {
	/*
	Returns the currently defining instance of t.

	As a side effect, collapses the list of type instances. The function Prune
	is used whenever a type expression has to be inspected: it will always
	return a type expression which is either an uninstantiated type variable or
	a type operator; i.e. it will skip instantiated variables, and will
	actually prune them from expressions to remove long chains of instantiated
	variables.

	Args:
		t: The type to be pruned

	Returns:
		An uninstantiated TypeVariable or a TypeOperator
	*/
	return t
}

func Fresh(t Type, nonGeneric []Type) Type {
	/*
	Makes a copy of a type expression.

	The type t is copied. The the generic variables are duplicated and the
	non_generic variables are shared.

	Args:
		t: A type to be copied.
		non_generic: A set of non-generic TypeVariables
	*/

	return freshRec(t, nonGeneric, make(map[TypeVariable]TypeVariable))
}

func freshRec(tp Type, nonGeneric []Type, mappings map[TypeVariable]TypeVariable) Type {
	p := Prune(tp)
	switch p.(type){
	case TypeVariable:
		if IsGeneric(p.(TypeVariable), nonGeneric){
			if _, ok := mappings[p.(TypeVariable)]; !ok {
				mappings[p.(TypeVariable)] = NewTypeVariable()
			}
		} else {
			return p
		}
	case TypeOperator:
		freshTypes := make([]Type,0)
		for _, el := range tp.(TypeOperator).Types {
			freshTypes = append(freshTypes, freshRec(el, nonGeneric, mappings))
		}

		f := TypeOperator{p.GetName(), freshTypes}
		return f
	}
	return tp
}

func OccursInType(v TypeVariable, type2 Type) bool{
	/*
	Checks whether a type variable occurs in a type expression.

	Note: Must be called with v pre-pruned

	Args:
		v:  The TypeVariable to be tested for
		type2: The type in which to search

	Returns:
		True if v occurs in type2, otherwise False
 	*/
	prunedT2 := Prune(type2)

	if prunedT2.GetName() == v.GetName() {
		return true
	}
	switch prunedT2.(type){
	case TypeOperator:
		return OccursIn(v, prunedT2.(TypeOperator).Types)
	}

	return false
}

func OccursIn(v TypeVariable, types []Type) bool {
	/*
	Checks whether a types variable occurs in any other types.

	Args:
		t:  The TypeVariable to be tested for
		types: The sequence of types in which to search

	Returns:
		True if t occurs in any of types, otherwise False
	*/
	for _, el := range types {
		if OccursInType(v, el) {
			return true
		}
	}
	return false
}

func IsGeneric(v TypeVariable, nonGeneric []Type) bool {
	/*
	Checks whether a given variable occurs in a list of non-generic variables

	Note that a variables in such a list may be instantiated to a type term,
	in which case the variables contained in the type term are considered
	non-generic.

	Note: Must be called with v pre-pruned

	Args:
	v: The TypeVariable to be tested for genericity
	non_generic: A set of non-generic TypeVariables

	Returns:
	True if v is a generic variable, otherwise False
	*/

	return !OccursIn(v, nonGeneric)
}
