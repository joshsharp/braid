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
								&litMatcher{
									pos:        position{line: 187, col: 22, offset: 5868},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 26, offset: 5872},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 28, offset: 5874},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 33, offset: 5879},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 44, offset: 5890},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 46, offset: 5892},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 50, offset: 5896},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 52, offset: 5898},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 58, offset: 5904},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 74, offset: 5920},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 76, offset: 5922},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 187, col: 81, offset: 5927},
										expr: &seqExpr{
											pos: position{line: 187, col: 82, offset: 5928},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 187, col: 82, offset: 5928},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 187, col: 86, offset: 5932},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 187, col: 88, offset: 5934},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 187, col: 104, offset: 5950},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 187, col: 108, offset: 5954},
									expr: &litMatcher{
										pos:        position{line: 187, col: 108, offset: 5954},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 113, offset: 5959},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 115, offset: 5961},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 119, offset: 5965},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 206, col: 1, offset: 6569},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 206, col: 1, offset: 6569},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 206, col: 1, offset: 6569},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 5, offset: 6573},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 206, col: 7, offset: 6575},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 206, col: 12, offset: 6580},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 206, col: 23, offset: 6591},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 206, col: 28, offset: 6596},
										expr: &seqExpr{
											pos: position{line: 206, col: 29, offset: 6597},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 206, col: 29, offset: 6597},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 206, col: 32, offset: 6600},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 49, offset: 6617},
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
			pos:  position{line: 223, col: 1, offset: 7054},
			expr: &choiceExpr{
				pos: position{line: 223, col: 14, offset: 7067},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 223, col: 14, offset: 7067},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 223, col: 14, offset: 7067},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 223, col: 14, offset: 7067},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 223, col: 16, offset: 7069},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 22, offset: 7075},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 223, col: 26, offset: 7079},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 223, col: 28, offset: 7081},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 39, offset: 7092},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 223, col: 42, offset: 7095},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 46, offset: 7099},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 223, col: 49, offset: 7102},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 223, col: 54, offset: 7107},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 223, col: 59, offset: 7112},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 229, col: 1, offset: 7231},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 229, col: 1, offset: 7231},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 229, col: 1, offset: 7231},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 229, col: 3, offset: 7233},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 9, offset: 7239},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 229, col: 13, offset: 7243},
									expr: &ruleRefExpr{
										pos:  position{line: 229, col: 14, offset: 7244},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 233, col: 1, offset: 7342},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 233, col: 1, offset: 7342},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 233, col: 1, offset: 7342},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 233, col: 3, offset: 7344},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 233, col: 9, offset: 7350},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 233, col: 13, offset: 7354},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 233, col: 15, offset: 7356},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 233, col: 26, offset: 7367},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 233, col: 29, offset: 7370},
									expr: &litMatcher{
										pos:        position{line: 233, col: 30, offset: 7371},
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
			pos:  position{line: 237, col: 1, offset: 7465},
			expr: &actionExpr{
				pos: position{line: 237, col: 12, offset: 7476},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 237, col: 12, offset: 7476},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 237, col: 12, offset: 7476},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 237, col: 14, offset: 7478},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 20, offset: 7484},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 24, offset: 7488},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 26, offset: 7490},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 39, offset: 7503},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 237, col: 42, offset: 7506},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 46, offset: 7510},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 49, offset: 7513},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 237, col: 53, offset: 7517},
								expr: &seqExpr{
									pos: position{line: 237, col: 54, offset: 7518},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 237, col: 54, offset: 7518},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 63, offset: 7527},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 67, offset: 7531},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 237, col: 71, offset: 7535},
								expr: &seqExpr{
									pos: position{line: 237, col: 72, offset: 7536},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 237, col: 72, offset: 7536},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 237, col: 74, offset: 7538},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 79, offset: 7543},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 81, offset: 7545},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 96, offset: 7560},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 237, col: 100, offset: 7564},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 104, offset: 7568},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 107, offset: 7571},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 237, col: 118, offset: 7582},
								expr: &ruleRefExpr{
									pos:  position{line: 237, col: 119, offset: 7583},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 131, offset: 7595},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 237, col: 133, offset: 7597},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 137, offset: 7601},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 263, col: 1, offset: 8196},
			expr: &actionExpr{
				pos: position{line: 263, col: 8, offset: 8203},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 263, col: 8, offset: 8203},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 263, col: 12, offset: 8207},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 263, col: 12, offset: 8207},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 263, col: 21, offset: 8216},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 269, col: 1, offset: 8333},
			expr: &choiceExpr{
				pos: position{line: 269, col: 10, offset: 8342},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 269, col: 10, offset: 8342},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 269, col: 10, offset: 8342},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 269, col: 10, offset: 8342},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 269, col: 12, offset: 8344},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 17, offset: 8349},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 269, col: 21, offset: 8353},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 269, col: 26, offset: 8358},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 36, offset: 8368},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 269, col: 39, offset: 8371},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 43, offset: 8375},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 269, col: 45, offset: 8377},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 269, col: 51, offset: 8383},
										expr: &ruleRefExpr{
											pos:  position{line: 269, col: 52, offset: 8384},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 64, offset: 8396},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 269, col: 67, offset: 8399},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 71, offset: 8403},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 269, col: 74, offset: 8406},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 81, offset: 8413},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 269, col: 84, offset: 8416},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 88, offset: 8420},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 269, col: 90, offset: 8422},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 269, col: 96, offset: 8428},
										expr: &ruleRefExpr{
											pos:  position{line: 269, col: 97, offset: 8429},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 269, col: 109, offset: 8441},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 269, col: 112, offset: 8444},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 288, col: 1, offset: 8947},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 288, col: 1, offset: 8947},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 288, col: 1, offset: 8947},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 288, col: 3, offset: 8949},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 8, offset: 8954},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 288, col: 12, offset: 8958},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 288, col: 17, offset: 8963},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 27, offset: 8973},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 288, col: 30, offset: 8976},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 34, offset: 8980},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 288, col: 36, offset: 8982},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 288, col: 42, offset: 8988},
										expr: &ruleRefExpr{
											pos:  position{line: 288, col: 43, offset: 8989},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 55, offset: 9001},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 288, col: 57, offset: 9003},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 288, col: 61, offset: 9007},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 288, col: 64, offset: 9010},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 288, col: 71, offset: 9017},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 288, col: 79, offset: 9025},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 300, col: 1, offset: 9355},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 300, col: 1, offset: 9355},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 300, col: 1, offset: 9355},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 300, col: 3, offset: 9357},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 8, offset: 9362},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 300, col: 12, offset: 9366},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 17, offset: 9371},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 27, offset: 9381},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 300, col: 30, offset: 9384},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 34, offset: 9388},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 300, col: 36, offset: 9390},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 300, col: 42, offset: 9396},
										expr: &ruleRefExpr{
											pos:  position{line: 300, col: 43, offset: 9397},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 55, offset: 9409},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 300, col: 58, offset: 9412},
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
			pos:  position{line: 312, col: 1, offset: 9710},
			expr: &choiceExpr{
				pos: position{line: 312, col: 8, offset: 9717},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 312, col: 8, offset: 9717},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 312, col: 8, offset: 9717},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 312, col: 8, offset: 9717},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 312, col: 10, offset: 9719},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 17, offset: 9726},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 312, col: 28, offset: 9737},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 312, col: 32, offset: 9741},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 35, offset: 9744},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 312, col: 48, offset: 9757},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 53, offset: 9762},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 326, col: 1, offset: 10086},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 326, col: 1, offset: 10086},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 326, col: 1, offset: 10086},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 326, col: 3, offset: 10088},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 6, offset: 10091},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 326, col: 19, offset: 10104},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 24, offset: 10109},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 340, col: 1, offset: 10426},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 340, col: 1, offset: 10426},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 340, col: 1, offset: 10426},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 340, col: 3, offset: 10428},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 6, offset: 10431},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 340, col: 19, offset: 10444},
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
			pos:  position{line: 347, col: 1, offset: 10615},
			expr: &actionExpr{
				pos: position{line: 347, col: 16, offset: 10630},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 347, col: 16, offset: 10630},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 347, col: 16, offset: 10630},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 347, col: 23, offset: 10637},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 347, col: 36, offset: 10650},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 347, col: 41, offset: 10655},
								expr: &seqExpr{
									pos: position{line: 347, col: 42, offset: 10656},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 347, col: 42, offset: 10656},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 347, col: 46, offset: 10660},
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
			pos:  position{line: 364, col: 1, offset: 11097},
			expr: &actionExpr{
				pos: position{line: 364, col: 12, offset: 11108},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 364, col: 12, offset: 11108},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 364, col: 12, offset: 11108},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 16, offset: 11112},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 364, col: 18, offset: 11114},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 27, offset: 11123},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 35, offset: 11131},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 364, col: 37, offset: 11133},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 364, col: 42, offset: 11138},
								expr: &seqExpr{
									pos: position{line: 364, col: 43, offset: 11139},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 364, col: 43, offset: 11139},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 47, offset: 11143},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 49, offset: 11145},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 59, offset: 11155},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 364, col: 61, offset: 11157},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 382, col: 1, offset: 11579},
			expr: &actionExpr{
				pos: position{line: 382, col: 11, offset: 11589},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 382, col: 11, offset: 11589},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 382, col: 11, offset: 11589},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 382, col: 16, offset: 11594},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 382, col: 27, offset: 11605},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 382, col: 29, offset: 11607},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 382, col: 34, offset: 11612},
								expr: &seqExpr{
									pos: position{line: 382, col: 35, offset: 11613},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 382, col: 35, offset: 11613},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 382, col: 39, offset: 11617},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 382, col: 41, offset: 11619},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 382, col: 59, offset: 11637},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 403, col: 1, offset: 12176},
			expr: &choiceExpr{
				pos: position{line: 403, col: 18, offset: 12193},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 403, col: 18, offset: 12193},
						name: "AnyType",
					},
					&actionExpr{
						pos: position{line: 404, col: 1, offset: 12204},
						run: (*parser).callonTypeAnnotation3,
						expr: &seqExpr{
							pos: position{line: 404, col: 1, offset: 12204},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 404, col: 1, offset: 12204},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 404, col: 8, offset: 12211},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 404, col: 11, offset: 12214},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 404, col: 16, offset: 12219},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 404, col: 25, offset: 12228},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 404, col: 27, offset: 12230},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 404, col: 32, offset: 12235},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 404, col: 34, offset: 12237},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 404, col: 38, offset: 12241},
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
			pos:  position{line: 413, col: 1, offset: 12457},
			expr: &choiceExpr{
				pos: position{line: 413, col: 11, offset: 12467},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 413, col: 11, offset: 12467},
						name: "ModuleName",
					},
					&ruleRefExpr{
						pos:  position{line: 413, col: 24, offset: 12480},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 413, col: 35, offset: 12491},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 415, col: 1, offset: 12506},
			expr: &choiceExpr{
				pos: position{line: 415, col: 13, offset: 12518},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 415, col: 13, offset: 12518},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 415, col: 13, offset: 12518},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 415, col: 13, offset: 12518},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 415, col: 17, offset: 12522},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 415, col: 19, offset: 12524},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 415, col: 28, offset: 12533},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 415, col: 40, offset: 12545},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 415, col: 42, offset: 12547},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 415, col: 47, offset: 12552},
										expr: &seqExpr{
											pos: position{line: 415, col: 48, offset: 12553},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 415, col: 48, offset: 12553},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 415, col: 52, offset: 12557},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 415, col: 54, offset: 12559},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 415, col: 68, offset: 12573},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 415, col: 70, offset: 12575},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 432, col: 1, offset: 12997},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 432, col: 1, offset: 12997},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 432, col: 1, offset: 12997},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 432, col: 5, offset: 13001},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 432, col: 7, offset: 13003},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 432, col: 16, offset: 13012},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 432, col: 21, offset: 13017},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 432, col: 23, offset: 13019},
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
			pos:  position{line: 437, col: 1, offset: 13124},
			expr: &actionExpr{
				pos: position{line: 437, col: 16, offset: 13139},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 437, col: 16, offset: 13139},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 437, col: 16, offset: 13139},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 437, col: 18, offset: 13141},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 437, col: 21, offset: 13144},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 437, col: 27, offset: 13150},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 437, col: 32, offset: 13155},
								expr: &seqExpr{
									pos: position{line: 437, col: 33, offset: 13156},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 437, col: 33, offset: 13156},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 437, col: 37, offset: 13160},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 437, col: 46, offset: 13169},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 437, col: 50, offset: 13173},
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
			pos:  position{line: 457, col: 1, offset: 13779},
			expr: &choiceExpr{
				pos: position{line: 457, col: 9, offset: 13787},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 457, col: 9, offset: 13787},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 21, offset: 13799},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 37, offset: 13815},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 48, offset: 13826},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 60, offset: 13838},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 459, col: 1, offset: 13851},
			expr: &actionExpr{
				pos: position{line: 459, col: 13, offset: 13863},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 459, col: 13, offset: 13863},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 459, col: 13, offset: 13863},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 459, col: 15, offset: 13865},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 21, offset: 13871},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 459, col: 35, offset: 13885},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 459, col: 40, offset: 13890},
								expr: &seqExpr{
									pos: position{line: 459, col: 41, offset: 13891},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 459, col: 41, offset: 13891},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 459, col: 45, offset: 13895},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 459, col: 61, offset: 13911},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 459, col: 65, offset: 13915},
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
			pos:  position{line: 492, col: 1, offset: 14808},
			expr: &actionExpr{
				pos: position{line: 492, col: 17, offset: 14824},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 492, col: 17, offset: 14824},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 492, col: 17, offset: 14824},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 492, col: 19, offset: 14826},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 492, col: 25, offset: 14832},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 492, col: 34, offset: 14841},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 492, col: 39, offset: 14846},
								expr: &seqExpr{
									pos: position{line: 492, col: 40, offset: 14847},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 492, col: 40, offset: 14847},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 44, offset: 14851},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 61, offset: 14868},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 492, col: 65, offset: 14872},
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
			pos:  position{line: 524, col: 1, offset: 15759},
			expr: &actionExpr{
				pos: position{line: 524, col: 12, offset: 15770},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 524, col: 12, offset: 15770},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 524, col: 12, offset: 15770},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 524, col: 14, offset: 15772},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 524, col: 20, offset: 15778},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 524, col: 30, offset: 15788},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 524, col: 35, offset: 15793},
								expr: &seqExpr{
									pos: position{line: 524, col: 36, offset: 15794},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 524, col: 36, offset: 15794},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 524, col: 40, offset: 15798},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 524, col: 52, offset: 15810},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 524, col: 56, offset: 15814},
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
			pos:  position{line: 556, col: 1, offset: 16702},
			expr: &actionExpr{
				pos: position{line: 556, col: 13, offset: 16714},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 556, col: 13, offset: 16714},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 556, col: 13, offset: 16714},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 556, col: 15, offset: 16716},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 556, col: 21, offset: 16722},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 556, col: 33, offset: 16734},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 556, col: 38, offset: 16739},
								expr: &seqExpr{
									pos: position{line: 556, col: 39, offset: 16740},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 556, col: 39, offset: 16740},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 556, col: 43, offset: 16744},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 556, col: 56, offset: 16757},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 556, col: 60, offset: 16761},
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
			pos:  position{line: 587, col: 1, offset: 17650},
			expr: &choiceExpr{
				pos: position{line: 587, col: 15, offset: 17664},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 587, col: 15, offset: 17664},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 587, col: 15, offset: 17664},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 587, col: 15, offset: 17664},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 587, col: 17, offset: 17666},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 587, col: 21, offset: 17670},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 587, col: 24, offset: 17673},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 587, col: 30, offset: 17679},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 587, col: 36, offset: 17685},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 587, col: 39, offset: 17688},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 590, col: 5, offset: 17811},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 592, col: 1, offset: 17818},
			expr: &choiceExpr{
				pos: position{line: 592, col: 12, offset: 17829},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 592, col: 12, offset: 17829},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 592, col: 30, offset: 17847},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 592, col: 49, offset: 17866},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 592, col: 64, offset: 17881},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 594, col: 1, offset: 17894},
			expr: &actionExpr{
				pos: position{line: 594, col: 19, offset: 17912},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 594, col: 21, offset: 17914},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 594, col: 21, offset: 17914},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 594, col: 28, offset: 17921},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 598, col: 1, offset: 18003},
			expr: &actionExpr{
				pos: position{line: 598, col: 20, offset: 18022},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 598, col: 22, offset: 18024},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 598, col: 22, offset: 18024},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 29, offset: 18031},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 36, offset: 18038},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 42, offset: 18044},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 48, offset: 18050},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 598, col: 56, offset: 18058},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 602, col: 1, offset: 18137},
			expr: &choiceExpr{
				pos: position{line: 602, col: 16, offset: 18152},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 602, col: 16, offset: 18152},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 602, col: 18, offset: 18154},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 602, col: 18, offset: 18154},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 602, col: 24, offset: 18160},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 605, col: 3, offset: 18243},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 605, col: 5, offset: 18245},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 608, col: 3, offset: 18325},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 608, col: 3, offset: 18325},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 612, col: 1, offset: 18406},
			expr: &actionExpr{
				pos: position{line: 612, col: 15, offset: 18420},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 612, col: 17, offset: 18422},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 612, col: 17, offset: 18422},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 612, col: 23, offset: 18428},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 616, col: 1, offset: 18510},
			expr: &choiceExpr{
				pos: position{line: 616, col: 9, offset: 18518},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 616, col: 9, offset: 18518},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 16, offset: 18525},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 31, offset: 18540},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 46, offset: 18555},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 618, col: 1, offset: 18562},
			expr: &choiceExpr{
				pos: position{line: 618, col: 14, offset: 18575},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 618, col: 14, offset: 18575},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 618, col: 29, offset: 18590},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 620, col: 1, offset: 18598},
			expr: &choiceExpr{
				pos: position{line: 620, col: 14, offset: 18611},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 620, col: 14, offset: 18611},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 620, col: 29, offset: 18626},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 622, col: 1, offset: 18638},
			expr: &actionExpr{
				pos: position{line: 622, col: 16, offset: 18653},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 622, col: 16, offset: 18653},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 622, col: 16, offset: 18653},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 622, col: 20, offset: 18657},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 622, col: 22, offset: 18659},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 622, col: 28, offset: 18665},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 622, col: 33, offset: 18670},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 622, col: 35, offset: 18672},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 622, col: 40, offset: 18677},
								expr: &seqExpr{
									pos: position{line: 622, col: 41, offset: 18678},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 622, col: 41, offset: 18678},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 622, col: 45, offset: 18682},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 622, col: 47, offset: 18684},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 622, col: 52, offset: 18689},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 622, col: 56, offset: 18693},
							expr: &litMatcher{
								pos:        position{line: 622, col: 56, offset: 18693},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 622, col: 61, offset: 18698},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 622, col: 63, offset: 18700},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 638, col: 1, offset: 19145},
			expr: &actionExpr{
				pos: position{line: 638, col: 19, offset: 19163},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 638, col: 19, offset: 19163},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 638, col: 19, offset: 19163},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 638, col: 24, offset: 19168},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 638, col: 35, offset: 19179},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 638, col: 37, offset: 19181},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 638, col: 42, offset: 19186},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 651, col: 1, offset: 19456},
			expr: &actionExpr{
				pos: position{line: 651, col: 18, offset: 19473},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 651, col: 18, offset: 19473},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 651, col: 18, offset: 19473},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 651, col: 23, offset: 19478},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 34, offset: 19489},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 651, col: 36, offset: 19491},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 40, offset: 19495},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 651, col: 42, offset: 19497},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 651, col: 52, offset: 19507},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 65, offset: 19520},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 651, col: 67, offset: 19522},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 651, col: 71, offset: 19526},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 651, col: 73, offset: 19528},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 651, col: 84, offset: 19539},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 651, col: 89, offset: 19544},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 651, col: 94, offset: 19549},
								expr: &seqExpr{
									pos: position{line: 651, col: 95, offset: 19550},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 651, col: 95, offset: 19550},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 99, offset: 19554},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 101, offset: 19556},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 114, offset: 19569},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 651, col: 116, offset: 19571},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 120, offset: 19575},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 651, col: 122, offset: 19577},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 651, col: 130, offset: 19585},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 671, col: 1, offset: 20175},
			expr: &actionExpr{
				pos: position{line: 671, col: 17, offset: 20191},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 671, col: 17, offset: 20191},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 671, col: 17, offset: 20191},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 671, col: 22, offset: 20196},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 675, col: 1, offset: 20269},
			expr: &actionExpr{
				pos: position{line: 675, col: 16, offset: 20284},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 675, col: 16, offset: 20284},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 675, col: 16, offset: 20284},
							expr: &ruleRefExpr{
								pos:  position{line: 675, col: 17, offset: 20285},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 675, col: 27, offset: 20295},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 675, col: 27, offset: 20295},
									expr: &charClassMatcher{
										pos:        position{line: 675, col: 27, offset: 20295},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 675, col: 34, offset: 20302},
									expr: &charClassMatcher{
										pos:        position{line: 675, col: 34, offset: 20302},
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
			pos:  position{line: 679, col: 1, offset: 20377},
			expr: &actionExpr{
				pos: position{line: 679, col: 14, offset: 20390},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 679, col: 15, offset: 20391},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 679, col: 15, offset: 20391},
							expr: &charClassMatcher{
								pos:        position{line: 679, col: 15, offset: 20391},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 679, col: 22, offset: 20398},
							expr: &charClassMatcher{
								pos:        position{line: 679, col: 22, offset: 20398},
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
			pos:  position{line: 683, col: 1, offset: 20473},
			expr: &choiceExpr{
				pos: position{line: 683, col: 9, offset: 20481},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 683, col: 9, offset: 20481},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 683, col: 9, offset: 20481},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 683, col: 9, offset: 20481},
									expr: &litMatcher{
										pos:        position{line: 683, col: 9, offset: 20481},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 683, col: 14, offset: 20486},
									expr: &charClassMatcher{
										pos:        position{line: 683, col: 14, offset: 20486},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 683, col: 21, offset: 20493},
									expr: &litMatcher{
										pos:        position{line: 683, col: 22, offset: 20494},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 690, col: 3, offset: 20669},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 690, col: 3, offset: 20669},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 690, col: 3, offset: 20669},
									expr: &litMatcher{
										pos:        position{line: 690, col: 3, offset: 20669},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 690, col: 8, offset: 20674},
									expr: &charClassMatcher{
										pos:        position{line: 690, col: 8, offset: 20674},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 690, col: 15, offset: 20681},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 690, col: 19, offset: 20685},
									expr: &charClassMatcher{
										pos:        position{line: 690, col: 19, offset: 20685},
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
						pos: position{line: 697, col: 3, offset: 20874},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 697, col: 3, offset: 20874},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 701, col: 3, offset: 20959},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 701, col: 3, offset: 20959},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 704, col: 3, offset: 21045},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 705, col: 3, offset: 21052},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 706, col: 3, offset: 21068},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 706, col: 3, offset: 21068},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 706, col: 3, offset: 21068},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 706, col: 7, offset: 21072},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 706, col: 12, offset: 21077},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 706, col: 12, offset: 21077},
												expr: &ruleRefExpr{
													pos:  position{line: 706, col: 13, offset: 21078},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 706, col: 25, offset: 21090,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 706, col: 28, offset: 21093},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 708, col: 5, offset: 21185},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 708, col: 20, offset: 21200},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 708, col: 37, offset: 21217},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 710, col: 1, offset: 21234},
			expr: &actionExpr{
				pos: position{line: 710, col: 8, offset: 21241},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 710, col: 8, offset: 21241},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 714, col: 1, offset: 21304},
			expr: &actionExpr{
				pos: position{line: 714, col: 17, offset: 21320},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 714, col: 17, offset: 21320},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 714, col: 17, offset: 21320},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 714, col: 21, offset: 21324},
							expr: &seqExpr{
								pos: position{line: 714, col: 22, offset: 21325},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 714, col: 22, offset: 21325},
										expr: &ruleRefExpr{
											pos:  position{line: 714, col: 23, offset: 21326},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 714, col: 35, offset: 21338,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 714, col: 39, offset: 21342},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 722, col: 1, offset: 21525},
			expr: &actionExpr{
				pos: position{line: 722, col: 10, offset: 21534},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 722, col: 11, offset: 21535},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 726, col: 1, offset: 21590},
			expr: &seqExpr{
				pos: position{line: 726, col: 12, offset: 21601},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 726, col: 13, offset: 21602},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 726, col: 13, offset: 21602},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 21, offset: 21610},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 28, offset: 21617},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 37, offset: 21626},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 48, offset: 21637},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 57, offset: 21646},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 66, offset: 21655},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 76, offset: 21665},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 726, col: 88, offset: 21677},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 726, col: 97, offset: 21686},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 726, col: 107, offset: 21696},
						expr: &oneOrMoreExpr{
							pos: position{line: 726, col: 108, offset: 21697},
							expr: &charClassMatcher{
								pos:        position{line: 726, col: 108, offset: 21697},
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
			pos:  position{line: 728, col: 1, offset: 21705},
			expr: &choiceExpr{
				pos: position{line: 728, col: 12, offset: 21716},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 728, col: 12, offset: 21716},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 728, col: 14, offset: 21718},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 728, col: 14, offset: 21718},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 24, offset: 21728},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 32, offset: 21736},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 41, offset: 21745},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 52, offset: 21756},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 61, offset: 21765},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 70, offset: 21774},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 728, col: 80, offset: 21784},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 731, col: 3, offset: 21882},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 733, col: 1, offset: 21888},
			expr: &charClassMatcher{
				pos:        position{line: 733, col: 15, offset: 21902},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 735, col: 1, offset: 21918},
			expr: &choiceExpr{
				pos: position{line: 735, col: 18, offset: 21935},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 735, col: 18, offset: 21935},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 735, col: 37, offset: 21954},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 737, col: 1, offset: 21969},
			expr: &charClassMatcher{
				pos:        position{line: 737, col: 20, offset: 21988},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 739, col: 1, offset: 22001},
			expr: &charClassMatcher{
				pos:        position{line: 739, col: 16, offset: 22016},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 741, col: 1, offset: 22023},
			expr: &charClassMatcher{
				pos:        position{line: 741, col: 23, offset: 22045},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 743, col: 1, offset: 22052},
			expr: &charClassMatcher{
				pos:        position{line: 743, col: 12, offset: 22063},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 745, col: 1, offset: 22074},
			expr: &choiceExpr{
				pos: position{line: 745, col: 22, offset: 22095},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 745, col: 22, offset: 22095},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 745, col: 33, offset: 22106},
						expr: &charClassMatcher{
							pos:        position{line: 745, col: 33, offset: 22106},
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
			pos:         position{line: 747, col: 1, offset: 22118},
			expr: &choiceExpr{
				pos: position{line: 747, col: 21, offset: 22138},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 747, col: 21, offset: 22138},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 747, col: 32, offset: 22149},
						expr: &charClassMatcher{
							pos:        position{line: 747, col: 32, offset: 22149},
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
			pos:         position{line: 749, col: 1, offset: 22161},
			expr: &oneOrMoreExpr{
				pos: position{line: 749, col: 33, offset: 22193},
				expr: &charClassMatcher{
					pos:        position{line: 749, col: 33, offset: 22193},
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
			pos:         position{line: 751, col: 1, offset: 22201},
			expr: &zeroOrMoreExpr{
				pos: position{line: 751, col: 32, offset: 22232},
				expr: &charClassMatcher{
					pos:        position{line: 751, col: 32, offset: 22232},
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
			pos:         position{line: 753, col: 1, offset: 22240},
			expr: &choiceExpr{
				pos: position{line: 753, col: 15, offset: 22254},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 753, col: 15, offset: 22254},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 753, col: 26, offset: 22265},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 753, col: 26, offset: 22265},
								expr: &charClassMatcher{
									pos:        position{line: 753, col: 26, offset: 22265},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 753, col: 35, offset: 22274},
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
			pos:  position{line: 755, col: 1, offset: 22280},
			expr: &oneOrMoreExpr{
				pos: position{line: 755, col: 12, offset: 22291},
				expr: &ruleRefExpr{
					pos:  position{line: 755, col: 13, offset: 22292},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 757, col: 1, offset: 22303},
			expr: &choiceExpr{
				pos: position{line: 757, col: 11, offset: 22313},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 757, col: 11, offset: 22313},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 11, offset: 22313},
								expr: &charClassMatcher{
									pos:        position{line: 757, col: 11, offset: 22313},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 22, offset: 22324},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 27, offset: 22329},
								expr: &seqExpr{
									pos: position{line: 757, col: 28, offset: 22330},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 757, col: 28, offset: 22330},
											expr: &charClassMatcher{
												pos:        position{line: 757, col: 29, offset: 22331},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 757, col: 34, offset: 22336,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 757, col: 38, offset: 22340},
								expr: &litMatcher{
									pos:        position{line: 757, col: 39, offset: 22341},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 757, col: 46, offset: 22348},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 46, offset: 22348},
								expr: &charClassMatcher{
									pos:        position{line: 757, col: 46, offset: 22348},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 57, offset: 22359},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 62, offset: 22364},
								expr: &seqExpr{
									pos: position{line: 757, col: 63, offset: 22365},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 757, col: 63, offset: 22365},
											expr: &litMatcher{
												pos:        position{line: 757, col: 64, offset: 22366},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 757, col: 69, offset: 22371,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 73, offset: 22375},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 757, col: 78, offset: 22380},
								expr: &charClassMatcher{
									pos:        position{line: 757, col: 78, offset: 22380},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 757, col: 84, offset: 22386},
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
			pos:  position{line: 759, col: 1, offset: 22392},
			expr: &notExpr{
				pos: position{line: 759, col: 7, offset: 22398},
				expr: &anyMatcher{
					line: 759, col: 8, offset: 22399,
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

	return VariantInstance{Name: name.(BasicAst).StringValue, Arguments: arguments}, nil
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
