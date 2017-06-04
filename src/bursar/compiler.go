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
		//fmt.Println("container")
		values := ""
		for _, el := range (a.Subvalues) {
			values += el.Compile(state)
		}
		return values
	default:
		return ""
	}
	return ""
}

func (a Assignment) Compile(state State) string {
	result := ""

	switch a.Right.(type) {
	case Func:
		result += "func " + a.Left[0].Compile(state)
		result += a.Right.Compile(state)
		return result
	
	default:
		left := make([]string,0)
		for _, el := range(a.Left){
			left = append(left, el.Compile(state))
		}
		
		result += strings.Join(left, ", ")
		result += " := "
		result += a.Right.Compile(state)
		return result + "\n"
	}
	
	return result
}

func (a If) Compile(state State) string {
	return "if { }"
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
	result := "("
	if len(a.Arguments) > 0 {
		args := make([]string,0)
		for _, el := range(a.Arguments){
			args = append(args, el.Compile(state))
		}
		result += strings.Join(args, ", ")
	}
	result += ") {\n"
	
	for _, el := range(a.Subvalues){
		result += el.Compile(state)
	}
	
	result += "\n}\n\n"
	
	return result
}
