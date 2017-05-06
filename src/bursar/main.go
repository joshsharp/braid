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

func main() {
	input := `
# test
let _, b = 5.0 + 6.5;
let cheesy = func item item2 {
    item ++ " and " ++ item2 ++ " with cheese"; # more test
}
let tester = func a {
    let result = if a > 100 {
        a + 1;
    } else if a > 50 {
        a + 20;
    } else {
        a + 2;
    }
    result;
}
cheesy "pineapple" "bbq sauce";
`
	fmt.Println(input)
	r := strings.NewReader(input)
	result, err := ParseReader("", r)
	ast := result.(Ast)
	fmt.Println("=", ast.Print(0))
	fmt.Println(err)
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Module",
			pos:  position{line: 188, col: 1, offset: 3294},
			expr: &actionExpr{
				pos: position{line: 188, col: 10, offset: 3303},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 188, col: 10, offset: 3303},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 188, col: 10, offset: 3303},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 188, col: 12, offset: 3305},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 17, offset: 3310},
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 27, offset: 3320},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 188, col: 29, offset: 3322},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 188, col: 34, offset: 3327},
								expr: &ruleRefExpr{
									pos:  position{line: 188, col: 35, offset: 3328},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 47, offset: 3340},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 49, offset: 3342},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 203, col: 1, offset: 3826},
			expr: &choiceExpr{
				pos: position{line: 203, col: 13, offset: 3838},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 203, col: 13, offset: 3838},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 203, col: 13, offset: 3838},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 203, col: 13, offset: 3838},
									val:        "#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 203, col: 17, offset: 3842},
									label: "comment",
									expr: &zeroOrMoreExpr{
										pos: position{line: 203, col: 25, offset: 3850},
										expr: &seqExpr{
											pos: position{line: 203, col: 26, offset: 3851},
											exprs: []interface{}{
												&notExpr{
													pos: position{line: 203, col: 26, offset: 3851},
													expr: &ruleRefExpr{
														pos:  position{line: 203, col: 27, offset: 3852},
														name: "EscapedChar",
													},
												},
												&anyMatcher{
													line: 203, col: 39, offset: 3864,
												},
											},
										},
									},
								},
								&andExpr{
									pos: position{line: 203, col: 43, offset: 3868},
									expr: &litMatcher{
										pos:        position{line: 203, col: 44, offset: 3869},
										val:        "\n",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 207, col: 1, offset: 3974},
						run: (*parser).callonStatement13,
						expr: &seqExpr{
							pos: position{line: 207, col: 1, offset: 3974},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 207, col: 1, offset: 3974},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 207, col: 7, offset: 3980},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 207, col: 10, offset: 3983},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 207, col: 12, offset: 3985},
										name: "Value",
									},
								},
								&labeledExpr{
									pos:   position{line: 207, col: 18, offset: 3991},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 207, col: 23, offset: 3996},
										expr: &seqExpr{
											pos: position{line: 207, col: 24, offset: 3997},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 207, col: 24, offset: 3997},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 207, col: 28, offset: 4001},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 207, col: 30, offset: 4003},
													name: "Value",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 207, col: 38, offset: 4011},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 207, col: 40, offset: 4013},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 207, col: 44, offset: 4017},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 207, col: 46, offset: 4019},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 207, col: 51, offset: 4024},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 207, col: 56, offset: 4029},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 220, col: 5, offset: 4470},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 222, col: 1, offset: 4476},
			expr: &actionExpr{
				pos: position{line: 222, col: 8, offset: 4483},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 222, col: 8, offset: 4483},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 222, col: 12, offset: 4487},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 222, col: 12, offset: 4487},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 222, col: 23, offset: 4498},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 222, col: 32, offset: 4507},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 227, col: 1, offset: 4589},
			expr: &choiceExpr{
				pos: position{line: 227, col: 10, offset: 4598},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 227, col: 10, offset: 4598},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 227, col: 10, offset: 4598},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 227, col: 10, offset: 4598},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 227, col: 15, offset: 4603},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 227, col: 18, offset: 4606},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 227, col: 23, offset: 4611},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 227, col: 33, offset: 4621},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 227, col: 35, offset: 4623},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 227, col: 39, offset: 4627},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 227, col: 41, offset: 4629},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 227, col: 47, offset: 4635},
										expr: &ruleRefExpr{
											pos:  position{line: 227, col: 48, offset: 4636},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 227, col: 60, offset: 4648},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 227, col: 62, offset: 4650},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 227, col: 66, offset: 4654},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 227, col: 68, offset: 4656},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 227, col: 75, offset: 4663},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 227, col: 77, offset: 4665},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 227, col: 85, offset: 4673},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 239, col: 1, offset: 5001},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 239, col: 1, offset: 5001},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 239, col: 1, offset: 5001},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 6, offset: 5006},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 239, col: 9, offset: 5009},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 239, col: 14, offset: 5014},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 24, offset: 5024},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 26, offset: 5026},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 30, offset: 5030},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 239, col: 32, offset: 5032},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 239, col: 38, offset: 5038},
										expr: &ruleRefExpr{
											pos:  position{line: 239, col: 39, offset: 5039},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 51, offset: 5051},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 53, offset: 5053},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 57, offset: 5057},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 59, offset: 5059},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 66, offset: 5066},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 68, offset: 5068},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 72, offset: 5072},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 239, col: 74, offset: 5074},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 239, col: 80, offset: 5080},
										expr: &ruleRefExpr{
											pos:  position{line: 239, col: 81, offset: 5081},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 93, offset: 5093},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 95, offset: 5095},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 99, offset: 5099},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 258, col: 1, offset: 5598},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 258, col: 1, offset: 5598},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 258, col: 1, offset: 5598},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 6, offset: 5603},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 258, col: 9, offset: 5606},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 258, col: 14, offset: 5611},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 24, offset: 5621},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 258, col: 26, offset: 5623},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 30, offset: 5627},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 258, col: 32, offset: 5629},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 258, col: 38, offset: 5635},
										expr: &ruleRefExpr{
											pos:  position{line: 258, col: 39, offset: 5636},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 258, col: 51, offset: 5648},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 55, offset: 5652},
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
			pos:  position{line: 270, col: 1, offset: 5946},
			expr: &actionExpr{
				pos: position{line: 270, col: 12, offset: 5957},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 270, col: 12, offset: 5957},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 270, col: 12, offset: 5957},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 19, offset: 5964},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 270, col: 22, offset: 5967},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 270, col: 26, offset: 5971},
								expr: &seqExpr{
									pos: position{line: 270, col: 27, offset: 5972},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 270, col: 27, offset: 5972},
											name: "Identifier",
										},
										&ruleRefExpr{
											pos:  position{line: 270, col: 38, offset: 5983},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 43, offset: 5988},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 270, col: 45, offset: 5990},
							expr: &litMatcher{
								pos:        position{line: 270, col: 45, offset: 5990},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 51, offset: 5996},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 270, col: 53, offset: 5998},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 57, offset: 6002},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 270, col: 59, offset: 6004},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 270, col: 70, offset: 6015},
								expr: &ruleRefExpr{
									pos:  position{line: 270, col: 71, offset: 6016},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 83, offset: 6028},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 270, col: 85, offset: 6030},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 89, offset: 6034},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 293, col: 1, offset: 6683},
			expr: &actionExpr{
				pos: position{line: 293, col: 16, offset: 6698},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 293, col: 16, offset: 6698},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 293, col: 16, offset: 6698},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 293, col: 18, offset: 6700},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 21, offset: 6703},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 293, col: 27, offset: 6709},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 293, col: 32, offset: 6714},
								expr: &ruleRefExpr{
									pos:  position{line: 293, col: 33, offset: 6715},
									name: "BinOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 41, offset: 6723},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 293, col: 43, offset: 6725},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 47, offset: 6729},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 306, col: 1, offset: 7140},
			expr: &choiceExpr{
				pos: position{line: 306, col: 9, offset: 7148},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 306, col: 9, offset: 7148},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 21, offset: 7160},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 37, offset: 7176},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 48, offset: 7187},
						name: "BinOpHigh",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 308, col: 1, offset: 7198},
			expr: &actionExpr{
				pos: position{line: 308, col: 13, offset: 7210},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 308, col: 13, offset: 7210},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 308, col: 13, offset: 7210},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 308, col: 15, offset: 7212},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 21, offset: 7218},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 308, col: 35, offset: 7232},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 308, col: 40, offset: 7237},
								expr: &seqExpr{
									pos: position{line: 308, col: 41, offset: 7238},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 308, col: 41, offset: 7238},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 44, offset: 7241},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 60, offset: 7257},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 63, offset: 7260},
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
			pos:  position{line: 326, col: 1, offset: 7822},
			expr: &actionExpr{
				pos: position{line: 326, col: 17, offset: 7838},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 326, col: 17, offset: 7838},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 326, col: 17, offset: 7838},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 326, col: 19, offset: 7840},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 25, offset: 7846},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 326, col: 34, offset: 7855},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 326, col: 39, offset: 7860},
								expr: &seqExpr{
									pos: position{line: 326, col: 40, offset: 7861},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 326, col: 40, offset: 7861},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 326, col: 43, offset: 7864},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 326, col: 60, offset: 7881},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 326, col: 63, offset: 7884},
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
			pos:  position{line: 345, col: 1, offset: 8446},
			expr: &actionExpr{
				pos: position{line: 345, col: 12, offset: 8457},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 345, col: 12, offset: 8457},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 345, col: 12, offset: 8457},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 345, col: 14, offset: 8459},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 345, col: 20, offset: 8465},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 345, col: 30, offset: 8475},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 345, col: 35, offset: 8480},
								expr: &seqExpr{
									pos: position{line: 345, col: 36, offset: 8481},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 345, col: 36, offset: 8481},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 345, col: 39, offset: 8484},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 345, col: 51, offset: 8496},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 345, col: 54, offset: 8499},
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
			pos:  position{line: 364, col: 1, offset: 9057},
			expr: &actionExpr{
				pos: position{line: 364, col: 13, offset: 9069},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 364, col: 13, offset: 9069},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 364, col: 13, offset: 9069},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 364, col: 15, offset: 9071},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 21, offset: 9077},
								name: "Value",
							},
						},
						&labeledExpr{
							pos:   position{line: 364, col: 27, offset: 9083},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 364, col: 32, offset: 9088},
								expr: &seqExpr{
									pos: position{line: 364, col: 33, offset: 9089},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 364, col: 33, offset: 9089},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 36, offset: 9092},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 49, offset: 9105},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 52, offset: 9108},
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
			pos:  position{line: 382, col: 1, offset: 9662},
			expr: &choiceExpr{
				pos: position{line: 382, col: 12, offset: 9673},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 382, col: 12, offset: 9673},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 382, col: 30, offset: 9691},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 382, col: 49, offset: 9710},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 382, col: 64, offset: 9725},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 384, col: 1, offset: 9738},
			expr: &actionExpr{
				pos: position{line: 384, col: 19, offset: 9756},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 384, col: 21, offset: 9758},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 384, col: 21, offset: 9758},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 384, col: 29, offset: 9766},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 384, col: 36, offset: 9773},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 388, col: 1, offset: 9872},
			expr: &actionExpr{
				pos: position{line: 388, col: 20, offset: 9891},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 388, col: 22, offset: 9893},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 388, col: 22, offset: 9893},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 388, col: 29, offset: 9900},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 388, col: 36, offset: 9907},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 388, col: 42, offset: 9913},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 388, col: 48, offset: 9919},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 388, col: 56, offset: 9927},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 392, col: 1, offset: 10033},
			expr: &choiceExpr{
				pos: position{line: 392, col: 16, offset: 10048},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 392, col: 16, offset: 10048},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 392, col: 18, offset: 10050},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 392, col: 18, offset: 10050},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 392, col: 25, offset: 10057},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 395, col: 3, offset: 10163},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 395, col: 5, offset: 10165},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 395, col: 5, offset: 10165},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 395, col: 11, offset: 10171},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 395, col: 17, offset: 10177},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 398, col: 3, offset: 10280},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 398, col: 3, offset: 10280},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 402, col: 1, offset: 10384},
			expr: &choiceExpr{
				pos: position{line: 402, col: 15, offset: 10398},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 402, col: 15, offset: 10398},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 402, col: 17, offset: 10400},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 402, col: 17, offset: 10400},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 402, col: 24, offset: 10407},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 405, col: 3, offset: 10513},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 405, col: 5, offset: 10515},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 405, col: 5, offset: 10515},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 405, col: 11, offset: 10521},
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
			pos:  position{line: 409, col: 1, offset: 10623},
			expr: &choiceExpr{
				pos: position{line: 409, col: 9, offset: 10631},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 409, col: 9, offset: 10631},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 409, col: 22, offset: 10644},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 409, col: 22, offset: 10644},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 409, col: 24, offset: 10646},
								name: "Const",
							},
						},
					},
					&actionExpr{
						pos: position{line: 412, col: 3, offset: 10682},
						run: (*parser).callonValue6,
						expr: &seqExpr{
							pos: position{line: 412, col: 3, offset: 10682},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 412, col: 3, offset: 10682},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 412, col: 7, offset: 10686},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 412, col: 12, offset: 10691},
										name: "Expr",
									},
								},
								&litMatcher{
									pos:        position{line: 412, col: 17, offset: 10696},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 416, col: 1, offset: 10732},
			expr: &actionExpr{
				pos: position{line: 416, col: 14, offset: 10745},
				run: (*parser).callonIdentifier1,
				expr: &choiceExpr{
					pos: position{line: 416, col: 15, offset: 10746},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 416, col: 15, offset: 10746},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 416, col: 15, offset: 10746},
									expr: &charClassMatcher{
										pos:        position{line: 416, col: 15, offset: 10746},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 416, col: 22, offset: 10753},
									expr: &charClassMatcher{
										pos:        position{line: 416, col: 22, offset: 10753},
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
							pos:        position{line: 416, col: 38, offset: 10769},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 420, col: 1, offset: 10870},
			expr: &choiceExpr{
				pos: position{line: 420, col: 9, offset: 10878},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 420, col: 9, offset: 10878},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 420, col: 9, offset: 10878},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 420, col: 9, offset: 10878},
									expr: &litMatcher{
										pos:        position{line: 420, col: 9, offset: 10878},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 420, col: 14, offset: 10883},
									expr: &charClassMatcher{
										pos:        position{line: 420, col: 14, offset: 10883},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 420, col: 21, offset: 10890},
									expr: &litMatcher{
										pos:        position{line: 420, col: 22, offset: 10891},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 427, col: 3, offset: 11067},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 427, col: 3, offset: 11067},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 427, col: 3, offset: 11067},
									expr: &litMatcher{
										pos:        position{line: 427, col: 3, offset: 11067},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 427, col: 8, offset: 11072},
									expr: &charClassMatcher{
										pos:        position{line: 427, col: 8, offset: 11072},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 427, col: 15, offset: 11079},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 427, col: 19, offset: 11083},
									expr: &charClassMatcher{
										pos:        position{line: 427, col: 19, offset: 11083},
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
						pos:        position{line: 434, col: 3, offset: 11273},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 434, col: 12, offset: 11282},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 434, col: 12, offset: 11282},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 440, col: 3, offset: 11483},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 440, col: 3, offset: 11483},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 443, col: 3, offset: 11546},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 443, col: 3, offset: 11546},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 443, col: 3, offset: 11546},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 443, col: 7, offset: 11550},
									expr: &seqExpr{
										pos: position{line: 443, col: 8, offset: 11551},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 443, col: 8, offset: 11551},
												expr: &ruleRefExpr{
													pos:  position{line: 443, col: 9, offset: 11552},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 443, col: 21, offset: 11564,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 25, offset: 11568},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 450, col: 3, offset: 11752},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 450, col: 3, offset: 11752},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 450, col: 3, offset: 11752},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 450, col: 7, offset: 11756},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 450, col: 12, offset: 11761},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 450, col: 12, offset: 11761},
												expr: &ruleRefExpr{
													pos:  position{line: 450, col: 13, offset: 11762},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 450, col: 25, offset: 11774,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 450, col: 28, offset: 11777},
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
			pos:  position{line: 454, col: 1, offset: 11868},
			expr: &charClassMatcher{
				pos:        position{line: 454, col: 15, offset: 11882},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 456, col: 1, offset: 11898},
			expr: &choiceExpr{
				pos: position{line: 456, col: 18, offset: 11915},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 456, col: 18, offset: 11915},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 456, col: 37, offset: 11934},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 458, col: 1, offset: 11949},
			expr: &charClassMatcher{
				pos:        position{line: 458, col: 20, offset: 11968},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 460, col: 1, offset: 11981},
			expr: &charClassMatcher{
				pos:        position{line: 460, col: 16, offset: 11996},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 462, col: 1, offset: 12003},
			expr: &charClassMatcher{
				pos:        position{line: 462, col: 23, offset: 12025},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 464, col: 1, offset: 12032},
			expr: &charClassMatcher{
				pos:        position{line: 464, col: 12, offset: 12043},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 466, col: 1, offset: 12054},
			expr: &oneOrMoreExpr{
				pos: position{line: 466, col: 22, offset: 12075},
				expr: &charClassMatcher{
					pos:        position{line: 466, col: 22, offset: 12075},
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
			pos:         position{line: 468, col: 1, offset: 12087},
			expr: &zeroOrMoreExpr{
				pos: position{line: 468, col: 18, offset: 12104},
				expr: &charClassMatcher{
					pos:        position{line: 468, col: 18, offset: 12104},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 470, col: 1, offset: 12116},
			expr: &notExpr{
				pos: position{line: 470, col: 7, offset: 12122},
				expr: &anyMatcher{
					line: 470, col: 8, offset: 12123,
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

func (c *current) onCompoundExpr1(op, rest interface{}) (interface{}, error) {
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

func (c *current) onValue3(v interface{}) (interface{}, error) {
	return v.(Ast), nil
}

func (p *parser) callonValue3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue3(stack["v"])
}

func (c *current) onValue6(expr interface{}) (interface{}, error) {
	return expr.(Ast), nil
}

func (p *parser) callonValue6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue6(stack["expr"])
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
