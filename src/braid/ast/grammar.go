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
						&litMatcher{
							pos:        position{line: 13, col: 12, offset: 153},
							val:        "module",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 21, offset: 162},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 25, offset: 166},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 30, offset: 171},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 41, offset: 182},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 44, offset: 185},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 49, offset: 190},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 67, offset: 208},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 69, offset: 210},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 74, offset: 215},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 75, offset: 216},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 95, offset: 236},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 97, offset: 238},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 785},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 805},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 805},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 32, offset: 816},
						name: "TypeDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 43, offset: 827},
						name: "ExternFunc",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 56, offset: 840},
						name: "ExternType",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 852},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 864},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 864},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 875},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 888},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 898},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 909},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 909},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 909},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 911},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 32, col: 19, offset: 916},
							name: "N",
						},
					},
				},
			},
		},
		{
			name: "ExternFunc",
			pos:  position{line: 46, col: 1, offset: 1247},
			expr: &actionExpr{
				pos: position{line: 46, col: 14, offset: 1260},
				run: (*parser).callonExternFunc1,
				expr: &seqExpr{
					pos: position{line: 46, col: 14, offset: 1260},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 14, offset: 1260},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 16, offset: 1262},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 25, offset: 1271},
							name: "__N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 29, offset: 1275},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 36, offset: 1282},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 40, offset: 1286},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 45, offset: 1291},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 56, offset: 1302},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 59, offset: 1305},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 63, offset: 1309},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 66, offset: 1312},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 77, offset: 1323},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 91, offset: 1337},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 94, offset: 1340},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 99, offset: 1345},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 5, offset: 1358},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 47, col: 8, offset: 1361},
							val:        "->",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 13, offset: 1366},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 16, offset: 1369},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 20, offset: 1373},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "ExternType",
			pos:  position{line: 53, col: 1, offset: 1583},
			expr: &choiceExpr{
				pos: position{line: 53, col: 14, offset: 1596},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 53, col: 14, offset: 1596},
						run: (*parser).callonExternType2,
						expr: &seqExpr{
							pos: position{line: 53, col: 14, offset: 1596},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 53, col: 14, offset: 1596},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 16, offset: 1598},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 25, offset: 1607},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 53, col: 29, offset: 1611},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 36, offset: 1618},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 39, offset: 1621},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 53, col: 44, offset: 1626},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 55, offset: 1637},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 57, offset: 1639},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 61, offset: 1643},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 5, offset: 1649},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 16, offset: 1660},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 30, offset: 1674},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 32, offset: 1676},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 36, offset: 1680},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 38, offset: 1682},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 44, offset: 1688},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 60, offset: 1704},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 62, offset: 1706},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 54, col: 67, offset: 1711},
										expr: &seqExpr{
											pos: position{line: 54, col: 68, offset: 1712},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 54, col: 68, offset: 1712},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 72, offset: 1716},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 74, offset: 1718},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 90, offset: 1734},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 54, col: 94, offset: 1738},
									expr: &litMatcher{
										pos:        position{line: 54, col: 94, offset: 1738},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 99, offset: 1743},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 101, offset: 1745},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 105, offset: 1749},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 73, col: 1, offset: 2298},
						run: (*parser).callonExternType34,
						expr: &seqExpr{
							pos: position{line: 73, col: 1, offset: 2298},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 73, col: 1, offset: 2298},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 3, offset: 2300},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 12, offset: 2309},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 73, col: 16, offset: 2313},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 23, offset: 2320},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 26, offset: 2323},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 73, col: 31, offset: 2328},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 42, offset: 2339},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 44, offset: 2341},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 48, offset: 2345},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 5, offset: 2351},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 74, col: 16, offset: 2362},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 30, offset: 2376},
									name: "N",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 82, col: 1, offset: 2570},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2581},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 82, col: 12, offset: 2581},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 82, col: 12, offset: 2581},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 82, col: 12, offset: 2581},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 14, offset: 2583},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 21, offset: 2590},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 24, offset: 2593},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 29, offset: 2598},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 40, offset: 2609},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 42, offset: 2611},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 49, offset: 2618},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 60, offset: 2629},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 62, offset: 2631},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 82, col: 66, offset: 2635},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 82, col: 72, offset: 2641},
										expr: &seqExpr{
											pos: position{line: 82, col: 73, offset: 2642},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 73, offset: 2642},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 76, offset: 2645},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 93, offset: 2662},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 100, col: 1, offset: 3149},
						run: (*parser).callonTypeDefn20,
						expr: &seqExpr{
							pos: position{line: 100, col: 1, offset: 3149},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 100, col: 1, offset: 3149},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 3, offset: 3151},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 10, offset: 3158},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 13, offset: 3161},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 18, offset: 3166},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 29, offset: 3177},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 31, offset: 3179},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 38, offset: 3186},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 49, offset: 3197},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 51, offset: 3199},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 55, offset: 3203},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 5, offset: 3209},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 9, offset: 3213},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 11, offset: 3215},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 17, offset: 3221},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 33, offset: 3237},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 35, offset: 3239},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 40, offset: 3244},
										expr: &seqExpr{
											pos: position{line: 101, col: 41, offset: 3245},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 101, col: 41, offset: 3245},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 45, offset: 3249},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 47, offset: 3251},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 63, offset: 3267},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 101, col: 67, offset: 3271},
									expr: &litMatcher{
										pos:        position{line: 101, col: 67, offset: 3271},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 72, offset: 3276},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 74, offset: 3278},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 78, offset: 3282},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 119, col: 1, offset: 3770},
						run: (*parser).callonTypeDefn50,
						expr: &seqExpr{
							pos: position{line: 119, col: 1, offset: 3770},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 119, col: 1, offset: 3770},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 3, offset: 3772},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 10, offset: 3779},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 13, offset: 3782},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 119, col: 18, offset: 3787},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 29, offset: 3798},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 31, offset: 3800},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 35, offset: 3804},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 5, offset: 3810},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 9, offset: 3814},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 11, offset: 3816},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 120, col: 17, offset: 3822},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 33, offset: 3838},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 35, offset: 3840},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 120, col: 40, offset: 3845},
										expr: &seqExpr{
											pos: position{line: 120, col: 41, offset: 3846},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 120, col: 41, offset: 3846},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 45, offset: 3850},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 47, offset: 3852},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 63, offset: 3868},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 120, col: 67, offset: 3872},
									expr: &litMatcher{
										pos:        position{line: 120, col: 67, offset: 3872},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 72, offset: 3877},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 74, offset: 3879},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 78, offset: 3883},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 138, col: 1, offset: 4387},
						run: (*parser).callonTypeDefn77,
						expr: &seqExpr{
							pos: position{line: 138, col: 1, offset: 4387},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 138, col: 1, offset: 4387},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 3, offset: 4389},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 10, offset: 4396},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 13, offset: 4399},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 18, offset: 4404},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 29, offset: 4415},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 31, offset: 4417},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 38, offset: 4424},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 49, offset: 4435},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 51, offset: 4437},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 55, offset: 4441},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 57, offset: 4443},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 138, col: 62, offset: 4448},
										expr: &ruleRefExpr{
											pos:  position{line: 138, col: 63, offset: 4449},
											name: "VariantConstructor",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 84, offset: 4470},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 152, col: 1, offset: 4863},
						run: (*parser).callonTypeDefn94,
						expr: &seqExpr{
							pos: position{line: 152, col: 1, offset: 4863},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 152, col: 1, offset: 4863},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 152, col: 3, offset: 4865},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 10, offset: 4872},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 152, col: 13, offset: 4875},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 152, col: 18, offset: 4880},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 29, offset: 4891},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 152, col: 31, offset: 4893},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 35, offset: 4897},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 152, col: 37, offset: 4899},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 152, col: 42, offset: 4904},
										expr: &ruleRefExpr{
											pos:  position{line: 152, col: 43, offset: 4905},
											name: "VariantConstructor",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 64, offset: 4926},
									name: "N",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "TypeParams",
			pos:  position{line: 167, col: 1, offset: 5314},
			expr: &actionExpr{
				pos: position{line: 167, col: 14, offset: 5327},
				run: (*parser).callonTypeParams1,
				expr: &seqExpr{
					pos: position{line: 167, col: 14, offset: 5327},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 167, col: 14, offset: 5327},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 167, col: 18, offset: 5331},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 24, offset: 5337},
								name: "TypeParameter",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 167, col: 38, offset: 5351},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 167, col: 40, offset: 5353},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 167, col: 45, offset: 5358},
								expr: &seqExpr{
									pos: position{line: 167, col: 46, offset: 5359},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 167, col: 46, offset: 5359},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 50, offset: 5363},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 52, offset: 5365},
											name: "TypeParameter",
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 66, offset: 5379},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 167, col: 70, offset: 5383},
							expr: &litMatcher{
								pos:        position{line: 167, col: 70, offset: 5383},
								val:        ",",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 167, col: 75, offset: 5388},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 183, col: 1, offset: 5799},
			expr: &actionExpr{
				pos: position{line: 183, col: 19, offset: 5817},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 183, col: 19, offset: 5817},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 19, offset: 5817},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 24, offset: 5822},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 37, offset: 5835},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 183, col: 39, offset: 5837},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 43, offset: 5841},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 183, col: 45, offset: 5843},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 48, offset: 5846},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 187, col: 1, offset: 5947},
			expr: &choiceExpr{
				pos: position{line: 187, col: 22, offset: 5968},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 187, col: 22, offset: 5968},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 187, col: 22, offset: 5968},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 187, col: 22, offset: 5968},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 24, offset: 5970},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 28, offset: 5974},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 30, offset: 5976},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 35, offset: 5981},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 46, offset: 5992},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 5, offset: 5999},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 9, offset: 6003},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 11, offset: 6005},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 17, offset: 6011},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 33, offset: 6027},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 35, offset: 6029},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 188, col: 40, offset: 6034},
										expr: &seqExpr{
											pos: position{line: 188, col: 41, offset: 6035},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 188, col: 41, offset: 6035},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 188, col: 45, offset: 6039},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 188, col: 47, offset: 6041},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 188, col: 63, offset: 6057},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 188, col: 67, offset: 6061},
									expr: &litMatcher{
										pos:        position{line: 188, col: 67, offset: 6061},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 72, offset: 6066},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 74, offset: 6068},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 207, col: 1, offset: 6674},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 207, col: 1, offset: 6674},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 207, col: 1, offset: 6674},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 207, col: 3, offset: 6676},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 207, col: 7, offset: 6680},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 207, col: 9, offset: 6682},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 207, col: 14, offset: 6687},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 207, col: 25, offset: 6698},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 207, col: 30, offset: 6703},
										expr: &seqExpr{
											pos: position{line: 207, col: 31, offset: 6704},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 207, col: 31, offset: 6704},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 207, col: 34, offset: 6707},
													name: "TypeAnnotation",
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
		},
		{
			name: "Assignment",
			pos:  position{line: 224, col: 1, offset: 7159},
			expr: &choiceExpr{
				pos: position{line: 224, col: 14, offset: 7172},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 224, col: 14, offset: 7172},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 224, col: 14, offset: 7172},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 224, col: 14, offset: 7172},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 224, col: 16, offset: 7174},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 22, offset: 7180},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 26, offset: 7184},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 28, offset: 7186},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 39, offset: 7197},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 224, col: 42, offset: 7200},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 46, offset: 7204},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 49, offset: 7207},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 54, offset: 7212},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 59, offset: 7217},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 230, col: 1, offset: 7336},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 230, col: 1, offset: 7336},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 230, col: 1, offset: 7336},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 230, col: 3, offset: 7338},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 230, col: 9, offset: 7344},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 230, col: 13, offset: 7348},
									expr: &ruleRefExpr{
										pos:  position{line: 230, col: 14, offset: 7349},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 234, col: 1, offset: 7447},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 234, col: 1, offset: 7447},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 234, col: 1, offset: 7447},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 234, col: 3, offset: 7449},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 234, col: 9, offset: 7455},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 234, col: 13, offset: 7459},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 234, col: 15, offset: 7461},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 234, col: 26, offset: 7472},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 234, col: 29, offset: 7475},
									expr: &litMatcher{
										pos:        position{line: 234, col: 30, offset: 7476},
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
			pos:  position{line: 238, col: 1, offset: 7570},
			expr: &actionExpr{
				pos: position{line: 238, col: 12, offset: 7581},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 238, col: 12, offset: 7581},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 238, col: 12, offset: 7581},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 238, col: 14, offset: 7583},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 20, offset: 7589},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 238, col: 24, offset: 7593},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 238, col: 26, offset: 7595},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 39, offset: 7608},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 238, col: 42, offset: 7611},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 46, offset: 7615},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 238, col: 49, offset: 7618},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 238, col: 53, offset: 7622},
								expr: &seqExpr{
									pos: position{line: 238, col: 54, offset: 7623},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 238, col: 54, offset: 7623},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 63, offset: 7632},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 238, col: 67, offset: 7636},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 238, col: 71, offset: 7640},
								expr: &seqExpr{
									pos: position{line: 238, col: 72, offset: 7641},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 238, col: 72, offset: 7641},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 238, col: 74, offset: 7643},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 79, offset: 7648},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 81, offset: 7650},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 96, offset: 7665},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 238, col: 100, offset: 7669},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 104, offset: 7673},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 238, col: 107, offset: 7676},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 238, col: 118, offset: 7687},
								expr: &ruleRefExpr{
									pos:  position{line: 238, col: 119, offset: 7688},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 131, offset: 7700},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 238, col: 133, offset: 7702},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 137, offset: 7706},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 264, col: 1, offset: 8301},
			expr: &actionExpr{
				pos: position{line: 264, col: 8, offset: 8308},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 264, col: 8, offset: 8308},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 264, col: 12, offset: 8312},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 264, col: 12, offset: 8312},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 264, col: 21, offset: 8321},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 270, col: 1, offset: 8438},
			expr: &choiceExpr{
				pos: position{line: 270, col: 10, offset: 8447},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 270, col: 10, offset: 8447},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 270, col: 10, offset: 8447},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 270, col: 10, offset: 8447},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 270, col: 12, offset: 8449},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 17, offset: 8454},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 21, offset: 8458},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 26, offset: 8463},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 36, offset: 8473},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 270, col: 39, offset: 8476},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 43, offset: 8480},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 45, offset: 8482},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 270, col: 51, offset: 8488},
										expr: &ruleRefExpr{
											pos:  position{line: 270, col: 52, offset: 8489},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 64, offset: 8501},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 270, col: 67, offset: 8504},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 71, offset: 8508},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 270, col: 74, offset: 8511},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 81, offset: 8518},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 270, col: 84, offset: 8521},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 88, offset: 8525},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 90, offset: 8527},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 270, col: 96, offset: 8533},
										expr: &ruleRefExpr{
											pos:  position{line: 270, col: 97, offset: 8534},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 109, offset: 8546},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 270, col: 112, offset: 8549},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 289, col: 1, offset: 9052},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 289, col: 1, offset: 9052},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 289, col: 1, offset: 9052},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 289, col: 3, offset: 9054},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 8, offset: 9059},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 289, col: 12, offset: 9063},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 289, col: 17, offset: 9068},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 27, offset: 9078},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 289, col: 30, offset: 9081},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 34, offset: 9085},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 289, col: 36, offset: 9087},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 289, col: 42, offset: 9093},
										expr: &ruleRefExpr{
											pos:  position{line: 289, col: 43, offset: 9094},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 55, offset: 9106},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 289, col: 57, offset: 9108},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 61, offset: 9112},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 289, col: 64, offset: 9115},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 289, col: 71, offset: 9122},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 289, col: 79, offset: 9130},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 301, col: 1, offset: 9460},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 301, col: 1, offset: 9460},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 301, col: 1, offset: 9460},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 301, col: 3, offset: 9462},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 8, offset: 9467},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 301, col: 12, offset: 9471},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 301, col: 17, offset: 9476},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 27, offset: 9486},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 301, col: 30, offset: 9489},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 34, offset: 9493},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 301, col: 36, offset: 9495},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 301, col: 42, offset: 9501},
										expr: &ruleRefExpr{
											pos:  position{line: 301, col: 43, offset: 9502},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 55, offset: 9514},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 301, col: 58, offset: 9517},
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
			pos:  position{line: 313, col: 1, offset: 9815},
			expr: &choiceExpr{
				pos: position{line: 313, col: 8, offset: 9822},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 313, col: 8, offset: 9822},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 313, col: 8, offset: 9822},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 313, col: 8, offset: 9822},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 313, col: 10, offset: 9824},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 17, offset: 9831},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 313, col: 28, offset: 9842},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 313, col: 32, offset: 9846},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 35, offset: 9849},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 313, col: 48, offset: 9862},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 53, offset: 9867},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 327, col: 1, offset: 10191},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 327, col: 1, offset: 10191},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 327, col: 1, offset: 10191},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 327, col: 3, offset: 10193},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 327, col: 6, offset: 10196},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 327, col: 19, offset: 10209},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 327, col: 24, offset: 10214},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 341, col: 1, offset: 10531},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 341, col: 1, offset: 10531},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 341, col: 1, offset: 10531},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 341, col: 3, offset: 10533},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 341, col: 6, offset: 10536},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 341, col: 19, offset: 10549},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "RecordAccess",
			pos:  position{line: 348, col: 1, offset: 10720},
			expr: &actionExpr{
				pos: position{line: 348, col: 16, offset: 10735},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 348, col: 16, offset: 10735},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 348, col: 16, offset: 10735},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 23, offset: 10742},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 348, col: 36, offset: 10755},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 348, col: 41, offset: 10760},
								expr: &seqExpr{
									pos: position{line: 348, col: 42, offset: 10761},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 348, col: 42, offset: 10761},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 46, offset: 10765},
											name: "VariableName",
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
			name: "ArrayAccess",
			pos:  position{line: 365, col: 1, offset: 11202},
			expr: &actionExpr{
				pos: position{line: 365, col: 15, offset: 11216},
				run: (*parser).callonArrayAccess1,
				expr: &seqExpr{
					pos: position{line: 365, col: 15, offset: 11216},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 365, col: 15, offset: 11216},
							label: "array",
							expr: &ruleRefExpr{
								pos:  position{line: 365, col: 21, offset: 11222},
								name: "VariableName",
							},
						},
						&litMatcher{
							pos:        position{line: 365, col: 34, offset: 11235},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 365, col: 38, offset: 11239},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 365, col: 40, offset: 11241},
								name: "Expr",
							},
						},
						&litMatcher{
							pos:        position{line: 365, col: 45, offset: 11246},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgsDefn",
			pos:  position{line: 369, col: 1, offset: 11332},
			expr: &actionExpr{
				pos: position{line: 369, col: 12, offset: 11343},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 369, col: 12, offset: 11343},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 369, col: 12, offset: 11343},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 369, col: 16, offset: 11347},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 369, col: 18, offset: 11349},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 369, col: 27, offset: 11358},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 369, col: 35, offset: 11366},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 369, col: 37, offset: 11368},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 369, col: 42, offset: 11373},
								expr: &seqExpr{
									pos: position{line: 369, col: 43, offset: 11374},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 369, col: 43, offset: 11374},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 369, col: 47, offset: 11378},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 369, col: 49, offset: 11380},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 369, col: 59, offset: 11390},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 369, col: 61, offset: 11392},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 387, col: 1, offset: 11814},
			expr: &actionExpr{
				pos: position{line: 387, col: 11, offset: 11824},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 387, col: 11, offset: 11824},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 387, col: 11, offset: 11824},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 387, col: 16, offset: 11829},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 387, col: 27, offset: 11840},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 387, col: 29, offset: 11842},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 387, col: 34, offset: 11847},
								expr: &seqExpr{
									pos: position{line: 387, col: 35, offset: 11848},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 387, col: 35, offset: 11848},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 387, col: 39, offset: 11852},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 387, col: 41, offset: 11854},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 387, col: 59, offset: 11872},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 408, col: 1, offset: 12411},
			expr: &choiceExpr{
				pos: position{line: 408, col: 18, offset: 12428},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 18, offset: 12428},
						name: "ArrayType",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 30, offset: 12440},
						name: "AnyType",
					},
					&actionExpr{
						pos: position{line: 409, col: 1, offset: 12451},
						run: (*parser).callonTypeAnnotation4,
						expr: &seqExpr{
							pos: position{line: 409, col: 1, offset: 12451},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 409, col: 1, offset: 12451},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 409, col: 8, offset: 12458},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 409, col: 11, offset: 12461},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 409, col: 16, offset: 12466},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 409, col: 25, offset: 12475},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 409, col: 27, offset: 12477},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 409, col: 32, offset: 12482},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 409, col: 34, offset: 12484},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 409, col: 38, offset: 12488},
										name: "TypeAnnotation",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 418, col: 1, offset: 12704},
			expr: &choiceExpr{
				pos: position{line: 418, col: 11, offset: 12714},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 418, col: 11, offset: 12714},
						name: "ModuleName",
					},
					&ruleRefExpr{
						pos:  position{line: 418, col: 24, offset: 12727},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 418, col: 35, offset: 12738},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 420, col: 1, offset: 12753},
			expr: &choiceExpr{
				pos: position{line: 420, col: 13, offset: 12765},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 420, col: 13, offset: 12765},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 420, col: 13, offset: 12765},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 420, col: 13, offset: 12765},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 420, col: 17, offset: 12769},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 420, col: 19, offset: 12771},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 420, col: 28, offset: 12780},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 420, col: 40, offset: 12792},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 420, col: 42, offset: 12794},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 420, col: 47, offset: 12799},
										expr: &seqExpr{
											pos: position{line: 420, col: 48, offset: 12800},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 420, col: 48, offset: 12800},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 420, col: 52, offset: 12804},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 420, col: 54, offset: 12806},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 420, col: 68, offset: 12820},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 420, col: 70, offset: 12822},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 437, col: 1, offset: 13244},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 437, col: 1, offset: 13244},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 437, col: 1, offset: 13244},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 437, col: 5, offset: 13248},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 437, col: 7, offset: 13250},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 437, col: 16, offset: 13259},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 437, col: 21, offset: 13264},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 437, col: 23, offset: 13266},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 442, col: 1, offset: 13372},
						run: (*parser).callonArguments25,
						expr: &seqExpr{
							pos: position{line: 442, col: 1, offset: 13372},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 442, col: 1, offset: 13372},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 442, col: 5, offset: 13376},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 442, col: 7, offset: 13378},
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
			name: "CompoundExpr",
			pos:  position{line: 446, col: 1, offset: 13432},
			expr: &actionExpr{
				pos: position{line: 446, col: 16, offset: 13447},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 446, col: 16, offset: 13447},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 446, col: 16, offset: 13447},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 446, col: 18, offset: 13449},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 446, col: 21, offset: 13452},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 446, col: 27, offset: 13458},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 446, col: 32, offset: 13463},
								expr: &seqExpr{
									pos: position{line: 446, col: 33, offset: 13464},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 446, col: 33, offset: 13464},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 446, col: 37, offset: 13468},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 446, col: 46, offset: 13477},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 446, col: 50, offset: 13481},
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
			pos:  position{line: 466, col: 1, offset: 14087},
			expr: &choiceExpr{
				pos: position{line: 466, col: 9, offset: 14095},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 466, col: 9, offset: 14095},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 466, col: 21, offset: 14107},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 466, col: 37, offset: 14123},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 466, col: 48, offset: 14134},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 466, col: 60, offset: 14146},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 468, col: 1, offset: 14159},
			expr: &actionExpr{
				pos: position{line: 468, col: 13, offset: 14171},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 468, col: 13, offset: 14171},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 468, col: 13, offset: 14171},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 468, col: 15, offset: 14173},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 468, col: 21, offset: 14179},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 468, col: 35, offset: 14193},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 468, col: 40, offset: 14198},
								expr: &seqExpr{
									pos: position{line: 468, col: 41, offset: 14199},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 468, col: 41, offset: 14199},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 468, col: 45, offset: 14203},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 468, col: 61, offset: 14219},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 468, col: 65, offset: 14223},
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
			pos:  position{line: 501, col: 1, offset: 15116},
			expr: &actionExpr{
				pos: position{line: 501, col: 17, offset: 15132},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 501, col: 17, offset: 15132},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 501, col: 17, offset: 15132},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 501, col: 19, offset: 15134},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 501, col: 25, offset: 15140},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 501, col: 34, offset: 15149},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 501, col: 39, offset: 15154},
								expr: &seqExpr{
									pos: position{line: 501, col: 40, offset: 15155},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 501, col: 40, offset: 15155},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 501, col: 44, offset: 15159},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 501, col: 61, offset: 15176},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 501, col: 65, offset: 15180},
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
			pos:  position{line: 533, col: 1, offset: 16067},
			expr: &actionExpr{
				pos: position{line: 533, col: 12, offset: 16078},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 533, col: 12, offset: 16078},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 533, col: 12, offset: 16078},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 533, col: 14, offset: 16080},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 20, offset: 16086},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 533, col: 30, offset: 16096},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 533, col: 35, offset: 16101},
								expr: &seqExpr{
									pos: position{line: 533, col: 36, offset: 16102},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 533, col: 36, offset: 16102},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 533, col: 40, offset: 16106},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 533, col: 52, offset: 16118},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 533, col: 56, offset: 16122},
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
			pos:  position{line: 565, col: 1, offset: 17010},
			expr: &actionExpr{
				pos: position{line: 565, col: 13, offset: 17022},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 565, col: 13, offset: 17022},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 565, col: 13, offset: 17022},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 565, col: 15, offset: 17024},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 565, col: 21, offset: 17030},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 565, col: 33, offset: 17042},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 565, col: 38, offset: 17047},
								expr: &seqExpr{
									pos: position{line: 565, col: 39, offset: 17048},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 565, col: 39, offset: 17048},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 565, col: 43, offset: 17052},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 565, col: 56, offset: 17065},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 565, col: 60, offset: 17069},
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
			pos:  position{line: 596, col: 1, offset: 17958},
			expr: &choiceExpr{
				pos: position{line: 596, col: 15, offset: 17972},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 596, col: 15, offset: 17972},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 596, col: 15, offset: 17972},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 596, col: 15, offset: 17972},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 596, col: 17, offset: 17974},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 596, col: 21, offset: 17978},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 596, col: 24, offset: 17981},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 596, col: 30, offset: 17987},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 596, col: 36, offset: 17993},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 596, col: 39, offset: 17996},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 5, offset: 18119},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 601, col: 1, offset: 18126},
			expr: &choiceExpr{
				pos: position{line: 601, col: 12, offset: 18137},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 601, col: 12, offset: 18137},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 601, col: 30, offset: 18155},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 601, col: 49, offset: 18174},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 601, col: 64, offset: 18189},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 603, col: 1, offset: 18202},
			expr: &actionExpr{
				pos: position{line: 603, col: 19, offset: 18220},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 603, col: 21, offset: 18222},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 603, col: 21, offset: 18222},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 603, col: 28, offset: 18229},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 607, col: 1, offset: 18311},
			expr: &actionExpr{
				pos: position{line: 607, col: 20, offset: 18330},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 607, col: 22, offset: 18332},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 607, col: 22, offset: 18332},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 607, col: 29, offset: 18339},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 607, col: 36, offset: 18346},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 607, col: 42, offset: 18352},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 607, col: 48, offset: 18358},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 607, col: 56, offset: 18366},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 611, col: 1, offset: 18445},
			expr: &choiceExpr{
				pos: position{line: 611, col: 16, offset: 18460},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 611, col: 16, offset: 18460},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 611, col: 18, offset: 18462},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 611, col: 18, offset: 18462},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 611, col: 25, offset: 18469},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 614, col: 3, offset: 18552},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 614, col: 5, offset: 18554},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 614, col: 5, offset: 18554},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 614, col: 11, offset: 18560},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 617, col: 3, offset: 18640},
						run: (*parser).callonOperatorHigh10,
						expr: &litMatcher{
							pos:        position{line: 617, col: 5, offset: 18642},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 620, col: 3, offset: 18722},
						run: (*parser).callonOperatorHigh12,
						expr: &litMatcher{
							pos:        position{line: 620, col: 3, offset: 18722},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 624, col: 1, offset: 18803},
			expr: &choiceExpr{
				pos: position{line: 624, col: 15, offset: 18817},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 624, col: 15, offset: 18817},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 624, col: 17, offset: 18819},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 624, col: 17, offset: 18819},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 624, col: 24, offset: 18826},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 627, col: 3, offset: 18910},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 627, col: 5, offset: 18912},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 627, col: 5, offset: 18912},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 627, col: 11, offset: 18918},
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
			pos:  position{line: 631, col: 1, offset: 18997},
			expr: &choiceExpr{
				pos: position{line: 631, col: 9, offset: 19005},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 631, col: 9, offset: 19005},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 631, col: 16, offset: 19012},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 631, col: 31, offset: 19027},
						name: "ArrayAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 631, col: 45, offset: 19041},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 631, col: 60, offset: 19056},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 633, col: 1, offset: 19063},
			expr: &choiceExpr{
				pos: position{line: 633, col: 14, offset: 19076},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 633, col: 14, offset: 19076},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 633, col: 29, offset: 19091},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 635, col: 1, offset: 19099},
			expr: &choiceExpr{
				pos: position{line: 635, col: 14, offset: 19112},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 635, col: 14, offset: 19112},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 635, col: 29, offset: 19127},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 637, col: 1, offset: 19139},
			expr: &actionExpr{
				pos: position{line: 637, col: 16, offset: 19154},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 637, col: 16, offset: 19154},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 637, col: 16, offset: 19154},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 637, col: 20, offset: 19158},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 637, col: 22, offset: 19160},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 637, col: 28, offset: 19166},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 637, col: 33, offset: 19171},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 637, col: 35, offset: 19173},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 637, col: 40, offset: 19178},
								expr: &seqExpr{
									pos: position{line: 637, col: 41, offset: 19179},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 637, col: 41, offset: 19179},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 637, col: 45, offset: 19183},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 637, col: 47, offset: 19185},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 637, col: 52, offset: 19190},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 637, col: 56, offset: 19194},
							expr: &litMatcher{
								pos:        position{line: 637, col: 56, offset: 19194},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 637, col: 61, offset: 19199},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 637, col: 63, offset: 19201},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 653, col: 1, offset: 19642},
			expr: &actionExpr{
				pos: position{line: 653, col: 19, offset: 19660},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 653, col: 19, offset: 19660},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 653, col: 19, offset: 19660},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 653, col: 24, offset: 19665},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 653, col: 35, offset: 19676},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 653, col: 37, offset: 19678},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 653, col: 42, offset: 19683},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 666, col: 1, offset: 19955},
			expr: &actionExpr{
				pos: position{line: 666, col: 18, offset: 19972},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 666, col: 18, offset: 19972},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 666, col: 18, offset: 19972},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 666, col: 23, offset: 19977},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 666, col: 34, offset: 19988},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 666, col: 36, offset: 19990},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 666, col: 40, offset: 19994},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 666, col: 42, offset: 19996},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 666, col: 52, offset: 20006},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 666, col: 65, offset: 20019},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 666, col: 67, offset: 20021},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 666, col: 71, offset: 20025},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 666, col: 73, offset: 20027},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 666, col: 84, offset: 20038},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 666, col: 89, offset: 20043},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 666, col: 94, offset: 20048},
								expr: &seqExpr{
									pos: position{line: 666, col: 95, offset: 20049},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 666, col: 95, offset: 20049},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 666, col: 99, offset: 20053},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 666, col: 101, offset: 20055},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 666, col: 114, offset: 20068},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 666, col: 116, offset: 20070},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 666, col: 120, offset: 20074},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 666, col: 122, offset: 20076},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 666, col: 130, offset: 20084},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 686, col: 1, offset: 20674},
			expr: &actionExpr{
				pos: position{line: 686, col: 17, offset: 20690},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 686, col: 17, offset: 20690},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 686, col: 17, offset: 20690},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 686, col: 22, offset: 20695},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 690, col: 1, offset: 20768},
			expr: &actionExpr{
				pos: position{line: 690, col: 16, offset: 20783},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 690, col: 16, offset: 20783},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 690, col: 16, offset: 20783},
							expr: &ruleRefExpr{
								pos:  position{line: 690, col: 17, offset: 20784},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 690, col: 27, offset: 20794},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 690, col: 27, offset: 20794},
									expr: &charClassMatcher{
										pos:        position{line: 690, col: 27, offset: 20794},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 690, col: 34, offset: 20801},
									expr: &charClassMatcher{
										pos:        position{line: 690, col: 34, offset: 20801},
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
			pos:  position{line: 694, col: 1, offset: 20876},
			expr: &actionExpr{
				pos: position{line: 694, col: 14, offset: 20889},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 694, col: 15, offset: 20890},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 694, col: 15, offset: 20890},
							expr: &charClassMatcher{
								pos:        position{line: 694, col: 15, offset: 20890},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 694, col: 22, offset: 20897},
							expr: &charClassMatcher{
								pos:        position{line: 694, col: 22, offset: 20897},
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
			pos:  position{line: 698, col: 1, offset: 20972},
			expr: &choiceExpr{
				pos: position{line: 698, col: 9, offset: 20980},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 698, col: 9, offset: 20980},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 698, col: 9, offset: 20980},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 698, col: 9, offset: 20980},
									expr: &litMatcher{
										pos:        position{line: 698, col: 9, offset: 20980},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 698, col: 14, offset: 20985},
									expr: &charClassMatcher{
										pos:        position{line: 698, col: 14, offset: 20985},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 698, col: 21, offset: 20992},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 698, col: 25, offset: 20996},
									expr: &charClassMatcher{
										pos:        position{line: 698, col: 25, offset: 20996},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 705, col: 3, offset: 21185},
						run: (*parser).callonConst11,
						expr: &seqExpr{
							pos: position{line: 705, col: 3, offset: 21185},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 705, col: 3, offset: 21185},
									expr: &litMatcher{
										pos:        position{line: 705, col: 3, offset: 21185},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 705, col: 8, offset: 21190},
									expr: &charClassMatcher{
										pos:        position{line: 705, col: 8, offset: 21190},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 705, col: 15, offset: 21197},
									expr: &litMatcher{
										pos:        position{line: 705, col: 16, offset: 21198},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 712, col: 3, offset: 21374},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 712, col: 3, offset: 21374},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 716, col: 3, offset: 21459},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 716, col: 3, offset: 21459},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 719, col: 3, offset: 21545},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 720, col: 3, offset: 21552},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 721, col: 3, offset: 21568},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 721, col: 3, offset: 21568},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 721, col: 3, offset: 21568},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 721, col: 7, offset: 21572},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 721, col: 12, offset: 21577},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 721, col: 12, offset: 21577},
												expr: &ruleRefExpr{
													pos:  position{line: 721, col: 13, offset: 21578},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 721, col: 25, offset: 21590,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 721, col: 28, offset: 21593},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 723, col: 5, offset: 21685},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 723, col: 20, offset: 21700},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 723, col: 37, offset: 21717},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 725, col: 1, offset: 21734},
			expr: &actionExpr{
				pos: position{line: 725, col: 8, offset: 21741},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 725, col: 8, offset: 21741},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 729, col: 1, offset: 21804},
			expr: &actionExpr{
				pos: position{line: 729, col: 17, offset: 21820},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 729, col: 17, offset: 21820},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 729, col: 17, offset: 21820},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 729, col: 21, offset: 21824},
							expr: &charClassMatcher{
								pos:        position{line: 729, col: 22, offset: 21825},
								val:        "[^\\r\\n\"]",
								chars:      []rune{'\r', '\n', '"'},
								ignoreCase: false,
								inverted:   true,
							},
						},
						&litMatcher{
							pos:        position{line: 729, col: 33, offset: 21836},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 737, col: 1, offset: 21996},
			expr: &actionExpr{
				pos: position{line: 737, col: 10, offset: 22005},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 737, col: 11, offset: 22006},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 741, col: 1, offset: 22061},
			expr: &seqExpr{
				pos: position{line: 741, col: 12, offset: 22072},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 741, col: 13, offset: 22073},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 741, col: 13, offset: 22073},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 21, offset: 22081},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 28, offset: 22088},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 37, offset: 22097},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 48, offset: 22108},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 57, offset: 22117},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 66, offset: 22126},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 76, offset: 22136},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 741, col: 88, offset: 22148},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 741, col: 97, offset: 22157},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 741, col: 107, offset: 22167},
						expr: &oneOrMoreExpr{
							pos: position{line: 741, col: 108, offset: 22168},
							expr: &charClassMatcher{
								pos:        position{line: 741, col: 108, offset: 22168},
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
			name: "ArrayType",
			pos:  position{line: 743, col: 1, offset: 22176},
			expr: &actionExpr{
				pos: position{line: 743, col: 13, offset: 22188},
				run: (*parser).callonArrayType1,
				expr: &seqExpr{
					pos: position{line: 743, col: 13, offset: 22188},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 743, col: 13, offset: 22188},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 743, col: 17, offset: 22192},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 743, col: 19, offset: 22194},
								name: "TypeAnnotation",
							},
						},
						&litMatcher{
							pos:        position{line: 743, col: 34, offset: 22209},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BaseType",
			pos:  position{line: 747, col: 1, offset: 22262},
			expr: &choiceExpr{
				pos: position{line: 747, col: 12, offset: 22273},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 747, col: 12, offset: 22273},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 747, col: 14, offset: 22275},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 747, col: 14, offset: 22275},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 747, col: 24, offset: 22285},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 747, col: 33, offset: 22294},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 747, col: 44, offset: 22305},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 747, col: 53, offset: 22314},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 747, col: 62, offset: 22323},
									val:        "float64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 747, col: 74, offset: 22335},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 750, col: 3, offset: 22433},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 752, col: 1, offset: 22439},
			expr: &charClassMatcher{
				pos:        position{line: 752, col: 15, offset: 22453},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 754, col: 1, offset: 22469},
			expr: &choiceExpr{
				pos: position{line: 754, col: 18, offset: 22486},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 754, col: 18, offset: 22486},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 754, col: 37, offset: 22505},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 756, col: 1, offset: 22520},
			expr: &charClassMatcher{
				pos:        position{line: 756, col: 20, offset: 22539},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 758, col: 1, offset: 22552},
			expr: &charClassMatcher{
				pos:        position{line: 758, col: 16, offset: 22567},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 760, col: 1, offset: 22574},
			expr: &charClassMatcher{
				pos:        position{line: 760, col: 23, offset: 22596},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 762, col: 1, offset: 22603},
			expr: &charClassMatcher{
				pos:        position{line: 762, col: 12, offset: 22614},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 764, col: 1, offset: 22625},
			expr: &choiceExpr{
				pos: position{line: 764, col: 22, offset: 22646},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 764, col: 22, offset: 22646},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 764, col: 33, offset: 22657},
						expr: &charClassMatcher{
							pos:        position{line: 764, col: 33, offset: 22657},
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
			displayName: "\"optwhitespace\"",
			pos:         position{line: 766, col: 1, offset: 22669},
			expr: &choiceExpr{
				pos: position{line: 766, col: 21, offset: 22689},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 766, col: 21, offset: 22689},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 766, col: 32, offset: 22700},
						expr: &charClassMatcher{
							pos:        position{line: 766, col: 32, offset: 22700},
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
			name:        "__N",
			displayName: "\"singleline_reqwhitepace\"",
			pos:         position{line: 768, col: 1, offset: 22712},
			expr: &oneOrMoreExpr{
				pos: position{line: 768, col: 33, offset: 22744},
				expr: &charClassMatcher{
					pos:        position{line: 768, col: 33, offset: 22744},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "_N",
			displayName: "\"singleline_optwhitepace\"",
			pos:         position{line: 770, col: 1, offset: 22752},
			expr: &zeroOrMoreExpr{
				pos: position{line: 770, col: 32, offset: 22783},
				expr: &charClassMatcher{
					pos:        position{line: 770, col: 32, offset: 22783},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "N",
			displayName: "\"newline\"",
			pos:         position{line: 772, col: 1, offset: 22791},
			expr: &choiceExpr{
				pos: position{line: 772, col: 15, offset: 22805},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 772, col: 15, offset: 22805},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 772, col: 26, offset: 22816},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 772, col: 26, offset: 22816},
								expr: &charClassMatcher{
									pos:        position{line: 772, col: 26, offset: 22816},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 772, col: 35, offset: 22825},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Comments",
			pos:  position{line: 774, col: 1, offset: 22831},
			expr: &oneOrMoreExpr{
				pos: position{line: 774, col: 12, offset: 22842},
				expr: &ruleRefExpr{
					pos:  position{line: 774, col: 13, offset: 22843},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 776, col: 1, offset: 22854},
			expr: &choiceExpr{
				pos: position{line: 776, col: 11, offset: 22864},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 776, col: 11, offset: 22864},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 776, col: 11, offset: 22864},
								expr: &charClassMatcher{
									pos:        position{line: 776, col: 11, offset: 22864},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 776, col: 22, offset: 22875},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 776, col: 27, offset: 22880},
								expr: &seqExpr{
									pos: position{line: 776, col: 28, offset: 22881},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 776, col: 28, offset: 22881},
											expr: &charClassMatcher{
												pos:        position{line: 776, col: 29, offset: 22882},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 776, col: 34, offset: 22887,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 776, col: 38, offset: 22891},
								expr: &litMatcher{
									pos:        position{line: 776, col: 39, offset: 22892},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 776, col: 46, offset: 22899},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 776, col: 46, offset: 22899},
								expr: &charClassMatcher{
									pos:        position{line: 776, col: 46, offset: 22899},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 776, col: 57, offset: 22910},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 776, col: 62, offset: 22915},
								expr: &seqExpr{
									pos: position{line: 776, col: 63, offset: 22916},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 776, col: 63, offset: 22916},
											expr: &litMatcher{
												pos:        position{line: 776, col: 64, offset: 22917},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 776, col: 69, offset: 22922,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 776, col: 73, offset: 22926},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 776, col: 78, offset: 22931},
								expr: &charClassMatcher{
									pos:        position{line: 776, col: 78, offset: 22931},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 776, col: 84, offset: 22937},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 778, col: 1, offset: 22943},
			expr: &notExpr{
				pos: position{line: 778, col: 7, offset: 22949},
				expr: &anyMatcher{
					line: 778, col: 8, offset: 22950,
				},
			},
		},
	},
}

func (c *current) onModule1(name, stat, rest interface{}) (interface{}, error) {
	//fmt.Println("beginning module")
	vals := rest.([]interface{})
	if len(vals) > 0 {
		//fmt.Println("multiple statements")
		subvalues := []Ast{stat.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return Module{Name: name.(Identifier).StringValue, Subvalues: subvalues, Imports: make(map[string]bool)}, nil
	} else {
		return Module{Name: name.(Identifier).StringValue, Subvalues: []Ast{stat.(Ast)}, Imports: make(map[string]bool)}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["name"], stack["stat"], stack["rest"])
}

func (c *current) onExprLine1(e interface{}) (interface{}, error) {

	// wrap calls as statements in an expr
	switch e.(type) {
	case Call:
		ex := Expr{Subvalues: []Ast{e.(Call)}, AsStatement: true}
		return ex, nil
	case Expr:
		ex := Expr{Subvalues: e.(Expr).Subvalues, AsStatement: true}
		return ex, nil
	}
	return e, nil
}

func (p *parser) callonExprLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExprLine1(stack["e"])
}

func (c *current) onExternFunc1(name, importName, args, ret interface{}) (interface{}, error) {

	return ExternFunc{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Arguments: args.(Container).Subvalues, ReturnAnnotation: ret.(BasicAst)}, nil
}

func (p *parser) callonExternFunc1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternFunc1(stack["name"], stack["importName"], stack["args"], stack["ret"])
}

func (c *current) onExternType2(name, importName, first, rest interface{}) (interface{}, error) {
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

	return ExternRecordType{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Fields: fields}, nil
}

func (p *parser) callonExternType2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternType2(stack["name"], stack["importName"], stack["first"], stack["rest"])
}

func (c *current) onExternType34(name, importName interface{}) (interface{}, error) {
	// record type
	fields := []RecordField{}

	return ExternRecordType{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Fields: fields}, nil
}

func (p *parser) callonExternType34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternType34(stack["name"], stack["importName"])
}

func (c *current) onTypeDefn2(name, params, types interface{}) (interface{}, error) {
	// alias type
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

	return AliasType{Name: name.(Identifier).StringValue, Params: params.(Container).Subvalues, Types: fields}, nil
}

func (p *parser) callonTypeDefn2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn2(stack["name"], stack["params"], stack["types"])
}

func (c *current) onTypeDefn20(name, params, first, rest interface{}) (interface{}, error) {
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

func (p *parser) callonTypeDefn20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn20(stack["name"], stack["params"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn50(name, first, rest interface{}) (interface{}, error) {
	// record type, no type params
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

func (p *parser) callonTypeDefn50() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn50(stack["name"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn77(name, params, rest interface{}) (interface{}, error) {
	// variant type with params
	constructors := []VariantConstructor{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		for _, v := range vals {
			constructors = append(constructors, v.(VariantConstructor))
		}
	}

	return Variant{Name: name.(Identifier).StringValue, Params: params.(Container).Subvalues, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn77() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn77(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onTypeDefn94(name, rest interface{}) (interface{}, error) {
	// variant type
	parameters := []Ast{}
	constructors := []VariantConstructor{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		for _, v := range vals {
			constructors = append(constructors, v.(VariantConstructor))
		}
	}

	return Variant{Name: name.(Identifier).StringValue, Params: parameters, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn94() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn94(stack["name"], stack["rest"])
}

func (c *current) onTypeParams1(first, rest interface{}) (interface{}, error) {
	args := []Ast{first.(Identifier)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Identifier)
			args = append(args, v)
		}
	}
	return Container{Subvalues: args}, nil
}

func (p *parser) callonTypeParams1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeParams1(stack["first"], stack["rest"])
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
	return nil, errors.New("Variable name or '_' (unused result) required here")
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

func (c *current) onFuncDefn1(i, ids, ret, statements interface{}) (interface{}, error) {
	//fmt.Println(string(c.text))
	subvalues := []Ast{}
	vals := statements.([]interface{})
	args := []Ast{}

	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}

	if ids != nil {
		vals = ids.([]interface{})
		args = vals[0].(Container).Subvalues
	}

	var retType Ast
	if ret != nil {
		vals = ret.([]interface{})
		retType = vals[3].(Ast)
	}

	return Func{Name: i.(Identifier).StringValue, Arguments: args, Subvalues: subvalues, ReturnAnnotation: retType}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["i"], stack["ids"], stack["ret"], stack["statements"])
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

func (c *current) onIfExpr2(expr, thens, elses interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr2(stack["expr"], stack["thens"], stack["elses"])
}

func (c *current) onIfExpr27(expr, thens, elseifs interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr27() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr27(stack["expr"], stack["thens"], stack["elseifs"])
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

func (c *current) onCall2(module, fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: module.(Identifier), Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall2(stack["module"], stack["fn"], stack["args"])
}

func (c *current) onCall12(fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: Identifier{}, Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall12(stack["fn"], stack["args"])
}

func (c *current) onCall19(fn interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	return Call{Module: Identifier{}, Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall19(stack["fn"])
}

func (c *current) onRecordAccess1(record, rest interface{}) (interface{}, error) {
	args := []Identifier{record.(Identifier)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Identifier)
			args = append(args, v)
		}
	}

	return RecordAccess{Identifiers: args}, nil
}

func (p *parser) callonRecordAccess1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordAccess1(stack["record"], stack["rest"])
}

func (c *current) onArrayAccess1(array, e interface{}) (interface{}, error) {
	return ArrayAccess{Identifier: array.(Identifier), Index: e.(Expr)}, nil
}

func (p *parser) callonArrayAccess1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArrayAccess1(stack["array"], stack["e"])
}

func (c *current) onArgsDefn1(argument, rest interface{}) (interface{}, error) {

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

func (p *parser) callonArgsDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgsDefn1(stack["argument"], stack["rest"])
}

func (c *current) onArgDefn1(name, anno interface{}) (interface{}, error) {
	//fmt.Println("parsing arg:", string(c.text))
	arg := name.(Identifier)

	if anno != nil {
		vals := anno.([]interface{})
		//fmt.Println(vals)
		//restSl := toIfaceSlice(vals[0])

		//switch vals[2].(type) {
		//case BasicAst:
		//    arg.Annotation = vals[2].(BasicAst).StringValue
		//case Identifier:
		//    arg.Annotation = vals[2].(Identifier).StringValue
		//}
		arg.Annotation = vals[2].(Ast)
	}
	//fmt.Println("parsed:", arg)
	return arg, nil
}

func (p *parser) callonArgDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgDefn1(stack["name"], stack["anno"])
}

func (c *current) onTypeAnnotation4(args, ret interface{}) (interface{}, error) {
	// TODO: return correct func type annotation
	vals := args.(Container)
	vals.Subvalues = append(vals.Subvalues, ret.(Ast))
	vals.Type = "FuncAnnotation"

	return vals, nil
}

func (p *parser) callonTypeAnnotation4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeAnnotation4(stack["args"], stack["ret"])
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

func (c *current) onArguments25() (interface{}, error) {
	return Container{Type: "Arguments"}, nil
}

func (p *parser) callonArguments25() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments25()
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

	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil

}

func (p *parser) callonBinOpBool1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpBool1(stack["first"], stack["rest"])
}

func (c *current) onBinOpEquality1(first, rest interface{}) (interface{}, error) {
	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil

}

func (p *parser) callonBinOpEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpEquality1(stack["first"], stack["rest"])
}

func (c *current) onBinOpLow1(first, rest interface{}) (interface{}, error) {
	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil

}

func (p *parser) callonBinOpLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpLow1(stack["first"], stack["rest"])
}

func (c *current) onBinOpHigh1(first, rest interface{}) (interface{}, error) {
	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil
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
	return Operator{StringValue: string(c.text), ValueType: BOOL}, nil
}

func (p *parser) callonOperatorBoolean1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorBoolean1()
}

func (c *current) onOperatorEquality1() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: BOOL}, nil
}

func (p *parser) callonOperatorEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorEquality1()
}

func (c *current) onOperatorHigh2() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: FLOAT}, nil
}

func (p *parser) callonOperatorHigh2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh2()
}

func (c *current) onOperatorHigh6() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: INT}, nil
}

func (p *parser) callonOperatorHigh6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh6()
}

func (c *current) onOperatorHigh10() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: INT}, nil
}

func (p *parser) callonOperatorHigh10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh10()
}

func (c *current) onOperatorHigh12() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh12()
}

func (c *current) onOperatorLow2() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: FLOAT}, nil
}

func (p *parser) callonOperatorLow2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow2()
}

func (c *current) onOperatorLow6() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: INT}, nil
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
	return Array{Subvalues: subvalues}, nil
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

	return VariantInstance{Name: name.(Identifier).StringValue, Arguments: arguments}, nil
}

func (p *parser) callonVariantInstance1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantInstance1(stack["name"], stack["args"])
}

func (c *current) onRecordInstance1(name, firstName, firstValue, rest interface{}) (interface{}, error) {
	instance := RecordInstance{Name: name.(Identifier).StringValue}
	instance.Values = make(map[string]Ast)

	vals := rest.([]interface{})
	instance.Values[firstName.(Identifier).StringValue] = firstValue.(Ast)

	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			k := restExpr[2].(Identifier).StringValue
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
	val, err := strconv.ParseFloat(string(c.text), 64)
	if err != nil {
		return nil, err
	}
	return BasicAst{Type: "Float", FloatValue: val, ValueType: FLOAT}, nil
}

func (p *parser) callonConst2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst2()
}

func (c *current) onConst11() (interface{}, error) {
	val, err := strconv.Atoi(string(c.text))
	if err != nil {
		return nil, err
	}
	return BasicAst{Type: "Integer", IntValue: val, ValueType: INT}, nil
}

func (p *parser) callonConst11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst11()
}

func (c *current) onConst19() (interface{}, error) {
	return BasicAst{Type: "Bool", BoolValue: true, ValueType: BOOL}, nil

}

func (p *parser) callonConst19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst19()
}

func (c *current) onConst21() (interface{}, error) {
	return BasicAst{Type: "Bool", BoolValue: false, ValueType: BOOL}, nil
}

func (p *parser) callonConst21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst21()
}

func (c *current) onConst25(val interface{}) (interface{}, error) {
	return BasicAst{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst25() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst25(stack["val"])
}

func (c *current) onUnit1() (interface{}, error) {
	return BasicAst{Type: "Unit", ValueType: NIL}, nil
}

func (p *parser) callonUnit1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnit1()
}

func (c *current) onStringLiteral1() (interface{}, error) {
	val := string(c.text[1 : len(c.text)-1])

	return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil

	//return nil, err
}

func (p *parser) callonStringLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringLiteral1()
}

func (c *current) onUnused1() (interface{}, error) {
	return Identifier{StringValue: "_"}, nil
}

func (p *parser) callonUnused1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnused1()
}

func (c *current) onArrayType1(t interface{}) (interface{}, error) {
	return ArrayType{Subtype: t.(Ast)}, nil
}

func (p *parser) callonArrayType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArrayType1(stack["t"])
}

func (c *current) onBaseType2() (interface{}, error) {
	return BasicAst{Type: "Type", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonBaseType2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBaseType2()
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
