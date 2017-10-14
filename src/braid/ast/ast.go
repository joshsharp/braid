package ast

import (
	"fmt"
	"strings"
)

type ValueType int

const (
	STRING = iota
	INT
	FLOAT
	NUMBER
	BOOL
	CHAR
	NIL
)

type Module struct {
	Name      string
	Subvalues []Ast
}

type BasicAst struct {
	Type         string
	StringValue  string
	CharValue    rune
	BoolValue    bool
	IntValue     int
	FloatValue   float64
	ValueType    ValueType
	InferredType Type
}

type Operator struct {
	StringValue  string
	ValueType    ValueType
	InferredType Type
}

type Comment struct {
	StringValue string
}

type Identifier struct {
	StringValue  string
	InferredType Type
}

type Expr struct {
	Type         string
	Subvalues    []Ast
	InferredType Type
}

type BinOp struct {
	Operator     Ast
	Left         Ast
	Right        Ast
	InferredType Type
}

type Container struct {
	Type         string
	Subvalues    []Ast
	InferredType Type
}

type ArrayType struct {
	Subvalues    []Ast
	InferredType Type
}

type Func struct {
	Name         string
	Arguments    []Ast
	Subvalues    []Ast
	InferredType Type
}

type Call struct {
	Module       Ast
	Function     Ast
	Arguments    []Ast
	InferredType Type
}

type VariantInstance struct {
	Name      string
	Arguments []Ast
}

type RecordInstance struct {
	Name   string
	Values map[string]Ast
}

type If struct {
	Condition    Ast
	Then         []Ast
	Else         []Ast
	InferredType Type
}

type Assignment struct {
	Left         Ast
	Right        Ast
	InferredType Type
}

type RecordType struct {
	Name         string
	Fields       []RecordField
	Constructors []Ast
}

type VariantType struct {
	Name         string
	Params       []Ast
	Constructors []VariantConstructor
}

type AliasType struct {
	Name   string
	Params []Ast
	Types  []Ast
}

type RecordField struct {
	Name string
	Type Ast
}

type VariantConstructor struct {
	Name   string
	Fields []Ast
}

type Ast interface {
	String() string
	Print(indent int) string
	Compile(state State) string
	GetInferredType() Type
}

func (a BasicAst) String() string {
	switch (a.ValueType) {
	case STRING:
		if a.Type == "Comment" {
			return fmt.Sprintf("//%s", a.StringValue)
		} else {
			return fmt.Sprintf("\"%s\"", a.StringValue)
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
	case NIL:
		return "nil"
	}
	return "()"
}

func (o Operator) SetInferredType(t Type) {
	o.InferredType = t
}


func (c Container) SetInferredType(t Type) {
	c.InferredType = t
}

func (m Module) SetInferredType(t Type) {

}

func (c Call) SetInferredType(t Type) {
	c.InferredType = t
}

func (e Expr) SetInferredType(t Type) {
	e.InferredType = t
}

func (b BinOp) SetInferredType(t Type) {
	b.InferredType = t
}

func (a Assignment) SetInferredType(t Type) {
	a.InferredType = t
}

func (c Comment) SetInferredType(t Type) {

}

func (i Identifier) SetInferredType(t Type) {
	i.InferredType = t
}

func (r RecordType) SetInferredType(t Type) {

}

func (v VariantType) SetInferredType(t Type) {

}

func (a ArrayType) SetInferredType(t Type) {
	a.InferredType = t
}

func (a BasicAst) SetInferredType(t Type) {
	a.InferredType = t
}

func (f Func) SetInferredType(t Type) {
	f.InferredType = t
}

func (f If) SetInferredType(t Type) {
	f.InferredType = t
}

func (o Operator) GetInferredType() Type {
	return o.InferredType
}

func (c Container) GetInferredType() Type {
	return c.InferredType
}

func (m Module) GetInferredType() Type {
	return Unit
}

func (c Call) GetInferredType() Type {
	return c.InferredType
}

func (e Expr) GetInferredType() Type {
	return e.InferredType
}

func (b BinOp) GetInferredType() Type {
	return b.InferredType
}

func (a Assignment) GetInferredType() Type {
	return a.InferredType
}

func (c Comment) GetInferredType() Type {
	return Unit
}

func (i Identifier) GetInferredType() Type {
	return i.InferredType
}

func (r RecordType) GetInferredType() Type {
	return Unit
}

func (v VariantType) GetInferredType() Type {
	return Unit
}

func (a ArrayType) GetInferredType() Type {
	return a.InferredType
}

func (a BasicAst) GetInferredType() Type {
	return a.InferredType
}

func (f Func) GetInferredType() Type {
	return f.InferredType
}

func (f If) GetInferredType() Type {
	return f.InferredType
}


func (o Operator) String() string {
	return o.StringValue
}


func (c Container) String() string {
	values := ""
	for _, el := range c.Subvalues {
		values += fmt.Sprint(el)
	}
	return values
}

func (m Module) String() string {
	return m.Name
}

func (c Call) String() string {
	return "Call to " + c.Function.String()
}

func (e Expr) String() string {
	values := ""
	for i, el := range e.Subvalues {
		if (i > 0 ) {
			values += " "
		}
		values += fmt.Sprint(el)
	}
	return values
}

func (b BinOp) String() string {
	values := ""
	values += fmt.Sprint(b.Left)
	values += fmt.Sprint(b.Operator)
	values += fmt.Sprint(b.Right)

	return values
}

func (a Assignment) String() string {
	return a.Left.String() + " = " + a.Right.String()
}

func (c Comment) String() string {
	return "#" + c.StringValue
}

func (i Identifier) String() string {
	return i.StringValue
}

func (r RecordType) String() string {
	return r.Name
}

func (v VariantType) String() string {
	return v.Name
}

func (a ArrayType) String() string {
	values := []string{}


	for _, el := range a.Subvalues {
		values = append(values, fmt.Sprint(el))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (m Module) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Module:\n"
	for _, el := range m.Subvalues {
		str += el.Print(indent + 1)
	}
	return str
}

func (c Container) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Container " + c.Type + ":\n"
	for _, el := range c.Subvalues {
		str += el.Print(indent + 1)
	}
	return str
}

func (c Expr) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Container " + c.Type + ":\n"
	for _, el := range c.Subvalues {
		str += el.Print(indent + 1)
	}
	return str
}

func (a ArrayType) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += fmt.Sprintf("Array: %s\n", a.InferredType)
	for _, el := range a.Subvalues {
		str += el.Print(indent + 1)
	}
	return str
}

func (a BasicAst) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += fmt.Sprintf("%s %s:\n", a.Type, a)
	return str
}

func (o Operator) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += fmt.Sprintf("%s:\n", o.StringValue)
	return str
}

func (c Comment) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += fmt.Sprintf("Comment: %s\n", c.StringValue)
	return str
}

func (i Identifier) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	return str + i.StringValue + "\n"
}

func (a Func) String() string {
	return "Func " + a.Name
}

func (i If) String() string {
	return "If"
}

func (a Func) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Func "
	str += a.Name
	if len(a.Arguments) > 0 {
		str += " (\n"
		for _, el := range (a.Arguments) {
			str += el.Print(indent + 1)
		}
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += ")"
	}
	str += fmt.Sprint(a.InferredType)
	str += "\n"
	for _, el := range (a.Subvalues) {
		str += el.Print(indent + 1)
	}
	return str
}

func (a Call) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Call:\n"
	if a.Module != nil {
		str += a.Module.Print(indent+1) + "."
	}
	str += a.Function.Print(indent + 1)

	if len(a.Arguments) > 0 {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "(\n"
		for _, el := range (a.Arguments) {
			str += el.Print(indent + 1)
		}
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += ")\n"
	}
	return str
}

func (a VariantInstance) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "VariantInstance: "
	str += a.Name + "\n"

	if len(a.Arguments) > 0 {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "(\n"
		for _, el := range (a.Arguments) {
			str += el.Print(indent + 1)
		}
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += ")\n"
	}
	return str
}

func (a RecordInstance) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "RecordInstance: "
	str += a.Name + "\n"

	if len(a.Values) > 0 {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "(\n"
		for key, el := range (a.Values) {
			str += key + ":"
			str += el.Print(indent + 1)
		}
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += ")\n"
	}
	return str
}

func (i If) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "If"
	if i.Condition != nil {
		str += ":\n"
		str += i.Condition.Print(indent + 1)

	}
	for _, el := range (i.Then) {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "Then:\n"
		str += el.Print(indent + 1)
	}
	for _, el := range (i.Else) {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "Else:\n"
		str += el.Print(indent + 1)

	}
	return str
}

func (b BinOp) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "BinOp \n"

	str += b.Left.Print(indent + 1)
	str += b.Operator.Print(indent + 1)
	str += b.Right.Print(indent + 1)

	return str
}

func (a Assignment) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Assignment:\n"
	str += a.Left.Print(indent + 1)
	str += a.Right.Print(indent + 1)

	return str
}

func (t RecordType) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Type Record:\n"

	return str
}

func (t VariantType) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Type Variant:\n"

	return str
}

func (t AliasType) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Type Alias:\n"

	return str
}
func (f RecordField) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Field:\n"

	return str
}

func (c VariantConstructor) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Constructor:\n"

	return str
}
