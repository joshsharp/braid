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
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 55, offset: 1571},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 57, offset: 1573},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 61, offset: 1577},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 5, offset: 1583},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 16, offset: 1594},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 30, offset: 1608},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 32, offset: 1610},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 36, offset: 1614},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 38, offset: 1616},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 44, offset: 1622},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 60, offset: 1638},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 62, offset: 1640},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 54, col: 67, offset: 1645},
										expr: &seqExpr{
											pos: position{line: 54, col: 68, offset: 1646},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 54, col: 68, offset: 1646},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 72, offset: 1650},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 74, offset: 1652},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 90, offset: 1668},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 54, col: 94, offset: 1672},
									expr: &litMatcher{
										pos:        position{line: 54, col: 94, offset: 1672},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 99, offset: 1677},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 101, offset: 1679},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 105, offset: 1683},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 73, col: 1, offset: 2232},
						run: (*parser).callonExternType34,
						expr: &seqExpr{
							pos: position{line: 73, col: 1, offset: 2232},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 73, col: 1, offset: 2232},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 3, offset: 2234},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 12, offset: 2243},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 73, col: 16, offset: 2247},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 23, offset: 2254},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 26, offset: 2257},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 73, col: 31, offset: 2262},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 42, offset: 2273},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 44, offset: 2275},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 48, offset: 2279},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 5, offset: 2285},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 74, col: 16, offset: 2296},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 30, offset: 2310},
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
			pos:  position{line: 82, col: 1, offset: 2504},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2515},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 82, col: 12, offset: 2515},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 82, col: 12, offset: 2515},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 82, col: 12, offset: 2515},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 19, offset: 2522},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 22, offset: 2525},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 27, offset: 2530},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 38, offset: 2541},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 40, offset: 2543},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 47, offset: 2550},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 58, offset: 2561},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 60, offset: 2563},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 82, col: 64, offset: 2567},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 82, col: 70, offset: 2573},
										expr: &seqExpr{
											pos: position{line: 82, col: 71, offset: 2574},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 71, offset: 2574},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 74, offset: 2577},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 91, offset: 2594},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 100, col: 1, offset: 3081},
						run: (*parser).callonTypeDefn19,
						expr: &seqExpr{
							pos: position{line: 100, col: 1, offset: 3081},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 100, col: 1, offset: 3081},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 8, offset: 3088},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 11, offset: 3091},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 16, offset: 3096},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 27, offset: 3107},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 29, offset: 3109},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 36, offset: 3116},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 47, offset: 3127},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 49, offset: 3129},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 53, offset: 3133},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 5, offset: 3139},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 9, offset: 3143},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 11, offset: 3145},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 17, offset: 3151},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 33, offset: 3167},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 35, offset: 3169},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 40, offset: 3174},
										expr: &seqExpr{
											pos: position{line: 101, col: 41, offset: 3175},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 101, col: 41, offset: 3175},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 45, offset: 3179},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 47, offset: 3181},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 63, offset: 3197},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 101, col: 67, offset: 3201},
									expr: &litMatcher{
										pos:        position{line: 101, col: 67, offset: 3201},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 72, offset: 3206},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 74, offset: 3208},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 78, offset: 3212},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 119, col: 1, offset: 3696},
						run: (*parser).callonTypeDefn48,
						expr: &seqExpr{
							pos: position{line: 119, col: 1, offset: 3696},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 119, col: 1, offset: 3696},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 8, offset: 3703},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 11, offset: 3706},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 119, col: 16, offset: 3711},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 27, offset: 3722},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 29, offset: 3724},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 33, offset: 3728},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 5, offset: 3734},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 9, offset: 3738},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 11, offset: 3740},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 120, col: 17, offset: 3746},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 33, offset: 3762},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 35, offset: 3764},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 120, col: 40, offset: 3769},
										expr: &seqExpr{
											pos: position{line: 120, col: 41, offset: 3770},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 120, col: 41, offset: 3770},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 45, offset: 3774},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 47, offset: 3776},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 63, offset: 3792},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 120, col: 67, offset: 3796},
									expr: &litMatcher{
										pos:        position{line: 120, col: 67, offset: 3796},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 72, offset: 3801},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 74, offset: 3803},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 78, offset: 3807},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 138, col: 1, offset: 4307},
						run: (*parser).callonTypeDefn74,
						expr: &seqExpr{
							pos: position{line: 138, col: 1, offset: 4307},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 138, col: 1, offset: 4307},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 8, offset: 4314},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 11, offset: 4317},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 16, offset: 4322},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 27, offset: 4333},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 29, offset: 4335},
									label: "params",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 36, offset: 4342},
										name: "TypeParams",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 47, offset: 4353},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 49, offset: 4355},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 53, offset: 4359},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 55, offset: 4361},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 138, col: 60, offset: 4366},
										expr: &ruleRefExpr{
											pos:  position{line: 138, col: 61, offset: 4367},
											name: "VariantConstructor",
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 152, col: 1, offset: 4767},
						run: (*parser).callonTypeDefn89,
						expr: &seqExpr{
							pos: position{line: 152, col: 1, offset: 4767},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 152, col: 1, offset: 4767},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 8, offset: 4774},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 152, col: 11, offset: 4777},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 152, col: 16, offset: 4782},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 27, offset: 4793},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 152, col: 29, offset: 4795},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 152, col: 33, offset: 4799},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 152, col: 35, offset: 4801},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 152, col: 40, offset: 4806},
										expr: &ruleRefExpr{
											pos:  position{line: 152, col: 41, offset: 4807},
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
			name: "TypeParams",
			pos:  position{line: 167, col: 1, offset: 5214},
			expr: &actionExpr{
				pos: position{line: 167, col: 14, offset: 5227},
				run: (*parser).callonTypeParams1,
				expr: &seqExpr{
					pos: position{line: 167, col: 14, offset: 5227},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 167, col: 14, offset: 5227},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 167, col: 18, offset: 5231},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 24, offset: 5237},
								name: "TypeParameter",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 167, col: 38, offset: 5251},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 167, col: 40, offset: 5253},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 167, col: 45, offset: 5258},
								expr: &seqExpr{
									pos: position{line: 167, col: 46, offset: 5259},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 167, col: 46, offset: 5259},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 50, offset: 5263},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 52, offset: 5265},
											name: "TypeParameter",
										},
										&ruleRefExpr{
											pos:  position{line: 167, col: 66, offset: 5279},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 167, col: 70, offset: 5283},
							expr: &litMatcher{
								pos:        position{line: 167, col: 70, offset: 5283},
								val:        ",",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 167, col: 75, offset: 5288},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 183, col: 1, offset: 5699},
			expr: &actionExpr{
				pos: position{line: 183, col: 19, offset: 5717},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 183, col: 19, offset: 5717},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 19, offset: 5717},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 24, offset: 5722},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 37, offset: 5735},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 183, col: 39, offset: 5737},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 43, offset: 5741},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 183, col: 45, offset: 5743},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 48, offset: 5746},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 187, col: 1, offset: 5847},
			expr: &choiceExpr{
				pos: position{line: 187, col: 22, offset: 5868},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 187, col: 22, offset: 5868},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 187, col: 22, offset: 5868},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 187, col: 22, offset: 5868},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 24, offset: 5870},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 28, offset: 5874},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 30, offset: 5876},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 35, offset: 5881},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 46, offset: 5892},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 48, offset: 5894},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 52, offset: 5898},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 54, offset: 5900},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 60, offset: 5906},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 76, offset: 5922},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 78, offset: 5924},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 187, col: 83, offset: 5929},
										expr: &seqExpr{
											pos: position{line: 187, col: 84, offset: 5930},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 187, col: 84, offset: 5930},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 187, col: 88, offset: 5934},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 187, col: 90, offset: 5936},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 187, col: 106, offset: 5952},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 187, col: 110, offset: 5956},
									expr: &litMatcher{
										pos:        position{line: 187, col: 110, offset: 5956},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 115, offset: 5961},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 117, offset: 5963},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 121, offset: 5967},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 206, col: 1, offset: 6571},
						run: (*parser).callonVariantConstructor27,
						expr: &seqExpr{
							pos: position{line: 206, col: 1, offset: 6571},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 206, col: 1, offset: 6571},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 206, col: 3, offset: 6573},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 7, offset: 6577},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 206, col: 9, offset: 6579},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 206, col: 14, offset: 6584},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 206, col: 25, offset: 6595},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 206, col: 30, offset: 6600},
										expr: &seqExpr{
											pos: position{line: 206, col: 31, offset: 6601},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 206, col: 31, offset: 6601},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 206, col: 34, offset: 6604},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 51, offset: 6621},
									name: "N",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 223, col: 1, offset: 7058},
			expr: &choiceExpr{
				pos: position{line: 223, col: 14, offset: 7071},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 223, col: 14, offset: 7071},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 223, col: 14, offset: 7071},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 223, col: 14, offset: 7071},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 223, col: 16, offset: 7073},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 22, offset: 7079},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 223, col: 26, offset: 7083},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 223, col: 28, offset: 7085},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 39, offset: 7096},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 223, col: 42, offset: 7099},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 46, offset: 7103},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 223, col: 49, offset: 7106},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 223, col: 54, offset: 7111},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 59, offset: 7116},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 229, col: 1, offset: 7235},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 229, col: 1, offset: 7235},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 229, col: 1, offset: 7235},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 229, col: 3, offset: 7237},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 9, offset: 7243},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 229, col: 13, offset: 7247},
									expr: &ruleRefExpr{
										pos:  position{line: 229, col: 14, offset: 7248},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 233, col: 1, offset: 7346},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 233, col: 1, offset: 7346},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 233, col: 1, offset: 7346},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 233, col: 3, offset: 7348},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 233, col: 9, offset: 7354},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 233, col: 13, offset: 7358},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 233, col: 15, offset: 7360},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 233, col: 26, offset: 7371},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 233, col: 29, offset: 7374},
									expr: &litMatcher{
										pos:        position{line: 233, col: 30, offset: 7375},
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
			pos:  position{line: 237, col: 1, offset: 7469},
			expr: &actionExpr{
				pos: position{line: 237, col: 12, offset: 7480},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 237, col: 12, offset: 7480},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 237, col: 12, offset: 7480},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 237, col: 14, offset: 7482},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 20, offset: 7488},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 24, offset: 7492},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 26, offset: 7494},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 39, offset: 7507},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 237, col: 42, offset: 7510},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 46, offset: 7514},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 49, offset: 7517},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 237, col: 53, offset: 7521},
								expr: &seqExpr{
									pos: position{line: 237, col: 54, offset: 7522},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 237, col: 54, offset: 7522},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 63, offset: 7531},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 67, offset: 7535},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 237, col: 71, offset: 7539},
								expr: &seqExpr{
									pos: position{line: 237, col: 72, offset: 7540},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 237, col: 72, offset: 7540},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 237, col: 74, offset: 7542},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 79, offset: 7547},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 81, offset: 7549},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 96, offset: 7564},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 237, col: 100, offset: 7568},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 104, offset: 7572},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 107, offset: 7575},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 237, col: 118, offset: 7586},
								expr: &ruleRefExpr{
									pos:  position{line: 237, col: 119, offset: 7587},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 131, offset: 7599},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 237, col: 133, offset: 7601},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 137, offset: 7605},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 263, col: 1, offset: 8200},
			expr: &actionExpr{
				pos: position{line: 263, col: 8, offset: 8207},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 263, col: 8, offset: 8207},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 263, col: 12, offset: 8211},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 263, col: 12, offset: 8211},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 263, col: 21, offset: 8220},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 269, col: 1, offset: 8337},
			expr: &choiceExpr{
				pos: position{line: 269, col: 10, offset: 8346},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 269, col: 10, offset: 8346},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 269, col: 10, offset: 8346},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 269, col: 10, offset: 8346},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 269, col: 12, offset: 8348},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 17, offset: 8353},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 269, col: 21, offset: 8357},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 269, col: 26, offset: 8362},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 36, offset: 8372},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 269, col: 39, offset: 8375},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 43, offset: 8379},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 269, col: 45, offset: 8381},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 269, col: 51, offset: 8387},
										expr: &ruleRefExpr{
											pos:  position{line: 269, col: 52, offset: 8388},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 64, offset: 8400},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 269, col: 67, offset: 8403},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 71, offset: 8407},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 269, col: 74, offset: 8410},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 81, offset: 8417},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 269, col: 84, offset: 8420},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 88, offset: 8424},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 269, col: 90, offset: 8426},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 269, col: 96, offset: 8432},
										expr: &ruleRefExpr{
											pos:  position{line: 269, col: 97, offset: 8433},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 109, offset: 8445},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 269, col: 112, offset: 8448},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 288, col: 1, offset: 8951},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 288, col: 1, offset: 8951},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 288, col: 1, offset: 8951},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 288, col: 3, offset: 8953},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 8, offset: 8958},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 288, col: 12, offset: 8962},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 288, col: 17, offset: 8967},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 27, offset: 8977},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 288, col: 30, offset: 8980},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 34, offset: 8984},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 288, col: 36, offset: 8986},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 288, col: 42, offset: 8992},
										expr: &ruleRefExpr{
											pos:  position{line: 288, col: 43, offset: 8993},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 55, offset: 9005},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 288, col: 57, offset: 9007},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 61, offset: 9011},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 288, col: 64, offset: 9014},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 288, col: 71, offset: 9021},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 288, col: 79, offset: 9029},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 300, col: 1, offset: 9359},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 300, col: 1, offset: 9359},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 300, col: 1, offset: 9359},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 300, col: 3, offset: 9361},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 8, offset: 9366},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 300, col: 12, offset: 9370},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 17, offset: 9375},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 27, offset: 9385},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 300, col: 30, offset: 9388},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 34, offset: 9392},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 300, col: 36, offset: 9394},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 300, col: 42, offset: 9400},
										expr: &ruleRefExpr{
											pos:  position{line: 300, col: 43, offset: 9401},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 55, offset: 9413},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 300, col: 58, offset: 9416},
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
			pos:  position{line: 312, col: 1, offset: 9714},
			expr: &choiceExpr{
				pos: position{line: 312, col: 8, offset: 9721},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 312, col: 8, offset: 9721},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 312, col: 8, offset: 9721},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 312, col: 8, offset: 9721},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 312, col: 10, offset: 9723},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 17, offset: 9730},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 312, col: 28, offset: 9741},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 312, col: 32, offset: 9745},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 35, offset: 9748},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 312, col: 48, offset: 9761},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 53, offset: 9766},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 326, col: 1, offset: 10090},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 326, col: 1, offset: 10090},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 326, col: 1, offset: 10090},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 326, col: 3, offset: 10092},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 6, offset: 10095},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 326, col: 19, offset: 10108},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 24, offset: 10113},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 340, col: 1, offset: 10430},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 340, col: 1, offset: 10430},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 340, col: 1, offset: 10430},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 340, col: 3, offset: 10432},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 6, offset: 10435},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 340, col: 19, offset: 10448},
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
			pos:  position{line: 347, col: 1, offset: 10619},
			expr: &actionExpr{
				pos: position{line: 347, col: 16, offset: 10634},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 347, col: 16, offset: 10634},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 347, col: 16, offset: 10634},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 347, col: 23, offset: 10641},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 347, col: 36, offset: 10654},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 347, col: 41, offset: 10659},
								expr: &seqExpr{
									pos: position{line: 347, col: 42, offset: 10660},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 347, col: 42, offset: 10660},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 347, col: 46, offset: 10664},
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
			pos:  position{line: 364, col: 1, offset: 11101},
			expr: &actionExpr{
				pos: position{line: 364, col: 12, offset: 11112},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 364, col: 12, offset: 11112},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 364, col: 12, offset: 11112},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 16, offset: 11116},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 364, col: 18, offset: 11118},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 27, offset: 11127},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 35, offset: 11135},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 364, col: 37, offset: 11137},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 364, col: 42, offset: 11142},
								expr: &seqExpr{
									pos: position{line: 364, col: 43, offset: 11143},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 364, col: 43, offset: 11143},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 47, offset: 11147},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 49, offset: 11149},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 59, offset: 11159},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 364, col: 61, offset: 11161},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 382, col: 1, offset: 11583},
			expr: &actionExpr{
				pos: position{line: 382, col: 11, offset: 11593},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 382, col: 11, offset: 11593},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 382, col: 11, offset: 11593},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 382, col: 16, offset: 11598},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 382, col: 27, offset: 11609},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 382, col: 29, offset: 11611},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 382, col: 34, offset: 11616},
								expr: &seqExpr{
									pos: position{line: 382, col: 35, offset: 11617},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 382, col: 35, offset: 11617},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 382, col: 39, offset: 11621},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 382, col: 41, offset: 11623},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 382, col: 59, offset: 11641},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 403, col: 1, offset: 12180},
			expr: &choiceExpr{
				pos: position{line: 403, col: 18, offset: 12197},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 403, col: 18, offset: 12197},
						name: "AnyType",
					},
					&actionExpr{
						pos: position{line: 404, col: 1, offset: 12208},
						run: (*parser).callonTypeAnnotation3,
						expr: &seqExpr{
							pos: position{line: 404, col: 1, offset: 12208},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 404, col: 1, offset: 12208},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 404, col: 8, offset: 12215},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 404, col: 11, offset: 12218},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 404, col: 16, offset: 12223},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 404, col: 25, offset: 12232},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 404, col: 27, offset: 12234},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 404, col: 32, offset: 12239},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 404, col: 34, offset: 12241},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 404, col: 38, offset: 12245},
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
			pos:  position{line: 413, col: 1, offset: 12461},
			expr: &choiceExpr{
				pos: position{line: 413, col: 11, offset: 12471},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 413, col: 11, offset: 12471},
						name: "ModuleName",
					},
					&ruleRefExpr{
						pos:  position{line: 413, col: 24, offset: 12484},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 413, col: 35, offset: 12495},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 415, col: 1, offset: 12510},
			expr: &choiceExpr{
				pos: position{line: 415, col: 13, offset: 12522},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 415, col: 13, offset: 12522},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 415, col: 13, offset: 12522},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 415, col: 13, offset: 12522},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 415, col: 17, offset: 12526},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 415, col: 19, offset: 12528},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 415, col: 28, offset: 12537},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 415, col: 40, offset: 12549},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 415, col: 42, offset: 12551},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 415, col: 47, offset: 12556},
										expr: &seqExpr{
											pos: position{line: 415, col: 48, offset: 12557},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 415, col: 48, offset: 12557},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 415, col: 52, offset: 12561},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 415, col: 54, offset: 12563},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 415, col: 68, offset: 12577},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 415, col: 70, offset: 12579},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 432, col: 1, offset: 13001},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 432, col: 1, offset: 13001},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 432, col: 1, offset: 13001},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 432, col: 5, offset: 13005},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 432, col: 7, offset: 13007},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 432, col: 16, offset: 13016},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 432, col: 21, offset: 13021},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 432, col: 23, offset: 13023},
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
			pos:  position{line: 437, col: 1, offset: 13128},
			expr: &actionExpr{
				pos: position{line: 437, col: 16, offset: 13143},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 437, col: 16, offset: 13143},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 437, col: 16, offset: 13143},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 437, col: 18, offset: 13145},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 437, col: 21, offset: 13148},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 437, col: 27, offset: 13154},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 437, col: 32, offset: 13159},
								expr: &seqExpr{
									pos: position{line: 437, col: 33, offset: 13160},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 437, col: 33, offset: 13160},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 437, col: 37, offset: 13164},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 437, col: 46, offset: 13173},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 437, col: 50, offset: 13177},
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
			pos:  position{line: 457, col: 1, offset: 13783},
			expr: &choiceExpr{
				pos: position{line: 457, col: 9, offset: 13791},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 457, col: 9, offset: 13791},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 21, offset: 13803},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 37, offset: 13819},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 48, offset: 13830},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 60, offset: 13842},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 459, col: 1, offset: 13855},
			expr: &actionExpr{
				pos: position{line: 459, col: 13, offset: 13867},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 459, col: 13, offset: 13867},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 459, col: 13, offset: 13867},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 459, col: 15, offset: 13869},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 21, offset: 13875},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 459, col: 35, offset: 13889},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 459, col: 40, offset: 13894},
								expr: &seqExpr{
									pos: position{line: 459, col: 41, offset: 13895},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 459, col: 41, offset: 13895},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 459, col: 45, offset: 13899},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 459, col: 61, offset: 13915},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 459, col: 65, offset: 13919},
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
			pos:  position{line: 492, col: 1, offset: 14812},
			expr: &actionExpr{
				pos: position{line: 492, col: 17, offset: 14828},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 492, col: 17, offset: 14828},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 492, col: 17, offset: 14828},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 492, col: 19, offset: 14830},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 492, col: 25, offset: 14836},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 492, col: 34, offset: 14845},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 492, col: 39, offset: 14850},
								expr: &seqExpr{
									pos: position{line: 492, col: 40, offset: 14851},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 492, col: 40, offset: 14851},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 44, offset: 14855},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 61, offset: 14872},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 65, offset: 14876},
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
			pos:  position{line: 524, col: 1, offset: 15763},
			expr: &actionExpr{
				pos: position{line: 524, col: 12, offset: 15774},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 524, col: 12, offset: 15774},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 524, col: 12, offset: 15774},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 524, col: 14, offset: 15776},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 524, col: 20, offset: 15782},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 524, col: 30, offset: 15792},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 524, col: 35, offset: 15797},
								expr: &seqExpr{
									pos: position{line: 524, col: 36, offset: 15798},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 524, col: 36, offset: 15798},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 524, col: 40, offset: 15802},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 524, col: 52, offset: 15814},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 524, col: 56, offset: 15818},
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
			pos:  position{line: 556, col: 1, offset: 16706},
			expr: &actionExpr{
				pos: position{line: 556, col: 13, offset: 16718},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 556, col: 13, offset: 16718},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 556, col: 13, offset: 16718},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 556, col: 15, offset: 16720},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 556, col: 21, offset: 16726},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 556, col: 33, offset: 16738},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 556, col: 38, offset: 16743},
								expr: &seqExpr{
									pos: position{line: 556, col: 39, offset: 16744},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 556, col: 39, offset: 16744},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 556, col: 43, offset: 16748},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 556, col: 56, offset: 16761},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 556, col: 60, offset: 16765},
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
			pos:  position{line: 587, col: 1, offset: 17654},
			expr: &choiceExpr{
				pos: position{line: 587, col: 15, offset: 17668},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 587, col: 15, offset: 17668},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 587, col: 15, offset: 17668},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 587, col: 15, offset: 17668},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 587, col: 17, offset: 17670},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 587, col: 21, offset: 17674},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 587, col: 24, offset: 17677},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 587, col: 30, offset: 17683},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 587, col: 36, offset: 17689},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 587, col: 39, offset: 17692},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 590, col: 5, offset: 17815},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 592, col: 1, offset: 17822},
			expr: &choiceExpr{
				pos: position{line: 592, col: 12, offset: 17833},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 592, col: 12, offset: 17833},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 592, col: 30, offset: 17851},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 592, col: 49, offset: 17870},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 592, col: 64, offset: 17885},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 594, col: 1, offset: 17898},
			expr: &actionExpr{
				pos: position{line: 594, col: 19, offset: 17916},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 594, col: 21, offset: 17918},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 594, col: 21, offset: 17918},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 594, col: 28, offset: 17925},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 598, col: 1, offset: 18007},
			expr: &actionExpr{
				pos: position{line: 598, col: 20, offset: 18026},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 598, col: 22, offset: 18028},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 598, col: 22, offset: 18028},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 29, offset: 18035},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 36, offset: 18042},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 42, offset: 18048},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 48, offset: 18054},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 56, offset: 18062},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 602, col: 1, offset: 18141},
			expr: &choiceExpr{
				pos: position{line: 602, col: 16, offset: 18156},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 602, col: 16, offset: 18156},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 602, col: 18, offset: 18158},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 602, col: 18, offset: 18158},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 602, col: 24, offset: 18164},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 605, col: 3, offset: 18247},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 605, col: 5, offset: 18249},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 608, col: 3, offset: 18329},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 608, col: 3, offset: 18329},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 612, col: 1, offset: 18410},
			expr: &actionExpr{
				pos: position{line: 612, col: 15, offset: 18424},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 612, col: 17, offset: 18426},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 612, col: 17, offset: 18426},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 612, col: 23, offset: 18432},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 616, col: 1, offset: 18514},
			expr: &choiceExpr{
				pos: position{line: 616, col: 9, offset: 18522},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 616, col: 9, offset: 18522},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 16, offset: 18529},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 31, offset: 18544},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 46, offset: 18559},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 618, col: 1, offset: 18566},
			expr: &choiceExpr{
				pos: position{line: 618, col: 14, offset: 18579},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 618, col: 14, offset: 18579},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 618, col: 29, offset: 18594},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 620, col: 1, offset: 18602},
			expr: &choiceExpr{
				pos: position{line: 620, col: 14, offset: 18615},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 620, col: 14, offset: 18615},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 620, col: 29, offset: 18630},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 622, col: 1, offset: 18642},
			expr: &actionExpr{
				pos: position{line: 622, col: 16, offset: 18657},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 622, col: 16, offset: 18657},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 622, col: 16, offset: 18657},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 622, col: 20, offset: 18661},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 622, col: 22, offset: 18663},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 622, col: 28, offset: 18669},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 622, col: 33, offset: 18674},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 622, col: 35, offset: 18676},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 622, col: 40, offset: 18681},
								expr: &seqExpr{
									pos: position{line: 622, col: 41, offset: 18682},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 622, col: 41, offset: 18682},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 622, col: 45, offset: 18686},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 622, col: 47, offset: 18688},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 622, col: 52, offset: 18693},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 622, col: 56, offset: 18697},
							expr: &litMatcher{
								pos:        position{line: 622, col: 56, offset: 18697},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 622, col: 61, offset: 18702},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 622, col: 63, offset: 18704},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 638, col: 1, offset: 19149},
			expr: &actionExpr{
				pos: position{line: 638, col: 19, offset: 19167},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 638, col: 19, offset: 19167},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 638, col: 19, offset: 19167},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 638, col: 24, offset: 19172},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 638, col: 35, offset: 19183},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 638, col: 37, offset: 19185},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 638, col: 42, offset: 19190},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 651, col: 1, offset: 19462},
			expr: &actionExpr{
				pos: position{line: 651, col: 18, offset: 19479},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 651, col: 18, offset: 19479},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 651, col: 18, offset: 19479},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 651, col: 23, offset: 19484},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 34, offset: 19495},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 651, col: 36, offset: 19497},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 40, offset: 19501},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 651, col: 42, offset: 19503},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 651, col: 52, offset: 19513},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 65, offset: 19526},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 651, col: 67, offset: 19528},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 71, offset: 19532},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 651, col: 73, offset: 19534},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 651, col: 84, offset: 19545},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 651, col: 89, offset: 19550},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 651, col: 94, offset: 19555},
								expr: &seqExpr{
									pos: position{line: 651, col: 95, offset: 19556},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 651, col: 95, offset: 19556},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 99, offset: 19560},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 101, offset: 19562},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 114, offset: 19575},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 651, col: 116, offset: 19577},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 120, offset: 19581},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 122, offset: 19583},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 651, col: 130, offset: 19591},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 671, col: 1, offset: 20181},
			expr: &actionExpr{
				pos: position{line: 671, col: 17, offset: 20197},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 671, col: 17, offset: 20197},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 671, col: 17, offset: 20197},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 671, col: 22, offset: 20202},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 675, col: 1, offset: 20275},
			expr: &actionExpr{
				pos: position{line: 675, col: 16, offset: 20290},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 675, col: 16, offset: 20290},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 675, col: 16, offset: 20290},
							expr: &ruleRefExpr{
								pos:  position{line: 675, col: 17, offset: 20291},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 675, col: 27, offset: 20301},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 675, col: 27, offset: 20301},
									expr: &charClassMatcher{
										pos:        position{line: 675, col: 27, offset: 20301},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 675, col: 34, offset: 20308},
									expr: &charClassMatcher{
										pos:        position{line: 675, col: 34, offset: 20308},
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
			pos:  position{line: 679, col: 1, offset: 20383},
			expr: &actionExpr{
				pos: position{line: 679, col: 14, offset: 20396},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 679, col: 15, offset: 20397},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 679, col: 15, offset: 20397},
							expr: &charClassMatcher{
								pos:        position{line: 679, col: 15, offset: 20397},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 679, col: 22, offset: 20404},
							expr: &charClassMatcher{
								pos:        position{line: 679, col: 22, offset: 20404},
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
			pos:  position{line: 683, col: 1, offset: 20479},
			expr: &choiceExpr{
				pos: position{line: 683, col: 9, offset: 20487},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 683, col: 9, offset: 20487},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 683, col: 9, offset: 20487},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 683, col: 9, offset: 20487},
									expr: &litMatcher{
										pos:        position{line: 683, col: 9, offset: 20487},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 683, col: 14, offset: 20492},
									expr: &charClassMatcher{
										pos:        position{line: 683, col: 14, offset: 20492},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 683, col: 21, offset: 20499},
									expr: &litMatcher{
										pos:        position{line: 683, col: 22, offset: 20500},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 690, col: 3, offset: 20675},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 690, col: 3, offset: 20675},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 690, col: 3, offset: 20675},
									expr: &litMatcher{
										pos:        position{line: 690, col: 3, offset: 20675},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 690, col: 8, offset: 20680},
									expr: &charClassMatcher{
										pos:        position{line: 690, col: 8, offset: 20680},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 690, col: 15, offset: 20687},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 690, col: 19, offset: 20691},
									expr: &charClassMatcher{
										pos:        position{line: 690, col: 19, offset: 20691},
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
						pos: position{line: 697, col: 3, offset: 20880},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 697, col: 3, offset: 20880},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 701, col: 3, offset: 20965},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 701, col: 3, offset: 20965},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 704, col: 3, offset: 21051},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 705, col: 3, offset: 21058},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 706, col: 3, offset: 21074},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 706, col: 3, offset: 21074},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 706, col: 3, offset: 21074},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 706, col: 7, offset: 21078},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 706, col: 12, offset: 21083},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 706, col: 12, offset: 21083},
												expr: &ruleRefExpr{
													pos:  position{line: 706, col: 13, offset: 21084},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 706, col: 25, offset: 21096,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 706, col: 28, offset: 21099},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 708, col: 5, offset: 21191},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 708, col: 20, offset: 21206},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 708, col: 37, offset: 21223},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 710, col: 1, offset: 21240},
			expr: &actionExpr{
				pos: position{line: 710, col: 8, offset: 21247},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 710, col: 8, offset: 21247},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 714, col: 1, offset: 21310},
			expr: &actionExpr{
				pos: position{line: 714, col: 17, offset: 21326},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 714, col: 17, offset: 21326},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 714, col: 17, offset: 21326},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 714, col: 21, offset: 21330},
							expr: &seqExpr{
								pos: position{line: 714, col: 22, offset: 21331},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 714, col: 22, offset: 21331},
										expr: &ruleRefExpr{
											pos:  position{line: 714, col: 23, offset: 21332},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 714, col: 35, offset: 21344,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 714, col: 39, offset: 21348},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 722, col: 1, offset: 21531},
			expr: &actionExpr{
				pos: position{line: 722, col: 10, offset: 21540},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 722, col: 11, offset: 21541},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 726, col: 1, offset: 21596},
			expr: &seqExpr{
				pos: position{line: 726, col: 12, offset: 21607},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 726, col: 13, offset: 21608},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 726, col: 13, offset: 21608},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 21, offset: 21616},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 28, offset: 21623},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 37, offset: 21632},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 48, offset: 21643},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 57, offset: 21652},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 66, offset: 21661},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 76, offset: 21671},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 88, offset: 21683},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 726, col: 97, offset: 21692},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 726, col: 107, offset: 21702},
						expr: &oneOrMoreExpr{
							pos: position{line: 726, col: 108, offset: 21703},
							expr: &charClassMatcher{
								pos:        position{line: 726, col: 108, offset: 21703},
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
			pos:  position{line: 728, col: 1, offset: 21711},
			expr: &choiceExpr{
				pos: position{line: 728, col: 12, offset: 21722},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 728, col: 12, offset: 21722},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 728, col: 14, offset: 21724},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 728, col: 14, offset: 21724},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 24, offset: 21734},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 32, offset: 21742},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 41, offset: 21751},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 52, offset: 21762},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 61, offset: 21771},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 70, offset: 21780},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 80, offset: 21790},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 731, col: 3, offset: 21888},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 733, col: 1, offset: 21894},
			expr: &charClassMatcher{
				pos:        position{line: 733, col: 15, offset: 21908},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 735, col: 1, offset: 21924},
			expr: &choiceExpr{
				pos: position{line: 735, col: 18, offset: 21941},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 735, col: 18, offset: 21941},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 735, col: 37, offset: 21960},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 737, col: 1, offset: 21975},
			expr: &charClassMatcher{
				pos:        position{line: 737, col: 20, offset: 21994},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 739, col: 1, offset: 22007},
			expr: &charClassMatcher{
				pos:        position{line: 739, col: 16, offset: 22022},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 741, col: 1, offset: 22029},
			expr: &charClassMatcher{
				pos:        position{line: 741, col: 23, offset: 22051},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 743, col: 1, offset: 22058},
			expr: &charClassMatcher{
				pos:        position{line: 743, col: 12, offset: 22069},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 745, col: 1, offset: 22080},
			expr: &choiceExpr{
				pos: position{line: 745, col: 22, offset: 22101},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 745, col: 22, offset: 22101},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 745, col: 33, offset: 22112},
						expr: &charClassMatcher{
							pos:        position{line: 745, col: 33, offset: 22112},
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
			pos:         position{line: 747, col: 1, offset: 22124},
			expr: &choiceExpr{
				pos: position{line: 747, col: 21, offset: 22144},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 747, col: 21, offset: 22144},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 747, col: 32, offset: 22155},
						expr: &charClassMatcher{
							pos:        position{line: 747, col: 32, offset: 22155},
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
			pos:         position{line: 749, col: 1, offset: 22167},
			expr: &oneOrMoreExpr{
				pos: position{line: 749, col: 33, offset: 22199},
				expr: &charClassMatcher{
					pos:        position{line: 749, col: 33, offset: 22199},
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
			pos:         position{line: 751, col: 1, offset: 22207},
			expr: &zeroOrMoreExpr{
				pos: position{line: 751, col: 32, offset: 22238},
				expr: &charClassMatcher{
					pos:        position{line: 751, col: 32, offset: 22238},
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
			pos:         position{line: 753, col: 1, offset: 22246},
			expr: &choiceExpr{
				pos: position{line: 753, col: 15, offset: 22260},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 753, col: 15, offset: 22260},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 753, col: 26, offset: 22271},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 753, col: 26, offset: 22271},
								expr: &charClassMatcher{
									pos:        position{line: 753, col: 26, offset: 22271},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 753, col: 35, offset: 22280},
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
			pos:  position{line: 755, col: 1, offset: 22286},
			expr: &oneOrMoreExpr{
				pos: position{line: 755, col: 12, offset: 22297},
				expr: &ruleRefExpr{
					pos:  position{line: 755, col: 13, offset: 22298},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 757, col: 1, offset: 22309},
			expr: &choiceExpr{
				pos: position{line: 757, col: 11, offset: 22319},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 757, col: 11, offset: 22319},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 11, offset: 22319},
								expr: &charClassMatcher{
									pos:        position{line: 757, col: 11, offset: 22319},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 22, offset: 22330},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 27, offset: 22335},
								expr: &seqExpr{
									pos: position{line: 757, col: 28, offset: 22336},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 757, col: 28, offset: 22336},
											expr: &charClassMatcher{
												pos:        position{line: 757, col: 29, offset: 22337},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 757, col: 34, offset: 22342,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 757, col: 38, offset: 22346},
								expr: &litMatcher{
									pos:        position{line: 757, col: 39, offset: 22347},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 757, col: 46, offset: 22354},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 46, offset: 22354},
								expr: &charClassMatcher{
									pos:        position{line: 757, col: 46, offset: 22354},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 57, offset: 22365},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 62, offset: 22370},
								expr: &seqExpr{
									pos: position{line: 757, col: 63, offset: 22371},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 757, col: 63, offset: 22371},
											expr: &litMatcher{
												pos:        position{line: 757, col: 64, offset: 22372},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 757, col: 69, offset: 22377,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 73, offset: 22381},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 78, offset: 22386},
								expr: &charClassMatcher{
									pos:        position{line: 757, col: 78, offset: 22386},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 84, offset: 22392},
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
			pos:  position{line: 759, col: 1, offset: 22398},
			expr: &notExpr{
				pos: position{line: 759, col: 7, offset: 22404},
				expr: &anyMatcher{
					line: 759, col: 8, offset: 22405,
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

func (c *current) onTypeDefn19(name, params, first, rest interface{}) (interface{}, error) {
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

func (p *parser) callonTypeDefn19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn19(stack["name"], stack["params"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn48(name, first, rest interface{}) (interface{}, error) {
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

func (p *parser) callonTypeDefn48() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn48(stack["name"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn74(name, params, rest interface{}) (interface{}, error) {
	// variant type
	constructors := []VariantConstructor{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		for _, v := range vals {
			constructors = append(constructors, v.(VariantConstructor))
		}
	}

	return VariantType{Name: name.(Identifier).StringValue, Params: params.(Container).Subvalues, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn74() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn74(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onTypeDefn89(name, rest interface{}) (interface{}, error) {
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

func (p *parser) callonTypeDefn89() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn89(stack["name"], stack["rest"])
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

func (c *current) onVariantConstructor27(name, rest interface{}) (interface{}, error) {
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

func (p *parser) callonVariantConstructor27() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor27(stack["name"], stack["rest"])
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

func (c *current) onTypeAnnotation3(args, ret interface{}) (interface{}, error) {
	// TODO: return correct func type annotation
	vals := args.(Container)
	vals.Subvalues = append(vals.Subvalues, ret.(Ast))
	vals.Type = "FuncAnnotation"

	return vals, nil
}

func (p *parser) callonTypeAnnotation3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeAnnotation3(stack["args"], stack["ret"])
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
