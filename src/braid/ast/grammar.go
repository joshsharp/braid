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
						name: "ExternFunc",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 56, offset: 774},
						name: "ExternType",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 786},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 798},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 798},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 809},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 822},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 832},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 843},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 843},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 843},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 845},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 32, col: 19, offset: 850},
							name: "N",
						},
					},
				},
			},
		},
		{
			name: "ExternFunc",
			pos:  position{line: 46, col: 1, offset: 1181},
			expr: &actionExpr{
				pos: position{line: 46, col: 14, offset: 1194},
				run: (*parser).callonExternFunc1,
				expr: &seqExpr{
					pos: position{line: 46, col: 14, offset: 1194},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 14, offset: 1194},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 16, offset: 1196},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 25, offset: 1205},
							name: "__N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 29, offset: 1209},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 36, offset: 1216},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 40, offset: 1220},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 45, offset: 1225},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 56, offset: 1236},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 59, offset: 1239},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 63, offset: 1243},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 66, offset: 1246},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 77, offset: 1257},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 91, offset: 1271},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 94, offset: 1274},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 99, offset: 1279},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 5, offset: 1292},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 47, col: 8, offset: 1295},
							val:        "->",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 13, offset: 1300},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 16, offset: 1303},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 20, offset: 1307},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "ExternType",
			pos:  position{line: 53, col: 1, offset: 1517},
			expr: &choiceExpr{
				pos: position{line: 53, col: 14, offset: 1530},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 53, col: 14, offset: 1530},
						run: (*parser).callonExternType2,
						expr: &seqExpr{
							pos: position{line: 53, col: 14, offset: 1530},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 53, col: 14, offset: 1530},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 16, offset: 1532},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 25, offset: 1541},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 53, col: 29, offset: 1545},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 36, offset: 1552},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 39, offset: 1555},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 53, col: 44, offset: 1560},
										name: "OptionalPointerName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 64, offset: 1580},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 66, offset: 1582},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 70, offset: 1586},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 5, offset: 1592},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 16, offset: 1603},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 30, offset: 1617},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 32, offset: 1619},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 36, offset: 1623},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 38, offset: 1625},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 44, offset: 1631},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 60, offset: 1647},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 62, offset: 1649},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 54, col: 67, offset: 1654},
										expr: &seqExpr{
											pos: position{line: 54, col: 68, offset: 1655},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 54, col: 68, offset: 1655},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 72, offset: 1659},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 74, offset: 1661},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 90, offset: 1677},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 54, col: 94, offset: 1681},
									expr: &litMatcher{
										pos:        position{line: 54, col: 94, offset: 1681},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 99, offset: 1686},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 101, offset: 1688},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 105, offset: 1692},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 73, col: 1, offset: 2241},
						run: (*parser).callonExternType34,
						expr: &seqExpr{
							pos: position{line: 73, col: 1, offset: 2241},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 73, col: 1, offset: 2241},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 3, offset: 2243},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 12, offset: 2252},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 73, col: 16, offset: 2256},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 23, offset: 2263},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 26, offset: 2266},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 73, col: 31, offset: 2271},
										name: "OptionalPointerName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 51, offset: 2291},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 53, offset: 2293},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 57, offset: 2297},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 5, offset: 2303},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 74, col: 16, offset: 2314},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 30, offset: 2328},
									name: "__",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 82, col: 1, offset: 2523},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2534},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 82, col: 12, offset: 2534},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 82, col: 12, offset: 2534},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 82, col: 12, offset: 2534},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 14, offset: 2536},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 21, offset: 2543},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 24, offset: 2546},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 29, offset: 2551},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 82, col: 40, offset: 2562},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 82, col: 47, offset: 2569},
										expr: &seqExpr{
											pos: position{line: 82, col: 48, offset: 2570},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 48, offset: 2570},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 51, offset: 2573},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 67, offset: 2589},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 69, offset: 2591},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 82, col: 73, offset: 2595},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 82, col: 79, offset: 2601},
										expr: &seqExpr{
											pos: position{line: 82, col: 80, offset: 2602},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 80, offset: 2602},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 83, offset: 2605},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 100, offset: 2622},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 101, col: 1, offset: 3118},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 101, col: 1, offset: 3118},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 101, col: 1, offset: 3118},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 3, offset: 3120},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 10, offset: 3127},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 13, offset: 3130},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 18, offset: 3135},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 101, col: 29, offset: 3146},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 36, offset: 3153},
										expr: &seqExpr{
											pos: position{line: 101, col: 37, offset: 3154},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 101, col: 37, offset: 3154},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 40, offset: 3157},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 56, offset: 3173},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 58, offset: 3175},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 62, offset: 3179},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 5, offset: 3185},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 9, offset: 3189},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 102, col: 11, offset: 3191},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 102, col: 17, offset: 3197},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 33, offset: 3213},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 102, col: 35, offset: 3215},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 102, col: 40, offset: 3220},
										expr: &seqExpr{
											pos: position{line: 102, col: 41, offset: 3221},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 102, col: 41, offset: 3221},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 45, offset: 3225},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 47, offset: 3227},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 63, offset: 3243},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 102, col: 67, offset: 3247},
									expr: &litMatcher{
										pos:        position{line: 102, col: 67, offset: 3247},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 72, offset: 3252},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 74, offset: 3254},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 78, offset: 3258},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 120, col: 1, offset: 3743},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 120, col: 1, offset: 3743},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 120, col: 1, offset: 3743},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 3, offset: 3745},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 10, offset: 3752},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 13, offset: 3755},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 120, col: 18, offset: 3760},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 120, col: 29, offset: 3771},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 120, col: 36, offset: 3778},
										expr: &seqExpr{
											pos: position{line: 120, col: 37, offset: 3779},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 120, col: 37, offset: 3779},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 40, offset: 3782},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 56, offset: 3798},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 58, offset: 3800},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 62, offset: 3804},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 64, offset: 3806},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 120, col: 69, offset: 3811},
										expr: &ruleRefExpr{
											pos:  position{line: 120, col: 70, offset: 3812},
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
			pos:  position{line: 135, col: 1, offset: 4219},
			expr: &actionExpr{
				pos: position{line: 135, col: 19, offset: 4237},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 135, col: 19, offset: 4237},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 135, col: 19, offset: 4237},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 24, offset: 4242},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 135, col: 37, offset: 4255},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 135, col: 39, offset: 4257},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 135, col: 43, offset: 4261},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 135, col: 45, offset: 4263},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 48, offset: 4266},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 139, col: 1, offset: 4367},
			expr: &choiceExpr{
				pos: position{line: 139, col: 22, offset: 4388},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 139, col: 22, offset: 4388},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 139, col: 22, offset: 4388},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 139, col: 22, offset: 4388},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 26, offset: 4392},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 28, offset: 4394},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 139, col: 33, offset: 4399},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 44, offset: 4410},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 139, col: 46, offset: 4412},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 50, offset: 4416},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 52, offset: 4418},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 139, col: 58, offset: 4424},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 74, offset: 4440},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 76, offset: 4442},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 139, col: 81, offset: 4447},
										expr: &seqExpr{
											pos: position{line: 139, col: 82, offset: 4448},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 139, col: 82, offset: 4448},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 86, offset: 4452},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 88, offset: 4454},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 104, offset: 4470},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 139, col: 108, offset: 4474},
									expr: &litMatcher{
										pos:        position{line: 139, col: 108, offset: 4474},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 113, offset: 4479},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 139, col: 115, offset: 4481},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 119, offset: 4485},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 158, col: 1, offset: 5090},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 158, col: 1, offset: 5090},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 158, col: 1, offset: 5090},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 158, col: 5, offset: 5094},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 158, col: 7, offset: 5096},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 158, col: 12, offset: 5101},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 158, col: 23, offset: 5112},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 158, col: 28, offset: 5117},
										expr: &seqExpr{
											pos: position{line: 158, col: 29, offset: 5118},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 158, col: 29, offset: 5118},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 158, col: 32, offset: 5121},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 158, col: 49, offset: 5138},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 175, col: 1, offset: 5575},
			expr: &choiceExpr{
				pos: position{line: 175, col: 14, offset: 5588},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 175, col: 14, offset: 5588},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 175, col: 14, offset: 5588},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 175, col: 14, offset: 5588},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 16, offset: 5590},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 22, offset: 5596},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 26, offset: 5600},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 28, offset: 5602},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 39, offset: 5613},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 175, col: 42, offset: 5616},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 46, offset: 5620},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 49, offset: 5623},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 54, offset: 5628},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 59, offset: 5633},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 181, col: 1, offset: 5752},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 181, col: 1, offset: 5752},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 181, col: 1, offset: 5752},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 181, col: 3, offset: 5754},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 9, offset: 5760},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 181, col: 13, offset: 5764},
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 14, offset: 5765},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 185, col: 1, offset: 5863},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 185, col: 1, offset: 5863},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 185, col: 1, offset: 5863},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 185, col: 3, offset: 5865},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 9, offset: 5871},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 185, col: 13, offset: 5875},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 185, col: 15, offset: 5877},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 26, offset: 5888},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 185, col: 29, offset: 5891},
									expr: &litMatcher{
										pos:        position{line: 185, col: 30, offset: 5892},
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
			pos:  position{line: 189, col: 1, offset: 5986},
			expr: &actionExpr{
				pos: position{line: 189, col: 12, offset: 5997},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 189, col: 12, offset: 5997},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 189, col: 12, offset: 5997},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 14, offset: 5999},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 20, offset: 6005},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 24, offset: 6009},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 26, offset: 6011},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 39, offset: 6024},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 189, col: 42, offset: 6027},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 46, offset: 6031},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 6034},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 53, offset: 6038},
								expr: &seqExpr{
									pos: position{line: 189, col: 54, offset: 6039},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 189, col: 54, offset: 6039},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 63, offset: 6048},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 67, offset: 6052},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 71, offset: 6056},
								expr: &seqExpr{
									pos: position{line: 189, col: 72, offset: 6057},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 189, col: 72, offset: 6057},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 189, col: 74, offset: 6059},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 79, offset: 6064},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 81, offset: 6066},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 96, offset: 6081},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 189, col: 100, offset: 6085},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 104, offset: 6089},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 107, offset: 6092},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 189, col: 118, offset: 6103},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 119, offset: 6104},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 131, offset: 6116},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 133, offset: 6118},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 137, offset: 6122},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 215, col: 1, offset: 6717},
			expr: &actionExpr{
				pos: position{line: 215, col: 8, offset: 6724},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 215, col: 8, offset: 6724},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 215, col: 12, offset: 6728},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 215, col: 12, offset: 6728},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 215, col: 21, offset: 6737},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 221, col: 1, offset: 6854},
			expr: &choiceExpr{
				pos: position{line: 221, col: 10, offset: 6863},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 221, col: 10, offset: 6863},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 221, col: 10, offset: 6863},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 221, col: 10, offset: 6863},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 221, col: 12, offset: 6865},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 17, offset: 6870},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 21, offset: 6874},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 26, offset: 6879},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 36, offset: 6889},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 39, offset: 6892},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 43, offset: 6896},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 45, offset: 6898},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 51, offset: 6904},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 52, offset: 6905},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 64, offset: 6917},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 67, offset: 6920},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 71, offset: 6924},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 74, offset: 6927},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 81, offset: 6934},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 84, offset: 6937},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 88, offset: 6941},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 90, offset: 6943},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 96, offset: 6949},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 97, offset: 6950},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 109, offset: 6962},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 112, offset: 6965},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 240, col: 1, offset: 7468},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 240, col: 1, offset: 7468},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 240, col: 1, offset: 7468},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 240, col: 3, offset: 7470},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 8, offset: 7475},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 240, col: 12, offset: 7479},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 17, offset: 7484},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 27, offset: 7494},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 240, col: 30, offset: 7497},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 34, offset: 7501},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 240, col: 36, offset: 7503},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 240, col: 42, offset: 7509},
										expr: &ruleRefExpr{
											pos:  position{line: 240, col: 43, offset: 7510},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 55, offset: 7522},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 240, col: 57, offset: 7524},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 61, offset: 7528},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 240, col: 64, offset: 7531},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 240, col: 71, offset: 7538},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 79, offset: 7546},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 252, col: 1, offset: 7876},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 252, col: 1, offset: 7876},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 252, col: 1, offset: 7876},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 252, col: 3, offset: 7878},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 8, offset: 7883},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 252, col: 12, offset: 7887},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 17, offset: 7892},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 27, offset: 7902},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 252, col: 30, offset: 7905},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 34, offset: 7909},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 252, col: 36, offset: 7911},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 252, col: 42, offset: 7917},
										expr: &ruleRefExpr{
											pos:  position{line: 252, col: 43, offset: 7918},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 55, offset: 7930},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 252, col: 58, offset: 7933},
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
			pos:  position{line: 264, col: 1, offset: 8231},
			expr: &choiceExpr{
				pos: position{line: 264, col: 8, offset: 8238},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 264, col: 8, offset: 8238},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 264, col: 8, offset: 8238},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 8, offset: 8238},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 10, offset: 8240},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 17, offset: 8247},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 28, offset: 8258},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 264, col: 32, offset: 8262},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 35, offset: 8265},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 264, col: 48, offset: 8278},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 53, offset: 8283},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 278, col: 1, offset: 8607},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 278, col: 1, offset: 8607},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 278, col: 1, offset: 8607},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 278, col: 3, offset: 8609},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 6, offset: 8612},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 278, col: 19, offset: 8625},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 24, offset: 8630},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 292, col: 1, offset: 8947},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 292, col: 1, offset: 8947},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 292, col: 1, offset: 8947},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 292, col: 3, offset: 8949},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 292, col: 6, offset: 8952},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 292, col: 19, offset: 8965},
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
			pos:  position{line: 299, col: 1, offset: 9136},
			expr: &actionExpr{
				pos: position{line: 299, col: 16, offset: 9151},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 299, col: 16, offset: 9151},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 299, col: 16, offset: 9151},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 299, col: 23, offset: 9158},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 299, col: 36, offset: 9171},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 299, col: 41, offset: 9176},
								expr: &seqExpr{
									pos: position{line: 299, col: 42, offset: 9177},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 299, col: 42, offset: 9177},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 299, col: 46, offset: 9181},
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
			name: "ArgsDefn",
			pos:  position{line: 316, col: 1, offset: 9618},
			expr: &actionExpr{
				pos: position{line: 316, col: 12, offset: 9629},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 316, col: 12, offset: 9629},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 316, col: 12, offset: 9629},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 16, offset: 9633},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 316, col: 18, offset: 9635},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 316, col: 27, offset: 9644},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 35, offset: 9652},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 316, col: 37, offset: 9654},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 316, col: 42, offset: 9659},
								expr: &seqExpr{
									pos: position{line: 316, col: 43, offset: 9660},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 316, col: 43, offset: 9660},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 316, col: 47, offset: 9664},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 316, col: 49, offset: 9666},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 59, offset: 9676},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 316, col: 61, offset: 9678},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 334, col: 1, offset: 10100},
			expr: &actionExpr{
				pos: position{line: 334, col: 11, offset: 10110},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 334, col: 11, offset: 10110},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 334, col: 11, offset: 10110},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 16, offset: 10115},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 27, offset: 10126},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 29, offset: 10128},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 334, col: 34, offset: 10133},
								expr: &seqExpr{
									pos: position{line: 334, col: 35, offset: 10134},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 334, col: 35, offset: 10134},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 39, offset: 10138},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 41, offset: 10140},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 59, offset: 10158},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 355, col: 1, offset: 10697},
			expr: &choiceExpr{
				pos: position{line: 355, col: 18, offset: 10714},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 355, col: 18, offset: 10714},
						name: "OptionalPointerName",
					},
					&ruleRefExpr{
						pos:  position{line: 355, col: 40, offset: 10736},
						name: "AnyType",
					},
					&actionExpr{
						pos: position{line: 356, col: 1, offset: 10747},
						run: (*parser).callonTypeAnnotation4,
						expr: &seqExpr{
							pos: position{line: 356, col: 1, offset: 10747},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 356, col: 1, offset: 10747},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 8, offset: 10754},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 356, col: 11, offset: 10757},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 356, col: 16, offset: 10762},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 25, offset: 10771},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 356, col: 27, offset: 10773},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 32, offset: 10778},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 356, col: 34, offset: 10780},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 356, col: 38, offset: 10784},
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
			pos:  position{line: 365, col: 1, offset: 11000},
			expr: &choiceExpr{
				pos: position{line: 365, col: 11, offset: 11010},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 365, col: 11, offset: 11010},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 365, col: 22, offset: 11021},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 367, col: 1, offset: 11036},
			expr: &choiceExpr{
				pos: position{line: 367, col: 13, offset: 11048},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 367, col: 13, offset: 11048},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 367, col: 13, offset: 11048},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 367, col: 13, offset: 11048},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 17, offset: 11052},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 19, offset: 11054},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 367, col: 28, offset: 11063},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 40, offset: 11075},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 42, offset: 11077},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 367, col: 47, offset: 11082},
										expr: &seqExpr{
											pos: position{line: 367, col: 48, offset: 11083},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 367, col: 48, offset: 11083},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 367, col: 52, offset: 11087},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 367, col: 54, offset: 11089},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 68, offset: 11103},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 367, col: 70, offset: 11105},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 1, offset: 11527},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 384, col: 1, offset: 11527},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 384, col: 1, offset: 11527},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 384, col: 5, offset: 11531},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 384, col: 7, offset: 11533},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 384, col: 16, offset: 11542},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 384, col: 21, offset: 11547},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 384, col: 23, offset: 11549},
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
			pos:  position{line: 389, col: 1, offset: 11654},
			expr: &actionExpr{
				pos: position{line: 389, col: 16, offset: 11669},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 389, col: 16, offset: 11669},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 389, col: 16, offset: 11669},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 389, col: 18, offset: 11671},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 389, col: 21, offset: 11674},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 389, col: 27, offset: 11680},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 389, col: 32, offset: 11685},
								expr: &seqExpr{
									pos: position{line: 389, col: 33, offset: 11686},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 389, col: 33, offset: 11686},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 37, offset: 11690},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 46, offset: 11699},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 50, offset: 11703},
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
			pos:  position{line: 409, col: 1, offset: 12309},
			expr: &choiceExpr{
				pos: position{line: 409, col: 9, offset: 12317},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 409, col: 9, offset: 12317},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 21, offset: 12329},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 37, offset: 12345},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 48, offset: 12356},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 60, offset: 12368},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 411, col: 1, offset: 12381},
			expr: &actionExpr{
				pos: position{line: 411, col: 13, offset: 12393},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 411, col: 13, offset: 12393},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 411, col: 13, offset: 12393},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 411, col: 15, offset: 12395},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 411, col: 21, offset: 12401},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 411, col: 35, offset: 12415},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 411, col: 40, offset: 12420},
								expr: &seqExpr{
									pos: position{line: 411, col: 41, offset: 12421},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 411, col: 41, offset: 12421},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 45, offset: 12425},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 61, offset: 12441},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 65, offset: 12445},
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
			pos:  position{line: 444, col: 1, offset: 13338},
			expr: &actionExpr{
				pos: position{line: 444, col: 17, offset: 13354},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 444, col: 17, offset: 13354},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 444, col: 17, offset: 13354},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 444, col: 19, offset: 13356},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 25, offset: 13362},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 444, col: 34, offset: 13371},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 444, col: 39, offset: 13376},
								expr: &seqExpr{
									pos: position{line: 444, col: 40, offset: 13377},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 444, col: 40, offset: 13377},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 44, offset: 13381},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 61, offset: 13398},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 65, offset: 13402},
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
			pos:  position{line: 476, col: 1, offset: 14289},
			expr: &actionExpr{
				pos: position{line: 476, col: 12, offset: 14300},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 476, col: 12, offset: 14300},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 476, col: 12, offset: 14300},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 476, col: 14, offset: 14302},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 476, col: 20, offset: 14308},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 476, col: 30, offset: 14318},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 476, col: 35, offset: 14323},
								expr: &seqExpr{
									pos: position{line: 476, col: 36, offset: 14324},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 476, col: 36, offset: 14324},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 40, offset: 14328},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 52, offset: 14340},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 56, offset: 14344},
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
			pos:  position{line: 508, col: 1, offset: 15232},
			expr: &actionExpr{
				pos: position{line: 508, col: 13, offset: 15244},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 508, col: 13, offset: 15244},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 508, col: 13, offset: 15244},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 508, col: 15, offset: 15246},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 21, offset: 15252},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 508, col: 33, offset: 15264},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 508, col: 38, offset: 15269},
								expr: &seqExpr{
									pos: position{line: 508, col: 39, offset: 15270},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 508, col: 39, offset: 15270},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 43, offset: 15274},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 56, offset: 15287},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 60, offset: 15291},
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
			pos:  position{line: 539, col: 1, offset: 16180},
			expr: &choiceExpr{
				pos: position{line: 539, col: 15, offset: 16194},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 539, col: 15, offset: 16194},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 539, col: 15, offset: 16194},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 539, col: 15, offset: 16194},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 539, col: 17, offset: 16196},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 21, offset: 16200},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 539, col: 24, offset: 16203},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 539, col: 30, offset: 16209},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 36, offset: 16215},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 539, col: 39, offset: 16218},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 542, col: 5, offset: 16341},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 544, col: 1, offset: 16348},
			expr: &choiceExpr{
				pos: position{line: 544, col: 12, offset: 16359},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 544, col: 12, offset: 16359},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 30, offset: 16377},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 49, offset: 16396},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 64, offset: 16411},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 546, col: 1, offset: 16424},
			expr: &actionExpr{
				pos: position{line: 546, col: 19, offset: 16442},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 546, col: 21, offset: 16444},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 546, col: 21, offset: 16444},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 546, col: 28, offset: 16451},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 550, col: 1, offset: 16533},
			expr: &actionExpr{
				pos: position{line: 550, col: 20, offset: 16552},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 550, col: 22, offset: 16554},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 550, col: 22, offset: 16554},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 29, offset: 16561},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 36, offset: 16568},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 42, offset: 16574},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 48, offset: 16580},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 56, offset: 16588},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 554, col: 1, offset: 16667},
			expr: &choiceExpr{
				pos: position{line: 554, col: 16, offset: 16682},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 554, col: 16, offset: 16682},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 554, col: 18, offset: 16684},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 554, col: 18, offset: 16684},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 554, col: 24, offset: 16690},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 557, col: 3, offset: 16773},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 557, col: 5, offset: 16775},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 560, col: 3, offset: 16855},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 560, col: 3, offset: 16855},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 564, col: 1, offset: 16936},
			expr: &actionExpr{
				pos: position{line: 564, col: 15, offset: 16950},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 564, col: 17, offset: 16952},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 564, col: 17, offset: 16952},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 564, col: 23, offset: 16958},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 568, col: 1, offset: 17040},
			expr: &choiceExpr{
				pos: position{line: 568, col: 9, offset: 17048},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 568, col: 9, offset: 17048},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 16, offset: 17055},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 31, offset: 17070},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 46, offset: 17085},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 570, col: 1, offset: 17092},
			expr: &choiceExpr{
				pos: position{line: 570, col: 14, offset: 17105},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 570, col: 14, offset: 17105},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 570, col: 29, offset: 17120},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 572, col: 1, offset: 17128},
			expr: &choiceExpr{
				pos: position{line: 572, col: 14, offset: 17141},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 572, col: 14, offset: 17141},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 572, col: 29, offset: 17156},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 574, col: 1, offset: 17168},
			expr: &actionExpr{
				pos: position{line: 574, col: 16, offset: 17183},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 574, col: 16, offset: 17183},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 574, col: 16, offset: 17183},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 20, offset: 17187},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 574, col: 22, offset: 17189},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 574, col: 28, offset: 17195},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 33, offset: 17200},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 574, col: 35, offset: 17202},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 574, col: 40, offset: 17207},
								expr: &seqExpr{
									pos: position{line: 574, col: 41, offset: 17208},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 574, col: 41, offset: 17208},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 45, offset: 17212},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 47, offset: 17214},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 52, offset: 17219},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 574, col: 56, offset: 17223},
							expr: &litMatcher{
								pos:        position{line: 574, col: 56, offset: 17223},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 61, offset: 17228},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 574, col: 63, offset: 17230},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 590, col: 1, offset: 17675},
			expr: &actionExpr{
				pos: position{line: 590, col: 19, offset: 17693},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 590, col: 19, offset: 17693},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 590, col: 19, offset: 17693},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 24, offset: 17698},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 590, col: 35, offset: 17709},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 590, col: 37, offset: 17711},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 42, offset: 17716},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 603, col: 1, offset: 17986},
			expr: &actionExpr{
				pos: position{line: 603, col: 18, offset: 18003},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 603, col: 18, offset: 18003},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 603, col: 18, offset: 18003},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 23, offset: 18008},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 34, offset: 18019},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 603, col: 36, offset: 18021},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 40, offset: 18025},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 603, col: 42, offset: 18027},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 52, offset: 18037},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 65, offset: 18050},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 603, col: 67, offset: 18052},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 71, offset: 18056},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 603, col: 73, offset: 18058},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 84, offset: 18069},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 603, col: 89, offset: 18074},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 603, col: 94, offset: 18079},
								expr: &seqExpr{
									pos: position{line: 603, col: 95, offset: 18080},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 603, col: 95, offset: 18080},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 99, offset: 18084},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 101, offset: 18086},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 114, offset: 18099},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 603, col: 116, offset: 18101},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 120, offset: 18105},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 122, offset: 18107},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 603, col: 130, offset: 18115},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 623, col: 1, offset: 18699},
			expr: &actionExpr{
				pos: position{line: 623, col: 17, offset: 18715},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 623, col: 17, offset: 18715},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 623, col: 17, offset: 18715},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 623, col: 22, offset: 18720},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 627, col: 1, offset: 18793},
			expr: &actionExpr{
				pos: position{line: 627, col: 16, offset: 18808},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 627, col: 16, offset: 18808},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 627, col: 16, offset: 18808},
							expr: &ruleRefExpr{
								pos:  position{line: 627, col: 17, offset: 18809},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 627, col: 27, offset: 18819},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 627, col: 27, offset: 18819},
									expr: &charClassMatcher{
										pos:        position{line: 627, col: 27, offset: 18819},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 627, col: 34, offset: 18826},
									expr: &charClassMatcher{
										pos:        position{line: 627, col: 34, offset: 18826},
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
			pos:  position{line: 631, col: 1, offset: 18901},
			expr: &actionExpr{
				pos: position{line: 631, col: 14, offset: 18914},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 631, col: 15, offset: 18915},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 631, col: 15, offset: 18915},
							expr: &charClassMatcher{
								pos:        position{line: 631, col: 15, offset: 18915},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 631, col: 22, offset: 18922},
							expr: &charClassMatcher{
								pos:        position{line: 631, col: 22, offset: 18922},
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
			name: "PointerType",
			pos:  position{line: 635, col: 1, offset: 18997},
			expr: &actionExpr{
				pos: position{line: 635, col: 15, offset: 19011},
				run: (*parser).callonPointerType1,
				expr: &seqExpr{
					pos: position{line: 635, col: 15, offset: 19011},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 635, col: 15, offset: 19011},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 635, col: 19, offset: 19015},
							label: "m",
							expr: &ruleRefExpr{
								pos:  position{line: 635, col: 21, offset: 19017},
								name: "ModuleName",
							},
						},
					},
				},
			},
		},
		{
			name: "OptionalPointerName",
			pos:  position{line: 639, col: 1, offset: 19088},
			expr: &choiceExpr{
				pos: position{line: 639, col: 23, offset: 19110},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 639, col: 23, offset: 19110},
						name: "ModuleName",
					},
					&ruleRefExpr{
						pos:  position{line: 639, col: 36, offset: 19123},
						name: "PointerType",
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 641, col: 1, offset: 19136},
			expr: &choiceExpr{
				pos: position{line: 641, col: 9, offset: 19144},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 641, col: 9, offset: 19144},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 641, col: 9, offset: 19144},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 641, col: 9, offset: 19144},
									expr: &litMatcher{
										pos:        position{line: 641, col: 9, offset: 19144},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 641, col: 14, offset: 19149},
									expr: &charClassMatcher{
										pos:        position{line: 641, col: 14, offset: 19149},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 641, col: 21, offset: 19156},
									expr: &litMatcher{
										pos:        position{line: 641, col: 22, offset: 19157},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 648, col: 3, offset: 19332},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 648, col: 3, offset: 19332},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 648, col: 3, offset: 19332},
									expr: &litMatcher{
										pos:        position{line: 648, col: 3, offset: 19332},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 648, col: 8, offset: 19337},
									expr: &charClassMatcher{
										pos:        position{line: 648, col: 8, offset: 19337},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 648, col: 15, offset: 19344},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 648, col: 19, offset: 19348},
									expr: &charClassMatcher{
										pos:        position{line: 648, col: 19, offset: 19348},
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
						pos: position{line: 655, col: 3, offset: 19537},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 655, col: 3, offset: 19537},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 659, col: 3, offset: 19622},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 659, col: 3, offset: 19622},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 662, col: 3, offset: 19708},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 663, col: 3, offset: 19715},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 664, col: 3, offset: 19731},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 664, col: 3, offset: 19731},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 664, col: 3, offset: 19731},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 664, col: 7, offset: 19735},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 664, col: 12, offset: 19740},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 664, col: 12, offset: 19740},
												expr: &ruleRefExpr{
													pos:  position{line: 664, col: 13, offset: 19741},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 664, col: 25, offset: 19753,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 664, col: 28, offset: 19756},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 666, col: 5, offset: 19848},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 666, col: 20, offset: 19863},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 666, col: 37, offset: 19880},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 668, col: 1, offset: 19897},
			expr: &actionExpr{
				pos: position{line: 668, col: 8, offset: 19904},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 668, col: 8, offset: 19904},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 672, col: 1, offset: 19967},
			expr: &actionExpr{
				pos: position{line: 672, col: 17, offset: 19983},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 672, col: 17, offset: 19983},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 672, col: 17, offset: 19983},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 672, col: 21, offset: 19987},
							expr: &seqExpr{
								pos: position{line: 672, col: 22, offset: 19988},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 672, col: 22, offset: 19988},
										expr: &ruleRefExpr{
											pos:  position{line: 672, col: 23, offset: 19989},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 672, col: 35, offset: 20001,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 672, col: 39, offset: 20005},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 680, col: 1, offset: 20188},
			expr: &actionExpr{
				pos: position{line: 680, col: 10, offset: 20197},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 680, col: 11, offset: 20198},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 684, col: 1, offset: 20253},
			expr: &seqExpr{
				pos: position{line: 684, col: 12, offset: 20264},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 684, col: 13, offset: 20265},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 684, col: 13, offset: 20265},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 21, offset: 20273},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 28, offset: 20280},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 37, offset: 20289},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 48, offset: 20300},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 57, offset: 20309},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 66, offset: 20318},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 76, offset: 20328},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 684, col: 88, offset: 20340},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 684, col: 97, offset: 20349},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 684, col: 107, offset: 20359},
						expr: &oneOrMoreExpr{
							pos: position{line: 684, col: 108, offset: 20360},
							expr: &charClassMatcher{
								pos:        position{line: 684, col: 108, offset: 20360},
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
			pos:  position{line: 686, col: 1, offset: 20368},
			expr: &choiceExpr{
				pos: position{line: 686, col: 12, offset: 20379},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 686, col: 12, offset: 20379},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 686, col: 14, offset: 20381},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 686, col: 14, offset: 20381},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 24, offset: 20391},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 32, offset: 20399},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 41, offset: 20408},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 52, offset: 20419},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 61, offset: 20428},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 70, offset: 20437},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 686, col: 80, offset: 20447},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 689, col: 3, offset: 20545},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 691, col: 1, offset: 20551},
			expr: &charClassMatcher{
				pos:        position{line: 691, col: 15, offset: 20565},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 693, col: 1, offset: 20581},
			expr: &choiceExpr{
				pos: position{line: 693, col: 18, offset: 20598},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 693, col: 18, offset: 20598},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 693, col: 37, offset: 20617},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 695, col: 1, offset: 20632},
			expr: &charClassMatcher{
				pos:        position{line: 695, col: 20, offset: 20651},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 697, col: 1, offset: 20664},
			expr: &charClassMatcher{
				pos:        position{line: 697, col: 16, offset: 20679},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 699, col: 1, offset: 20686},
			expr: &charClassMatcher{
				pos:        position{line: 699, col: 23, offset: 20708},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 701, col: 1, offset: 20715},
			expr: &charClassMatcher{
				pos:        position{line: 701, col: 12, offset: 20726},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 703, col: 1, offset: 20737},
			expr: &choiceExpr{
				pos: position{line: 703, col: 22, offset: 20758},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 703, col: 22, offset: 20758},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 703, col: 33, offset: 20769},
						expr: &charClassMatcher{
							pos:        position{line: 703, col: 33, offset: 20769},
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
			pos:         position{line: 705, col: 1, offset: 20781},
			expr: &choiceExpr{
				pos: position{line: 705, col: 21, offset: 20801},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 705, col: 21, offset: 20801},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 705, col: 32, offset: 20812},
						expr: &charClassMatcher{
							pos:        position{line: 705, col: 32, offset: 20812},
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
			pos:         position{line: 707, col: 1, offset: 20824},
			expr: &oneOrMoreExpr{
				pos: position{line: 707, col: 33, offset: 20856},
				expr: &charClassMatcher{
					pos:        position{line: 707, col: 33, offset: 20856},
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
			pos:         position{line: 709, col: 1, offset: 20864},
			expr: &zeroOrMoreExpr{
				pos: position{line: 709, col: 32, offset: 20895},
				expr: &charClassMatcher{
					pos:        position{line: 709, col: 32, offset: 20895},
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
			pos:         position{line: 711, col: 1, offset: 20903},
			expr: &choiceExpr{
				pos: position{line: 711, col: 15, offset: 20917},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 711, col: 15, offset: 20917},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 711, col: 26, offset: 20928},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 711, col: 26, offset: 20928},
								expr: &charClassMatcher{
									pos:        position{line: 711, col: 26, offset: 20928},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 711, col: 35, offset: 20937},
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
			pos:  position{line: 713, col: 1, offset: 20943},
			expr: &oneOrMoreExpr{
				pos: position{line: 713, col: 12, offset: 20954},
				expr: &ruleRefExpr{
					pos:  position{line: 713, col: 13, offset: 20955},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 715, col: 1, offset: 20966},
			expr: &choiceExpr{
				pos: position{line: 715, col: 11, offset: 20976},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 715, col: 11, offset: 20976},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 715, col: 11, offset: 20976},
								expr: &charClassMatcher{
									pos:        position{line: 715, col: 11, offset: 20976},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 715, col: 22, offset: 20987},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 715, col: 27, offset: 20992},
								expr: &seqExpr{
									pos: position{line: 715, col: 28, offset: 20993},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 715, col: 28, offset: 20993},
											expr: &charClassMatcher{
												pos:        position{line: 715, col: 29, offset: 20994},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 715, col: 34, offset: 20999,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 715, col: 38, offset: 21003},
								expr: &litMatcher{
									pos:        position{line: 715, col: 39, offset: 21004},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 715, col: 46, offset: 21011},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 715, col: 46, offset: 21011},
								expr: &charClassMatcher{
									pos:        position{line: 715, col: 46, offset: 21011},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 715, col: 57, offset: 21022},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 715, col: 62, offset: 21027},
								expr: &seqExpr{
									pos: position{line: 715, col: 63, offset: 21028},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 715, col: 63, offset: 21028},
											expr: &litMatcher{
												pos:        position{line: 715, col: 64, offset: 21029},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 715, col: 69, offset: 21034,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 715, col: 73, offset: 21038},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 715, col: 78, offset: 21043},
								expr: &charClassMatcher{
									pos:        position{line: 715, col: 78, offset: 21043},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 715, col: 84, offset: 21049},
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
			pos:  position{line: 717, col: 1, offset: 21055},
			expr: &notExpr{
				pos: position{line: 717, col: 7, offset: 21061},
				expr: &anyMatcher{
					line: 717, col: 8, offset: 21062,
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

func (c *current) onPointerType1(m interface{}) (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
}

func (p *parser) callonPointerType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPointerType1(stack["m"])
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
