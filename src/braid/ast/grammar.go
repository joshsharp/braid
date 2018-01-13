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
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "ExternType",
			pos:  position{line: 53, col: 1, offset: 1511},
			expr: &choiceExpr{
				pos: position{line: 53, col: 14, offset: 1524},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 53, col: 14, offset: 1524},
						run: (*parser).callonExternType2,
						expr: &seqExpr{
							pos: position{line: 53, col: 14, offset: 1524},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 53, col: 14, offset: 1524},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 16, offset: 1526},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 25, offset: 1535},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 53, col: 29, offset: 1539},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 36, offset: 1546},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 39, offset: 1549},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 53, col: 44, offset: 1554},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 55, offset: 1565},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 57, offset: 1567},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 61, offset: 1571},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 5, offset: 1577},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 16, offset: 1588},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 30, offset: 1602},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 32, offset: 1604},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 36, offset: 1608},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 38, offset: 1610},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 44, offset: 1616},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 60, offset: 1632},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 62, offset: 1634},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 54, col: 67, offset: 1639},
										expr: &seqExpr{
											pos: position{line: 54, col: 68, offset: 1640},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 54, col: 68, offset: 1640},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 72, offset: 1644},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 74, offset: 1646},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 54, col: 90, offset: 1662},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 54, col: 94, offset: 1666},
									expr: &litMatcher{
										pos:        position{line: 54, col: 94, offset: 1666},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 99, offset: 1671},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 54, col: 101, offset: 1673},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 105, offset: 1677},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 73, col: 1, offset: 2226},
						run: (*parser).callonExternType34,
						expr: &seqExpr{
							pos: position{line: 73, col: 1, offset: 2226},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 73, col: 1, offset: 2226},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 3, offset: 2228},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 12, offset: 2237},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 73, col: 16, offset: 2241},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 23, offset: 2248},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 26, offset: 2251},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 73, col: 31, offset: 2256},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 42, offset: 2267},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 73, col: 44, offset: 2269},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 48, offset: 2273},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 74, col: 5, offset: 2279},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 74, col: 16, offset: 2290},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 74, col: 30, offset: 2304},
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
			pos:  position{line: 82, col: 1, offset: 2499},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2510},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 82, col: 12, offset: 2510},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 82, col: 12, offset: 2510},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 82, col: 12, offset: 2510},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 14, offset: 2512},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 21, offset: 2519},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 82, col: 24, offset: 2522},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 29, offset: 2527},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 82, col: 40, offset: 2538},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 82, col: 47, offset: 2545},
										expr: &seqExpr{
											pos: position{line: 82, col: 48, offset: 2546},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 48, offset: 2546},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 51, offset: 2549},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 67, offset: 2565},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 82, col: 69, offset: 2567},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 82, col: 73, offset: 2571},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 82, col: 79, offset: 2577},
										expr: &seqExpr{
											pos: position{line: 82, col: 80, offset: 2578},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 82, col: 80, offset: 2578},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 82, col: 83, offset: 2581},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 100, offset: 2598},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 101, col: 1, offset: 3094},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 101, col: 1, offset: 3094},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 101, col: 1, offset: 3094},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 3, offset: 3096},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 10, offset: 3103},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 13, offset: 3106},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 18, offset: 3111},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 101, col: 29, offset: 3122},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 36, offset: 3129},
										expr: &seqExpr{
											pos: position{line: 101, col: 37, offset: 3130},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 101, col: 37, offset: 3130},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 40, offset: 3133},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 56, offset: 3149},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 58, offset: 3151},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 62, offset: 3155},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 5, offset: 3161},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 9, offset: 3165},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 102, col: 11, offset: 3167},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 102, col: 17, offset: 3173},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 33, offset: 3189},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 102, col: 35, offset: 3191},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 102, col: 40, offset: 3196},
										expr: &seqExpr{
											pos: position{line: 102, col: 41, offset: 3197},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 102, col: 41, offset: 3197},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 45, offset: 3201},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 47, offset: 3203},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 102, col: 63, offset: 3219},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 102, col: 67, offset: 3223},
									expr: &litMatcher{
										pos:        position{line: 102, col: 67, offset: 3223},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 72, offset: 3228},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 74, offset: 3230},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 78, offset: 3234},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 120, col: 1, offset: 3719},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 120, col: 1, offset: 3719},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 120, col: 1, offset: 3719},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 3, offset: 3721},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 10, offset: 3728},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 13, offset: 3731},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 120, col: 18, offset: 3736},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 120, col: 29, offset: 3747},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 120, col: 36, offset: 3754},
										expr: &seqExpr{
											pos: position{line: 120, col: 37, offset: 3755},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 120, col: 37, offset: 3755},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 120, col: 40, offset: 3758},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 56, offset: 3774},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 120, col: 58, offset: 3776},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 120, col: 62, offset: 3780},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 120, col: 64, offset: 3782},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 120, col: 69, offset: 3787},
										expr: &ruleRefExpr{
											pos:  position{line: 120, col: 70, offset: 3788},
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
			pos:  position{line: 135, col: 1, offset: 4195},
			expr: &actionExpr{
				pos: position{line: 135, col: 19, offset: 4213},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 135, col: 19, offset: 4213},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 135, col: 19, offset: 4213},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 24, offset: 4218},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 135, col: 37, offset: 4231},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 135, col: 39, offset: 4233},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 135, col: 43, offset: 4237},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 135, col: 45, offset: 4239},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 48, offset: 4242},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 139, col: 1, offset: 4343},
			expr: &choiceExpr{
				pos: position{line: 139, col: 22, offset: 4364},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 139, col: 22, offset: 4364},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 139, col: 22, offset: 4364},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 139, col: 22, offset: 4364},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 26, offset: 4368},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 28, offset: 4370},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 139, col: 33, offset: 4375},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 44, offset: 4386},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 139, col: 46, offset: 4388},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 50, offset: 4392},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 52, offset: 4394},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 139, col: 58, offset: 4400},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 74, offset: 4416},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 139, col: 76, offset: 4418},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 139, col: 81, offset: 4423},
										expr: &seqExpr{
											pos: position{line: 139, col: 82, offset: 4424},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 139, col: 82, offset: 4424},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 86, offset: 4428},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 88, offset: 4430},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 104, offset: 4446},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 139, col: 108, offset: 4450},
									expr: &litMatcher{
										pos:        position{line: 139, col: 108, offset: 4450},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 113, offset: 4455},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 139, col: 115, offset: 4457},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 139, col: 119, offset: 4461},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 158, col: 1, offset: 5066},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 158, col: 1, offset: 5066},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 158, col: 1, offset: 5066},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 158, col: 5, offset: 5070},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 158, col: 7, offset: 5072},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 158, col: 12, offset: 5077},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 158, col: 23, offset: 5088},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 158, col: 28, offset: 5093},
										expr: &seqExpr{
											pos: position{line: 158, col: 29, offset: 5094},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 158, col: 29, offset: 5094},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 158, col: 32, offset: 5097},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 158, col: 49, offset: 5114},
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
			pos:  position{line: 175, col: 1, offset: 5551},
			expr: &choiceExpr{
				pos: position{line: 175, col: 14, offset: 5564},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 175, col: 14, offset: 5564},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 175, col: 14, offset: 5564},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 175, col: 14, offset: 5564},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 16, offset: 5566},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 22, offset: 5572},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 26, offset: 5576},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 28, offset: 5578},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 39, offset: 5589},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 175, col: 42, offset: 5592},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 46, offset: 5596},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 49, offset: 5599},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 54, offset: 5604},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 59, offset: 5609},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 181, col: 1, offset: 5728},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 181, col: 1, offset: 5728},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 181, col: 1, offset: 5728},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 181, col: 3, offset: 5730},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 9, offset: 5736},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 181, col: 13, offset: 5740},
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 14, offset: 5741},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 185, col: 1, offset: 5839},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 185, col: 1, offset: 5839},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 185, col: 1, offset: 5839},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 185, col: 3, offset: 5841},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 9, offset: 5847},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 185, col: 13, offset: 5851},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 185, col: 15, offset: 5853},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 26, offset: 5864},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 185, col: 29, offset: 5867},
									expr: &litMatcher{
										pos:        position{line: 185, col: 30, offset: 5868},
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
			pos:  position{line: 189, col: 1, offset: 5962},
			expr: &actionExpr{
				pos: position{line: 189, col: 12, offset: 5973},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 189, col: 12, offset: 5973},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 189, col: 12, offset: 5973},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 14, offset: 5975},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 20, offset: 5981},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 24, offset: 5985},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 26, offset: 5987},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 39, offset: 6000},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 189, col: 42, offset: 6003},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 46, offset: 6007},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 6010},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 53, offset: 6014},
								expr: &seqExpr{
									pos: position{line: 189, col: 54, offset: 6015},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 189, col: 54, offset: 6015},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 63, offset: 6024},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 67, offset: 6028},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 71, offset: 6032},
								expr: &seqExpr{
									pos: position{line: 189, col: 72, offset: 6033},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 189, col: 72, offset: 6033},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 189, col: 74, offset: 6035},
											val:        "->",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 79, offset: 6040},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 81, offset: 6042},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 189, col: 96, offset: 6057},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 189, col: 100, offset: 6061},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 104, offset: 6065},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 107, offset: 6068},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 189, col: 118, offset: 6079},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 119, offset: 6080},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 131, offset: 6092},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 133, offset: 6094},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 137, offset: 6098},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 215, col: 1, offset: 6693},
			expr: &actionExpr{
				pos: position{line: 215, col: 8, offset: 6700},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 215, col: 8, offset: 6700},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 215, col: 12, offset: 6704},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 215, col: 12, offset: 6704},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 215, col: 21, offset: 6713},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 221, col: 1, offset: 6830},
			expr: &choiceExpr{
				pos: position{line: 221, col: 10, offset: 6839},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 221, col: 10, offset: 6839},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 221, col: 10, offset: 6839},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 221, col: 10, offset: 6839},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 221, col: 12, offset: 6841},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 17, offset: 6846},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 21, offset: 6850},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 26, offset: 6855},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 36, offset: 6865},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 39, offset: 6868},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 43, offset: 6872},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 45, offset: 6874},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 51, offset: 6880},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 52, offset: 6881},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 64, offset: 6893},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 67, offset: 6896},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 71, offset: 6900},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 74, offset: 6903},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 81, offset: 6910},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 221, col: 84, offset: 6913},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 88, offset: 6917},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 90, offset: 6919},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 96, offset: 6925},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 97, offset: 6926},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 109, offset: 6938},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 112, offset: 6941},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 240, col: 1, offset: 7444},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 240, col: 1, offset: 7444},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 240, col: 1, offset: 7444},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 240, col: 3, offset: 7446},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 8, offset: 7451},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 240, col: 12, offset: 7455},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 17, offset: 7460},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 27, offset: 7470},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 240, col: 30, offset: 7473},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 34, offset: 7477},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 240, col: 36, offset: 7479},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 240, col: 42, offset: 7485},
										expr: &ruleRefExpr{
											pos:  position{line: 240, col: 43, offset: 7486},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 55, offset: 7498},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 240, col: 57, offset: 7500},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 61, offset: 7504},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 240, col: 64, offset: 7507},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 240, col: 71, offset: 7514},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 79, offset: 7522},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 252, col: 1, offset: 7852},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 252, col: 1, offset: 7852},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 252, col: 1, offset: 7852},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 252, col: 3, offset: 7854},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 8, offset: 7859},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 252, col: 12, offset: 7863},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 17, offset: 7868},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 27, offset: 7878},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 252, col: 30, offset: 7881},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 34, offset: 7885},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 252, col: 36, offset: 7887},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 252, col: 42, offset: 7893},
										expr: &ruleRefExpr{
											pos:  position{line: 252, col: 43, offset: 7894},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 55, offset: 7906},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 252, col: 58, offset: 7909},
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
			pos:  position{line: 264, col: 1, offset: 8207},
			expr: &choiceExpr{
				pos: position{line: 264, col: 8, offset: 8214},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 264, col: 8, offset: 8214},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 264, col: 8, offset: 8214},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 8, offset: 8214},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 10, offset: 8216},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 17, offset: 8223},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 28, offset: 8234},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 264, col: 32, offset: 8238},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 35, offset: 8241},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 264, col: 48, offset: 8254},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 53, offset: 8259},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 278, col: 1, offset: 8583},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 278, col: 1, offset: 8583},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 278, col: 1, offset: 8583},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 278, col: 3, offset: 8585},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 6, offset: 8588},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 278, col: 19, offset: 8601},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 24, offset: 8606},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 292, col: 1, offset: 8923},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 292, col: 1, offset: 8923},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 292, col: 1, offset: 8923},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 292, col: 3, offset: 8925},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 292, col: 6, offset: 8928},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 292, col: 19, offset: 8941},
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
			pos:  position{line: 299, col: 1, offset: 9112},
			expr: &actionExpr{
				pos: position{line: 299, col: 16, offset: 9127},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 299, col: 16, offset: 9127},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 299, col: 16, offset: 9127},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 299, col: 23, offset: 9134},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 299, col: 36, offset: 9147},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 299, col: 41, offset: 9152},
								expr: &seqExpr{
									pos: position{line: 299, col: 42, offset: 9153},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 299, col: 42, offset: 9153},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 299, col: 46, offset: 9157},
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
			pos:  position{line: 316, col: 1, offset: 9594},
			expr: &actionExpr{
				pos: position{line: 316, col: 12, offset: 9605},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 316, col: 12, offset: 9605},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 316, col: 12, offset: 9605},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 16, offset: 9609},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 316, col: 18, offset: 9611},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 316, col: 27, offset: 9620},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 35, offset: 9628},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 316, col: 37, offset: 9630},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 316, col: 42, offset: 9635},
								expr: &seqExpr{
									pos: position{line: 316, col: 43, offset: 9636},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 316, col: 43, offset: 9636},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 316, col: 47, offset: 9640},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 316, col: 49, offset: 9642},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 316, col: 59, offset: 9652},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 316, col: 61, offset: 9654},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 334, col: 1, offset: 10076},
			expr: &actionExpr{
				pos: position{line: 334, col: 11, offset: 10086},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 334, col: 11, offset: 10086},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 334, col: 11, offset: 10086},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 16, offset: 10091},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 27, offset: 10102},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 29, offset: 10104},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 334, col: 34, offset: 10109},
								expr: &seqExpr{
									pos: position{line: 334, col: 35, offset: 10110},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 334, col: 35, offset: 10110},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 39, offset: 10114},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 41, offset: 10116},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 59, offset: 10134},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 355, col: 1, offset: 10673},
			expr: &choiceExpr{
				pos: position{line: 355, col: 18, offset: 10690},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 355, col: 18, offset: 10690},
						name: "AnyType",
					},
					&ruleRefExpr{
						pos:  position{line: 355, col: 28, offset: 10700},
						name: "ModuleName",
					},
					&actionExpr{
						pos: position{line: 356, col: 1, offset: 10713},
						run: (*parser).callonTypeAnnotation4,
						expr: &seqExpr{
							pos: position{line: 356, col: 1, offset: 10713},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 356, col: 1, offset: 10713},
									val:        "func",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 8, offset: 10720},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 356, col: 11, offset: 10723},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 356, col: 16, offset: 10728},
										name: "ArgsDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 25, offset: 10737},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 356, col: 27, offset: 10739},
									val:        "->",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 356, col: 32, offset: 10744},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 356, col: 34, offset: 10746},
									label: "ret",
									expr: &ruleRefExpr{
										pos:  position{line: 356, col: 38, offset: 10750},
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
			pos:  position{line: 365, col: 1, offset: 10966},
			expr: &choiceExpr{
				pos: position{line: 365, col: 11, offset: 10976},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 365, col: 11, offset: 10976},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 365, col: 22, offset: 10987},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 367, col: 1, offset: 11002},
			expr: &choiceExpr{
				pos: position{line: 367, col: 13, offset: 11014},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 367, col: 13, offset: 11014},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 367, col: 13, offset: 11014},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 367, col: 13, offset: 11014},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 17, offset: 11018},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 19, offset: 11020},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 367, col: 28, offset: 11029},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 40, offset: 11041},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 367, col: 42, offset: 11043},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 367, col: 47, offset: 11048},
										expr: &seqExpr{
											pos: position{line: 367, col: 48, offset: 11049},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 367, col: 48, offset: 11049},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 367, col: 52, offset: 11053},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 367, col: 54, offset: 11055},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 367, col: 68, offset: 11069},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 367, col: 70, offset: 11071},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 1, offset: 11493},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 384, col: 1, offset: 11493},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 384, col: 1, offset: 11493},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 384, col: 5, offset: 11497},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 384, col: 7, offset: 11499},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 384, col: 16, offset: 11508},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 384, col: 21, offset: 11513},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 384, col: 23, offset: 11515},
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
			pos:  position{line: 389, col: 1, offset: 11620},
			expr: &actionExpr{
				pos: position{line: 389, col: 16, offset: 11635},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 389, col: 16, offset: 11635},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 389, col: 16, offset: 11635},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 389, col: 18, offset: 11637},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 389, col: 21, offset: 11640},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 389, col: 27, offset: 11646},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 389, col: 32, offset: 11651},
								expr: &seqExpr{
									pos: position{line: 389, col: 33, offset: 11652},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 389, col: 33, offset: 11652},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 37, offset: 11656},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 46, offset: 11665},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 389, col: 50, offset: 11669},
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
			pos:  position{line: 409, col: 1, offset: 12275},
			expr: &choiceExpr{
				pos: position{line: 409, col: 9, offset: 12283},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 409, col: 9, offset: 12283},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 21, offset: 12295},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 37, offset: 12311},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 48, offset: 12322},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 60, offset: 12334},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 411, col: 1, offset: 12347},
			expr: &actionExpr{
				pos: position{line: 411, col: 13, offset: 12359},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 411, col: 13, offset: 12359},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 411, col: 13, offset: 12359},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 411, col: 15, offset: 12361},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 411, col: 21, offset: 12367},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 411, col: 35, offset: 12381},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 411, col: 40, offset: 12386},
								expr: &seqExpr{
									pos: position{line: 411, col: 41, offset: 12387},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 411, col: 41, offset: 12387},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 45, offset: 12391},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 61, offset: 12407},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 411, col: 65, offset: 12411},
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
			pos:  position{line: 444, col: 1, offset: 13304},
			expr: &actionExpr{
				pos: position{line: 444, col: 17, offset: 13320},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 444, col: 17, offset: 13320},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 444, col: 17, offset: 13320},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 444, col: 19, offset: 13322},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 25, offset: 13328},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 444, col: 34, offset: 13337},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 444, col: 39, offset: 13342},
								expr: &seqExpr{
									pos: position{line: 444, col: 40, offset: 13343},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 444, col: 40, offset: 13343},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 44, offset: 13347},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 61, offset: 13364},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 444, col: 65, offset: 13368},
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
			pos:  position{line: 476, col: 1, offset: 14255},
			expr: &actionExpr{
				pos: position{line: 476, col: 12, offset: 14266},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 476, col: 12, offset: 14266},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 476, col: 12, offset: 14266},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 476, col: 14, offset: 14268},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 476, col: 20, offset: 14274},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 476, col: 30, offset: 14284},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 476, col: 35, offset: 14289},
								expr: &seqExpr{
									pos: position{line: 476, col: 36, offset: 14290},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 476, col: 36, offset: 14290},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 40, offset: 14294},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 52, offset: 14306},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 476, col: 56, offset: 14310},
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
			pos:  position{line: 508, col: 1, offset: 15198},
			expr: &actionExpr{
				pos: position{line: 508, col: 13, offset: 15210},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 508, col: 13, offset: 15210},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 508, col: 13, offset: 15210},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 508, col: 15, offset: 15212},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 21, offset: 15218},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 508, col: 33, offset: 15230},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 508, col: 38, offset: 15235},
								expr: &seqExpr{
									pos: position{line: 508, col: 39, offset: 15236},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 508, col: 39, offset: 15236},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 43, offset: 15240},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 56, offset: 15253},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 508, col: 60, offset: 15257},
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
			pos:  position{line: 539, col: 1, offset: 16146},
			expr: &choiceExpr{
				pos: position{line: 539, col: 15, offset: 16160},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 539, col: 15, offset: 16160},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 539, col: 15, offset: 16160},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 539, col: 15, offset: 16160},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 539, col: 17, offset: 16162},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 21, offset: 16166},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 539, col: 24, offset: 16169},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 539, col: 30, offset: 16175},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 36, offset: 16181},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 539, col: 39, offset: 16184},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 542, col: 5, offset: 16307},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 544, col: 1, offset: 16314},
			expr: &choiceExpr{
				pos: position{line: 544, col: 12, offset: 16325},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 544, col: 12, offset: 16325},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 30, offset: 16343},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 49, offset: 16362},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 544, col: 64, offset: 16377},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 546, col: 1, offset: 16390},
			expr: &actionExpr{
				pos: position{line: 546, col: 19, offset: 16408},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 546, col: 21, offset: 16410},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 546, col: 21, offset: 16410},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 546, col: 28, offset: 16417},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 550, col: 1, offset: 16499},
			expr: &actionExpr{
				pos: position{line: 550, col: 20, offset: 16518},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 550, col: 22, offset: 16520},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 550, col: 22, offset: 16520},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 29, offset: 16527},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 36, offset: 16534},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 42, offset: 16540},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 48, offset: 16546},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 550, col: 56, offset: 16554},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 554, col: 1, offset: 16633},
			expr: &choiceExpr{
				pos: position{line: 554, col: 16, offset: 16648},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 554, col: 16, offset: 16648},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 554, col: 18, offset: 16650},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 554, col: 18, offset: 16650},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 554, col: 24, offset: 16656},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 557, col: 3, offset: 16739},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 557, col: 5, offset: 16741},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 560, col: 3, offset: 16821},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 560, col: 3, offset: 16821},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 564, col: 1, offset: 16902},
			expr: &actionExpr{
				pos: position{line: 564, col: 15, offset: 16916},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 564, col: 17, offset: 16918},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 564, col: 17, offset: 16918},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 564, col: 23, offset: 16924},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 568, col: 1, offset: 17006},
			expr: &choiceExpr{
				pos: position{line: 568, col: 9, offset: 17014},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 568, col: 9, offset: 17014},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 16, offset: 17021},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 31, offset: 17036},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 568, col: 46, offset: 17051},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 570, col: 1, offset: 17058},
			expr: &choiceExpr{
				pos: position{line: 570, col: 14, offset: 17071},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 570, col: 14, offset: 17071},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 570, col: 29, offset: 17086},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 572, col: 1, offset: 17094},
			expr: &choiceExpr{
				pos: position{line: 572, col: 14, offset: 17107},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 572, col: 14, offset: 17107},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 572, col: 29, offset: 17122},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 574, col: 1, offset: 17134},
			expr: &actionExpr{
				pos: position{line: 574, col: 16, offset: 17149},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 574, col: 16, offset: 17149},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 574, col: 16, offset: 17149},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 20, offset: 17153},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 574, col: 22, offset: 17155},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 574, col: 28, offset: 17161},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 33, offset: 17166},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 574, col: 35, offset: 17168},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 574, col: 40, offset: 17173},
								expr: &seqExpr{
									pos: position{line: 574, col: 41, offset: 17174},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 574, col: 41, offset: 17174},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 45, offset: 17178},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 47, offset: 17180},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 574, col: 52, offset: 17185},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 574, col: 56, offset: 17189},
							expr: &litMatcher{
								pos:        position{line: 574, col: 56, offset: 17189},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 574, col: 61, offset: 17194},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 574, col: 63, offset: 17196},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 590, col: 1, offset: 17641},
			expr: &actionExpr{
				pos: position{line: 590, col: 19, offset: 17659},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 590, col: 19, offset: 17659},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 590, col: 19, offset: 17659},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 24, offset: 17664},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 590, col: 35, offset: 17675},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 590, col: 37, offset: 17677},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 42, offset: 17682},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 603, col: 1, offset: 17952},
			expr: &actionExpr{
				pos: position{line: 603, col: 18, offset: 17969},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 603, col: 18, offset: 17969},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 603, col: 18, offset: 17969},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 23, offset: 17974},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 34, offset: 17985},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 603, col: 36, offset: 17987},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 40, offset: 17991},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 603, col: 42, offset: 17993},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 52, offset: 18003},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 65, offset: 18016},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 603, col: 67, offset: 18018},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 603, col: 71, offset: 18022},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 603, col: 73, offset: 18024},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 603, col: 84, offset: 18035},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 603, col: 89, offset: 18040},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 603, col: 94, offset: 18045},
								expr: &seqExpr{
									pos: position{line: 603, col: 95, offset: 18046},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 603, col: 95, offset: 18046},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 99, offset: 18050},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 101, offset: 18052},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 114, offset: 18065},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 603, col: 116, offset: 18067},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 120, offset: 18071},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 603, col: 122, offset: 18073},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 603, col: 130, offset: 18081},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 623, col: 1, offset: 18665},
			expr: &actionExpr{
				pos: position{line: 623, col: 17, offset: 18681},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 623, col: 17, offset: 18681},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 623, col: 17, offset: 18681},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 623, col: 22, offset: 18686},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 627, col: 1, offset: 18759},
			expr: &actionExpr{
				pos: position{line: 627, col: 16, offset: 18774},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 627, col: 16, offset: 18774},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 627, col: 16, offset: 18774},
							expr: &ruleRefExpr{
								pos:  position{line: 627, col: 17, offset: 18775},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 627, col: 27, offset: 18785},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 627, col: 27, offset: 18785},
									expr: &charClassMatcher{
										pos:        position{line: 627, col: 27, offset: 18785},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 627, col: 34, offset: 18792},
									expr: &charClassMatcher{
										pos:        position{line: 627, col: 34, offset: 18792},
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
			pos:  position{line: 631, col: 1, offset: 18867},
			expr: &actionExpr{
				pos: position{line: 631, col: 14, offset: 18880},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 631, col: 15, offset: 18881},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 631, col: 15, offset: 18881},
							expr: &charClassMatcher{
								pos:        position{line: 631, col: 15, offset: 18881},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 631, col: 22, offset: 18888},
							expr: &charClassMatcher{
								pos:        position{line: 631, col: 22, offset: 18888},
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
			pos:  position{line: 635, col: 1, offset: 18963},
			expr: &choiceExpr{
				pos: position{line: 635, col: 9, offset: 18971},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 635, col: 9, offset: 18971},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 635, col: 9, offset: 18971},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 635, col: 9, offset: 18971},
									expr: &litMatcher{
										pos:        position{line: 635, col: 9, offset: 18971},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 635, col: 14, offset: 18976},
									expr: &charClassMatcher{
										pos:        position{line: 635, col: 14, offset: 18976},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 635, col: 21, offset: 18983},
									expr: &litMatcher{
										pos:        position{line: 635, col: 22, offset: 18984},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 642, col: 3, offset: 19159},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 642, col: 3, offset: 19159},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 642, col: 3, offset: 19159},
									expr: &litMatcher{
										pos:        position{line: 642, col: 3, offset: 19159},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 642, col: 8, offset: 19164},
									expr: &charClassMatcher{
										pos:        position{line: 642, col: 8, offset: 19164},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 642, col: 15, offset: 19171},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 642, col: 19, offset: 19175},
									expr: &charClassMatcher{
										pos:        position{line: 642, col: 19, offset: 19175},
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
						pos: position{line: 649, col: 3, offset: 19364},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 649, col: 3, offset: 19364},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 653, col: 3, offset: 19449},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 653, col: 3, offset: 19449},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 656, col: 3, offset: 19535},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 657, col: 3, offset: 19542},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 658, col: 3, offset: 19558},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 658, col: 3, offset: 19558},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 658, col: 3, offset: 19558},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 658, col: 7, offset: 19562},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 658, col: 12, offset: 19567},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 658, col: 12, offset: 19567},
												expr: &ruleRefExpr{
													pos:  position{line: 658, col: 13, offset: 19568},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 658, col: 25, offset: 19580,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 658, col: 28, offset: 19583},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 660, col: 5, offset: 19675},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 660, col: 20, offset: 19690},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 660, col: 37, offset: 19707},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 662, col: 1, offset: 19724},
			expr: &actionExpr{
				pos: position{line: 662, col: 8, offset: 19731},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 662, col: 8, offset: 19731},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 666, col: 1, offset: 19794},
			expr: &actionExpr{
				pos: position{line: 666, col: 17, offset: 19810},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 666, col: 17, offset: 19810},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 666, col: 17, offset: 19810},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 666, col: 21, offset: 19814},
							expr: &seqExpr{
								pos: position{line: 666, col: 22, offset: 19815},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 666, col: 22, offset: 19815},
										expr: &ruleRefExpr{
											pos:  position{line: 666, col: 23, offset: 19816},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 666, col: 35, offset: 19828,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 666, col: 39, offset: 19832},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 674, col: 1, offset: 20015},
			expr: &actionExpr{
				pos: position{line: 674, col: 10, offset: 20024},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 674, col: 11, offset: 20025},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 678, col: 1, offset: 20080},
			expr: &seqExpr{
				pos: position{line: 678, col: 12, offset: 20091},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 678, col: 13, offset: 20092},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 678, col: 13, offset: 20092},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 21, offset: 20100},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 28, offset: 20107},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 37, offset: 20116},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 48, offset: 20127},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 57, offset: 20136},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 66, offset: 20145},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 76, offset: 20155},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 678, col: 88, offset: 20167},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 678, col: 97, offset: 20176},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 678, col: 107, offset: 20186},
						expr: &oneOrMoreExpr{
							pos: position{line: 678, col: 108, offset: 20187},
							expr: &charClassMatcher{
								pos:        position{line: 678, col: 108, offset: 20187},
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
			pos:  position{line: 680, col: 1, offset: 20195},
			expr: &choiceExpr{
				pos: position{line: 680, col: 12, offset: 20206},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 680, col: 12, offset: 20206},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 680, col: 14, offset: 20208},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 680, col: 14, offset: 20208},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 24, offset: 20218},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 32, offset: 20226},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 41, offset: 20235},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 52, offset: 20246},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 61, offset: 20255},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 70, offset: 20264},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 680, col: 80, offset: 20274},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 683, col: 3, offset: 20372},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 685, col: 1, offset: 20378},
			expr: &charClassMatcher{
				pos:        position{line: 685, col: 15, offset: 20392},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 687, col: 1, offset: 20408},
			expr: &choiceExpr{
				pos: position{line: 687, col: 18, offset: 20425},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 687, col: 18, offset: 20425},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 687, col: 37, offset: 20444},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 689, col: 1, offset: 20459},
			expr: &charClassMatcher{
				pos:        position{line: 689, col: 20, offset: 20478},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 691, col: 1, offset: 20491},
			expr: &charClassMatcher{
				pos:        position{line: 691, col: 16, offset: 20506},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 693, col: 1, offset: 20513},
			expr: &charClassMatcher{
				pos:        position{line: 693, col: 23, offset: 20535},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 695, col: 1, offset: 20542},
			expr: &charClassMatcher{
				pos:        position{line: 695, col: 12, offset: 20553},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 697, col: 1, offset: 20564},
			expr: &choiceExpr{
				pos: position{line: 697, col: 22, offset: 20585},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 697, col: 22, offset: 20585},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 697, col: 33, offset: 20596},
						expr: &charClassMatcher{
							pos:        position{line: 697, col: 33, offset: 20596},
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
			pos:         position{line: 699, col: 1, offset: 20608},
			expr: &choiceExpr{
				pos: position{line: 699, col: 21, offset: 20628},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 699, col: 21, offset: 20628},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 699, col: 32, offset: 20639},
						expr: &charClassMatcher{
							pos:        position{line: 699, col: 32, offset: 20639},
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
			pos:         position{line: 701, col: 1, offset: 20651},
			expr: &oneOrMoreExpr{
				pos: position{line: 701, col: 33, offset: 20683},
				expr: &charClassMatcher{
					pos:        position{line: 701, col: 33, offset: 20683},
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
			pos:         position{line: 703, col: 1, offset: 20691},
			expr: &zeroOrMoreExpr{
				pos: position{line: 703, col: 32, offset: 20722},
				expr: &charClassMatcher{
					pos:        position{line: 703, col: 32, offset: 20722},
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
			pos:         position{line: 705, col: 1, offset: 20730},
			expr: &choiceExpr{
				pos: position{line: 705, col: 15, offset: 20744},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 705, col: 15, offset: 20744},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 705, col: 26, offset: 20755},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 705, col: 26, offset: 20755},
								expr: &charClassMatcher{
									pos:        position{line: 705, col: 26, offset: 20755},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 705, col: 35, offset: 20764},
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
			pos:  position{line: 707, col: 1, offset: 20770},
			expr: &oneOrMoreExpr{
				pos: position{line: 707, col: 12, offset: 20781},
				expr: &ruleRefExpr{
					pos:  position{line: 707, col: 13, offset: 20782},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 709, col: 1, offset: 20793},
			expr: &choiceExpr{
				pos: position{line: 709, col: 11, offset: 20803},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 709, col: 11, offset: 20803},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 11, offset: 20803},
								expr: &charClassMatcher{
									pos:        position{line: 709, col: 11, offset: 20803},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 22, offset: 20814},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 27, offset: 20819},
								expr: &seqExpr{
									pos: position{line: 709, col: 28, offset: 20820},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 709, col: 28, offset: 20820},
											expr: &charClassMatcher{
												pos:        position{line: 709, col: 29, offset: 20821},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 709, col: 34, offset: 20826,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 709, col: 38, offset: 20830},
								expr: &litMatcher{
									pos:        position{line: 709, col: 39, offset: 20831},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 709, col: 46, offset: 20838},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 46, offset: 20838},
								expr: &charClassMatcher{
									pos:        position{line: 709, col: 46, offset: 20838},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 57, offset: 20849},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 62, offset: 20854},
								expr: &seqExpr{
									pos: position{line: 709, col: 63, offset: 20855},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 709, col: 63, offset: 20855},
											expr: &litMatcher{
												pos:        position{line: 709, col: 64, offset: 20856},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 709, col: 69, offset: 20861,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 73, offset: 20865},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 709, col: 78, offset: 20870},
								expr: &charClassMatcher{
									pos:        position{line: 709, col: 78, offset: 20870},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 709, col: 84, offset: 20876},
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
			pos:  position{line: 711, col: 1, offset: 20882},
			expr: &notExpr{
				pos: position{line: 711, col: 7, offset: 20888},
				expr: &anyMatcher{
					line: 711, col: 8, offset: 20889,
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
