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
								name: "ReturnTypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "ExternType",
			pos:  position{line: 53, col: 1, offset: 1584},
			expr: &choiceExpr{
				pos: position{line: 53, col: 14, offset: 1597},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 53, col: 14, offset: 1597},
						run: (*parser).callonExternType2,
						expr: &seqExpr{
							pos: position{line: 53, col: 14, offset: 1597},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 53, col: 14, offset: 1597},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 16, offset: 1599},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 25, offset: 1608},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 53, col: 29, offset: 1612},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 36, offset: 1619},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 39, offset: 1622},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 53, col: 44, offset: 1627},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 55, offset: 1638},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 57, offset: 1640},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 61, offset: 1644},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 5, offset: 1650},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 16, offset: 1661},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 30, offset: 1675},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 32, offset: 1677},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 36, offset: 1681},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 38, offset: 1683},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 44, offset: 1689},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 60, offset: 1705},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 62, offset: 1707},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 54, col: 67, offset: 1712},
										expr: &seqExpr{
											pos: position{line: 54, col: 68, offset: 1713},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 54, col: 68, offset: 1713},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 72, offset: 1717},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 74, offset: 1719},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 90, offset: 1735},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 54, col: 94, offset: 1739},
									expr: &litMatcher{
										pos:        position{line: 54, col: 94, offset: 1739},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 99, offset: 1744},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 101, offset: 1746},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 105, offset: 1750},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 73, col: 1, offset: 2299},
						run: (*parser).callonExternType34,
						expr: &seqExpr{
							pos: position{line: 73, col: 1, offset: 2299},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 73, col: 1, offset: 2299},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 3, offset: 2301},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 12, offset: 2310},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 73, col: 16, offset: 2314},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 23, offset: 2321},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 26, offset: 2324},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 73, col: 31, offset: 2329},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 42, offset: 2340},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 44, offset: 2342},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 48, offset: 2346},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 5, offset: 2352},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 74, col: 16, offset: 2363},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 30, offset: 2377},
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
			pos:  position{line: 82, col: 1, offset: 2571},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2582},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 82, col: 12, offset: 2582},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 82, col: 12, offset: 2582},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 82, col: 12, offset: 2582},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 14, offset: 2584},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 21, offset: 2591},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 24, offset: 2594},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 29, offset: 2599},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 40, offset: 2610},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 42, offset: 2612},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 49, offset: 2619},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 60, offset: 2630},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 62, offset: 2632},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 82, col: 66, offset: 2636},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 82, col: 72, offset: 2642},
										expr: &seqExpr{
											pos: position{line: 82, col: 73, offset: 2643},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 73, offset: 2643},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 76, offset: 2646},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 93, offset: 2663},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 100, col: 1, offset: 3150},
						run: (*parser).callonTypeDefn20,
						expr: &seqExpr{
							pos: position{line: 100, col: 1, offset: 3150},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 100, col: 1, offset: 3150},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 3, offset: 3152},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 10, offset: 3159},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 13, offset: 3162},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 18, offset: 3167},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 29, offset: 3178},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 31, offset: 3180},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 38, offset: 3187},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 49, offset: 3198},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 51, offset: 3200},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 55, offset: 3204},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 5, offset: 3210},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 9, offset: 3214},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 11, offset: 3216},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 17, offset: 3222},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 33, offset: 3238},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 35, offset: 3240},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 40, offset: 3245},
										expr: &seqExpr{
											pos: position{line: 101, col: 41, offset: 3246},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 101, col: 41, offset: 3246},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 45, offset: 3250},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 47, offset: 3252},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 63, offset: 3268},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 101, col: 67, offset: 3272},
									expr: &litMatcher{
										pos:        position{line: 101, col: 67, offset: 3272},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 72, offset: 3277},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 74, offset: 3279},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 78, offset: 3283},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 119, col: 1, offset: 3771},
						run: (*parser).callonTypeDefn50,
						expr: &seqExpr{
							pos: position{line: 119, col: 1, offset: 3771},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 119, col: 1, offset: 3771},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 3, offset: 3773},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 10, offset: 3780},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 13, offset: 3783},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 119, col: 18, offset: 3788},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 29, offset: 3799},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 31, offset: 3801},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 35, offset: 3805},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 5, offset: 3811},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 9, offset: 3815},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 11, offset: 3817},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 120, col: 17, offset: 3823},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 33, offset: 3839},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 35, offset: 3841},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 120, col: 40, offset: 3846},
										expr: &seqExpr{
											pos: position{line: 120, col: 41, offset: 3847},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 120, col: 41, offset: 3847},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 45, offset: 3851},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 47, offset: 3853},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 63, offset: 3869},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 120, col: 67, offset: 3873},
									expr: &litMatcher{
										pos:        position{line: 120, col: 67, offset: 3873},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 72, offset: 3878},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 74, offset: 3880},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 78, offset: 3884},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 138, col: 1, offset: 4388},
						run: (*parser).callonTypeDefn77,
						expr: &seqExpr{
							pos: position{line: 138, col: 1, offset: 4388},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 138, col: 1, offset: 4388},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 3, offset: 4390},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 10, offset: 4397},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 13, offset: 4400},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 18, offset: 4405},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 29, offset: 4416},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 31, offset: 4418},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 38, offset: 4425},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 49, offset: 4436},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 51, offset: 4438},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 55, offset: 4442},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 57, offset: 4444},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 138, col: 62, offset: 4449},
										expr: &ruleRefExpr{
											pos:  position{line: 138, col: 63, offset: 4450},
											name: "VariantConstructor",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 84, offset: 4471},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 152, col: 1, offset: 4864},
						run: (*parser).callonTypeDefn94,
						expr: &seqExpr{
							pos: position{line: 152, col: 1, offset: 4864},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 152, col: 1, offset: 4864},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 152, col: 3, offset: 4866},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 10, offset: 4873},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 152, col: 13, offset: 4876},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 152, col: 18, offset: 4881},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 29, offset: 4892},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 152, col: 31, offset: 4894},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 35, offset: 4898},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 152, col: 37, offset: 4900},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 152, col: 42, offset: 4905},
										expr: &ruleRefExpr{
											pos:  position{line: 152, col: 43, offset: 4906},
											name: "VariantConstructor",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 64, offset: 4927},
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
			pos:  position{line: 167, col: 1, offset: 5315},
			expr: &actionExpr{
				pos: position{line: 167, col: 14, offset: 5328},
				run: (*parser).callonTypeParams1,
				expr: &seqExpr{
					pos: position{line: 167, col: 14, offset: 5328},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 167, col: 14, offset: 5328},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 167, col: 18, offset: 5332},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 24, offset: 5338},
								name: "TypeParameter",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 167, col: 38, offset: 5352},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 167, col: 40, offset: 5354},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 167, col: 45, offset: 5359},
								expr: &seqExpr{
									pos: position{line: 167, col: 46, offset: 5360},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 167, col: 46, offset: 5360},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 50, offset: 5364},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 52, offset: 5366},
											name: "TypeParameter",
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 66, offset: 5380},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 167, col: 70, offset: 5384},
							expr: &litMatcher{
								pos:        position{line: 167, col: 70, offset: 5384},
								val:        ",",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 167, col: 75, offset: 5389},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 183, col: 1, offset: 5800},
			expr: &actionExpr{
				pos: position{line: 183, col: 19, offset: 5818},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 183, col: 19, offset: 5818},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 19, offset: 5818},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 24, offset: 5823},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 37, offset: 5836},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 183, col: 39, offset: 5838},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 43, offset: 5842},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 183, col: 45, offset: 5844},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 48, offset: 5847},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 187, col: 1, offset: 5948},
			expr: &choiceExpr{
				pos: position{line: 187, col: 22, offset: 5969},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 187, col: 22, offset: 5969},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 187, col: 22, offset: 5969},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 187, col: 22, offset: 5969},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 24, offset: 5971},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 28, offset: 5975},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 30, offset: 5977},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 35, offset: 5982},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 46, offset: 5993},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 5, offset: 6000},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 9, offset: 6004},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 11, offset: 6006},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 17, offset: 6012},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 33, offset: 6028},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 35, offset: 6030},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 188, col: 40, offset: 6035},
										expr: &seqExpr{
											pos: position{line: 188, col: 41, offset: 6036},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 188, col: 41, offset: 6036},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 188, col: 45, offset: 6040},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 188, col: 47, offset: 6042},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 188, col: 63, offset: 6058},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 188, col: 67, offset: 6062},
									expr: &litMatcher{
										pos:        position{line: 188, col: 67, offset: 6062},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 72, offset: 6067},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 74, offset: 6069},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 207, col: 1, offset: 6675},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 207, col: 1, offset: 6675},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 207, col: 1, offset: 6675},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 207, col: 3, offset: 6677},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 207, col: 7, offset: 6681},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 207, col: 9, offset: 6683},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 207, col: 14, offset: 6688},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 207, col: 25, offset: 6699},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 207, col: 30, offset: 6704},
										expr: &seqExpr{
											pos: position{line: 207, col: 31, offset: 6705},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 207, col: 31, offset: 6705},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 207, col: 34, offset: 6708},
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
			pos:  position{line: 224, col: 1, offset: 7160},
			expr: &choiceExpr{
				pos: position{line: 224, col: 14, offset: 7173},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 224, col: 14, offset: 7173},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 224, col: 14, offset: 7173},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 224, col: 14, offset: 7173},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 224, col: 16, offset: 7175},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 22, offset: 7181},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 26, offset: 7185},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 28, offset: 7187},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 39, offset: 7198},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 224, col: 42, offset: 7201},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 46, offset: 7205},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 49, offset: 7208},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 54, offset: 7213},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 59, offset: 7218},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 230, col: 1, offset: 7337},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 230, col: 1, offset: 7337},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 230, col: 1, offset: 7337},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 230, col: 3, offset: 7339},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 230, col: 9, offset: 7345},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 230, col: 13, offset: 7349},
									expr: &ruleRefExpr{
										pos:  position{line: 230, col: 14, offset: 7350},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 234, col: 1, offset: 7448},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 234, col: 1, offset: 7448},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 234, col: 1, offset: 7448},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 234, col: 3, offset: 7450},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 234, col: 9, offset: 7456},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 234, col: 13, offset: 7460},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 234, col: 15, offset: 7462},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 234, col: 26, offset: 7473},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 234, col: 29, offset: 7476},
									expr: &litMatcher{
										pos:        position{line: 234, col: 30, offset: 7477},
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
			pos:  position{line: 238, col: 1, offset: 7571},
			expr: &actionExpr{
				pos: position{line: 238, col: 12, offset: 7582},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 238, col: 12, offset: 7582},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 238, col: 12, offset: 7582},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 238, col: 14, offset: 7584},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 20, offset: 7590},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 238, col: 24, offset: 7594},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 238, col: 26, offset: 7596},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 39, offset: 7609},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 238, col: 42, offset: 7612},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 46, offset: 7616},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 238, col: 49, offset: 7619},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 238, col: 53, offset: 7623},
								expr: &seqExpr{
									pos: position{line: 238, col: 54, offset: 7624},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 238, col: 54, offset: 7624},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 63, offset: 7633},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 238, col: 67, offset: 7637},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 238, col: 71, offset: 7641},
								expr: &seqExpr{
									pos: position{line: 238, col: 72, offset: 7642},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 238, col: 72, offset: 7642},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 238, col: 74, offset: 7644},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 79, offset: 7649},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 81, offset: 7651},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 238, col: 96, offset: 7666},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 238, col: 100, offset: 7670},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 104, offset: 7674},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 238, col: 107, offset: 7677},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 238, col: 118, offset: 7688},
								expr: &ruleRefExpr{
									pos:  position{line: 238, col: 119, offset: 7689},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 131, offset: 7701},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 238, col: 133, offset: 7703},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 238, col: 137, offset: 7707},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 264, col: 1, offset: 8302},
			expr: &actionExpr{
				pos: position{line: 264, col: 8, offset: 8309},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 264, col: 8, offset: 8309},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 264, col: 12, offset: 8313},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 264, col: 12, offset: 8313},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 264, col: 21, offset: 8322},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 270, col: 1, offset: 8439},
			expr: &choiceExpr{
				pos: position{line: 270, col: 10, offset: 8448},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 270, col: 10, offset: 8448},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 270, col: 10, offset: 8448},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 270, col: 10, offset: 8448},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 270, col: 12, offset: 8450},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 17, offset: 8455},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 21, offset: 8459},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 26, offset: 8464},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 36, offset: 8474},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 270, col: 39, offset: 8477},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 43, offset: 8481},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 45, offset: 8483},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 270, col: 51, offset: 8489},
										expr: &ruleRefExpr{
											pos:  position{line: 270, col: 52, offset: 8490},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 64, offset: 8502},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 270, col: 67, offset: 8505},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 71, offset: 8509},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 270, col: 74, offset: 8512},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 81, offset: 8519},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 270, col: 84, offset: 8522},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 88, offset: 8526},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 90, offset: 8528},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 270, col: 96, offset: 8534},
										expr: &ruleRefExpr{
											pos:  position{line: 270, col: 97, offset: 8535},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 109, offset: 8547},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 270, col: 112, offset: 8550},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 289, col: 1, offset: 9053},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 289, col: 1, offset: 9053},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 289, col: 1, offset: 9053},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 289, col: 3, offset: 9055},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 8, offset: 9060},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 289, col: 12, offset: 9064},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 289, col: 17, offset: 9069},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 27, offset: 9079},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 289, col: 30, offset: 9082},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 34, offset: 9086},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 289, col: 36, offset: 9088},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 289, col: 42, offset: 9094},
										expr: &ruleRefExpr{
											pos:  position{line: 289, col: 43, offset: 9095},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 55, offset: 9107},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 289, col: 57, offset: 9109},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 289, col: 61, offset: 9113},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 289, col: 64, offset: 9116},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 289, col: 71, offset: 9123},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 289, col: 79, offset: 9131},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 301, col: 1, offset: 9461},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 301, col: 1, offset: 9461},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 301, col: 1, offset: 9461},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 301, col: 3, offset: 9463},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 8, offset: 9468},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 301, col: 12, offset: 9472},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 301, col: 17, offset: 9477},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 27, offset: 9487},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 301, col: 30, offset: 9490},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 34, offset: 9494},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 301, col: 36, offset: 9496},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 301, col: 42, offset: 9502},
										expr: &ruleRefExpr{
											pos:  position{line: 301, col: 43, offset: 9503},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 55, offset: 9515},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 301, col: 58, offset: 9518},
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
			pos:  position{line: 313, col: 1, offset: 9816},
			expr: &choiceExpr{
				pos: position{line: 313, col: 8, offset: 9823},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 313, col: 8, offset: 9823},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 313, col: 8, offset: 9823},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 313, col: 8, offset: 9823},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 313, col: 10, offset: 9825},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 17, offset: 9832},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 313, col: 28, offset: 9843},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 313, col: 32, offset: 9847},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 35, offset: 9850},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 313, col: 48, offset: 9863},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 53, offset: 9868},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 327, col: 1, offset: 10192},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 327, col: 1, offset: 10192},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 327, col: 1, offset: 10192},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 327, col: 3, offset: 10194},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 327, col: 6, offset: 10197},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 327, col: 19, offset: 10210},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 327, col: 24, offset: 10215},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 341, col: 1, offset: 10532},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 341, col: 1, offset: 10532},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 341, col: 1, offset: 10532},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 341, col: 3, offset: 10534},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 341, col: 6, offset: 10537},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 341, col: 19, offset: 10550},
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
			pos:  position{line: 348, col: 1, offset: 10721},
			expr: &actionExpr{
				pos: position{line: 348, col: 16, offset: 10736},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 348, col: 16, offset: 10736},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 348, col: 16, offset: 10736},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 23, offset: 10743},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 348, col: 36, offset: 10756},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 348, col: 41, offset: 10761},
								expr: &seqExpr{
									pos: position{line: 348, col: 42, offset: 10762},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 348, col: 42, offset: 10762},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 46, offset: 10766},
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
			pos:  position{line: 365, col: 1, offset: 11203},
			expr: &actionExpr{
				pos: position{line: 365, col: 15, offset: 11217},
				run: (*parser).callonArrayAccess1,
				expr: &seqExpr{
					pos: position{line: 365, col: 15, offset: 11217},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 365, col: 15, offset: 11217},
							label: "array",
							expr: &ruleRefExpr{
								pos:  position{line: 365, col: 21, offset: 11223},
								name: "VariableName",
							},
						},
						&litMatcher{
							pos:        position{line: 365, col: 34, offset: 11236},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 365, col: 38, offset: 11240},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 365, col: 40, offset: 11242},
								name: "Expr",
							},
						},
						&litMatcher{
							pos:        position{line: 365, col: 45, offset: 11247},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgsDefn",
			pos:  position{line: 369, col: 1, offset: 11333},
			expr: &choiceExpr{
				pos: position{line: 369, col: 12, offset: 11344},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 369, col: 12, offset: 11344},
						run: (*parser).callonArgsDefn2,
						expr: &seqExpr{
							pos: position{line: 369, col: 12, offset: 11344},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 369, col: 12, offset: 11344},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 369, col: 16, offset: 11348},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 369, col: 18, offset: 11350},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 369, col: 27, offset: 11359},
										name: "ArgDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 369, col: 35, offset: 11367},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 369, col: 37, offset: 11369},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 369, col: 42, offset: 11374},
										expr: &seqExpr{
											pos: position{line: 369, col: 43, offset: 11375},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 369, col: 43, offset: 11375},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 369, col: 47, offset: 11379},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 369, col: 49, offset: 11381},
													name: "ArgDefn",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 369, col: 59, offset: 11391},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 369, col: 61, offset: 11393},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 387, col: 1, offset: 11816},
						run: (*parser).callonArgsDefn17,
						expr: &seqExpr{
							pos: position{line: 387, col: 1, offset: 11816},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 387, col: 1, offset: 11816},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 387, col: 5, offset: 11820},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 387, col: 7, offset: 11822},
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
			name: "ArgDefn",
			pos:  position{line: 391, col: 1, offset: 11876},
			expr: &actionExpr{
				pos: position{line: 391, col: 11, offset: 11886},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 391, col: 11, offset: 11886},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 391, col: 11, offset: 11886},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 391, col: 16, offset: 11891},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 391, col: 27, offset: 11902},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 391, col: 29, offset: 11904},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 391, col: 34, offset: 11909},
								expr: &seqExpr{
									pos: position{line: 391, col: 35, offset: 11910},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 391, col: 35, offset: 11910},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 391, col: 39, offset: 11914},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 391, col: 41, offset: 11916},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 391, col: 59, offset: 11934},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ReturnTypeAnnotation",
			pos:  position{line: 412, col: 1, offset: 12473},
			expr: &choiceExpr{
				pos: position{line: 412, col: 24, offset: 12496},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 412, col: 24, offset: 12496},
						name: "TypeAnnotation",
					},
					&actionExpr{
						pos: position{line: 413, col: 1, offset: 12514},
						run: (*parser).callonReturnTypeAnnotation3,
						expr: &seqExpr{
							pos: position{line: 413, col: 1, offset: 12514},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 413, col: 1, offset: 12514},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 413, col: 5, offset: 12518},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 413, col: 7, offset: 12520},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 413, col: 9, offset: 12522},
										name: "TypeAnnotation",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 413, col: 24, offset: 12537},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 413, col: 26, offset: 12539},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 413, col: 31, offset: 12544},
										expr: &seqExpr{
											pos: position{line: 413, col: 32, offset: 12545},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 413, col: 32, offset: 12545},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 413, col: 36, offset: 12549},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 413, col: 38, offset: 12551},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 413, col: 55, offset: 12568},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 413, col: 57, offset: 12570},
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
			name: "TypeAnnotation",
			pos:  position{line: 432, col: 1, offset: 12969},
			expr: &choiceExpr{
				pos: position{line: 432, col: 18, offset: 12986},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 432, col: 18, offset: 12986},
						name: "ArrayType",
					},
					&ruleRefExpr{
						pos:  position{line: 432, col: 30, offset: 12998},
						name: "AnyType",
					},
					&actionExpr{
						pos: position{line: 433, col: 1, offset: 13009},
						run: (*parser).callonTypeAnnotation4,
						expr: &seqExpr{
							pos: position{line: 433, col: 1, offset: 13009},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 433, col: 1, offset: 13009},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 433, col: 8, offset: 13016},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 433, col: 11, offset: 13019},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 433, col: 16, offset: 13024},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 433, col: 25, offset: 13033},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 433, col: 27, offset: 13035},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 433, col: 32, offset: 13040},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 433, col: 34, offset: 13042},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 433, col: 38, offset: 13046},
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
			pos:  position{line: 442, col: 1, offset: 13262},
			expr: &choiceExpr{
				pos: position{line: 442, col: 11, offset: 13272},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 442, col: 11, offset: 13272},
						name: "ModuleName",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 24, offset: 13285},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 35, offset: 13296},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 444, col: 1, offset: 13311},
			expr: &choiceExpr{
				pos: position{line: 444, col: 13, offset: 13323},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 444, col: 13, offset: 13323},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 444, col: 13, offset: 13323},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 444, col: 13, offset: 13323},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 444, col: 17, offset: 13327},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 444, col: 19, offset: 13329},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 444, col: 28, offset: 13338},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 444, col: 40, offset: 13350},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 444, col: 42, offset: 13352},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 444, col: 47, offset: 13357},
										expr: &seqExpr{
											pos: position{line: 444, col: 48, offset: 13358},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 444, col: 48, offset: 13358},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 444, col: 52, offset: 13362},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 444, col: 54, offset: 13364},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 444, col: 68, offset: 13378},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 444, col: 70, offset: 13380},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 461, col: 1, offset: 13802},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 461, col: 1, offset: 13802},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 461, col: 1, offset: 13802},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 461, col: 5, offset: 13806},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 461, col: 7, offset: 13808},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 461, col: 16, offset: 13817},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 461, col: 21, offset: 13822},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 461, col: 23, offset: 13824},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 466, col: 1, offset: 13930},
						run: (*parser).callonArguments25,
						expr: &seqExpr{
							pos: position{line: 466, col: 1, offset: 13930},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 466, col: 1, offset: 13930},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 466, col: 5, offset: 13934},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 466, col: 7, offset: 13936},
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
			pos:  position{line: 470, col: 1, offset: 13990},
			expr: &actionExpr{
				pos: position{line: 470, col: 16, offset: 14005},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 470, col: 16, offset: 14005},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 470, col: 16, offset: 14005},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 470, col: 18, offset: 14007},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 470, col: 21, offset: 14010},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 470, col: 27, offset: 14016},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 470, col: 32, offset: 14021},
								expr: &seqExpr{
									pos: position{line: 470, col: 33, offset: 14022},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 470, col: 33, offset: 14022},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 470, col: 37, offset: 14026},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 470, col: 46, offset: 14035},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 470, col: 50, offset: 14039},
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
			pos:  position{line: 490, col: 1, offset: 14645},
			expr: &choiceExpr{
				pos: position{line: 490, col: 9, offset: 14653},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 490, col: 9, offset: 14653},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 490, col: 21, offset: 14665},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 490, col: 37, offset: 14681},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 490, col: 48, offset: 14692},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 490, col: 60, offset: 14704},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 492, col: 1, offset: 14717},
			expr: &actionExpr{
				pos: position{line: 492, col: 13, offset: 14729},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 492, col: 13, offset: 14729},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 492, col: 13, offset: 14729},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 492, col: 15, offset: 14731},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 492, col: 21, offset: 14737},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 492, col: 35, offset: 14751},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 492, col: 40, offset: 14756},
								expr: &seqExpr{
									pos: position{line: 492, col: 41, offset: 14757},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 492, col: 41, offset: 14757},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 45, offset: 14761},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 61, offset: 14777},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 65, offset: 14781},
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
			pos:  position{line: 525, col: 1, offset: 15674},
			expr: &actionExpr{
				pos: position{line: 525, col: 17, offset: 15690},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 525, col: 17, offset: 15690},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 525, col: 17, offset: 15690},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 525, col: 19, offset: 15692},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 525, col: 25, offset: 15698},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 525, col: 34, offset: 15707},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 525, col: 39, offset: 15712},
								expr: &seqExpr{
									pos: position{line: 525, col: 40, offset: 15713},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 525, col: 40, offset: 15713},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 525, col: 44, offset: 15717},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 525, col: 61, offset: 15734},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 525, col: 65, offset: 15738},
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
			pos:  position{line: 557, col: 1, offset: 16625},
			expr: &actionExpr{
				pos: position{line: 557, col: 12, offset: 16636},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 557, col: 12, offset: 16636},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 557, col: 12, offset: 16636},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 557, col: 14, offset: 16638},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 557, col: 20, offset: 16644},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 557, col: 30, offset: 16654},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 557, col: 35, offset: 16659},
								expr: &seqExpr{
									pos: position{line: 557, col: 36, offset: 16660},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 557, col: 36, offset: 16660},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 557, col: 40, offset: 16664},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 557, col: 52, offset: 16676},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 557, col: 56, offset: 16680},
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
			pos:  position{line: 589, col: 1, offset: 17568},
			expr: &actionExpr{
				pos: position{line: 589, col: 13, offset: 17580},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 589, col: 13, offset: 17580},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 589, col: 13, offset: 17580},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 589, col: 15, offset: 17582},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 589, col: 21, offset: 17588},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 589, col: 33, offset: 17600},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 589, col: 38, offset: 17605},
								expr: &seqExpr{
									pos: position{line: 589, col: 39, offset: 17606},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 589, col: 39, offset: 17606},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 589, col: 43, offset: 17610},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 589, col: 56, offset: 17623},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 589, col: 60, offset: 17627},
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
			pos:  position{line: 620, col: 1, offset: 18516},
			expr: &choiceExpr{
				pos: position{line: 620, col: 15, offset: 18530},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 620, col: 15, offset: 18530},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 620, col: 15, offset: 18530},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 620, col: 15, offset: 18530},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 620, col: 17, offset: 18532},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 620, col: 21, offset: 18536},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 620, col: 24, offset: 18539},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 620, col: 30, offset: 18545},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 620, col: 36, offset: 18551},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 620, col: 39, offset: 18554},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 623, col: 5, offset: 18677},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 625, col: 1, offset: 18684},
			expr: &choiceExpr{
				pos: position{line: 625, col: 12, offset: 18695},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 625, col: 12, offset: 18695},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 625, col: 30, offset: 18713},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 625, col: 49, offset: 18732},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 625, col: 64, offset: 18747},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 627, col: 1, offset: 18760},
			expr: &actionExpr{
				pos: position{line: 627, col: 19, offset: 18778},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 627, col: 21, offset: 18780},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 627, col: 21, offset: 18780},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 627, col: 28, offset: 18787},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 631, col: 1, offset: 18869},
			expr: &actionExpr{
				pos: position{line: 631, col: 20, offset: 18888},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 631, col: 22, offset: 18890},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 631, col: 22, offset: 18890},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 631, col: 29, offset: 18897},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 631, col: 36, offset: 18904},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 631, col: 42, offset: 18910},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 631, col: 48, offset: 18916},
							val:        "!=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 631, col: 55, offset: 18923},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 635, col: 1, offset: 19002},
			expr: &choiceExpr{
				pos: position{line: 635, col: 16, offset: 19017},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 635, col: 16, offset: 19017},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 635, col: 18, offset: 19019},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 635, col: 18, offset: 19019},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 635, col: 25, offset: 19026},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 638, col: 3, offset: 19109},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 638, col: 5, offset: 19111},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 638, col: 5, offset: 19111},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 638, col: 11, offset: 19117},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 641, col: 3, offset: 19197},
						run: (*parser).callonOperatorHigh10,
						expr: &litMatcher{
							pos:        position{line: 641, col: 5, offset: 19199},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 644, col: 3, offset: 19279},
						run: (*parser).callonOperatorHigh12,
						expr: &litMatcher{
							pos:        position{line: 644, col: 3, offset: 19279},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 648, col: 1, offset: 19360},
			expr: &choiceExpr{
				pos: position{line: 648, col: 15, offset: 19374},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 648, col: 15, offset: 19374},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 648, col: 17, offset: 19376},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 648, col: 17, offset: 19376},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 648, col: 24, offset: 19383},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 651, col: 3, offset: 19467},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 651, col: 5, offset: 19469},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 651, col: 5, offset: 19469},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 651, col: 11, offset: 19475},
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
			pos:  position{line: 655, col: 1, offset: 19554},
			expr: &choiceExpr{
				pos: position{line: 655, col: 9, offset: 19562},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 655, col: 9, offset: 19562},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 655, col: 16, offset: 19569},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 655, col: 31, offset: 19584},
						name: "ArrayAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 655, col: 45, offset: 19598},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 655, col: 60, offset: 19613},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 657, col: 1, offset: 19620},
			expr: &choiceExpr{
				pos: position{line: 657, col: 14, offset: 19633},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 657, col: 14, offset: 19633},
						run: (*parser).callonAssignable2,
						expr: &seqExpr{
							pos: position{line: 657, col: 14, offset: 19633},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 657, col: 14, offset: 19633},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 657, col: 20, offset: 19639},
										name: "SubAssignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 657, col: 34, offset: 19653},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 657, col: 36, offset: 19655},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 657, col: 41, offset: 19660},
										expr: &seqExpr{
											pos: position{line: 657, col: 42, offset: 19661},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 657, col: 42, offset: 19661},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 657, col: 46, offset: 19665},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 657, col: 48, offset: 19667},
													name: "SubAssignable",
												},
												&ruleRefExpr{
													pos:  position{line: 657, col: 62, offset: 19681},
													name: "_",
												},
											},
										},
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 672, col: 3, offset: 20098},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 672, col: 18, offset: 20113},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "SubAssignable",
			pos:  position{line: 674, col: 1, offset: 20121},
			expr: &choiceExpr{
				pos: position{line: 674, col: 17, offset: 20137},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 674, col: 17, offset: 20137},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 674, col: 32, offset: 20152},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 676, col: 1, offset: 20160},
			expr: &choiceExpr{
				pos: position{line: 676, col: 14, offset: 20173},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 676, col: 14, offset: 20173},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 676, col: 29, offset: 20188},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 678, col: 1, offset: 20200},
			expr: &actionExpr{
				pos: position{line: 678, col: 16, offset: 20215},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 678, col: 16, offset: 20215},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 678, col: 16, offset: 20215},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 678, col: 20, offset: 20219},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 678, col: 22, offset: 20221},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 678, col: 28, offset: 20227},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 678, col: 33, offset: 20232},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 678, col: 35, offset: 20234},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 678, col: 40, offset: 20239},
								expr: &seqExpr{
									pos: position{line: 678, col: 41, offset: 20240},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 678, col: 41, offset: 20240},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 678, col: 45, offset: 20244},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 678, col: 47, offset: 20246},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 678, col: 52, offset: 20251},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 678, col: 56, offset: 20255},
							expr: &litMatcher{
								pos:        position{line: 678, col: 56, offset: 20255},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 678, col: 61, offset: 20260},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 678, col: 63, offset: 20262},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 694, col: 1, offset: 20703},
			expr: &actionExpr{
				pos: position{line: 694, col: 19, offset: 20721},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 694, col: 19, offset: 20721},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 694, col: 19, offset: 20721},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 694, col: 24, offset: 20726},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 694, col: 35, offset: 20737},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 694, col: 37, offset: 20739},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 694, col: 42, offset: 20744},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 707, col: 1, offset: 21016},
			expr: &actionExpr{
				pos: position{line: 707, col: 18, offset: 21033},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 707, col: 18, offset: 21033},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 707, col: 18, offset: 21033},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 707, col: 23, offset: 21038},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 707, col: 34, offset: 21049},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 707, col: 36, offset: 21051},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 707, col: 40, offset: 21055},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 707, col: 42, offset: 21057},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 707, col: 52, offset: 21067},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 707, col: 65, offset: 21080},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 707, col: 67, offset: 21082},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 707, col: 71, offset: 21086},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 707, col: 73, offset: 21088},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 707, col: 84, offset: 21099},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 707, col: 89, offset: 21104},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 707, col: 94, offset: 21109},
								expr: &seqExpr{
									pos: position{line: 707, col: 95, offset: 21110},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 707, col: 95, offset: 21110},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 707, col: 99, offset: 21114},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 707, col: 101, offset: 21116},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 707, col: 114, offset: 21129},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 707, col: 116, offset: 21131},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 707, col: 120, offset: 21135},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 707, col: 122, offset: 21137},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 707, col: 130, offset: 21145},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 727, col: 1, offset: 21735},
			expr: &actionExpr{
				pos: position{line: 727, col: 17, offset: 21751},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 727, col: 17, offset: 21751},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 727, col: 17, offset: 21751},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 727, col: 22, offset: 21756},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 731, col: 1, offset: 21829},
			expr: &actionExpr{
				pos: position{line: 731, col: 16, offset: 21844},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 731, col: 16, offset: 21844},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 731, col: 16, offset: 21844},
							expr: &ruleRefExpr{
								pos:  position{line: 731, col: 17, offset: 21845},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 731, col: 27, offset: 21855},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 731, col: 27, offset: 21855},
									expr: &charClassMatcher{
										pos:        position{line: 731, col: 27, offset: 21855},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 731, col: 34, offset: 21862},
									expr: &charClassMatcher{
										pos:        position{line: 731, col: 34, offset: 21862},
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
			pos:  position{line: 735, col: 1, offset: 21937},
			expr: &actionExpr{
				pos: position{line: 735, col: 14, offset: 21950},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 735, col: 15, offset: 21951},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 735, col: 15, offset: 21951},
							expr: &charClassMatcher{
								pos:        position{line: 735, col: 15, offset: 21951},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 735, col: 22, offset: 21958},
							expr: &charClassMatcher{
								pos:        position{line: 735, col: 22, offset: 21958},
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
			pos:  position{line: 739, col: 1, offset: 22033},
			expr: &choiceExpr{
				pos: position{line: 739, col: 9, offset: 22041},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 739, col: 9, offset: 22041},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 739, col: 9, offset: 22041},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 739, col: 9, offset: 22041},
									expr: &litMatcher{
										pos:        position{line: 739, col: 9, offset: 22041},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 739, col: 14, offset: 22046},
									expr: &charClassMatcher{
										pos:        position{line: 739, col: 14, offset: 22046},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 739, col: 21, offset: 22053},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 739, col: 25, offset: 22057},
									expr: &charClassMatcher{
										pos:        position{line: 739, col: 25, offset: 22057},
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
						pos: position{line: 746, col: 3, offset: 22246},
						run: (*parser).callonConst11,
						expr: &seqExpr{
							pos: position{line: 746, col: 3, offset: 22246},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 746, col: 3, offset: 22246},
									expr: &litMatcher{
										pos:        position{line: 746, col: 3, offset: 22246},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 746, col: 8, offset: 22251},
									expr: &charClassMatcher{
										pos:        position{line: 746, col: 8, offset: 22251},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 746, col: 15, offset: 22258},
									expr: &litMatcher{
										pos:        position{line: 746, col: 16, offset: 22259},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 753, col: 3, offset: 22435},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 753, col: 3, offset: 22435},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 757, col: 3, offset: 22520},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 757, col: 3, offset: 22520},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 760, col: 3, offset: 22606},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 761, col: 3, offset: 22613},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 762, col: 3, offset: 22629},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 762, col: 3, offset: 22629},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 762, col: 3, offset: 22629},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 762, col: 7, offset: 22633},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 762, col: 12, offset: 22638},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 762, col: 12, offset: 22638},
												expr: &ruleRefExpr{
													pos:  position{line: 762, col: 13, offset: 22639},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 762, col: 25, offset: 22651,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 762, col: 28, offset: 22654},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 764, col: 5, offset: 22746},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 764, col: 20, offset: 22761},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 764, col: 37, offset: 22778},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 766, col: 1, offset: 22795},
			expr: &actionExpr{
				pos: position{line: 766, col: 8, offset: 22802},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 766, col: 8, offset: 22802},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 770, col: 1, offset: 22865},
			expr: &actionExpr{
				pos: position{line: 770, col: 17, offset: 22881},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 770, col: 17, offset: 22881},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 770, col: 17, offset: 22881},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 770, col: 21, offset: 22885},
							expr: &charClassMatcher{
								pos:        position{line: 770, col: 22, offset: 22886},
								val:        "[^\\r\\n\"]",
								chars:      []rune{'\r', '\n', '"'},
								ignoreCase: false,
								inverted:   true,
							},
						},
						&litMatcher{
							pos:        position{line: 770, col: 33, offset: 22897},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 778, col: 1, offset: 23057},
			expr: &actionExpr{
				pos: position{line: 778, col: 10, offset: 23066},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 778, col: 11, offset: 23067},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 782, col: 1, offset: 23122},
			expr: &seqExpr{
				pos: position{line: 782, col: 12, offset: 23133},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 782, col: 13, offset: 23134},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 782, col: 13, offset: 23134},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 21, offset: 23142},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 28, offset: 23149},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 37, offset: 23158},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 48, offset: 23169},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 57, offset: 23178},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 66, offset: 23187},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 76, offset: 23197},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 782, col: 88, offset: 23209},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 782, col: 97, offset: 23218},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 782, col: 107, offset: 23228},
						expr: &oneOrMoreExpr{
							pos: position{line: 782, col: 108, offset: 23229},
							expr: &charClassMatcher{
								pos:        position{line: 782, col: 108, offset: 23229},
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
			pos:  position{line: 784, col: 1, offset: 23237},
			expr: &actionExpr{
				pos: position{line: 784, col: 13, offset: 23249},
				run: (*parser).callonArrayType1,
				expr: &seqExpr{
					pos: position{line: 784, col: 13, offset: 23249},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 784, col: 13, offset: 23249},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 784, col: 17, offset: 23253},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 784, col: 19, offset: 23255},
								name: "TypeAnnotation",
							},
						},
						&litMatcher{
							pos:        position{line: 784, col: 34, offset: 23270},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BaseType",
			pos:  position{line: 788, col: 1, offset: 23323},
			expr: &choiceExpr{
				pos: position{line: 788, col: 12, offset: 23334},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 788, col: 12, offset: 23334},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 788, col: 14, offset: 23336},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 788, col: 14, offset: 23336},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 788, col: 24, offset: 23346},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 788, col: 33, offset: 23355},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 788, col: 44, offset: 23366},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 788, col: 53, offset: 23375},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 788, col: 62, offset: 23384},
									val:        "float64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 788, col: 74, offset: 23396},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 791, col: 3, offset: 23494},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 793, col: 1, offset: 23500},
			expr: &charClassMatcher{
				pos:        position{line: 793, col: 15, offset: 23514},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 795, col: 1, offset: 23530},
			expr: &choiceExpr{
				pos: position{line: 795, col: 18, offset: 23547},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 795, col: 18, offset: 23547},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 795, col: 37, offset: 23566},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 797, col: 1, offset: 23581},
			expr: &charClassMatcher{
				pos:        position{line: 797, col: 20, offset: 23600},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 799, col: 1, offset: 23613},
			expr: &charClassMatcher{
				pos:        position{line: 799, col: 16, offset: 23628},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 801, col: 1, offset: 23635},
			expr: &charClassMatcher{
				pos:        position{line: 801, col: 23, offset: 23657},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 803, col: 1, offset: 23664},
			expr: &charClassMatcher{
				pos:        position{line: 803, col: 12, offset: 23675},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 805, col: 1, offset: 23686},
			expr: &choiceExpr{
				pos: position{line: 805, col: 22, offset: 23707},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 805, col: 22, offset: 23707},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 805, col: 33, offset: 23718},
						expr: &charClassMatcher{
							pos:        position{line: 805, col: 33, offset: 23718},
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
			pos:         position{line: 807, col: 1, offset: 23730},
			expr: &choiceExpr{
				pos: position{line: 807, col: 21, offset: 23750},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 807, col: 21, offset: 23750},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 807, col: 32, offset: 23761},
						expr: &charClassMatcher{
							pos:        position{line: 807, col: 32, offset: 23761},
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
			pos:         position{line: 809, col: 1, offset: 23773},
			expr: &oneOrMoreExpr{
				pos: position{line: 809, col: 33, offset: 23805},
				expr: &charClassMatcher{
					pos:        position{line: 809, col: 33, offset: 23805},
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
			pos:         position{line: 811, col: 1, offset: 23813},
			expr: &zeroOrMoreExpr{
				pos: position{line: 811, col: 32, offset: 23844},
				expr: &charClassMatcher{
					pos:        position{line: 811, col: 32, offset: 23844},
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
			pos:         position{line: 813, col: 1, offset: 23852},
			expr: &choiceExpr{
				pos: position{line: 813, col: 15, offset: 23866},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 813, col: 15, offset: 23866},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 813, col: 26, offset: 23877},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 813, col: 26, offset: 23877},
								expr: &charClassMatcher{
									pos:        position{line: 813, col: 26, offset: 23877},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 813, col: 35, offset: 23886},
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
			pos:  position{line: 815, col: 1, offset: 23892},
			expr: &oneOrMoreExpr{
				pos: position{line: 815, col: 12, offset: 23903},
				expr: &ruleRefExpr{
					pos:  position{line: 815, col: 13, offset: 23904},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 817, col: 1, offset: 23915},
			expr: &choiceExpr{
				pos: position{line: 817, col: 11, offset: 23925},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 817, col: 11, offset: 23925},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 817, col: 11, offset: 23925},
								expr: &charClassMatcher{
									pos:        position{line: 817, col: 11, offset: 23925},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 817, col: 22, offset: 23936},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 817, col: 27, offset: 23941},
								expr: &seqExpr{
									pos: position{line: 817, col: 28, offset: 23942},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 817, col: 28, offset: 23942},
											expr: &charClassMatcher{
												pos:        position{line: 817, col: 29, offset: 23943},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 817, col: 34, offset: 23948,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 817, col: 38, offset: 23952},
								expr: &litMatcher{
									pos:        position{line: 817, col: 39, offset: 23953},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 817, col: 46, offset: 23960},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 817, col: 46, offset: 23960},
								expr: &charClassMatcher{
									pos:        position{line: 817, col: 46, offset: 23960},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 817, col: 57, offset: 23971},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 817, col: 62, offset: 23976},
								expr: &seqExpr{
									pos: position{line: 817, col: 63, offset: 23977},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 817, col: 63, offset: 23977},
											expr: &litMatcher{
												pos:        position{line: 817, col: 64, offset: 23978},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 817, col: 69, offset: 23983,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 817, col: 73, offset: 23987},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 817, col: 78, offset: 23992},
								expr: &charClassMatcher{
									pos:        position{line: 817, col: 78, offset: 23992},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 817, col: 84, offset: 23998},
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
			pos:  position{line: 819, col: 1, offset: 24004},
			expr: &notExpr{
				pos: position{line: 819, col: 7, offset: 24010},
				expr: &anyMatcher{
					line: 819, col: 8, offset: 24011,
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
		Arguments: args.(Container).Subvalues, ReturnAnnotation: ret.(Ast)}, nil
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

func (c *current) onArgsDefn2(argument, rest interface{}) (interface{}, error) {

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

func (p *parser) callonArgsDefn2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgsDefn2(stack["argument"], stack["rest"])
}

func (c *current) onArgsDefn17() (interface{}, error) {
	return Container{Type: "Arguments"}, nil
}

func (p *parser) callonArgsDefn17() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgsDefn17()
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

func (c *current) onReturnTypeAnnotation3(t, rest interface{}) (interface{}, error) {

	args := []Ast{t.(Ast)}

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

	return ReturnTuple{Subvalues: args}, nil
}

func (p *parser) callonReturnTypeAnnotation3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReturnTypeAnnotation3(stack["t"], stack["rest"])
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

func (c *current) onAssignable2(first, rest interface{}) (interface{}, error) {
	args := []Ast{first.(Ast)}

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
	return Container{Type: "Assignable", Subvalues: args}, nil
}

func (p *parser) callonAssignable2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignable2(stack["first"], stack["rest"])
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
