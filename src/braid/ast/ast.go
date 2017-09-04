package ast

import (
	"fmt"
)

type ValueType int

const (
    STRING = iota
    INT
    FLOAT
    BOOL
    CHAR
    NIL
)

type Module struct {
    Name string
    Subvalues []Ast
}

type BasicAst struct {
    Type string
    StringValue string
    CharValue rune
    BoolValue bool
    IntValue int
    FloatValue float64
    ValueType ValueType
}

type Operator struct {
	StringValue string
	ValueType ValueType
}

type Comment struct {
    StringValue string
}

type Identifier struct {
	StringValue string
}

type Expr struct {
	Type string
	Subvalues []Ast
}

type Container struct {
	Type string
	Subvalues []Ast
}

type ArrayType struct {
	Subvalues []Ast
}

type Func struct {
    Name string
    Arguments []Ast
    Subvalues []Ast
}

type Call struct {
    Module Ast
    Function Ast
    Arguments []Ast
}

type VariantInstance struct {
    Name string
    Arguments []Ast
}

type RecordInstance struct {
    Name string
    Values map[string]Ast
}

type If struct {
    Condition Ast
    Then []Ast
    Else []Ast
}

type Assignment struct {
    Left Ast
    Right Ast
}

type RecordType struct {
    Name string
    Fields []RecordField
    Constructors []Ast
}

type VariantType struct {
    Name string
    Params []Ast
    Constructors []VariantConstructor
}

type AliasType struct {
    Name string
    Params []Ast
    Types []Ast
}

type RecordField struct {
    Name string
    Type Ast
}

type VariantConstructor struct {
    Name string
    Fields []Ast
}

type Ast interface {
    String() string
    Print(indent int) string
    Compile(state State) string
}


func (a BasicAst) String() string {
    switch (a.ValueType){
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

func (o Operator) String() string {
	return o.StringValue
}

func (c Container) String() string {
	values := ""
	for _, el := range c.Subvalues{
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
	for i, el := range e.Subvalues{
		if (i > 0 ){
			values += " "
		}
		values += fmt.Sprint(el)
	}
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
	values := ""
	for _, el := range a.Subvalues{
		values += fmt.Sprint(el)
	}
	return values
}

func (m Module) Print(indent int) string {
    str := ""

    for i := 0; i < indent; i++ {
        str += "  "
    }
    str += "Module:\n"
    for _, el := range m.Subvalues{
        str += el.Print(indent+1)
    }
    return str
}

func (c Container) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Container " + c.Type + ":\n"
	for _, el := range c.Subvalues{
		str += el.Print(indent+1)
	}
	return str
}

func (c Expr) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Container " + c.Type + ":\n"
	for _, el := range c.Subvalues{
		str += el.Print(indent+1)
	}
	return str
}

func (a ArrayType) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Array: \n"
	for _, el := range a.Subvalues{
		str += el.Print(indent+1)
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
	return i.StringValue
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
        for _, el := range(a.Arguments){
            str += el.Print(indent + 1)
        }
        for i := 0; i < indent; i++ {
            str += "  "
        }
        str += ")"
    }
    str += "\n"
    for _, el := range(a.Subvalues){
        str += el.Print(indent+1)
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
        str += a.Module.Print(indent + 1) + "."
    }
    str += a.Function.Print(indent + 1)

    if len(a.Arguments) > 0 {
        for i := 0; i < indent; i++ {
            str += "  "
        }
        str += "(\n"
        for _, el := range(a.Arguments){
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
        for _, el := range(a.Arguments){
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
        for key, el := range(a.Values){
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
    for _, el := range(i.Then){
        for i := 0; i < indent; i++ {
            str += "  "
        }
        str += "Then:\n"
        str += el.Print(indent+1)
    }
    for _, el := range(i.Else){
        for i := 0; i < indent; i++ {
            str += "  "
        }
        str += "Else:\n"
        str += el.Print(indent+1)

    }
    return str
}

func (a Assignment) Print(indent int) string {
    str := ""

    for i := 0; i < indent; i++ {
        str += "  "
    }
    str += "Assignment:\n"
	str += a.Left.Print(indent+1)

    str += a.Right.Print(indent+1)

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

