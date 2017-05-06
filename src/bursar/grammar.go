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
			pos:  position{line: 161, col: 1, offset: 2778},
			expr: &actionExpr{
				pos: position{line: 161, col: 10, offset: 2787},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 161, col: 10, offset: 2787},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 161, col: 10, offset: 2787},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 12, offset: 2789},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 17, offset: 2794},
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 27, offset: 2804},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 29, offset: 2806},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 161, col: 34, offset: 2811},
								expr: &ruleRefExpr{
									pos:  position{line: 161, col: 35, offset: 2812},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 47, offset: 2824},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 49, offset: 2826},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 176, col: 1, offset: 3310},
			expr: &choiceExpr{
				pos: position{line: 176, col: 13, offset: 3322},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 176, col: 13, offset: 3322},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 176, col: 13, offset: 3322},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 176, col: 13, offset: 3322},
									val:        "#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 176, col: 17, offset: 3326},
									label: "comment",
									expr: &zeroOrMoreExpr{
										pos: position{line: 176, col: 25, offset: 3334},
										expr: &seqExpr{
											pos: position{line: 176, col: 26, offset: 3335},
											exprs: []interface{}{
												&notExpr{
													pos: position{line: 176, col: 26, offset: 3335},
													expr: &ruleRefExpr{
														pos:  position{line: 176, col: 27, offset: 3336},
														name: "EscapedChar",
													},
												},
												&anyMatcher{
													line: 176, col: 39, offset: 3348,
												},
											},
										},
									},
								},
								&andExpr{
									pos: position{line: 176, col: 43, offset: 3352},
									expr: &litMatcher{
										pos:        position{line: 176, col: 44, offset: 3353},
										val:        "\n",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 180, col: 1, offset: 3458},
						run: (*parser).callonStatement13,
						expr: &seqExpr{
							pos: position{line: 180, col: 1, offset: 3458},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 180, col: 1, offset: 3458},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 7, offset: 3464},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 10, offset: 3467},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 12, offset: 3469},
										name: "Identifier",
									},
								},
								&labeledExpr{
									pos:   position{line: 180, col: 23, offset: 3480},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 180, col: 28, offset: 3485},
										expr: &seqExpr{
											pos: position{line: 180, col: 29, offset: 3486},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 180, col: 29, offset: 3486},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 180, col: 33, offset: 3490},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 180, col: 35, offset: 3492},
													name: "Identifier",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 48, offset: 3505},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 180, col: 50, offset: 3507},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 54, offset: 3511},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 56, offset: 3513},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 61, offset: 3518},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 66, offset: 3523},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 5, offset: 3964},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 195, col: 1, offset: 3970},
			expr: &actionExpr{
				pos: position{line: 195, col: 8, offset: 3977},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 195, col: 8, offset: 3977},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 195, col: 12, offset: 3981},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 195, col: 12, offset: 3981},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 195, col: 23, offset: 3992},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 195, col: 32, offset: 4001},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 200, col: 1, offset: 4083},
			expr: &choiceExpr{
				pos: position{line: 200, col: 10, offset: 4092},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 200, col: 10, offset: 4092},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 200, col: 10, offset: 4092},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 200, col: 10, offset: 4092},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 15, offset: 4097},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 18, offset: 4100},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 23, offset: 4105},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 33, offset: 4115},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 35, offset: 4117},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 39, offset: 4121},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 41, offset: 4123},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 47, offset: 4129},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 48, offset: 4130},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 60, offset: 4142},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 62, offset: 4144},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 66, offset: 4148},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 68, offset: 4150},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 75, offset: 4157},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 77, offset: 4159},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 85, offset: 4167},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 212, col: 1, offset: 4495},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 212, col: 1, offset: 4495},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 212, col: 1, offset: 4495},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 6, offset: 4500},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 9, offset: 4503},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 14, offset: 4508},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 24, offset: 4518},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 26, offset: 4520},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 30, offset: 4524},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 32, offset: 4526},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 38, offset: 4532},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 39, offset: 4533},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 51, offset: 4545},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 53, offset: 4547},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 57, offset: 4551},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 59, offset: 4553},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 66, offset: 4560},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 68, offset: 4562},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 72, offset: 4566},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 74, offset: 4568},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 80, offset: 4574},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 81, offset: 4575},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 93, offset: 4587},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 95, offset: 4589},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 99, offset: 4593},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 231, col: 1, offset: 5092},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 231, col: 1, offset: 5092},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 231, col: 1, offset: 5092},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 6, offset: 5097},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 9, offset: 5100},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 14, offset: 5105},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 24, offset: 5115},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 231, col: 26, offset: 5117},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 30, offset: 5121},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 32, offset: 5123},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 231, col: 38, offset: 5129},
										expr: &ruleRefExpr{
											pos:  position{line: 231, col: 39, offset: 5130},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 231, col: 51, offset: 5142},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 55, offset: 5146},
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
			pos:  position{line: 243, col: 1, offset: 5440},
			expr: &actionExpr{
				pos: position{line: 243, col: 12, offset: 5451},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 243, col: 12, offset: 5451},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 243, col: 12, offset: 5451},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 19, offset: 5458},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 243, col: 22, offset: 5461},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 243, col: 26, offset: 5465},
								expr: &seqExpr{
									pos: position{line: 243, col: 27, offset: 5466},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 243, col: 27, offset: 5466},
											name: "Identifier",
										},
										&ruleRefExpr{
											pos:  position{line: 243, col: 38, offset: 5477},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 43, offset: 5482},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 243, col: 45, offset: 5484},
							expr: &litMatcher{
								pos:        position{line: 243, col: 45, offset: 5484},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 51, offset: 5490},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 243, col: 53, offset: 5492},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 57, offset: 5496},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 243, col: 59, offset: 5498},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 243, col: 70, offset: 5509},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 71, offset: 5510},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 83, offset: 5522},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 243, col: 85, offset: 5524},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 89, offset: 5528},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 266, col: 1, offset: 6169},
			expr: &actionExpr{
				pos: position{line: 266, col: 16, offset: 6184},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 266, col: 16, offset: 6184},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 266, col: 16, offset: 6184},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 266, col: 18, offset: 6186},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 21, offset: 6189},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 266, col: 27, offset: 6195},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 266, col: 32, offset: 6200},
								expr: &ruleRefExpr{
									pos:  position{line: 266, col: 33, offset: 6201},
									name: "BinOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 41, offset: 6209},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 266, col: 43, offset: 6211},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 47, offset: 6215},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 280, col: 1, offset: 6667},
			expr: &choiceExpr{
				pos: position{line: 280, col: 9, offset: 6675},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 280, col: 9, offset: 6675},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 280, col: 21, offset: 6687},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 280, col: 37, offset: 6703},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 280, col: 48, offset: 6714},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 280, col: 60, offset: 6726},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 282, col: 1, offset: 6739},
			expr: &actionExpr{
				pos: position{line: 282, col: 13, offset: 6751},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 282, col: 13, offset: 6751},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 282, col: 13, offset: 6751},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 282, col: 15, offset: 6753},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 21, offset: 6759},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 282, col: 35, offset: 6773},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 282, col: 40, offset: 6778},
								expr: &seqExpr{
									pos: position{line: 282, col: 41, offset: 6779},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 282, col: 41, offset: 6779},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 282, col: 44, offset: 6782},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 282, col: 60, offset: 6798},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 282, col: 63, offset: 6801},
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
			pos:  position{line: 301, col: 1, offset: 7408},
			expr: &actionExpr{
				pos: position{line: 301, col: 17, offset: 7424},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 301, col: 17, offset: 7424},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 301, col: 17, offset: 7424},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 301, col: 19, offset: 7426},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 301, col: 25, offset: 7432},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 301, col: 34, offset: 7441},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 301, col: 39, offset: 7446},
								expr: &seqExpr{
									pos: position{line: 301, col: 40, offset: 7447},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 301, col: 40, offset: 7447},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 301, col: 43, offset: 7450},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 301, col: 60, offset: 7467},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 301, col: 63, offset: 7470},
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
			pos:  position{line: 321, col: 1, offset: 8075},
			expr: &actionExpr{
				pos: position{line: 321, col: 12, offset: 8086},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 321, col: 12, offset: 8086},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 321, col: 12, offset: 8086},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 321, col: 14, offset: 8088},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 321, col: 20, offset: 8094},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 321, col: 30, offset: 8104},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 321, col: 35, offset: 8109},
								expr: &seqExpr{
									pos: position{line: 321, col: 36, offset: 8110},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 321, col: 36, offset: 8110},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 321, col: 39, offset: 8113},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 321, col: 51, offset: 8125},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 321, col: 54, offset: 8128},
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
			pos:  position{line: 341, col: 1, offset: 8730},
			expr: &actionExpr{
				pos: position{line: 341, col: 13, offset: 8742},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 341, col: 13, offset: 8742},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 341, col: 13, offset: 8742},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 341, col: 15, offset: 8744},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 341, col: 21, offset: 8750},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 341, col: 33, offset: 8762},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 341, col: 38, offset: 8767},
								expr: &seqExpr{
									pos: position{line: 341, col: 39, offset: 8768},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 341, col: 39, offset: 8768},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 341, col: 42, offset: 8771},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 341, col: 55, offset: 8784},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 341, col: 58, offset: 8787},
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
			pos:  position{line: 360, col: 1, offset: 9392},
			expr: &choiceExpr{
				pos: position{line: 360, col: 15, offset: 9406},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 360, col: 15, offset: 9406},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 360, col: 15, offset: 9406},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 360, col: 15, offset: 9406},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 360, col: 19, offset: 9410},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 360, col: 21, offset: 9412},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 360, col: 27, offset: 9418},
										name: "BinOp",
									},
								},
								&litMatcher{
									pos:        position{line: 360, col: 33, offset: 9424},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 365, col: 5, offset: 9602},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 367, col: 1, offset: 9609},
			expr: &choiceExpr{
				pos: position{line: 367, col: 12, offset: 9620},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 367, col: 12, offset: 9620},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 367, col: 30, offset: 9638},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 367, col: 49, offset: 9657},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 367, col: 64, offset: 9672},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 369, col: 1, offset: 9685},
			expr: &actionExpr{
				pos: position{line: 369, col: 19, offset: 9703},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 369, col: 21, offset: 9705},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 369, col: 21, offset: 9705},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 369, col: 29, offset: 9713},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 369, col: 36, offset: 9720},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 373, col: 1, offset: 9819},
			expr: &actionExpr{
				pos: position{line: 373, col: 20, offset: 9838},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 373, col: 22, offset: 9840},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 373, col: 22, offset: 9840},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 29, offset: 9847},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 36, offset: 9854},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 42, offset: 9860},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 48, offset: 9866},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 56, offset: 9874},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 377, col: 1, offset: 9980},
			expr: &choiceExpr{
				pos: position{line: 377, col: 16, offset: 9995},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 377, col: 16, offset: 9995},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 377, col: 18, offset: 9997},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 377, col: 18, offset: 9997},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 377, col: 25, offset: 10004},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 380, col: 3, offset: 10110},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 380, col: 5, offset: 10112},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 380, col: 5, offset: 10112},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 380, col: 11, offset: 10118},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 380, col: 17, offset: 10124},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 383, col: 3, offset: 10227},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 383, col: 3, offset: 10227},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 387, col: 1, offset: 10331},
			expr: &choiceExpr{
				pos: position{line: 387, col: 15, offset: 10345},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 387, col: 15, offset: 10345},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 387, col: 17, offset: 10347},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 387, col: 17, offset: 10347},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 387, col: 24, offset: 10354},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 390, col: 3, offset: 10460},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 390, col: 5, offset: 10462},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 390, col: 5, offset: 10462},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 390, col: 11, offset: 10468},
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
			pos:  position{line: 394, col: 1, offset: 10570},
			expr: &choiceExpr{
				pos: position{line: 394, col: 9, offset: 10578},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 394, col: 9, offset: 10578},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 394, col: 22, offset: 10591},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 394, col: 22, offset: 10591},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 394, col: 24, offset: 10593},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 399, col: 1, offset: 10629},
			expr: &actionExpr{
				pos: position{line: 399, col: 14, offset: 10642},
				run: (*parser).callonIdentifier1,
				expr: &choiceExpr{
					pos: position{line: 399, col: 15, offset: 10643},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 399, col: 15, offset: 10643},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 399, col: 15, offset: 10643},
									expr: &charClassMatcher{
										pos:        position{line: 399, col: 15, offset: 10643},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 399, col: 22, offset: 10650},
									expr: &charClassMatcher{
										pos:        position{line: 399, col: 22, offset: 10650},
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
							pos:        position{line: 399, col: 38, offset: 10666},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 403, col: 1, offset: 10767},
			expr: &choiceExpr{
				pos: position{line: 403, col: 9, offset: 10775},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 403, col: 9, offset: 10775},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 403, col: 9, offset: 10775},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 403, col: 9, offset: 10775},
									expr: &litMatcher{
										pos:        position{line: 403, col: 9, offset: 10775},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 403, col: 14, offset: 10780},
									expr: &charClassMatcher{
										pos:        position{line: 403, col: 14, offset: 10780},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 403, col: 21, offset: 10787},
									expr: &litMatcher{
										pos:        position{line: 403, col: 22, offset: 10788},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 410, col: 3, offset: 10964},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 410, col: 3, offset: 10964},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 410, col: 3, offset: 10964},
									expr: &litMatcher{
										pos:        position{line: 410, col: 3, offset: 10964},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 410, col: 8, offset: 10969},
									expr: &charClassMatcher{
										pos:        position{line: 410, col: 8, offset: 10969},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 410, col: 15, offset: 10976},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 410, col: 19, offset: 10980},
									expr: &charClassMatcher{
										pos:        position{line: 410, col: 19, offset: 10980},
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
						pos:        position{line: 417, col: 3, offset: 11170},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 417, col: 12, offset: 11179},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 417, col: 12, offset: 11179},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 423, col: 3, offset: 11380},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 423, col: 3, offset: 11380},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 426, col: 3, offset: 11443},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 426, col: 3, offset: 11443},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 426, col: 3, offset: 11443},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 426, col: 7, offset: 11447},
									expr: &seqExpr{
										pos: position{line: 426, col: 8, offset: 11448},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 426, col: 8, offset: 11448},
												expr: &ruleRefExpr{
													pos:  position{line: 426, col: 9, offset: 11449},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 426, col: 21, offset: 11461,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 426, col: 25, offset: 11465},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 433, col: 3, offset: 11649},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 433, col: 3, offset: 11649},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 433, col: 3, offset: 11649},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 433, col: 7, offset: 11653},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 433, col: 12, offset: 11658},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 433, col: 12, offset: 11658},
												expr: &ruleRefExpr{
													pos:  position{line: 433, col: 13, offset: 11659},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 433, col: 25, offset: 11671,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 433, col: 28, offset: 11674},
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
			pos:  position{line: 437, col: 1, offset: 11765},
			expr: &charClassMatcher{
				pos:        position{line: 437, col: 15, offset: 11779},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 439, col: 1, offset: 11795},
			expr: &choiceExpr{
				pos: position{line: 439, col: 18, offset: 11812},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 439, col: 18, offset: 11812},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 439, col: 37, offset: 11831},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 441, col: 1, offset: 11846},
			expr: &charClassMatcher{
				pos:        position{line: 441, col: 20, offset: 11865},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 443, col: 1, offset: 11878},
			expr: &charClassMatcher{
				pos:        position{line: 443, col: 16, offset: 11893},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 445, col: 1, offset: 11900},
			expr: &charClassMatcher{
				pos:        position{line: 445, col: 23, offset: 11922},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 447, col: 1, offset: 11929},
			expr: &charClassMatcher{
				pos:        position{line: 447, col: 12, offset: 11940},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 449, col: 1, offset: 11951},
			expr: &oneOrMoreExpr{
				pos: position{line: 449, col: 22, offset: 11972},
				expr: &charClassMatcher{
					pos:        position{line: 449, col: 22, offset: 11972},
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
			pos:         position{line: 451, col: 1, offset: 11984},
			expr: &zeroOrMoreExpr{
				pos: position{line: 451, col: 18, offset: 12001},
				expr: &charClassMatcher{
					pos:        position{line: 451, col: 18, offset: 12001},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 453, col: 1, offset: 12013},
			expr: &notExpr{
				pos: position{line: 453, col: 7, offset: 12019},
				expr: &anyMatcher{
					line: 453, col: 8, offset: 12020,
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

func (c *current) onStatement2(comment interface{}) (interface{}, error) {
	return BasicAst{Type: "Comment", StringValue: string(c.text[1:]), ValueType: STRING}, nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["comment"])
}

func (c *current) onStatement13(i, rest, expr interface{}) (interface{}, error) {
	fmt.Printf("assignment: %s\n", string(c.text))
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
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
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

	subvalues := []Ast{first.(Ast)}
	return BasicAst{Type: "BinOpParens", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
