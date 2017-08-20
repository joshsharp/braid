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
						pos: position{line: 60, col: 1, offset: 1536},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1536},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1536},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1538},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1545},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1548},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1553},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1564},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1571},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1572},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1572},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1575},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1591},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1593},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1597},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1603},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1607},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1609},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1615},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1631},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1633},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1638},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1639},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1639},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1643},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1645},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1661},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1665},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1665},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1670},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1672},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1676},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2159},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2159},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2159},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2161},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2168},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2171},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2176},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2187},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2194},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2195},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2195},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2198},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2214},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2216},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2220},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2222},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2227},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2228},
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
			pos:  position{line: 94, col: 1, offset: 2633},
			expr: &actionExpr{
				pos: position{line: 94, col: 19, offset: 2651},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 94, col: 19, offset: 2651},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 19, offset: 2651},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 24, offset: 2656},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 37, offset: 2669},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 39, offset: 2671},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 43, offset: 2675},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 45, offset: 2677},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 48, offset: 2680},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 98, col: 1, offset: 2772},
			expr: &choiceExpr{
				pos: position{line: 98, col: 22, offset: 2793},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 22, offset: 2793},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 98, col: 22, offset: 2793},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 22, offset: 2793},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 26, offset: 2797},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 28, offset: 2799},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 33, offset: 2804},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 44, offset: 2815},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 46, offset: 2817},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 50, offset: 2821},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 52, offset: 2823},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 58, offset: 2829},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 74, offset: 2845},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 76, offset: 2847},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 98, col: 81, offset: 2852},
										expr: &seqExpr{
											pos: position{line: 98, col: 82, offset: 2853},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 98, col: 82, offset: 2853},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 86, offset: 2857},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 88, offset: 2859},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 104, offset: 2875},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 108, offset: 2879},
									expr: &litMatcher{
										pos:        position{line: 98, col: 108, offset: 2879},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 113, offset: 2884},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 115, offset: 2886},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 119, offset: 2890},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 1, offset: 3491},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 117, col: 1, offset: 3491},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 117, col: 1, offset: 3491},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 5, offset: 3495},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 7, offset: 3497},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 12, offset: 3502},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 23, offset: 3513},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 28, offset: 3518},
										expr: &seqExpr{
											pos: position{line: 117, col: 29, offset: 3519},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 117, col: 29, offset: 3519},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 32, offset: 3522},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 42, offset: 3532},
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
			pos:  position{line: 134, col: 1, offset: 3967},
			expr: &choiceExpr{
				pos: position{line: 134, col: 11, offset: 3977},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 3977},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 22, offset: 3988},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 136, col: 1, offset: 4003},
			expr: &choiceExpr{
				pos: position{line: 136, col: 14, offset: 4016},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 14, offset: 4016},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 136, col: 14, offset: 4016},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 14, offset: 4016},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 16, offset: 4018},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 22, offset: 4024},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 4027},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 27, offset: 4029},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 38, offset: 4040},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 40, offset: 4042},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 44, offset: 4046},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 46, offset: 4048},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 51, offset: 4053},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 56, offset: 4058},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 142, col: 1, offset: 4177},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 142, col: 1, offset: 4177},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 142, col: 1, offset: 4177},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 142, col: 3, offset: 4179},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 142, col: 9, offset: 4185},
									name: "__",
								},
								&notExpr{
									pos: position{line: 142, col: 12, offset: 4188},
									expr: &ruleRefExpr{
										pos:  position{line: 142, col: 13, offset: 4189},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 146, col: 1, offset: 4297},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 146, col: 1, offset: 4297},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 146, col: 1, offset: 4297},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 146, col: 3, offset: 4299},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 9, offset: 4305},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 146, col: 12, offset: 4308},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 146, col: 14, offset: 4310},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 25, offset: 4321},
									name: "_",
								},
								&notExpr{
									pos: position{line: 146, col: 27, offset: 4323},
									expr: &litMatcher{
										pos:        position{line: 146, col: 28, offset: 4324},
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
			pos:  position{line: 150, col: 1, offset: 4418},
			expr: &actionExpr{
				pos: position{line: 150, col: 12, offset: 4429},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 150, col: 12, offset: 4429},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 150, col: 12, offset: 4429},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 14, offset: 4431},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 20, offset: 4437},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 23, offset: 4440},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 25, offset: 4442},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 38, offset: 4455},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 40, offset: 4457},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 44, offset: 4461},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 46, offset: 4463},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 53, offset: 4470},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 56, offset: 4473},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 150, col: 60, offset: 4477},
								expr: &seqExpr{
									pos: position{line: 150, col: 61, offset: 4478},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 150, col: 61, offset: 4478},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 150, col: 74, offset: 4491},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 79, offset: 4496},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 81, offset: 4498},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 85, offset: 4502},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 88, offset: 4505},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 150, col: 99, offset: 4516},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 100, offset: 4517},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 112, offset: 4529},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 114, offset: 4531},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 118, offset: 4535},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 173, col: 1, offset: 5189},
			expr: &actionExpr{
				pos: position{line: 173, col: 8, offset: 5196},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 173, col: 8, offset: 5196},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 173, col: 12, offset: 5200},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 173, col: 12, offset: 5200},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 173, col: 21, offset: 5209},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 173, col: 28, offset: 5216},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 179, col: 1, offset: 5333},
			expr: &choiceExpr{
				pos: position{line: 179, col: 10, offset: 5342},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 179, col: 10, offset: 5342},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 179, col: 10, offset: 5342},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 179, col: 10, offset: 5342},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 15, offset: 5347},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 18, offset: 5350},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 23, offset: 5355},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 33, offset: 5365},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 35, offset: 5367},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 39, offset: 5371},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 41, offset: 5373},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 179, col: 47, offset: 5379},
										expr: &ruleRefExpr{
											pos:  position{line: 179, col: 48, offset: 5380},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 60, offset: 5392},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 62, offset: 5394},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 66, offset: 5398},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 68, offset: 5400},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 75, offset: 5407},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 77, offset: 5409},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 85, offset: 5417},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 191, col: 1, offset: 5747},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 191, col: 1, offset: 5747},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 191, col: 1, offset: 5747},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 6, offset: 5752},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 9, offset: 5755},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 14, offset: 5760},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 24, offset: 5770},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 26, offset: 5772},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 30, offset: 5776},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 32, offset: 5778},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 191, col: 38, offset: 5784},
										expr: &ruleRefExpr{
											pos:  position{line: 191, col: 39, offset: 5785},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 51, offset: 5797},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 191, col: 54, offset: 5800},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 58, offset: 5804},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 60, offset: 5806},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 67, offset: 5813},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 69, offset: 5815},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 73, offset: 5819},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 75, offset: 5821},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 191, col: 81, offset: 5827},
										expr: &ruleRefExpr{
											pos:  position{line: 191, col: 82, offset: 5828},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 94, offset: 5840},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 191, col: 97, offset: 5843},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 210, col: 1, offset: 6346},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 210, col: 1, offset: 6346},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 210, col: 1, offset: 6346},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 6, offset: 6351},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 210, col: 9, offset: 6354},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 210, col: 14, offset: 6359},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 24, offset: 6369},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 210, col: 26, offset: 6371},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 30, offset: 6375},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 210, col: 32, offset: 6377},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 210, col: 38, offset: 6383},
										expr: &ruleRefExpr{
											pos:  position{line: 210, col: 39, offset: 6384},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 51, offset: 6396},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 210, col: 54, offset: 6399},
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
			pos:  position{line: 222, col: 1, offset: 6697},
			expr: &choiceExpr{
				pos: position{line: 222, col: 8, offset: 6704},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 222, col: 8, offset: 6704},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 222, col: 8, offset: 6704},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 222, col: 8, offset: 6704},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 222, col: 10, offset: 6706},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 17, offset: 6713},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 222, col: 28, offset: 6724},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 222, col: 32, offset: 6728},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 35, offset: 6731},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 222, col: 48, offset: 6744},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 53, offset: 6749},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 222, col: 63, offset: 6759},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 236, col: 1, offset: 7061},
						run: (*parser).callonCall13,
						expr: &seqExpr{
							pos: position{line: 236, col: 1, offset: 7061},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 236, col: 1, offset: 7061},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 3, offset: 7063},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 6, offset: 7066},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 236, col: 19, offset: 7079},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 24, offset: 7084},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 34, offset: 7094},
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
			pos:  position{line: 250, col: 1, offset: 7386},
			expr: &choiceExpr{
				pos: position{line: 250, col: 13, offset: 7398},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 250, col: 13, offset: 7398},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 250, col: 13, offset: 7398},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 250, col: 13, offset: 7398},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 17, offset: 7402},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 19, offset: 7404},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 28, offset: 7413},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 40, offset: 7425},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 42, offset: 7427},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 250, col: 47, offset: 7432},
										expr: &seqExpr{
											pos: position{line: 250, col: 48, offset: 7433},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 250, col: 48, offset: 7433},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 250, col: 52, offset: 7437},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 250, col: 54, offset: 7439},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 68, offset: 7453},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 250, col: 70, offset: 7455},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 267, col: 1, offset: 7877},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 267, col: 1, offset: 7877},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 267, col: 1, offset: 7877},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 5, offset: 7881},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 267, col: 7, offset: 7883},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 267, col: 16, offset: 7892},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 21, offset: 7897},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 267, col: 23, offset: 7899},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 272, col: 1, offset: 8005},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 274, col: 1, offset: 8011},
			expr: &actionExpr{
				pos: position{line: 274, col: 16, offset: 8026},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 274, col: 16, offset: 8026},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 274, col: 16, offset: 8026},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 274, col: 18, offset: 8028},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 274, col: 21, offset: 8031},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 274, col: 27, offset: 8037},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 274, col: 32, offset: 8042},
								expr: &seqExpr{
									pos: position{line: 274, col: 33, offset: 8043},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 274, col: 33, offset: 8043},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 274, col: 36, offset: 8046},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 274, col: 45, offset: 8055},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 274, col: 48, offset: 8058},
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
			pos:  position{line: 294, col: 1, offset: 8682},
			expr: &choiceExpr{
				pos: position{line: 294, col: 9, offset: 8690},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 294, col: 9, offset: 8690},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 21, offset: 8702},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 37, offset: 8718},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 48, offset: 8729},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 60, offset: 8741},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 296, col: 1, offset: 8754},
			expr: &actionExpr{
				pos: position{line: 296, col: 13, offset: 8766},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 296, col: 13, offset: 8766},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 296, col: 13, offset: 8766},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 296, col: 15, offset: 8768},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 21, offset: 8774},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 35, offset: 8788},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 296, col: 40, offset: 8793},
								expr: &seqExpr{
									pos: position{line: 296, col: 41, offset: 8794},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 296, col: 41, offset: 8794},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 44, offset: 8797},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 60, offset: 8813},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 63, offset: 8816},
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
			pos:  position{line: 315, col: 1, offset: 9402},
			expr: &actionExpr{
				pos: position{line: 315, col: 17, offset: 9418},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 315, col: 17, offset: 9418},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 315, col: 17, offset: 9418},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 315, col: 19, offset: 9420},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 315, col: 25, offset: 9426},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 315, col: 34, offset: 9435},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 315, col: 39, offset: 9440},
								expr: &seqExpr{
									pos: position{line: 315, col: 40, offset: 9441},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 315, col: 40, offset: 9441},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 43, offset: 9444},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 60, offset: 9461},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 63, offset: 9464},
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
			pos:  position{line: 335, col: 1, offset: 10048},
			expr: &actionExpr{
				pos: position{line: 335, col: 12, offset: 10059},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 335, col: 12, offset: 10059},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 335, col: 12, offset: 10059},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 335, col: 14, offset: 10061},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 335, col: 20, offset: 10067},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 335, col: 30, offset: 10077},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 335, col: 35, offset: 10082},
								expr: &seqExpr{
									pos: position{line: 335, col: 36, offset: 10083},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 335, col: 36, offset: 10083},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 39, offset: 10086},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 51, offset: 10098},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 54, offset: 10101},
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
			pos:  position{line: 355, col: 1, offset: 10682},
			expr: &actionExpr{
				pos: position{line: 355, col: 13, offset: 10694},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 355, col: 13, offset: 10694},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 355, col: 13, offset: 10694},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 355, col: 15, offset: 10696},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 355, col: 21, offset: 10702},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 355, col: 33, offset: 10714},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 355, col: 38, offset: 10719},
								expr: &seqExpr{
									pos: position{line: 355, col: 39, offset: 10720},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 355, col: 39, offset: 10720},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 42, offset: 10723},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 55, offset: 10736},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 58, offset: 10739},
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
			pos:  position{line: 374, col: 1, offset: 11322},
			expr: &choiceExpr{
				pos: position{line: 374, col: 15, offset: 11336},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 374, col: 15, offset: 11336},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 374, col: 15, offset: 11336},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 374, col: 15, offset: 11336},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 374, col: 19, offset: 11340},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 374, col: 21, offset: 11342},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 374, col: 27, offset: 11348},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 374, col: 33, offset: 11354},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 374, col: 35, offset: 11356},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 377, col: 5, offset: 11484},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 379, col: 1, offset: 11491},
			expr: &choiceExpr{
				pos: position{line: 379, col: 12, offset: 11502},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 379, col: 12, offset: 11502},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 379, col: 30, offset: 11520},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 379, col: 49, offset: 11539},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 379, col: 64, offset: 11554},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 381, col: 1, offset: 11567},
			expr: &actionExpr{
				pos: position{line: 381, col: 19, offset: 11585},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 381, col: 21, offset: 11587},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 381, col: 21, offset: 11587},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 381, col: 29, offset: 11595},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 381, col: 36, offset: 11602},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 385, col: 1, offset: 11701},
			expr: &actionExpr{
				pos: position{line: 385, col: 20, offset: 11720},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 385, col: 22, offset: 11722},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 385, col: 22, offset: 11722},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 29, offset: 11729},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 36, offset: 11736},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 42, offset: 11742},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 48, offset: 11748},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 385, col: 56, offset: 11756},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 389, col: 1, offset: 11862},
			expr: &choiceExpr{
				pos: position{line: 389, col: 16, offset: 11877},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 389, col: 16, offset: 11877},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 389, col: 18, offset: 11879},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 389, col: 18, offset: 11879},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 389, col: 25, offset: 11886},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 392, col: 3, offset: 11992},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 392, col: 5, offset: 11994},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 392, col: 5, offset: 11994},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 392, col: 11, offset: 12000},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 392, col: 17, offset: 12006},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 395, col: 3, offset: 12109},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 395, col: 3, offset: 12109},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 399, col: 1, offset: 12213},
			expr: &choiceExpr{
				pos: position{line: 399, col: 15, offset: 12227},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 399, col: 15, offset: 12227},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 399, col: 17, offset: 12229},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 399, col: 17, offset: 12229},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 399, col: 24, offset: 12236},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 402, col: 3, offset: 12342},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 402, col: 5, offset: 12344},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 402, col: 5, offset: 12344},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 402, col: 11, offset: 12350},
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
			pos:  position{line: 406, col: 1, offset: 12452},
			expr: &choiceExpr{
				pos: position{line: 406, col: 9, offset: 12460},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 406, col: 9, offset: 12460},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 406, col: 24, offset: 12475},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 408, col: 1, offset: 12482},
			expr: &choiceExpr{
				pos: position{line: 408, col: 14, offset: 12495},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 14, offset: 12495},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 29, offset: 12510},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 410, col: 1, offset: 12518},
			expr: &choiceExpr{
				pos: position{line: 410, col: 14, offset: 12531},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 410, col: 14, offset: 12531},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 410, col: 29, offset: 12546},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 412, col: 1, offset: 12558},
			expr: &actionExpr{
				pos: position{line: 412, col: 16, offset: 12573},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 412, col: 16, offset: 12573},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 412, col: 16, offset: 12573},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 20, offset: 12577},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 412, col: 22, offset: 12579},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 412, col: 28, offset: 12585},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 33, offset: 12590},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 412, col: 35, offset: 12592},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 412, col: 40, offset: 12597},
								expr: &seqExpr{
									pos: position{line: 412, col: 41, offset: 12598},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 412, col: 41, offset: 12598},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 412, col: 45, offset: 12602},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 412, col: 47, offset: 12604},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 412, col: 52, offset: 12609},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 412, col: 56, offset: 12613},
							expr: &litMatcher{
								pos:        position{line: 412, col: 56, offset: 12613},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 61, offset: 12618},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 412, col: 63, offset: 12620},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 428, col: 1, offset: 13065},
			expr: &actionExpr{
				pos: position{line: 428, col: 19, offset: 13083},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 428, col: 19, offset: 13083},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 428, col: 19, offset: 13083},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 428, col: 24, offset: 13088},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 428, col: 35, offset: 13099},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 428, col: 37, offset: 13101},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 428, col: 42, offset: 13106},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 441, col: 1, offset: 13376},
			expr: &actionExpr{
				pos: position{line: 441, col: 18, offset: 13393},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 441, col: 18, offset: 13393},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 441, col: 18, offset: 13393},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 441, col: 23, offset: 13398},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 34, offset: 13409},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 441, col: 36, offset: 13411},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 40, offset: 13415},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 441, col: 42, offset: 13417},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 441, col: 52, offset: 13427},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 65, offset: 13440},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 441, col: 67, offset: 13442},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 441, col: 71, offset: 13446},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 441, col: 73, offset: 13448},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 441, col: 84, offset: 13459},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 441, col: 89, offset: 13464},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 441, col: 94, offset: 13469},
								expr: &seqExpr{
									pos: position{line: 441, col: 95, offset: 13470},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 441, col: 95, offset: 13470},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 99, offset: 13474},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 101, offset: 13476},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 114, offset: 13489},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 441, col: 116, offset: 13491},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 120, offset: 13495},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 441, col: 122, offset: 13497},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 441, col: 130, offset: 13505},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 461, col: 1, offset: 14089},
			expr: &actionExpr{
				pos: position{line: 461, col: 17, offset: 14105},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 461, col: 17, offset: 14105},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 461, col: 17, offset: 14105},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 461, col: 22, offset: 14110},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 465, col: 1, offset: 14217},
			expr: &actionExpr{
				pos: position{line: 465, col: 16, offset: 14232},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 465, col: 16, offset: 14232},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 465, col: 16, offset: 14232},
							expr: &ruleRefExpr{
								pos:  position{line: 465, col: 17, offset: 14233},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 465, col: 27, offset: 14243},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 465, col: 27, offset: 14243},
									expr: &charClassMatcher{
										pos:        position{line: 465, col: 27, offset: 14243},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 465, col: 34, offset: 14250},
									expr: &charClassMatcher{
										pos:        position{line: 465, col: 34, offset: 14250},
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
			pos:  position{line: 469, col: 1, offset: 14360},
			expr: &actionExpr{
				pos: position{line: 469, col: 14, offset: 14373},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 469, col: 15, offset: 14374},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 469, col: 15, offset: 14374},
							expr: &charClassMatcher{
								pos:        position{line: 469, col: 15, offset: 14374},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 469, col: 22, offset: 14381},
							expr: &charClassMatcher{
								pos:        position{line: 469, col: 22, offset: 14381},
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
			pos:  position{line: 473, col: 1, offset: 14491},
			expr: &choiceExpr{
				pos: position{line: 473, col: 9, offset: 14499},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 473, col: 9, offset: 14499},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 473, col: 9, offset: 14499},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 473, col: 9, offset: 14499},
									expr: &litMatcher{
										pos:        position{line: 473, col: 9, offset: 14499},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 473, col: 14, offset: 14504},
									expr: &charClassMatcher{
										pos:        position{line: 473, col: 14, offset: 14504},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 473, col: 21, offset: 14511},
									expr: &litMatcher{
										pos:        position{line: 473, col: 22, offset: 14512},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 480, col: 3, offset: 14687},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 480, col: 3, offset: 14687},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 480, col: 3, offset: 14687},
									expr: &litMatcher{
										pos:        position{line: 480, col: 3, offset: 14687},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 480, col: 8, offset: 14692},
									expr: &charClassMatcher{
										pos:        position{line: 480, col: 8, offset: 14692},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 480, col: 15, offset: 14699},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 480, col: 19, offset: 14703},
									expr: &charClassMatcher{
										pos:        position{line: 480, col: 19, offset: 14703},
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
						pos:        position{line: 487, col: 3, offset: 14892},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 487, col: 12, offset: 14901},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 487, col: 12, offset: 14901},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 493, col: 3, offset: 15102},
						name: "Unit",
					},
					&actionExpr{
						pos: position{line: 494, col: 3, offset: 15109},
						run: (*parser).callonConst23,
						expr: &seqExpr{
							pos: position{line: 494, col: 3, offset: 15109},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 494, col: 3, offset: 15109},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 494, col: 7, offset: 15113},
									expr: &seqExpr{
										pos: position{line: 494, col: 8, offset: 15114},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 494, col: 8, offset: 15114},
												expr: &ruleRefExpr{
													pos:  position{line: 494, col: 9, offset: 15115},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 494, col: 21, offset: 15127,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 494, col: 25, offset: 15131},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 501, col: 3, offset: 15315},
						run: (*parser).callonConst32,
						expr: &seqExpr{
							pos: position{line: 501, col: 3, offset: 15315},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 501, col: 3, offset: 15315},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 501, col: 7, offset: 15319},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 501, col: 12, offset: 15324},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 501, col: 12, offset: 15324},
												expr: &ruleRefExpr{
													pos:  position{line: 501, col: 13, offset: 15325},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 501, col: 25, offset: 15337,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 501, col: 28, offset: 15340},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 5, offset: 15432},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 20, offset: 15447},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 37, offset: 15464},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 505, col: 1, offset: 15481},
			expr: &actionExpr{
				pos: position{line: 505, col: 8, offset: 15488},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 505, col: 8, offset: 15488},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 509, col: 1, offset: 15551},
			expr: &actionExpr{
				pos: position{line: 509, col: 10, offset: 15560},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 509, col: 11, offset: 15561},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 513, col: 1, offset: 15651},
			expr: &seqExpr{
				pos: position{line: 513, col: 12, offset: 15662},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 513, col: 13, offset: 15663},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 513, col: 13, offset: 15663},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 21, offset: 15671},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 28, offset: 15678},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 37, offset: 15687},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 46, offset: 15696},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 55, offset: 15705},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 64, offset: 15714},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 74, offset: 15724},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 513, col: 86, offset: 15736},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 513, col: 95, offset: 15745},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 513, col: 105, offset: 15755},
						expr: &oneOrMoreExpr{
							pos: position{line: 513, col: 106, offset: 15756},
							expr: &charClassMatcher{
								pos:        position{line: 513, col: 106, offset: 15756},
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
			pos:  position{line: 515, col: 1, offset: 15764},
			expr: &actionExpr{
				pos: position{line: 515, col: 12, offset: 15775},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 515, col: 14, offset: 15777},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 515, col: 14, offset: 15777},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 22, offset: 15785},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 31, offset: 15794},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 42, offset: 15805},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 51, offset: 15814},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 60, offset: 15823},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 515, col: 70, offset: 15833},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 519, col: 1, offset: 15932},
			expr: &charClassMatcher{
				pos:        position{line: 519, col: 15, offset: 15946},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 521, col: 1, offset: 15962},
			expr: &choiceExpr{
				pos: position{line: 521, col: 18, offset: 15979},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 521, col: 18, offset: 15979},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 521, col: 37, offset: 15998},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 523, col: 1, offset: 16013},
			expr: &charClassMatcher{
				pos:        position{line: 523, col: 20, offset: 16032},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 525, col: 1, offset: 16045},
			expr: &charClassMatcher{
				pos:        position{line: 525, col: 16, offset: 16060},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 527, col: 1, offset: 16067},
			expr: &charClassMatcher{
				pos:        position{line: 527, col: 23, offset: 16089},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 529, col: 1, offset: 16096},
			expr: &charClassMatcher{
				pos:        position{line: 529, col: 12, offset: 16107},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 531, col: 1, offset: 16118},
			expr: &choiceExpr{
				pos: position{line: 531, col: 22, offset: 16139},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 531, col: 22, offset: 16139},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 531, col: 32, offset: 16149},
						expr: &charClassMatcher{
							pos:        position{line: 531, col: 32, offset: 16149},
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
			pos:         position{line: 533, col: 1, offset: 16161},
			expr: &zeroOrMoreExpr{
				pos: position{line: 533, col: 18, offset: 16178},
				expr: &charClassMatcher{
					pos:        position{line: 533, col: 18, offset: 16178},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 535, col: 1, offset: 16190},
			expr: &notExpr{
				pos: position{line: 535, col: 7, offset: 16196},
				expr: &anyMatcher{
					line: 535, col: 8, offset: 16197,
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

	return AliasType{Name: name.(BasicAst).StringValue, Params: parameters, Types: fields}, nil
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

	return RecordType{Name: name.(BasicAst).StringValue, Fields: fields}, nil
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

	return VariantType{Name: name.(BasicAst).StringValue, Params: parameters, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn54() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn54(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{Name: name.(BasicAst).StringValue, Type: t.(Ast)}, nil
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

	record := RecordType{Name: name.(BasicAst).StringValue, Fields: fields}
	return VariantConstructor{Name: name.(BasicAst).StringValue, Fields: []Ast{record}}, nil
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

	return VariantConstructor{Name: name.(BasicAst).StringValue, Fields: params}, nil
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
	return Func{Name: i.(BasicAst).StringValue, Arguments: args, Subvalues: subvalues}, nil
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

		return Container{Type: "CompoundExpr", Subvalues: subvalues}, nil
	} else {
		return Container{Type: "CompoundExpr", Subvalues: []Ast{op.(Ast)}}, nil
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
		return Container{Type: "BinOpBool", Subvalues: subvalues}, nil
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
		return Container{Type: "BinOpEquality", Subvalues: subvalues}, nil
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
		return Container{Type: "BinOpLow", Subvalues: subvalues}, nil
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
		return Container{Type: "BinOpHigh", Subvalues: subvalues}, nil
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
	return Container{Type: "BinOpParens", Subvalues: []Ast{first.(Ast)}}, nil
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
	return BasicAst{Type: "TypeParam", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonTypeParameter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeParameter1()
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
	return BasicAst{Type: "Identifier", StringValue: "_", ValueType: STRING}, nil
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
