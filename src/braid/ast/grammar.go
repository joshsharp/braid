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
			pos:  position{line: 82, col: 1, offset: 2505},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2516},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 82, col: 12, offset: 2516},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 82, col: 12, offset: 2516},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 82, col: 12, offset: 2516},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 14, offset: 2518},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 21, offset: 2525},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 24, offset: 2528},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 29, offset: 2533},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 82, col: 40, offset: 2544},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 82, col: 47, offset: 2551},
										expr: &seqExpr{
											pos: position{line: 82, col: 48, offset: 2552},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 48, offset: 2552},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 51, offset: 2555},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 67, offset: 2571},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 69, offset: 2573},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 82, col: 73, offset: 2577},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 82, col: 79, offset: 2583},
										expr: &seqExpr{
											pos: position{line: 82, col: 80, offset: 2584},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 80, offset: 2584},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 83, offset: 2587},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 100, offset: 2604},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 101, col: 1, offset: 3100},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 101, col: 1, offset: 3100},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 101, col: 1, offset: 3100},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 3, offset: 3102},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 10, offset: 3109},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 13, offset: 3112},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 18, offset: 3117},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 101, col: 29, offset: 3128},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 36, offset: 3135},
										expr: &seqExpr{
											pos: position{line: 101, col: 37, offset: 3136},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 101, col: 37, offset: 3136},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 40, offset: 3139},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 56, offset: 3155},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 58, offset: 3157},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 62, offset: 3161},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 5, offset: 3167},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 9, offset: 3171},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 102, col: 11, offset: 3173},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 102, col: 17, offset: 3179},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 33, offset: 3195},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 102, col: 35, offset: 3197},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 102, col: 40, offset: 3202},
										expr: &seqExpr{
											pos: position{line: 102, col: 41, offset: 3203},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 102, col: 41, offset: 3203},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 45, offset: 3207},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 47, offset: 3209},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 63, offset: 3225},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 102, col: 67, offset: 3229},
									expr: &litMatcher{
										pos:        position{line: 102, col: 67, offset: 3229},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 72, offset: 3234},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 74, offset: 3236},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 78, offset: 3240},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 120, col: 1, offset: 3725},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 120, col: 1, offset: 3725},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 120, col: 1, offset: 3725},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 3, offset: 3727},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 10, offset: 3734},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 13, offset: 3737},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 120, col: 18, offset: 3742},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 120, col: 29, offset: 3753},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 120, col: 36, offset: 3760},
										expr: &seqExpr{
											pos: position{line: 120, col: 37, offset: 3761},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 120, col: 37, offset: 3761},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 40, offset: 3764},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 56, offset: 3780},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 58, offset: 3782},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 62, offset: 3786},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 64, offset: 3788},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 120, col: 69, offset: 3793},
										expr: &ruleRefExpr{
											pos:  position{line: 120, col: 70, offset: 3794},
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
			pos:  position{line: 135, col: 1, offset: 4201},
			expr: &actionExpr{
				pos: position{line: 135, col: 19, offset: 4219},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 135, col: 19, offset: 4219},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 135, col: 19, offset: 4219},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 24, offset: 4224},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 135, col: 37, offset: 4237},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 135, col: 39, offset: 4239},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 135, col: 43, offset: 4243},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 135, col: 45, offset: 4245},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 48, offset: 4248},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 139, col: 1, offset: 4349},
			expr: &choiceExpr{
				pos: position{line: 139, col: 22, offset: 4370},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 139, col: 22, offset: 4370},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 139, col: 22, offset: 4370},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 139, col: 22, offset: 4370},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 26, offset: 4374},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 28, offset: 4376},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 139, col: 33, offset: 4381},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 44, offset: 4392},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 139, col: 46, offset: 4394},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 50, offset: 4398},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 52, offset: 4400},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 139, col: 58, offset: 4406},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 74, offset: 4422},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 76, offset: 4424},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 139, col: 81, offset: 4429},
										expr: &seqExpr{
											pos: position{line: 139, col: 82, offset: 4430},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 139, col: 82, offset: 4430},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 86, offset: 4434},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 88, offset: 4436},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 104, offset: 4452},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 139, col: 108, offset: 4456},
									expr: &litMatcher{
										pos:        position{line: 139, col: 108, offset: 4456},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 113, offset: 4461},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 139, col: 115, offset: 4463},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 119, offset: 4467},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 158, col: 1, offset: 5072},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 158, col: 1, offset: 5072},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 158, col: 1, offset: 5072},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 158, col: 5, offset: 5076},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 158, col: 7, offset: 5078},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 158, col: 12, offset: 5083},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 158, col: 23, offset: 5094},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 158, col: 28, offset: 5099},
										expr: &seqExpr{
											pos: position{line: 158, col: 29, offset: 5100},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 158, col: 29, offset: 5100},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 158, col: 32, offset: 5103},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 158, col: 49, offset: 5120},
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
			pos:  position{line: 175, col: 1, offset: 5557},
			expr: &choiceExpr{
				pos: position{line: 175, col: 14, offset: 5570},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 175, col: 14, offset: 5570},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 175, col: 14, offset: 5570},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 175, col: 14, offset: 5570},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 16, offset: 5572},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 22, offset: 5578},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 26, offset: 5582},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 28, offset: 5584},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 39, offset: 5595},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 175, col: 42, offset: 5598},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 46, offset: 5602},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 49, offset: 5605},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 54, offset: 5610},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 59, offset: 5615},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 181, col: 1, offset: 5734},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 181, col: 1, offset: 5734},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 181, col: 1, offset: 5734},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 181, col: 3, offset: 5736},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 9, offset: 5742},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 181, col: 13, offset: 5746},
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 14, offset: 5747},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 185, col: 1, offset: 5845},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 185, col: 1, offset: 5845},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 185, col: 1, offset: 5845},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 185, col: 3, offset: 5847},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 9, offset: 5853},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 185, col: 13, offset: 5857},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 185, col: 15, offset: 5859},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 26, offset: 5870},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 185, col: 29, offset: 5873},
									expr: &litMatcher{
										pos:        position{line: 185, col: 30, offset: 5874},
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
			pos:  position{line: 189, col: 1, offset: 5968},
			expr: &actionExpr{
				pos: position{line: 189, col: 12, offset: 5979},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 189, col: 12, offset: 5979},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 189, col: 12, offset: 5979},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 14, offset: 5981},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 20, offset: 5987},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 24, offset: 5991},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 26, offset: 5993},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 39, offset: 6006},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 189, col: 42, offset: 6009},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 46, offset: 6013},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 6016},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 53, offset: 6020},
								expr: &seqExpr{
									pos: position{line: 189, col: 54, offset: 6021},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 189, col: 54, offset: 6021},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 63, offset: 6030},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 67, offset: 6034},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 71, offset: 6038},
								expr: &seqExpr{
									pos: position{line: 189, col: 72, offset: 6039},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 189, col: 72, offset: 6039},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 189, col: 74, offset: 6041},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 79, offset: 6046},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 81, offset: 6048},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 96, offset: 6063},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 189, col: 100, offset: 6067},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 104, offset: 6071},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 107, offset: 6074},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 189, col: 118, offset: 6085},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 119, offset: 6086},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 131, offset: 6098},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 133, offset: 6100},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 137, offset: 6104},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 215, col: 1, offset: 6699},
			expr: &actionExpr{
				pos: position{line: 215, col: 8, offset: 6706},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 215, col: 8, offset: 6706},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 215, col: 12, offset: 6710},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 215, col: 12, offset: 6710},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 215, col: 21, offset: 6719},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 221, col: 1, offset: 6836},
			expr: &choiceExpr{
				pos: position{line: 221, col: 10, offset: 6845},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 221, col: 10, offset: 6845},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 221, col: 10, offset: 6845},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 221, col: 10, offset: 6845},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 221, col: 12, offset: 6847},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 17, offset: 6852},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 21, offset: 6856},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 26, offset: 6861},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 36, offset: 6871},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 39, offset: 6874},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 43, offset: 6878},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 45, offset: 6880},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 51, offset: 6886},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 52, offset: 6887},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 64, offset: 6899},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 67, offset: 6902},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 71, offset: 6906},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 74, offset: 6909},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 81, offset: 6916},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 84, offset: 6919},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 88, offset: 6923},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 90, offset: 6925},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 96, offset: 6931},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 97, offset: 6932},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 109, offset: 6944},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 112, offset: 6947},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 240, col: 1, offset: 7450},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 240, col: 1, offset: 7450},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 240, col: 1, offset: 7450},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 240, col: 3, offset: 7452},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 8, offset: 7457},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 240, col: 12, offset: 7461},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 17, offset: 7466},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 27, offset: 7476},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 240, col: 30, offset: 7479},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 34, offset: 7483},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 240, col: 36, offset: 7485},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 240, col: 42, offset: 7491},
										expr: &ruleRefExpr{
											pos:  position{line: 240, col: 43, offset: 7492},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 55, offset: 7504},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 240, col: 57, offset: 7506},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 61, offset: 7510},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 240, col: 64, offset: 7513},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 240, col: 71, offset: 7520},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 79, offset: 7528},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 252, col: 1, offset: 7858},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 252, col: 1, offset: 7858},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 252, col: 1, offset: 7858},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 252, col: 3, offset: 7860},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 8, offset: 7865},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 252, col: 12, offset: 7869},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 17, offset: 7874},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 27, offset: 7884},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 252, col: 30, offset: 7887},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 34, offset: 7891},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 252, col: 36, offset: 7893},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 252, col: 42, offset: 7899},
										expr: &ruleRefExpr{
											pos:  position{line: 252, col: 43, offset: 7900},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 55, offset: 7912},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 252, col: 58, offset: 7915},
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
			pos:  position{line: 264, col: 1, offset: 8213},
			expr: &choiceExpr{
				pos: position{line: 264, col: 8, offset: 8220},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 264, col: 8, offset: 8220},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 264, col: 8, offset: 8220},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 8, offset: 8220},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 10, offset: 8222},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 17, offset: 8229},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 28, offset: 8240},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 264, col: 32, offset: 8244},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 35, offset: 8247},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 264, col: 48, offset: 8260},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 53, offset: 8265},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 278, col: 1, offset: 8589},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 278, col: 1, offset: 8589},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 278, col: 1, offset: 8589},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 278, col: 3, offset: 8591},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 6, offset: 8594},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 278, col: 19, offset: 8607},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 24, offset: 8612},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 292, col: 1, offset: 8929},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 292, col: 1, offset: 8929},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 292, col: 1, offset: 8929},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 292, col: 3, offset: 8931},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 292, col: 6, offset: 8934},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 292, col: 19, offset: 8947},
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
			pos:  position{line: 299, col: 1, offset: 9118},
			expr: &actionExpr{
				pos: position{line: 299, col: 16, offset: 9133},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 299, col: 16, offset: 9133},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 299, col: 16, offset: 9133},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 299, col: 23, offset: 9140},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 299, col: 36, offset: 9153},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 299, col: 41, offset: 9158},
								expr: &seqExpr{
									pos: position{line: 299, col: 42, offset: 9159},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 299, col: 42, offset: 9159},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 299, col: 46, offset: 9163},
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
			pos:  position{line: 316, col: 1, offset: 9600},
			expr: &actionExpr{
				pos: position{line: 316, col: 12, offset: 9611},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 316, col: 12, offset: 9611},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 316, col: 12, offset: 9611},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 16, offset: 9615},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 316, col: 18, offset: 9617},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 316, col: 27, offset: 9626},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 35, offset: 9634},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 316, col: 37, offset: 9636},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 316, col: 42, offset: 9641},
								expr: &seqExpr{
									pos: position{line: 316, col: 43, offset: 9642},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 316, col: 43, offset: 9642},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 316, col: 47, offset: 9646},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 316, col: 49, offset: 9648},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 59, offset: 9658},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 316, col: 61, offset: 9660},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 334, col: 1, offset: 10082},
			expr: &actionExpr{
				pos: position{line: 334, col: 11, offset: 10092},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 334, col: 11, offset: 10092},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 334, col: 11, offset: 10092},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 16, offset: 10097},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 27, offset: 10108},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 29, offset: 10110},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 334, col: 34, offset: 10115},
								expr: &seqExpr{
									pos: position{line: 334, col: 35, offset: 10116},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 334, col: 35, offset: 10116},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 39, offset: 10120},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 41, offset: 10122},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 59, offset: 10140},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 355, col: 1, offset: 10679},
			expr: &choiceExpr{
				pos: position{line: 355, col: 18, offset: 10696},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 355, col: 18, offset: 10696},
						name: "AnyType",
					},
					&actionExpr{
						pos: position{line: 356, col: 1, offset: 10707},
						run: (*parser).callonTypeAnnotation3,
						expr: &seqExpr{
							pos: position{line: 356, col: 1, offset: 10707},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 356, col: 1, offset: 10707},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 8, offset: 10714},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 356, col: 11, offset: 10717},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 356, col: 16, offset: 10722},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 25, offset: 10731},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 356, col: 27, offset: 10733},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 32, offset: 10738},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 356, col: 34, offset: 10740},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 356, col: 38, offset: 10744},
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
			pos:  position{line: 365, col: 1, offset: 10960},
			expr: &choiceExpr{
				pos: position{line: 365, col: 11, offset: 10970},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 365, col: 11, offset: 10970},
						name: "ModuleName",
					},
					&ruleRefExpr{
						pos:  position{line: 365, col: 24, offset: 10983},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 365, col: 35, offset: 10994},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 367, col: 1, offset: 11009},
			expr: &choiceExpr{
				pos: position{line: 367, col: 13, offset: 11021},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 367, col: 13, offset: 11021},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 367, col: 13, offset: 11021},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 367, col: 13, offset: 11021},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 17, offset: 11025},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 19, offset: 11027},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 367, col: 28, offset: 11036},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 40, offset: 11048},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 42, offset: 11050},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 367, col: 47, offset: 11055},
										expr: &seqExpr{
											pos: position{line: 367, col: 48, offset: 11056},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 367, col: 48, offset: 11056},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 367, col: 52, offset: 11060},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 367, col: 54, offset: 11062},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 68, offset: 11076},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 367, col: 70, offset: 11078},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 1, offset: 11500},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 384, col: 1, offset: 11500},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 384, col: 1, offset: 11500},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 384, col: 5, offset: 11504},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 384, col: 7, offset: 11506},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 384, col: 16, offset: 11515},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 384, col: 21, offset: 11520},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 384, col: 23, offset: 11522},
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
			pos:  position{line: 389, col: 1, offset: 11627},
			expr: &actionExpr{
				pos: position{line: 389, col: 16, offset: 11642},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 389, col: 16, offset: 11642},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 389, col: 16, offset: 11642},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 389, col: 18, offset: 11644},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 389, col: 21, offset: 11647},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 389, col: 27, offset: 11653},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 389, col: 32, offset: 11658},
								expr: &seqExpr{
									pos: position{line: 389, col: 33, offset: 11659},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 389, col: 33, offset: 11659},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 37, offset: 11663},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 46, offset: 11672},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 50, offset: 11676},
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
			pos:  position{line: 409, col: 1, offset: 12282},
			expr: &choiceExpr{
				pos: position{line: 409, col: 9, offset: 12290},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 409, col: 9, offset: 12290},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 21, offset: 12302},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 37, offset: 12318},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 48, offset: 12329},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 60, offset: 12341},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 411, col: 1, offset: 12354},
			expr: &actionExpr{
				pos: position{line: 411, col: 13, offset: 12366},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 411, col: 13, offset: 12366},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 411, col: 13, offset: 12366},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 411, col: 15, offset: 12368},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 411, col: 21, offset: 12374},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 411, col: 35, offset: 12388},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 411, col: 40, offset: 12393},
								expr: &seqExpr{
									pos: position{line: 411, col: 41, offset: 12394},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 411, col: 41, offset: 12394},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 45, offset: 12398},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 61, offset: 12414},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 65, offset: 12418},
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
			pos:  position{line: 444, col: 1, offset: 13311},
			expr: &actionExpr{
				pos: position{line: 444, col: 17, offset: 13327},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 444, col: 17, offset: 13327},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 444, col: 17, offset: 13327},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 444, col: 19, offset: 13329},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 25, offset: 13335},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 444, col: 34, offset: 13344},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 444, col: 39, offset: 13349},
								expr: &seqExpr{
									pos: position{line: 444, col: 40, offset: 13350},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 444, col: 40, offset: 13350},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 44, offset: 13354},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 61, offset: 13371},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 65, offset: 13375},
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
			pos:  position{line: 476, col: 1, offset: 14262},
			expr: &actionExpr{
				pos: position{line: 476, col: 12, offset: 14273},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 476, col: 12, offset: 14273},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 476, col: 12, offset: 14273},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 476, col: 14, offset: 14275},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 476, col: 20, offset: 14281},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 476, col: 30, offset: 14291},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 476, col: 35, offset: 14296},
								expr: &seqExpr{
									pos: position{line: 476, col: 36, offset: 14297},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 476, col: 36, offset: 14297},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 40, offset: 14301},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 52, offset: 14313},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 56, offset: 14317},
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
			pos:  position{line: 508, col: 1, offset: 15205},
			expr: &actionExpr{
				pos: position{line: 508, col: 13, offset: 15217},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 508, col: 13, offset: 15217},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 508, col: 13, offset: 15217},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 508, col: 15, offset: 15219},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 21, offset: 15225},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 508, col: 33, offset: 15237},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 508, col: 38, offset: 15242},
								expr: &seqExpr{
									pos: position{line: 508, col: 39, offset: 15243},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 508, col: 39, offset: 15243},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 43, offset: 15247},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 56, offset: 15260},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 60, offset: 15264},
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
			pos:  position{line: 539, col: 1, offset: 16153},
			expr: &choiceExpr{
				pos: position{line: 539, col: 15, offset: 16167},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 539, col: 15, offset: 16167},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 539, col: 15, offset: 16167},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 539, col: 15, offset: 16167},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 539, col: 17, offset: 16169},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 21, offset: 16173},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 539, col: 24, offset: 16176},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 539, col: 30, offset: 16182},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 36, offset: 16188},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 539, col: 39, offset: 16191},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 542, col: 5, offset: 16314},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 544, col: 1, offset: 16321},
			expr: &choiceExpr{
				pos: position{line: 544, col: 12, offset: 16332},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 544, col: 12, offset: 16332},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 30, offset: 16350},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 49, offset: 16369},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 64, offset: 16384},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 546, col: 1, offset: 16397},
			expr: &actionExpr{
				pos: position{line: 546, col: 19, offset: 16415},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 546, col: 21, offset: 16417},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 546, col: 21, offset: 16417},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 546, col: 28, offset: 16424},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 550, col: 1, offset: 16506},
			expr: &actionExpr{
				pos: position{line: 550, col: 20, offset: 16525},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 550, col: 22, offset: 16527},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 550, col: 22, offset: 16527},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 29, offset: 16534},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 36, offset: 16541},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 42, offset: 16547},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 48, offset: 16553},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 56, offset: 16561},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 554, col: 1, offset: 16640},
			expr: &choiceExpr{
				pos: position{line: 554, col: 16, offset: 16655},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 554, col: 16, offset: 16655},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 554, col: 18, offset: 16657},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 554, col: 18, offset: 16657},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 554, col: 24, offset: 16663},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 557, col: 3, offset: 16746},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 557, col: 5, offset: 16748},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 560, col: 3, offset: 16828},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 560, col: 3, offset: 16828},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 564, col: 1, offset: 16909},
			expr: &actionExpr{
				pos: position{line: 564, col: 15, offset: 16923},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 564, col: 17, offset: 16925},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 564, col: 17, offset: 16925},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 564, col: 23, offset: 16931},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 568, col: 1, offset: 17013},
			expr: &choiceExpr{
				pos: position{line: 568, col: 9, offset: 17021},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 568, col: 9, offset: 17021},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 16, offset: 17028},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 31, offset: 17043},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 46, offset: 17058},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 570, col: 1, offset: 17065},
			expr: &choiceExpr{
				pos: position{line: 570, col: 14, offset: 17078},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 570, col: 14, offset: 17078},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 570, col: 29, offset: 17093},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 572, col: 1, offset: 17101},
			expr: &choiceExpr{
				pos: position{line: 572, col: 14, offset: 17114},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 572, col: 14, offset: 17114},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 572, col: 29, offset: 17129},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 574, col: 1, offset: 17141},
			expr: &actionExpr{
				pos: position{line: 574, col: 16, offset: 17156},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 574, col: 16, offset: 17156},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 574, col: 16, offset: 17156},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 20, offset: 17160},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 574, col: 22, offset: 17162},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 574, col: 28, offset: 17168},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 33, offset: 17173},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 574, col: 35, offset: 17175},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 574, col: 40, offset: 17180},
								expr: &seqExpr{
									pos: position{line: 574, col: 41, offset: 17181},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 574, col: 41, offset: 17181},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 45, offset: 17185},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 47, offset: 17187},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 52, offset: 17192},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 574, col: 56, offset: 17196},
							expr: &litMatcher{
								pos:        position{line: 574, col: 56, offset: 17196},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 61, offset: 17201},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 574, col: 63, offset: 17203},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 590, col: 1, offset: 17648},
			expr: &actionExpr{
				pos: position{line: 590, col: 19, offset: 17666},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 590, col: 19, offset: 17666},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 590, col: 19, offset: 17666},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 24, offset: 17671},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 590, col: 35, offset: 17682},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 590, col: 37, offset: 17684},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 42, offset: 17689},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 603, col: 1, offset: 17959},
			expr: &actionExpr{
				pos: position{line: 603, col: 18, offset: 17976},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 603, col: 18, offset: 17976},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 603, col: 18, offset: 17976},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 23, offset: 17981},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 34, offset: 17992},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 603, col: 36, offset: 17994},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 40, offset: 17998},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 603, col: 42, offset: 18000},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 52, offset: 18010},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 65, offset: 18023},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 603, col: 67, offset: 18025},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 71, offset: 18029},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 603, col: 73, offset: 18031},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 84, offset: 18042},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 603, col: 89, offset: 18047},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 603, col: 94, offset: 18052},
								expr: &seqExpr{
									pos: position{line: 603, col: 95, offset: 18053},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 603, col: 95, offset: 18053},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 99, offset: 18057},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 101, offset: 18059},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 114, offset: 18072},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 603, col: 116, offset: 18074},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 120, offset: 18078},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 122, offset: 18080},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 603, col: 130, offset: 18088},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 623, col: 1, offset: 18678},
			expr: &actionExpr{
				pos: position{line: 623, col: 17, offset: 18694},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 623, col: 17, offset: 18694},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 623, col: 17, offset: 18694},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 623, col: 22, offset: 18699},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 627, col: 1, offset: 18772},
			expr: &actionExpr{
				pos: position{line: 627, col: 16, offset: 18787},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 627, col: 16, offset: 18787},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 627, col: 16, offset: 18787},
							expr: &ruleRefExpr{
								pos:  position{line: 627, col: 17, offset: 18788},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 627, col: 27, offset: 18798},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 627, col: 27, offset: 18798},
									expr: &charClassMatcher{
										pos:        position{line: 627, col: 27, offset: 18798},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 627, col: 34, offset: 18805},
									expr: &charClassMatcher{
										pos:        position{line: 627, col: 34, offset: 18805},
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
			pos:  position{line: 631, col: 1, offset: 18880},
			expr: &actionExpr{
				pos: position{line: 631, col: 14, offset: 18893},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 631, col: 15, offset: 18894},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 631, col: 15, offset: 18894},
							expr: &charClassMatcher{
								pos:        position{line: 631, col: 15, offset: 18894},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 631, col: 22, offset: 18901},
							expr: &charClassMatcher{
								pos:        position{line: 631, col: 22, offset: 18901},
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
			pos:  position{line: 635, col: 1, offset: 18976},
			expr: &choiceExpr{
				pos: position{line: 635, col: 9, offset: 18984},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 635, col: 9, offset: 18984},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 635, col: 9, offset: 18984},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 635, col: 9, offset: 18984},
									expr: &litMatcher{
										pos:        position{line: 635, col: 9, offset: 18984},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 635, col: 14, offset: 18989},
									expr: &charClassMatcher{
										pos:        position{line: 635, col: 14, offset: 18989},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 635, col: 21, offset: 18996},
									expr: &litMatcher{
										pos:        position{line: 635, col: 22, offset: 18997},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 642, col: 3, offset: 19172},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 642, col: 3, offset: 19172},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 642, col: 3, offset: 19172},
									expr: &litMatcher{
										pos:        position{line: 642, col: 3, offset: 19172},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 642, col: 8, offset: 19177},
									expr: &charClassMatcher{
										pos:        position{line: 642, col: 8, offset: 19177},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 642, col: 15, offset: 19184},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 642, col: 19, offset: 19188},
									expr: &charClassMatcher{
										pos:        position{line: 642, col: 19, offset: 19188},
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
						pos: position{line: 649, col: 3, offset: 19377},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 649, col: 3, offset: 19377},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 653, col: 3, offset: 19462},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 653, col: 3, offset: 19462},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 656, col: 3, offset: 19548},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 657, col: 3, offset: 19555},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 658, col: 3, offset: 19571},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 658, col: 3, offset: 19571},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 658, col: 3, offset: 19571},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 658, col: 7, offset: 19575},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 658, col: 12, offset: 19580},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 658, col: 12, offset: 19580},
												expr: &ruleRefExpr{
													pos:  position{line: 658, col: 13, offset: 19581},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 658, col: 25, offset: 19593,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 658, col: 28, offset: 19596},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 660, col: 5, offset: 19688},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 660, col: 20, offset: 19703},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 660, col: 37, offset: 19720},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 662, col: 1, offset: 19737},
			expr: &actionExpr{
				pos: position{line: 662, col: 8, offset: 19744},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 662, col: 8, offset: 19744},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 666, col: 1, offset: 19807},
			expr: &actionExpr{
				pos: position{line: 666, col: 17, offset: 19823},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 666, col: 17, offset: 19823},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 666, col: 17, offset: 19823},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 666, col: 21, offset: 19827},
							expr: &seqExpr{
								pos: position{line: 666, col: 22, offset: 19828},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 666, col: 22, offset: 19828},
										expr: &ruleRefExpr{
											pos:  position{line: 666, col: 23, offset: 19829},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 666, col: 35, offset: 19841,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 666, col: 39, offset: 19845},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 674, col: 1, offset: 20028},
			expr: &actionExpr{
				pos: position{line: 674, col: 10, offset: 20037},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 674, col: 11, offset: 20038},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 678, col: 1, offset: 20093},
			expr: &seqExpr{
				pos: position{line: 678, col: 12, offset: 20104},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 678, col: 13, offset: 20105},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 678, col: 13, offset: 20105},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 21, offset: 20113},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 28, offset: 20120},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 37, offset: 20129},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 48, offset: 20140},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 57, offset: 20149},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 66, offset: 20158},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 76, offset: 20168},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 88, offset: 20180},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 678, col: 97, offset: 20189},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 678, col: 107, offset: 20199},
						expr: &oneOrMoreExpr{
							pos: position{line: 678, col: 108, offset: 20200},
							expr: &charClassMatcher{
								pos:        position{line: 678, col: 108, offset: 20200},
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
			pos:  position{line: 680, col: 1, offset: 20208},
			expr: &choiceExpr{
				pos: position{line: 680, col: 12, offset: 20219},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 680, col: 12, offset: 20219},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 680, col: 14, offset: 20221},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 680, col: 14, offset: 20221},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 24, offset: 20231},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 32, offset: 20239},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 41, offset: 20248},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 52, offset: 20259},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 61, offset: 20268},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 70, offset: 20277},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 80, offset: 20287},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 683, col: 3, offset: 20385},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 685, col: 1, offset: 20391},
			expr: &charClassMatcher{
				pos:        position{line: 685, col: 15, offset: 20405},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 687, col: 1, offset: 20421},
			expr: &choiceExpr{
				pos: position{line: 687, col: 18, offset: 20438},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 687, col: 18, offset: 20438},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 687, col: 37, offset: 20457},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 689, col: 1, offset: 20472},
			expr: &charClassMatcher{
				pos:        position{line: 689, col: 20, offset: 20491},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 691, col: 1, offset: 20504},
			expr: &charClassMatcher{
				pos:        position{line: 691, col: 16, offset: 20519},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 693, col: 1, offset: 20526},
			expr: &charClassMatcher{
				pos:        position{line: 693, col: 23, offset: 20548},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 695, col: 1, offset: 20555},
			expr: &charClassMatcher{
				pos:        position{line: 695, col: 12, offset: 20566},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 697, col: 1, offset: 20577},
			expr: &choiceExpr{
				pos: position{line: 697, col: 22, offset: 20598},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 697, col: 22, offset: 20598},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 697, col: 33, offset: 20609},
						expr: &charClassMatcher{
							pos:        position{line: 697, col: 33, offset: 20609},
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
			pos:         position{line: 699, col: 1, offset: 20621},
			expr: &choiceExpr{
				pos: position{line: 699, col: 21, offset: 20641},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 699, col: 21, offset: 20641},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 699, col: 32, offset: 20652},
						expr: &charClassMatcher{
							pos:        position{line: 699, col: 32, offset: 20652},
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
			pos:         position{line: 701, col: 1, offset: 20664},
			expr: &oneOrMoreExpr{
				pos: position{line: 701, col: 33, offset: 20696},
				expr: &charClassMatcher{
					pos:        position{line: 701, col: 33, offset: 20696},
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
			pos:         position{line: 703, col: 1, offset: 20704},
			expr: &zeroOrMoreExpr{
				pos: position{line: 703, col: 32, offset: 20735},
				expr: &charClassMatcher{
					pos:        position{line: 703, col: 32, offset: 20735},
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
			pos:         position{line: 705, col: 1, offset: 20743},
			expr: &choiceExpr{
				pos: position{line: 705, col: 15, offset: 20757},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 705, col: 15, offset: 20757},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 705, col: 26, offset: 20768},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 705, col: 26, offset: 20768},
								expr: &charClassMatcher{
									pos:        position{line: 705, col: 26, offset: 20768},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 705, col: 35, offset: 20777},
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
			pos:  position{line: 707, col: 1, offset: 20783},
			expr: &oneOrMoreExpr{
				pos: position{line: 707, col: 12, offset: 20794},
				expr: &ruleRefExpr{
					pos:  position{line: 707, col: 13, offset: 20795},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 709, col: 1, offset: 20806},
			expr: &choiceExpr{
				pos: position{line: 709, col: 11, offset: 20816},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 709, col: 11, offset: 20816},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 11, offset: 20816},
								expr: &charClassMatcher{
									pos:        position{line: 709, col: 11, offset: 20816},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 22, offset: 20827},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 27, offset: 20832},
								expr: &seqExpr{
									pos: position{line: 709, col: 28, offset: 20833},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 709, col: 28, offset: 20833},
											expr: &charClassMatcher{
												pos:        position{line: 709, col: 29, offset: 20834},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 709, col: 34, offset: 20839,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 709, col: 38, offset: 20843},
								expr: &litMatcher{
									pos:        position{line: 709, col: 39, offset: 20844},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 709, col: 46, offset: 20851},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 46, offset: 20851},
								expr: &charClassMatcher{
									pos:        position{line: 709, col: 46, offset: 20851},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 57, offset: 20862},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 62, offset: 20867},
								expr: &seqExpr{
									pos: position{line: 709, col: 63, offset: 20868},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 709, col: 63, offset: 20868},
											expr: &litMatcher{
												pos:        position{line: 709, col: 64, offset: 20869},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 709, col: 69, offset: 20874,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 73, offset: 20878},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 78, offset: 20883},
								expr: &charClassMatcher{
									pos:        position{line: 709, col: 78, offset: 20883},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 84, offset: 20889},
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
			pos:  position{line: 711, col: 1, offset: 20895},
			expr: &notExpr{
				pos: position{line: 711, col: 7, offset: 20901},
				expr: &anyMatcher{
					line: 711, col: 8, offset: 20902,
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
