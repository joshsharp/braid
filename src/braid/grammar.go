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
					&ruleRefExpr{
						pos:  position{line: 28, col: 42, offset: 736},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 746},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 758},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 758},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 768},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 779},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 792},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 802},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 813},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 813},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 813},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 815},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 820},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 821},
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
			pos:  position{line: 36, col: 1, offset: 849},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 859},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 859},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 36, col: 11, offset: 859},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 15, offset: 863},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 23, offset: 871},
								expr: &seqExpr{
									pos: position{line: 36, col: 24, offset: 872},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 24, offset: 872},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 25, offset: 873},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 37, offset: 885,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 41, offset: 889},
							expr: &litMatcher{
								pos:        position{line: 36, col: 42, offset: 890},
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
			pos:  position{line: 41, col: 1, offset: 1040},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 1051},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 1051},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 1051},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 1051},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 1053},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 1060},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 1063},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 1068},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 42, offset: 1081},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 49, offset: 1088},
										expr: &seqExpr{
											pos: position{line: 41, col: 50, offset: 1089},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 50, offset: 1089},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 53, offset: 1092},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 69, offset: 1108},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 71, offset: 1110},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 75, offset: 1114},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 81, offset: 1120},
										expr: &seqExpr{
											pos: position{line: 41, col: 82, offset: 1121},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 82, offset: 1121},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 85, offset: 1124},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 95, offset: 1134},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 46, col: 1, offset: 1224},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 46, col: 1, offset: 1224},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 46, col: 1, offset: 1224},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 46, col: 3, offset: 1226},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 10, offset: 1233},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 46, col: 13, offset: 1236},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 46, col: 18, offset: 1241},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 46, col: 31, offset: 1254},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 46, col: 38, offset: 1261},
										expr: &seqExpr{
											pos: position{line: 46, col: 39, offset: 1262},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 46, col: 39, offset: 1262},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 46, col: 42, offset: 1265},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 58, offset: 1281},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 46, col: 60, offset: 1283},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 64, offset: 1287},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 47, col: 5, offset: 1293},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 9, offset: 1297},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 47, col: 11, offset: 1299},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 47, col: 17, offset: 1305},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 33, offset: 1321},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 47, col: 35, offset: 1323},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 47, col: 40, offset: 1328},
										expr: &seqExpr{
											pos: position{line: 47, col: 41, offset: 1329},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 47, col: 41, offset: 1329},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 47, col: 45, offset: 1333},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 47, col: 47, offset: 1335},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 47, col: 63, offset: 1351},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 47, col: 67, offset: 1355},
									expr: &litMatcher{
										pos:        position{line: 47, col: 67, offset: 1355},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 72, offset: 1360},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 47, col: 74, offset: 1362},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 78, offset: 1366},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 65, col: 1, offset: 1825},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 65, col: 1, offset: 1825},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 65, col: 1, offset: 1825},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 3, offset: 1827},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 10, offset: 1834},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 65, col: 13, offset: 1837},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 65, col: 18, offset: 1842},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 65, col: 31, offset: 1855},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 65, col: 38, offset: 1862},
										expr: &seqExpr{
											pos: position{line: 65, col: 39, offset: 1863},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 65, col: 39, offset: 1863},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 65, col: 42, offset: 1866},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 58, offset: 1882},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 60, offset: 1884},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 64, offset: 1888},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 65, col: 66, offset: 1890},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 65, col: 71, offset: 1895},
										expr: &ruleRefExpr{
											pos:  position{line: 65, col: 72, offset: 1896},
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
			name: "VariantConstructor",
			pos:  position{line: 70, col: 1, offset: 2005},
			expr: &actionExpr{
				pos: position{line: 70, col: 22, offset: 2026},
				run: (*parser).callonVariantConstructor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 22, offset: 2026},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 70, col: 22, offset: 2026},
							val:        "|",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 26, offset: 2030},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 70, col: 28, offset: 2032},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 33, offset: 2037},
								name: "ModuleName",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 44, offset: 2048},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 70, col: 49, offset: 2053},
								expr: &seqExpr{
									pos: position{line: 70, col: 50, offset: 2054},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 70, col: 50, offset: 2054},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 70, col: 53, offset: 2057},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 63, offset: 2067},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 74, col: 1, offset: 2144},
			expr: &actionExpr{
				pos: position{line: 74, col: 19, offset: 2162},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 74, col: 19, offset: 2162},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 74, col: 19, offset: 2162},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 74, col: 24, offset: 2167},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 37, offset: 2180},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 74, col: 39, offset: 2182},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 43, offset: 2186},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 74, col: 45, offset: 2188},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 74, col: 48, offset: 2191},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 78, col: 1, offset: 2283},
			expr: &choiceExpr{
				pos: position{line: 78, col: 11, offset: 2293},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 78, col: 11, offset: 2293},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 22, offset: 2304},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 80, col: 1, offset: 2319},
			expr: &choiceExpr{
				pos: position{line: 80, col: 14, offset: 2332},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 80, col: 14, offset: 2332},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 80, col: 14, offset: 2332},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 80, col: 14, offset: 2332},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 80, col: 16, offset: 2334},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 80, col: 22, offset: 2340},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 80, col: 25, offset: 2343},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 80, col: 27, offset: 2345},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 80, col: 38, offset: 2356},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 80, col: 43, offset: 2361},
										expr: &seqExpr{
											pos: position{line: 80, col: 44, offset: 2362},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 80, col: 44, offset: 2362},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 80, col: 48, offset: 2366},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 80, col: 50, offset: 2368},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 80, col: 63, offset: 2381},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 80, col: 65, offset: 2383},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 80, col: 69, offset: 2387},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 80, col: 71, offset: 2389},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 80, col: 76, offset: 2394},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 80, col: 81, offset: 2399},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 95, col: 1, offset: 2838},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 95, col: 1, offset: 2838},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 95, col: 1, offset: 2838},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 95, col: 3, offset: 2840},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 95, col: 9, offset: 2846},
									name: "__",
								},
								&notExpr{
									pos: position{line: 95, col: 12, offset: 2849},
									expr: &ruleRefExpr{
										pos:  position{line: 95, col: 13, offset: 2850},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 99, col: 1, offset: 2958},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 99, col: 1, offset: 2958},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 99, col: 1, offset: 2958},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 99, col: 3, offset: 2960},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 99, col: 9, offset: 2966},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 99, col: 12, offset: 2969},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 99, col: 14, offset: 2971},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 99, col: 25, offset: 2982},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 99, col: 30, offset: 2987},
										expr: &seqExpr{
											pos: position{line: 99, col: 31, offset: 2988},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 99, col: 31, offset: 2988},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 99, col: 35, offset: 2992},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 99, col: 37, offset: 2994},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 99, col: 50, offset: 3007},
									name: "_",
								},
								&notExpr{
									pos: position{line: 99, col: 52, offset: 3009},
									expr: &litMatcher{
										pos:        position{line: 99, col: 53, offset: 3010},
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
			pos:  position{line: 103, col: 1, offset: 3104},
			expr: &actionExpr{
				pos: position{line: 103, col: 12, offset: 3115},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 103, col: 12, offset: 3115},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 103, col: 12, offset: 3115},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 103, col: 14, offset: 3117},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 20, offset: 3123},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 103, col: 23, offset: 3126},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 103, col: 25, offset: 3128},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 38, offset: 3141},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 103, col: 40, offset: 3143},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 44, offset: 3147},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 103, col: 46, offset: 3149},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 53, offset: 3156},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 103, col: 56, offset: 3159},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 103, col: 60, offset: 3163},
								expr: &seqExpr{
									pos: position{line: 103, col: 61, offset: 3164},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 103, col: 61, offset: 3164},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 103, col: 74, offset: 3177},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 79, offset: 3182},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 103, col: 81, offset: 3184},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 85, offset: 3188},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 103, col: 88, offset: 3191},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 103, col: 99, offset: 3202},
								expr: &ruleRefExpr{
									pos:  position{line: 103, col: 100, offset: 3203},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 112, offset: 3215},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 103, col: 114, offset: 3217},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 103, col: 118, offset: 3221},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 126, col: 1, offset: 3896},
			expr: &actionExpr{
				pos: position{line: 126, col: 8, offset: 3903},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 126, col: 8, offset: 3903},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 126, col: 12, offset: 3907},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 126, col: 12, offset: 3907},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 126, col: 21, offset: 3916},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 126, col: 28, offset: 3923},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 131, col: 1, offset: 4018},
			expr: &choiceExpr{
				pos: position{line: 131, col: 10, offset: 4027},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 131, col: 10, offset: 4027},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 131, col: 10, offset: 4027},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 131, col: 10, offset: 4027},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 131, col: 15, offset: 4032},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 131, col: 18, offset: 4035},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 131, col: 23, offset: 4040},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 131, col: 33, offset: 4050},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 131, col: 35, offset: 4052},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 131, col: 39, offset: 4056},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 131, col: 41, offset: 4058},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 131, col: 47, offset: 4064},
										expr: &ruleRefExpr{
											pos:  position{line: 131, col: 48, offset: 4065},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 131, col: 60, offset: 4077},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 131, col: 62, offset: 4079},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 131, col: 66, offset: 4083},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 131, col: 68, offset: 4085},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 131, col: 75, offset: 4092},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 131, col: 77, offset: 4094},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 131, col: 85, offset: 4102},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 143, col: 1, offset: 4432},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 143, col: 1, offset: 4432},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 143, col: 1, offset: 4432},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 6, offset: 4437},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 143, col: 9, offset: 4440},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 143, col: 14, offset: 4445},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 24, offset: 4455},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 143, col: 26, offset: 4457},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 30, offset: 4461},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 143, col: 32, offset: 4463},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 143, col: 38, offset: 4469},
										expr: &ruleRefExpr{
											pos:  position{line: 143, col: 39, offset: 4470},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 51, offset: 4482},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 143, col: 54, offset: 4485},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 58, offset: 4489},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 143, col: 60, offset: 4491},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 67, offset: 4498},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 143, col: 69, offset: 4500},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 73, offset: 4504},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 143, col: 75, offset: 4506},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 143, col: 81, offset: 4512},
										expr: &ruleRefExpr{
											pos:  position{line: 143, col: 82, offset: 4513},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 143, col: 94, offset: 4525},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 143, col: 97, offset: 4528},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 162, col: 1, offset: 5031},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 162, col: 1, offset: 5031},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 162, col: 1, offset: 5031},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 162, col: 6, offset: 5036},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 162, col: 9, offset: 5039},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 162, col: 14, offset: 5044},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 162, col: 24, offset: 5054},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 162, col: 26, offset: 5056},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 162, col: 30, offset: 5060},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 162, col: 32, offset: 5062},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 162, col: 38, offset: 5068},
										expr: &ruleRefExpr{
											pos:  position{line: 162, col: 39, offset: 5069},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 162, col: 51, offset: 5081},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 162, col: 54, offset: 5084},
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
			pos:  position{line: 175, col: 1, offset: 5383},
			expr: &choiceExpr{
				pos: position{line: 175, col: 8, offset: 5390},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 175, col: 8, offset: 5390},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 175, col: 8, offset: 5390},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 175, col: 8, offset: 5390},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 15, offset: 5397},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 175, col: 26, offset: 5408},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 175, col: 30, offset: 5412},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 33, offset: 5415},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 175, col: 46, offset: 5428},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 175, col: 56, offset: 5438},
										expr: &seqExpr{
											pos: position{line: 175, col: 57, offset: 5439},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 175, col: 57, offset: 5439},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 175, col: 60, offset: 5442},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 74, offset: 5456},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 193, col: 1, offset: 5954},
						run: (*parser).callonCall15,
						expr: &seqExpr{
							pos: position{line: 193, col: 1, offset: 5954},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 193, col: 1, offset: 5954},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 3, offset: 5956},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 193, col: 6, offset: 5959},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 193, col: 19, offset: 5972},
									label: "arguments",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 29, offset: 5982},
										expr: &seqExpr{
											pos: position{line: 193, col: 30, offset: 5983},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 193, col: 30, offset: 5983},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 193, col: 33, offset: 5986},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 47, offset: 6000},
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
			pos:  position{line: 211, col: 1, offset: 6510},
			expr: &actionExpr{
				pos: position{line: 211, col: 16, offset: 6525},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 211, col: 16, offset: 6525},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 211, col: 16, offset: 6525},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 211, col: 18, offset: 6527},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 211, col: 21, offset: 6530},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 211, col: 27, offset: 6536},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 211, col: 32, offset: 6541},
								expr: &seqExpr{
									pos: position{line: 211, col: 33, offset: 6542},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 211, col: 33, offset: 6542},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 211, col: 36, offset: 6545},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 211, col: 45, offset: 6554},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 211, col: 48, offset: 6557},
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
			pos:  position{line: 231, col: 1, offset: 7221},
			expr: &choiceExpr{
				pos: position{line: 231, col: 9, offset: 7229},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 231, col: 9, offset: 7229},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 231, col: 21, offset: 7241},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 231, col: 37, offset: 7257},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 231, col: 48, offset: 7268},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 231, col: 60, offset: 7280},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 233, col: 1, offset: 7293},
			expr: &actionExpr{
				pos: position{line: 233, col: 13, offset: 7305},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 233, col: 13, offset: 7305},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 233, col: 13, offset: 7305},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 233, col: 15, offset: 7307},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 233, col: 21, offset: 7313},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 233, col: 35, offset: 7327},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 233, col: 40, offset: 7332},
								expr: &seqExpr{
									pos: position{line: 233, col: 41, offset: 7333},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 233, col: 41, offset: 7333},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 233, col: 44, offset: 7336},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 233, col: 60, offset: 7352},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 233, col: 63, offset: 7355},
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
			pos:  position{line: 252, col: 1, offset: 7961},
			expr: &actionExpr{
				pos: position{line: 252, col: 17, offset: 7977},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 252, col: 17, offset: 7977},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 252, col: 17, offset: 7977},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 252, col: 19, offset: 7979},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 25, offset: 7985},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 252, col: 34, offset: 7994},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 252, col: 39, offset: 7999},
								expr: &seqExpr{
									pos: position{line: 252, col: 40, offset: 8000},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 252, col: 40, offset: 8000},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 252, col: 43, offset: 8003},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 252, col: 60, offset: 8020},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 252, col: 63, offset: 8023},
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
			pos:  position{line: 272, col: 1, offset: 8627},
			expr: &actionExpr{
				pos: position{line: 272, col: 12, offset: 8638},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 272, col: 12, offset: 8638},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 272, col: 12, offset: 8638},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 272, col: 14, offset: 8640},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 272, col: 20, offset: 8646},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 272, col: 30, offset: 8656},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 272, col: 35, offset: 8661},
								expr: &seqExpr{
									pos: position{line: 272, col: 36, offset: 8662},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 272, col: 36, offset: 8662},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 272, col: 39, offset: 8665},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 272, col: 51, offset: 8677},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 272, col: 54, offset: 8680},
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
			pos:  position{line: 292, col: 1, offset: 9281},
			expr: &actionExpr{
				pos: position{line: 292, col: 13, offset: 9293},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 292, col: 13, offset: 9293},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 292, col: 13, offset: 9293},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 292, col: 15, offset: 9295},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 292, col: 21, offset: 9301},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 292, col: 33, offset: 9313},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 292, col: 38, offset: 9318},
								expr: &seqExpr{
									pos: position{line: 292, col: 39, offset: 9319},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 292, col: 39, offset: 9319},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 292, col: 42, offset: 9322},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 292, col: 55, offset: 9335},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 292, col: 58, offset: 9338},
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
			pos:  position{line: 311, col: 1, offset: 9942},
			expr: &choiceExpr{
				pos: position{line: 311, col: 15, offset: 9956},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 311, col: 15, offset: 9956},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 311, col: 15, offset: 9956},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 311, col: 15, offset: 9956},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 311, col: 19, offset: 9960},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 311, col: 21, offset: 9962},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 311, col: 27, offset: 9968},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 311, col: 33, offset: 9974},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 311, col: 35, offset: 9976},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 314, col: 5, offset: 10125},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 316, col: 1, offset: 10132},
			expr: &choiceExpr{
				pos: position{line: 316, col: 12, offset: 10143},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 316, col: 12, offset: 10143},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 316, col: 30, offset: 10161},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 316, col: 49, offset: 10180},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 316, col: 64, offset: 10195},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 318, col: 1, offset: 10208},
			expr: &actionExpr{
				pos: position{line: 318, col: 19, offset: 10226},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 318, col: 21, offset: 10228},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 318, col: 21, offset: 10228},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 318, col: 29, offset: 10236},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 318, col: 36, offset: 10243},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 322, col: 1, offset: 10342},
			expr: &actionExpr{
				pos: position{line: 322, col: 20, offset: 10361},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 322, col: 22, offset: 10363},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 322, col: 22, offset: 10363},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 322, col: 29, offset: 10370},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 322, col: 36, offset: 10377},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 322, col: 42, offset: 10383},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 322, col: 48, offset: 10389},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 322, col: 56, offset: 10397},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 326, col: 1, offset: 10503},
			expr: &choiceExpr{
				pos: position{line: 326, col: 16, offset: 10518},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 326, col: 16, offset: 10518},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 326, col: 18, offset: 10520},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 326, col: 18, offset: 10520},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 326, col: 25, offset: 10527},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 329, col: 3, offset: 10633},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 329, col: 5, offset: 10635},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 329, col: 5, offset: 10635},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 329, col: 11, offset: 10641},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 329, col: 17, offset: 10647},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 332, col: 3, offset: 10750},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 332, col: 3, offset: 10750},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 336, col: 1, offset: 10854},
			expr: &choiceExpr{
				pos: position{line: 336, col: 15, offset: 10868},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 336, col: 15, offset: 10868},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 336, col: 17, offset: 10870},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 336, col: 17, offset: 10870},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 336, col: 24, offset: 10877},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 339, col: 3, offset: 10983},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 339, col: 5, offset: 10985},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 339, col: 5, offset: 10985},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 339, col: 11, offset: 10991},
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
			pos:  position{line: 343, col: 1, offset: 11093},
			expr: &choiceExpr{
				pos: position{line: 343, col: 9, offset: 11101},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 343, col: 9, offset: 11101},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 343, col: 24, offset: 11116},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 345, col: 1, offset: 11123},
			expr: &choiceExpr{
				pos: position{line: 345, col: 14, offset: 11136},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 345, col: 14, offset: 11136},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 345, col: 29, offset: 11151},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 347, col: 1, offset: 11159},
			expr: &choiceExpr{
				pos: position{line: 347, col: 14, offset: 11172},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 347, col: 14, offset: 11172},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 347, col: 29, offset: 11187},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 349, col: 1, offset: 11199},
			expr: &actionExpr{
				pos: position{line: 349, col: 16, offset: 11214},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 349, col: 16, offset: 11214},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 349, col: 16, offset: 11214},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 349, col: 20, offset: 11218},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 349, col: 22, offset: 11220},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 349, col: 28, offset: 11226},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 349, col: 33, offset: 11231},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 349, col: 35, offset: 11233},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 349, col: 40, offset: 11238},
								expr: &seqExpr{
									pos: position{line: 349, col: 41, offset: 11239},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 349, col: 41, offset: 11239},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 349, col: 45, offset: 11243},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 349, col: 47, offset: 11245},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 349, col: 52, offset: 11250},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 349, col: 56, offset: 11254},
							expr: &litMatcher{
								pos:        position{line: 349, col: 56, offset: 11254},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 349, col: 61, offset: 11259},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 349, col: 63, offset: 11261},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 365, col: 1, offset: 11742},
			expr: &actionExpr{
				pos: position{line: 365, col: 17, offset: 11758},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 365, col: 17, offset: 11758},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 365, col: 17, offset: 11758},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 365, col: 22, offset: 11763},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 369, col: 1, offset: 11871},
			expr: &actionExpr{
				pos: position{line: 369, col: 16, offset: 11886},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 369, col: 16, offset: 11886},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 369, col: 16, offset: 11886},
							expr: &ruleRefExpr{
								pos:  position{line: 369, col: 17, offset: 11887},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 369, col: 27, offset: 11897},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 369, col: 27, offset: 11897},
									expr: &charClassMatcher{
										pos:        position{line: 369, col: 27, offset: 11897},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 369, col: 34, offset: 11904},
									expr: &charClassMatcher{
										pos:        position{line: 369, col: 34, offset: 11904},
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
			pos:  position{line: 373, col: 1, offset: 12015},
			expr: &actionExpr{
				pos: position{line: 373, col: 14, offset: 12028},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 373, col: 15, offset: 12029},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 373, col: 15, offset: 12029},
							expr: &charClassMatcher{
								pos:        position{line: 373, col: 15, offset: 12029},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 373, col: 22, offset: 12036},
							expr: &charClassMatcher{
								pos:        position{line: 373, col: 22, offset: 12036},
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
			pos:  position{line: 377, col: 1, offset: 12147},
			expr: &choiceExpr{
				pos: position{line: 377, col: 9, offset: 12155},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 377, col: 9, offset: 12155},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 377, col: 9, offset: 12155},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 377, col: 9, offset: 12155},
									expr: &litMatcher{
										pos:        position{line: 377, col: 9, offset: 12155},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 377, col: 14, offset: 12160},
									expr: &charClassMatcher{
										pos:        position{line: 377, col: 14, offset: 12160},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 377, col: 21, offset: 12167},
									expr: &litMatcher{
										pos:        position{line: 377, col: 22, offset: 12168},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 3, offset: 12344},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 384, col: 3, offset: 12344},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 384, col: 3, offset: 12344},
									expr: &litMatcher{
										pos:        position{line: 384, col: 3, offset: 12344},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 384, col: 8, offset: 12349},
									expr: &charClassMatcher{
										pos:        position{line: 384, col: 8, offset: 12349},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 384, col: 15, offset: 12356},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 384, col: 19, offset: 12360},
									expr: &charClassMatcher{
										pos:        position{line: 384, col: 19, offset: 12360},
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
						pos:        position{line: 391, col: 3, offset: 12550},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 391, col: 12, offset: 12559},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 391, col: 12, offset: 12559},
							val:        "false",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 397, col: 3, offset: 12760},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 397, col: 3, offset: 12760},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 400, col: 3, offset: 12823},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 400, col: 3, offset: 12823},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 400, col: 3, offset: 12823},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 400, col: 7, offset: 12827},
									expr: &seqExpr{
										pos: position{line: 400, col: 8, offset: 12828},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 400, col: 8, offset: 12828},
												expr: &ruleRefExpr{
													pos:  position{line: 400, col: 9, offset: 12829},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 400, col: 21, offset: 12841,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 400, col: 25, offset: 12845},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 407, col: 3, offset: 13029},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 407, col: 3, offset: 13029},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 407, col: 3, offset: 13029},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 407, col: 7, offset: 13033},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 407, col: 12, offset: 13038},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 407, col: 12, offset: 13038},
												expr: &ruleRefExpr{
													pos:  position{line: 407, col: 13, offset: 13039},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 407, col: 25, offset: 13051,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 407, col: 28, offset: 13054},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 5, offset: 13146},
						name: "ArrayLiteral",
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 411, col: 1, offset: 13160},
			expr: &actionExpr{
				pos: position{line: 411, col: 10, offset: 13169},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 411, col: 11, offset: 13170},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 415, col: 1, offset: 13271},
			expr: &seqExpr{
				pos: position{line: 415, col: 12, offset: 13282},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 415, col: 13, offset: 13283},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 415, col: 13, offset: 13283},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 21, offset: 13291},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 28, offset: 13298},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 37, offset: 13307},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 46, offset: 13316},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 55, offset: 13325},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 64, offset: 13334},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 74, offset: 13344},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 415, col: 86, offset: 13356},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 415, col: 95, offset: 13365},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 415, col: 105, offset: 13375},
						expr: &oneOrMoreExpr{
							pos: position{line: 415, col: 106, offset: 13376},
							expr: &charClassMatcher{
								pos:        position{line: 415, col: 106, offset: 13376},
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
			pos:  position{line: 417, col: 1, offset: 13384},
			expr: &actionExpr{
				pos: position{line: 417, col: 12, offset: 13395},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 417, col: 14, offset: 13397},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 417, col: 14, offset: 13397},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 417, col: 22, offset: 13405},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 417, col: 31, offset: 13414},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 417, col: 42, offset: 13425},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 417, col: 51, offset: 13434},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 417, col: 60, offset: 13443},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 417, col: 70, offset: 13453},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 421, col: 1, offset: 13552},
			expr: &charClassMatcher{
				pos:        position{line: 421, col: 15, offset: 13566},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 423, col: 1, offset: 13582},
			expr: &choiceExpr{
				pos: position{line: 423, col: 18, offset: 13599},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 423, col: 18, offset: 13599},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 423, col: 37, offset: 13618},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 425, col: 1, offset: 13633},
			expr: &charClassMatcher{
				pos:        position{line: 425, col: 20, offset: 13652},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 427, col: 1, offset: 13665},
			expr: &charClassMatcher{
				pos:        position{line: 427, col: 16, offset: 13680},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 429, col: 1, offset: 13687},
			expr: &charClassMatcher{
				pos:        position{line: 429, col: 23, offset: 13709},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 431, col: 1, offset: 13716},
			expr: &charClassMatcher{
				pos:        position{line: 431, col: 12, offset: 13727},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 433, col: 1, offset: 13738},
			expr: &oneOrMoreExpr{
				pos: position{line: 433, col: 22, offset: 13759},
				expr: &charClassMatcher{
					pos:        position{line: 433, col: 22, offset: 13759},
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
			pos:         position{line: 435, col: 1, offset: 13771},
			expr: &zeroOrMoreExpr{
				pos: position{line: 435, col: 18, offset: 13788},
				expr: &charClassMatcher{
					pos:        position{line: 435, col: 18, offset: 13788},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 437, col: 1, offset: 13800},
			expr: &notExpr{
				pos: position{line: 437, col: 7, offset: 13806},
				expr: &anyMatcher{
					line: 437, col: 8, offset: 13807,
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

func (c *current) onTypeDefn2(name, params, types interface{}) (interface{}, error) {
	// alias type
	return AliasType{Name: name.(BasicAst).StringValue}, nil
}

func (p *parser) callonTypeDefn2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn2(stack["name"], stack["params"], stack["types"])
}

func (c *current) onTypeDefn22(name, params, first, rest interface{}) (interface{}, error) {
	// record type
	fields := []Ast{first.(Ast)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
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
	return VariantType{Name: name.(BasicAst).StringValue}, nil
}

func (p *parser) callonTypeDefn54() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn54(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onVariantConstructor1(name, rest interface{}) (interface{}, error) {
	return VariantConstructor{Name: name.(BasicAst).StringValue}, nil
}

func (p *parser) callonVariantConstructor1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor1(stack["name"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{Name: name.(BasicAst).StringValue, Type: t.(Ast)}, nil
}

func (p *parser) callonRecordFieldDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordFieldDefn1(stack["name"], stack["t"])
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
