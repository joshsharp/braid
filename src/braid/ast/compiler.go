package ast

import (
	"fmt"
	"strings"
)


func (m Module) Compile(state State) string {
	values := "package main\n\n"
	for _, el := range m.Subvalues {
		values += el.Compile(state)
	}
	return values
}

func (a BasicAst) Compile(state State) string {
	switch a.ValueType {
	case STRING:
		switch a.Type {
		case "Comment":
			return fmt.Sprintf("//%s\n", a.StringValue)
		case "String":
			return fmt.Sprintf("\"%s\"", a.StringValue)
		default:
			return fmt.Sprintf("%s", a.StringValue)
		}
	case CHAR:
		return fmt.Sprintf("'%s'", string(a.CharValue))
	case INT:
		return fmt.Sprintf("%d", a.IntValue)
	case FLOAT:
		return fmt.Sprintf("%f", a.FloatValue)
	case BOOL:
		if a.BoolValue {
			return "true"
		}
		return "false"

	default:
		return ""
	}
	return ""
}

func (o Operator) Compile(state State) string {
	return fmt.Sprintf(" %s ", o.StringValue)
}

func (c Comment) Compile(state State) string {
	return fmt.Sprintf("//%s\n", c.StringValue)
}

func (i Identifier) Compile(state State) string {
	return i.StringValue
}

func (a ArrayType) Compile(state State) string {
	fmt.Println(a.Print(0))
	values := fmt.Sprintf("[]%s{", a.InferredType.GetName())
	for _, el := range a.Subvalues {
		values += el.Compile(state) + ","
	}
	return values + "}"
}

func (c Container) Compile(state State) string {
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

func (e Expr) Compile(state State) string {
	switch e.Type {
	case "BinOpParens":
		values := "("
		for _, el := range e.Subvalues {
			values += el.Compile(state)
		}
		return values + ")"
	default:
		values := ""
		for _, el := range e.Subvalues {
			values += el.Compile(state)
		}
		return values
	}
}

func (a Assignment) Compile(state State) string {
	result := ""

	result += a.Left.Compile(state)
	result += " := "
	result += a.Right.Compile(state)
	return result + "\n"

}

func (a If) Compile(state State) string {
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

func (b BinOp) Compile(state State) string {
	result := ""

	result += b.Left.Compile(state)
	result += b.Operator.Compile(state)
	result += b.Right.Compile(state)

	return result

}

func (a Call) Compile(state State) string {
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
	result += ")"

	return result
}

func (a VariantInstance) Compile(state State) string {
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

func (a RecordInstance) Compile(state State) string {
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

func (a Func) Compile(state State) string {

	types := a.InferredType.(Function).Types
	typesLen := len(types)

	result := "func " + a.Name + " ("
	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for i, el := range a.Arguments {
			argName := el.Compile(state)
			argType := types[i].GetName()
			arg := fmt.Sprintf("%s %s", argName, argType)
			args = append(args, arg)
		}
		result += strings.Join(args, ", ")
	}


	result += fmt.Sprintf(") %s {\n", types[typesLen-1].GetName() )

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

func (a AliasType) Compile(state State) string {
	// TODO: Only compile once we have concrete implementations
	return "type " + a.Name + " int32\n\n"
}

func (r RecordType) Compile(state State) string {
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

func (v VariantType) Compile(state State) string {
	// TODO: Only compile once we have concrete implementations
	str := "type " + v.Name + " interface {\n" +
		"\tsumtype()\n" +
		"}\n\n"

	for _, el := range v.Constructors {
		str += el.Compile(state)
	}

	return str
}

func (c VariantConstructor) Compile(state State) string {
	str := "type " + c.Name + " struct {\n"
	for i, el := range c.Fields {
		str += fmt.Sprintf("\tF%d", i) + " " + el.Compile(state)
	}
	str += "\n}\n\n"

	// implement sealed
	str += "func (*" + c.Name + ") sumtype() {}\n\n"

	return str
}

func (f RecordField) Compile(state State) string {
	return f.Name + " " + f.Type.Compile(state) + "\n"
}
