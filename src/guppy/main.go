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

type PrintableValue interface {
	PrintValue() string
}

func (a *Ast) PrintValue() string {
	switch a.ValueType {
	case STRING:
		return a.StringValue
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
	str += fmt.Sprintf("%s %s:\n", a.Type, a.PrintValue())
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

func main() {
	input := `-4 + -50.0 * 2 + 66.7 + "hi {boo}" / 'nn'`
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
			name: "Expr",
			pos:  position{line: 79, col: 1, offset: 1340},
			expr: &actionExpr{
				pos: position{line: 79, col: 8, offset: 1347},
				run: (*parser).callonExpr1,
				expr: &seqExpr{
					pos: position{line: 79, col: 8, offset: 1347},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 79, col: 8, offset: 1347},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 79, col: 11, offset: 1350},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 79, col: 17, offset: 1356},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 79, col: 22, offset: 1361},
								expr: &ruleRefExpr{
									pos:  position{line: 79, col: 23, offset: 1362},
									name: "BinOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 79, col: 31, offset: 1370},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 79, col: 33, offset: 1372},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 92, col: 1, offset: 1759},
			expr: &choiceExpr{
				pos: position{line: 92, col: 9, offset: 1767},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 92, col: 9, offset: 1767},
						run: (*parser).callonBinOp2,
						expr: &seqExpr{
							pos: position{line: 92, col: 9, offset: 1767},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 92, col: 9, offset: 1767},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 11, offset: 1769},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 92, col: 17, offset: 1775},
										name: "Value",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 23, offset: 1781},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 26, offset: 1784},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 92, col: 29, offset: 1787},
										name: "Operator",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 38, offset: 1796},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 41, offset: 1799},
									label: "second",
									expr: &ruleRefExpr{
										pos:  position{line: 92, col: 48, offset: 1806},
										name: "Expr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 97, col: 3, offset: 1950},
						run: (*parser).callonBinOp13,
						expr: &seqExpr{
							pos: position{line: 97, col: 3, offset: 1950},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 97, col: 3, offset: 1950},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 97, col: 5, offset: 1952},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 97, col: 11, offset: 1958},
										name: "Value",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 97, col: 17, offset: 1964},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 97, col: 20, offset: 1967},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 97, col: 23, offset: 1970},
										name: "Operator",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 97, col: 32, offset: 1979},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 97, col: 35, offset: 1982},
									label: "second",
									expr: &ruleRefExpr{
										pos:  position{line: 97, col: 42, offset: 1989},
										name: "Value",
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
			pos:  position{line: 103, col: 1, offset: 2129},
			expr: &choiceExpr{
				pos: position{line: 103, col: 12, offset: 2140},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 103, col: 12, offset: 2140},
						run: (*parser).callonOperator2,
						expr: &choiceExpr{
							pos: position{line: 103, col: 14, offset: 2142},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 103, col: 14, offset: 2142},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 103, col: 20, offset: 2148},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 103, col: 26, offset: 2154},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 103, col: 33, offset: 2161},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 106, col: 3, offset: 2251},
						run: (*parser).callonOperator8,
						expr: &choiceExpr{
							pos: position{line: 106, col: 5, offset: 2253},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 106, col: 5, offset: 2253},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 106, col: 11, offset: 2259},
									val:        "-",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 106, col: 17, offset: 2265},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 106, col: 24, offset: 2272},
									val:        "-.",
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
			pos:  position{line: 110, col: 1, offset: 2361},
			expr: &actionExpr{
				pos: position{line: 110, col: 9, offset: 2369},
				run: (*parser).callonValue1,
				expr: &labeledExpr{
					pos:   position{line: 110, col: 9, offset: 2369},
					label: "v",
					expr: &ruleRefExpr{
						pos:  position{line: 110, col: 11, offset: 2371},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 114, col: 1, offset: 2406},
			expr: &choiceExpr{
				pos: position{line: 114, col: 9, offset: 2414},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 114, col: 9, offset: 2414},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 114, col: 9, offset: 2414},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 114, col: 9, offset: 2414},
									expr: &litMatcher{
										pos:        position{line: 114, col: 9, offset: 2414},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 114, col: 14, offset: 2419},
									expr: &charClassMatcher{
										pos:        position{line: 114, col: 14, offset: 2419},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 114, col: 21, offset: 2426},
									expr: &litMatcher{
										pos:        position{line: 114, col: 22, offset: 2427},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 121, col: 3, offset: 2598},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 121, col: 3, offset: 2598},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 121, col: 3, offset: 2598},
									expr: &litMatcher{
										pos:        position{line: 121, col: 3, offset: 2598},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 121, col: 8, offset: 2603},
									expr: &charClassMatcher{
										pos:        position{line: 121, col: 8, offset: 2603},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 121, col: 15, offset: 2610},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 121, col: 19, offset: 2614},
									expr: &charClassMatcher{
										pos:        position{line: 121, col: 19, offset: 2614},
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
						pos:        position{line: 128, col: 3, offset: 2799},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 128, col: 12, offset: 2808},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 128, col: 12, offset: 2808},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 134, col: 3, offset: 2999},
						run: (*parser).callonConst22,
						expr: &seqExpr{
							pos: position{line: 134, col: 3, offset: 2999},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 134, col: 3, offset: 2999},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 134, col: 7, offset: 3003},
									expr: &seqExpr{
										pos: position{line: 134, col: 8, offset: 3004},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 134, col: 8, offset: 3004},
												expr: &ruleRefExpr{
													pos:  position{line: 134, col: 9, offset: 3005},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 134, col: 21, offset: 3017,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 134, col: 25, offset: 3021},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 141, col: 3, offset: 3200},
						run: (*parser).callonConst31,
						expr: &seqExpr{
							pos: position{line: 141, col: 3, offset: 3200},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 141, col: 3, offset: 3200},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 141, col: 7, offset: 3204},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 141, col: 12, offset: 3209},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 141, col: 12, offset: 3209},
												expr: &ruleRefExpr{
													pos:  position{line: 141, col: 13, offset: 3210},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 141, col: 25, offset: 3222,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 141, col: 28, offset: 3225},
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
			pos:  position{line: 145, col: 1, offset: 3316},
			expr: &charClassMatcher{
				pos:        position{line: 145, col: 15, offset: 3330},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 147, col: 1, offset: 3346},
			expr: &choiceExpr{
				pos: position{line: 147, col: 18, offset: 3363},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 147, col: 18, offset: 3363},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 37, offset: 3382},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 149, col: 1, offset: 3397},
			expr: &charClassMatcher{
				pos:        position{line: 149, col: 20, offset: 3416},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 151, col: 1, offset: 3429},
			expr: &charClassMatcher{
				pos:        position{line: 151, col: 16, offset: 3444},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 153, col: 1, offset: 3451},
			expr: &charClassMatcher{
				pos:        position{line: 153, col: 23, offset: 3473},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 155, col: 1, offset: 3480},
			expr: &charClassMatcher{
				pos:        position{line: 155, col: 12, offset: 3491},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 157, col: 1, offset: 3502},
			expr: &oneOrMoreExpr{
				pos: position{line: 157, col: 22, offset: 3523},
				expr: &charClassMatcher{
					pos:        position{line: 157, col: 22, offset: 3523},
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
			pos:         position{line: 159, col: 1, offset: 3535},
			expr: &zeroOrMoreExpr{
				pos: position{line: 159, col: 18, offset: 3552},
				expr: &charClassMatcher{
					pos:        position{line: 159, col: 18, offset: 3552},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 161, col: 1, offset: 3564},
			expr: &notExpr{
				pos: position{line: 161, col: 7, offset: 3570},
				expr: &anyMatcher{
					line: 161, col: 8, offset: 3571,
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

func (c *current) onBinOp2(first, op, second interface{}) (interface{}, error) {
	ast := Ast{Type: "BinOp", Subvalues: []Ast{first.(Ast), op.(Ast), second.(Ast)},
		ValueType: CONTAINER}
	return ast, nil
}

func (p *parser) callonBinOp2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOp2(stack["first"], stack["op"], stack["second"])
}

func (c *current) onBinOp13(first, op, second interface{}) (interface{}, error) {
	ast := Ast{Type: "BinOp", Subvalues: []Ast{first.(Ast), op.(Ast), second.(Ast)},
		ValueType: CONTAINER}
	return ast, nil
}

func (p *parser) callonBinOp13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOp13(stack["first"], stack["op"], stack["second"])
}

func (c *current) onOperator2() (interface{}, error) {
	return Ast{Type: "Op", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperator2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperator2()
}

func (c *current) onOperator8() (interface{}, error) {
	return Ast{Type: "Op", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperator8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperator8()
}

func (c *current) onValue1(v interface{}) (interface{}, error) {
	return v.(Ast), nil
}

func (p *parser) callonValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue1(stack["v"])
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
	return Ast{Type: "String", StringValue: string(c.text), ValueType: STRING}, nil
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
