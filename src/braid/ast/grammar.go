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
			pos:  position{line: 28, col: 1, offset: 719},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 739},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 739},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 32, offset: 750},
						name: "TypeDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 43, offset: 761},
						name: "ExternDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 773},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 785},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 785},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 796},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 809},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 819},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 830},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 830},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 830},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 832},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 32, col: 19, offset: 837},
							name: "N",
						},
					},
				},
			},
		},
		{
			name: "ExternDefn",
			pos:  position{line: 46, col: 1, offset: 1195},
			expr: &actionExpr{
				pos: position{line: 46, col: 14, offset: 1208},
				run: (*parser).callonExternDefn1,
				expr: &seqExpr{
					pos: position{line: 46, col: 14, offset: 1208},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 14, offset: 1208},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 16, offset: 1210},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 25, offset: 1219},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 29, offset: 1223},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 34, offset: 1228},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 45, offset: 1239},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 48, offset: 1242},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 52, offset: 1246},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 55, offset: 1249},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 66, offset: 1260},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 80, offset: 1274},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 83, offset: 1277},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 88, offset: 1282},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 97, offset: 1291},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 100, offset: 1294},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 104, offset: 1298},
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 52, col: 1, offset: 1510},
			expr: &choiceExpr{
				pos: position{line: 52, col: 12, offset: 1521},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 52, col: 12, offset: 1521},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 52, col: 12, offset: 1521},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 12, offset: 1521},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 14, offset: 1523},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 21, offset: 1530},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 24, offset: 1533},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 29, offset: 1538},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 52, col: 40, offset: 1549},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 52, col: 47, offset: 1556},
										expr: &seqExpr{
											pos: position{line: 52, col: 48, offset: 1557},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 48, offset: 1557},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 51, offset: 1560},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 67, offset: 1576},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 69, offset: 1578},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 52, col: 73, offset: 1582},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 52, col: 79, offset: 1588},
										expr: &seqExpr{
											pos: position{line: 52, col: 80, offset: 1589},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 80, offset: 1589},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 83, offset: 1592},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 93, offset: 1602},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 71, col: 1, offset: 2098},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 71, col: 1, offset: 2098},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 71, col: 1, offset: 2098},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 71, col: 3, offset: 2100},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 10, offset: 2107},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 71, col: 13, offset: 2110},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 71, col: 18, offset: 2115},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 71, col: 29, offset: 2126},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 71, col: 36, offset: 2133},
										expr: &seqExpr{
											pos: position{line: 71, col: 37, offset: 2134},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 71, col: 37, offset: 2134},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 71, col: 40, offset: 2137},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 56, offset: 2153},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 71, col: 58, offset: 2155},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 62, offset: 2159},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 5, offset: 2165},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 9, offset: 2169},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 11, offset: 2171},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 72, col: 17, offset: 2177},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 33, offset: 2193},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 35, offset: 2195},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 72, col: 40, offset: 2200},
										expr: &seqExpr{
											pos: position{line: 72, col: 41, offset: 2201},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 72, col: 41, offset: 2201},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 45, offset: 2205},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 47, offset: 2207},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 63, offset: 2223},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 72, col: 67, offset: 2227},
									expr: &litMatcher{
										pos:        position{line: 72, col: 67, offset: 2227},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 72, offset: 2232},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 74, offset: 2234},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 78, offset: 2238},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 90, col: 1, offset: 2723},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 90, col: 1, offset: 2723},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 90, col: 1, offset: 2723},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 90, col: 3, offset: 2725},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 10, offset: 2732},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 90, col: 13, offset: 2735},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 90, col: 18, offset: 2740},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 90, col: 29, offset: 2751},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 90, col: 36, offset: 2758},
										expr: &seqExpr{
											pos: position{line: 90, col: 37, offset: 2759},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 90, col: 37, offset: 2759},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 90, col: 40, offset: 2762},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 56, offset: 2778},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 90, col: 58, offset: 2780},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 62, offset: 2784},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 90, col: 64, offset: 2786},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 90, col: 69, offset: 2791},
										expr: &ruleRefExpr{
											pos:  position{line: 90, col: 70, offset: 2792},
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
			pos:  position{line: 105, col: 1, offset: 3199},
			expr: &actionExpr{
				pos: position{line: 105, col: 19, offset: 3217},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 105, col: 19, offset: 3217},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 105, col: 19, offset: 3217},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 24, offset: 3222},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 37, offset: 3235},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 105, col: 39, offset: 3237},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 43, offset: 3241},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 105, col: 45, offset: 3243},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 48, offset: 3246},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 109, col: 1, offset: 3340},
			expr: &choiceExpr{
				pos: position{line: 109, col: 22, offset: 3361},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 109, col: 22, offset: 3361},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 109, col: 22, offset: 3361},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 109, col: 22, offset: 3361},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 26, offset: 3365},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 28, offset: 3367},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 109, col: 33, offset: 3372},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 44, offset: 3383},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 109, col: 46, offset: 3385},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 50, offset: 3389},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 52, offset: 3391},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 109, col: 58, offset: 3397},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 74, offset: 3413},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 76, offset: 3415},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 109, col: 81, offset: 3420},
										expr: &seqExpr{
											pos: position{line: 109, col: 82, offset: 3421},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 109, col: 82, offset: 3421},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 86, offset: 3425},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 88, offset: 3427},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 104, offset: 3443},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 109, col: 108, offset: 3447},
									expr: &litMatcher{
										pos:        position{line: 109, col: 108, offset: 3447},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 113, offset: 3452},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 109, col: 115, offset: 3454},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 119, offset: 3458},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 128, col: 1, offset: 4063},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 128, col: 1, offset: 4063},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 128, col: 1, offset: 4063},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 5, offset: 4067},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 128, col: 7, offset: 4069},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 128, col: 12, offset: 4074},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 128, col: 23, offset: 4085},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 128, col: 28, offset: 4090},
										expr: &seqExpr{
											pos: position{line: 128, col: 29, offset: 4091},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 128, col: 29, offset: 4091},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 128, col: 32, offset: 4094},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 42, offset: 4104},
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
			pos:  position{line: 145, col: 1, offset: 4541},
			expr: &choiceExpr{
				pos: position{line: 145, col: 11, offset: 4551},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 145, col: 11, offset: 4551},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 145, col: 22, offset: 4562},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 147, col: 1, offset: 4577},
			expr: &choiceExpr{
				pos: position{line: 147, col: 14, offset: 4590},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 147, col: 14, offset: 4590},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 147, col: 14, offset: 4590},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 147, col: 14, offset: 4590},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 147, col: 16, offset: 4592},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 22, offset: 4598},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 26, offset: 4602},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 28, offset: 4604},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 39, offset: 4615},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 147, col: 42, offset: 4618},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 46, offset: 4622},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 49, offset: 4625},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 54, offset: 4630},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 59, offset: 4635},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 153, col: 1, offset: 4754},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 153, col: 1, offset: 4754},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 153, col: 1, offset: 4754},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 153, col: 3, offset: 4756},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 153, col: 9, offset: 4762},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 153, col: 13, offset: 4766},
									expr: &ruleRefExpr{
										pos:  position{line: 153, col: 14, offset: 4767},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 157, col: 1, offset: 4865},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 157, col: 1, offset: 4865},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 157, col: 1, offset: 4865},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 157, col: 3, offset: 4867},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 9, offset: 4873},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 157, col: 13, offset: 4877},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 157, col: 15, offset: 4879},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 26, offset: 4890},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 157, col: 29, offset: 4893},
									expr: &litMatcher{
										pos:        position{line: 157, col: 30, offset: 4894},
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
			pos:  position{line: 161, col: 1, offset: 4988},
			expr: &actionExpr{
				pos: position{line: 161, col: 12, offset: 4999},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 161, col: 12, offset: 4999},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 161, col: 12, offset: 4999},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 14, offset: 5001},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 20, offset: 5007},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 24, offset: 5011},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 26, offset: 5013},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 39, offset: 5026},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 161, col: 42, offset: 5029},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 46, offset: 5033},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 49, offset: 5036},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 161, col: 53, offset: 5040},
								expr: &seqExpr{
									pos: position{line: 161, col: 54, offset: 5041},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 54, offset: 5041},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 63, offset: 5050},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 67, offset: 5054},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 161, col: 71, offset: 5058},
								expr: &seqExpr{
									pos: position{line: 161, col: 72, offset: 5059},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 72, offset: 5059},
											name: "AnyType",
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 80, offset: 5067},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 161, col: 84, offset: 5071},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 88, offset: 5075},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 91, offset: 5078},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 161, col: 102, offset: 5089},
								expr: &ruleRefExpr{
									pos:  position{line: 161, col: 103, offset: 5090},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 115, offset: 5102},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 117, offset: 5104},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 121, offset: 5108},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 187, col: 1, offset: 5718},
			expr: &actionExpr{
				pos: position{line: 187, col: 8, offset: 5725},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 187, col: 8, offset: 5725},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 187, col: 12, offset: 5729},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 187, col: 12, offset: 5729},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 187, col: 21, offset: 5738},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 193, col: 1, offset: 5855},
			expr: &choiceExpr{
				pos: position{line: 193, col: 10, offset: 5864},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 193, col: 10, offset: 5864},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 193, col: 10, offset: 5864},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 193, col: 10, offset: 5864},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 193, col: 12, offset: 5866},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 17, offset: 5871},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 21, offset: 5875},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 193, col: 26, offset: 5880},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 36, offset: 5890},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 193, col: 39, offset: 5893},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 43, offset: 5897},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 45, offset: 5899},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 51, offset: 5905},
										expr: &ruleRefExpr{
											pos:  position{line: 193, col: 52, offset: 5906},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 64, offset: 5918},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 193, col: 67, offset: 5921},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 71, offset: 5925},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 193, col: 74, offset: 5928},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 81, offset: 5935},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 193, col: 84, offset: 5938},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 88, offset: 5942},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 90, offset: 5944},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 96, offset: 5950},
										expr: &ruleRefExpr{
											pos:  position{line: 193, col: 97, offset: 5951},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 109, offset: 5963},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 193, col: 112, offset: 5966},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 212, col: 1, offset: 6469},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 212, col: 1, offset: 6469},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 212, col: 1, offset: 6469},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 3, offset: 6471},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 8, offset: 6476},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 12, offset: 6480},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 17, offset: 6485},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 27, offset: 6495},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 212, col: 30, offset: 6498},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 34, offset: 6502},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 36, offset: 6504},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 42, offset: 6510},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 43, offset: 6511},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 55, offset: 6523},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 57, offset: 6525},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 61, offset: 6529},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 212, col: 64, offset: 6532},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 212, col: 71, offset: 6539},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 79, offset: 6547},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 224, col: 1, offset: 6877},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 224, col: 1, offset: 6877},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 224, col: 1, offset: 6877},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 224, col: 3, offset: 6879},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 8, offset: 6884},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 12, offset: 6888},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 17, offset: 6893},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 27, offset: 6903},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 224, col: 30, offset: 6906},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 34, offset: 6910},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 36, offset: 6912},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 224, col: 42, offset: 6918},
										expr: &ruleRefExpr{
											pos:  position{line: 224, col: 43, offset: 6919},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 55, offset: 6931},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 224, col: 58, offset: 6934},
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
			pos:  position{line: 236, col: 1, offset: 7232},
			expr: &choiceExpr{
				pos: position{line: 236, col: 8, offset: 7239},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 236, col: 8, offset: 7239},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 236, col: 8, offset: 7239},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 236, col: 8, offset: 7239},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 10, offset: 7241},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 17, offset: 7248},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 236, col: 28, offset: 7259},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 236, col: 32, offset: 7263},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 35, offset: 7266},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 236, col: 48, offset: 7279},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 53, offset: 7284},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 250, col: 1, offset: 7608},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 250, col: 1, offset: 7608},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 250, col: 1, offset: 7608},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 3, offset: 7610},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 6, offset: 7613},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 250, col: 19, offset: 7626},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 24, offset: 7631},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 264, col: 1, offset: 7948},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 264, col: 1, offset: 7948},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 1, offset: 7948},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 3, offset: 7950},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 6, offset: 7953},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 19, offset: 7966},
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
			name: "ArgsDefn",
			pos:  position{line: 271, col: 1, offset: 8137},
			expr: &actionExpr{
				pos: position{line: 271, col: 12, offset: 8148},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 271, col: 12, offset: 8148},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 271, col: 12, offset: 8148},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 271, col: 16, offset: 8152},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 271, col: 18, offset: 8154},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 271, col: 27, offset: 8163},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 271, col: 35, offset: 8171},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 271, col: 37, offset: 8173},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 271, col: 42, offset: 8178},
								expr: &seqExpr{
									pos: position{line: 271, col: 43, offset: 8179},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 271, col: 43, offset: 8179},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 271, col: 47, offset: 8183},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 271, col: 49, offset: 8185},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 271, col: 59, offset: 8195},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 271, col: 61, offset: 8197},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 289, col: 1, offset: 8619},
			expr: &actionExpr{
				pos: position{line: 289, col: 11, offset: 8629},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 289, col: 11, offset: 8629},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 289, col: 11, offset: 8629},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 16, offset: 8634},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 27, offset: 8645},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 29, offset: 8647},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 289, col: 34, offset: 8652},
								expr: &seqExpr{
									pos: position{line: 289, col: 35, offset: 8653},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 289, col: 35, offset: 8653},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 39, offset: 8657},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 41, offset: 8659},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 52, offset: 8670},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 309, col: 1, offset: 9158},
			expr: &choiceExpr{
				pos: position{line: 309, col: 13, offset: 9170},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 309, col: 13, offset: 9170},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 309, col: 13, offset: 9170},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 309, col: 13, offset: 9170},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 17, offset: 9174},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 309, col: 19, offset: 9176},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 309, col: 28, offset: 9185},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 40, offset: 9197},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 309, col: 42, offset: 9199},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 309, col: 47, offset: 9204},
										expr: &seqExpr{
											pos: position{line: 309, col: 48, offset: 9205},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 309, col: 48, offset: 9205},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 309, col: 52, offset: 9209},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 309, col: 54, offset: 9211},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 68, offset: 9225},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 309, col: 70, offset: 9227},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 326, col: 1, offset: 9649},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 326, col: 1, offset: 9649},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 326, col: 1, offset: 9649},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 5, offset: 9653},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 326, col: 7, offset: 9655},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 16, offset: 9664},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 21, offset: 9669},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 326, col: 23, offset: 9671},
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
			pos:  position{line: 331, col: 1, offset: 9776},
			expr: &actionExpr{
				pos: position{line: 331, col: 16, offset: 9791},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 331, col: 16, offset: 9791},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 331, col: 16, offset: 9791},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 331, col: 18, offset: 9793},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 331, col: 21, offset: 9796},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 331, col: 27, offset: 9802},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 331, col: 32, offset: 9807},
								expr: &seqExpr{
									pos: position{line: 331, col: 33, offset: 9808},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 331, col: 33, offset: 9808},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 37, offset: 9812},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 46, offset: 9821},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 50, offset: 9825},
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
			pos:  position{line: 351, col: 1, offset: 10431},
			expr: &choiceExpr{
				pos: position{line: 351, col: 9, offset: 10439},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 351, col: 9, offset: 10439},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 21, offset: 10451},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 37, offset: 10467},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 48, offset: 10478},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 60, offset: 10490},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 353, col: 1, offset: 10503},
			expr: &actionExpr{
				pos: position{line: 353, col: 13, offset: 10515},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 353, col: 13, offset: 10515},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 353, col: 13, offset: 10515},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 353, col: 15, offset: 10517},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 353, col: 21, offset: 10523},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 353, col: 35, offset: 10537},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 353, col: 40, offset: 10542},
								expr: &seqExpr{
									pos: position{line: 353, col: 41, offset: 10543},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 353, col: 41, offset: 10543},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 45, offset: 10547},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 61, offset: 10563},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 65, offset: 10567},
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
			pos:  position{line: 386, col: 1, offset: 11460},
			expr: &actionExpr{
				pos: position{line: 386, col: 17, offset: 11476},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 386, col: 17, offset: 11476},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 386, col: 17, offset: 11476},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 386, col: 19, offset: 11478},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 25, offset: 11484},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 386, col: 34, offset: 11493},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 386, col: 39, offset: 11498},
								expr: &seqExpr{
									pos: position{line: 386, col: 40, offset: 11499},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 386, col: 40, offset: 11499},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 44, offset: 11503},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 61, offset: 11520},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 65, offset: 11524},
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
			pos:  position{line: 418, col: 1, offset: 12411},
			expr: &actionExpr{
				pos: position{line: 418, col: 12, offset: 12422},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 418, col: 12, offset: 12422},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 418, col: 12, offset: 12422},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 418, col: 14, offset: 12424},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 418, col: 20, offset: 12430},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 418, col: 30, offset: 12440},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 418, col: 35, offset: 12445},
								expr: &seqExpr{
									pos: position{line: 418, col: 36, offset: 12446},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 418, col: 36, offset: 12446},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 40, offset: 12450},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 52, offset: 12462},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 56, offset: 12466},
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
			pos:  position{line: 450, col: 1, offset: 13354},
			expr: &actionExpr{
				pos: position{line: 450, col: 13, offset: 13366},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 450, col: 13, offset: 13366},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 450, col: 13, offset: 13366},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 450, col: 15, offset: 13368},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 450, col: 21, offset: 13374},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 450, col: 33, offset: 13386},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 450, col: 38, offset: 13391},
								expr: &seqExpr{
									pos: position{line: 450, col: 39, offset: 13392},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 450, col: 39, offset: 13392},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 43, offset: 13396},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 56, offset: 13409},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 60, offset: 13413},
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
			pos:  position{line: 481, col: 1, offset: 14302},
			expr: &choiceExpr{
				pos: position{line: 481, col: 15, offset: 14316},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 481, col: 15, offset: 14316},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 481, col: 15, offset: 14316},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 481, col: 15, offset: 14316},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 481, col: 17, offset: 14318},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 481, col: 21, offset: 14322},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 481, col: 24, offset: 14325},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 481, col: 30, offset: 14331},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 481, col: 36, offset: 14337},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 481, col: 39, offset: 14340},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 484, col: 5, offset: 14463},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 486, col: 1, offset: 14470},
			expr: &choiceExpr{
				pos: position{line: 486, col: 12, offset: 14481},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 486, col: 12, offset: 14481},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 486, col: 30, offset: 14499},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 486, col: 49, offset: 14518},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 486, col: 64, offset: 14533},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 488, col: 1, offset: 14546},
			expr: &actionExpr{
				pos: position{line: 488, col: 19, offset: 14564},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 488, col: 21, offset: 14566},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 488, col: 21, offset: 14566},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 488, col: 28, offset: 14573},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 492, col: 1, offset: 14655},
			expr: &actionExpr{
				pos: position{line: 492, col: 20, offset: 14674},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 492, col: 22, offset: 14676},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 492, col: 22, offset: 14676},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 29, offset: 14683},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 36, offset: 14690},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 42, offset: 14696},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 48, offset: 14702},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 56, offset: 14710},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 496, col: 1, offset: 14789},
			expr: &choiceExpr{
				pos: position{line: 496, col: 16, offset: 14804},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 496, col: 16, offset: 14804},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 496, col: 18, offset: 14806},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 496, col: 18, offset: 14806},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 496, col: 24, offset: 14812},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 499, col: 3, offset: 14895},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 499, col: 5, offset: 14897},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 502, col: 3, offset: 14977},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 502, col: 3, offset: 14977},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 506, col: 1, offset: 15058},
			expr: &actionExpr{
				pos: position{line: 506, col: 15, offset: 15072},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 506, col: 17, offset: 15074},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 506, col: 17, offset: 15074},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 506, col: 23, offset: 15080},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 510, col: 1, offset: 15162},
			expr: &choiceExpr{
				pos: position{line: 510, col: 9, offset: 15170},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 510, col: 9, offset: 15170},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 510, col: 16, offset: 15177},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 510, col: 31, offset: 15192},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 512, col: 1, offset: 15199},
			expr: &choiceExpr{
				pos: position{line: 512, col: 14, offset: 15212},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 512, col: 14, offset: 15212},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 512, col: 29, offset: 15227},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 514, col: 1, offset: 15235},
			expr: &choiceExpr{
				pos: position{line: 514, col: 14, offset: 15248},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 514, col: 14, offset: 15248},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 514, col: 29, offset: 15263},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 516, col: 1, offset: 15275},
			expr: &actionExpr{
				pos: position{line: 516, col: 16, offset: 15290},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 516, col: 16, offset: 15290},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 516, col: 16, offset: 15290},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 20, offset: 15294},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 516, col: 22, offset: 15296},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 516, col: 28, offset: 15302},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 33, offset: 15307},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 516, col: 35, offset: 15309},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 516, col: 40, offset: 15314},
								expr: &seqExpr{
									pos: position{line: 516, col: 41, offset: 15315},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 516, col: 41, offset: 15315},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 516, col: 45, offset: 15319},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 516, col: 47, offset: 15321},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 516, col: 52, offset: 15326},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 516, col: 56, offset: 15330},
							expr: &litMatcher{
								pos:        position{line: 516, col: 56, offset: 15330},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 61, offset: 15335},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 516, col: 63, offset: 15337},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 532, col: 1, offset: 15782},
			expr: &actionExpr{
				pos: position{line: 532, col: 19, offset: 15800},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 532, col: 19, offset: 15800},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 532, col: 19, offset: 15800},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 532, col: 24, offset: 15805},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 532, col: 35, offset: 15816},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 532, col: 37, offset: 15818},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 532, col: 42, offset: 15823},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 545, col: 1, offset: 16093},
			expr: &actionExpr{
				pos: position{line: 545, col: 18, offset: 16110},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 545, col: 18, offset: 16110},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 545, col: 18, offset: 16110},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 23, offset: 16115},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 34, offset: 16126},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 545, col: 36, offset: 16128},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 40, offset: 16132},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 545, col: 42, offset: 16134},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 52, offset: 16144},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 65, offset: 16157},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 545, col: 67, offset: 16159},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 71, offset: 16163},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 545, col: 73, offset: 16165},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 84, offset: 16176},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 545, col: 89, offset: 16181},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 545, col: 94, offset: 16186},
								expr: &seqExpr{
									pos: position{line: 545, col: 95, offset: 16187},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 545, col: 95, offset: 16187},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 99, offset: 16191},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 101, offset: 16193},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 114, offset: 16206},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 545, col: 116, offset: 16208},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 120, offset: 16212},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 122, offset: 16214},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 545, col: 130, offset: 16222},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 565, col: 1, offset: 16806},
			expr: &actionExpr{
				pos: position{line: 565, col: 17, offset: 16822},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 565, col: 17, offset: 16822},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 565, col: 17, offset: 16822},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 565, col: 22, offset: 16827},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 569, col: 1, offset: 16900},
			expr: &actionExpr{
				pos: position{line: 569, col: 16, offset: 16915},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 569, col: 16, offset: 16915},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 569, col: 16, offset: 16915},
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 17, offset: 16916},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 569, col: 27, offset: 16926},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 569, col: 27, offset: 16926},
									expr: &charClassMatcher{
										pos:        position{line: 569, col: 27, offset: 16926},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 569, col: 34, offset: 16933},
									expr: &charClassMatcher{
										pos:        position{line: 569, col: 34, offset: 16933},
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
			pos:  position{line: 573, col: 1, offset: 17008},
			expr: &actionExpr{
				pos: position{line: 573, col: 14, offset: 17021},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 573, col: 15, offset: 17022},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 573, col: 15, offset: 17022},
							expr: &charClassMatcher{
								pos:        position{line: 573, col: 15, offset: 17022},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 573, col: 22, offset: 17029},
							expr: &charClassMatcher{
								pos:        position{line: 573, col: 22, offset: 17029},
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
			pos:  position{line: 577, col: 1, offset: 17104},
			expr: &choiceExpr{
				pos: position{line: 577, col: 9, offset: 17112},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 577, col: 9, offset: 17112},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 577, col: 9, offset: 17112},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 577, col: 9, offset: 17112},
									expr: &litMatcher{
										pos:        position{line: 577, col: 9, offset: 17112},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 577, col: 14, offset: 17117},
									expr: &charClassMatcher{
										pos:        position{line: 577, col: 14, offset: 17117},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 577, col: 21, offset: 17124},
									expr: &litMatcher{
										pos:        position{line: 577, col: 22, offset: 17125},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 584, col: 3, offset: 17300},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 584, col: 3, offset: 17300},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 584, col: 3, offset: 17300},
									expr: &litMatcher{
										pos:        position{line: 584, col: 3, offset: 17300},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 584, col: 8, offset: 17305},
									expr: &charClassMatcher{
										pos:        position{line: 584, col: 8, offset: 17305},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 584, col: 15, offset: 17312},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 584, col: 19, offset: 17316},
									expr: &charClassMatcher{
										pos:        position{line: 584, col: 19, offset: 17316},
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
						pos: position{line: 591, col: 3, offset: 17505},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 591, col: 3, offset: 17505},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 595, col: 3, offset: 17590},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 595, col: 3, offset: 17590},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 598, col: 3, offset: 17676},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 3, offset: 17683},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 600, col: 3, offset: 17699},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 600, col: 3, offset: 17699},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 600, col: 3, offset: 17699},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 600, col: 7, offset: 17703},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 600, col: 12, offset: 17708},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 600, col: 12, offset: 17708},
												expr: &ruleRefExpr{
													pos:  position{line: 600, col: 13, offset: 17709},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 600, col: 25, offset: 17721,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 600, col: 28, offset: 17724},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 602, col: 5, offset: 17816},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 602, col: 20, offset: 17831},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 602, col: 37, offset: 17848},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 604, col: 1, offset: 17865},
			expr: &actionExpr{
				pos: position{line: 604, col: 8, offset: 17872},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 604, col: 8, offset: 17872},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 608, col: 1, offset: 17935},
			expr: &actionExpr{
				pos: position{line: 608, col: 17, offset: 17951},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 608, col: 17, offset: 17951},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 608, col: 17, offset: 17951},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 608, col: 21, offset: 17955},
							expr: &seqExpr{
								pos: position{line: 608, col: 22, offset: 17956},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 608, col: 22, offset: 17956},
										expr: &ruleRefExpr{
											pos:  position{line: 608, col: 23, offset: 17957},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 608, col: 35, offset: 17969,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 608, col: 39, offset: 17973},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 616, col: 1, offset: 18156},
			expr: &actionExpr{
				pos: position{line: 616, col: 10, offset: 18165},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 616, col: 11, offset: 18166},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 620, col: 1, offset: 18221},
			expr: &seqExpr{
				pos: position{line: 620, col: 12, offset: 18232},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 620, col: 13, offset: 18233},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 620, col: 13, offset: 18233},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 21, offset: 18241},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 28, offset: 18248},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 37, offset: 18257},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 46, offset: 18266},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 55, offset: 18275},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 64, offset: 18284},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 74, offset: 18294},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 86, offset: 18306},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 620, col: 95, offset: 18315},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 620, col: 105, offset: 18325},
						expr: &oneOrMoreExpr{
							pos: position{line: 620, col: 106, offset: 18326},
							expr: &charClassMatcher{
								pos:        position{line: 620, col: 106, offset: 18326},
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
			pos:  position{line: 622, col: 1, offset: 18334},
			expr: &choiceExpr{
				pos: position{line: 622, col: 12, offset: 18345},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 622, col: 12, offset: 18345},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 622, col: 14, offset: 18347},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 622, col: 14, offset: 18347},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 22, offset: 18355},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 31, offset: 18364},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 42, offset: 18375},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 51, offset: 18384},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 60, offset: 18393},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 70, offset: 18403},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 625, col: 3, offset: 18500},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 627, col: 1, offset: 18506},
			expr: &charClassMatcher{
				pos:        position{line: 627, col: 15, offset: 18520},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 629, col: 1, offset: 18536},
			expr: &choiceExpr{
				pos: position{line: 629, col: 18, offset: 18553},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 629, col: 18, offset: 18553},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 629, col: 37, offset: 18572},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 631, col: 1, offset: 18587},
			expr: &charClassMatcher{
				pos:        position{line: 631, col: 20, offset: 18606},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 633, col: 1, offset: 18619},
			expr: &charClassMatcher{
				pos:        position{line: 633, col: 16, offset: 18634},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 635, col: 1, offset: 18641},
			expr: &charClassMatcher{
				pos:        position{line: 635, col: 23, offset: 18663},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 637, col: 1, offset: 18670},
			expr: &charClassMatcher{
				pos:        position{line: 637, col: 12, offset: 18681},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 639, col: 1, offset: 18692},
			expr: &choiceExpr{
				pos: position{line: 639, col: 22, offset: 18713},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 639, col: 22, offset: 18713},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 639, col: 33, offset: 18724},
						expr: &charClassMatcher{
							pos:        position{line: 639, col: 33, offset: 18724},
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
			pos:         position{line: 641, col: 1, offset: 18736},
			expr: &choiceExpr{
				pos: position{line: 641, col: 21, offset: 18756},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 641, col: 21, offset: 18756},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 641, col: 32, offset: 18767},
						expr: &charClassMatcher{
							pos:        position{line: 641, col: 32, offset: 18767},
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
			pos:         position{line: 643, col: 1, offset: 18779},
			expr: &oneOrMoreExpr{
				pos: position{line: 643, col: 33, offset: 18811},
				expr: &charClassMatcher{
					pos:        position{line: 643, col: 33, offset: 18811},
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
			pos:         position{line: 645, col: 1, offset: 18819},
			expr: &zeroOrMoreExpr{
				pos: position{line: 645, col: 32, offset: 18850},
				expr: &charClassMatcher{
					pos:        position{line: 645, col: 32, offset: 18850},
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
			pos:         position{line: 647, col: 1, offset: 18858},
			expr: &choiceExpr{
				pos: position{line: 647, col: 15, offset: 18872},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 647, col: 15, offset: 18872},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 647, col: 26, offset: 18883},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 647, col: 26, offset: 18883},
								expr: &charClassMatcher{
									pos:        position{line: 647, col: 26, offset: 18883},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 647, col: 35, offset: 18892},
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
			pos:  position{line: 649, col: 1, offset: 18898},
			expr: &oneOrMoreExpr{
				pos: position{line: 649, col: 12, offset: 18909},
				expr: &ruleRefExpr{
					pos:  position{line: 649, col: 13, offset: 18910},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 651, col: 1, offset: 18921},
			expr: &seqExpr{
				pos: position{line: 651, col: 11, offset: 18931},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 651, col: 11, offset: 18931},
						expr: &charClassMatcher{
							pos:        position{line: 651, col: 11, offset: 18931},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 651, col: 22, offset: 18942},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 651, col: 26, offset: 18946},
						expr: &seqExpr{
							pos: position{line: 651, col: 27, offset: 18947},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 651, col: 27, offset: 18947},
									expr: &charClassMatcher{
										pos:        position{line: 651, col: 28, offset: 18948},
										val:        "[\\n]",
										chars:      []rune{'\n'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&anyMatcher{
									line: 651, col: 33, offset: 18953,
								},
							},
						},
					},
					&andExpr{
						pos: position{line: 651, col: 37, offset: 18957},
						expr: &litMatcher{
							pos:        position{line: 651, col: 38, offset: 18958},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 653, col: 1, offset: 18964},
			expr: &notExpr{
				pos: position{line: 653, col: 7, offset: 18970},
				expr: &anyMatcher{
					line: 653, col: 8, offset: 18971,
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
		return Module{Name: name.(Identifier).StringValue, Subvalues: subvalues}, nil
	} else {
		return Module{Name: name.(Identifier).StringValue, Subvalues: []Ast{stat.(Ast)}}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["name"], stack["stat"], stack["rest"])
}

func (c *current) onExprLine1(e interface{}) (interface{}, error) {
	fmt.Println("line:", e)
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

func (c *current) onExternDefn1(name, importName, args, ret interface{}) (interface{}, error) {

	return Extern{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Arguments: args.(Container).Subvalues, ReturnAnnotation: ret.(BasicAst).StringValue}, nil
}

func (p *parser) callonExternDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternDefn1(stack["name"], stack["importName"], stack["args"], stack["ret"])
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

	retType := ""
	if ret != nil {
		vals = ret.([]interface{})
		retType = vals[0].(BasicAst).StringValue
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

		switch vals[2].(type) {
		case BasicAst:
			arg.Annotation = vals[2].(BasicAst).StringValue
		case Identifier:
			arg.Annotation = vals[2].(Identifier).StringValue
		}
	}
	//fmt.Println("parsed:", arg)
	return arg, nil
}

func (p *parser) callonArgDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgDefn1(stack["name"], stack["anno"])
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
	return Operator{StringValue: string(c.text), ValueType: NUMBER}, nil
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

func (c *current) onOperatorHigh8() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh8()
}

func (c *current) onOperatorLow1() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: NUMBER}, nil
}

func (p *parser) callonOperatorLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow1()
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
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
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
