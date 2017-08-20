package ast

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
			pos:  position{line: 13, col: 1, offset: 142},
			expr: &actionExpr{
				pos: position{line: 13, col: 10, offset: 151},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 13, col: 10, offset: 151},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 13, col: 10, offset: 151},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 12, offset: 153},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 17, offset: 158},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 35, offset: 176},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 37, offset: 178},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 42, offset: 183},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 43, offset: 184},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 63, offset: 204},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 65, offset: 206},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 633},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 653},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 653},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 31, offset: 663},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 42, offset: 674},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 684},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 696},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 696},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 706},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 717},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 730},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 54, offset: 737},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 747},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 758},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 758},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 758},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 760},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 765},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 766},
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
			pos:  position{line: 36, col: 1, offset: 794},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 804},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 804},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 36, col: 11, offset: 804},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 36, col: 13, offset: 806},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 17, offset: 810},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 25, offset: 818},
								expr: &seqExpr{
									pos: position{line: 36, col: 26, offset: 819},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 26, offset: 819},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 27, offset: 820},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 39, offset: 832,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 43, offset: 836},
							expr: &litMatcher{
								pos:        position{line: 36, col: 44, offset: 837},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 41, col: 1, offset: 950},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 961},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 961},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 961},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 961},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 963},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 970},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 973},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 978},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 989},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 996},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 997},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 997},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 1000},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1016},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1018},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1022},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1028},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1029},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1029},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1032},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1042},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1538},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1538},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1538},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1540},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1547},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1550},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1555},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1566},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1573},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1574},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1574},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1577},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1593},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1595},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1599},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1605},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1609},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1611},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1617},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1633},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1635},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1640},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1641},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1641},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1645},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1647},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1663},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1667},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1667},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1672},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1674},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1678},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2163},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2163},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2163},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2165},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2172},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2175},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2180},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2191},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2198},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2199},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2199},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2202},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2218},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2220},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2224},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2226},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2231},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2232},
											name: "VariantConstructor",
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
			name: "RecordFieldDefn",
			pos:  position{line: 94, col: 1, offset: 2639},
			expr: &actionExpr{
				pos: position{line: 94, col: 19, offset: 2657},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 94, col: 19, offset: 2657},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 19, offset: 2657},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 24, offset: 2662},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 37, offset: 2675},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 39, offset: 2677},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 43, offset: 2681},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 45, offset: 2683},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 48, offset: 2686},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 98, col: 1, offset: 2780},
			expr: &choiceExpr{
				pos: position{line: 98, col: 22, offset: 2801},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 22, offset: 2801},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 98, col: 22, offset: 2801},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 22, offset: 2801},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 26, offset: 2805},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 28, offset: 2807},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 33, offset: 2812},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 44, offset: 2823},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 46, offset: 2825},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 50, offset: 2829},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 52, offset: 2831},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 58, offset: 2837},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 74, offset: 2853},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 76, offset: 2855},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 98, col: 81, offset: 2860},
										expr: &seqExpr{
											pos: position{line: 98, col: 82, offset: 2861},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 98, col: 82, offset: 2861},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 86, offset: 2865},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 88, offset: 2867},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 104, offset: 2883},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 108, offset: 2887},
									expr: &litMatcher{
										pos:        position{line: 98, col: 108, offset: 2887},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 113, offset: 2892},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 115, offset: 2894},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 119, offset: 2898},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 1, offset: 3503},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 117, col: 1, offset: 3503},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 117, col: 1, offset: 3503},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 5, offset: 3507},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 7, offset: 3509},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 12, offset: 3514},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 23, offset: 3525},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 28, offset: 3530},
										expr: &seqExpr{
											pos: position{line: 117, col: 29, offset: 3531},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 117, col: 29, offset: 3531},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 32, offset: 3534},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 42, offset: 3544},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 134, col: 1, offset: 3981},
			expr: &choiceExpr{
				pos: position{line: 134, col: 11, offset: 3991},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 3991},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 22, offset: 4002},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 136, col: 1, offset: 4017},
			expr: &choiceExpr{
				pos: position{line: 136, col: 14, offset: 4030},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 14, offset: 4030},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 136, col: 14, offset: 4030},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 14, offset: 4030},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 16, offset: 4032},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 22, offset: 4038},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 4041},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 27, offset: 4043},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 38, offset: 4054},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 40, offset: 4056},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 44, offset: 4060},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 46, offset: 4062},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 51, offset: 4067},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 56, offset: 4072},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 142, col: 1, offset: 4191},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 142, col: 1, offset: 4191},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 142, col: 1, offset: 4191},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 142, col: 3, offset: 4193},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 142, col: 9, offset: 4199},
									name: "__",
								},
								&notExpr{
									pos: position{line: 142, col: 12, offset: 4202},
									expr: &ruleRefExpr{
										pos:  position{line: 142, col: 13, offset: 4203},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 146, col: 1, offset: 4311},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 146, col: 1, offset: 4311},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 146, col: 1, offset: 4311},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 146, col: 3, offset: 4313},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 9, offset: 4319},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 146, col: 12, offset: 4322},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 146, col: 14, offset: 4324},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 25, offset: 4335},
									name: "_",
								},
								&notExpr{
									pos: position{line: 146, col: 27, offset: 4337},
									expr: &litMatcher{
										pos:        position{line: 146, col: 28, offset: 4338},
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
			pos:  position{line: 150, col: 1, offset: 4432},
			expr: &actionExpr{
				pos: position{line: 150, col: 12, offset: 4443},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 150, col: 12, offset: 4443},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 150, col: 12, offset: 4443},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 14, offset: 4445},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 20, offset: 4451},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 23, offset: 4454},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 25, offset: 4456},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 38, offset: 4469},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 40, offset: 4471},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 44, offset: 4475},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 46, offset: 4477},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 53, offset: 4484},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 56, offset: 4487},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 150, col: 60, offset: 4491},
								expr: &seqExpr{
									pos: position{line: 150, col: 61, offset: 4492},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 150, col: 61, offset: 4492},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 150, col: 74, offset: 4505},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 79, offset: 4510},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 81, offset: 4512},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 85, offset: 4516},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 88, offset: 4519},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 150, col: 99, offset: 4530},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 100, offset: 4531},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 112, offset: 4543},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 114, offset: 4545},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 118, offset: 4549},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 173, col: 1, offset: 5205},
			expr: &actionExpr{
				pos: position{line: 173, col: 8, offset: 5212},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 173, col: 8, offset: 5212},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 173, col: 12, offset: 5216},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 173, col: 12, offset: 5216},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 173, col: 21, offset: 5225},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 173, col: 28, offset: 5232},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 179, col: 1, offset: 5349},
			expr: &choiceExpr{
				pos: position{line: 179, col: 10, offset: 5358},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 179, col: 10, offset: 5358},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 179, col: 10, offset: 5358},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 179, col: 10, offset: 5358},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 15, offset: 5363},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 18, offset: 5366},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 23, offset: 5371},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 33, offset: 5381},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 35, offset: 5383},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 39, offset: 5387},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 41, offset: 5389},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 179, col: 47, offset: 5395},
										expr: &ruleRefExpr{
											pos:  position{line: 179, col: 48, offset: 5396},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 60, offset: 5408},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 62, offset: 5410},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 66, offset: 5414},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 68, offset: 5416},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 75, offset: 5423},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 77, offset: 5425},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 85, offset: 5433},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 191, col: 1, offset: 5763},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 191, col: 1, offset: 5763},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 191, col: 1, offset: 5763},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 6, offset: 5768},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 9, offset: 5771},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 14, offset: 5776},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 24, offset: 5786},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 26, offset: 5788},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 30, offset: 5792},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 32, offset: 5794},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 191, col: 38, offset: 5800},
										expr: &ruleRefExpr{
											pos:  position{line: 191, col: 39, offset: 5801},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 51, offset: 5813},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 191, col: 54, offset: 5816},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 58, offset: 5820},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 60, offset: 5822},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 67, offset: 5829},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 69, offset: 5831},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 73, offset: 5835},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 75, offset: 5837},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 191, col: 81, offset: 5843},
										expr: &ruleRefExpr{
											pos:  position{line: 191, col: 82, offset: 5844},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 94, offset: 5856},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 191, col: 97, offset: 5859},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 210, col: 1, offset: 6362},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 210, col: 1, offset: 6362},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 210, col: 1, offset: 6362},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 6, offset: 6367},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 210, col: 9, offset: 6370},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 210, col: 14, offset: 6375},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 24, offset: 6385},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 210, col: 26, offset: 6387},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 30, offset: 6391},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 210, col: 32, offset: 6393},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 210, col: 38, offset: 6399},
										expr: &ruleRefExpr{
											pos:  position{line: 210, col: 39, offset: 6400},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 51, offset: 6412},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 210, col: 54, offset: 6415},
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
			pos:  position{line: 222, col: 1, offset: 6713},
			expr: &choiceExpr{
				pos: position{line: 222, col: 8, offset: 6720},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 222, col: 8, offset: 6720},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 222, col: 8, offset: 6720},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 222, col: 8, offset: 6720},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 222, col: 10, offset: 6722},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 17, offset: 6729},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 222, col: 28, offset: 6740},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 222, col: 32, offset: 6744},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 35, offset: 6747},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 222, col: 48, offset: 6760},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 53, offset: 6765},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 222, col: 63, offset: 6775},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 236, col: 1, offset: 7077},
						run: (*parser).callonCall13,
						expr: &seqExpr{
							pos: position{line: 236, col: 1, offset: 7077},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 236, col: 1, offset: 7077},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 3, offset: 7079},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 6, offset: 7082},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 236, col: 19, offset: 7095},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 24, offset: 7100},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 34, offset: 7110},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 250, col: 1, offset: 7402},
			expr: &choiceExpr{
				pos: position{line: 250, col: 13, offset: 7414},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 250, col: 13, offset: 7414},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 250, col: 13, offset: 7414},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 250, col: 13, offset: 7414},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 17, offset: 7418},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 19, offset: 7420},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 28, offset: 7429},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 40, offset: 7441},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 42, offset: 7443},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 250, col: 47, offset: 7448},
										expr: &seqExpr{
											pos: position{line: 250, col: 48, offset: 7449},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 250, col: 48, offset: 7449},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 250, col: 52, offset: 7453},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 250, col: 54, offset: 7455},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 68, offset: 7469},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 250, col: 70, offset: 7471},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 267, col: 1, offset: 7893},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 267, col: 1, offset: 7893},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 267, col: 1, offset: 7893},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 5, offset: 7897},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 267, col: 7, offset: 7899},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 267, col: 16, offset: 7908},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 21, offset: 7913},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 267, col: 23, offset: 7915},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 272, col: 1, offset: 8021},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 274, col: 1, offset: 8027},
			expr: &actionExpr{
				pos: position{line: 274, col: 16, offset: 8042},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 274, col: 16, offset: 8042},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 274, col: 16, offset: 8042},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 274, col: 18, offset: 8044},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 274, col: 21, offset: 8047},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 274, col: 27, offset: 8053},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 274, col: 32, offset: 8058},
								expr: &seqExpr{
									pos: position{line: 274, col: 33, offset: 8059},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 274, col: 33, offset: 8059},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 274, col: 36, offset: 8062},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 274, col: 45, offset: 8071},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 274, col: 48, offset: 8074},
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
			pos:  position{line: 294, col: 1, offset: 8680},
			expr: &choiceExpr{
				pos: position{line: 294, col: 9, offset: 8688},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 294, col: 9, offset: 8688},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 21, offset: 8700},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 37, offset: 8716},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 48, offset: 8727},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 60, offset: 8739},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 296, col: 1, offset: 8752},
			expr: &actionExpr{
				pos: position{line: 296, col: 13, offset: 8764},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 296, col: 13, offset: 8764},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 296, col: 13, offset: 8764},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 296, col: 15, offset: 8766},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 21, offset: 8772},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 35, offset: 8786},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 296, col: 40, offset: 8791},
								expr: &seqExpr{
									pos: position{line: 296, col: 41, offset: 8792},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 296, col: 41, offset: 8792},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 44, offset: 8795},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 60, offset: 8811},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 63, offset: 8814},
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
			pos:  position{line: 315, col: 1, offset: 9395},
			expr: &actionExpr{
				pos: position{line: 315, col: 17, offset: 9411},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 315, col: 17, offset: 9411},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 315, col: 17, offset: 9411},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 315, col: 19, offset: 9413},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 315, col: 25, offset: 9419},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 315, col: 34, offset: 9428},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 315, col: 39, offset: 9433},
								expr: &seqExpr{
									pos: position{line: 315, col: 40, offset: 9434},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 315, col: 40, offset: 9434},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 43, offset: 9437},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 60, offset: 9454},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 63, offset: 9457},
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
			pos:  position{line: 335, col: 1, offset: 10036},
			expr: &actionExpr{
				pos: position{line: 335, col: 12, offset: 10047},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 335, col: 12, offset: 10047},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 335, col: 12, offset: 10047},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 335, col: 14, offset: 10049},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 335, col: 20, offset: 10055},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 335, col: 30, offset: 10065},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 335, col: 35, offset: 10070},
								expr: &seqExpr{
									pos: position{line: 335, col: 36, offset: 10071},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 335, col: 36, offset: 10071},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 39, offset: 10074},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 51, offset: 10086},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 54, offset: 10089},
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
			pos:  position{line: 355, col: 1, offset: 10665},
			expr: &actionExpr{
				pos: position{line: 355, col: 13, offset: 10677},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 355, col: 13, offset: 10677},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 355, col: 13, offset: 10677},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 355, col: 15, offset: 10679},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 355, col: 21, offset: 10685},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 355, col: 33, offset: 10697},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 355, col: 38, offset: 10702},
								expr: &seqExpr{
									pos: position{line: 355, col: 39, offset: 10703},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 355, col: 39, offset: 10703},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 42, offset: 10706},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 55, offset: 10719},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 58, offset: 10722},
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
			pos:  position{line: 374, col: 1, offset: 11300},
			expr: &choiceExpr{
				pos: position{line: 374, col: 15, offset: 11314},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 374, col: 15, offset: 11314},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 374, col: 15, offset: 11314},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 374, col: 15, offset: 11314},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 374, col: 19, offset: 11318},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 374, col: 21, offset: 11320},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 374, col: 27, offset: 11326},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 374, col: 33, offset: 11332},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 374, col: 35, offset: 11334},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 377, col: 5, offset: 11457},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 379, col: 1, offset: 11464},
			expr: &choiceExpr{
				pos: position{line: 379, col: 12, offset: 11475},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 379, col: 12, offset: 11475},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 379, col: 30, offset: 11493},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 379, col: 49, offset: 11512},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 379, col: 64, offset: 11527},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 381, col: 1, offset: 11540},
			expr: &actionExpr{
				pos: position{line: 381, col: 19, offset: 11558},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 381, col: 21, offset: 11560},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 381, col: 21, offset: 11560},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 381, col: 29, offset: 11568},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 381, col: 36, offset: 11575},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 385, col: 1, offset: 11674},
			expr: &actionExpr{
				pos: position{line: 385, col: 20, offset: 11693},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 385, col: 22, offset: 11695},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 385, col: 22, offset: 11695},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 29, offset: 11702},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 36, offset: 11709},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 42, offset: 11715},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 48, offset: 11721},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 56, offset: 11729},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 389, col: 1, offset: 11835},
			expr: &choiceExpr{
				pos: position{line: 389, col: 16, offset: 11850},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 389, col: 16, offset: 11850},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 389, col: 18, offset: 11852},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 389, col: 18, offset: 11852},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 389, col: 25, offset: 11859},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 392, col: 3, offset: 11965},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 392, col: 5, offset: 11967},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 392, col: 5, offset: 11967},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 392, col: 11, offset: 11973},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 392, col: 17, offset: 11979},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 395, col: 3, offset: 12082},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 395, col: 3, offset: 12082},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 399, col: 1, offset: 12186},
			expr: &choiceExpr{
				pos: position{line: 399, col: 15, offset: 12200},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 399, col: 15, offset: 12200},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 399, col: 17, offset: 12202},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 399, col: 17, offset: 12202},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 399, col: 24, offset: 12209},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 402, col: 3, offset: 12315},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 402, col: 5, offset: 12317},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 402, col: 5, offset: 12317},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 402, col: 11, offset: 12323},
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
			pos:  position{line: 406, col: 1, offset: 12425},
			expr: &choiceExpr{
				pos: position{line: 406, col: 9, offset: 12433},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 406, col: 9, offset: 12433},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 406, col: 24, offset: 12448},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 408, col: 1, offset: 12455},
			expr: &choiceExpr{
				pos: position{line: 408, col: 14, offset: 12468},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 14, offset: 12468},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 29, offset: 12483},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 410, col: 1, offset: 12491},
			expr: &choiceExpr{
				pos: position{line: 410, col: 14, offset: 12504},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 410, col: 14, offset: 12504},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 410, col: 29, offset: 12519},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 412, col: 1, offset: 12531},
			expr: &actionExpr{
				pos: position{line: 412, col: 16, offset: 12546},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 412, col: 16, offset: 12546},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 412, col: 16, offset: 12546},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 20, offset: 12550},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 412, col: 22, offset: 12552},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 412, col: 28, offset: 12558},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 33, offset: 12563},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 412, col: 35, offset: 12565},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 412, col: 40, offset: 12570},
								expr: &seqExpr{
									pos: position{line: 412, col: 41, offset: 12571},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 412, col: 41, offset: 12571},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 412, col: 45, offset: 12575},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 412, col: 47, offset: 12577},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 412, col: 52, offset: 12582},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 412, col: 56, offset: 12586},
							expr: &litMatcher{
								pos:        position{line: 412, col: 56, offset: 12586},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 61, offset: 12591},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 412, col: 63, offset: 12593},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 428, col: 1, offset: 13038},
			expr: &actionExpr{
				pos: position{line: 428, col: 19, offset: 13056},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 428, col: 19, offset: 13056},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 428, col: 19, offset: 13056},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 428, col: 24, offset: 13061},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 428, col: 35, offset: 13072},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 428, col: 37, offset: 13074},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 428, col: 42, offset: 13079},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 441, col: 1, offset: 13349},
			expr: &actionExpr{
				pos: position{line: 441, col: 18, offset: 13366},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 441, col: 18, offset: 13366},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 441, col: 18, offset: 13366},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 441, col: 23, offset: 13371},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 34, offset: 13382},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 441, col: 36, offset: 13384},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 40, offset: 13388},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 441, col: 42, offset: 13390},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 441, col: 52, offset: 13400},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 65, offset: 13413},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 441, col: 67, offset: 13415},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 71, offset: 13419},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 441, col: 73, offset: 13421},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 441, col: 84, offset: 13432},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 441, col: 89, offset: 13437},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 441, col: 94, offset: 13442},
								expr: &seqExpr{
									pos: position{line: 441, col: 95, offset: 13443},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 441, col: 95, offset: 13443},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 99, offset: 13447},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 101, offset: 13449},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 114, offset: 13462},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 441, col: 116, offset: 13464},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 120, offset: 13468},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 122, offset: 13470},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 441, col: 130, offset: 13478},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 461, col: 1, offset: 14062},
			expr: &actionExpr{
				pos: position{line: 461, col: 17, offset: 14078},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 461, col: 17, offset: 14078},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 461, col: 17, offset: 14078},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 461, col: 22, offset: 14083},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 465, col: 1, offset: 14156},
			expr: &actionExpr{
				pos: position{line: 465, col: 16, offset: 14171},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 465, col: 16, offset: 14171},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 465, col: 16, offset: 14171},
							expr: &ruleRefExpr{
								pos:  position{line: 465, col: 17, offset: 14172},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 465, col: 27, offset: 14182},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 465, col: 27, offset: 14182},
									expr: &charClassMatcher{
										pos:        position{line: 465, col: 27, offset: 14182},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 465, col: 34, offset: 14189},
									expr: &charClassMatcher{
										pos:        position{line: 465, col: 34, offset: 14189},
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
			pos:  position{line: 469, col: 1, offset: 14264},
			expr: &actionExpr{
				pos: position{line: 469, col: 14, offset: 14277},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 469, col: 15, offset: 14278},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 469, col: 15, offset: 14278},
							expr: &charClassMatcher{
								pos:        position{line: 469, col: 15, offset: 14278},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 469, col: 22, offset: 14285},
							expr: &charClassMatcher{
								pos:        position{line: 469, col: 22, offset: 14285},
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
			pos:  position{line: 473, col: 1, offset: 14360},
			expr: &choiceExpr{
				pos: position{line: 473, col: 9, offset: 14368},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 473, col: 9, offset: 14368},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 473, col: 9, offset: 14368},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 473, col: 9, offset: 14368},
									expr: &litMatcher{
										pos:        position{line: 473, col: 9, offset: 14368},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 473, col: 14, offset: 14373},
									expr: &charClassMatcher{
										pos:        position{line: 473, col: 14, offset: 14373},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 473, col: 21, offset: 14380},
									expr: &litMatcher{
										pos:        position{line: 473, col: 22, offset: 14381},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 480, col: 3, offset: 14556},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 480, col: 3, offset: 14556},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 480, col: 3, offset: 14556},
									expr: &litMatcher{
										pos:        position{line: 480, col: 3, offset: 14556},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 480, col: 8, offset: 14561},
									expr: &charClassMatcher{
										pos:        position{line: 480, col: 8, offset: 14561},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 480, col: 15, offset: 14568},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 480, col: 19, offset: 14572},
									expr: &charClassMatcher{
										pos:        position{line: 480, col: 19, offset: 14572},
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
						pos:        position{line: 487, col: 3, offset: 14761},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 487, col: 12, offset: 14770},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 487, col: 12, offset: 14770},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 493, col: 3, offset: 14971},
						name: "Unit",
					},
					&actionExpr{
						pos: position{line: 494, col: 3, offset: 14978},
						run: (*parser).callonConst23,
						expr: &seqExpr{
							pos: position{line: 494, col: 3, offset: 14978},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 494, col: 3, offset: 14978},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 494, col: 7, offset: 14982},
									expr: &seqExpr{
										pos: position{line: 494, col: 8, offset: 14983},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 494, col: 8, offset: 14983},
												expr: &ruleRefExpr{
													pos:  position{line: 494, col: 9, offset: 14984},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 494, col: 21, offset: 14996,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 494, col: 25, offset: 15000},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 501, col: 3, offset: 15184},
						run: (*parser).callonConst32,
						expr: &seqExpr{
							pos: position{line: 501, col: 3, offset: 15184},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 501, col: 3, offset: 15184},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 501, col: 7, offset: 15188},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 501, col: 12, offset: 15193},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 501, col: 12, offset: 15193},
												expr: &ruleRefExpr{
													pos:  position{line: 501, col: 13, offset: 15194},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 501, col: 25, offset: 15206,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 501, col: 28, offset: 15209},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 5, offset: 15301},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 20, offset: 15316},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 37, offset: 15333},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 505, col: 1, offset: 15350},
			expr: &actionExpr{
				pos: position{line: 505, col: 8, offset: 15357},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 505, col: 8, offset: 15357},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 509, col: 1, offset: 15420},
			expr: &actionExpr{
				pos: position{line: 509, col: 10, offset: 15429},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 509, col: 11, offset: 15430},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 513, col: 1, offset: 15485},
			expr: &seqExpr{
				pos: position{line: 513, col: 12, offset: 15496},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 513, col: 13, offset: 15497},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 513, col: 13, offset: 15497},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 21, offset: 15505},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 28, offset: 15512},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 37, offset: 15521},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 46, offset: 15530},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 55, offset: 15539},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 64, offset: 15548},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 74, offset: 15558},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 86, offset: 15570},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 513, col: 95, offset: 15579},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 513, col: 105, offset: 15589},
						expr: &oneOrMoreExpr{
							pos: position{line: 513, col: 106, offset: 15590},
							expr: &charClassMatcher{
								pos:        position{line: 513, col: 106, offset: 15590},
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
			name: "BaseType",
			pos:  position{line: 515, col: 1, offset: 15598},
			expr: &actionExpr{
				pos: position{line: 515, col: 12, offset: 15609},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 515, col: 14, offset: 15611},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 515, col: 14, offset: 15611},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 22, offset: 15619},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 31, offset: 15628},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 42, offset: 15639},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 51, offset: 15648},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 60, offset: 15657},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 70, offset: 15667},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 519, col: 1, offset: 15766},
			expr: &charClassMatcher{
				pos:        position{line: 519, col: 15, offset: 15780},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 521, col: 1, offset: 15796},
			expr: &choiceExpr{
				pos: position{line: 521, col: 18, offset: 15813},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 521, col: 18, offset: 15813},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 521, col: 37, offset: 15832},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 523, col: 1, offset: 15847},
			expr: &charClassMatcher{
				pos:        position{line: 523, col: 20, offset: 15866},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 525, col: 1, offset: 15879},
			expr: &charClassMatcher{
				pos:        position{line: 525, col: 16, offset: 15894},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 527, col: 1, offset: 15901},
			expr: &charClassMatcher{
				pos:        position{line: 527, col: 23, offset: 15923},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 529, col: 1, offset: 15930},
			expr: &charClassMatcher{
				pos:        position{line: 529, col: 12, offset: 15941},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 531, col: 1, offset: 15952},
			expr: &choiceExpr{
				pos: position{line: 531, col: 22, offset: 15973},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 531, col: 22, offset: 15973},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 531, col: 32, offset: 15983},
						expr: &charClassMatcher{
							pos:        position{line: 531, col: 32, offset: 15983},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 533, col: 1, offset: 15995},
			expr: &zeroOrMoreExpr{
				pos: position{line: 533, col: 18, offset: 16012},
				expr: &charClassMatcher{
					pos:        position{line: 533, col: 18, offset: 16012},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 535, col: 1, offset: 16024},
			expr: &notExpr{
				pos: position{line: 535, col: 7, offset: 16030},
				expr: &anyMatcher{
					line: 535, col: 8, offset: 16031,
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
		return Module{Name: "", Subvalues: subvalues}, nil
	} else {
		return Module{Name: "", Subvalues: []Ast{stat.(Ast)}}, nil
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
	return Comment{StringValue: string(c.text[1:])}, nil
}

func (p *parser) callonComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment1(stack["comment"])
}

func (c *current) onTypeDefn2(name, params, types interface{}) (interface{}, error) {
	// alias type
	parameters := []Ast{}
	fields := []Ast{}

	vals := types.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(types)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			fields = append(fields, v)
		}
	}

	return AliasType{Name: name.(Identifier).StringValue, Params: parameters, Types: fields}, nil
}

func (p *parser) callonTypeDefn2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn2(stack["name"], stack["params"], stack["types"])
}

func (c *current) onTypeDefn22(name, params, first, rest interface{}) (interface{}, error) {
	// record type
	fields := []RecordField{first.(RecordField)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(RecordField)
			fields = append(fields, v)
		}
	}

	return RecordType{Name: name.(Identifier).StringValue, Fields: fields}, nil
}

func (p *parser) callonTypeDefn22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn22(stack["name"], stack["params"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn54(name, params, rest interface{}) (interface{}, error) {
	// variant type
	parameters := []Ast{}
	constructors := []VariantConstructor{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		for _, v := range vals {
			constructors = append(constructors, v.(VariantConstructor))
		}
	}

	return VariantType{Name: name.(Identifier).StringValue, Params: parameters, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn54() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn54(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{Name: name.(Identifier).StringValue, Type: t.(Ast)}, nil
}

func (p *parser) callonRecordFieldDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordFieldDefn1(stack["name"], stack["t"])
}

func (c *current) onVariantConstructor2(name, first, rest interface{}) (interface{}, error) {
	// variant constructor with a record type
	fields := []RecordField{first.(RecordField)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(RecordField)
			fields = append(fields, v)
		}
	}

	record := RecordType{Name: name.(Identifier).StringValue, Fields: fields}
	return VariantConstructor{Name: name.(Identifier).StringValue, Fields: []Ast{record}}, nil
}

func (p *parser) callonVariantConstructor2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor2(stack["name"], stack["first"], stack["rest"])
}

func (c *current) onVariantConstructor26(name, rest interface{}) (interface{}, error) {
	params := []Ast{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			params = append(params, v)
		}
	}

	return VariantConstructor{Name: name.(Identifier).StringValue, Fields: params}, nil
}

func (p *parser) callonVariantConstructor26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor26(stack["name"], stack["rest"])
}

func (c *current) onAssignment2(i, expr interface{}) (interface{}, error) {
	//fmt.Println("assignment:", string(c.text))

	return Assignment{Left: i.(Ast), Right: expr.(Ast)}, nil
}

func (p *parser) callonAssignment2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment2(stack["i"], stack["expr"])
}

func (c *current) onAssignment15() (interface{}, error) {
	return nil, errors.New("Variable name or '_' (unused result character) required here")
}

func (p *parser) callonAssignment15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment15()
}

func (c *current) onAssignment22(i interface{}) (interface{}, error) {
	return nil, errors.New("When assigning a value to a variable, you must use '='")
}

func (p *parser) callonAssignment22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment22(stack["i"])
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
	return Func{Name: i.(Identifier).StringValue, Arguments: args, Subvalues: subvalues}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["i"], stack["ids"], stack["statements"])
}

func (c *current) onExpr1(ex interface{}) (interface{}, error) {
	//fmt.Printf("top-level expr: %s\n", string(c.text))
	//fmt.Println(ex)
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

func (c *current) onCall2(module, fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: module.(Ast), Function: fn.(Ast), Arguments: arguments}, nil
}

func (p *parser) callonCall2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall2(stack["module"], stack["fn"], stack["args"])
}

func (c *current) onCall13(fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: nil, Function: fn.(Ast), Arguments: arguments}, nil
}

func (p *parser) callonCall13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall13(stack["fn"], stack["args"])
}

func (c *current) onArguments2(argument, rest interface{}) (interface{}, error) {
	args := []Ast{argument.(Ast)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			args = append(args, v)
		}
	}

	return Container{Type: "Arguments", Subvalues: args}, nil
}

func (p *parser) callonArguments2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments2(stack["argument"], stack["rest"])
}

func (c *current) onArguments17(argument interface{}) (interface{}, error) {
	args := []Ast{argument.(Ast)}
	return Container{Type: "Arguments", Subvalues: args}, nil
}

func (p *parser) callonArguments17() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments17(stack["argument"])
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

		return Expr{Type: "Compound", Subvalues: subvalues}, nil
	} else {
		return Expr{Type: "Compound", Subvalues: []Ast{op.(Ast)}}, nil
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
		return Expr{Type: "BinOpBool", Subvalues: subvalues}, nil
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
		return Expr{Type: "BinOpEquality", Subvalues: subvalues}, nil
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
		return Expr{Type: "BinOpLow", Subvalues: subvalues}, nil
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
		return Expr{Type: "BinOpHigh", Subvalues: subvalues}, nil
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
	return Expr{Type: "BinOpParens", Subvalues: []Ast{first.(Ast)}}, nil
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
	return ArrayType{Subvalues: subvalues}, nil
}

func (p *parser) callonArrayLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArrayLiteral1(stack["first"], stack["rest"])
}

func (c *current) onVariantInstance1(name, args interface{}) (interface{}, error) {
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return VariantInstance{Name: name.(BasicAst).StringValue, Arguments: arguments}, nil
}

func (p *parser) callonVariantInstance1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantInstance1(stack["name"], stack["args"])
}

func (c *current) onRecordInstance1(name, firstName, firstValue, rest interface{}) (interface{}, error) {
	instance := RecordInstance{Name: name.(BasicAst).StringValue}
	instance.Values = make(map[string]Ast)

	vals := rest.([]interface{})
	instance.Values[firstName.(BasicAst).StringValue] = firstValue.(Ast)

	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			k := restExpr[2].(BasicAst).StringValue
			v := restExpr[6].(Ast)
			instance.Values[k] = v
		}
	}
	return instance, nil
}

func (p *parser) callonRecordInstance1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordInstance1(stack["name"], stack["firstName"], stack["firstValue"], stack["rest"])
}

func (c *current) onTypeParameter1() (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
}

func (p *parser) callonTypeParameter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeParameter1()
}

func (c *current) onVariableName1() (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
}

func (p *parser) callonVariableName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableName1()
}

func (c *current) onModuleName1() (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
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

func (c *current) onConst23() (interface{}, error) {
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonConst23() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst23()
}

func (c *current) onConst32(val interface{}) (interface{}, error) {
	return BasicAst{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst32() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst32(stack["val"])
}

func (c *current) onUnit1() (interface{}, error) {
	return BasicAst{Type: "Unit", ValueType: NIL}, nil
}

func (p *parser) callonUnit1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnit1()
}

func (c *current) onUnused1() (interface{}, error) {
	return Identifier{StringValue: "_"}, nil
}

func (p *parser) callonUnused1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnused1()
}

func (c *current) onBaseType1() (interface{}, error) {
	return BasicAst{Type: "Type", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonBaseType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBaseType1()
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
