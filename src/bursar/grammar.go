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

type Assignment struct {
	Left  []Ast
	Right Ast
}

type Ast interface {
	Print(indent int) string
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
	return "Func"
}

func (i If) String() string {
	return "If"
}

func (a Func) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Func"
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

func (a Assignment) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Assignment:\n"

	for _, el := range a.Left {
		str += el.Print(indent + 1)
	}
	str += a.Right.Print(indent + 1)

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
			pos:  position{line: 161, col: 1, offset: 2780},
			expr: &actionExpr{
				pos: position{line: 161, col: 10, offset: 2789},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 161, col: 10, offset: 2789},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 161, col: 10, offset: 2789},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 12, offset: 2791},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 17, offset: 2796},
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 27, offset: 2806},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 29, offset: 2808},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 161, col: 34, offset: 2813},
								expr: &ruleRefExpr{
									pos:  position{line: 161, col: 35, offset: 2814},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 47, offset: 2826},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 49, offset: 2828},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 176, col: 1, offset: 3312},
			expr: &choiceExpr{
				pos: position{line: 176, col: 13, offset: 3324},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 176, col: 13, offset: 3324},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 176, col: 13, offset: 3324},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 176, col: 13, offset: 3324},
									val:        "#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 176, col: 17, offset: 3328},
									label: "comment",
									expr: &zeroOrMoreExpr{
										pos: position{line: 176, col: 25, offset: 3336},
										expr: &seqExpr{
											pos: position{line: 176, col: 26, offset: 3337},
											exprs: []interface{}{
												&notExpr{
													pos: position{line: 176, col: 26, offset: 3337},
													expr: &ruleRefExpr{
														pos:  position{line: 176, col: 27, offset: 3338},
														name: "EscapedChar",
													},
												},
												&anyMatcher{
													line: 176, col: 39, offset: 3350,
												},
											},
										},
									},
								},
								&andExpr{
									pos: position{line: 176, col: 43, offset: 3354},
									expr: &litMatcher{
										pos:        position{line: 176, col: 44, offset: 3355},
										val:        "\n",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 181, col: 1, offset: 3504},
						run: (*parser).callonStatement13,
						expr: &seqExpr{
							pos: position{line: 181, col: 1, offset: 3504},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 181, col: 1, offset: 3504},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 181, col: 3, offset: 3506},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 9, offset: 3512},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 181, col: 12, offset: 3515},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 14, offset: 3517},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 181, col: 25, offset: 3528},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 181, col: 30, offset: 3533},
										expr: &seqExpr{
											pos: position{line: 181, col: 31, offset: 3534},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 181, col: 31, offset: 3534},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 181, col: 35, offset: 3538},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 181, col: 37, offset: 3540},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 50, offset: 3553},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 181, col: 52, offset: 3555},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 56, offset: 3559},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 181, col: 58, offset: 3561},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 63, offset: 3566},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 68, offset: 3571},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 194, col: 5, offset: 4008},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 196, col: 1, offset: 4014},
			expr: &actionExpr{
				pos: position{line: 196, col: 8, offset: 4021},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 196, col: 8, offset: 4021},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 196, col: 12, offset: 4025},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 196, col: 12, offset: 4025},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 196, col: 23, offset: 4036},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 196, col: 32, offset: 4045},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 201, col: 1, offset: 4137},
			expr: &choiceExpr{
				pos: position{line: 201, col: 10, offset: 4146},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 201, col: 10, offset: 4146},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 201, col: 10, offset: 4146},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 201, col: 10, offset: 4146},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 15, offset: 4151},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 201, col: 18, offset: 4154},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 201, col: 23, offset: 4159},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 33, offset: 4169},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 201, col: 35, offset: 4171},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 39, offset: 4175},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 201, col: 41, offset: 4177},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 201, col: 47, offset: 4183},
										expr: &ruleRefExpr{
											pos:  position{line: 201, col: 48, offset: 4184},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 60, offset: 4196},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 201, col: 62, offset: 4198},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 66, offset: 4202},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 201, col: 68, offset: 4204},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 75, offset: 4211},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 201, col: 77, offset: 4213},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 201, col: 85, offset: 4221},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 213, col: 1, offset: 4549},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 213, col: 1, offset: 4549},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 213, col: 1, offset: 4549},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 6, offset: 4554},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 9, offset: 4557},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 213, col: 14, offset: 4562},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 24, offset: 4572},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 26, offset: 4574},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 30, offset: 4578},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 32, offset: 4580},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 213, col: 38, offset: 4586},
										expr: &ruleRefExpr{
											pos:  position{line: 213, col: 39, offset: 4587},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 51, offset: 4599},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 53, offset: 4601},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 57, offset: 4605},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 59, offset: 4607},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 66, offset: 4614},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 68, offset: 4616},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 72, offset: 4620},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 74, offset: 4622},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 213, col: 80, offset: 4628},
										expr: &ruleRefExpr{
											pos:  position{line: 213, col: 81, offset: 4629},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 93, offset: 4641},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 95, offset: 4643},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 99, offset: 4647},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 232, col: 1, offset: 5146},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 232, col: 1, offset: 5146},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 232, col: 1, offset: 5146},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 6, offset: 5151},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 232, col: 9, offset: 5154},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 14, offset: 5159},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 24, offset: 5169},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 232, col: 26, offset: 5171},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 30, offset: 5175},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 232, col: 32, offset: 5177},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 232, col: 38, offset: 5183},
										expr: &ruleRefExpr{
											pos:  position{line: 232, col: 39, offset: 5184},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 232, col: 51, offset: 5196},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 55, offset: 5200},
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
			pos:  position{line: 244, col: 1, offset: 5494},
			expr: &actionExpr{
				pos: position{line: 244, col: 12, offset: 5505},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 244, col: 12, offset: 5505},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 244, col: 12, offset: 5505},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 19, offset: 5512},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 244, col: 22, offset: 5515},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 244, col: 26, offset: 5519},
								expr: &seqExpr{
									pos: position{line: 244, col: 27, offset: 5520},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 244, col: 27, offset: 5520},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 244, col: 40, offset: 5533},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 45, offset: 5538},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 244, col: 47, offset: 5540},
							expr: &litMatcher{
								pos:        position{line: 244, col: 47, offset: 5540},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 53, offset: 5546},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 244, col: 55, offset: 5548},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 59, offset: 5552},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 244, col: 61, offset: 5554},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 244, col: 72, offset: 5565},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 73, offset: 5566},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 85, offset: 5578},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 244, col: 87, offset: 5580},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 91, offset: 5584},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 267, col: 1, offset: 6225},
			expr: &actionExpr{
				pos: position{line: 267, col: 16, offset: 6240},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 267, col: 16, offset: 6240},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 267, col: 16, offset: 6240},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 267, col: 18, offset: 6242},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 21, offset: 6245},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 267, col: 27, offset: 6251},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 267, col: 32, offset: 6256},
								expr: &seqExpr{
									pos: position{line: 267, col: 33, offset: 6257},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 267, col: 33, offset: 6257},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 267, col: 36, offset: 6260},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 267, col: 45, offset: 6269},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 267, col: 48, offset: 6272},
											name: "BinOp",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 267, col: 56, offset: 6280},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 267, col: 58, offset: 6282},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 267, col: 62, offset: 6286},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 287, col: 1, offset: 6945},
			expr: &choiceExpr{
				pos: position{line: 287, col: 9, offset: 6953},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 287, col: 9, offset: 6953},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 287, col: 21, offset: 6965},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 287, col: 37, offset: 6981},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 287, col: 48, offset: 6992},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 287, col: 60, offset: 7004},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 289, col: 1, offset: 7017},
			expr: &actionExpr{
				pos: position{line: 289, col: 13, offset: 7029},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 289, col: 13, offset: 7029},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 289, col: 13, offset: 7029},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 15, offset: 7031},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 21, offset: 7037},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 289, col: 35, offset: 7051},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 289, col: 40, offset: 7056},
								expr: &seqExpr{
									pos: position{line: 289, col: 41, offset: 7057},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 289, col: 41, offset: 7057},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 44, offset: 7060},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 60, offset: 7076},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 63, offset: 7079},
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
			pos:  position{line: 308, col: 1, offset: 7686},
			expr: &actionExpr{
				pos: position{line: 308, col: 17, offset: 7702},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 308, col: 17, offset: 7702},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 308, col: 17, offset: 7702},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 308, col: 19, offset: 7704},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 25, offset: 7710},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 308, col: 34, offset: 7719},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 308, col: 39, offset: 7724},
								expr: &seqExpr{
									pos: position{line: 308, col: 40, offset: 7725},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 308, col: 40, offset: 7725},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 43, offset: 7728},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 60, offset: 7745},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 63, offset: 7748},
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
			pos:  position{line: 328, col: 1, offset: 8353},
			expr: &actionExpr{
				pos: position{line: 328, col: 12, offset: 8364},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 328, col: 12, offset: 8364},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 328, col: 12, offset: 8364},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 328, col: 14, offset: 8366},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 328, col: 20, offset: 8372},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 328, col: 30, offset: 8382},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 328, col: 35, offset: 8387},
								expr: &seqExpr{
									pos: position{line: 328, col: 36, offset: 8388},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 328, col: 36, offset: 8388},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 328, col: 39, offset: 8391},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 328, col: 51, offset: 8403},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 328, col: 54, offset: 8406},
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
			pos:  position{line: 348, col: 1, offset: 9008},
			expr: &actionExpr{
				pos: position{line: 348, col: 13, offset: 9020},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 348, col: 13, offset: 9020},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 348, col: 13, offset: 9020},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 348, col: 15, offset: 9022},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 21, offset: 9028},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 348, col: 33, offset: 9040},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 348, col: 38, offset: 9045},
								expr: &seqExpr{
									pos: position{line: 348, col: 39, offset: 9046},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 348, col: 39, offset: 9046},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 42, offset: 9049},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 55, offset: 9062},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 58, offset: 9065},
											name: "BinOpParens",
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
			name: "BinOpParens",
			pos:  position{line: 367, col: 1, offset: 9670},
			expr: &choiceExpr{
				pos: position{line: 367, col: 15, offset: 9684},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 367, col: 15, offset: 9684},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 367, col: 15, offset: 9684},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 367, col: 15, offset: 9684},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 19, offset: 9688},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 21, offset: 9690},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 367, col: 27, offset: 9696},
										name: "BinOp",
									},
								},
								&litMatcher{
									pos:        position{line: 367, col: 33, offset: 9702},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 370, col: 5, offset: 9852},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 372, col: 1, offset: 9859},
			expr: &choiceExpr{
				pos: position{line: 372, col: 12, offset: 9870},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 372, col: 12, offset: 9870},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 372, col: 30, offset: 9888},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 372, col: 49, offset: 9907},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 372, col: 64, offset: 9922},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 374, col: 1, offset: 9935},
			expr: &actionExpr{
				pos: position{line: 374, col: 19, offset: 9953},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 374, col: 21, offset: 9955},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 374, col: 21, offset: 9955},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 374, col: 29, offset: 9963},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 374, col: 36, offset: 9970},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 378, col: 1, offset: 10069},
			expr: &actionExpr{
				pos: position{line: 378, col: 20, offset: 10088},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 378, col: 22, offset: 10090},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 378, col: 22, offset: 10090},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 378, col: 29, offset: 10097},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 378, col: 36, offset: 10104},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 378, col: 42, offset: 10110},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 378, col: 48, offset: 10116},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 378, col: 56, offset: 10124},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 382, col: 1, offset: 10230},
			expr: &choiceExpr{
				pos: position{line: 382, col: 16, offset: 10245},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 382, col: 16, offset: 10245},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 382, col: 18, offset: 10247},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 382, col: 18, offset: 10247},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 382, col: 25, offset: 10254},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 385, col: 3, offset: 10360},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 385, col: 5, offset: 10362},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 385, col: 5, offset: 10362},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 385, col: 11, offset: 10368},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 385, col: 17, offset: 10374},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 388, col: 3, offset: 10477},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 388, col: 3, offset: 10477},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 392, col: 1, offset: 10581},
			expr: &choiceExpr{
				pos: position{line: 392, col: 15, offset: 10595},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 392, col: 15, offset: 10595},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 392, col: 17, offset: 10597},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 392, col: 17, offset: 10597},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 392, col: 24, offset: 10604},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 395, col: 3, offset: 10710},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 395, col: 5, offset: 10712},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 395, col: 5, offset: 10712},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 395, col: 11, offset: 10718},
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
			pos:  position{line: 399, col: 1, offset: 10820},
			expr: &choiceExpr{
				pos: position{line: 399, col: 9, offset: 10828},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 399, col: 9, offset: 10828},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 399, col: 22, offset: 10841},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 399, col: 22, offset: 10841},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 399, col: 24, offset: 10843},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 403, col: 1, offset: 10878},
			expr: &choiceExpr{
				pos: position{line: 403, col: 14, offset: 10891},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 403, col: 14, offset: 10891},
						name: "Identifier",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 27, offset: 10904},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 405, col: 1, offset: 10912},
			expr: &choiceExpr{
				pos: position{line: 405, col: 14, offset: 10925},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 405, col: 14, offset: 10925},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 405, col: 29, offset: 10940},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 407, col: 1, offset: 10952},
			expr: &actionExpr{
				pos: position{line: 407, col: 16, offset: 10967},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 407, col: 17, offset: 10968},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 407, col: 17, offset: 10968},
							expr: &charClassMatcher{
								pos:        position{line: 407, col: 17, offset: 10968},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 407, col: 24, offset: 10975},
							expr: &charClassMatcher{
								pos:        position{line: 407, col: 24, offset: 10975},
								val:        "[a-zA-Z0-9_]",
								chars:      []rune{'_'},
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "ModuleName",
			pos:  position{line: 411, col: 1, offset: 11086},
			expr: &actionExpr{
				pos: position{line: 411, col: 14, offset: 11099},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 411, col: 15, offset: 11100},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 411, col: 15, offset: 11100},
							expr: &charClassMatcher{
								pos:        position{line: 411, col: 15, offset: 11100},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 411, col: 22, offset: 11107},
							expr: &charClassMatcher{
								pos:        position{line: 411, col: 22, offset: 11107},
								val:        "[a-zA-Z0-9_]",
								chars:      []rune{'_'},
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 415, col: 1, offset: 11218},
			expr: &choiceExpr{
				pos: position{line: 415, col: 9, offset: 11226},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 415, col: 9, offset: 11226},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 415, col: 9, offset: 11226},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 415, col: 9, offset: 11226},
									expr: &litMatcher{
										pos:        position{line: 415, col: 9, offset: 11226},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 415, col: 14, offset: 11231},
									expr: &charClassMatcher{
										pos:        position{line: 415, col: 14, offset: 11231},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 415, col: 21, offset: 11238},
									expr: &litMatcher{
										pos:        position{line: 415, col: 22, offset: 11239},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 422, col: 3, offset: 11415},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 422, col: 3, offset: 11415},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 422, col: 3, offset: 11415},
									expr: &litMatcher{
										pos:        position{line: 422, col: 3, offset: 11415},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 422, col: 8, offset: 11420},
									expr: &charClassMatcher{
										pos:        position{line: 422, col: 8, offset: 11420},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 15, offset: 11427},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 422, col: 19, offset: 11431},
									expr: &charClassMatcher{
										pos:        position{line: 422, col: 19, offset: 11431},
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
						pos:        position{line: 429, col: 3, offset: 11621},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 429, col: 12, offset: 11630},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 429, col: 12, offset: 11630},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 435, col: 3, offset: 11831},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 435, col: 3, offset: 11831},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 438, col: 3, offset: 11894},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 438, col: 3, offset: 11894},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 438, col: 3, offset: 11894},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 438, col: 7, offset: 11898},
									expr: &seqExpr{
										pos: position{line: 438, col: 8, offset: 11899},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 438, col: 8, offset: 11899},
												expr: &ruleRefExpr{
													pos:  position{line: 438, col: 9, offset: 11900},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 438, col: 21, offset: 11912,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 25, offset: 11916},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 445, col: 3, offset: 12100},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 445, col: 3, offset: 12100},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 445, col: 3, offset: 12100},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 445, col: 7, offset: 12104},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 445, col: 12, offset: 12109},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 445, col: 12, offset: 12109},
												expr: &ruleRefExpr{
													pos:  position{line: 445, col: 13, offset: 12110},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 445, col: 25, offset: 12122,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 445, col: 28, offset: 12125},
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
			name: "Unused",
			pos:  position{line: 449, col: 1, offset: 12216},
			expr: &actionExpr{
				pos: position{line: 449, col: 10, offset: 12225},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 449, col: 11, offset: 12226},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 453, col: 1, offset: 12327},
			expr: &charClassMatcher{
				pos:        position{line: 453, col: 15, offset: 12341},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 455, col: 1, offset: 12357},
			expr: &choiceExpr{
				pos: position{line: 455, col: 18, offset: 12374},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 455, col: 18, offset: 12374},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 455, col: 37, offset: 12393},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 457, col: 1, offset: 12408},
			expr: &charClassMatcher{
				pos:        position{line: 457, col: 20, offset: 12427},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 459, col: 1, offset: 12440},
			expr: &charClassMatcher{
				pos:        position{line: 459, col: 16, offset: 12455},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 461, col: 1, offset: 12462},
			expr: &charClassMatcher{
				pos:        position{line: 461, col: 23, offset: 12484},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 463, col: 1, offset: 12491},
			expr: &charClassMatcher{
				pos:        position{line: 463, col: 12, offset: 12502},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 465, col: 1, offset: 12513},
			expr: &oneOrMoreExpr{
				pos: position{line: 465, col: 22, offset: 12534},
				expr: &charClassMatcher{
					pos:        position{line: 465, col: 22, offset: 12534},
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
			pos:         position{line: 467, col: 1, offset: 12546},
			expr: &zeroOrMoreExpr{
				pos: position{line: 467, col: 18, offset: 12563},
				expr: &charClassMatcher{
					pos:        position{line: 467, col: 18, offset: 12563},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 469, col: 1, offset: 12575},
			expr: &notExpr{
				pos: position{line: 469, col: 7, offset: 12581},
				expr: &anyMatcher{
					line: 469, col: 8, offset: 12582,
				},
			},
		},
	},
}

func (c *current) onModule1(stat, rest interface{}) (interface{}, error) {
	fmt.Println("beginning module")
	vals := rest.([]interface{})
	if len(vals) > 0 {
		fmt.Println("multiple statements")
		subvalues := []Ast{stat.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return BasicAst{Type: "Module", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return BasicAst{Type: "Module", Subvalues: []Ast{stat.(Ast)}, ValueType: CONTAINER}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["stat"], stack["rest"])
}

func (c *current) onStatement2(comment interface{}) (interface{}, error) {
	fmt.Println("comment:", string(c.text))
	return BasicAst{Type: "Comment", StringValue: string(c.text[1:]), ValueType: STRING}, nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["comment"])
}

func (c *current) onStatement13(i, rest, expr interface{}) (interface{}, error) {
	fmt.Println("assignment:", string(c.text))
	vals := []Ast{i.(Ast)}
	if len(rest.([]interface{})) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			vals = append(vals, v)
		}
	}
	return Assignment{Left: vals, Right: expr.(Ast)}, nil
}

func (p *parser) callonStatement13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement13(stack["i"], stack["rest"], stack["expr"])
}

func (c *current) onExpr1(ex interface{}) (interface{}, error) {
	fmt.Printf("top-level expr: %s\n", string(c.text))
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
	fmt.Println(string(c.text))
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

func (c *current) onCompoundExpr1(op, rest interface{}) (interface{}, error) {
	//fmt.Println("compound", op, rest);
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{op.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}

		return BasicAst{Type: "CompoundExpr", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return BasicAst{Type: "CompoundExpr", Subvalues: []Ast{op.(Ast)}, ValueType: CONTAINER}, nil
	}
}

func (p *parser) callonCompoundExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCompoundExpr1(stack["op"], stack["rest"])
}

func (c *current) onBinOpBool1(first, rest interface{}) (interface{}, error) {
	//fmt.Println("binopbool", first, rest);
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
	//fmt.Println("binopeq", first, rest);
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
	//fmt.Println("binoplow", first, rest);
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
	//fmt.Println("binophigh", first, rest);
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

func (c *current) onBinOpParens2(first interface{}) (interface{}, error) {
	//fmt.Println("binopparens", first);
	return BasicAst{Type: "BinOpParens", Subvalues: []Ast{first.(Ast)}, ValueType: CONTAINER}, nil
}

func (p *parser) callonBinOpParens2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpParens2(stack["first"])
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

func (c *current) onValue3(v interface{}) (interface{}, error) {
	return v.(Ast), nil
}

func (p *parser) callonValue3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue3(stack["v"])
}

func (c *current) onVariableName1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonVariableName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableName1()
}

func (c *current) onModuleName1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonModuleName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModuleName1()
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

func (c *current) onUnused1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonUnused1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnused1()
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
