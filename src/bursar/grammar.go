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

type Call struct {
	Module    Ast
	Function  Ast
	Arguments []Ast
	ValueType ValueType
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

func (a Call) Print(indent int) string {
	str := ""

	for i := 0; i < indent; i++ {
		str += "  "
	}
	str += "Call:\n"
	if a.Module.(BasicAst).Type != "" {
		str += a.Module.Print(indent + 1)
	}
	str += a.Function.Print(indent + 1)

	if len(a.Arguments) > 0 {
		for i := 0; i < indent; i++ {
			str += "  "
		}
		str += "(\n"
		for _, el := range a.Arguments {
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
			pos:  position{line: 196, col: 1, offset: 3475},
			expr: &actionExpr{
				pos: position{line: 196, col: 10, offset: 3484},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 196, col: 10, offset: 3484},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 196, col: 10, offset: 3484},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 196, col: 12, offset: 3486},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 17, offset: 3491},
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 27, offset: 3501},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 196, col: 29, offset: 3503},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 196, col: 34, offset: 3508},
								expr: &ruleRefExpr{
									pos:  position{line: 196, col: 35, offset: 3509},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 47, offset: 3521},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 49, offset: 3523},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 211, col: 1, offset: 4007},
			expr: &choiceExpr{
				pos: position{line: 211, col: 13, offset: 4019},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 211, col: 13, offset: 4019},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 211, col: 13, offset: 4019},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 211, col: 13, offset: 4019},
									val:        "#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 211, col: 17, offset: 4023},
									label: "comment",
									expr: &zeroOrMoreExpr{
										pos: position{line: 211, col: 25, offset: 4031},
										expr: &seqExpr{
											pos: position{line: 211, col: 26, offset: 4032},
											exprs: []interface{}{
												&notExpr{
													pos: position{line: 211, col: 26, offset: 4032},
													expr: &ruleRefExpr{
														pos:  position{line: 211, col: 27, offset: 4033},
														name: "EscapedChar",
													},
												},
												&anyMatcher{
													line: 211, col: 39, offset: 4045,
												},
											},
										},
									},
								},
								&andExpr{
									pos: position{line: 211, col: 43, offset: 4049},
									expr: &litMatcher{
										pos:        position{line: 211, col: 44, offset: 4050},
										val:        "\n",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 216, col: 1, offset: 4199},
						run: (*parser).callonStatement13,
						expr: &seqExpr{
							pos: position{line: 216, col: 1, offset: 4199},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 216, col: 1, offset: 4199},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 216, col: 3, offset: 4201},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 216, col: 9, offset: 4207},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 216, col: 12, offset: 4210},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 216, col: 14, offset: 4212},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 216, col: 25, offset: 4223},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 216, col: 30, offset: 4228},
										expr: &seqExpr{
											pos: position{line: 216, col: 31, offset: 4229},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 216, col: 31, offset: 4229},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 216, col: 35, offset: 4233},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 216, col: 37, offset: 4235},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 216, col: 50, offset: 4248},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 216, col: 52, offset: 4250},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 216, col: 56, offset: 4254},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 216, col: 58, offset: 4256},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 216, col: 63, offset: 4261},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 216, col: 68, offset: 4266},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 229, col: 5, offset: 4703},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 231, col: 1, offset: 4709},
			expr: &actionExpr{
				pos: position{line: 231, col: 8, offset: 4716},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 231, col: 8, offset: 4716},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 231, col: 12, offset: 4720},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 231, col: 12, offset: 4720},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 231, col: 23, offset: 4731},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 231, col: 30, offset: 4738},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 231, col: 39, offset: 4747},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 236, col: 1, offset: 4839},
			expr: &choiceExpr{
				pos: position{line: 236, col: 10, offset: 4848},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 236, col: 10, offset: 4848},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 236, col: 10, offset: 4848},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 236, col: 10, offset: 4848},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 15, offset: 4853},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 18, offset: 4856},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 23, offset: 4861},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 33, offset: 4871},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 236, col: 35, offset: 4873},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 39, offset: 4877},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 41, offset: 4879},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 236, col: 47, offset: 4885},
										expr: &ruleRefExpr{
											pos:  position{line: 236, col: 48, offset: 4886},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 60, offset: 4898},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 236, col: 62, offset: 4900},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 66, offset: 4904},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 236, col: 68, offset: 4906},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 75, offset: 4913},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 77, offset: 4915},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 85, offset: 4923},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 248, col: 1, offset: 5251},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 248, col: 1, offset: 5251},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 248, col: 1, offset: 5251},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 6, offset: 5256},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 248, col: 9, offset: 5259},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 248, col: 14, offset: 5264},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 24, offset: 5274},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 248, col: 26, offset: 5276},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 30, offset: 5280},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 248, col: 32, offset: 5282},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 248, col: 38, offset: 5288},
										expr: &ruleRefExpr{
											pos:  position{line: 248, col: 39, offset: 5289},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 51, offset: 5301},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 248, col: 53, offset: 5303},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 57, offset: 5307},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 248, col: 59, offset: 5309},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 66, offset: 5316},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 248, col: 68, offset: 5318},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 72, offset: 5322},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 248, col: 74, offset: 5324},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 248, col: 80, offset: 5330},
										expr: &ruleRefExpr{
											pos:  position{line: 248, col: 81, offset: 5331},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 93, offset: 5343},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 248, col: 95, offset: 5345},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 99, offset: 5349},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 267, col: 1, offset: 5848},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 267, col: 1, offset: 5848},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 267, col: 1, offset: 5848},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 6, offset: 5853},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 267, col: 9, offset: 5856},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 267, col: 14, offset: 5861},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 24, offset: 5871},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 267, col: 26, offset: 5873},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 30, offset: 5877},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 267, col: 32, offset: 5879},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 267, col: 38, offset: 5885},
										expr: &ruleRefExpr{
											pos:  position{line: 267, col: 39, offset: 5886},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 267, col: 51, offset: 5898},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 55, offset: 5902},
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
			pos:  position{line: 279, col: 1, offset: 6196},
			expr: &actionExpr{
				pos: position{line: 279, col: 12, offset: 6207},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 279, col: 12, offset: 6207},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 279, col: 12, offset: 6207},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 19, offset: 6214},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 279, col: 22, offset: 6217},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 279, col: 26, offset: 6221},
								expr: &seqExpr{
									pos: position{line: 279, col: 27, offset: 6222},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 279, col: 27, offset: 6222},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 279, col: 40, offset: 6235},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 45, offset: 6240},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 279, col: 47, offset: 6242},
							expr: &litMatcher{
								pos:        position{line: 279, col: 47, offset: 6242},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 53, offset: 6248},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 279, col: 55, offset: 6250},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 59, offset: 6254},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 279, col: 61, offset: 6256},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 279, col: 72, offset: 6267},
								expr: &ruleRefExpr{
									pos:  position{line: 279, col: 73, offset: 6268},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 85, offset: 6280},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 279, col: 87, offset: 6282},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 91, offset: 6286},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Call",
			pos:  position{line: 302, col: 1, offset: 6927},
			expr: &choiceExpr{
				pos: position{line: 302, col: 8, offset: 6934},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 302, col: 8, offset: 6934},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 302, col: 8, offset: 6934},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 302, col: 8, offset: 6934},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 15, offset: 6941},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 302, col: 26, offset: 6952},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 302, col: 30, offset: 6956},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 33, offset: 6959},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 302, col: 46, offset: 6972},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 302, col: 56, offset: 6982},
										expr: &seqExpr{
											pos: position{line: 302, col: 57, offset: 6983},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 302, col: 57, offset: 6983},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 302, col: 60, offset: 6986},
													name: "Value",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 302, col: 68, offset: 6994},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 302, col: 72, offset: 6998},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 319, col: 3, offset: 7494},
						run: (*parser).callonCall16,
						expr: &seqExpr{
							pos: position{line: 319, col: 3, offset: 7494},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 319, col: 3, offset: 7494},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 319, col: 6, offset: 7497},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 319, col: 19, offset: 7510},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 319, col: 29, offset: 7520},
										expr: &seqExpr{
											pos: position{line: 319, col: 30, offset: 7521},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 319, col: 30, offset: 7521},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 319, col: 33, offset: 7524},
													name: "Value",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 319, col: 41, offset: 7532},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 319, col: 45, offset: 7536},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 337, col: 1, offset: 8042},
			expr: &actionExpr{
				pos: position{line: 337, col: 16, offset: 8057},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 337, col: 16, offset: 8057},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 337, col: 16, offset: 8057},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 337, col: 18, offset: 8059},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 337, col: 21, offset: 8062},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 337, col: 27, offset: 8068},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 337, col: 32, offset: 8073},
								expr: &seqExpr{
									pos: position{line: 337, col: 33, offset: 8074},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 337, col: 33, offset: 8074},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 337, col: 36, offset: 8077},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 337, col: 45, offset: 8086},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 337, col: 48, offset: 8089},
											name: "BinOp",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 337, col: 56, offset: 8097},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 337, col: 58, offset: 8099},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 337, col: 62, offset: 8103},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 357, col: 1, offset: 8762},
			expr: &choiceExpr{
				pos: position{line: 357, col: 9, offset: 8770},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 357, col: 9, offset: 8770},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 357, col: 21, offset: 8782},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 357, col: 37, offset: 8798},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 357, col: 48, offset: 8809},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 357, col: 60, offset: 8821},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 359, col: 1, offset: 8834},
			expr: &actionExpr{
				pos: position{line: 359, col: 13, offset: 8846},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 359, col: 13, offset: 8846},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 359, col: 13, offset: 8846},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 359, col: 15, offset: 8848},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 359, col: 21, offset: 8854},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 359, col: 35, offset: 8868},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 359, col: 40, offset: 8873},
								expr: &seqExpr{
									pos: position{line: 359, col: 41, offset: 8874},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 359, col: 41, offset: 8874},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 359, col: 44, offset: 8877},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 359, col: 60, offset: 8893},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 359, col: 63, offset: 8896},
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
			pos:  position{line: 378, col: 1, offset: 9503},
			expr: &actionExpr{
				pos: position{line: 378, col: 17, offset: 9519},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 378, col: 17, offset: 9519},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 378, col: 17, offset: 9519},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 378, col: 19, offset: 9521},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 378, col: 25, offset: 9527},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 378, col: 34, offset: 9536},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 378, col: 39, offset: 9541},
								expr: &seqExpr{
									pos: position{line: 378, col: 40, offset: 9542},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 378, col: 40, offset: 9542},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 378, col: 43, offset: 9545},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 378, col: 60, offset: 9562},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 378, col: 63, offset: 9565},
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
			pos:  position{line: 398, col: 1, offset: 10170},
			expr: &actionExpr{
				pos: position{line: 398, col: 12, offset: 10181},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 398, col: 12, offset: 10181},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 398, col: 12, offset: 10181},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 398, col: 14, offset: 10183},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 398, col: 20, offset: 10189},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 398, col: 30, offset: 10199},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 398, col: 35, offset: 10204},
								expr: &seqExpr{
									pos: position{line: 398, col: 36, offset: 10205},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 398, col: 36, offset: 10205},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 398, col: 39, offset: 10208},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 398, col: 51, offset: 10220},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 398, col: 54, offset: 10223},
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
			pos:  position{line: 418, col: 1, offset: 10825},
			expr: &actionExpr{
				pos: position{line: 418, col: 13, offset: 10837},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 418, col: 13, offset: 10837},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 418, col: 13, offset: 10837},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 418, col: 15, offset: 10839},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 418, col: 21, offset: 10845},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 418, col: 33, offset: 10857},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 418, col: 38, offset: 10862},
								expr: &seqExpr{
									pos: position{line: 418, col: 39, offset: 10863},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 418, col: 39, offset: 10863},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 42, offset: 10866},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 55, offset: 10879},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 58, offset: 10882},
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
			pos:  position{line: 437, col: 1, offset: 11487},
			expr: &choiceExpr{
				pos: position{line: 437, col: 15, offset: 11501},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 437, col: 15, offset: 11501},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 437, col: 15, offset: 11501},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 437, col: 15, offset: 11501},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 437, col: 19, offset: 11505},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 437, col: 21, offset: 11507},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 437, col: 27, offset: 11513},
										name: "BinOp",
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 33, offset: 11519},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 440, col: 5, offset: 11669},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 442, col: 1, offset: 11676},
			expr: &choiceExpr{
				pos: position{line: 442, col: 12, offset: 11687},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 442, col: 12, offset: 11687},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 30, offset: 11705},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 49, offset: 11724},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 64, offset: 11739},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 444, col: 1, offset: 11752},
			expr: &actionExpr{
				pos: position{line: 444, col: 19, offset: 11770},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 444, col: 21, offset: 11772},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 444, col: 21, offset: 11772},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 444, col: 29, offset: 11780},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 444, col: 36, offset: 11787},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 448, col: 1, offset: 11886},
			expr: &actionExpr{
				pos: position{line: 448, col: 20, offset: 11905},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 448, col: 22, offset: 11907},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 448, col: 22, offset: 11907},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 448, col: 29, offset: 11914},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 448, col: 36, offset: 11921},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 448, col: 42, offset: 11927},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 448, col: 48, offset: 11933},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 448, col: 56, offset: 11941},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 452, col: 1, offset: 12047},
			expr: &choiceExpr{
				pos: position{line: 452, col: 16, offset: 12062},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 452, col: 16, offset: 12062},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 452, col: 18, offset: 12064},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 452, col: 18, offset: 12064},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 452, col: 25, offset: 12071},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 455, col: 3, offset: 12177},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 455, col: 5, offset: 12179},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 455, col: 5, offset: 12179},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 455, col: 11, offset: 12185},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 455, col: 17, offset: 12191},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 458, col: 3, offset: 12294},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 458, col: 3, offset: 12294},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 462, col: 1, offset: 12398},
			expr: &choiceExpr{
				pos: position{line: 462, col: 15, offset: 12412},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 462, col: 15, offset: 12412},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 462, col: 17, offset: 12414},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 462, col: 17, offset: 12414},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 462, col: 24, offset: 12421},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 465, col: 3, offset: 12527},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 465, col: 5, offset: 12529},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 465, col: 5, offset: 12529},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 465, col: 11, offset: 12535},
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
			pos:  position{line: 469, col: 1, offset: 12637},
			expr: &choiceExpr{
				pos: position{line: 469, col: 9, offset: 12645},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 469, col: 9, offset: 12645},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 469, col: 22, offset: 12658},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 469, col: 22, offset: 12658},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 469, col: 24, offset: 12660},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 473, col: 1, offset: 12695},
			expr: &choiceExpr{
				pos: position{line: 473, col: 14, offset: 12708},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 473, col: 14, offset: 12708},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 473, col: 29, offset: 12723},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 475, col: 1, offset: 12731},
			expr: &choiceExpr{
				pos: position{line: 475, col: 14, offset: 12744},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 475, col: 14, offset: 12744},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 29, offset: 12759},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 477, col: 1, offset: 12771},
			expr: &actionExpr{
				pos: position{line: 477, col: 16, offset: 12786},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 477, col: 17, offset: 12787},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 477, col: 17, offset: 12787},
							expr: &charClassMatcher{
								pos:        position{line: 477, col: 17, offset: 12787},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 477, col: 24, offset: 12794},
							expr: &charClassMatcher{
								pos:        position{line: 477, col: 24, offset: 12794},
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
			pos:  position{line: 481, col: 1, offset: 12905},
			expr: &actionExpr{
				pos: position{line: 481, col: 14, offset: 12918},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 481, col: 15, offset: 12919},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 481, col: 15, offset: 12919},
							expr: &charClassMatcher{
								pos:        position{line: 481, col: 15, offset: 12919},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 481, col: 22, offset: 12926},
							expr: &charClassMatcher{
								pos:        position{line: 481, col: 22, offset: 12926},
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
			pos:  position{line: 485, col: 1, offset: 13037},
			expr: &choiceExpr{
				pos: position{line: 485, col: 9, offset: 13045},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 485, col: 9, offset: 13045},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 485, col: 9, offset: 13045},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 485, col: 9, offset: 13045},
									expr: &litMatcher{
										pos:        position{line: 485, col: 9, offset: 13045},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 485, col: 14, offset: 13050},
									expr: &charClassMatcher{
										pos:        position{line: 485, col: 14, offset: 13050},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 485, col: 21, offset: 13057},
									expr: &litMatcher{
										pos:        position{line: 485, col: 22, offset: 13058},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 492, col: 3, offset: 13234},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 492, col: 3, offset: 13234},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 492, col: 3, offset: 13234},
									expr: &litMatcher{
										pos:        position{line: 492, col: 3, offset: 13234},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 492, col: 8, offset: 13239},
									expr: &charClassMatcher{
										pos:        position{line: 492, col: 8, offset: 13239},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 492, col: 15, offset: 13246},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 492, col: 19, offset: 13250},
									expr: &charClassMatcher{
										pos:        position{line: 492, col: 19, offset: 13250},
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
						pos:        position{line: 499, col: 3, offset: 13440},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 499, col: 12, offset: 13449},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 499, col: 12, offset: 13449},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 505, col: 3, offset: 13650},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 505, col: 3, offset: 13650},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 508, col: 3, offset: 13713},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 508, col: 3, offset: 13713},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 508, col: 3, offset: 13713},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 508, col: 7, offset: 13717},
									expr: &seqExpr{
										pos: position{line: 508, col: 8, offset: 13718},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 508, col: 8, offset: 13718},
												expr: &ruleRefExpr{
													pos:  position{line: 508, col: 9, offset: 13719},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 508, col: 21, offset: 13731,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 508, col: 25, offset: 13735},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 515, col: 3, offset: 13919},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 515, col: 3, offset: 13919},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 515, col: 3, offset: 13919},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 515, col: 7, offset: 13923},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 515, col: 12, offset: 13928},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 515, col: 12, offset: 13928},
												expr: &ruleRefExpr{
													pos:  position{line: 515, col: 13, offset: 13929},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 515, col: 25, offset: 13941,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 515, col: 28, offset: 13944},
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
			pos:  position{line: 519, col: 1, offset: 14035},
			expr: &actionExpr{
				pos: position{line: 519, col: 10, offset: 14044},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 519, col: 11, offset: 14045},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 523, col: 1, offset: 14146},
			expr: &charClassMatcher{
				pos:        position{line: 523, col: 15, offset: 14160},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 525, col: 1, offset: 14176},
			expr: &choiceExpr{
				pos: position{line: 525, col: 18, offset: 14193},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 525, col: 18, offset: 14193},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 525, col: 37, offset: 14212},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 527, col: 1, offset: 14227},
			expr: &charClassMatcher{
				pos:        position{line: 527, col: 20, offset: 14246},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 529, col: 1, offset: 14259},
			expr: &charClassMatcher{
				pos:        position{line: 529, col: 16, offset: 14274},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 531, col: 1, offset: 14281},
			expr: &charClassMatcher{
				pos:        position{line: 531, col: 23, offset: 14303},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 533, col: 1, offset: 14310},
			expr: &charClassMatcher{
				pos:        position{line: 533, col: 12, offset: 14321},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 535, col: 1, offset: 14332},
			expr: &oneOrMoreExpr{
				pos: position{line: 535, col: 22, offset: 14353},
				expr: &charClassMatcher{
					pos:        position{line: 535, col: 22, offset: 14353},
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
			pos:         position{line: 537, col: 1, offset: 14365},
			expr: &zeroOrMoreExpr{
				pos: position{line: 537, col: 18, offset: 14382},
				expr: &charClassMatcher{
					pos:        position{line: 537, col: 18, offset: 14382},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 539, col: 1, offset: 14394},
			expr: &notExpr{
				pos: position{line: 539, col: 7, offset: 14400},
				expr: &anyMatcher{
					line: 539, col: 8, offset: 14401,
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

func (c *current) onCall2(module, fn, arguments interface{}) (interface{}, error) {
	fmt.Println("call", string(c.text))

	args := []Ast{}
	vals := arguments.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(arguments)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			args = append(args, v)
		}
	}

	return Call{Module: module.(Ast), Function: fn.(Ast), Arguments: args, ValueType: CONTAINER}, nil
}

func (p *parser) callonCall2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall2(stack["module"], stack["fn"], stack["arguments"])
}

func (c *current) onCall16(fn, arguments interface{}) (interface{}, error) {
	fmt.Println("call", string(c.text))
	var mod BasicAst
	args := []Ast{}
	vals := arguments.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(arguments)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			args = append(args, v)
		}
	}

	return Call{Module: mod, Function: fn.(Ast), Arguments: args, ValueType: CONTAINER}, nil
}

func (p *parser) callonCall16() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall16(stack["fn"], stack["arguments"])
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
	fmt.Println("adding an error at", pos)
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
	
	//fmt.Println("pos updated:", p.pt)
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
			//fmt.Println("the first rule failed", p.cur.pos)
			p.addErrAt(errNoMatch, p.cur.pos)
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
