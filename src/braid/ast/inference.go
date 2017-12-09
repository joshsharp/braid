package ast

import (
	"fmt"
	"strings"
)

var (
	nextId         int = 0
	nextVarName        = "a"
	nextTempId     int = 0
	Boolean            = TypeOperator{"bool", []Type{}}
	Integer            = TypeOperator{"int64", []Type{}}
	Float              = TypeOperator{"float64", []Type{}}
	Number             = TypeOperator{"number", []Type{Float, Integer}}
	String             = TypeOperator{"string", []Type{}}
	Rune               = TypeOperator{"rune", []Type{}}
	Unit               = TypeOperator{"()", []Type{}}
	MainReturnType     = TypeOperator{" ", []Type{}}
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
	Name  string
	Types []Type
}

type Function struct {
	Name  string
	External string
	Types []Type
	Env   State
}

type Type interface {
	GetName() string
	GetType() string
}

type State struct {
	Env map[string]Type
	UsedVariables map[string]bool
}

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
	t.Name = "'" + nextVarName
	nextVarName = string(rune(int(nextVarName[0]) + 1))
	return t
}

func NewTempVariable() string {

	nextId += 1
	name := fmt.Sprintf("__temp_%d", nextTempId)
	nextTempId += 1
	return name
}

func Infer(node Ast, env *State, nonGeneric []Type) (Ast, error) {
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

	switch node.(type) {

	case Module:
		node := node.(Module)
		statements := node.Subvalues

		var newStatements []Ast

		for _, s := range statements {
			switch s.(type) {
			case Comment:
				continue
			default:
				//fmt.Printf("Encountered %s\n", s.String())
				t, err := Infer(s, env, nonGeneric)
				if err != nil {
					return nil, err
				} else {
					newStatements = append(newStatements, t)
					//fmt.Printf("Module infer %s: %s\n", s.String(), t.GetInferredType())
				}
			}
		}
		node.Subvalues = newStatements
		//for _, s := range node.Subvalues {
		//
		//	fmt.Printf("Module infer %s: %s\n", s.String(), s.GetInferredType())
		//}

		//node.InferredType = Unit
		return node, nil
	case Extern:
		node := node.(Extern)

		// make our function type
		fType := Function{Name: node.Name, External:"__go_" + node.Import, Types: []Type{}}

		// grab inferred types of args
		if len(node.Arguments) > 0 {
			for _, el := range node.Arguments {
				// lookup the type from its annotation
				//fmt.Println("looking up arg type annotation:", el.(Identifier).Annotation)
				anno, err := GetTypeFromAnnotation(el.(Identifier).Annotation)
				if err != nil {
					return nil, err
				}
				fType.Types = append(fType.Types, anno)
			}
		}

		// lookup return type annotation
		//fmt.Println("looking up return type annotation:", node.ReturnAnnotation)
		anno, err := GetTypeFromAnnotation(node.ReturnAnnotation)
		if err != nil {
			return nil, err
		}
		fType.Types = append(fType.Types, anno)
		(*env).Env[node.Name] = fType
		node.InferredType = fType

		return node, nil

	case BasicAst:
		node := node.(BasicAst)
		switch node.ValueType {
		case CHAR:
			node.InferredType = Rune
			return node, nil
		case INT:
			node.InferredType = Integer
			return node, nil
		case FLOAT:
			node.InferredType = Float
			return node, nil
		case BOOL:
			node.InferredType = Boolean
			return node, nil
		case STRING:
			node.InferredType = String
			return node, nil
		}

	case Operator:
		node := node.(Operator)
		switch node.ValueType {
		case NUMBER:
			node.InferredType = Number
			return node, nil
		case CHAR:
			node.InferredType = Rune
			return node, nil
		case INT:
			node.InferredType = Integer
			return node, nil
		case FLOAT:
			node.InferredType = Float
			return node, nil
		case BOOL:
			node.InferredType = Boolean
			return node, nil
		case STRING:
			node.InferredType = String
			return node, nil
		}

	case Comment:
		node := node.(Comment)
		//node.InferredType = Unit
		return node, nil

	case Assignment:
		node := node.(Assignment)
		right := node.Right
		//fmt.Printf("Encountered %s\n", right.String())
		rightSide, err := Infer(right, env, nonGeneric)
		if err != nil {
			return nil, err
		}

		///fmt.Println("right side type:", rightSide.GetInferredType())
		if node.Left.(Identifier).StringValue != "_" {
			name := node.Left.(Identifier).StringValue

			if _, ok := (*env).Env[name]; ok {
				return nil, InferenceError{"Cannot assign to " + name + ", it is already declared"}
			}

			(*env).Env[name] = rightSide.GetInferredType()

		}

		left := Identifier{StringValue: node.Left.(Identifier).StringValue,
			InferredType: rightSide.GetInferredType()}

		// check in case this is a typevar already stored
		if t, ok := (*env).Env[rightSide.GetInferredType().GetName()]; ok {
			(*env).Env[left.StringValue] = t
			left.InferredType = t
		}

		node.Right = rightSide
		node.Left = left
		node.InferredType = left.GetInferredType()

		return node, nil

	case Identifier:
		node := node.(Identifier)
		if node.StringValue == "_" {
			node.InferredType = NewTypeVariable()
			return node, nil
		}
		t, err := GetType(node.StringValue, *env, nonGeneric)
		if err != nil {
			return nil, err
		}
		(*env).UsedVariables[node.StringValue] = true
		node.InferredType = t
		return node, nil

	case Call:
		node := node.(Call)
		fType := (*env).Env[node.Function.StringValue]
		if fType == nil {
			return nil, InferenceError{"Do not know the type of function " + node.Function.StringValue}
		}

		// TODO: unify call args and func args
		// infer call args so they get marked as used, and eventually unify with func defn args
		for _, el := range(node.Arguments){
			_, err := Infer(el, env, nonGeneric)
			if err != nil {
				return nil, err
			}
		}

		types := fType.(Function).Types
		node.InferredType = types[len(types)-1]
		(*env).UsedVariables[node.Function.StringValue] = true

		if fType.(Function).External != "" {
			node.Module = Identifier{StringValue: strings.SplitN(fType.(Function).External,".",2)[0]}
			node.Function = Identifier{StringValue: strings.SplitN(fType.(Function).External,".",2)[1]}
		}
		//fmt.Println((*env).UsedVariables)
		return node, nil

	case BinOp:
		//fmt.Printf("Encountered %s\n", node.String())
		node := node.(BinOp)
		operator := node.Operator
		op, err := Infer(operator, env, nonGeneric)
		//fmt.Printf("Encountered %s\n", op.String())
		if err != nil {
			return nil, err
		}

		l := node.Left

		left, err := Infer(l, env, nonGeneric)

		if err != nil {
			return nil, err
		}

		r := node.Right
		right, err := Infer(r, env, nonGeneric)
		if err != nil {
			return nil, err
		}

		node.Left = left
		node.Right = right
		node.Operator = op
		lType := left.GetInferredType()
		rType := right.GetInferredType()
		err = Unify(&lType, &rType, env)
		if err != nil {
			return nil, err
		}

		// number is a convenience for how go handles ops with floats and ints
		// if we see it, be more specific by using the left type
		if op.GetInferredType().GetName() == Number.GetName() {
			node.InferredType = lType
			// that they've been unified
		} else {
			node.InferredType = op.GetInferredType()
		}

		return node, nil

	case Container:
		node := node.(Container)

		return node, nil

	case If:
		node := node.(If)

		node.TempVar = NewTempVariable()
		(*env).UsedVariables[node.TempVar] = true
		//fmt.Println("new temp var", node.TempVar)

		ifAst := node.Condition
		condition, err := Infer(ifAst, env, nonGeneric)
		if err != nil {
			return nil, err
		}

		if condition.GetInferredType().GetType() != Boolean.GetType() {
			return nil, InferenceError{"Condition must be a boolean"}
		}

		statements := node.Then
		var thenType Type
		var elseType Type

		// infer all statements
		newStatements := []Ast{}

		for i, s := range statements {
			switch s.(type) {
			case BasicAst, Expr, BinOp:
				t, err := Infer(s, env, nonGeneric)
				thenType = t.GetInferredType()

				if err != nil {
					return nil, err
				} else {

					if i == len(statements)-1 {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}

						newStatements = append(newStatements, assign)
					} else {
						newStatements = append(newStatements, t)
					}
					//fmt.Printf("Infer Then: %s\n", thenType.GetName())
				}

			case Assignment:
				t, err := Infer(s, env, nonGeneric)
				thenType = t.GetInferredType()

				if err != nil {
					return nil, err
				} else {
					newStatements = append(newStatements, t)
					if i == len(statements)-1 {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t.(Assignment).Left, Update: true}
						newStatements = append(newStatements, assign)
					}
					//fmt.Printf("Infer Then: %s\n", thenType.GetName())
				}
			default:

				t, err := Infer(s, env, nonGeneric)
				thenType = t.GetInferredType()

				if err != nil {
					return nil, err
				} else {
					newStatements = append(newStatements, t)
					if i == len(statements)-1 {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}

						newStatements = append(newStatements, assign)
					}
					//fmt.Printf("Infer Then: %s\n", thenType.GetName())
				}
			}
		}

		node.Then = newStatements

		statements = node.Else
		newStatements = []Ast{}

		// infer all statements
		for i, s := range statements {
			switch s.(type) {
			case BasicAst, BinOp, Expr:
				t, err := Infer(s, env, nonGeneric)
				thenType = t.GetInferredType()

				if err != nil {
					return nil, err
				} else {

					if i == len(statements)-1 {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}
						newStatements = append(newStatements, assign)
					} else {
						newStatements = append(newStatements, t)
					}
					//fmt.Printf("Infer Then: %s\n", thenType.GetName())
				}
			case Assignment:
				t, err := Infer(s, env, nonGeneric)
				thenType = t.GetInferredType()

				if err != nil {
					return nil, err
				} else {
					newStatements = append(newStatements, t)
					if i == len(statements)-1 {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t.(Assignment).Left, Update: true}
						newStatements = append(newStatements, assign)
					}
					//fmt.Printf("Infer Then: %s\n", thenType.GetName())
				}
			default:

				t, err := Infer(s, env, nonGeneric)
				thenType = t.GetInferredType()

				if err != nil {
					return nil, err
				} else {
					newStatements = append(newStatements, t)
					if i == len(statements)-1 {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}
						newStatements = append(newStatements, assign)
					}
					//fmt.Printf("Infer Then: %s\n", thenType.GetName())
				}
			}
		}

		node.Else = newStatements

		if elseType != nil {
			err = Unify(&thenType, &elseType, env)
			if err != nil {
				return nil, err
			}
			node.InferredType = elseType

		} else {
			node.InferredType = thenType
		}

		return node, nil

	case ArrayType:
		node := node.(ArrayType)
		var lastType Type
		var newValues []Ast

		for _, s := range node.Subvalues {
			t, err := Infer(s, env, nonGeneric)
			tType := t.GetInferredType()

			if err != nil {
				return nil, err
			}
			if lastType != nil {

				err := Unify(&tType, &lastType, env)
				if err != nil {
					return nil, err
				}
			}
			lastType = tType
			newValues = append(newValues, t)

			if err != nil {
				//fmt.Println(err.Error())
				return nil, err
			} else {
				//fmt.Printf("Infer %s: %s\n", s.String(), lastType.GetName())
			}
		}
		node.Subvalues = newValues
		node.InferredType = lastType
		//fmt.Println("Array type is", lastType.GetName())
		return node, nil

	case Expr:
		node := node.(Expr)
		var lastType Type
		newValues := []Ast{}

		for _, s := range node.Subvalues {
			//fmt.Printf("Encountered %s\n", s.String())
			t, err := Infer(s, env, nonGeneric)
			tType := t.GetInferredType()

			if err != nil {
				return nil, err
			}
			if lastType != nil {

				err := Unify(&tType, &lastType, env)
				if err != nil {
					return nil, err
				}
			}
			lastType = tType
			newValues = append(newValues, t)

			//fmt.Printf("Infer %s: %s\n", s.String(), lastType.GetName())

		}
		node.Subvalues = newValues
		// why are we updating the type from the env here?
		if t, ok := (*env).Env[lastType.GetName()]; ok {
			lastType = t
		}
		node.InferredType = lastType
		return node, nil

	case Func:
		node := node.(Func)
		statements := node.Subvalues
		var newStatements []Ast
		var lastType Type
		var newEnv = State{Env:make(map[string]Type), UsedVariables:make(map[string]bool)}
		CopyState(*env, newEnv)

		// init
		// argument names as type variables ready to be filled
		if len(node.Arguments) > 0 {
			for _, el := range node.Arguments {
				newEnv.Env[el.(Identifier).StringValue] = NewTypeVariable()
			}
		}

		if _, ok := (*env).Env[node.Name]; ok {
			return nil, InferenceError{"Cannot redeclare func " + node.Name + ", it is already declared"}
		}

		// infer all statements
		for i, s := range statements {
			switch s.(type) {
			case If:
				t, err := Infer(s, &newEnv, nonGeneric)

				if err != nil {
					return nil, err
				} else {
					lastType = t.GetInferredType()
					newStatements = append(newStatements, t)
					// if last, replace with its equivalent return
					if i == len(statements)-1 && node.Name != "main" {
						returnAst := Return{Value: Identifier{StringValue: t.(If).TempVar}}
						newEnv.UsedVariables[t.(If).TempVar] = true
						newStatements = append(newStatements, returnAst)
					}
				}
			case Expr:
				t, err := Infer(s, &newEnv, nonGeneric)

				if err != nil {
					return nil, err
				} else {
					lastType = t.GetInferredType()
					// if last, replace with its equivalent return
					if i == len(statements)-1 && node.Name != "main" {
						returnAst := Return{Value: t}
						newStatements = append(newStatements, returnAst)
					} else {
						newStatements = append(newStatements, t)
					}
				}
			case BinOp:
				t, err := Infer(s, &newEnv, nonGeneric)

				if err != nil {
					return nil, err
				} else {
					lastType = t.GetInferredType()

					// if last, replace with its equivalent return
					if i == len(statements)-1 && node.Name != "main" {
						returnAst := Return{Value: t}
						newStatements = append(newStatements, returnAst)
					} else {
						newStatements = append(newStatements, t)
					}
				}
			case Assignment:
				t, err := Infer(s, &newEnv, nonGeneric)

				if err != nil {
					return nil, err
				} else {
					lastType = t.GetInferredType()
					newStatements = append(newStatements, t)
					// if last, replace with its equivalent return
					if i == len(statements)-1 && node.Name != "main" {

						returnAst := Return{Value: t.(Assignment).Left}
						newEnv.UsedVariables[t.(Assignment).Left.(Identifier).StringValue] = true
						//fmt.Printf("Adding %s as used", t.(Assignment).Left.(Identifier).StringValue)
						newStatements = append(newStatements, returnAst)
					}
				}
			default:
				//fmt.Printf("Encountered %s\n", s.String())
				t, err := Infer(s, &newEnv, nonGeneric)

				if err != nil {
					return nil, err
				} else {
					lastType = t.GetInferredType()

					newStatements = append(newStatements, t)
					//fmt.Printf("Func infer %s: %s\n", s.String(), lastType)

					if i == len(statements)-1 && node.Name != "main" {
						returnAst := Return{Value: t}
						newStatements = append(newStatements, returnAst)
					}
				}
			}
		}

		if node.Name != "main" {
			node.Subvalues = newStatements

			// make our function type
			fType := Function{Name: node.Name, Types: []Type{}}

			// grab inferred types of args
			if len(node.Arguments) > 0 {
				for _, el := range node.Arguments {
					fType.Types = append(fType.Types, newEnv.Env[el.(Identifier).StringValue])
				}
			}

			// now the final type is the return type
			fType.Types = append(fType.Types, lastType)
			DiffState(*env, newEnv)
			fType.Env = newEnv

			(*env).Env[node.Name] = fType
			node.InferredType = fType

		} else {
			// omit return
			node.Subvalues = newStatements //[:len(newStatements)-1]

			// make our function type
			fType := Function{Name: node.Name, Types: []Type{}}

			// grab inferred types of args
			if len(node.Arguments) > 0 {
				for _, el := range node.Arguments {
					fType.Types = append(fType.Types, newEnv.Env[el.(Identifier).StringValue])
				}
			}

			fType.Types = []Type{}
			DiffState(*env, newEnv)
			fType.Env = newEnv

			(*env).Env[node.Name] = fType
			node.InferredType = fType
		}
		return node, nil

	default:
		panic("Don't know this type: " + node.String())
	}

	return nil, InferenceError{"Don't know this type: " + node.String()}
}

func GetTypeFromAnnotation(name string)(Type, error){
	types := make(map[string]Type)
	types["int64"] = Integer
	types["float64"] = Float
	types["string"] = String
	types["rune"] = Rune
	types["bool"] = Boolean
	types["()"] = Unit

	if val, ok := types[name]; ok {
		return val, nil
	}
	return nil, InferenceError{fmt.Sprintf("Do not know annotated type '%s'", name)}
}

func GetType(name string, env State, nonGeneric []Type) (Type, error) {
	if _, ok := env.Env[name]; ok {
		return Fresh(env.Env[name].(Type), nonGeneric), nil
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

	switch a.(type) {
	case TypeVariable:
		if a.(TypeVariable).GetName() != b.GetName() {
			if OccursInType(a.(TypeVariable), b) {
				return InferenceError{"Recursive unification"}
			}

			newA := a.(TypeVariable)
			newA.Instance = b
			//fmt.Printf("Unify %s is now %s\n", a.GetName(), b.GetName())
			(*env).Env[a.GetName()] = newA

			// try updating other refs to this typevariable
			for k, v := range (*env).Env {
				if v.GetName() == a.GetName() {
					(*env).Env[k] = b
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
			if a.GetName() != b.GetName() || aTypeLen != bTypeLen {
				if len(b.(TypeOperator).Types) > 0 {
					return Unify(&b, &a, env)
				} else if aTypeLen > 0 {
					if a.(TypeOperator).Types[aTypeLen-1].GetName() ==
						b.(TypeOperator).GetName() {
						return nil
					}
				}
				return InferenceError{fmt.Sprintf("Type mismatch: %s != %s", a.GetName(), b.GetName())}
			}
			// we know that the types must match because they didn't pass into that last condition
			for i, el := range a.(TypeOperator).Types {
				err := Unify(&el, &b.(TypeOperator).Types[i], env)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return InferenceError{fmt.Sprintf("Types not unified: %s and %s", a, b)}
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

	switch t.(type) {
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
	switch p.(type) {
	case TypeVariable:
		if IsGeneric(p.(TypeVariable), nonGeneric) {
			if _, ok := mappings[p.(TypeVariable)]; !ok {
				mappings[p.(TypeVariable)] = NewTypeVariable()
			}
		} else {
			return p
		}
	case TypeOperator:
		freshTypes := make([]Type, 0)
		for _, el := range tp.(TypeOperator).Types {
			freshTypes = append(freshTypes, freshRec(el, nonGeneric, mappings))
		}

		f := TypeOperator{p.GetName(), freshTypes}
		return f
	}
	return tp
}

func OccursInType(v TypeVariable, type2 Type) bool {
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
	switch prunedT2.(type) {
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

func CopyState(existing State, copy State) {
	for k, v := range existing.Env {
		copy.Env[k] = v
	}
	for k, v := range existing.UsedVariables {
		copy.UsedVariables[k] = v
	}
}

func DiffState(existing State, copy State) {
	for k, _ := range existing.Env {
		delete(copy.Env, k)
	}
	for k, _ := range existing.UsedVariables {
		delete(copy.UsedVariables, k)
	}
}
