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
			pos:  position{line: 28, col: 1, offset: 634},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 654},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 654},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 31, offset: 664},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 42, offset: 675},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 685},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 697},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 697},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 707},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 718},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 731},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 741},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 752},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 752},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 752},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 754},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 759},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 760},
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
			pos:  position{line: 36, col: 1, offset: 788},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 798},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 798},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 36, col: 11, offset: 798},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 15, offset: 802},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 23, offset: 810},
								expr: &seqExpr{
									pos: position{line: 36, col: 24, offset: 811},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 24, offset: 811},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 25, offset: 812},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 37, offset: 824,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 41, offset: 828},
							expr: &litMatcher{
								pos:        position{line: 36, col: 42, offset: 829},
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
			pos:  position{line: 41, col: 1, offset: 979},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 990},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 990},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 990},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 990},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 992},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 999},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 1002},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 1007},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 1018},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 1025},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 1026},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 1026},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 1029},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1045},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1047},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1051},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1057},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1058},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1058},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1061},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1071},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1565},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1565},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1565},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1567},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1574},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1577},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1582},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1593},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1600},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1601},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1601},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1604},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1620},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1622},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1626},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1632},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1636},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1638},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1644},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1660},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1662},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1667},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1668},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1668},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1672},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1674},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1690},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1694},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1694},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1699},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1701},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1705},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2188},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2188},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2188},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2190},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2197},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2200},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2205},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2216},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2223},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2224},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2224},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2227},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2243},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2245},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2249},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2251},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2256},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2257},
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
			pos:  position{line: 94, col: 1, offset: 2662},
			expr: &actionExpr{
				pos: position{line: 94, col: 22, offset: 2683},
				run: (*parser).callonVariantConstructor1,
				expr: &seqExpr{
					pos: position{line: 94, col: 22, offset: 2683},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 94, col: 22, offset: 2683},
							val:        "|",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 26, offset: 2687},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 28, offset: 2689},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 33, offset: 2694},
								name: "ModuleName",
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 44, offset: 2705},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 94, col: 49, offset: 2710},
								expr: &seqExpr{
									pos: position{line: 94, col: 50, offset: 2711},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 94, col: 50, offset: 2711},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 94, col: 53, offset: 2714},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 63, offset: 2724},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 111, col: 1, offset: 3159},
			expr: &actionExpr{
				pos: position{line: 111, col: 19, offset: 3177},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 111, col: 19, offset: 3177},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 111, col: 19, offset: 3177},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 111, col: 24, offset: 3182},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 111, col: 37, offset: 3195},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 111, col: 39, offset: 3197},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 111, col: 43, offset: 3201},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 111, col: 45, offset: 3203},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 111, col: 48, offset: 3206},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 115, col: 1, offset: 3298},
			expr: &choiceExpr{
				pos: position{line: 115, col: 11, offset: 3308},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 115, col: 11, offset: 3308},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 22, offset: 3319},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 117, col: 1, offset: 3334},
			expr: &choiceExpr{
				pos: position{line: 117, col: 14, offset: 3347},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 117, col: 14, offset: 3347},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 117, col: 14, offset: 3347},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 117, col: 14, offset: 3347},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 117, col: 16, offset: 3349},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 22, offset: 3355},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 25, offset: 3358},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 27, offset: 3360},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 38, offset: 3371},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 43, offset: 3376},
										expr: &seqExpr{
											pos: position{line: 117, col: 44, offset: 3377},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 117, col: 44, offset: 3377},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 48, offset: 3381},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 50, offset: 3383},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 63, offset: 3396},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 117, col: 65, offset: 3398},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 69, offset: 3402},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 71, offset: 3404},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 76, offset: 3409},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 81, offset: 3414},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 132, col: 1, offset: 3853},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 132, col: 1, offset: 3853},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 132, col: 1, offset: 3853},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 132, col: 3, offset: 3855},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 132, col: 9, offset: 3861},
									name: "__",
								},
								&notExpr{
									pos: position{line: 132, col: 12, offset: 3864},
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 13, offset: 3865},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 136, col: 1, offset: 3973},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 136, col: 1, offset: 3973},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 1, offset: 3973},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 3, offset: 3975},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 9, offset: 3981},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 12, offset: 3984},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 14, offset: 3986},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 3997},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 136, col: 30, offset: 4002},
										expr: &seqExpr{
											pos: position{line: 136, col: 31, offset: 4003},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 136, col: 31, offset: 4003},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 35, offset: 4007},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 37, offset: 4009},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 50, offset: 4022},
									name: "_",
								},
								&notExpr{
									pos: position{line: 136, col: 52, offset: 4024},
									expr: &litMatcher{
										pos:        position{line: 136, col: 53, offset: 4025},
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
			pos:  position{line: 140, col: 1, offset: 4119},
			expr: &actionExpr{
				pos: position{line: 140, col: 12, offset: 4130},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 140, col: 12, offset: 4130},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 140, col: 12, offset: 4130},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 140, col: 14, offset: 4132},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 20, offset: 4138},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 140, col: 23, offset: 4141},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 140, col: 25, offset: 4143},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 38, offset: 4156},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 140, col: 40, offset: 4158},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 44, offset: 4162},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 140, col: 46, offset: 4164},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 53, offset: 4171},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 140, col: 56, offset: 4174},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 140, col: 60, offset: 4178},
								expr: &seqExpr{
									pos: position{line: 140, col: 61, offset: 4179},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 140, col: 61, offset: 4179},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 140, col: 74, offset: 4192},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 79, offset: 4197},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 140, col: 81, offset: 4199},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 85, offset: 4203},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 140, col: 88, offset: 4206},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 140, col: 99, offset: 4217},
								expr: &ruleRefExpr{
									pos:  position{line: 140, col: 100, offset: 4218},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 112, offset: 4230},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 140, col: 114, offset: 4232},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 140, col: 118, offset: 4236},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 163, col: 1, offset: 4911},
			expr: &actionExpr{
				pos: position{line: 163, col: 8, offset: 4918},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 163, col: 8, offset: 4918},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 163, col: 12, offset: 4922},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 163, col: 12, offset: 4922},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 163, col: 21, offset: 4931},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 163, col: 28, offset: 4938},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 168, col: 1, offset: 5033},
			expr: &choiceExpr{
				pos: position{line: 168, col: 10, offset: 5042},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 168, col: 10, offset: 5042},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 168, col: 10, offset: 5042},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 168, col: 10, offset: 5042},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 15, offset: 5047},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 168, col: 18, offset: 5050},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 168, col: 23, offset: 5055},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 33, offset: 5065},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 168, col: 35, offset: 5067},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 39, offset: 5071},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 168, col: 41, offset: 5073},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 168, col: 47, offset: 5079},
										expr: &ruleRefExpr{
											pos:  position{line: 168, col: 48, offset: 5080},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 60, offset: 5092},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 168, col: 62, offset: 5094},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 66, offset: 5098},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 168, col: 68, offset: 5100},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 75, offset: 5107},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 168, col: 77, offset: 5109},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 168, col: 85, offset: 5117},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 180, col: 1, offset: 5447},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 180, col: 1, offset: 5447},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 180, col: 1, offset: 5447},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 6, offset: 5452},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 9, offset: 5455},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 14, offset: 5460},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 24, offset: 5470},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 180, col: 26, offset: 5472},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 30, offset: 5476},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 32, offset: 5478},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 180, col: 38, offset: 5484},
										expr: &ruleRefExpr{
											pos:  position{line: 180, col: 39, offset: 5485},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 51, offset: 5497},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 180, col: 54, offset: 5500},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 58, offset: 5504},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 180, col: 60, offset: 5506},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 67, offset: 5513},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 180, col: 69, offset: 5515},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 73, offset: 5519},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 180, col: 75, offset: 5521},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 180, col: 81, offset: 5527},
										expr: &ruleRefExpr{
											pos:  position{line: 180, col: 82, offset: 5528},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 94, offset: 5540},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 180, col: 97, offset: 5543},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 199, col: 1, offset: 6046},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 199, col: 1, offset: 6046},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 199, col: 1, offset: 6046},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 199, col: 6, offset: 6051},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 199, col: 9, offset: 6054},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 199, col: 14, offset: 6059},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 199, col: 24, offset: 6069},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 199, col: 26, offset: 6071},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 199, col: 30, offset: 6075},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 199, col: 32, offset: 6077},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 199, col: 38, offset: 6083},
										expr: &ruleRefExpr{
											pos:  position{line: 199, col: 39, offset: 6084},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 199, col: 51, offset: 6096},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 199, col: 54, offset: 6099},
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
			pos:  position{line: 212, col: 1, offset: 6398},
			expr: &choiceExpr{
				pos: position{line: 212, col: 8, offset: 6405},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 212, col: 8, offset: 6405},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 212, col: 8, offset: 6405},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 212, col: 8, offset: 6405},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 15, offset: 6412},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 212, col: 26, offset: 6423},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 212, col: 30, offset: 6427},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 33, offset: 6430},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 212, col: 46, offset: 6443},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 212, col: 56, offset: 6453},
										expr: &seqExpr{
											pos: position{line: 212, col: 57, offset: 6454},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 212, col: 57, offset: 6454},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 212, col: 60, offset: 6457},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 74, offset: 6471},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 230, col: 1, offset: 6969},
						run: (*parser).callonCall15,
						expr: &seqExpr{
							pos: position{line: 230, col: 1, offset: 6969},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 230, col: 1, offset: 6969},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 230, col: 3, offset: 6971},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 230, col: 6, offset: 6974},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 230, col: 19, offset: 6987},
									label: "arguments",
									expr: &oneOrMoreExpr{
										pos: position{line: 230, col: 29, offset: 6997},
										expr: &seqExpr{
											pos: position{line: 230, col: 30, offset: 6998},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 230, col: 30, offset: 6998},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 230, col: 33, offset: 7001},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 230, col: 47, offset: 7015},
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
			pos:  position{line: 248, col: 1, offset: 7525},
			expr: &actionExpr{
				pos: position{line: 248, col: 16, offset: 7540},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 248, col: 16, offset: 7540},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 248, col: 16, offset: 7540},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 248, col: 18, offset: 7542},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 21, offset: 7545},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 27, offset: 7551},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 248, col: 32, offset: 7556},
								expr: &seqExpr{
									pos: position{line: 248, col: 33, offset: 7557},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 248, col: 33, offset: 7557},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 248, col: 36, offset: 7560},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 248, col: 45, offset: 7569},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 248, col: 48, offset: 7572},
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
			pos:  position{line: 268, col: 1, offset: 8236},
			expr: &choiceExpr{
				pos: position{line: 268, col: 9, offset: 8244},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 268, col: 9, offset: 8244},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 268, col: 21, offset: 8256},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 268, col: 37, offset: 8272},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 268, col: 48, offset: 8283},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 268, col: 60, offset: 8295},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 270, col: 1, offset: 8308},
			expr: &actionExpr{
				pos: position{line: 270, col: 13, offset: 8320},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 270, col: 13, offset: 8320},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 270, col: 13, offset: 8320},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 270, col: 15, offset: 8322},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 270, col: 21, offset: 8328},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 270, col: 35, offset: 8342},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 270, col: 40, offset: 8347},
								expr: &seqExpr{
									pos: position{line: 270, col: 41, offset: 8348},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 270, col: 41, offset: 8348},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 270, col: 44, offset: 8351},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 270, col: 60, offset: 8367},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 270, col: 63, offset: 8370},
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
			pos:  position{line: 289, col: 1, offset: 8976},
			expr: &actionExpr{
				pos: position{line: 289, col: 17, offset: 8992},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 289, col: 17, offset: 8992},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 289, col: 17, offset: 8992},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 19, offset: 8994},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 25, offset: 9000},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 289, col: 34, offset: 9009},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 289, col: 39, offset: 9014},
								expr: &seqExpr{
									pos: position{line: 289, col: 40, offset: 9015},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 289, col: 40, offset: 9015},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 43, offset: 9018},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 60, offset: 9035},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 63, offset: 9038},
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
			pos:  position{line: 309, col: 1, offset: 9642},
			expr: &actionExpr{
				pos: position{line: 309, col: 12, offset: 9653},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 309, col: 12, offset: 9653},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 309, col: 12, offset: 9653},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 309, col: 14, offset: 9655},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 309, col: 20, offset: 9661},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 309, col: 30, offset: 9671},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 309, col: 35, offset: 9676},
								expr: &seqExpr{
									pos: position{line: 309, col: 36, offset: 9677},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 309, col: 36, offset: 9677},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 309, col: 39, offset: 9680},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 309, col: 51, offset: 9692},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 309, col: 54, offset: 9695},
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
			pos:  position{line: 329, col: 1, offset: 10296},
			expr: &actionExpr{
				pos: position{line: 329, col: 13, offset: 10308},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 329, col: 13, offset: 10308},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 329, col: 13, offset: 10308},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 329, col: 15, offset: 10310},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 329, col: 21, offset: 10316},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 329, col: 33, offset: 10328},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 329, col: 38, offset: 10333},
								expr: &seqExpr{
									pos: position{line: 329, col: 39, offset: 10334},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 329, col: 39, offset: 10334},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 329, col: 42, offset: 10337},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 329, col: 55, offset: 10350},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 329, col: 58, offset: 10353},
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
			pos:  position{line: 348, col: 1, offset: 10957},
			expr: &choiceExpr{
				pos: position{line: 348, col: 15, offset: 10971},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 348, col: 15, offset: 10971},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 348, col: 15, offset: 10971},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 348, col: 15, offset: 10971},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 348, col: 19, offset: 10975},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 348, col: 21, offset: 10977},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 348, col: 27, offset: 10983},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 348, col: 33, offset: 10989},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 348, col: 35, offset: 10991},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 5, offset: 11140},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 353, col: 1, offset: 11147},
			expr: &choiceExpr{
				pos: position{line: 353, col: 12, offset: 11158},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 353, col: 12, offset: 11158},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 353, col: 30, offset: 11176},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 353, col: 49, offset: 11195},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 353, col: 64, offset: 11210},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 355, col: 1, offset: 11223},
			expr: &actionExpr{
				pos: position{line: 355, col: 19, offset: 11241},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 355, col: 21, offset: 11243},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 355, col: 21, offset: 11243},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 355, col: 29, offset: 11251},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 355, col: 36, offset: 11258},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 359, col: 1, offset: 11357},
			expr: &actionExpr{
				pos: position{line: 359, col: 20, offset: 11376},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 359, col: 22, offset: 11378},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 359, col: 22, offset: 11378},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 359, col: 29, offset: 11385},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 359, col: 36, offset: 11392},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 359, col: 42, offset: 11398},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 359, col: 48, offset: 11404},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 359, col: 56, offset: 11412},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 363, col: 1, offset: 11518},
			expr: &choiceExpr{
				pos: position{line: 363, col: 16, offset: 11533},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 363, col: 16, offset: 11533},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 363, col: 18, offset: 11535},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 363, col: 18, offset: 11535},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 363, col: 25, offset: 11542},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 366, col: 3, offset: 11648},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 366, col: 5, offset: 11650},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 366, col: 5, offset: 11650},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 366, col: 11, offset: 11656},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 366, col: 17, offset: 11662},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 369, col: 3, offset: 11765},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 369, col: 3, offset: 11765},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 373, col: 1, offset: 11869},
			expr: &choiceExpr{
				pos: position{line: 373, col: 15, offset: 11883},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 373, col: 15, offset: 11883},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 373, col: 17, offset: 11885},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 373, col: 17, offset: 11885},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 373, col: 24, offset: 11892},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 376, col: 3, offset: 11998},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 376, col: 5, offset: 12000},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 376, col: 5, offset: 12000},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 376, col: 11, offset: 12006},
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
			pos:  position{line: 380, col: 1, offset: 12108},
			expr: &choiceExpr{
				pos: position{line: 380, col: 9, offset: 12116},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 380, col: 9, offset: 12116},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 380, col: 24, offset: 12131},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 382, col: 1, offset: 12138},
			expr: &choiceExpr{
				pos: position{line: 382, col: 14, offset: 12151},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 382, col: 14, offset: 12151},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 382, col: 29, offset: 12166},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 384, col: 1, offset: 12174},
			expr: &choiceExpr{
				pos: position{line: 384, col: 14, offset: 12187},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 384, col: 14, offset: 12187},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 384, col: 29, offset: 12202},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 386, col: 1, offset: 12214},
			expr: &actionExpr{
				pos: position{line: 386, col: 16, offset: 12229},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 386, col: 16, offset: 12229},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 386, col: 16, offset: 12229},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 386, col: 20, offset: 12233},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 386, col: 22, offset: 12235},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 28, offset: 12241},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 386, col: 33, offset: 12246},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 386, col: 35, offset: 12248},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 386, col: 40, offset: 12253},
								expr: &seqExpr{
									pos: position{line: 386, col: 41, offset: 12254},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 386, col: 41, offset: 12254},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 45, offset: 12258},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 47, offset: 12260},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 52, offset: 12265},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 386, col: 56, offset: 12269},
							expr: &litMatcher{
								pos:        position{line: 386, col: 56, offset: 12269},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 386, col: 61, offset: 12274},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 386, col: 63, offset: 12276},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 402, col: 1, offset: 12757},
			expr: &actionExpr{
				pos: position{line: 402, col: 17, offset: 12773},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 402, col: 17, offset: 12773},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 402, col: 17, offset: 12773},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 402, col: 22, offset: 12778},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 406, col: 1, offset: 12886},
			expr: &actionExpr{
				pos: position{line: 406, col: 16, offset: 12901},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 406, col: 16, offset: 12901},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 406, col: 16, offset: 12901},
							expr: &ruleRefExpr{
								pos:  position{line: 406, col: 17, offset: 12902},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 406, col: 27, offset: 12912},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 406, col: 27, offset: 12912},
									expr: &charClassMatcher{
										pos:        position{line: 406, col: 27, offset: 12912},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 406, col: 34, offset: 12919},
									expr: &charClassMatcher{
										pos:        position{line: 406, col: 34, offset: 12919},
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
			pos:  position{line: 410, col: 1, offset: 13030},
			expr: &actionExpr{
				pos: position{line: 410, col: 14, offset: 13043},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 410, col: 15, offset: 13044},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 410, col: 15, offset: 13044},
							expr: &charClassMatcher{
								pos:        position{line: 410, col: 15, offset: 13044},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 410, col: 22, offset: 13051},
							expr: &charClassMatcher{
								pos:        position{line: 410, col: 22, offset: 13051},
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
			pos:  position{line: 414, col: 1, offset: 13162},
			expr: &choiceExpr{
				pos: position{line: 414, col: 9, offset: 13170},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 414, col: 9, offset: 13170},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 414, col: 9, offset: 13170},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 414, col: 9, offset: 13170},
									expr: &litMatcher{
										pos:        position{line: 414, col: 9, offset: 13170},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 414, col: 14, offset: 13175},
									expr: &charClassMatcher{
										pos:        position{line: 414, col: 14, offset: 13175},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 414, col: 21, offset: 13182},
									expr: &litMatcher{
										pos:        position{line: 414, col: 22, offset: 13183},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 421, col: 3, offset: 13359},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 421, col: 3, offset: 13359},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 421, col: 3, offset: 13359},
									expr: &litMatcher{
										pos:        position{line: 421, col: 3, offset: 13359},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 421, col: 8, offset: 13364},
									expr: &charClassMatcher{
										pos:        position{line: 421, col: 8, offset: 13364},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 421, col: 15, offset: 13371},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 421, col: 19, offset: 13375},
									expr: &charClassMatcher{
										pos:        position{line: 421, col: 19, offset: 13375},
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
						pos:        position{line: 428, col: 3, offset: 13565},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 428, col: 12, offset: 13574},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 428, col: 12, offset: 13574},
							val:        "false",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 434, col: 3, offset: 13775},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 434, col: 3, offset: 13775},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 437, col: 3, offset: 13838},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 437, col: 3, offset: 13838},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 437, col: 3, offset: 13838},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 437, col: 7, offset: 13842},
									expr: &seqExpr{
										pos: position{line: 437, col: 8, offset: 13843},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 437, col: 8, offset: 13843},
												expr: &ruleRefExpr{
													pos:  position{line: 437, col: 9, offset: 13844},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 437, col: 21, offset: 13856,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 25, offset: 13860},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 444, col: 3, offset: 14044},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 444, col: 3, offset: 14044},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 444, col: 3, offset: 14044},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 444, col: 7, offset: 14048},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 444, col: 12, offset: 14053},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 444, col: 12, offset: 14053},
												expr: &ruleRefExpr{
													pos:  position{line: 444, col: 13, offset: 14054},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 444, col: 25, offset: 14066,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 444, col: 28, offset: 14069},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 446, col: 5, offset: 14161},
						name: "ArrayLiteral",
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 448, col: 1, offset: 14175},
			expr: &actionExpr{
				pos: position{line: 448, col: 10, offset: 14184},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 448, col: 11, offset: 14185},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 452, col: 1, offset: 14286},
			expr: &seqExpr{
				pos: position{line: 452, col: 12, offset: 14297},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 452, col: 13, offset: 14298},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 452, col: 13, offset: 14298},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 21, offset: 14306},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 28, offset: 14313},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 37, offset: 14322},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 46, offset: 14331},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 55, offset: 14340},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 64, offset: 14349},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 74, offset: 14359},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 452, col: 86, offset: 14371},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 452, col: 95, offset: 14380},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 452, col: 105, offset: 14390},
						expr: &oneOrMoreExpr{
							pos: position{line: 452, col: 106, offset: 14391},
							expr: &charClassMatcher{
								pos:        position{line: 452, col: 106, offset: 14391},
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
			pos:  position{line: 454, col: 1, offset: 14399},
			expr: &actionExpr{
				pos: position{line: 454, col: 12, offset: 14410},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 454, col: 14, offset: 14412},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 454, col: 14, offset: 14412},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 454, col: 22, offset: 14420},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 454, col: 31, offset: 14429},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 454, col: 42, offset: 14440},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 454, col: 51, offset: 14449},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 454, col: 60, offset: 14458},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 454, col: 70, offset: 14468},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 458, col: 1, offset: 14567},
			expr: &charClassMatcher{
				pos:        position{line: 458, col: 15, offset: 14581},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 460, col: 1, offset: 14597},
			expr: &choiceExpr{
				pos: position{line: 460, col: 18, offset: 14614},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 460, col: 18, offset: 14614},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 460, col: 37, offset: 14633},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 462, col: 1, offset: 14648},
			expr: &charClassMatcher{
				pos:        position{line: 462, col: 20, offset: 14667},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 464, col: 1, offset: 14680},
			expr: &charClassMatcher{
				pos:        position{line: 464, col: 16, offset: 14695},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 466, col: 1, offset: 14702},
			expr: &charClassMatcher{
				pos:        position{line: 466, col: 23, offset: 14724},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 468, col: 1, offset: 14731},
			expr: &charClassMatcher{
				pos:        position{line: 468, col: 12, offset: 14742},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 470, col: 1, offset: 14753},
			expr: &oneOrMoreExpr{
				pos: position{line: 470, col: 22, offset: 14774},
				expr: &charClassMatcher{
					pos:        position{line: 470, col: 22, offset: 14774},
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
			pos:         position{line: 472, col: 1, offset: 14786},
			expr: &zeroOrMoreExpr{
				pos: position{line: 472, col: 18, offset: 14803},
				expr: &charClassMatcher{
					pos:        position{line: 472, col: 18, offset: 14803},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 474, col: 1, offset: 14815},
			expr: &notExpr{
				pos: position{line: 474, col: 7, offset: 14821},
				expr: &anyMatcher{
					line: 474, col: 8, offset: 14822,
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
	return BasicAst{Type: "Comment", StringValue: string(c.text[1:]), ValueType: STRING}, nil
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

func (c *current) onVariantConstructor1(name, rest interface{}) (interface{}, error) {
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
