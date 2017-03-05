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
	input := `-4 + -50.0 * 2 + 66.7 + "hi {boo}" / 'n'`
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
			name: "Expr",
			pos:  position{line: 77, col: 1, offset: 1365},
			expr: &actionExpr{
				pos: position{line: 77, col: 8, offset: 1372},
				run: (*parser).callonExpr1,
				expr: &seqExpr{
					pos: position{line: 77, col: 8, offset: 1372},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 77, col: 8, offset: 1372},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 11, offset: 1375},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 77, col: 17, offset: 1381},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 77, col: 22, offset: 1386},
								expr: &ruleRefExpr{
									pos:  position{line: 77, col: 23, offset: 1387},
									name: "BinOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 31, offset: 1395},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 33, offset: 1397},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 90, col: 1, offset: 1784},
			expr: &choiceExpr{
				pos: position{line: 90, col: 9, offset: 1792},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 90, col: 9, offset: 1792},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 20, offset: 1803},
						name: "BinOpHigh",
					},
				},
			},
		},
		{
			name: "BinOpLow",
			pos:  position{line: 92, col: 1, offset: 1814},
			expr: &actionExpr{
				pos: position{line: 92, col: 12, offset: 1825},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 92, col: 12, offset: 1825},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 92, col: 12, offset: 1825},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 92, col: 14, offset: 1827},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 20, offset: 1833},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 92, col: 30, offset: 1843},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 92, col: 35, offset: 1848},
								expr: &seqExpr{
									pos: position{line: 92, col: 36, offset: 1849},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 92, col: 36, offset: 1849},
											name: "__",
										},
										&labeledExpr{
											pos:   position{line: 92, col: 39, offset: 1852},
											label: "op",
											expr: &ruleRefExpr{
												pos:  position{line: 92, col: 42, offset: 1855},
												name: "OperatorLow",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 92, col: 54, offset: 1867},
											name: "__",
										},
										&labeledExpr{
											pos:   position{line: 92, col: 57, offset: 1870},
											label: "second",
											expr: &ruleRefExpr{
												pos:  position{line: 92, col: 64, offset: 1877},
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
		},
		{
			name: "BinOpHigh",
			pos:  position{line: 113, col: 1, offset: 2501},
			expr: &actionExpr{
				pos: position{line: 113, col: 13, offset: 2513},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 113, col: 13, offset: 2513},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 113, col: 13, offset: 2513},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 113, col: 15, offset: 2515},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 21, offset: 2521},
								name: "Value",
							},
						},
						&labeledExpr{
							pos:   position{line: 113, col: 27, offset: 2527},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 113, col: 32, offset: 2532},
								expr: &seqExpr{
									pos: position{line: 113, col: 33, offset: 2533},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 113, col: 33, offset: 2533},
											name: "__",
										},
										&labeledExpr{
											pos:   position{line: 113, col: 36, offset: 2536},
											label: "op",
											expr: &ruleRefExpr{
												pos:  position{line: 113, col: 39, offset: 2539},
												name: "OperatorHigh",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 113, col: 52, offset: 2552},
											name: "__",
										},
										&labeledExpr{
											pos:   position{line: 113, col: 55, offset: 2555},
											label: "second",
											expr: &ruleRefExpr{
												pos:  position{line: 113, col: 62, offset: 2562},
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
		},
		{
			name: "Operator",
			pos:  position{line: 130, col: 1, offset: 3046},
			expr: &choiceExpr{
				pos: position{line: 130, col: 12, offset: 3057},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 12, offset: 3057},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 27, offset: 3072},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 132, col: 1, offset: 3085},
			expr: &actionExpr{
				pos: position{line: 132, col: 16, offset: 3100},
				run: (*parser).callonOperatorHigh1,
				expr: &choiceExpr{
					pos: position{line: 132, col: 18, offset: 3102},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 132, col: 18, offset: 3102},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 24, offset: 3108},
							val:        "/",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 30, offset: 3114},
							val:        "/.",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 37, offset: 3121},
							val:        "*.",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 136, col: 1, offset: 3210},
			expr: &actionExpr{
				pos: position{line: 136, col: 15, offset: 3224},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 136, col: 17, offset: 3226},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 136, col: 17, offset: 3226},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 136, col: 23, offset: 3232},
							val:        "-",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 136, col: 29, offset: 3238},
							val:        "+.",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 136, col: 36, offset: 3245},
							val:        "-.",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 140, col: 1, offset: 3334},
			expr: &choiceExpr{
				pos: position{line: 140, col: 9, offset: 3342},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 140, col: 9, offset: 3342},
						run: (*parser).callonValue2,
						expr: &labeledExpr{
							pos:   position{line: 140, col: 9, offset: 3342},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 140, col: 11, offset: 3344},
								name: "Const",
							},
						},
					},
					&actionExpr{
						pos: position{line: 143, col: 3, offset: 3380},
						run: (*parser).callonValue5,
						expr: &seqExpr{
							pos: position{line: 143, col: 3, offset: 3380},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 143, col: 3, offset: 3380},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 143, col: 7, offset: 3384},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 143, col: 12, offset: 3389},
										name: "Expr",
									},
								},
								&litMatcher{
									pos:        position{line: 143, col: 17, offset: 3394},
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
			name: "Const",
			pos:  position{line: 147, col: 1, offset: 3430},
			expr: &choiceExpr{
				pos: position{line: 147, col: 9, offset: 3438},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 147, col: 9, offset: 3438},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 147, col: 9, offset: 3438},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 147, col: 9, offset: 3438},
									expr: &litMatcher{
										pos:        position{line: 147, col: 9, offset: 3438},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 147, col: 14, offset: 3443},
									expr: &charClassMatcher{
										pos:        position{line: 147, col: 14, offset: 3443},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 147, col: 21, offset: 3450},
									expr: &litMatcher{
										pos:        position{line: 147, col: 22, offset: 3451},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 154, col: 3, offset: 3622},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 154, col: 3, offset: 3622},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 154, col: 3, offset: 3622},
									expr: &litMatcher{
										pos:        position{line: 154, col: 3, offset: 3622},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 154, col: 8, offset: 3627},
									expr: &charClassMatcher{
										pos:        position{line: 154, col: 8, offset: 3627},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 154, col: 15, offset: 3634},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 154, col: 19, offset: 3638},
									expr: &charClassMatcher{
										pos:        position{line: 154, col: 19, offset: 3638},
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
						pos:        position{line: 161, col: 3, offset: 3823},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 161, col: 12, offset: 3832},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 161, col: 12, offset: 3832},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 167, col: 3, offset: 4023},
						run: (*parser).callonConst22,
						expr: &seqExpr{
							pos: position{line: 167, col: 3, offset: 4023},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 167, col: 3, offset: 4023},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 167, col: 7, offset: 4027},
									expr: &seqExpr{
										pos: position{line: 167, col: 8, offset: 4028},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 167, col: 8, offset: 4028},
												expr: &ruleRefExpr{
													pos:  position{line: 167, col: 9, offset: 4029},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 167, col: 21, offset: 4041,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 167, col: 25, offset: 4045},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 174, col: 3, offset: 4224},
						run: (*parser).callonConst31,
						expr: &seqExpr{
							pos: position{line: 174, col: 3, offset: 4224},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 174, col: 3, offset: 4224},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 174, col: 7, offset: 4228},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 174, col: 12, offset: 4233},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 174, col: 12, offset: 4233},
												expr: &ruleRefExpr{
													pos:  position{line: 174, col: 13, offset: 4234},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 174, col: 25, offset: 4246,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 174, col: 28, offset: 4249},
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
			pos:  position{line: 178, col: 1, offset: 4335},
			expr: &charClassMatcher{
				pos:        position{line: 178, col: 15, offset: 4349},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 180, col: 1, offset: 4365},
			expr: &choiceExpr{
				pos: position{line: 180, col: 18, offset: 4382},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 180, col: 18, offset: 4382},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 180, col: 37, offset: 4401},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 182, col: 1, offset: 4416},
			expr: &charClassMatcher{
				pos:        position{line: 182, col: 20, offset: 4435},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 184, col: 1, offset: 4448},
			expr: &charClassMatcher{
				pos:        position{line: 184, col: 16, offset: 4463},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 186, col: 1, offset: 4470},
			expr: &charClassMatcher{
				pos:        position{line: 186, col: 23, offset: 4492},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 188, col: 1, offset: 4499},
			expr: &charClassMatcher{
				pos:        position{line: 188, col: 12, offset: 4510},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 190, col: 1, offset: 4521},
			expr: &oneOrMoreExpr{
				pos: position{line: 190, col: 22, offset: 4542},
				expr: &charClassMatcher{
					pos:        position{line: 190, col: 22, offset: 4542},
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
			pos:         position{line: 192, col: 1, offset: 4554},
			expr: &zeroOrMoreExpr{
				pos: position{line: 192, col: 18, offset: 4571},
				expr: &charClassMatcher{
					pos:        position{line: 192, col: 18, offset: 4571},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 194, col: 1, offset: 4583},
			expr: &notExpr{
				pos: position{line: 194, col: 7, offset: 4589},
				expr: &anyMatcher{
					line: 194, col: 8, offset: 4590,
				},
			},
		},
	},
}

func (c *current) onExpr1(op, rest interface{}) (interface{}, error) {
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

func (p *parser) callonExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr1(stack["op"], stack["rest"])
}

func (c *current) onBinOpLow1(first, rest interface{}) (interface{}, error) {
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return Ast{Type: "Expr", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return first.(Ast), nil
	}

	//ast := Ast{Type:"BinOp", Subvalues:[]Ast{first.(Ast), op.(Ast), second.(Ast)},
	//ValueType: CONTAINER}
	//return ast, nil
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
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return Ast{Type: "Expr", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return first.(Ast), nil
	}
}

func (p *parser) callonBinOpHigh1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpHigh1(stack["first"], stack["rest"])
}

func (c *current) onOperatorHigh1() (interface{}, error) {
	return Ast{Type: "Op", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh1()
}

func (c *current) onOperatorLow1() (interface{}, error) {
	return Ast{Type: "Op", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow1()
}

func (c *current) onValue2(v interface{}) (interface{}, error) {
	return v.(Ast), nil
}

func (p *parser) callonValue2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue2(stack["v"])
}

func (c *current) onValue5(expr interface{}) (interface{}, error) {
	return expr.(Ast), nil
}

func (p *parser) callonValue5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue5(stack["expr"])
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
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return Ast{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonConst22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst22()
}

func (c *current) onConst31(val interface{}) (interface{}, error) {
	return Ast{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst31(stack["val"])
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
