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

type Ast struct {
	Type        string
	StringValue string
	CharValue   rune
	BoolValue   bool
	IntValue    int
	FloatValue  float64
	ValueType   ValueType
	Subvalues   []Ast
}

type Printable interface {
	Print(indent int) string
}

func main() {
	input := `
let a = -4 + 55.0 > 99 or "hi" ++ 'm';
let b = "cheese" ++ "ham";
let c = 5.0;
34 + 5;
let go = fun num -> {
    num + 5;
};
`
	fmt.Println(input)
	r := strings.NewReader(input)
	result, err := ParseReader("", r)
	ast := result.(Ast)
	fmt.Println("=", ast.Print(0))
	fmt.Println(err)
}

func (a *Ast) String() string {
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

func (a *Ast) Print(indent int) string {
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
			pos:  position{line: 85, col: 1, offset: 1458},
			expr: &actionExpr{
				pos: position{line: 85, col: 10, offset: 1467},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 85, col: 10, offset: 1467},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 85, col: 10, offset: 1467},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 85, col: 12, offset: 1469},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 17, offset: 1474},
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 27, offset: 1484},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 85, col: 29, offset: 1486},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 85, col: 34, offset: 1491},
								expr: &ruleRefExpr{
									pos:  position{line: 85, col: 35, offset: 1492},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 47, offset: 1504},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 49, offset: 1506},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 98, col: 1, offset: 1901},
			expr: &choiceExpr{
				pos: position{line: 98, col: 13, offset: 1913},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 13, offset: 1913},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 98, col: 13, offset: 1913},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 13, offset: 1913},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 19, offset: 1919},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 22, offset: 1922},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 24, offset: 1924},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 35, offset: 1935},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 37, offset: 1937},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 41, offset: 1941},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 43, offset: 1943},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 48, offset: 1948},
										name: "Expr",
									},
								},
								&litMatcher{
									pos:        position{line: 98, col: 53, offset: 1953},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 57, offset: 1957},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 101, col: 5, offset: 2116},
						run: (*parser).callonStatement15,
						expr: &seqExpr{
							pos: position{line: 101, col: 5, offset: 2116},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 101, col: 5, offset: 2116},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 10, offset: 2121},
										name: "Expr",
									},
								},
								&litMatcher{
									pos:        position{line: 101, col: 15, offset: 2126},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 19, offset: 2130},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 105, col: 1, offset: 2164},
			expr: &choiceExpr{
				pos: position{line: 105, col: 8, offset: 2171},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 105, col: 8, offset: 2171},
						run: (*parser).callonExpr2,
						expr: &seqExpr{
							pos: position{line: 105, col: 8, offset: 2171},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 105, col: 8, offset: 2171},
									val:        "fun",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 105, col: 14, offset: 2177},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 105, col: 17, offset: 2180},
									label: "ids",
									expr: &zeroOrMoreExpr{
										pos: position{line: 105, col: 21, offset: 2184},
										expr: &ruleRefExpr{
											pos:  position{line: 105, col: 22, offset: 2185},
											name: "Identifier",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 105, col: 35, offset: 2198},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 105, col: 37, offset: 2200},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 105, col: 42, offset: 2205},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 105, col: 44, offset: 2207},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 105, col: 48, offset: 2211},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 105, col: 50, offset: 2213},
									label: "statements",
									expr: &oneOrMoreExpr{
										pos: position{line: 105, col: 61, offset: 2224},
										expr: &ruleRefExpr{
											pos:  position{line: 105, col: 62, offset: 2225},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 105, col: 74, offset: 2237},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 105, col: 76, offset: 2239},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 105, col: 80, offset: 2243},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 116, col: 3, offset: 2558},
						run: (*parser).callonExpr20,
						expr: &seqExpr{
							pos: position{line: 116, col: 3, offset: 2558},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 116, col: 3, offset: 2558},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 116, col: 5, offset: 2560},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 116, col: 8, offset: 2563},
										name: "BinOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 116, col: 14, offset: 2569},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 116, col: 19, offset: 2574},
										expr: &ruleRefExpr{
											pos:  position{line: 116, col: 20, offset: 2575},
											name: "BinOp",
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
			name: "BinOp",
			pos:  position{line: 129, col: 1, offset: 2966},
			expr: &choiceExpr{
				pos: position{line: 129, col: 9, offset: 2974},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 129, col: 9, offset: 2974},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 21, offset: 2986},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 37, offset: 3002},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 48, offset: 3013},
						name: "BinOpHigh",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 131, col: 1, offset: 3024},
			expr: &actionExpr{
				pos: position{line: 131, col: 13, offset: 3036},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 131, col: 13, offset: 3036},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 131, col: 13, offset: 3036},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 131, col: 15, offset: 3038},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 131, col: 21, offset: 3044},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 131, col: 35, offset: 3058},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 131, col: 40, offset: 3063},
								expr: &seqExpr{
									pos: position{line: 131, col: 41, offset: 3064},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 131, col: 41, offset: 3064},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 131, col: 44, offset: 3067},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 131, col: 60, offset: 3083},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 131, col: 63, offset: 3086},
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
			pos:  position{line: 149, col: 1, offset: 3639},
			expr: &actionExpr{
				pos: position{line: 149, col: 17, offset: 3655},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 149, col: 17, offset: 3655},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 149, col: 17, offset: 3655},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 149, col: 19, offset: 3657},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 25, offset: 3663},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 149, col: 34, offset: 3672},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 149, col: 39, offset: 3677},
								expr: &seqExpr{
									pos: position{line: 149, col: 40, offset: 3678},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 149, col: 40, offset: 3678},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 149, col: 43, offset: 3681},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 149, col: 60, offset: 3698},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 149, col: 63, offset: 3701},
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
			pos:  position{line: 168, col: 1, offset: 4250},
			expr: &actionExpr{
				pos: position{line: 168, col: 12, offset: 4261},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 168, col: 12, offset: 4261},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 168, col: 12, offset: 4261},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 168, col: 14, offset: 4263},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 168, col: 20, offset: 4269},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 168, col: 30, offset: 4279},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 168, col: 35, offset: 4284},
								expr: &seqExpr{
									pos: position{line: 168, col: 36, offset: 4285},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 168, col: 36, offset: 4285},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 168, col: 39, offset: 4288},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 168, col: 51, offset: 4300},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 168, col: 54, offset: 4303},
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
			pos:  position{line: 187, col: 1, offset: 4853},
			expr: &actionExpr{
				pos: position{line: 187, col: 13, offset: 4865},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 187, col: 13, offset: 4865},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 187, col: 13, offset: 4865},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 187, col: 15, offset: 4867},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 21, offset: 4873},
								name: "Value",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 27, offset: 4879},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 187, col: 32, offset: 4884},
								expr: &seqExpr{
									pos: position{line: 187, col: 33, offset: 4885},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 187, col: 33, offset: 4885},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 187, col: 36, offset: 4888},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 187, col: 49, offset: 4901},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 187, col: 52, offset: 4904},
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
			pos:  position{line: 205, col: 1, offset: 5449},
			expr: &choiceExpr{
				pos: position{line: 205, col: 12, offset: 5460},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 205, col: 12, offset: 5460},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 205, col: 30, offset: 5478},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 205, col: 49, offset: 5497},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 205, col: 64, offset: 5512},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 207, col: 1, offset: 5525},
			expr: &actionExpr{
				pos: position{line: 207, col: 19, offset: 5543},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 207, col: 21, offset: 5545},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 207, col: 21, offset: 5545},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 207, col: 29, offset: 5553},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 207, col: 36, offset: 5560},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 211, col: 1, offset: 5654},
			expr: &actionExpr{
				pos: position{line: 211, col: 20, offset: 5673},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 211, col: 22, offset: 5675},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 211, col: 22, offset: 5675},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 211, col: 29, offset: 5682},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 211, col: 36, offset: 5689},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 211, col: 42, offset: 5695},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 211, col: 48, offset: 5701},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 211, col: 56, offset: 5709},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 215, col: 1, offset: 5810},
			expr: &choiceExpr{
				pos: position{line: 215, col: 16, offset: 5825},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 215, col: 16, offset: 5825},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 215, col: 18, offset: 5827},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 215, col: 18, offset: 5827},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 215, col: 25, offset: 5834},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 218, col: 3, offset: 5935},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 218, col: 5, offset: 5937},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 218, col: 5, offset: 5937},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 218, col: 11, offset: 5943},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 218, col: 17, offset: 5949},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 221, col: 3, offset: 6047},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 221, col: 3, offset: 6047},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 225, col: 1, offset: 6146},
			expr: &choiceExpr{
				pos: position{line: 225, col: 15, offset: 6160},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 225, col: 15, offset: 6160},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 225, col: 17, offset: 6162},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 225, col: 17, offset: 6162},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 225, col: 24, offset: 6169},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 228, col: 3, offset: 6270},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 228, col: 5, offset: 6272},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 228, col: 5, offset: 6272},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 228, col: 11, offset: 6278},
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
			pos:  position{line: 232, col: 1, offset: 6375},
			expr: &choiceExpr{
				pos: position{line: 232, col: 9, offset: 6383},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 232, col: 9, offset: 6383},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 232, col: 22, offset: 6396},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 232, col: 22, offset: 6396},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 24, offset: 6398},
								name: "Const",
							},
						},
					},
					&actionExpr{
						pos: position{line: 235, col: 3, offset: 6434},
						run: (*parser).callonValue6,
						expr: &seqExpr{
							pos: position{line: 235, col: 3, offset: 6434},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 235, col: 3, offset: 6434},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 235, col: 7, offset: 6438},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 235, col: 12, offset: 6443},
										name: "Expr",
									},
								},
								&litMatcher{
									pos:        position{line: 235, col: 17, offset: 6448},
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
			pos:  position{line: 239, col: 1, offset: 6484},
			expr: &actionExpr{
				pos: position{line: 239, col: 14, offset: 6497},
				run: (*parser).callonIdentifier1,
				expr: &choiceExpr{
					pos: position{line: 239, col: 15, offset: 6498},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 239, col: 15, offset: 6498},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 239, col: 15, offset: 6498},
									expr: &charClassMatcher{
										pos:        position{line: 239, col: 15, offset: 6498},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 239, col: 22, offset: 6505},
									expr: &charClassMatcher{
										pos:        position{line: 239, col: 22, offset: 6505},
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
							pos:        position{line: 239, col: 38, offset: 6521},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 243, col: 1, offset: 6617},
			expr: &choiceExpr{
				pos: position{line: 243, col: 9, offset: 6625},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 243, col: 9, offset: 6625},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 243, col: 9, offset: 6625},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 243, col: 9, offset: 6625},
									expr: &litMatcher{
										pos:        position{line: 243, col: 9, offset: 6625},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 243, col: 14, offset: 6630},
									expr: &charClassMatcher{
										pos:        position{line: 243, col: 14, offset: 6630},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 243, col: 21, offset: 6637},
									expr: &litMatcher{
										pos:        position{line: 243, col: 22, offset: 6638},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 250, col: 3, offset: 6809},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 250, col: 3, offset: 6809},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 250, col: 3, offset: 6809},
									expr: &litMatcher{
										pos:        position{line: 250, col: 3, offset: 6809},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 250, col: 8, offset: 6814},
									expr: &charClassMatcher{
										pos:        position{line: 250, col: 8, offset: 6814},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 250, col: 15, offset: 6821},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 250, col: 19, offset: 6825},
									expr: &charClassMatcher{
										pos:        position{line: 250, col: 19, offset: 6825},
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
						pos:        position{line: 257, col: 3, offset: 7010},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 257, col: 12, offset: 7019},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 257, col: 12, offset: 7019},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 263, col: 3, offset: 7210},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 263, col: 3, offset: 7210},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 266, col: 3, offset: 7268},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 266, col: 3, offset: 7268},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 266, col: 3, offset: 7268},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 266, col: 7, offset: 7272},
									expr: &seqExpr{
										pos: position{line: 266, col: 8, offset: 7273},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 266, col: 8, offset: 7273},
												expr: &ruleRefExpr{
													pos:  position{line: 266, col: 9, offset: 7274},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 266, col: 21, offset: 7286,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 266, col: 25, offset: 7290},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 273, col: 3, offset: 7469},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 273, col: 3, offset: 7469},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 273, col: 3, offset: 7469},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 273, col: 7, offset: 7473},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 273, col: 12, offset: 7478},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 273, col: 12, offset: 7478},
												expr: &ruleRefExpr{
													pos:  position{line: 273, col: 13, offset: 7479},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 273, col: 25, offset: 7491,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 273, col: 28, offset: 7494},
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
			pos:  position{line: 277, col: 1, offset: 7580},
			expr: &charClassMatcher{
				pos:        position{line: 277, col: 15, offset: 7594},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 279, col: 1, offset: 7610},
			expr: &choiceExpr{
				pos: position{line: 279, col: 18, offset: 7627},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 279, col: 18, offset: 7627},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 279, col: 37, offset: 7646},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 281, col: 1, offset: 7661},
			expr: &charClassMatcher{
				pos:        position{line: 281, col: 20, offset: 7680},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 283, col: 1, offset: 7693},
			expr: &charClassMatcher{
				pos:        position{line: 283, col: 16, offset: 7708},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 285, col: 1, offset: 7715},
			expr: &charClassMatcher{
				pos:        position{line: 285, col: 23, offset: 7737},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 287, col: 1, offset: 7744},
			expr: &charClassMatcher{
				pos:        position{line: 287, col: 12, offset: 7755},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 289, col: 1, offset: 7766},
			expr: &oneOrMoreExpr{
				pos: position{line: 289, col: 22, offset: 7787},
				expr: &charClassMatcher{
					pos:        position{line: 289, col: 22, offset: 7787},
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
			pos:         position{line: 291, col: 1, offset: 7799},
			expr: &zeroOrMoreExpr{
				pos: position{line: 291, col: 18, offset: 7816},
				expr: &charClassMatcher{
					pos:        position{line: 291, col: 18, offset: 7816},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 293, col: 1, offset: 7828},
			expr: &notExpr{
				pos: position{line: 293, col: 7, offset: 7834},
				expr: &anyMatcher{
					line: 293, col: 8, offset: 7835,
				},
			},
		},
	},
}

func (c *current) onModule1(expr, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{expr.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return Ast{Type: "Module", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return Ast{Type: "Module", Subvalues: []Ast{expr.(Ast)}, ValueType: CONTAINER}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["expr"], stack["rest"])
}

func (c *current) onStatement2(i, expr interface{}) (interface{}, error) {
	fmt.Printf("assignment: %s\n", string(c.text))
	return Ast{Type: "Assignment", Subvalues: []Ast{i.(Ast), expr.(Ast)}, ValueType: CONTAINER}, nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["i"], stack["expr"])
}

func (c *current) onStatement15(expr interface{}) (interface{}, error) {
	return expr.(Ast), nil
}

func (p *parser) callonStatement15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement15(stack["expr"])
}

func (c *current) onExpr2(ids, statements interface{}) (interface{}, error) {
	fmt.Println("func", string(c.text))
	subvalues := []Ast{}
	vals := statements.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	return Ast{Type: "Func", Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr2(stack["ids"], stack["statements"])
}

func (c *current) onExpr20(op, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{op.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return Ast{Type: "Expr", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return Ast{Type: "Expr", Subvalues: []Ast{op.(Ast)}, ValueType: CONTAINER}, nil
	}
}

func (p *parser) callonExpr20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr20(stack["op"], stack["rest"])
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
		return Ast{Type: "BinOp", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
		return Ast{Type: "BinOp", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
		return Ast{Type: "BinOp", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
		return Ast{Type: "BinOp", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
	return Ast{Type: "BoolOp", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorBoolean1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorBoolean1()
}

func (c *current) onOperatorEquality1() (interface{}, error) {
	return Ast{Type: "EqualityOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorEquality1()
}

func (c *current) onOperatorHigh2() (interface{}, error) {
	return Ast{Type: "FloatOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh2()
}

func (c *current) onOperatorHigh6() (interface{}, error) {
	return Ast{Type: "IntOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh6()
}

func (c *current) onOperatorHigh11() (interface{}, error) {
	return Ast{Type: "StringOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh11()
}

func (c *current) onOperatorLow2() (interface{}, error) {
	return Ast{Type: "FloatOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorLow2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow2()
}

func (c *current) onOperatorLow6() (interface{}, error) {
	return Ast{Type: "IntOperator", StringValue: string(c.text), ValueType: STRING}, nil
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
	return Ast{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
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
	return Ast{Type: "Integer", IntValue: val, ValueType: INT}, nil
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
	return Ast{Type: "Float", FloatValue: val, ValueType: FLOAT}, nil
}

func (p *parser) callonConst10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst10()
}

func (c *current) onConst20() (interface{}, error) {
	if string(c.text) == "True" {
		return Ast{Type: "Bool", BoolValue: true, ValueType: BOOL}, nil
	}
	return Ast{Type: "Bool", BoolValue: false, ValueType: BOOL}, nil
}

func (p *parser) callonConst20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst20()
}

func (c *current) onConst22() (interface{}, error) {
	return Ast{Type: "Nil", ValueType: NIL}, nil
}

func (p *parser) callonConst22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst22()
}

func (c *current) onConst24() (interface{}, error) {
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return Ast{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonConst24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst24()
}

func (c *current) onConst33(val interface{}) (interface{}, error) {
	return Ast{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
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
