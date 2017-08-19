package compiler

import (
	"fmt"
	"strings"
	"braid/types"
	"braid/ast"
)

type Ast interface {
	Compile(state types.State) string
}

func (m ast.Module) Compile(state types.State) string {
	values := "package main\n\n"
	for _, el := range m.Subvalues {
		values += el.Compile(state)
	}
	return values
}

func (a ast.BasicAst) Compile(state types.State) string {
	switch a.ValueType {
	case ast.STRING:
		switch a.Type {
		case "Comment":
			return fmt.Sprintf("//%s\n", a.StringValue)
		case "String":
			return fmt.Sprintf("\"%s\"", a.StringValue)
		case "StringOperator":
			return fmt.Sprintf(" %s ", a.StringValue)
		case "IntOperator":
			return fmt.Sprintf(" %s ", a.StringValue)
		case "FloatOperator":
			return fmt.Sprintf(" %s ", a.StringValue)
		default:
			return fmt.Sprintf("%s", a.StringValue)
		}
	case ast.CHAR:
		return fmt.Sprintf("'%s'", string(a.CharValue))
	case ast.INT:
		return fmt.Sprintf("%d", a.IntValue)
	case ast.FLOAT:
		return fmt.Sprintf("%f", a.FloatValue)
	case ast.BOOL:
		if a.BoolValue {
			return "true"
		}
		return "false"

	default:
		return ""
	}
	return ""
}

func (a ast.ArrayType) Compile(state types.State) string {
	values := "[]int{"
	for _, el := range a.Subvalues {
		values += el.Compile(state) + ","
	}
	return values + "}"
}

func (c ast.Container) Compile(state types.State) string {
	switch c.Type {
	case "BinOpParens":
		values := "("
		for _, el := range c.Subvalues {
			values += el.Compile(state)
		}
		return values + ")"
	default:
		values := ""
		for _, el := range c.Subvalues {
			values += el.Compile(state)
		}
		return values
	}
}

func (a ast.Assignment) Compile(state types.State) string {
	result := ""

	left := make([]string, 0)
	for _, el := range a.Left {
		left = append(left, el.Compile(state))
	}

	result += strings.Join(left, ", ")
	result += " := "
	result += a.Right.Compile(state)
	return result + "\n"

}

func (a ast.If) Compile(state types.State) string {
	result := "\nif "

	result += "(" + a.Condition.Compile(state) + ") {\n"
	then := ""

	for _, el := range a.Then {
		// compile each sub AST
		// make a result then indent each line
		then += el.Compile(state)
	}

	for _, el := range strings.Split(then, "\n") {
		result += "\t" + el + "\n"
	}

	result += "}"
	if a.Else == nil {
		return result + "\n\n"
	}

	result += " else {\n"
	elser := ""

	for _, el := range a.Else {
		// compile each sub AST
		// make a result then indent each line
		elser += el.Compile(state)
	}

	for _, el := range strings.Split(elser, "\n") {
		result += "\t" + el + "\n"
	}

	result += "}\n"

	return result

}

func (a ast.Call) Compile(state types.State) string {
	result := ""
	if a.Module != nil {
		result += a.Module.Compile(state) + "."
	}
	result += a.Function.Compile(state) + "("
	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for _, el := range a.Arguments {
			args = append(args, el.Compile(state))
		}
		result += strings.Join(args, ", ")
	}
	result += ")\n"

	return result
}

func (a ast.VariantInstance) Compile(state types.State) string {
	result := ""
	result += a.Name + "{"
	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for _, el := range a.Arguments {
			args = append(args, el.Compile(state))
		}
		result += strings.Join(args, ", ")
	}
	result += "}\n"

	return result
}

func (a ast.RecordInstance) Compile(state types.State) string {
	result := ""
	result += a.Name + "{"
	if len(a.Values) > 0 {
		args := make([]string, 0)
		for key, el := range a.Values {
			val := ""
			val += key + ": "
			val += el.Compile(state)
			args = append(args, val)
		}
		result += strings.Join(args, ", ")
	}
	result += "}\n"

	return result
}

func (a ast.Func) Compile(state types.State) string {
	// TODO: Only compile once we have concrete implementations
	result := "func " + a.Name + " ("
	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for _, el := range a.Arguments {
			args = append(args, el.Compile(state))
		}
		result += strings.Join(args, ", ")
	}
	result += ") {\n"

	inner := ""

	for _, el := range a.Subvalues {
		// compile each sub AST
		// make a result then indent each line
		inner += el.Compile(state)
	}

	for _, el := range strings.Split(inner, "\n") {
		result += "\t" + el + "\n"
	}

	result += "}\n\n"

	return result
}

func (a ast.AliasType) Compile(state types.State) string {
	// TODO: Only compile once we have concrete implementations
	return "type " + a.Name + " int32\n\n"
}

func (r ast.RecordType) Compile(state types.State) string {
	// TODO: Only compile once we have concrete implementations
	str := "type " + r.Name + " struct {\n"

	inner := ""

	for _, el := range r.Fields {
		// compile each sub AST
		// make a result then indent each line
		inner += el.Compile(state)
	}

	for _, el := range strings.Split(inner, "\n") {
		str += "\t" + el + "\n"
	}

	str += "}\n\n"
	return str
}

func (v ast.VariantType) Compile(state types.State) string {
	// TODO: Only compile once we have concrete implementations
	str := "type " + v.Name + " interface {\n" +
		"\tsumtype()\n" +
		"}\n\n"

	for _, el := range v.Constructors {
		str += el.Compile(state)
	}

	return str
}

func (c ast.VariantConstructor) Compile(state types.State) string {
	str := "type " + c.Name + " struct {\n"
	for i, el := range c.Fields {
		str += fmt.Sprintf("\tF%d", i) + " " + el.Compile(state)
	}
	str += "\n}\n\n"

	// implement sealed
	str += "func (*" + c.Name + ") sumtype() {}\n\n"

	return str
}

func (f ast.RecordField) Compile(state types.State) string {
	return f.Name + " " + f.Type.Compile(state) + "\n"
}
