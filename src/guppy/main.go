package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type ValueType int

const (
	STRING = iota
	INT
	FLOAT
	BOOL
	CHAR
	CONTAINER
	NIL
)

type BasicAst struct {
	Type        string
	StringValue string
	CharValue   rune
	BoolValue   bool
	IntValue    int
	FloatValue  float64
	ValueType   ValueType
	Subvalues   []Ast
}

type Func struct {
	Arguments []Ast
	ValueType ValueType
	Subvalues []Ast
}

type If struct {
	Condition Ast
	Then      []Ast
	Else      []Ast
}

type Ast interface {
	Print(indent int) string
}

func main() {
	input := `

let a = 5.0 + 6.5;
let adder = fun num {
    if num > 0 and num < 10 {
        20 + num;
    } else {
        5 + num;
    }
}
let cheesy = fun item {
    item ++ " with cheese";
}
let tester = fun a {
    let result = if a > 100 {
        a + 1;
    } else if a > 50 {
        a + 5;
    } else if a > 20 {
        a + 10;
    } else {
        a + 20;
    }
    result;
}
`
	fmt.Println(input)
	r := strings.NewReader(input)
	result, err := ParseReader("", r)
	ast := result.(Ast)
	fmt.Println("=", ast.Print(0))
	fmt.Println(err)
}

func (a BasicAst) String() string {
	switch a.ValueType {
	case STRING:
		return fmt.Sprintf("\"%s\"", a.StringValue)
	case CHAR:
		return fmt.Sprintf("'%s'", string(a.CharValue))
	case INT:
		return fmt.Sprintf("%d", a.IntValue)
	case FLOAT:
		return fmt.Sprintf("%f", a.FloatValue)
	}
	return "()"
}

func (a BasicAst) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += fmt.Sprintf("%s %s:\n", a.Type, a)
	for _, el := range a.Subvalues {
		str += el.Print(indent + 1)
	}
	return str
}

func (a Func) String() string {
	return "Fun"
}

func (i If) String() string {
	return "If"
}

func (a Func) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Fun"
	if len(a.Arguments) > 0 {
		str += " (\n"
		for _, el := range a.Arguments {
			str += el.Print(indent + 1)
		}
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += ")\n"
	}
	for _, el := range a.Subvalues {
		str += el.Print(indent + 1)
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
	for _, el := range i.Then {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "Then:\n"
		str += el.Print(indent + 1)
	}
	for _, el := range i.Else {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "Else:\n"
		str += el.Print(indent + 1)

	}
	return str
}

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Module",
			pos:  position{line: 175, col: 1, offset: 3027},
			expr: &actionExpr{
				pos: position{line: 175, col: 10, offset: 3036},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 175, col: 10, offset: 3036},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 175, col: 10, offset: 3036},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 175, col: 12, offset: 3038},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 175, col: 17, offset: 3043},
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 175, col: 27, offset: 3053},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 175, col: 29, offset: 3055},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 175, col: 34, offset: 3060},
								expr: &ruleRefExpr{
									pos:  position{line: 175, col: 35, offset: 3061},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 175, col: 47, offset: 3073},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 175, col: 49, offset: 3075},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 190, col: 1, offset: 3559},
			expr: &choiceExpr{
				pos: position{line: 190, col: 13, offset: 3571},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 190, col: 13, offset: 3571},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 190, col: 13, offset: 3571},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 190, col: 13, offset: 3571},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 19, offset: 3577},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 190, col: 22, offset: 3580},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 190, col: 24, offset: 3582},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 35, offset: 3593},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 190, col: 37, offset: 3595},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 41, offset: 3599},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 190, col: 43, offset: 3601},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 190, col: 48, offset: 3606},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 53, offset: 3611},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 5, offset: 3776},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 195, col: 1, offset: 3782},
			expr: &actionExpr{
				pos: position{line: 195, col: 8, offset: 3789},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 195, col: 8, offset: 3789},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 195, col: 12, offset: 3793},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 195, col: 12, offset: 3793},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 195, col: 23, offset: 3804},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 195, col: 32, offset: 3813},
								name: "BinExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 200, col: 1, offset: 3890},
			expr: &choiceExpr{
				pos: position{line: 200, col: 10, offset: 3899},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 200, col: 10, offset: 3899},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 200, col: 10, offset: 3899},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 200, col: 10, offset: 3899},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 15, offset: 3904},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 18, offset: 3907},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 23, offset: 3912},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 33, offset: 3922},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 35, offset: 3924},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 39, offset: 3928},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 41, offset: 3930},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 47, offset: 3936},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 48, offset: 3937},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 60, offset: 3949},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 62, offset: 3951},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 66, offset: 3955},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 68, offset: 3957},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 75, offset: 3964},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 77, offset: 3966},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 85, offset: 3974},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 212, col: 1, offset: 4302},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 212, col: 1, offset: 4302},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 212, col: 1, offset: 4302},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 6, offset: 4307},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 9, offset: 4310},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 14, offset: 4315},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 24, offset: 4325},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 26, offset: 4327},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 30, offset: 4331},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 32, offset: 4333},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 38, offset: 4339},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 39, offset: 4340},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 51, offset: 4352},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 53, offset: 4354},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 57, offset: 4358},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 59, offset: 4360},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 66, offset: 4367},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 68, offset: 4369},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 72, offset: 4373},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 74, offset: 4375},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 80, offset: 4381},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 81, offset: 4382},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 93, offset: 4394},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 95, offset: 4396},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 99, offset: 4400},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 231, col: 1, offset: 4899},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 231, col: 1, offset: 4899},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 231, col: 1, offset: 4899},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 6, offset: 4904},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 9, offset: 4907},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 14, offset: 4912},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 24, offset: 4922},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 231, col: 26, offset: 4924},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 30, offset: 4928},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 32, offset: 4930},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 231, col: 38, offset: 4936},
										expr: &ruleRefExpr{
											pos:  position{line: 231, col: 39, offset: 4937},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 231, col: 51, offset: 4949},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 55, offset: 4953},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FuncDefn",
			pos:  position{line: 243, col: 1, offset: 5247},
			expr: &actionExpr{
				pos: position{line: 243, col: 12, offset: 5258},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 243, col: 12, offset: 5258},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 243, col: 12, offset: 5258},
							val:        "fun",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 18, offset: 5264},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 243, col: 21, offset: 5267},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 243, col: 25, offset: 5271},
								expr: &seqExpr{
									pos: position{line: 243, col: 26, offset: 5272},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 243, col: 26, offset: 5272},
											name: "Identifier",
										},
										&ruleRefExpr{
											pos:  position{line: 243, col: 37, offset: 5283},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 42, offset: 5288},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 243, col: 44, offset: 5290},
							expr: &litMatcher{
								pos:        position{line: 243, col: 44, offset: 5290},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 50, offset: 5296},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 243, col: 52, offset: 5298},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 56, offset: 5302},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 243, col: 58, offset: 5304},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 243, col: 69, offset: 5315},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 70, offset: 5316},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 82, offset: 5328},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 243, col: 84, offset: 5330},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 88, offset: 5334},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinExpr",
			pos:  position{line: 266, col: 1, offset: 5983},
			expr: &actionExpr{
				pos: position{line: 266, col: 11, offset: 5993},
				run: (*parser).callonBinExpr1,
				expr: &seqExpr{
					pos: position{line: 266, col: 11, offset: 5993},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 266, col: 11, offset: 5993},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 266, col: 13, offset: 5995},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 16, offset: 5998},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 266, col: 22, offset: 6004},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 266, col: 27, offset: 6009},
								expr: &ruleRefExpr{
									pos:  position{line: 266, col: 28, offset: 6010},
									name: "BinOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 36, offset: 6018},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 266, col: 38, offset: 6020},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 42, offset: 6024},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 279, col: 1, offset: 6425},
			expr: &choiceExpr{
				pos: position{line: 279, col: 9, offset: 6433},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 279, col: 9, offset: 6433},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 279, col: 21, offset: 6445},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 279, col: 37, offset: 6461},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 279, col: 48, offset: 6472},
						name: "BinOpHigh",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 281, col: 1, offset: 6483},
			expr: &actionExpr{
				pos: position{line: 281, col: 13, offset: 6495},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 281, col: 13, offset: 6495},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 281, col: 13, offset: 6495},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 281, col: 15, offset: 6497},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 281, col: 21, offset: 6503},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 281, col: 35, offset: 6517},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 281, col: 40, offset: 6522},
								expr: &seqExpr{
									pos: position{line: 281, col: 41, offset: 6523},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 281, col: 41, offset: 6523},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 281, col: 44, offset: 6526},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 281, col: 60, offset: 6542},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 281, col: 63, offset: 6545},
											name: "BinOpEquality",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpEquality",
			pos:  position{line: 299, col: 1, offset: 7107},
			expr: &actionExpr{
				pos: position{line: 299, col: 17, offset: 7123},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 299, col: 17, offset: 7123},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 299, col: 17, offset: 7123},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 299, col: 19, offset: 7125},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 299, col: 25, offset: 7131},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 299, col: 34, offset: 7140},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 299, col: 39, offset: 7145},
								expr: &seqExpr{
									pos: position{line: 299, col: 40, offset: 7146},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 299, col: 40, offset: 7146},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 299, col: 43, offset: 7149},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 299, col: 60, offset: 7166},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 299, col: 63, offset: 7169},
											name: "BinOpLow",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpLow",
			pos:  position{line: 318, col: 1, offset: 7731},
			expr: &actionExpr{
				pos: position{line: 318, col: 12, offset: 7742},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 318, col: 12, offset: 7742},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 318, col: 12, offset: 7742},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 318, col: 14, offset: 7744},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 318, col: 20, offset: 7750},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 318, col: 30, offset: 7760},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 318, col: 35, offset: 7765},
								expr: &seqExpr{
									pos: position{line: 318, col: 36, offset: 7766},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 318, col: 36, offset: 7766},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 318, col: 39, offset: 7769},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 318, col: 51, offset: 7781},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 318, col: 54, offset: 7784},
											name: "BinOpHigh",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpHigh",
			pos:  position{line: 337, col: 1, offset: 8342},
			expr: &actionExpr{
				pos: position{line: 337, col: 13, offset: 8354},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 337, col: 13, offset: 8354},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 337, col: 13, offset: 8354},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 337, col: 15, offset: 8356},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 337, col: 21, offset: 8362},
								name: "Value",
							},
						},
						&labeledExpr{
							pos:   position{line: 337, col: 27, offset: 8368},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 337, col: 32, offset: 8373},
								expr: &seqExpr{
									pos: position{line: 337, col: 33, offset: 8374},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 337, col: 33, offset: 8374},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 337, col: 36, offset: 8377},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 337, col: 49, offset: 8390},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 337, col: 52, offset: 8393},
											name: "Value",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 355, col: 1, offset: 8947},
			expr: &choiceExpr{
				pos: position{line: 355, col: 12, offset: 8958},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 355, col: 12, offset: 8958},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 355, col: 30, offset: 8976},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 355, col: 49, offset: 8995},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 355, col: 64, offset: 9010},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 357, col: 1, offset: 9023},
			expr: &actionExpr{
				pos: position{line: 357, col: 19, offset: 9041},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 357, col: 21, offset: 9043},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 357, col: 21, offset: 9043},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 357, col: 29, offset: 9051},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 357, col: 36, offset: 9058},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 361, col: 1, offset: 9157},
			expr: &actionExpr{
				pos: position{line: 361, col: 20, offset: 9176},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 361, col: 22, offset: 9178},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 361, col: 22, offset: 9178},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 361, col: 29, offset: 9185},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 361, col: 36, offset: 9192},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 361, col: 42, offset: 9198},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 361, col: 48, offset: 9204},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 361, col: 56, offset: 9212},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 365, col: 1, offset: 9318},
			expr: &choiceExpr{
				pos: position{line: 365, col: 16, offset: 9333},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 365, col: 16, offset: 9333},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 365, col: 18, offset: 9335},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 365, col: 18, offset: 9335},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 365, col: 25, offset: 9342},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 368, col: 3, offset: 9448},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 368, col: 5, offset: 9450},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 368, col: 5, offset: 9450},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 368, col: 11, offset: 9456},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 368, col: 17, offset: 9462},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 371, col: 3, offset: 9565},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 371, col: 3, offset: 9565},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 375, col: 1, offset: 9669},
			expr: &choiceExpr{
				pos: position{line: 375, col: 15, offset: 9683},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 375, col: 15, offset: 9683},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 375, col: 17, offset: 9685},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 375, col: 17, offset: 9685},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 375, col: 24, offset: 9692},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 378, col: 3, offset: 9798},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 378, col: 5, offset: 9800},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 378, col: 5, offset: 9800},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 378, col: 11, offset: 9806},
									val:        "-",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 382, col: 1, offset: 9908},
			expr: &choiceExpr{
				pos: position{line: 382, col: 9, offset: 9916},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 382, col: 9, offset: 9916},
						run: (*parser).callonValue2,
						expr: &seqExpr{
							pos: position{line: 382, col: 9, offset: 9916},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 382, col: 9, offset: 9916},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 382, col: 13, offset: 9920},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 382, col: 18, offset: 9925},
										name: "Expr",
									},
								},
								&litMatcher{
									pos:        position{line: 382, col: 23, offset: 9930},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 385, col: 3, offset: 9967},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 385, col: 16, offset: 9980},
						run: (*parser).callonValue9,
						expr: &labeledExpr{
							pos:   position{line: 385, col: 16, offset: 9980},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 385, col: 18, offset: 9982},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 389, col: 1, offset: 10017},
			expr: &actionExpr{
				pos: position{line: 389, col: 14, offset: 10030},
				run: (*parser).callonIdentifier1,
				expr: &choiceExpr{
					pos: position{line: 389, col: 15, offset: 10031},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 389, col: 15, offset: 10031},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 389, col: 15, offset: 10031},
									expr: &charClassMatcher{
										pos:        position{line: 389, col: 15, offset: 10031},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 389, col: 22, offset: 10038},
									expr: &charClassMatcher{
										pos:        position{line: 389, col: 22, offset: 10038},
										val:        "[a-zA-Z0-9_]",
										chars:      []rune{'_'},
										ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 389, col: 38, offset: 10054},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 393, col: 1, offset: 10155},
			expr: &choiceExpr{
				pos: position{line: 393, col: 9, offset: 10163},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 393, col: 9, offset: 10163},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 393, col: 9, offset: 10163},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 393, col: 9, offset: 10163},
									expr: &litMatcher{
										pos:        position{line: 393, col: 9, offset: 10163},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 393, col: 14, offset: 10168},
									expr: &charClassMatcher{
										pos:        position{line: 393, col: 14, offset: 10168},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 393, col: 21, offset: 10175},
									expr: &litMatcher{
										pos:        position{line: 393, col: 22, offset: 10176},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 400, col: 3, offset: 10352},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 400, col: 3, offset: 10352},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 400, col: 3, offset: 10352},
									expr: &litMatcher{
										pos:        position{line: 400, col: 3, offset: 10352},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 400, col: 8, offset: 10357},
									expr: &charClassMatcher{
										pos:        position{line: 400, col: 8, offset: 10357},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 400, col: 15, offset: 10364},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 400, col: 19, offset: 10368},
									expr: &charClassMatcher{
										pos:        position{line: 400, col: 19, offset: 10368},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 407, col: 3, offset: 10558},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 407, col: 12, offset: 10567},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 407, col: 12, offset: 10567},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 413, col: 3, offset: 10768},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 413, col: 3, offset: 10768},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 416, col: 3, offset: 10831},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 416, col: 3, offset: 10831},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 416, col: 3, offset: 10831},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 416, col: 7, offset: 10835},
									expr: &seqExpr{
										pos: position{line: 416, col: 8, offset: 10836},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 416, col: 8, offset: 10836},
												expr: &ruleRefExpr{
													pos:  position{line: 416, col: 9, offset: 10837},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 416, col: 21, offset: 10849,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 25, offset: 10853},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 423, col: 3, offset: 11037},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 423, col: 3, offset: 11037},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 423, col: 3, offset: 11037},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 423, col: 7, offset: 11041},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 423, col: 12, offset: 11046},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 423, col: 12, offset: 11046},
												expr: &ruleRefExpr{
													pos:  position{line: 423, col: 13, offset: 11047},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 423, col: 25, offset: 11059,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 423, col: 28, offset: 11062},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 427, col: 1, offset: 11153},
			expr: &charClassMatcher{
				pos:        position{line: 427, col: 15, offset: 11167},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 429, col: 1, offset: 11183},
			expr: &choiceExpr{
				pos: position{line: 429, col: 18, offset: 11200},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 429, col: 18, offset: 11200},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 429, col: 37, offset: 11219},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 431, col: 1, offset: 11234},
			expr: &charClassMatcher{
				pos:        position{line: 431, col: 20, offset: 11253},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 433, col: 1, offset: 11266},
			expr: &charClassMatcher{
				pos:        position{line: 433, col: 16, offset: 11281},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 435, col: 1, offset: 11288},
			expr: &charClassMatcher{
				pos:        position{line: 435, col: 23, offset: 11310},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 437, col: 1, offset: 11317},
			expr: &charClassMatcher{
				pos:        position{line: 437, col: 12, offset: 11328},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 439, col: 1, offset: 11339},
			expr: &oneOrMoreExpr{
				pos: position{line: 439, col: 22, offset: 11360},
				expr: &charClassMatcher{
					pos:        position{line: 439, col: 22, offset: 11360},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 441, col: 1, offset: 11372},
			expr: &zeroOrMoreExpr{
				pos: position{line: 441, col: 18, offset: 11389},
				expr: &charClassMatcher{
					pos:        position{line: 441, col: 18, offset: 11389},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 443, col: 1, offset: 11401},
			expr: &notExpr{
				pos: position{line: 443, col: 7, offset: 11407},
				expr: &anyMatcher{
					line: 443, col: 8, offset: 11408,
				},
			},
		},
	},
}

func (c *current) onModule1(expr, rest interface{}) (interface{}, error) {
	fmt.Println("beginning module")
	vals := rest.([]interface{})
	if len(vals) > 0 {
		fmt.Println("multiple statements")
		subvalues := []Ast{expr.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return BasicAst{Type: "Module", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return BasicAst{Type: "Module", Subvalues: []Ast{expr.(Ast)}, ValueType: CONTAINER}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["expr"], stack["rest"])
}

func (c *current) onStatement2(i, expr interface{}) (interface{}, error) {
	fmt.Printf("assignment: %s\n", string(c.text))
	return BasicAst{Type: "Assignment", Subvalues: []Ast{i.(Ast), expr.(Ast)}, ValueType: CONTAINER}, nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["i"], stack["expr"])
}

func (c *current) onExpr1(ex interface{}) (interface{}, error) {
	fmt.Printf("expr: %s\n", string(c.text))
	return ex, nil
}

func (p *parser) callonExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr1(stack["ex"])
}

func (c *current) onIfExpr2(expr, thens, elseifs interface{}) (interface{}, error) {
	fmt.Printf("if: %s\n", string(c.text))
	subvalues := []Ast{}
	vals := thens.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	return If{Condition: expr.(Ast), Then: subvalues, Else: []Ast{elseifs.(Ast)}}, nil
}

func (p *parser) callonIfExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr2(stack["expr"], stack["thens"], stack["elseifs"])
}

func (c *current) onIfExpr21(expr, thens, elses interface{}) (interface{}, error) {
	fmt.Printf("if: %s\n", string(c.text))
	subvalues := []Ast{}
	vals := thens.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	elsevalues := []Ast{}
	vals = elses.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			elsevalues = append(elsevalues, el.(Ast))
		}
	}
	return If{Condition: expr.(Ast), Then: subvalues, Else: elsevalues}, nil
}

func (p *parser) callonIfExpr21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr21(stack["expr"], stack["thens"], stack["elses"])
}

func (c *current) onIfExpr46(expr, thens interface{}) (interface{}, error) {
	fmt.Printf("if: %s\n", string(c.text))
	subvalues := []Ast{}
	vals := thens.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	return If{Condition: expr.(Ast), Then: subvalues}, nil
}

func (p *parser) callonIfExpr46() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr46(stack["expr"], stack["thens"])
}

func (c *current) onFuncDefn1(ids, statements interface{}) (interface{}, error) {
	fmt.Println("func", string(c.text))
	subvalues := []Ast{}
	args := []Ast{}
	vals := statements.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	vals = ids.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(ids)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[0].(Ast)
			args = append(args, v)
		}
	}
	return Func{Arguments: args, Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["ids"], stack["statements"])
}

func (c *current) onBinExpr1(op, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{op.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return BasicAst{Type: "BinExpr", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return BasicAst{Type: "BinExpr", Subvalues: []Ast{op.(Ast)}, ValueType: CONTAINER}, nil
	}
}

func (p *parser) callonBinExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinExpr1(stack["op"], stack["rest"])
}

func (c *current) onBinOpBool1(first, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return BasicAst{Type: "BinOpBool", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return first.(Ast), nil
	}
}

func (p *parser) callonBinOpBool1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpBool1(stack["first"], stack["rest"])
}

func (c *current) onBinOpEquality1(first, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return BasicAst{Type: "BinOpEquality", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return first.(Ast), nil
	}

}

func (p *parser) callonBinOpEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpEquality1(stack["first"], stack["rest"])
}

func (c *current) onBinOpLow1(first, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return BasicAst{Type: "BinOpLow", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return first.(Ast), nil
	}

}

func (p *parser) callonBinOpLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpLow1(stack["first"], stack["rest"])
}

func (c *current) onBinOpHigh1(first, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return BasicAst{Type: "BinOpHigh", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return first.(Ast), nil
	}
}

func (p *parser) callonBinOpHigh1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpHigh1(stack["first"], stack["rest"])
}

func (c *current) onOperatorBoolean1() (interface{}, error) {
	return BasicAst{Type: "BoolOp", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorBoolean1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorBoolean1()
}

func (c *current) onOperatorEquality1() (interface{}, error) {
	return BasicAst{Type: "EqualityOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorEquality1()
}

func (c *current) onOperatorHigh2() (interface{}, error) {
	return BasicAst{Type: "FloatOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh2()
}

func (c *current) onOperatorHigh6() (interface{}, error) {
	return BasicAst{Type: "IntOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh6()
}

func (c *current) onOperatorHigh11() (interface{}, error) {
	return BasicAst{Type: "StringOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh11()
}

func (c *current) onOperatorLow2() (interface{}, error) {
	return BasicAst{Type: "FloatOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorLow2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow2()
}

func (c *current) onOperatorLow6() (interface{}, error) {
	return BasicAst{Type: "IntOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorLow6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow6()
}

func (c *current) onValue2(expr interface{}) (interface{}, error) {
	return expr.(Ast), nil
}

func (p *parser) callonValue2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue2(stack["expr"])
}

func (c *current) onValue9(v interface{}) (interface{}, error) {
	return v.(Ast), nil
}

func (p *parser) callonValue9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue9(stack["v"])
}

func (c *current) onIdentifier1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1()
}

func (c *current) onConst2() (interface{}, error) {
	val, err := strconv.Atoi(string(c.text))
	if err != nil {
		return nil, err
	}
	return BasicAst{Type: "Integer", IntValue: val, ValueType: INT}, nil
}

func (p *parser) callonConst2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst2()
}

func (c *current) onConst10() (interface{}, error) {
	val, err := strconv.ParseFloat(string(c.text), 64)
	if err != nil {
		return nil, err
	}
	return BasicAst{Type: "Float", FloatValue: val, ValueType: FLOAT}, nil
}

func (p *parser) callonConst10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst10()
}

func (c *current) onConst20() (interface{}, error) {
	if string(c.text) == "True" {
		return BasicAst{Type: "Bool", BoolValue: true, ValueType: BOOL}, nil
	}
	return BasicAst{Type: "Bool", BoolValue: false, ValueType: BOOL}, nil
}

func (p *parser) callonConst20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst20()
}

func (c *current) onConst22() (interface{}, error) {
	return BasicAst{Type: "Nil", ValueType: NIL}, nil
}

func (p *parser) callonConst22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst22()
}

func (c *current) onConst24() (interface{}, error) {
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonConst24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst24()
}

func (c *current) onConst33(val interface{}) (interface{}, error) {
	return BasicAst{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst33() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst33(stack["val"])
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
