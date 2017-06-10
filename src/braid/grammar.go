package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

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
			pos:  position{line: 13, col: 1, offset: 143},
			expr: &actionExpr{
				pos: position{line: 13, col: 10, offset: 152},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 13, col: 10, offset: 152},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 13, col: 10, offset: 152},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 12, offset: 154},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 17, offset: 159},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 35, offset: 177},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 37, offset: 179},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 42, offset: 184},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 43, offset: 185},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 63, offset: 205},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 65, offset: 207},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 695},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 715},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 715},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 31, offset: 725},
						name: "FuncDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 735},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 747},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 747},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 757},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 768},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 781},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 32, col: 1, offset: 787},
			expr: &actionExpr{
				pos: position{line: 32, col: 11, offset: 797},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 32, col: 11, offset: 797},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 32, col: 11, offset: 797},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 32, col: 15, offset: 801},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 32, col: 23, offset: 809},
								expr: &seqExpr{
									pos: position{line: 32, col: 24, offset: 810},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 32, col: 24, offset: 810},
											expr: &ruleRefExpr{
												pos:  position{line: 32, col: 25, offset: 811},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 32, col: 37, offset: 823,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 32, col: 41, offset: 827},
							expr: &litMatcher{
								pos:        position{line: 32, col: 42, offset: 828},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 37, col: 1, offset: 978},
			expr: &choiceExpr{
				pos: position{line: 37, col: 14, offset: 991},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 37, col: 14, offset: 991},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 37, col: 14, offset: 991},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 37, col: 14, offset: 991},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 37, col: 16, offset: 993},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 22, offset: 999},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 25, offset: 1002},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 27, offset: 1004},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 37, col: 38, offset: 1015},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 37, col: 43, offset: 1020},
										expr: &seqExpr{
											pos: position{line: 37, col: 44, offset: 1021},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 37, col: 44, offset: 1021},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 37, col: 48, offset: 1025},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 37, col: 50, offset: 1027},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 63, offset: 1040},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 37, col: 65, offset: 1042},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 69, offset: 1046},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 71, offset: 1048},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 76, offset: 1053},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 81, offset: 1058},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 52, col: 1, offset: 1497},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 52, col: 1, offset: 1497},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 1, offset: 1497},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 3, offset: 1499},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 9, offset: 1505},
									name: "__",
								},
								&notExpr{
									pos: position{line: 52, col: 12, offset: 1508},
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 13, offset: 1509},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 56, col: 1, offset: 1617},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 56, col: 1, offset: 1617},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 56, col: 1, offset: 1617},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 56, col: 3, offset: 1619},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 56, col: 9, offset: 1625},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 56, col: 12, offset: 1628},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 56, col: 14, offset: 1630},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 56, col: 25, offset: 1641},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 56, col: 30, offset: 1646},
										expr: &seqExpr{
											pos: position{line: 56, col: 31, offset: 1647},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 56, col: 31, offset: 1647},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 56, col: 35, offset: 1651},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 56, col: 37, offset: 1653},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 56, col: 50, offset: 1666},
									name: "_",
								},
								&notExpr{
									pos: position{line: 56, col: 52, offset: 1668},
									expr: &litMatcher{
										pos:        position{line: 56, col: 53, offset: 1669},
										val:        "=",
										ignoreCase: false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FuncDefn",
			pos:  position{line: 60, col: 1, offset: 1763},
			expr: &actionExpr{
				pos: position{line: 60, col: 12, offset: 1774},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 60, col: 12, offset: 1774},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 60, col: 12, offset: 1774},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 60, col: 14, offset: 1776},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 20, offset: 1782},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 60, col: 23, offset: 1785},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 25, offset: 1787},
								name: "Assignable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 36, offset: 1798},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 60, col: 38, offset: 1800},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 42, offset: 1804},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 60, col: 44, offset: 1806},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 51, offset: 1813},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 60, col: 54, offset: 1816},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 60, col: 58, offset: 1820},
								expr: &seqExpr{
									pos: position{line: 60, col: 59, offset: 1821},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 60, col: 59, offset: 1821},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 60, col: 72, offset: 1834},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 77, offset: 1839},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 60, col: 79, offset: 1841},
							expr: &litMatcher{
								pos:        position{line: 60, col: 79, offset: 1841},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 85, offset: 1847},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 60, col: 87, offset: 1849},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 91, offset: 1853},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 60, col: 93, offset: 1855},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 60, col: 104, offset: 1866},
								expr: &ruleRefExpr{
									pos:  position{line: 60, col: 105, offset: 1867},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 117, offset: 1879},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 60, col: 119, offset: 1881},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 123, offset: 1885},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 83, col: 1, offset: 2560},
			expr: &actionExpr{
				pos: position{line: 83, col: 8, offset: 2567},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 83, col: 8, offset: 2567},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 83, col: 12, offset: 2571},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 83, col: 12, offset: 2571},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 83, col: 19, offset: 2578},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 83, col: 28, offset: 2587},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 88, col: 1, offset: 2681},
			expr: &choiceExpr{
				pos: position{line: 88, col: 10, offset: 2690},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 88, col: 10, offset: 2690},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 88, col: 10, offset: 2690},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 88, col: 10, offset: 2690},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 15, offset: 2695},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 18, offset: 2698},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 88, col: 23, offset: 2703},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 33, offset: 2713},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 88, col: 35, offset: 2715},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 39, offset: 2719},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 41, offset: 2721},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 88, col: 47, offset: 2727},
										expr: &ruleRefExpr{
											pos:  position{line: 88, col: 48, offset: 2728},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 60, offset: 2740},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 88, col: 62, offset: 2742},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 66, offset: 2746},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 88, col: 68, offset: 2748},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 75, offset: 2755},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 77, offset: 2757},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 88, col: 85, offset: 2765},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 100, col: 1, offset: 3095},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 100, col: 1, offset: 3095},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 100, col: 1, offset: 3095},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 6, offset: 3100},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 9, offset: 3103},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 14, offset: 3108},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 24, offset: 3118},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 26, offset: 3120},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 30, offset: 3124},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 32, offset: 3126},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 100, col: 38, offset: 3132},
										expr: &ruleRefExpr{
											pos:  position{line: 100, col: 39, offset: 3133},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 51, offset: 3145},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 53, offset: 3147},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 57, offset: 3151},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 59, offset: 3153},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 66, offset: 3160},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 68, offset: 3162},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 72, offset: 3166},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 74, offset: 3168},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 100, col: 80, offset: 3174},
										expr: &ruleRefExpr{
											pos:  position{line: 100, col: 81, offset: 3175},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 93, offset: 3187},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 95, offset: 3189},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 99, offset: 3193},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 119, col: 1, offset: 3694},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 119, col: 1, offset: 3694},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 119, col: 1, offset: 3694},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 6, offset: 3699},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 9, offset: 3702},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 119, col: 14, offset: 3707},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 24, offset: 3717},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 26, offset: 3719},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 30, offset: 3723},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 32, offset: 3725},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 119, col: 38, offset: 3731},
										expr: &ruleRefExpr{
											pos:  position{line: 119, col: 39, offset: 3732},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 119, col: 51, offset: 3744},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 55, offset: 3748},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Call",
			pos:  position{line: 132, col: 1, offset: 4045},
			expr: &choiceExpr{
				pos: position{line: 132, col: 8, offset: 4052},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 132, col: 8, offset: 4052},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 132, col: 8, offset: 4052},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 132, col: 8, offset: 4052},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 15, offset: 4059},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 132, col: 26, offset: 4070},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 132, col: 30, offset: 4074},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 33, offset: 4077},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 132, col: 46, offset: 4090},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 132, col: 56, offset: 4100},
										expr: &seqExpr{
											pos: position{line: 132, col: 57, offset: 4101},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 132, col: 57, offset: 4101},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 132, col: 60, offset: 4104},
													name: "Value",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 132, col: 68, offset: 4112},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 132, col: 72, offset: 4116},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 150, col: 1, offset: 4614},
						run: (*parser).callonCall16,
						expr: &seqExpr{
							pos: position{line: 150, col: 1, offset: 4614},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 150, col: 1, offset: 4614},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 150, col: 4, offset: 4617},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 150, col: 17, offset: 4630},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 150, col: 27, offset: 4640},
										expr: &seqExpr{
											pos: position{line: 150, col: 28, offset: 4641},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 150, col: 28, offset: 4641},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 150, col: 31, offset: 4644},
													name: "Value",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 150, col: 39, offset: 4652},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 150, col: 43, offset: 4656},
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
			pos:  position{line: 168, col: 1, offset: 5166},
			expr: &actionExpr{
				pos: position{line: 168, col: 16, offset: 5181},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 168, col: 16, offset: 5181},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 168, col: 16, offset: 5181},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 168, col: 18, offset: 5183},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 168, col: 21, offset: 5186},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 168, col: 27, offset: 5192},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 168, col: 32, offset: 5197},
								expr: &seqExpr{
									pos: position{line: 168, col: 33, offset: 5198},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 168, col: 33, offset: 5198},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 168, col: 36, offset: 5201},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 168, col: 45, offset: 5210},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 168, col: 48, offset: 5213},
											name: "BinOp",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 168, col: 56, offset: 5221},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 168, col: 58, offset: 5223},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 168, col: 62, offset: 5227},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 188, col: 1, offset: 5886},
			expr: &choiceExpr{
				pos: position{line: 188, col: 9, offset: 5894},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 188, col: 9, offset: 5894},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 21, offset: 5906},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 37, offset: 5922},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 48, offset: 5933},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 60, offset: 5945},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 190, col: 1, offset: 5958},
			expr: &actionExpr{
				pos: position{line: 190, col: 13, offset: 5970},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 190, col: 13, offset: 5970},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 190, col: 13, offset: 5970},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 190, col: 15, offset: 5972},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 190, col: 21, offset: 5978},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 190, col: 35, offset: 5992},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 190, col: 40, offset: 5997},
								expr: &seqExpr{
									pos: position{line: 190, col: 41, offset: 5998},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 190, col: 41, offset: 5998},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 190, col: 44, offset: 6001},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 190, col: 60, offset: 6017},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 190, col: 63, offset: 6020},
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
			pos:  position{line: 209, col: 1, offset: 6627},
			expr: &actionExpr{
				pos: position{line: 209, col: 17, offset: 6643},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 209, col: 17, offset: 6643},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 209, col: 17, offset: 6643},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 209, col: 19, offset: 6645},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 209, col: 25, offset: 6651},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 209, col: 34, offset: 6660},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 209, col: 39, offset: 6665},
								expr: &seqExpr{
									pos: position{line: 209, col: 40, offset: 6666},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 209, col: 40, offset: 6666},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 209, col: 43, offset: 6669},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 209, col: 60, offset: 6686},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 209, col: 63, offset: 6689},
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
			pos:  position{line: 229, col: 1, offset: 7294},
			expr: &actionExpr{
				pos: position{line: 229, col: 12, offset: 7305},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 229, col: 12, offset: 7305},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 229, col: 12, offset: 7305},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 229, col: 14, offset: 7307},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 20, offset: 7313},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 229, col: 30, offset: 7323},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 229, col: 35, offset: 7328},
								expr: &seqExpr{
									pos: position{line: 229, col: 36, offset: 7329},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 229, col: 36, offset: 7329},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 229, col: 39, offset: 7332},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 229, col: 51, offset: 7344},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 229, col: 54, offset: 7347},
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
			pos:  position{line: 249, col: 1, offset: 7949},
			expr: &actionExpr{
				pos: position{line: 249, col: 13, offset: 7961},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 249, col: 13, offset: 7961},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 249, col: 13, offset: 7961},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 249, col: 15, offset: 7963},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 249, col: 21, offset: 7969},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 249, col: 33, offset: 7981},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 249, col: 38, offset: 7986},
								expr: &seqExpr{
									pos: position{line: 249, col: 39, offset: 7987},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 249, col: 39, offset: 7987},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 249, col: 42, offset: 7990},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 249, col: 55, offset: 8003},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 249, col: 58, offset: 8006},
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
			pos:  position{line: 268, col: 1, offset: 8611},
			expr: &choiceExpr{
				pos: position{line: 268, col: 15, offset: 8625},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 268, col: 15, offset: 8625},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 268, col: 15, offset: 8625},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 268, col: 15, offset: 8625},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 268, col: 19, offset: 8629},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 268, col: 21, offset: 8631},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 268, col: 27, offset: 8637},
										name: "BinOp",
									},
								},
								&litMatcher{
									pos:        position{line: 268, col: 33, offset: 8643},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 271, col: 5, offset: 8793},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 273, col: 1, offset: 8800},
			expr: &choiceExpr{
				pos: position{line: 273, col: 12, offset: 8811},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 273, col: 12, offset: 8811},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 30, offset: 8829},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 49, offset: 8848},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 64, offset: 8863},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 275, col: 1, offset: 8876},
			expr: &actionExpr{
				pos: position{line: 275, col: 19, offset: 8894},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 275, col: 21, offset: 8896},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 275, col: 21, offset: 8896},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 29, offset: 8904},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 36, offset: 8911},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 279, col: 1, offset: 9010},
			expr: &actionExpr{
				pos: position{line: 279, col: 20, offset: 9029},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 279, col: 22, offset: 9031},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 279, col: 22, offset: 9031},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 279, col: 29, offset: 9038},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 279, col: 36, offset: 9045},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 279, col: 42, offset: 9051},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 279, col: 48, offset: 9057},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 279, col: 56, offset: 9065},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 283, col: 1, offset: 9171},
			expr: &choiceExpr{
				pos: position{line: 283, col: 16, offset: 9186},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 283, col: 16, offset: 9186},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 283, col: 18, offset: 9188},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 283, col: 18, offset: 9188},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 283, col: 25, offset: 9195},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 286, col: 3, offset: 9301},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 286, col: 5, offset: 9303},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 286, col: 5, offset: 9303},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 286, col: 11, offset: 9309},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 286, col: 17, offset: 9315},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 289, col: 3, offset: 9418},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 289, col: 3, offset: 9418},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 293, col: 1, offset: 9522},
			expr: &choiceExpr{
				pos: position{line: 293, col: 15, offset: 9536},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 293, col: 15, offset: 9536},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 293, col: 17, offset: 9538},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 293, col: 17, offset: 9538},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 293, col: 24, offset: 9545},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 296, col: 3, offset: 9651},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 296, col: 5, offset: 9653},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 296, col: 5, offset: 9653},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 296, col: 11, offset: 9659},
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
			pos:  position{line: 300, col: 1, offset: 9761},
			expr: &choiceExpr{
				pos: position{line: 300, col: 9, offset: 9769},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 300, col: 9, offset: 9769},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 300, col: 22, offset: 9782},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 300, col: 22, offset: 9782},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 300, col: 24, offset: 9784},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 304, col: 1, offset: 9819},
			expr: &choiceExpr{
				pos: position{line: 304, col: 14, offset: 9832},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 304, col: 14, offset: 9832},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 304, col: 29, offset: 9847},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 306, col: 1, offset: 9855},
			expr: &choiceExpr{
				pos: position{line: 306, col: 14, offset: 9868},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 306, col: 14, offset: 9868},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 29, offset: 9883},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 308, col: 1, offset: 9895},
			expr: &actionExpr{
				pos: position{line: 308, col: 16, offset: 9910},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 308, col: 16, offset: 9910},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 308, col: 16, offset: 9910},
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 17, offset: 9911},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 308, col: 27, offset: 9921},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 308, col: 27, offset: 9921},
									expr: &charClassMatcher{
										pos:        position{line: 308, col: 27, offset: 9921},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 308, col: 34, offset: 9928},
									expr: &charClassMatcher{
										pos:        position{line: 308, col: 34, offset: 9928},
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
			},
		},
		{
			name: "ModuleName",
			pos:  position{line: 312, col: 1, offset: 10039},
			expr: &actionExpr{
				pos: position{line: 312, col: 14, offset: 10052},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 312, col: 15, offset: 10053},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 312, col: 15, offset: 10053},
							expr: &charClassMatcher{
								pos:        position{line: 312, col: 15, offset: 10053},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 312, col: 22, offset: 10060},
							expr: &charClassMatcher{
								pos:        position{line: 312, col: 22, offset: 10060},
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
			pos:  position{line: 316, col: 1, offset: 10171},
			expr: &choiceExpr{
				pos: position{line: 316, col: 9, offset: 10179},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 316, col: 9, offset: 10179},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 316, col: 9, offset: 10179},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 316, col: 9, offset: 10179},
									expr: &litMatcher{
										pos:        position{line: 316, col: 9, offset: 10179},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 316, col: 14, offset: 10184},
									expr: &charClassMatcher{
										pos:        position{line: 316, col: 14, offset: 10184},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 316, col: 21, offset: 10191},
									expr: &litMatcher{
										pos:        position{line: 316, col: 22, offset: 10192},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 323, col: 3, offset: 10368},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 323, col: 3, offset: 10368},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 323, col: 3, offset: 10368},
									expr: &litMatcher{
										pos:        position{line: 323, col: 3, offset: 10368},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 323, col: 8, offset: 10373},
									expr: &charClassMatcher{
										pos:        position{line: 323, col: 8, offset: 10373},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 323, col: 15, offset: 10380},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 323, col: 19, offset: 10384},
									expr: &charClassMatcher{
										pos:        position{line: 323, col: 19, offset: 10384},
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
						pos:        position{line: 330, col: 3, offset: 10574},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 330, col: 12, offset: 10583},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 330, col: 12, offset: 10583},
							val:        "false",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 336, col: 3, offset: 10784},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 336, col: 3, offset: 10784},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 339, col: 3, offset: 10847},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 339, col: 3, offset: 10847},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 339, col: 3, offset: 10847},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 339, col: 7, offset: 10851},
									expr: &seqExpr{
										pos: position{line: 339, col: 8, offset: 10852},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 339, col: 8, offset: 10852},
												expr: &ruleRefExpr{
													pos:  position{line: 339, col: 9, offset: 10853},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 339, col: 21, offset: 10865,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 339, col: 25, offset: 10869},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 346, col: 3, offset: 11053},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 346, col: 3, offset: 11053},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 346, col: 3, offset: 11053},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 346, col: 7, offset: 11057},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 346, col: 12, offset: 11062},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 346, col: 12, offset: 11062},
												expr: &ruleRefExpr{
													pos:  position{line: 346, col: 13, offset: 11063},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 346, col: 25, offset: 11075,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 346, col: 28, offset: 11078},
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
			pos:  position{line: 350, col: 1, offset: 11169},
			expr: &actionExpr{
				pos: position{line: 350, col: 10, offset: 11178},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 350, col: 11, offset: 11179},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 354, col: 1, offset: 11280},
			expr: &seqExpr{
				pos: position{line: 354, col: 12, offset: 11291},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 354, col: 13, offset: 11292},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 354, col: 13, offset: 11292},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 21, offset: 11300},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 28, offset: 11307},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 37, offset: 11316},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 46, offset: 11325},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 55, offset: 11334},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 64, offset: 11343},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 354, col: 74, offset: 11353},
								val:        "mutable",
								ignoreCase: false,
							},
						},
					},
					&notExpr{
						pos: position{line: 354, col: 86, offset: 11365},
						expr: &oneOrMoreExpr{
							pos: position{line: 354, col: 87, offset: 11366},
							expr: &charClassMatcher{
								pos:        position{line: 354, col: 87, offset: 11366},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 356, col: 1, offset: 11374},
			expr: &charClassMatcher{
				pos:        position{line: 356, col: 15, offset: 11388},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 358, col: 1, offset: 11404},
			expr: &choiceExpr{
				pos: position{line: 358, col: 18, offset: 11421},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 358, col: 18, offset: 11421},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 358, col: 37, offset: 11440},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 360, col: 1, offset: 11455},
			expr: &charClassMatcher{
				pos:        position{line: 360, col: 20, offset: 11474},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 362, col: 1, offset: 11487},
			expr: &charClassMatcher{
				pos:        position{line: 362, col: 16, offset: 11502},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 364, col: 1, offset: 11509},
			expr: &charClassMatcher{
				pos:        position{line: 364, col: 23, offset: 11531},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 366, col: 1, offset: 11538},
			expr: &charClassMatcher{
				pos:        position{line: 366, col: 12, offset: 11549},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 368, col: 1, offset: 11560},
			expr: &oneOrMoreExpr{
				pos: position{line: 368, col: 22, offset: 11581},
				expr: &charClassMatcher{
					pos:        position{line: 368, col: 22, offset: 11581},
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
			pos:         position{line: 370, col: 1, offset: 11593},
			expr: &zeroOrMoreExpr{
				pos: position{line: 370, col: 18, offset: 11610},
				expr: &charClassMatcher{
					pos:        position{line: 370, col: 18, offset: 11610},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 372, col: 1, offset: 11622},
			expr: &notExpr{
				pos: position{line: 372, col: 7, offset: 11628},
				expr: &anyMatcher{
					line: 372, col: 8, offset: 11629,
				},
			},
		},
	},
}

func (c *current) onModule1(stat, rest interface{}) (interface{}, error) {
	//fmt.Println("beginning module")
	vals := rest.([]interface{})
	if len(vals) > 0 {
		//fmt.Println("multiple statements")
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

func (c *current) onComment1(comment interface{}) (interface{}, error) {
	//fmt.Println("comment:", string(c.text))
	return BasicAst{Type: "Comment", StringValue: string(c.text[1:]), ValueType: STRING}, nil
}

func (p *parser) callonComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment1(stack["comment"])
}

func (c *current) onAssignment2(i, rest, expr interface{}) (interface{}, error) {
	//fmt.Println("assignment:", string(c.text))
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

func (p *parser) callonAssignment2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment2(stack["i"], stack["rest"], stack["expr"])
}

func (c *current) onAssignment21() (interface{}, error) {
	return nil, errors.New("Variable name or '_' (unused result character) required here")
}

func (p *parser) callonAssignment21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment21()
}

func (c *current) onAssignment28(i, rest interface{}) (interface{}, error) {
	return nil, errors.New("When assigning a value to a variable, you must use '='")
}

func (p *parser) callonAssignment28() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment28(stack["i"], stack["rest"])
}

func (c *current) onFuncDefn1(i, ids, statements interface{}) (interface{}, error) {
	//fmt.Println(string(c.text))
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
	return Func{Name: i.(BasicAst).StringValue, Arguments: args, Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["i"], stack["ids"], stack["statements"])
}

func (c *current) onExpr1(ex interface{}) (interface{}, error) {
	//fmt.Printf("top-level expr: %s\n", string(c.text))
	return ex, nil
}

func (p *parser) callonExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr1(stack["ex"])
}

func (c *current) onIfExpr2(expr, thens, elseifs interface{}) (interface{}, error) {
	//fmt.Printf("if: %s\n", string(c.text))
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
	//fmt.Printf("if: %s\n", string(c.text))
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
	//fmt.Printf("if: %s\n", string(c.text))
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

func (c *current) onCall2(module, fn, arguments interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))

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
	//fmt.Println("call", string(c.text))
	//var mod BasicAst
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

	return Call{Module: nil, Function: fn.(Ast), Arguments: args, ValueType: CONTAINER}, nil
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
	if string(c.text) == "true" {
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
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()
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
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename:        filename,
		errs:            new(errList),
		data:            b,
		pt:              savepoint{position: position{line: 1}},
		recover:         true,
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make(map[string]struct{}),
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

	depth   int
	recover bool
	debug   bool

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

	// parse fail
	maxFailPos            position
	maxFailExpected       map[string]struct{}
	maxFailInvertExpected bool
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
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
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
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = make(map[string]struct{})
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected[want] = struct{}{}
	}
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
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			expected := make([]string, 0, len(p.maxFailExpected))
			eof := false
			if _, ok := p.maxFailExpected["!."]; ok {
				delete(p.maxFailExpected, "!.")
				eof = true
			}
			for k := range p.maxFailExpected {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
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
	var ok bool
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
			p.addErrAt(err, start.position, []string{})
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
		p.failAt(true, start.position, ".")
		return p.sliceFrom(start), true
	}
	p.failAt(false, p.pt.position, ".")
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt
	// can't match EOF
	if cur == utf8.RuneError {
		p.failAt(false, start.position, chr.val)
		return nil, false
	}
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
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

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
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
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
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
