package ast

import (
	"fmt"
)

var (
	nextId int  = 0
	nextVarName = "'a"
	Integer = TypeOperator{"int",[]Type{}}
	Boolean = TypeOperator{"bool",[]Type{}}
	Float = TypeOperator{"float",[]Type{}}
	String = TypeOperator{"string",[]Type{}}
	Rune = TypeOperator{"rune",[]Type{}}
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
	Name string
	Types []Type
}

type Type interface {
	GetName() string
	GetType() string
}

type State map[string]Type

func (t TypeVariable) GetName() string {
	return t.Name
}

func (t TypeOperator) GetName() string {
	return t.Name
}

func (t TypeVariable) GetType() string {
	if t.Instance != nil {
		return t.Instance.GetType()
	}
	return t.Name
}

func (t TypeOperator) GetType() string {
	return t.Name
}

func (t Function) GetName() string {
	return t.Name
}

func (t Function) GetType() string {
	name := "("
	for i, el := range t.Types {
		if i > 0 {
			name += " -> "
		}
		name += el.GetName()
	}
	name += ")"
	return name
}

func NewTypeVariable() TypeVariable {
	t := TypeVariable{}
	t.Id = nextId
	nextId += 1
	t.Name = nextVarName
	nextVarName = "'" + string(rune(int(nextVarName[1])+1))
	return t
}

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
			switch s.(type) {
			case Comment:
				continue
			default:
				t, err := Infer(s, env, nonGeneric)
				if err != nil {
					return nil, err
				} else {
					fmt.Printf("Infer %s: %s %s\n", s.String(), t.GetName(), t.GetType())
				}
			}

		}
		return Unit, nil

	case BasicAst:
		switch node.(BasicAst).ValueType {
		case CHAR:
			return Rune, nil
		case INT:
			return Integer, nil
		case FLOAT:
			return Float, nil
		case BOOL:
			return Boolean, nil
		case STRING:
			return String, nil
		}

	case Operator:
		switch node.(Operator).ValueType {
		case CHAR:
			return Rune, nil
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
		if node.(Assignment).Left.(Identifier).StringValue != "_" {
			(*env)[node.(Assignment).Left.(Identifier).StringValue] = rightSide
		}
		return rightSide, nil

	case Identifier:
		if node.(Identifier).StringValue == "_" {
			return NewTypeVariable(), nil
		}
		return GetType(node.(Identifier).StringValue, *env, nonGeneric)

	case Call:
		if (*env)[node.(Call).Function.(Identifier).StringValue] != nil {
			types := (*env)[node.(Call).Function.(Identifier).StringValue].(Function).Types
			return types[len(types)-1], nil
		}
		return nil, InferenceError{"Do not know the type of function " + node.(Call).Function.(Identifier).StringValue}

	case If:
		return NewTypeVariable(), nil

	case Container:
		// TODO: Do we use this concretely anywhere?
		return List, nil

	case ArrayType:

		var lastType Type
		for _, s := range node.(ArrayType).Subvalues {
			t, err := Infer(s, env, nonGeneric)
			if err != nil {
				return nil, err
			}
			if lastType != nil {
				err := Unify(&t, &lastType, env)
				if err != nil {
					return nil, err
				}
			}
			lastType = t

			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			} else {
				fmt.Printf("Infer %s: %s\n", s.String(), lastType.GetName())
			}
		}
		return lastType, nil

	case Expr:
		var lastType Type
		for _, s := range node.(Expr).Subvalues {
			t, err := Infer(s, env, nonGeneric)
			if err != nil {
				return nil, err
			}
			if lastType != nil {
				err := Unify(&t, &lastType, env)
				if err != nil {
					return nil, err
				}
			}
			lastType = t

			if err != nil {
				return nil, err
			} else {
				fmt.Printf("Infer %s: %s\n", s.String(), lastType.GetName())
			}
		}
		return lastType, nil

	case Func:
		statements := node.(Func).Subvalues
		var lastType Type
		var newEnv = env

		// init argument names as type variables ready to be filled
		if len(node.(Func).Arguments) > 0 {
			for _, el := range node.(Func).Arguments {
				(*newEnv)[el.(Identifier).StringValue] = NewTypeVariable()
			}
		}

		// infer all statements
		for _, s := range statements {
			switch s.(type) {
			case Comment:
				continue
			default:
				t, err := Infer(s, newEnv, nonGeneric)
				lastType = t

				if err != nil {
					return nil, err
				} else {
					fmt.Printf("Infer %s: %s\n", s.String(), lastType.GetName())
				}
			}
		}

		// make our function type
		fType := Function{Name: node.(Func).Name, Types: []Type{}}

		// grab inferred types of args
		if len(node.(Func).Arguments) > 0 {
			for _, el := range node.(Func).Arguments {
				fType.Types = append(fType.Types, (*newEnv)[el.(Identifier).StringValue])
			}
		}

		// now the final type is the return type
		fType.Types = append(fType.Types, lastType)

		(*env)[node.(Func).Name] = fType
		return fType, nil

	default:
		panic("Don't know this type: " + node.String())
	}

	return nil, InferenceError{"Don't know this type: " + node.String()}
}

func GetType(name string, env State, nonGeneric []Type) (Type, error) {
	if _, ok := env[name]; ok {
		return Fresh(env[name].(Type), nonGeneric), nil
	}
	return nil, InferenceError{"Undefined symbol " + name}
}

func Unify(t1 *Type, t2 *Type, env *State) error {
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

	//fmt.Println("Unify", *t1, *t2)

	switch a.(type){
	case TypeVariable:
		if a.(TypeVariable).GetName() != b.GetName(){
			if OccursInType(a.(TypeVariable), b){
				return InferenceError{"Recursive unification"}
			}

			// TODO: need to be able to assign to `a` here
			newA := a.(TypeVariable)
			newA.Instance = b
			fmt.Printf("Unify %s is now %s\n", a.GetName(), b.GetName())
			(*env)[a.GetName()] = newA

			// try updating other refs to this typevariable
			for k, v := range *env {
				if v.GetName() == a.GetName(){
					(*env)[k] = b
				}
			}

		}
		return nil
	case TypeOperator:
		switch b.(type) {
		case TypeVariable:
			return Unify(&b, &a, env)
		case TypeOperator:
			aTypeLen := len(a.(TypeOperator).Types)
			bTypeLen := len(b.(TypeOperator).Types)
			if a.GetName() != b.GetName() || aTypeLen != bTypeLen{
				if len(b.(TypeOperator).Types) > 0 {
					return Unify(&b, &a, env)
				} else if aTypeLen > 0 {
					if a.(TypeOperator).Types[aTypeLen - 1].GetName() ==
						b.(TypeOperator).GetName() {
						return nil
					}
				}
				return InferenceError{fmt.Sprintf("Type mismatch: %s != %s", a.GetName(), b.GetName())}
			}
			// we know that the types must match because they didn't pass into that last condition
			for i, el := range a.(TypeOperator).Types {
				Unify(&el, &b.(TypeOperator).Types[i], env)
			}
			return nil
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


	switch t.(type){
	case TypeVariable:
		if t.(TypeVariable).Instance != nil {
			newInstance := Prune(t.(TypeVariable).Instance)
			return newInstance
		}
	}
	return t
}

func Fresh(t Type, nonGeneric []Type) Type {
	/*
	Makes a copy of a type expression.

	The type t is copied. Then the generic variables are duplicated and the
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
