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
								name: "Statement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 27, offset: 169},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 29, offset: 171},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 34, offset: 176},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 35, offset: 177},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 47, offset: 189},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 49, offset: 191},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 28, col: 1, offset: 679},
			expr: &choiceExpr{
				pos: position{line: 28, col: 13, offset: 691},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 28, col: 13, offset: 691},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 28, col: 13, offset: 691},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 28, col: 13, offset: 691},
									val:        "#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 28, col: 17, offset: 695},
									label: "comment",
									expr: &zeroOrMoreExpr{
										pos: position{line: 28, col: 25, offset: 703},
										expr: &seqExpr{
											pos: position{line: 28, col: 26, offset: 704},
											exprs: []interface{}{
												&notExpr{
													pos: position{line: 28, col: 26, offset: 704},
													expr: &ruleRefExpr{
														pos:  position{line: 28, col: 27, offset: 705},
														name: "EscapedChar",
													},
												},
												&anyMatcher{
													line: 28, col: 39, offset: 717,
												},
											},
										},
									},
								},
								&andExpr{
									pos: position{line: 28, col: 43, offset: 721},
									expr: &litMatcher{
										pos:        position{line: 28, col: 44, offset: 722},
										val:        "\n",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 33, col: 1, offset: 873},
						run: (*parser).callonStatement13,
						expr: &seqExpr{
							pos: position{line: 33, col: 1, offset: 873},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 33, col: 1, offset: 873},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 33, col: 3, offset: 875},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 33, col: 9, offset: 881},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 33, col: 12, offset: 884},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 33, col: 14, offset: 886},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 33, col: 25, offset: 897},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 33, col: 30, offset: 902},
										expr: &seqExpr{
											pos: position{line: 33, col: 31, offset: 903},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 33, col: 31, offset: 903},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 33, col: 35, offset: 907},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 33, col: 37, offset: 909},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 33, col: 50, offset: 922},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 33, col: 52, offset: 924},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 33, col: 56, offset: 928},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 33, col: 58, offset: 930},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 33, col: 63, offset: 935},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 33, col: 68, offset: 940},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 48, col: 1, offset: 1379},
						run: (*parser).callonStatement32,
						expr: &seqExpr{
							pos: position{line: 48, col: 1, offset: 1379},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 48, col: 1, offset: 1379},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 48, col: 3, offset: 1381},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 48, col: 9, offset: 1387},
									name: "__",
								},
								&notExpr{
									pos: position{line: 48, col: 12, offset: 1390},
									expr: &ruleRefExpr{
										pos:  position{line: 48, col: 13, offset: 1391},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 52, col: 1, offset: 1499},
						run: (*parser).callonStatement39,
						expr: &seqExpr{
							pos: position{line: 52, col: 1, offset: 1499},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 1, offset: 1499},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 3, offset: 1501},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 9, offset: 1507},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 12, offset: 1510},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 14, offset: 1512},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 52, col: 25, offset: 1523},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 52, col: 30, offset: 1528},
										expr: &seqExpr{
											pos: position{line: 52, col: 31, offset: 1529},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 52, col: 31, offset: 1529},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 35, offset: 1533},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 37, offset: 1535},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 50, offset: 1548},
									name: "_",
								},
								&notExpr{
									pos: position{line: 52, col: 52, offset: 1550},
									expr: &litMatcher{
										pos:        position{line: 52, col: 53, offset: 1551},
										val:        "=",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 3, offset: 1646},
						name: "Expr",
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 57, col: 1, offset: 1652},
			expr: &actionExpr{
				pos: position{line: 57, col: 8, offset: 1659},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 57, col: 8, offset: 1659},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 57, col: 12, offset: 1663},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 57, col: 12, offset: 1663},
								name: "FuncDefn",
							},
							&ruleRefExpr{
								pos:  position{line: 57, col: 23, offset: 1674},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 57, col: 30, offset: 1681},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 57, col: 39, offset: 1690},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 62, col: 1, offset: 1784},
			expr: &choiceExpr{
				pos: position{line: 62, col: 10, offset: 1793},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 62, col: 10, offset: 1793},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 62, col: 10, offset: 1793},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 62, col: 10, offset: 1793},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 15, offset: 1798},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 62, col: 18, offset: 1801},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 62, col: 23, offset: 1806},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 33, offset: 1816},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 62, col: 35, offset: 1818},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 39, offset: 1822},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 62, col: 41, offset: 1824},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 62, col: 47, offset: 1830},
										expr: &ruleRefExpr{
											pos:  position{line: 62, col: 48, offset: 1831},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 60, offset: 1843},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 62, col: 62, offset: 1845},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 66, offset: 1849},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 62, col: 68, offset: 1851},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 75, offset: 1858},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 62, col: 77, offset: 1860},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 62, col: 85, offset: 1868},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 74, col: 1, offset: 2198},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 74, col: 1, offset: 2198},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 74, col: 1, offset: 2198},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 6, offset: 2203},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 9, offset: 2206},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 74, col: 14, offset: 2211},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 24, offset: 2221},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 74, col: 26, offset: 2223},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 30, offset: 2227},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 32, offset: 2229},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 74, col: 38, offset: 2235},
										expr: &ruleRefExpr{
											pos:  position{line: 74, col: 39, offset: 2236},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 51, offset: 2248},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 74, col: 53, offset: 2250},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 57, offset: 2254},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 74, col: 59, offset: 2256},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 66, offset: 2263},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 74, col: 68, offset: 2265},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 72, offset: 2269},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 74, offset: 2271},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 74, col: 80, offset: 2277},
										expr: &ruleRefExpr{
											pos:  position{line: 74, col: 81, offset: 2278},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 93, offset: 2290},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 74, col: 95, offset: 2292},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 99, offset: 2296},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 93, col: 1, offset: 2797},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 93, col: 1, offset: 2797},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 93, col: 1, offset: 2797},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 6, offset: 2802},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 93, col: 9, offset: 2805},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 93, col: 14, offset: 2810},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 24, offset: 2820},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 93, col: 26, offset: 2822},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 30, offset: 2826},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 93, col: 32, offset: 2828},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 93, col: 38, offset: 2834},
										expr: &ruleRefExpr{
											pos:  position{line: 93, col: 39, offset: 2835},
											name: "Statement",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 93, col: 51, offset: 2847},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 55, offset: 2851},
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
			pos:  position{line: 105, col: 1, offset: 3147},
			expr: &actionExpr{
				pos: position{line: 105, col: 12, offset: 3158},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 105, col: 12, offset: 3158},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 12, offset: 3158},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 19, offset: 3165},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 105, col: 22, offset: 3168},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 105, col: 26, offset: 3172},
								expr: &seqExpr{
									pos: position{line: 105, col: 27, offset: 3173},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 105, col: 27, offset: 3173},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 105, col: 40, offset: 3186},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 45, offset: 3191},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 105, col: 47, offset: 3193},
							expr: &litMatcher{
								pos:        position{line: 105, col: 47, offset: 3193},
								val:        "->",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 53, offset: 3199},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 105, col: 55, offset: 3201},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 59, offset: 3205},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 105, col: 61, offset: 3207},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 105, col: 72, offset: 3218},
								expr: &ruleRefExpr{
									pos:  position{line: 105, col: 73, offset: 3219},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 85, offset: 3231},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 105, col: 87, offset: 3233},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 91, offset: 3237},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Call",
			pos:  position{line: 128, col: 1, offset: 3880},
			expr: &choiceExpr{
				pos: position{line: 128, col: 8, offset: 3887},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 128, col: 8, offset: 3887},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 128, col: 8, offset: 3887},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 128, col: 8, offset: 3887},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 128, col: 15, offset: 3894},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 128, col: 26, offset: 3905},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 128, col: 30, offset: 3909},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 128, col: 33, offset: 3912},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 128, col: 46, offset: 3925},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 128, col: 56, offset: 3935},
										expr: &seqExpr{
											pos: position{line: 128, col: 57, offset: 3936},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 128, col: 57, offset: 3936},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 128, col: 60, offset: 3939},
													name: "Value",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 128, col: 68, offset: 3947},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 72, offset: 3951},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 146, col: 1, offset: 4449},
						run: (*parser).callonCall16,
						expr: &seqExpr{
							pos: position{line: 146, col: 1, offset: 4449},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 146, col: 1, offset: 4449},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 146, col: 4, offset: 4452},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 146, col: 17, offset: 4465},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 146, col: 27, offset: 4475},
										expr: &seqExpr{
											pos: position{line: 146, col: 28, offset: 4476},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 146, col: 28, offset: 4476},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 146, col: 31, offset: 4479},
													name: "Value",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 146, col: 39, offset: 4487},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 43, offset: 4491},
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
			pos:  position{line: 164, col: 1, offset: 5001},
			expr: &actionExpr{
				pos: position{line: 164, col: 16, offset: 5016},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 164, col: 16, offset: 5016},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 164, col: 16, offset: 5016},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 164, col: 18, offset: 5018},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 164, col: 21, offset: 5021},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 164, col: 27, offset: 5027},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 164, col: 32, offset: 5032},
								expr: &seqExpr{
									pos: position{line: 164, col: 33, offset: 5033},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 164, col: 33, offset: 5033},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 164, col: 36, offset: 5036},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 164, col: 45, offset: 5045},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 164, col: 48, offset: 5048},
											name: "BinOp",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 164, col: 56, offset: 5056},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 164, col: 58, offset: 5058},
							val:        ";",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 164, col: 62, offset: 5062},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 184, col: 1, offset: 5721},
			expr: &choiceExpr{
				pos: position{line: 184, col: 9, offset: 5729},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 184, col: 9, offset: 5729},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 21, offset: 5741},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 37, offset: 5757},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 48, offset: 5768},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 60, offset: 5780},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 186, col: 1, offset: 5793},
			expr: &actionExpr{
				pos: position{line: 186, col: 13, offset: 5805},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 186, col: 13, offset: 5805},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 186, col: 13, offset: 5805},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 186, col: 15, offset: 5807},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 186, col: 21, offset: 5813},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 186, col: 35, offset: 5827},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 186, col: 40, offset: 5832},
								expr: &seqExpr{
									pos: position{line: 186, col: 41, offset: 5833},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 186, col: 41, offset: 5833},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 186, col: 44, offset: 5836},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 186, col: 60, offset: 5852},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 186, col: 63, offset: 5855},
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
			pos:  position{line: 205, col: 1, offset: 6462},
			expr: &actionExpr{
				pos: position{line: 205, col: 17, offset: 6478},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 205, col: 17, offset: 6478},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 205, col: 17, offset: 6478},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 205, col: 19, offset: 6480},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 25, offset: 6486},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 205, col: 34, offset: 6495},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 205, col: 39, offset: 6500},
								expr: &seqExpr{
									pos: position{line: 205, col: 40, offset: 6501},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 205, col: 40, offset: 6501},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 205, col: 43, offset: 6504},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 205, col: 60, offset: 6521},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 205, col: 63, offset: 6524},
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
			pos:  position{line: 225, col: 1, offset: 7129},
			expr: &actionExpr{
				pos: position{line: 225, col: 12, offset: 7140},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 225, col: 12, offset: 7140},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 225, col: 12, offset: 7140},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 225, col: 14, offset: 7142},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 20, offset: 7148},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 30, offset: 7158},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 225, col: 35, offset: 7163},
								expr: &seqExpr{
									pos: position{line: 225, col: 36, offset: 7164},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 225, col: 36, offset: 7164},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 225, col: 39, offset: 7167},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 225, col: 51, offset: 7179},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 225, col: 54, offset: 7182},
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
			pos:  position{line: 245, col: 1, offset: 7784},
			expr: &actionExpr{
				pos: position{line: 245, col: 13, offset: 7796},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 245, col: 13, offset: 7796},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 245, col: 13, offset: 7796},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 245, col: 15, offset: 7798},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 245, col: 21, offset: 7804},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 245, col: 33, offset: 7816},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 245, col: 38, offset: 7821},
								expr: &seqExpr{
									pos: position{line: 245, col: 39, offset: 7822},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 245, col: 39, offset: 7822},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 245, col: 42, offset: 7825},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 245, col: 55, offset: 7838},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 245, col: 58, offset: 7841},
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
			pos:  position{line: 264, col: 1, offset: 8446},
			expr: &choiceExpr{
				pos: position{line: 264, col: 15, offset: 8460},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 264, col: 15, offset: 8460},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 264, col: 15, offset: 8460},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 264, col: 15, offset: 8460},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 264, col: 19, offset: 8464},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 21, offset: 8466},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 27, offset: 8472},
										name: "BinOp",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 33, offset: 8478},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 267, col: 5, offset: 8628},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 269, col: 1, offset: 8635},
			expr: &choiceExpr{
				pos: position{line: 269, col: 12, offset: 8646},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 269, col: 12, offset: 8646},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 269, col: 30, offset: 8664},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 269, col: 49, offset: 8683},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 269, col: 64, offset: 8698},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 271, col: 1, offset: 8711},
			expr: &actionExpr{
				pos: position{line: 271, col: 19, offset: 8729},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 271, col: 21, offset: 8731},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 271, col: 21, offset: 8731},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 271, col: 29, offset: 8739},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 271, col: 36, offset: 8746},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 275, col: 1, offset: 8845},
			expr: &actionExpr{
				pos: position{line: 275, col: 20, offset: 8864},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 275, col: 22, offset: 8866},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 275, col: 22, offset: 8866},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 29, offset: 8873},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 36, offset: 8880},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 42, offset: 8886},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 48, offset: 8892},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 275, col: 56, offset: 8900},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 279, col: 1, offset: 9006},
			expr: &choiceExpr{
				pos: position{line: 279, col: 16, offset: 9021},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 279, col: 16, offset: 9021},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 279, col: 18, offset: 9023},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 279, col: 18, offset: 9023},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 279, col: 25, offset: 9030},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 282, col: 3, offset: 9136},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 282, col: 5, offset: 9138},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 282, col: 5, offset: 9138},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 282, col: 11, offset: 9144},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 282, col: 17, offset: 9150},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 285, col: 3, offset: 9253},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 285, col: 3, offset: 9253},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 289, col: 1, offset: 9357},
			expr: &choiceExpr{
				pos: position{line: 289, col: 15, offset: 9371},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 289, col: 15, offset: 9371},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 289, col: 17, offset: 9373},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 289, col: 17, offset: 9373},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 289, col: 24, offset: 9380},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 292, col: 3, offset: 9486},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 292, col: 5, offset: 9488},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 292, col: 5, offset: 9488},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 292, col: 11, offset: 9494},
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
			pos:  position{line: 296, col: 1, offset: 9596},
			expr: &choiceExpr{
				pos: position{line: 296, col: 9, offset: 9604},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 296, col: 9, offset: 9604},
						name: "Identifier",
					},
					&actionExpr{
						pos: position{line: 296, col: 22, offset: 9617},
						run: (*parser).callonValue3,
						expr: &labeledExpr{
							pos:   position{line: 296, col: 22, offset: 9617},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 24, offset: 9619},
								name: "Const",
							},
						},
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 300, col: 1, offset: 9654},
			expr: &choiceExpr{
				pos: position{line: 300, col: 14, offset: 9667},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 300, col: 14, offset: 9667},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 29, offset: 9682},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 302, col: 1, offset: 9690},
			expr: &choiceExpr{
				pos: position{line: 302, col: 14, offset: 9703},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 302, col: 14, offset: 9703},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 302, col: 29, offset: 9718},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 304, col: 1, offset: 9730},
			expr: &actionExpr{
				pos: position{line: 304, col: 16, offset: 9745},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 304, col: 16, offset: 9745},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 304, col: 16, offset: 9745},
							expr: &ruleRefExpr{
								pos:  position{line: 304, col: 17, offset: 9746},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 304, col: 27, offset: 9756},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 304, col: 27, offset: 9756},
									expr: &charClassMatcher{
										pos:        position{line: 304, col: 27, offset: 9756},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 304, col: 34, offset: 9763},
									expr: &charClassMatcher{
										pos:        position{line: 304, col: 34, offset: 9763},
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
			pos:  position{line: 308, col: 1, offset: 9874},
			expr: &actionExpr{
				pos: position{line: 308, col: 14, offset: 9887},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 308, col: 15, offset: 9888},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 308, col: 15, offset: 9888},
							expr: &charClassMatcher{
								pos:        position{line: 308, col: 15, offset: 9888},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 308, col: 22, offset: 9895},
							expr: &charClassMatcher{
								pos:        position{line: 308, col: 22, offset: 9895},
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
			pos:  position{line: 312, col: 1, offset: 10006},
			expr: &choiceExpr{
				pos: position{line: 312, col: 9, offset: 10014},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 312, col: 9, offset: 10014},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 312, col: 9, offset: 10014},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 312, col: 9, offset: 10014},
									expr: &litMatcher{
										pos:        position{line: 312, col: 9, offset: 10014},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 312, col: 14, offset: 10019},
									expr: &charClassMatcher{
										pos:        position{line: 312, col: 14, offset: 10019},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 312, col: 21, offset: 10026},
									expr: &litMatcher{
										pos:        position{line: 312, col: 22, offset: 10027},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 319, col: 3, offset: 10203},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 319, col: 3, offset: 10203},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 319, col: 3, offset: 10203},
									expr: &litMatcher{
										pos:        position{line: 319, col: 3, offset: 10203},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 319, col: 8, offset: 10208},
									expr: &charClassMatcher{
										pos:        position{line: 319, col: 8, offset: 10208},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 319, col: 15, offset: 10215},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 319, col: 19, offset: 10219},
									expr: &charClassMatcher{
										pos:        position{line: 319, col: 19, offset: 10219},
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
						pos:        position{line: 326, col: 3, offset: 10409},
						val:        "True",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 326, col: 12, offset: 10418},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 326, col: 12, offset: 10418},
							val:        "False",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 332, col: 3, offset: 10619},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 332, col: 3, offset: 10619},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 335, col: 3, offset: 10682},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 335, col: 3, offset: 10682},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 335, col: 3, offset: 10682},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 335, col: 7, offset: 10686},
									expr: &seqExpr{
										pos: position{line: 335, col: 8, offset: 10687},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 335, col: 8, offset: 10687},
												expr: &ruleRefExpr{
													pos:  position{line: 335, col: 9, offset: 10688},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 335, col: 21, offset: 10700,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 335, col: 25, offset: 10704},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 342, col: 3, offset: 10888},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 342, col: 3, offset: 10888},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 342, col: 3, offset: 10888},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 342, col: 7, offset: 10892},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 342, col: 12, offset: 10897},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 342, col: 12, offset: 10897},
												expr: &ruleRefExpr{
													pos:  position{line: 342, col: 13, offset: 10898},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 342, col: 25, offset: 10910,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 342, col: 28, offset: 10913},
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
			pos:  position{line: 346, col: 1, offset: 11004},
			expr: &actionExpr{
				pos: position{line: 346, col: 10, offset: 11013},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 346, col: 11, offset: 11014},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 350, col: 1, offset: 11115},
			expr: &seqExpr{
				pos: position{line: 350, col: 12, offset: 11126},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 350, col: 13, offset: 11127},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 350, col: 13, offset: 11127},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 350, col: 21, offset: 11135},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 350, col: 28, offset: 11142},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 350, col: 37, offset: 11151},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 350, col: 46, offset: 11160},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 350, col: 55, offset: 11169},
								val:        "True",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 350, col: 64, offset: 11178},
								val:        "False",
								ignoreCase: false,
							},
						},
					},
					&notExpr{
						pos: position{line: 350, col: 74, offset: 11188},
						expr: &oneOrMoreExpr{
							pos: position{line: 350, col: 75, offset: 11189},
							expr: &charClassMatcher{
								pos:        position{line: 350, col: 75, offset: 11189},
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
			pos:  position{line: 352, col: 1, offset: 11197},
			expr: &charClassMatcher{
				pos:        position{line: 352, col: 15, offset: 11211},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 354, col: 1, offset: 11227},
			expr: &choiceExpr{
				pos: position{line: 354, col: 18, offset: 11244},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 354, col: 18, offset: 11244},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 354, col: 37, offset: 11263},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 356, col: 1, offset: 11278},
			expr: &charClassMatcher{
				pos:        position{line: 356, col: 20, offset: 11297},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 358, col: 1, offset: 11310},
			expr: &charClassMatcher{
				pos:        position{line: 358, col: 16, offset: 11325},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 360, col: 1, offset: 11332},
			expr: &charClassMatcher{
				pos:        position{line: 360, col: 23, offset: 11354},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 362, col: 1, offset: 11361},
			expr: &charClassMatcher{
				pos:        position{line: 362, col: 12, offset: 11372},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 364, col: 1, offset: 11383},
			expr: &oneOrMoreExpr{
				pos: position{line: 364, col: 22, offset: 11404},
				expr: &charClassMatcher{
					pos:        position{line: 364, col: 22, offset: 11404},
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
			pos:         position{line: 366, col: 1, offset: 11416},
			expr: &zeroOrMoreExpr{
				pos: position{line: 366, col: 18, offset: 11433},
				expr: &charClassMatcher{
					pos:        position{line: 366, col: 18, offset: 11433},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 368, col: 1, offset: 11445},
			expr: &notExpr{
				pos: position{line: 368, col: 7, offset: 11451},
				expr: &anyMatcher{
					line: 368, col: 8, offset: 11452,
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

func (c *current) onStatement2(comment interface{}) (interface{}, error) {
	//fmt.Println("comment:", string(c.text))
	return BasicAst{Type: "Comment", StringValue: string(c.text[1:]), ValueType: STRING}, nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["comment"])
}

func (c *current) onStatement13(i, rest, expr interface{}) (interface{}, error) {
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

func (p *parser) callonStatement13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement13(stack["i"], stack["rest"], stack["expr"])
}

func (c *current) onStatement32() (interface{}, error) {
	return nil, errors.New("Variable name or '_' (unused result character) required here")
}

func (p *parser) callonStatement32() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement32()
}

func (c *current) onStatement39(i, rest interface{}) (interface{}, error) {
	return nil, errors.New("When assigning a value to a variable, you must use '='")
}

func (p *parser) callonStatement39() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement39(stack["i"], stack["rest"])
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

func (c *current) onFuncDefn1(ids, statements interface{}) (interface{}, error) {
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
	return Func{Arguments: args, Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["ids"], stack["statements"])
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

// FailureTracking creates an Option to set failureTracking flag to b.
// When set to true, this causes the parser to track the farthest failures
// and report them, if the parsing of the input fails.
//
// The default is false.
func FailureTracking(b bool) Option {
	return func(p *parser) Option {
		old := p.failureTracking
		p.failureTracking = b
		return FailureTracking(old)
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
	failureTracking       bool
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
			if p.failureTracking {
				// If parsing fails, but no errors have been recorded, the expected values
				// for the farthest parser position are returned as error.
				var expected []string
				eof := ""
				if _, ok := p.maxFailExpected["!."]; ok {
					delete(p.maxFailExpected, "!.")
					if len(p.maxFailExpected) > 0 {
						eof = " or EOF"
					} else {
						eof = "EOF"
					}
				}
				for k := range p.maxFailExpected {
					expected = append(expected, k)
				}
				sort.Strings(expected)
				p.addErrAt(errors.New("no match found, expected: "+strings.Join(expected, ", ")+eof), p.maxFailPos)
			} else {
				// make sure this doesn't go out silently
				p.addErr(errNoMatch)
			}
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

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if p.failureTracking && fail == p.maxFailInvertExpected {
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
