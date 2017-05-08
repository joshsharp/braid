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
						pos: position{line: 180, col: 1, offset: 3460},
						run: (*parser).callonStatement13,
						expr: &seqExpr{
							pos: position{line: 180, col: 1, offset: 3460},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 180, col: 1, offset: 3460},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 7, offset: 3466},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 10, offset: 3469},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 12, offset: 3471},
										name: "Identifier",
									},
								},
								&labeledExpr{
									pos:   position{line: 180, col: 23, offset: 3482},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 180, col: 28, offset: 3487},
										expr: &seqExpr{
											pos: position{line: 180, col: 29, offset: 3488},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 180, col: 29, offset: 3488},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 180, col: 33, offset: 3492},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 180, col: 35, offset: 3494},
													name: "Identifier",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 48, offset: 3507},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 180, col: 50, offset: 3509},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 54, offset: 3513},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 56, offset: 3515},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 61, offset: 3520},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 66, offset: 3525},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 5, offset: 3962},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 195, col: 1, offset: 3968},
			expr: &actionExpr{
				pos: position{line: 195, col: 8, offset: 3975},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 195, col: 8, offset: 3975},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 195, col: 12, offset: 3979},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 195, col: 12, offset: 3979},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 195, col: 23, offset: 3990},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 195, col: 32, offset: 3999},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 200, col: 1, offset: 4091},
			expr: &choiceExpr{
				pos: position{line: 200, col: 10, offset: 4100},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 200, col: 10, offset: 4100},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 200, col: 10, offset: 4100},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 200, col: 10, offset: 4100},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 15, offset: 4105},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 18, offset: 4108},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 23, offset: 4113},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 33, offset: 4123},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 35, offset: 4125},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 39, offset: 4129},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 41, offset: 4131},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 47, offset: 4137},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 48, offset: 4138},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 60, offset: 4150},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 62, offset: 4152},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 66, offset: 4156},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 68, offset: 4158},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 75, offset: 4165},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 77, offset: 4167},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 85, offset: 4175},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 212, col: 1, offset: 4503},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 212, col: 1, offset: 4503},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 212, col: 1, offset: 4503},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 6, offset: 4508},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 9, offset: 4511},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 14, offset: 4516},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 24, offset: 4526},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 26, offset: 4528},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 30, offset: 4532},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 32, offset: 4534},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 38, offset: 4540},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 39, offset: 4541},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 51, offset: 4553},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 53, offset: 4555},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 57, offset: 4559},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 59, offset: 4561},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 66, offset: 4568},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 68, offset: 4570},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 72, offset: 4574},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 74, offset: 4576},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 80, offset: 4582},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 81, offset: 4583},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 93, offset: 4595},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 95, offset: 4597},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 99, offset: 4601},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 231, col: 1, offset: 5100},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 231, col: 1, offset: 5100},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 231, col: 1, offset: 5100},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 6, offset: 5105},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 9, offset: 5108},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 14, offset: 5113},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 24, offset: 5123},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 231, col: 26, offset: 5125},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 30, offset: 5129},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 32, offset: 5131},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 231, col: 38, offset: 5137},
										expr: &ruleRefExpr{
											pos:  position{line: 231, col: 39, offset: 5138},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 231, col: 51, offset: 5150},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 55, offset: 5154},
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
			pos:  position{line: 243, col: 1, offset: 5448},
			expr: &actionExpr{
				pos: position{line: 243, col: 12, offset: 5459},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 243, col: 12, offset: 5459},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 243, col: 12, offset: 5459},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 19, offset: 5466},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 243, col: 22, offset: 5469},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 243, col: 26, offset: 5473},
								expr: &seqExpr{
									pos: position{line: 243, col: 27, offset: 5474},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 243, col: 27, offset: 5474},
											name: "Identifier",
										},
										&ruleRefExpr{
											pos:  position{line: 243, col: 38, offset: 5485},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 43, offset: 5490},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 243, col: 45, offset: 5492},
							expr: &litMatcher{
								pos:        position{line: 243, col: 45, offset: 5492},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 51, offset: 5498},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 243, col: 53, offset: 5500},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 57, offset: 5504},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 243, col: 59, offset: 5506},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 243, col: 70, offset: 5517},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 71, offset: 5518},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 83, offset: 5530},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 243, col: 85, offset: 5532},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 89, offset: 5536},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 266, col: 1, offset: 6177},
			expr: &actionExpr{
				pos: position{line: 266, col: 16, offset: 6192},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 266, col: 16, offset: 6192},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 266, col: 16, offset: 6192},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 266, col: 18, offset: 6194},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 21, offset: 6197},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 266, col: 27, offset: 6203},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 266, col: 32, offset: 6208},
								expr: &seqExpr{
									pos: position{line: 266, col: 33, offset: 6209},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 266, col: 33, offset: 6209},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 266, col: 36, offset: 6212},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 266, col: 45, offset: 6221},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 266, col: 48, offset: 6224},
											name: "BinOp",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 56, offset: 6232},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 266, col: 58, offset: 6234},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 62, offset: 6238},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 286, col: 1, offset: 6897},
			expr: &choiceExpr{
				pos: position{line: 286, col: 9, offset: 6905},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 286, col: 9, offset: 6905},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 21, offset: 6917},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 37, offset: 6933},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 48, offset: 6944},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 60, offset: 6956},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 288, col: 1, offset: 6969},
			expr: &actionExpr{
				pos: position{line: 288, col: 13, offset: 6981},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 288, col: 13, offset: 6981},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 288, col: 13, offset: 6981},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 288, col: 15, offset: 6983},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 288, col: 21, offset: 6989},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 288, col: 35, offset: 7003},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 288, col: 40, offset: 7008},
								expr: &seqExpr{
									pos: position{line: 288, col: 41, offset: 7009},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 288, col: 41, offset: 7009},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 288, col: 44, offset: 7012},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 288, col: 60, offset: 7028},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 288, col: 63, offset: 7031},
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
			pos:  position{line: 307, col: 1, offset: 7638},
			expr: &actionExpr{
				pos: position{line: 307, col: 17, offset: 7654},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 307, col: 17, offset: 7654},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 307, col: 17, offset: 7654},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 307, col: 19, offset: 7656},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 307, col: 25, offset: 7662},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 307, col: 34, offset: 7671},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 307, col: 39, offset: 7676},
								expr: &seqExpr{
									pos: position{line: 307, col: 40, offset: 7677},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 307, col: 40, offset: 7677},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 307, col: 43, offset: 7680},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 307, col: 60, offset: 7697},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 307, col: 63, offset: 7700},
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
			pos:  position{line: 327, col: 1, offset: 8305},
			expr: &actionExpr{
				pos: position{line: 327, col: 12, offset: 8316},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 327, col: 12, offset: 8316},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 327, col: 12, offset: 8316},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 327, col: 14, offset: 8318},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 327, col: 20, offset: 8324},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 327, col: 30, offset: 8334},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 327, col: 35, offset: 8339},
								expr: &seqExpr{
									pos: position{line: 327, col: 36, offset: 8340},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 327, col: 36, offset: 8340},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 327, col: 39, offset: 8343},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 327, col: 51, offset: 8355},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 327, col: 54, offset: 8358},
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
			pos:  position{line: 347, col: 1, offset: 8960},
			expr: &actionExpr{
				pos: position{line: 347, col: 13, offset: 8972},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 347, col: 13, offset: 8972},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 347, col: 13, offset: 8972},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 347, col: 15, offset: 8974},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 347, col: 21, offset: 8980},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 347, col: 33, offset: 8992},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 347, col: 38, offset: 8997},
								expr: &seqExpr{
									pos: position{line: 347, col: 39, offset: 8998},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 347, col: 39, offset: 8998},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 347, col: 42, offset: 9001},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 347, col: 55, offset: 9014},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 347, col: 58, offset: 9017},
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
			pos:  position{line: 366, col: 1, offset: 9622},
			expr: &choiceExpr{
				pos: position{line: 366, col: 15, offset: 9636},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 366, col: 15, offset: 9636},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 366, col: 15, offset: 9636},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 366, col: 15, offset: 9636},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 366, col: 19, offset: 9640},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 366, col: 21, offset: 9642},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 366, col: 27, offset: 9648},
										name: "BinOp",
									},
								},
								&litMatcher{
									pos:        position{line: 366, col: 33, offset: 9654},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 369, col: 5, offset: 9804},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 371, col: 1, offset: 9811},
			expr: &choiceExpr{
				pos: position{line: 371, col: 12, offset: 9822},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 371, col: 12, offset: 9822},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 371, col: 30, offset: 9840},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 371, col: 49, offset: 9859},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 371, col: 64, offset: 9874},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 373, col: 1, offset: 9887},
			expr: &actionExpr{
				pos: position{line: 373, col: 19, offset: 9905},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 373, col: 21, offset: 9907},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 373, col: 21, offset: 9907},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 29, offset: 9915},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 373, col: 36, offset: 9922},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 377, col: 1, offset: 10021},
			expr: &actionExpr{
				pos: position{line: 377, col: 20, offset: 10040},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 377, col: 22, offset: 10042},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 377, col: 22, offset: 10042},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 377, col: 29, offset: 10049},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 377, col: 36, offset: 10056},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 377, col: 42, offset: 10062},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 377, col: 48, offset: 10068},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 377, col: 56, offset: 10076},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 381, col: 1, offset: 10182},
			expr: &choiceExpr{
				pos: position{line: 381, col: 16, offset: 10197},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 381, col: 16, offset: 10197},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 381, col: 18, offset: 10199},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 381, col: 18, offset: 10199},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 381, col: 25, offset: 10206},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 3, offset: 10312},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 384, col: 5, offset: 10314},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 384, col: 5, offset: 10314},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 384, col: 11, offset: 10320},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 384, col: 17, offset: 10326},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 387, col: 3, offset: 10429},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 387, col: 3, offset: 10429},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 391, col: 1, offset: 10533},
			expr: &choiceExpr{
				pos: position{line: 391, col: 15, offset: 10547},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 391, col: 15, offset: 10547},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 391, col: 17, offset: 10549},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 391, col: 17, offset: 10549},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 391, col: 24, offset: 10556},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 394, col: 3, offset: 10662},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 394, col: 5, offset: 10664},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 394, col: 5, offset: 10664},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 394, col: 11, offset: 10670},
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
			pos:  position{line: 398, col: 1, offset: 10772},
			expr: &choiceExpr{
				pos: position{line: 398, col: 9, offset: 10780},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 398, col: 9, offset: 10780},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 398, col: 22, offset: 10793},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 398, col: 22, offset: 10793},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 398, col: 24, offset: 10795},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 403, col: 1, offset: 10831},
			expr: &actionExpr{
				pos: position{line: 403, col: 14, offset: 10844},
				run: (*parser).callonIdentifier1,
				expr: &choiceExpr{
					pos: position{line: 403, col: 15, offset: 10845},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 403, col: 15, offset: 10845},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 403, col: 15, offset: 10845},
									expr: &charClassMatcher{
										pos:        position{line: 403, col: 15, offset: 10845},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 403, col: 22, offset: 10852},
									expr: &charClassMatcher{
										pos:        position{line: 403, col: 22, offset: 10852},
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
							pos:        position{line: 403, col: 38, offset: 10868},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 407, col: 1, offset: 10969},
			expr: &choiceExpr{
				pos: position{line: 407, col: 9, offset: 10977},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 407, col: 9, offset: 10977},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 407, col: 9, offset: 10977},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 407, col: 9, offset: 10977},
									expr: &litMatcher{
										pos:        position{line: 407, col: 9, offset: 10977},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 407, col: 14, offset: 10982},
									expr: &charClassMatcher{
										pos:        position{line: 407, col: 14, offset: 10982},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 407, col: 21, offset: 10989},
									expr: &litMatcher{
										pos:        position{line: 407, col: 22, offset: 10990},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 414, col: 3, offset: 11166},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 414, col: 3, offset: 11166},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 414, col: 3, offset: 11166},
									expr: &litMatcher{
										pos:        position{line: 414, col: 3, offset: 11166},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 414, col: 8, offset: 11171},
									expr: &charClassMatcher{
										pos:        position{line: 414, col: 8, offset: 11171},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 414, col: 15, offset: 11178},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 414, col: 19, offset: 11182},
									expr: &charClassMatcher{
										pos:        position{line: 414, col: 19, offset: 11182},
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
						pos:        position{line: 421, col: 3, offset: 11372},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 421, col: 12, offset: 11381},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 421, col: 12, offset: 11381},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 427, col: 3, offset: 11582},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 427, col: 3, offset: 11582},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 430, col: 3, offset: 11645},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 430, col: 3, offset: 11645},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 430, col: 3, offset: 11645},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 430, col: 7, offset: 11649},
									expr: &seqExpr{
										pos: position{line: 430, col: 8, offset: 11650},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 430, col: 8, offset: 11650},
												expr: &ruleRefExpr{
													pos:  position{line: 430, col: 9, offset: 11651},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 430, col: 21, offset: 11663,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 430, col: 25, offset: 11667},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 437, col: 3, offset: 11851},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 437, col: 3, offset: 11851},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 437, col: 3, offset: 11851},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 437, col: 7, offset: 11855},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 437, col: 12, offset: 11860},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 437, col: 12, offset: 11860},
												expr: &ruleRefExpr{
													pos:  position{line: 437, col: 13, offset: 11861},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 437, col: 25, offset: 11873,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 28, offset: 11876},
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
			pos:  position{line: 441, col: 1, offset: 11967},
			expr: &charClassMatcher{
				pos:        position{line: 441, col: 15, offset: 11981},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 443, col: 1, offset: 11997},
			expr: &choiceExpr{
				pos: position{line: 443, col: 18, offset: 12014},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 443, col: 18, offset: 12014},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 443, col: 37, offset: 12033},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 445, col: 1, offset: 12048},
			expr: &charClassMatcher{
				pos:        position{line: 445, col: 20, offset: 12067},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 447, col: 1, offset: 12080},
			expr: &charClassMatcher{
				pos:        position{line: 447, col: 16, offset: 12095},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 449, col: 1, offset: 12102},
			expr: &charClassMatcher{
				pos:        position{line: 449, col: 23, offset: 12124},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 451, col: 1, offset: 12131},
			expr: &charClassMatcher{
				pos:        position{line: 451, col: 12, offset: 12142},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 453, col: 1, offset: 12153},
			expr: &oneOrMoreExpr{
				pos: position{line: 453, col: 22, offset: 12174},
				expr: &charClassMatcher{
					pos:        position{line: 453, col: 22, offset: 12174},
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
			pos:         position{line: 455, col: 1, offset: 12186},
			expr: &zeroOrMoreExpr{
				pos: position{line: 455, col: 18, offset: 12203},
				expr: &charClassMatcher{
					pos:        position{line: 455, col: 18, offset: 12203},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 457, col: 1, offset: 12215},
			expr: &notExpr{
				pos: position{line: 457, col: 7, offset: 12221},
				expr: &anyMatcher{
					line: 457, col: 8, offset: 12222,
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
