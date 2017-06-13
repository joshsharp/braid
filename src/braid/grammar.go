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
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 791},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 802},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 802},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 802},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 804},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 809},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 810},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 36, col: 1, offset: 838},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 848},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 848},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 36, col: 11, offset: 848},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 15, offset: 852},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 23, offset: 860},
								expr: &seqExpr{
									pos: position{line: 36, col: 24, offset: 861},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 24, offset: 861},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 25, offset: 862},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 37, offset: 874,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 41, offset: 878},
							expr: &litMatcher{
								pos:        position{line: 36, col: 42, offset: 879},
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
			pos:  position{line: 41, col: 1, offset: 1029},
			expr: &choiceExpr{
				pos: position{line: 41, col: 14, offset: 1042},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 14, offset: 1042},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 41, col: 14, offset: 1042},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 14, offset: 1042},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 16, offset: 1044},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 22, offset: 1050},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 25, offset: 1053},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 27, offset: 1055},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 38, offset: 1066},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 43, offset: 1071},
										expr: &seqExpr{
											pos: position{line: 41, col: 44, offset: 1072},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 41, col: 44, offset: 1072},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 1076},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 50, offset: 1078},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 63, offset: 1091},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 65, offset: 1093},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 69, offset: 1097},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 71, offset: 1099},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 76, offset: 1104},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 81, offset: 1109},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 56, col: 1, offset: 1548},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 56, col: 1, offset: 1548},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 56, col: 1, offset: 1548},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 56, col: 3, offset: 1550},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 56, col: 9, offset: 1556},
									name: "__",
								},
								&notExpr{
									pos: position{line: 56, col: 12, offset: 1559},
									expr: &ruleRefExpr{
										pos:  position{line: 56, col: 13, offset: 1560},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1668},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1668},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1668},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1670},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 9, offset: 1676},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 12, offset: 1679},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 14, offset: 1681},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 25, offset: 1692},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 30, offset: 1697},
										expr: &seqExpr{
											pos: position{line: 60, col: 31, offset: 1698},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 60, col: 31, offset: 1698},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 35, offset: 1702},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1704},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 50, offset: 1717},
									name: "_",
								},
								&notExpr{
									pos: position{line: 60, col: 52, offset: 1719},
									expr: &litMatcher{
										pos:        position{line: 60, col: 53, offset: 1720},
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
			pos:  position{line: 64, col: 1, offset: 1814},
			expr: &actionExpr{
				pos: position{line: 64, col: 12, offset: 1825},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 64, col: 12, offset: 1825},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 64, col: 12, offset: 1825},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 64, col: 14, offset: 1827},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 20, offset: 1833},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 64, col: 23, offset: 1836},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 64, col: 25, offset: 1838},
								name: "Assignable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 36, offset: 1849},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 64, col: 38, offset: 1851},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 42, offset: 1855},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 64, col: 44, offset: 1857},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 51, offset: 1864},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 64, col: 54, offset: 1867},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 64, col: 58, offset: 1871},
								expr: &seqExpr{
									pos: position{line: 64, col: 59, offset: 1872},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 64, col: 59, offset: 1872},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 64, col: 72, offset: 1885},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 77, offset: 1890},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 64, col: 79, offset: 1892},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 83, offset: 1896},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 64, col: 86, offset: 1899},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 64, col: 97, offset: 1910},
								expr: &ruleRefExpr{
									pos:  position{line: 64, col: 98, offset: 1911},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 110, offset: 1923},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 64, col: 112, offset: 1925},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 64, col: 116, offset: 1929},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 87, col: 1, offset: 2604},
			expr: &actionExpr{
				pos: position{line: 87, col: 8, offset: 2611},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 87, col: 8, offset: 2611},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 87, col: 12, offset: 2615},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 87, col: 12, offset: 2615},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 87, col: 21, offset: 2624},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 87, col: 28, offset: 2631},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 92, col: 1, offset: 2726},
			expr: &choiceExpr{
				pos: position{line: 92, col: 10, offset: 2735},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 92, col: 10, offset: 2735},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 92, col: 10, offset: 2735},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 92, col: 10, offset: 2735},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 15, offset: 2740},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 18, offset: 2743},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 92, col: 23, offset: 2748},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 33, offset: 2758},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 92, col: 35, offset: 2760},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 39, offset: 2764},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 41, offset: 2766},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 92, col: 47, offset: 2772},
										expr: &ruleRefExpr{
											pos:  position{line: 92, col: 48, offset: 2773},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 60, offset: 2785},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 92, col: 62, offset: 2787},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 66, offset: 2791},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 92, col: 68, offset: 2793},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 75, offset: 2800},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 77, offset: 2802},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 92, col: 85, offset: 2810},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 104, col: 1, offset: 3140},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 104, col: 1, offset: 3140},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 104, col: 1, offset: 3140},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 6, offset: 3145},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 104, col: 9, offset: 3148},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 104, col: 14, offset: 3153},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 24, offset: 3163},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 104, col: 26, offset: 3165},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 30, offset: 3169},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 104, col: 32, offset: 3171},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 104, col: 38, offset: 3177},
										expr: &ruleRefExpr{
											pos:  position{line: 104, col: 39, offset: 3178},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 51, offset: 3190},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 104, col: 54, offset: 3193},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 58, offset: 3197},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 104, col: 60, offset: 3199},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 67, offset: 3206},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 104, col: 69, offset: 3208},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 73, offset: 3212},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 104, col: 75, offset: 3214},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 104, col: 81, offset: 3220},
										expr: &ruleRefExpr{
											pos:  position{line: 104, col: 82, offset: 3221},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 94, offset: 3233},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 104, col: 97, offset: 3236},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 123, col: 1, offset: 3739},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 123, col: 1, offset: 3739},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 123, col: 1, offset: 3739},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 123, col: 6, offset: 3744},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 123, col: 9, offset: 3747},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 123, col: 14, offset: 3752},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 123, col: 24, offset: 3762},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 123, col: 26, offset: 3764},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 123, col: 30, offset: 3768},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 123, col: 32, offset: 3770},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 123, col: 38, offset: 3776},
										expr: &ruleRefExpr{
											pos:  position{line: 123, col: 39, offset: 3777},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 123, col: 51, offset: 3789},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 123, col: 54, offset: 3792},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Call",
			pos:  position{line: 136, col: 1, offset: 4091},
			expr: &choiceExpr{
				pos: position{line: 136, col: 8, offset: 4098},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 8, offset: 4098},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 136, col: 8, offset: 4098},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 136, col: 8, offset: 4098},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 15, offset: 4105},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 136, col: 26, offset: 4116},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 136, col: 30, offset: 4120},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 33, offset: 4123},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 136, col: 46, offset: 4136},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 136, col: 56, offset: 4146},
										expr: &seqExpr{
											pos: position{line: 136, col: 57, offset: 4147},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 136, col: 57, offset: 4147},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 60, offset: 4150},
													name: "Value",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 68, offset: 4158},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 154, col: 1, offset: 4656},
						run: (*parser).callonCall15,
						expr: &seqExpr{
							pos: position{line: 154, col: 1, offset: 4656},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 154, col: 1, offset: 4656},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 154, col: 3, offset: 4658},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 154, col: 6, offset: 4661},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 154, col: 19, offset: 4674},
									label: "arguments",
									expr: &oneOrMoreExpr{
										pos: position{line: 154, col: 29, offset: 4684},
										expr: &seqExpr{
											pos: position{line: 154, col: 30, offset: 4685},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 154, col: 30, offset: 4685},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 154, col: 33, offset: 4688},
													name: "Value",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 41, offset: 4696},
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
			pos:  position{line: 172, col: 1, offset: 5206},
			expr: &actionExpr{
				pos: position{line: 173, col: 1, offset: 5221},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 173, col: 1, offset: 5221},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 173, col: 1, offset: 5221},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 173, col: 3, offset: 5223},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 173, col: 6, offset: 5226},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 173, col: 12, offset: 5232},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 173, col: 17, offset: 5237},
								expr: &seqExpr{
									pos: position{line: 173, col: 18, offset: 5238},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 173, col: 18, offset: 5238},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 21, offset: 5241},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 30, offset: 5250},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 33, offset: 5253},
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
			pos:  position{line: 196, col: 1, offset: 5920},
			expr: &choiceExpr{
				pos: position{line: 196, col: 9, offset: 5928},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 196, col: 9, offset: 5928},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 196, col: 21, offset: 5940},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 196, col: 37, offset: 5956},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 196, col: 48, offset: 5967},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 196, col: 60, offset: 5979},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 198, col: 1, offset: 5992},
			expr: &actionExpr{
				pos: position{line: 198, col: 13, offset: 6004},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 198, col: 13, offset: 6004},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 198, col: 13, offset: 6004},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 198, col: 15, offset: 6006},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 198, col: 21, offset: 6012},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 198, col: 35, offset: 6026},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 198, col: 40, offset: 6031},
								expr: &seqExpr{
									pos: position{line: 198, col: 41, offset: 6032},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 198, col: 41, offset: 6032},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 198, col: 44, offset: 6035},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 198, col: 60, offset: 6051},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 198, col: 63, offset: 6054},
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
			pos:  position{line: 217, col: 1, offset: 6660},
			expr: &actionExpr{
				pos: position{line: 217, col: 17, offset: 6676},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 217, col: 17, offset: 6676},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 217, col: 17, offset: 6676},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 217, col: 19, offset: 6678},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 25, offset: 6684},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 217, col: 34, offset: 6693},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 217, col: 39, offset: 6698},
								expr: &seqExpr{
									pos: position{line: 217, col: 40, offset: 6699},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 217, col: 40, offset: 6699},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 217, col: 43, offset: 6702},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 217, col: 60, offset: 6719},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 217, col: 63, offset: 6722},
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
			pos:  position{line: 237, col: 1, offset: 7326},
			expr: &actionExpr{
				pos: position{line: 237, col: 12, offset: 7337},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 237, col: 12, offset: 7337},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 237, col: 12, offset: 7337},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 14, offset: 7339},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 20, offset: 7345},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 30, offset: 7355},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 237, col: 35, offset: 7360},
								expr: &seqExpr{
									pos: position{line: 237, col: 36, offset: 7361},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 237, col: 36, offset: 7361},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 39, offset: 7364},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 51, offset: 7376},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 54, offset: 7379},
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
			pos:  position{line: 257, col: 1, offset: 7980},
			expr: &actionExpr{
				pos: position{line: 257, col: 13, offset: 7992},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 257, col: 13, offset: 7992},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 257, col: 13, offset: 7992},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 257, col: 15, offset: 7994},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 257, col: 21, offset: 8000},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 257, col: 33, offset: 8012},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 257, col: 38, offset: 8017},
								expr: &seqExpr{
									pos: position{line: 257, col: 39, offset: 8018},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 257, col: 39, offset: 8018},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 257, col: 42, offset: 8021},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 257, col: 55, offset: 8034},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 257, col: 58, offset: 8037},
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
			pos:  position{line: 276, col: 1, offset: 8641},
			expr: &choiceExpr{
				pos: position{line: 276, col: 15, offset: 8655},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 276, col: 15, offset: 8655},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 276, col: 15, offset: 8655},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 276, col: 15, offset: 8655},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 276, col: 19, offset: 8659},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 276, col: 21, offset: 8661},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 276, col: 27, offset: 8667},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 276, col: 33, offset: 8673},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 276, col: 35, offset: 8675},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 279, col: 5, offset: 8824},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 281, col: 1, offset: 8831},
			expr: &choiceExpr{
				pos: position{line: 281, col: 12, offset: 8842},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 281, col: 12, offset: 8842},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 281, col: 30, offset: 8860},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 281, col: 49, offset: 8879},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 281, col: 64, offset: 8894},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 283, col: 1, offset: 8907},
			expr: &actionExpr{
				pos: position{line: 283, col: 19, offset: 8925},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 283, col: 21, offset: 8927},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 283, col: 21, offset: 8927},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 283, col: 29, offset: 8935},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 283, col: 36, offset: 8942},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 287, col: 1, offset: 9041},
			expr: &actionExpr{
				pos: position{line: 287, col: 20, offset: 9060},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 287, col: 22, offset: 9062},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 287, col: 22, offset: 9062},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 287, col: 29, offset: 9069},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 287, col: 36, offset: 9076},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 287, col: 42, offset: 9082},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 287, col: 48, offset: 9088},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 287, col: 56, offset: 9096},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 291, col: 1, offset: 9202},
			expr: &choiceExpr{
				pos: position{line: 291, col: 16, offset: 9217},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 291, col: 16, offset: 9217},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 291, col: 18, offset: 9219},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 291, col: 18, offset: 9219},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 291, col: 25, offset: 9226},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 294, col: 3, offset: 9332},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 294, col: 5, offset: 9334},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 294, col: 5, offset: 9334},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 294, col: 11, offset: 9340},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 294, col: 17, offset: 9346},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 297, col: 3, offset: 9449},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 297, col: 3, offset: 9449},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 301, col: 1, offset: 9553},
			expr: &choiceExpr{
				pos: position{line: 301, col: 15, offset: 9567},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 301, col: 15, offset: 9567},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 301, col: 17, offset: 9569},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 301, col: 17, offset: 9569},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 301, col: 24, offset: 9576},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 304, col: 3, offset: 9682},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 304, col: 5, offset: 9684},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 304, col: 5, offset: 9684},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 304, col: 11, offset: 9690},
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
			pos:  position{line: 308, col: 1, offset: 9792},
			expr: &choiceExpr{
				pos: position{line: 308, col: 9, offset: 9800},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 308, col: 9, offset: 9800},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 308, col: 24, offset: 9815},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 310, col: 1, offset: 9822},
			expr: &choiceExpr{
				pos: position{line: 310, col: 14, offset: 9835},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 310, col: 14, offset: 9835},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 310, col: 29, offset: 9850},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 312, col: 1, offset: 9858},
			expr: &choiceExpr{
				pos: position{line: 312, col: 14, offset: 9871},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 312, col: 14, offset: 9871},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 312, col: 29, offset: 9886},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 314, col: 1, offset: 9898},
			expr: &actionExpr{
				pos: position{line: 314, col: 16, offset: 9913},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 314, col: 16, offset: 9913},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 314, col: 16, offset: 9913},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 20, offset: 9917},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 314, col: 22, offset: 9919},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 314, col: 28, offset: 9925},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 33, offset: 9930},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 314, col: 35, offset: 9932},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 314, col: 40, offset: 9937},
								expr: &seqExpr{
									pos: position{line: 314, col: 41, offset: 9938},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 314, col: 41, offset: 9938},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 314, col: 45, offset: 9942},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 314, col: 47, offset: 9944},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 314, col: 52, offset: 9949},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 314, col: 56, offset: 9953},
							expr: &litMatcher{
								pos:        position{line: 314, col: 56, offset: 9953},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 61, offset: 9958},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 314, col: 63, offset: 9960},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 330, col: 1, offset: 10441},
			expr: &actionExpr{
				pos: position{line: 330, col: 16, offset: 10456},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 330, col: 16, offset: 10456},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 330, col: 16, offset: 10456},
							expr: &ruleRefExpr{
								pos:  position{line: 330, col: 17, offset: 10457},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 330, col: 27, offset: 10467},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 330, col: 27, offset: 10467},
									expr: &charClassMatcher{
										pos:        position{line: 330, col: 27, offset: 10467},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 330, col: 34, offset: 10474},
									expr: &charClassMatcher{
										pos:        position{line: 330, col: 34, offset: 10474},
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
			pos:  position{line: 334, col: 1, offset: 10585},
			expr: &actionExpr{
				pos: position{line: 334, col: 14, offset: 10598},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 334, col: 15, offset: 10599},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 334, col: 15, offset: 10599},
							expr: &charClassMatcher{
								pos:        position{line: 334, col: 15, offset: 10599},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 334, col: 22, offset: 10606},
							expr: &charClassMatcher{
								pos:        position{line: 334, col: 22, offset: 10606},
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
			pos:  position{line: 338, col: 1, offset: 10717},
			expr: &choiceExpr{
				pos: position{line: 338, col: 9, offset: 10725},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 338, col: 9, offset: 10725},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 338, col: 9, offset: 10725},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 338, col: 9, offset: 10725},
									expr: &litMatcher{
										pos:        position{line: 338, col: 9, offset: 10725},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 338, col: 14, offset: 10730},
									expr: &charClassMatcher{
										pos:        position{line: 338, col: 14, offset: 10730},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 338, col: 21, offset: 10737},
									expr: &litMatcher{
										pos:        position{line: 338, col: 22, offset: 10738},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 345, col: 3, offset: 10914},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 345, col: 3, offset: 10914},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 345, col: 3, offset: 10914},
									expr: &litMatcher{
										pos:        position{line: 345, col: 3, offset: 10914},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 345, col: 8, offset: 10919},
									expr: &charClassMatcher{
										pos:        position{line: 345, col: 8, offset: 10919},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 345, col: 15, offset: 10926},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 345, col: 19, offset: 10930},
									expr: &charClassMatcher{
										pos:        position{line: 345, col: 19, offset: 10930},
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
						pos:        position{line: 352, col: 3, offset: 11120},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 352, col: 12, offset: 11129},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 352, col: 12, offset: 11129},
							val:        "false",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 358, col: 3, offset: 11330},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 358, col: 3, offset: 11330},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 361, col: 3, offset: 11393},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 361, col: 3, offset: 11393},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 361, col: 3, offset: 11393},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 361, col: 7, offset: 11397},
									expr: &seqExpr{
										pos: position{line: 361, col: 8, offset: 11398},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 361, col: 8, offset: 11398},
												expr: &ruleRefExpr{
													pos:  position{line: 361, col: 9, offset: 11399},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 361, col: 21, offset: 11411,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 361, col: 25, offset: 11415},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 368, col: 3, offset: 11599},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 368, col: 3, offset: 11599},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 368, col: 3, offset: 11599},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 368, col: 7, offset: 11603},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 368, col: 12, offset: 11608},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 368, col: 12, offset: 11608},
												expr: &ruleRefExpr{
													pos:  position{line: 368, col: 13, offset: 11609},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 368, col: 25, offset: 11621,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 368, col: 28, offset: 11624},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 370, col: 5, offset: 11716},
						name: "ArrayLiteral",
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 372, col: 1, offset: 11730},
			expr: &actionExpr{
				pos: position{line: 372, col: 10, offset: 11739},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 372, col: 11, offset: 11740},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 376, col: 1, offset: 11841},
			expr: &seqExpr{
				pos: position{line: 376, col: 12, offset: 11852},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 376, col: 13, offset: 11853},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 376, col: 13, offset: 11853},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 21, offset: 11861},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 28, offset: 11868},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 37, offset: 11877},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 46, offset: 11886},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 55, offset: 11895},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 64, offset: 11904},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 376, col: 74, offset: 11914},
								val:        "mutable",
								ignoreCase: false,
							},
						},
					},
					&notExpr{
						pos: position{line: 376, col: 86, offset: 11926},
						expr: &oneOrMoreExpr{
							pos: position{line: 376, col: 87, offset: 11927},
							expr: &charClassMatcher{
								pos:        position{line: 376, col: 87, offset: 11927},
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
			pos:  position{line: 378, col: 1, offset: 11935},
			expr: &charClassMatcher{
				pos:        position{line: 378, col: 15, offset: 11949},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 380, col: 1, offset: 11965},
			expr: &choiceExpr{
				pos: position{line: 380, col: 18, offset: 11982},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 380, col: 18, offset: 11982},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 380, col: 37, offset: 12001},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 382, col: 1, offset: 12016},
			expr: &charClassMatcher{
				pos:        position{line: 382, col: 20, offset: 12035},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 384, col: 1, offset: 12048},
			expr: &charClassMatcher{
				pos:        position{line: 384, col: 16, offset: 12063},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 386, col: 1, offset: 12070},
			expr: &charClassMatcher{
				pos:        position{line: 386, col: 23, offset: 12092},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 388, col: 1, offset: 12099},
			expr: &charClassMatcher{
				pos:        position{line: 388, col: 12, offset: 12110},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 390, col: 1, offset: 12121},
			expr: &oneOrMoreExpr{
				pos: position{line: 390, col: 22, offset: 12142},
				expr: &charClassMatcher{
					pos:        position{line: 390, col: 22, offset: 12142},
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
			pos:         position{line: 392, col: 1, offset: 12154},
			expr: &zeroOrMoreExpr{
				pos: position{line: 392, col: 18, offset: 12171},
				expr: &charClassMatcher{
					pos:        position{line: 392, col: 18, offset: 12171},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 394, col: 1, offset: 12183},
			expr: &notExpr{
				pos: position{line: 394, col: 7, offset: 12189},
				expr: &anyMatcher{
					line: 394, col: 8, offset: 12190,
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

func (c *current) onExprLine1(e interface{}) (interface{}, error) {
	return e, nil
}

func (p *parser) callonExprLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExprLine1(stack["e"])
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

func (c *current) onIfExpr45(expr, thens interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr45() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr45(stack["expr"], stack["thens"])
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

func (c *current) onCall15(fn, arguments interface{}) (interface{}, error) {
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

func (p *parser) callonCall15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall15(stack["fn"], stack["arguments"])
}

func (c *current) onCompoundExpr1(op, rest interface{}) (interface{}, error) {
	//fmt.Println("compound", op, rest)
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
	//fmt.Println("binopbool", first, rest)
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
	//fmt.Println("binopeq", first, rest)
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
	//fmt.Println("binoplow", first, rest)
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
	//fmt.Println("binophigh", first, rest)
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
	//fmt.Println("binopparens", first)
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

func (c *current) onArrayLiteral1(first, rest interface{}) (interface{}, error) {
	// rest:(_ ',' _ Expr)*
	vals := rest.([]interface{})
	subvalues := []Ast{first.(Ast)}
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			subvalues = append(subvalues, v)
		}
	}
	return BasicAst{Type: "Array", Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonArrayLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArrayLiteral1(stack["first"], stack["rest"])
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
