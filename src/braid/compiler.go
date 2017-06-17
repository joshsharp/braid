package main

import (
	"fmt"
	"strings"
)

func (a BasicAst) Compile(state State) string {
	switch (a.ValueType){
	case STRING:
		switch (a.Type){
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
	case CONTAINER:
		switch (a.Type){
		case "Array":
			values := "["
			for _, el := range (a.Subvalues) {
				values += el.Compile(state) + ","
			}
			return values + "]"
		case "BinOpParens":
			values := "("
			for _, el := range (a.Subvalues) {
				values += el.Compile(state)
			}
			return values + ")"
		default:
			values := ""
			for _, el := range (a.Subvalues) {
				values += el.Compile(state)
			}
			return values
		}
		
		//fmt.Println("container")
		
	default:
		return ""
	}
	return ""
}

func (a Assignment) Compile(state State) string {
	result := ""

	left := make([]string,0)
	for _, el := range(a.Left){
		left = append(left, el.Compile(state))
	}
	
	result += strings.Join(left, ", ")
	result += " := "
	result += a.Right.Compile(state)
	return result + "\n"
	
}

func (a If) Compile(state State) string {
	result := "\nif "
	
	result += "(" + a.Condition.Compile(state) + ") {\n"
	then := ""
	
	for _, el := range(a.Then){
		// compile each sub AST
		// make a result then indent each line
		then += el.Compile(state)
	}
	
	for _, el := range(strings.Split(then, "\n")){
		result += "\t" + el + "\n"
	}
	
	result += "}"
	if a.Else == nil {
		return result + "\n\n"
	}
	
	result += " else {\n"
	elser := ""
	
	for _, el := range(a.Else){
		// compile each sub AST
		// make a result then indent each line
		elser += el.Compile(state)
	}
	
	for _, el := range(strings.Split(elser, "\n")){
		result += "\t" + el + "\n"
	}
	
	result += "}\n"
	
	return result
	
	
}

func (a Call) Compile(state State) string {
	result := ""
	if a.Module != nil {
		result += a.Module.Compile(state) + "."
	}
	result += a.Function.Compile(state) + "("
	if len(a.Arguments) > 0 {
		args := make([]string,0)
		for _, el := range(a.Arguments){
			args = append(args, el.Compile(state))
		}
		result += strings.Join(args, ", ")
	}
	result += ")"
	
	return result
}

func (a Func) Compile(state State) string {
	result := "func " + a.Name + " ("
	if len(a.Arguments) > 0 {
		args := make([]string,0)
		for _, el := range(a.Arguments){
			args = append(args, el.Compile(state))
		}
		result += strings.Join(args, ", ")
	}
	result += ") {\n"
	
	inner := ""
	
	for _, el := range(a.Subvalues){
		// compile each sub AST
		// make a result then indent each line
		inner += el.Compile(state)
	}
	
	for _, el := range(strings.Split(inner, "\n")){
		result += "\t" + el + "\n"
	}
	
	result += "}\n\n"
	
	return result
}

func (a AliasType) Compile(state State) string{
	return ""
}

func (r RecordType) Compile(state State) string{
	str := "type " + r.Name + " struct {\n"
	
	inner := ""
	
	for _, el := range(r.Fields){
		// compile each sub AST
		// make a result then indent each line
		inner += el.Compile(state)
	}
	
	for _, el := range(strings.Split(inner, "\n")){
		str += "\t" + el + "\n"
	}
	
	str += "}\n\n"
	return str
}

func (v VariantType) Compile(state State) string {
	return ""
}

func (f RecordField) Compile(state State) string {
	return f.Name + " " + f.Type.Compile(state) + "\n"
}