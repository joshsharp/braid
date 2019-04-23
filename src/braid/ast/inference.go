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
	Byte               = TypeOperator{"byte", []Type{}}
	Rune               = TypeOperator{"rune", []Type{}}
	Unit               = TypeOperator{"()", []Type{}}
	MainReturnType     = TypeOperator{" ", []Type{}}
	Any                = TypeOperator{"interface{}", []Type{}}
)

type Type interface {
	GetName() string
	GetType() string
}

type State struct {
	Env           map[string]Type
	UsedVariables map[string]bool
	Module        *Module
}

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
	Name     string
	External string
	Types    []Type
	Env      State
}

type List struct {
	Name  string
	Types []Type
}

type Record struct {
	Name   string
	Params []string
	Fields map[string]Type
}

type VariantInstanceType struct {
	Name        string
	Constructor string
	Types       []Type
}

type VariantType struct {
	Name         string
	Params       []string
	Constructors []string
}

type VariantConstructorType struct {
	Name   string
	Parent string
	Types  []Type
}

func (v VariantInstanceType) GetName() string {
	return v.Name
}

func (v VariantInstanceType) GetType() string {
	return v.Name
}

func (r Record) GetName() string {
	return r.Name
}

func (r Record) GetType() string {
	return r.Name
}

func (l List) GetName() string {
	return l.Name
}

func (l List) GetType() string {
	return l.Name
}

func (v VariantType) GetName() string {
	return v.Name
}

func (v VariantType) GetType() string {
	return v.Name
}

func (v VariantConstructorType) GetName() string {
	return v.Name
}

func (v VariantConstructorType) GetType() string {
	return v.Parent + "." + v.Name
}

func (f Function) String() string {
	str := "func " + f.Name + ":\n"
	for _, el := range f.Types {
		str += fmt.Sprintf("\t%s\n", el)
	}
	return str
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
	name := "func("
	for i, el := range t.Types[:len(t.Types)-1] {
		if i > 0 {
			name += ", "
		}
		name += el.GetType()
	}
	name += ") "
	name += t.Types[len(t.Types)-1].GetType()
	return name
}

func NewTypeVariable() Type {
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

func (node Module) Infer(env *State, nonGeneric []Type) (Ast, error) {
	//node := node.(Module)
	statements := node.Subvalues

	var newStatements []Ast

	for _, s := range statements {
		switch s.(type) {
		case Comment:
			continue
		default:
			//fmt.Printf("Encountered %s\n", s.String())
			t, err := s.Infer(env, nonGeneric)
			if err != nil {
				return nil, err
			} else {
				if t != nil {
					newStatements = append(newStatements, t)
				}
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
	env.Module = &node
	return node, nil
}

func (node BasicAst) Infer(env *State, nonGeneric []Type) (Ast, error) {

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
	case NIL:
		node.InferredType = Unit
		return node, nil
	}
	panic("Unknown type")
}

func (node Operator) Infer(env *State, nonGeneric []Type) (Ast, error) {
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
	panic("Unknown type")
}

func (node Comment) Infer(env *State, nonGeneric []Type) (Ast, error) {
	return node, nil
}

func (node Assignment) Infer(env *State, nonGeneric []Type) (Ast, error) {
	right := node.Right
	//fmt.Printf("Encountered %s\n", right.String())
	rightSide, err := right.Infer(env, nonGeneric)
	if err != nil {
		return nil, err
	}

	if rightSide.GetInferredType().GetType() == Unit.GetType() {
		return nil, InferenceError{"Cannot assign a value of nil"}
	}

	switch node.Left.(type) {
	case Identifier:
		//fmt.Println("right side type:", rightSide.GetInferredType())
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
	case Container:
		left := node.Left.(Container)
		// multiple assignables
		switch rightSide.GetInferredType().(type) {
		case List:

			if len(left.Subvalues) != len(rightSide.GetInferredType().(List).Types) {
				return nil, InferenceError{"Number of identifiers on left side does not much number of return arguments"}
			}

			for i, el := range left.Subvalues {
				name := el.(Identifier).StringValue

				if _, ok := (*env).Env[name]; ok {
					return nil, InferenceError{"Cannot assign to " + name + ", it is already declared"}
				}

				(*env).Env[name] = rightSide.GetInferredType().(List).Types[i]
			}
		default:
			return nil, InferenceError{"Cannot unpack multiple values for " + right.String()}
		}

		node.Right = rightSide
		node.Left = left
		node.InferredType = right.GetInferredType()

	}

	return node, nil

}

func (node Identifier) Infer(env *State, nonGeneric []Type) (Ast, error) {
	if node.StringValue == "_" {
		node.InferredType = NewTypeVariable()
		return node, nil
	}

	if node.StringValue[0] == '\'' {
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
}

func (node Call) Infer(env *State, nonGeneric []Type) (Ast, error) {
	fType := (*env).Env[node.Function.StringValue]
	if fType == nil {
		return nil, InferenceError{"Do not know the type of function " + node.Function.StringValue}
	}

	// remove one type for return type
	if len(node.Arguments) != len(fType.(Function).Types)-1 {
		return nil, InferenceError{fmt.Sprintf("Called function %s with %d argument%s, but it takes %d",
			node.Function, len(node.Arguments),
			func() string {
				if len(node.Arguments) == 1 {
					return ""
				}
				return "s"
			}(),
			len(fType.(Function).Types)-1)}
	}
	// infer call args so they get marked as used, and eventually unify with func defn args
	for i, el := range node.Arguments {

		fArg := fType.(Function).Types[i]

		// infer arg
		el, err := el.Infer(env, nonGeneric)
		if err != nil {
			return nil, err
		}

		t := el.GetInferredType()

		// now unify matching func arg type
		err = Unify(t, fArg, env)
		if err != nil {
			return nil, err
		}
		node.Arguments[i] = el

	}

	types := fType.(Function).Types
	node.InferredType = types[len(types)-1]
	(*env).UsedVariables[node.Function.StringValue] = true

	// if external func, rewrite to original func name and module
	if fType.(Function).External != "" {
		if strings.Contains(fType.(Function).External, ".") {
			node.Module = Identifier{StringValue: strings.SplitN(fType.(Function).External, ".", 2)[0]}
			node.Function = Identifier{StringValue: strings.SplitN(fType.(Function).External, ".", 2)[1]}
		} else {
			node.Module = Identifier{StringValue: ""}
			node.Function = Identifier{StringValue: fType.(Function).External}
			fmt.Printf("Setting %s to %s directly\n", node.Function.StringValue, fType.(Function).External)
		}

	}
	//fmt.Println((*env).UsedVariables)
	return node, nil
}

func (node BinOp) Infer(env *State, nonGeneric []Type) (Ast, error) {
	//fmt.Printf("Encountered %s\n", node.String())

	operator := node.Operator
	op, err := operator.Infer(env, nonGeneric)
	//fmt.Printf("Encountered %s\n", op.String())
	if err != nil {
		return nil, err
	}

	l := node.Left

	left, err := l.Infer(env, nonGeneric)

	if err != nil {
		return nil, err
	}

	r := node.Right
	right, err := r.Infer(env, nonGeneric)
	if err != nil {
		return nil, err
	}

	node.Left = left
	node.Right = right
	node.Operator = op
	lType := left.GetInferredType()
	rType := right.GetInferredType()
	err = Unify(lType, rType, env)
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
		Unify(lType, node.InferredType, env)
		Unify(rType, node.InferredType, env)
	}

	return node, nil

}

func (node Container) Infer(env *State, nonGeneric []Type) (Ast, error) {
	return node, nil
}

func (node ArrayType) Infer(env *State, nonGeneric []Type) (Ast, error) {
	return node, nil
}

func (node ArrayAccess) Infer(env *State, nonGeneric []Type) (Ast, error) {
	// look up type of identifier, should be array
	t, err := GetType(node.Identifier.StringValue, *env, nonGeneric)
	if err != nil {
		return nil, err
	}

	switch t.(type) {
	case List:
		//pass
	default:
		return nil, InferenceError{"Cannot access the index of a non-array value"}
	}

	// get subtype - this is our type
	node.InferredType = t.(List).Types[0]

	// make sure index is an int
	index, err := node.Index.Infer(env, nonGeneric)
	if err != nil {
		return nil, err
	}
	node.Index = index
	indexType := index.GetInferredType()

	err = Unify(indexType, Integer, env)

	if err != nil {
		return nil, err
	}
	(*env).UsedVariables[node.Identifier.StringValue] = true
	return node, nil
}

func (node If) Infer(env *State, nonGeneric []Type) (Ast, error) {

	node.TempVar = NewTempVariable()
	(*env).UsedVariables[node.TempVar] = true
	//fmt.Println("new temp var", node.TempVar)

	ifAst := node.Condition
	condition, err := ifAst.Infer(env, nonGeneric)
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
			t, err := s.Infer(env, nonGeneric)
			thenType = t.GetInferredType()

			if err != nil {
				return nil, err
			} else {

				if i == len(statements)-1 {
					if thenType.GetName() != Unit.GetName() {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}
						newStatements = append(newStatements, assign)
					} else {
						newStatements = append(newStatements, t)
					}

				} else {
					newStatements = append(newStatements, t)
				}
				//fmt.Printf("Infer Then: %s\n", thenType.GetName())
			}

		case Assignment:
			t, err := s.Infer(env, nonGeneric)
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

			t, err := s.Infer(env, nonGeneric)
			thenType = t.GetInferredType()

			if err != nil {
				return nil, err
			} else {
				newStatements = append(newStatements, t)
				if i == len(statements)-1 {
					if thenType.GetName() != Unit.GetName() {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}

						newStatements = append(newStatements, assign)
					}
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
			t, err := s.Infer(env, nonGeneric)
			thenType = t.GetInferredType()

			if err != nil {
				return nil, err
			} else {

				if i == len(statements)-1 {
					if thenType.GetName() != Unit.GetName() {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}
						newStatements = append(newStatements, assign)
					} else {
						newStatements = append(newStatements, t)
					}
				} else {
					newStatements = append(newStatements, t)
				}
				//fmt.Printf("Infer Then: %s\n", thenType.GetName())
			}
		case Assignment:
			t, err := s.Infer(env, nonGeneric)
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

			t, err := s.Infer(env, nonGeneric)
			thenType = t.GetInferredType()

			if err != nil {
				return nil, err
			} else {
				newStatements = append(newStatements, t)
				if i == len(statements)-1 {
					if thenType.GetName() != Unit.GetName() {
						assign := Assignment{Left: Identifier{StringValue: node.TempVar},
							Right: t, Update: true}
						newStatements = append(newStatements, assign)
					}
				}
				//fmt.Printf("Infer Then: %s\n", thenType.GetName())
			}
		}
	}

	node.Else = newStatements

	if elseType != nil {
		err = Unify(thenType, elseType, env)
		if err != nil {
			return nil, err
		}
		node.InferredType = elseType

	} else {
		node.InferredType = thenType
	}

	return node, nil

}

func (node Array) Infer(env *State, nonGeneric []Type) (Ast, error) {

	var lastType Type
	var newValues []Ast

	for _, s := range node.Subvalues {
		t, err := s.Infer(env, nonGeneric)
		tType := t.GetInferredType()

		if err != nil {
			return nil, err
		}
		if lastType != nil {

			err := Unify(tType, lastType, env)
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
	node.InferredType = List{Name: fmt.Sprintf("[]%s", lastType.GetName()), Types: []Type{lastType}}
	//fmt.Println("Array type is", lastType.GetName())
	return node, nil

}

func (node Expr) Infer(env *State, nonGeneric []Type) (Ast, error) {

	var lastType Type
	newValues := []Ast{}

	for _, s := range node.Subvalues {
		//fmt.Printf("Encountered %s\n", s.String())
		t, err := s.Infer(env, nonGeneric)

		if err != nil {
			return nil, err
		}
		tType := t.GetInferredType()
		if lastType != nil {

			err := Unify(tType, lastType, env)
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

}

func (node Func) Infer(env *State, nonGeneric []Type) (Ast, error) {
	statements := node.Subvalues
	var newStatements []Ast
	var lastType Type
	var returnAnnotationType Type
	var newEnv = State{Env: make(map[string]Type), UsedVariables: make(map[string]bool)}
	CopyState(*env, newEnv)

	// init
	// argument names as type variables ready to be filled
	if len(node.Arguments) > 0 {
		for _, el := range node.Arguments {
			newType := NewTypeVariable()
			newEnv.Env[el.(Identifier).StringValue] = newType
			if el.(Identifier).Annotation != nil {
				t, err := GetTypeFromAnnotation(el.(Identifier).Annotation, env)
				if err != nil {
					return nil, InferenceError{fmt.Sprintf("Cannot understand type annotation %s: %s",
						el.(Identifier).StringValue,
						el.(Identifier).Annotation)}
				}

				Unify(t, newType, &newEnv)
				newEnv.Env[el.(Identifier).StringValue] = t
			}
		}
	}

	if node.ReturnAnnotation != nil {
		t, err := GetTypeFromAnnotation(node.ReturnAnnotation, env)
		if err != nil {
			return nil, InferenceError{fmt.Sprintf("Cannot understand type annotation %s: %s",
				node.ReturnAnnotation.(Identifier).StringValue,
				node.ReturnAnnotation.(Identifier).Annotation)}
		}
		returnAnnotationType = t
	}

	if _, ok := (*env).Env[node.Name]; ok {
		return nil, InferenceError{"Cannot redeclare func " + node.Name + ", it is already declared"}
	}

	// infer all statements
	for i, s := range statements {
		switch s.(type) {
		case If:
			t, err := s.Infer(&newEnv, nonGeneric)

			if err != nil {
				return nil, err
			} else {
				lastType = t.GetInferredType()

				newStatements = append(newStatements, t)
				// if last, replace with its equivalent return
				if i == len(statements)-1 && node.Name != "main" {
					if lastType.GetName() != Unit.GetName() {
						returnAst := Return{Value: Identifier{StringValue: t.(If).TempVar}}
						newEnv.UsedVariables[t.(If).TempVar] = true
						newStatements = append(newStatements, returnAst)
					}
				}

			}
		case Expr:
			t, err := s.Infer(&newEnv, nonGeneric)

			if err != nil {
				return nil, err
			} else {
				lastType = t.GetInferredType()
				// if last, replace with its equivalent return
				if i == len(statements)-1 && lastType.GetName() != Unit.GetName() && node.Name != "main" {
					returnAst := Return{Value: t}
					newStatements = append(newStatements, returnAst)

				} else {
					newStatements = append(newStatements, t)
				}
			}
		case BinOp:
			t, err := s.Infer(&newEnv, nonGeneric)

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
			t, err := s.Infer(&newEnv, nonGeneric)

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
			t, err := s.Infer(&newEnv, nonGeneric)

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

	if returnAnnotationType != nil {
		err := Unify(returnAnnotationType, lastType, &newEnv)
		if err != nil {
			return nil, InferenceError{fmt.Sprintf("Type annotation does not match return type: %s != %s",
				node.ReturnAnnotation.String(), lastType.GetName())}
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
}

func (node RecordField) Infer(env *State, nonGeneric []Type) (Ast, error) {

	var t Type
	var err error
	switch node.Type.(type) {
	case BasicAst:
		t, err = GetTypeFromAnnotation(node.Type.(BasicAst), env)
	case Identifier:
		t, err = GetTypeFromAnnotation(node.Type.(Identifier), env)
	case Container:
		t, err = GetTypeFromAnnotation(node.Type.(Container), env)
	default:
		panic(fmt.Sprintf("Do not know this type %s", node.Type))
	}

	if err != nil {
		return nil, err
	}
	node.InferredType = t
	return node, nil

}

func (node RecordType) Infer(env *State, nonGeneric []Type) (Ast, error) {

	for i, el := range node.Fields {
		field, err := el.Infer(env, nonGeneric)
		if err != nil {
			return nil, err
		}
		node.Fields[i] = field.(RecordField)
	}

	var params []string
	for _, p := range node.Params {
		params = append(params, p.String())
	}

	fields := make(map[string]Type)
	for _, p := range node.Fields {
		fields[p.Name] = p.InferredType
	}

	env.Env[node.Name] = Record{Name: node.Name, Params: params, Fields: fields}
	return node, nil
}

func (node Variant) Infer(env *State, nonGeneric []Type) (Ast, error) {

	for i, el := range node.Constructors {
		newEnv := State{Env: make(map[string]Type)}
		CopyState(*env, newEnv)
		newEnv.Env["__parent__"] = VariantType{Name: node.Name}
		cons, err := el.Infer(&newEnv, nonGeneric)
		if err != nil {
			return nil, err
		}
		node.Constructors[i] = cons.(VariantConstructor)
		env.Env[el.Name] = cons.(VariantConstructor).GetInferredType()
	}

	var params []string
	for _, p := range node.Params {
		params = append(params, p.String())
	}

	// get all of the types for each constructor, store in parent type
	consTypes := make([]string, 0)
	for _, c := range node.Constructors {
		consTypes = append(consTypes, c.GetInferredType().GetName())
	}

	env.Env[node.Name] = VariantType{Name: node.Name, Params: params, Constructors: consTypes}
	node.InferredType = env.Env[node.Name]

	// concrete := node.InferredType.(VariantType).checkConcrete()
	// if concrete {
	// 	env.Module.ConcreteTypes = append(env.Module.ConcreteTypes, node)
	// }

	return node, nil
}

func (node VariantType) checkConcrete() bool {
	// concrete := true

	// for _, constructor := range node.Constructors {
	// 	for _, f := range constructor.(VariantConstructorType).Types {
	// 		//fields = append(fields, f.GetName())
	// 		fmt.Printf("check concrete: %s\n", f.GetName())
	// 		if f.GetName()[0] == '\'' {
	// 			concrete = false
	// 			break
	// 		}
	// 	}
	// 	//fields[p.Name] = p.InferredType
	// }
	// fmt.Println(node.GetName(), "concrete status:", concrete)
	// return concrete
	return true
}

func (node VariantConstructor) Infer(env *State, nonGeneric []Type) (Ast, error) {
	parent := env.Env["__parent__"].(VariantType)

	types := make([]Type, 0)

	for i, field := range node.Fields {
		field, err := field.Infer(env, nonGeneric)
		if err != nil {
			return nil, err
		}
		node.Fields[i] = field
		types = append(types, field.GetInferredType())
	}

	node.InferredType = VariantConstructorType{Name: node.Name, Parent: parent.GetName(), Types: types}
	env.Env[node.Name] = node.InferredType

	return node, nil
}

func (node VariantInstance) Infer(env *State, nonGeneric []Type) (Ast, error) {
	vType, ok := env.Env[node.Name]
	pName := vType.(VariantConstructorType).Parent
	if !ok {
		return nil, InferenceError{"Don't know the type of this variable: " + node.Name}
	}

	if len(node.Arguments) != len(vType.(VariantConstructorType).Types) {
		return nil, InferenceError{fmt.Sprintf("%s required %d arguments, but was supplied with %d",
			node.Name, len(vType.(VariantConstructorType).Types), len(node.Arguments))}
	}

	types := []Type{}

	for i, t := range vType.(VariantConstructorType).Types {
		arg := node.Arguments[i]
		arg, err := arg.Infer(env, nonGeneric)
		if err != nil {
			return nil, err
		}
		argType := arg.GetInferredType()

		//fmt.Println(node.Name, "arg type is", argType)
		err = Unify(argType, t, env)
		if err != nil {
			return nil, InferenceError{fmt.Sprintf("Argument %d to %s doesn't match type %s", i+1, node.Name, t.GetName())}
		}
		types = append(types, argType)
	}

	pType := env.Env[vType.(VariantConstructorType).Parent].(VariantType)
	for i, el := range pType.Constructors {
		if el == node.Name {
			node.Constructor = i
			break
		}
	}
	node.InferredType = VariantInstanceType{Name: pName, Types: types}
	// switch this to be a parent type
	node.Name = pName
	return node, nil
}

func (node ExternRecordType) Infer(env *State, nonGeneric []Type) (Ast, error) {
	for i, el := range node.Fields {
		field, err := el.Infer(env, nonGeneric)
		if err != nil {
			return nil, err
		}
		node.Fields[i] = field.(RecordField)
	}

	var params []string

	fields := make(map[string]Type)
	for _, p := range node.Fields {
		fields[p.Name] = p.InferredType
	}

	env.Env[node.Name] = Record{Name: node.Name, Params: params, Fields: fields}
	return node, nil
}

func (node RecordInstance) Infer(env *State, nonGeneric []Type) (Ast, error) {
	el, ok := env.Env[node.Name]
	if !ok {
		return nil, InferenceError{"Don't know this type: " + node.String()}
	}
	node.InferredType = el
	return node, nil
}

func (node RecordAccess) Infer(env *State, nonGeneric []Type) (Ast, error) {
	if len(node.Identifiers) < 2 {
		return nil, InferenceError{"This record type has no fields: " + node.String()}
	}
	varName := node.Identifiers[0].StringValue
	recordType, ok := env.Env[varName]
	if !ok {
		return nil, InferenceError{"Don't know the type of this variable: " + varName}
	}
	fieldName := node.Identifiers[1].StringValue
	fieldType := recordType.(Record).Fields[fieldName]
	node.Identifiers[1].StringValue = strings.Title(node.Identifiers[1].StringValue)
	node.InferredType = fieldType
	env.UsedVariables[varName] = true
	return node, nil
}

func (node Return) Infer(env *State, nonGeneric []Type) (Ast, error) {
	// added by inference passes, should not need checking
	return node, nil
}

func (node ReturnTuple) Infer(env *State, nonGeneric []Type) (Ast, error) {
	// added by inference passes, should not need checking
	return node, nil
}

func (node ExternFunc) Infer(env *State, nonGeneric []Type) (Ast, error) {
	// make our function type
	var fType Function
	if HasImportPath(node.Import) {
		fType = Function{Name: node.Name, External: "__go_" + StripImportPath(node.Import), Types: []Type{}}
	} else {
		fType = Function{Name: node.Name, External: node.Import, Types: []Type{}}
	}

	// grab inferred types of args
	if len(node.Arguments) > 0 {
		for _, el := range node.Arguments {
			// lookup the type from its annotation
			//fmt.Println("looking up arg type annotation:", el.(Identifier).Annotation)
			anno, err := GetTypeFromAnnotation(el.(Identifier).Annotation, env)
			if err != nil {
				return nil, err
			}
			fType.Types = append(fType.Types, anno)
		}
	}

	// lookup return type annotation
	//fmt.Println("looking up return type annotation:", node.ReturnAnnotation)
	anno, err := GetTypeFromAnnotation(node.ReturnAnnotation, env)
	if err != nil {
		return nil, err
	}
	fType.Types = append(fType.Types, anno)
	(*env).Env[node.Name] = fType
	node.InferredType = fType

	return node, nil
}

func GetTypeFromAnnotation(name Ast, env *State) (Type, error) {
	types := make(map[string]Type)
	types["int64"] = Integer
	types["float64"] = Float
	types["string"] = String
	types["rune"] = Rune
	types["byte"] = Byte
	types["bool"] = Boolean
	types["()"] = Unit
	types["'any"] = Any

	switch name.(type) {
	case BasicAst:
		tName := name.(BasicAst).StringValue
		if val, ok := types[tName]; ok {
			return val, nil
		}

		if val, ok := env.Env[tName]; ok {
			return val, nil
		}
	case Identifier:
		tName := name.(Identifier).StringValue
		if val, ok := types[tName]; ok {
			return val, nil
		}

		if val, ok := env.Env[tName]; ok {
			return val, nil
		}
	case Container:
		// containers are a list of types for a func (last one is return type)
		// we need to make sure each type in here matches too
		c := name.(Container)
		switch c.Type {
		case "FuncAnnotation":
			var types []Type
			for _, el := range c.Subvalues {
				t, err := GetTypeFromAnnotation(el, env)
				if err != nil {
					return nil, err
				}
				types = append(types, t)
			}

			return Function{Name: NewTempVariable(), Types: types}, nil
		}
	case ReturnTuple:
		tuple := name.(ReturnTuple)
		var names []string

		var types []Type

		for _, el := range tuple.Subvalues {
			t, err := GetTypeFromAnnotation(el, env)
			if err != nil {
				return nil, err
			}
			types = append(types, t)
			names = append(names, t.GetName())
		}

		return List{Name: fmt.Sprintf("(%s)", strings.Join(names, ",")), Types: types}, nil
	case ArrayType:
		c := name.(ArrayType)
		var types []Type

		t, err := GetTypeFromAnnotation(c.Subtype, env)
		if err != nil {
			return nil, err
		}
		types = append(types, t)

		return List{Name: fmt.Sprintf("[]%s", t.GetName()), Types: types}, nil
	}

	return nil, InferenceError{fmt.Sprintf("Do not know annotated type '%s'", name)}
}

func GetType(name string, env State, nonGeneric []Type) (Type, error) {
	if _, ok := env.Env[name]; ok {
		return Fresh(env.Env[name].(Type), nonGeneric), nil
	}
	return nil, InferenceError{"Undefined type " + name}
}

func Unify(t1 Type, t2 Type, env *State) error {
	/* Unify the two types t1 and t2.

	Makes the types t1 and t2 the same.

	Args:
		t1: The first type to be made equivalent
		t2: The second type to be be equivalent

	Returns:
		An error if the types cannot be unified.
	*/
	a := Prune(t1)
	b := Prune(t2)

	if a.GetName() == Any.GetName() || b.GetName() == Any.GetName() {
		// Any unifies with everything
		return nil
	}

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
			return Unify(b, a, env)
		case TypeOperator:
			aTypeLen := len(a.(TypeOperator).Types)
			bTypeLen := len(b.(TypeOperator).Types)
			if a.GetName() != b.GetName() || aTypeLen != bTypeLen {
				if len(b.(TypeOperator).Types) > 0 {
					return Unify(b, a, env)
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
				err := Unify(el, b.(TypeOperator).Types[i], env)
				if err != nil {
					return err
				}
			}
			return nil
		}
	case Function:
		switch b.(type) {
		case Function:
			aTypeLen := len(a.(Function).Types)
			bTypeLen := len(b.(Function).Types)
			if aTypeLen != bTypeLen {
				return InferenceError{fmt.Sprintf("Type mismatch: %s != %s (different number of arguments)", a.GetName(), b.GetName())}
			}
			// we know that the types must match because they didn't pass into that last condition
			for i, el := range a.(Function).Types {
				err := Unify(el, b.(Function).Types[i], env)
				if err != nil {
					return err
				}
			}
			return nil
		}
	case Record:
		switch b.(type) {
		case Record:
			aFieldLen := len(a.(Record).Fields)
			bFieldLen := len(b.(Record).Fields)
			if aFieldLen != bFieldLen {
				return InferenceError{fmt.Sprintf("Type mismatch: %s != %s (different number of fields)", a.GetName(), b.GetName())}
			}

			for i, el := range a.(Record).Fields {
				field := b.(Record).Fields[i]
				err := Unify(el, field, env)
				if err != nil {
					return err
				}
			}

			return nil
		}
	}

	return InferenceError{fmt.Sprintf("No match for these types, not unified:\n* %s\n* %s", a.GetName(), b.GetName())}
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
				mappings[p.(TypeVariable)] = NewTypeVariable().(TypeVariable)
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
