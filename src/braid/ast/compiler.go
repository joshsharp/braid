package ast

import (
	"fmt"
	"strings"
)

func GetImportPath(imprt string) string {
	pathParts := strings.Split(imprt, ".")
	return pathParts[0]
}

func StripImportPath(extern string) string {
	pathParts := strings.Split(extern, "/")
	return pathParts[len(pathParts)-1]
}

func (m Module) Compile(state State) (string, State) {
	values := fmt.Sprintf("package %s\n\n", strings.ToLower(m.Name))
	for _, el := range m.Subvalues {
		value, s := el.Compile(state)
		values += value
		state = s
	}
	return values, state
}

func (a BasicAst) Compile(state State) (string, State) {
	switch a.ValueType {
	case STRING:
		switch a.Type {
		case "Comment":
			return fmt.Sprintf("//%s\n", a.StringValue), state
		case "String":
			return fmt.Sprintf("\"%s\"", a.StringValue), state
		default:
			return fmt.Sprintf("%s", a.StringValue), state
		}
	case CHAR:
		return fmt.Sprintf("'%s'", string(a.CharValue)), state
	case INT:
		return fmt.Sprintf("%d", a.IntValue), state
	case FLOAT:
		return fmt.Sprintf("%f", a.FloatValue), state
	case BOOL:
		if a.BoolValue {
			return "true", state
		}
		return "false", state
	case NIL:
		return "nil", state
	default:
		return "", state
	}
	return "", state
}

func (o Operator) Compile(state State) (string, State) {
	ops := map[string]string{
		"+":  "+",
		"-":  "-",
		"*":  "*",
		"/":  "/",
		"<":  "<",
		">":  ">",
		"==": "==",
		"++": "+",
	}

	return fmt.Sprintf(" %s ", ops[o.StringValue]), state
}

func (c Comment) Compile(state State) (string, State) {
	return fmt.Sprintf("//%s\n", c.StringValue), state
}

func (i Identifier) Compile(state State) (string, State) {
	return i.StringValue, state
}

func (a ArrayType) Compile(state State) (string, State) {
	fmt.Println(a.Print(0))
	values := fmt.Sprintf("[]%s{", a.InferredType.GetName())
	for _, el := range a.Subvalues {
		value, s := el.Compile(state)
		values += value + ","
		state = s
	}
	return values + "}", state
}

func (c Container) Compile(state State) (string, State) {
	switch c.Type {
	case "BinOpParens":
		values := "("
		for _, el := range c.Subvalues {

			value, s := el.Compile(state)
			values += value
			state = s
		}
		return values + ")", state
	default:
		values := ""
		for _, el := range c.Subvalues {
			value, s := el.Compile(state)
			values += value
			state = s
		}
		return values, state
	}
}

func (e Expr) Compile(state State) (string, State) {
	switch e.Type {
	case "BinOpParens":
		values := "("
		for _, el := range e.Subvalues {

			value, s := el.Compile(state)
			values += value
			state = s
		}
		return values + ")", state
	default:
		values := ""
		for _, el := range e.Subvalues {
			value, s := el.Compile(state)
			values += value
			state = s
		}
		if e.AsStatement {
			values += "\n"
		}
		return values, state
	}
}

func (a RecordAccess) Compile(state State) (string, State) {
	var bits []string
	for _, el := range a.Identifiers {
		bits = append(bits, el.String())
	}
	val := strings.Join(bits, ".")
	return val, state
}

func (a Assignment) Compile(state State) (string, State) {
	result := ""

	var varName string

	if _, ok := state.UsedVariables[a.Left.(Identifier).StringValue]; !ok {
		// if this identifier is in not here, means it's unused
		// so return '_'
		varName = "_"
	} else {
		value, s := a.Left.Compile(state)
		state = s
		varName = value
	}

	switch a.Right.(type) {
	case If:
		value, s := a.Right.Compile(state)
		state = s
		result += value + "\n"

		result += varName
		if a.Update || varName == "_" {
			result += " = "
		} else {
			result += " := "
		}

		result += a.Right.(If).TempVar
	default:

		result += varName
		if a.Update || varName == "_" {
			result += " = "
		} else {
			result += " := "
		}

		value, s := a.Right.Compile(state)
		result += value
		state = s
	}

	return result + "\n", state

}

func (r Return) Compile(state State) (string, State) {
	result := "\nreturn "
	value, s := r.Value.Compile(state)
	result += value
	state = s

	return result, state

}

func (a If) Compile(state State) (string, State) {
	result := fmt.Sprintf("var %s %s\n", a.TempVar, a.InferredType.GetName())
	result += "\nif "

	value, s := a.Condition.Compile(state)
	result += value + " {\n"
	state = s
	then := ""

	for _, el := range a.Then {
		// compile each sub AST
		// make a result then indent each line
		value, s := el.Compile(state)
		state = s
		then += value
	}

	for _, el := range strings.Split(then, "\n") {
		result += "\t" + el + "\n"
	}

	result += "}"
	if a.Else == nil {
		return result + "\n\n", state
	}

	result += " else {\n"
	elser := ""

	for _, el := range a.Else {
		// compile each sub AST
		// make a result then indent each line
		value, s := el.Compile(state)
		state = s
		elser += value
	}

	for _, el := range strings.Split(elser, "\n") {
		result += "\t" + el + "\n"
	}

	result += "}\n"

	return result, state

}

func (b BinOp) Compile(state State) (string, State) {
	result := ""

	value, s := b.Left.Compile(state)
	state = s
	result += value
	value, s = b.Operator.Compile(state)
	state = s
	result += value
	value, s = b.Right.Compile(state)
	state = s
	result += value

	return result, state

}

func (a Call) Compile(state State) (string, State) {
	result := ""
	if a.Module.StringValue != "" {
		value := StripImportPath(a.Module.StringValue)
		result += value + "."
	}
	value, s := a.Function.Compile(state)
	state = s
	result += value + "("
	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for _, el := range a.Arguments {
			value, s := el.Compile(state)
			state = s

			args = append(args, value)
		}
		result += strings.Join(args, ", ")
	}
	result += ")"

	return result, state
}

func (a VariantInstance) Compile(state State) (string, State) {
	result := ""
	result += a.Name + "{"
	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for _, el := range a.Arguments {
			value, s := el.Compile(state)
			state = s
			args = append(args, value)
		}
		result += strings.Join(args, ", ")
	}
	result += "}\n"

	return result, state
}

func (a RecordInstance) Compile(state State) (string, State) {
	result := ""
	result += a.Name + "{"
	if len(a.Values) > 0 {
		args := make([]string, 0)
		for key, el := range a.Values {
			val := ""
			val += strings.Title(key) + ": "
			value, s := el.Compile(state)
			val += value
			state = s
			args = append(args, val)
		}
		result += strings.Join(args, ", ")
	}
	result += "}\n"

	return result, state
}

func (e ExternRecordType) Compile(state State) (string, State) {
	str := ""
	path := GetImportPath(e.Import)
	name := "__go_" + StripImportPath(path)

	if _, ok := state.Imports[name]; ok {
		return "", state
	}

	state.Imports[name] = true
	str += fmt.Sprintf("import %s \"%s\"\n", name, path)
	return str, state

}

func (e ExternFunc) Compile(state State) (string, State) {
	// TODO: handle nested packages

	path := GetImportPath(e.Import)
	name := "__go_" + StripImportPath(path)

	if _, ok := state.Imports[name]; ok {
		return "", state
	}

	state.Imports[name] = true

	// TODO: handle tracking whether functions are actually called - not sure how to get root state
	//if _, ok := state.UsedVariables[e.Name]; !ok {
	//	return fmt.Sprintf("import _ \"%s\"\n", path[0]), state
	//} else {
	state.UsedVariables[e.Name] = true
	return fmt.Sprintf("import %s \"%s\"\n", name, path), state
	//}
}

func (a Func) Compile(state State) (string, State) {

	types := a.InferredType.(Function).Types
	typesLen := len(types)

	for _, el := range types {
		if el.GetName()[0] == '\'' {
			return fmt.Sprintf("// func `%s` not added, not concrete\n", a.Name), state
		}
	}

	result := ""

	if _, ok := state.Env["scope"]; ok {
		var varName string

		if _, ok := state.UsedVariables[a.Name]; !ok {
			// if this identifier is in not here, means it's unused
			// so return '_'
			varName = "_"
		} else {
			varName = a.Name
		}

		result += varName + " := func ("
	} else {
		result += "func " + a.Name + " ("
	}

	if len(a.Arguments) > 0 {
		args := make([]string, 0)
		for i, el := range a.Arguments {
			argName, s := el.Compile(state)
			argType := types[i].GetName()
			arg := fmt.Sprintf("%s %s", argName, argType)
			args = append(args, arg)
			state = s
		}
		result += strings.Join(args, ", ")

	}
	result += ") "
	if typesLen > 0 {
		result += fmt.Sprintf("%s {\n", types[typesLen-1].GetName())
	} else {
		result += "{\n"
	}

	inner := ""
	//innerState := State{Env:make(map[string]Type), UsedVariables:make(map[string]bool)}
	//CopyState(newState, innerState)
	newState := state.Env[a.Name].(Function).Env
	newState.Imports = state.Imports
	newState.Env["scope"] = Function{}

	for _, el := range a.Subvalues {
		// compile each sub AST
		// make a result then indent each line
		value, s := el.Compile(newState)

		inner += value
		newState = s
	}

	lines := strings.Split(inner, "\n")

	for _, el := range lines {
		result += "\t" + el + "\n"
	}

	result += "}\n\n"

	return result, state
}

func (a AliasType) Compile(state State) (string, State) {
	// TODO: Only compile once we have concrete implementations
	return "type " + a.Name + " int32\n\n", state
}

func (r RecordType) Compile(state State) (string, State) {
	// TODO: Only compile once we have concrete implementations
	str := "type " + r.Name + " struct {\n"

	inner := ""

	for _, el := range r.Fields {
		// compile each sub AST
		// make a result then indent each line
		value, s := el.Compile(state)
		inner += value
		state = s
	}

	for _, el := range strings.Split(inner, "\n") {
		str += "\t" + el + "\n"
	}

	str += "}\n\n"
	return str, state
}

func (v VariantType) Compile(state State) (string, State) {
	// TODO: Only compile once we have concrete implementations
	str := "type " + v.Name + " interface {\n" +
		"\tsumtype()\n" +
		"}\n\n"

	for _, el := range v.Constructors {
		value, s := el.Compile(state)

		str += value
		state = s
	}

	return str, state
}

func (c VariantConstructor) Compile(state State) (string, State) {
	str := "type " + c.Name + " struct {\n"
	for i, el := range c.Fields {
		value, s := el.Compile(state)
		state = s
		str += fmt.Sprintf("\tF%d", i) + " " + value
	}
	str += "\n}\n\n"

	// implement sealed
	str += "func (*" + c.Name + ") sumtype() {}\n\n"

	return str, state
}

func (f RecordField) Compile(state State) (string, State) {
	value, s := f.Type.Compile(state)
	state = s
	return strings.Title(f.Name) + " " + value + "\n", state
}
