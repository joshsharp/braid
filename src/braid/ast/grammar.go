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
							pos:  position{line: 46, col: 108, offset: 1288},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 111, offset: 1291},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 115, offset: 1295},
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "ExternType",
			pos:  position{line: 52, col: 1, offset: 1511},
			expr: &choiceExpr{
				pos: position{line: 52, col: 14, offset: 1524},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 52, col: 14, offset: 1524},
						run: (*parser).callonExternType2,
						expr: &seqExpr{
							pos: position{line: 52, col: 14, offset: 1524},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 14, offset: 1524},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 16, offset: 1526},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 25, offset: 1535},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 52, col: 29, offset: 1539},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 36, offset: 1546},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 39, offset: 1549},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 44, offset: 1554},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 55, offset: 1565},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 57, offset: 1567},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 61, offset: 1571},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 5, offset: 1577},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 53, col: 16, offset: 1588},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 30, offset: 1602},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 32, offset: 1604},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 36, offset: 1608},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 38, offset: 1610},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 53, col: 44, offset: 1616},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 60, offset: 1632},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 53, col: 62, offset: 1634},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 53, col: 67, offset: 1639},
										expr: &seqExpr{
											pos: position{line: 53, col: 68, offset: 1640},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 53, col: 68, offset: 1640},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 53, col: 72, offset: 1644},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 53, col: 74, offset: 1646},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 53, col: 90, offset: 1662},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 53, col: 94, offset: 1666},
									expr: &litMatcher{
										pos:        position{line: 53, col: 94, offset: 1666},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 99, offset: 1671},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 53, col: 101, offset: 1673},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 105, offset: 1677},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 72, col: 1, offset: 2226},
						run: (*parser).callonExternType34,
						expr: &seqExpr{
							pos: position{line: 72, col: 1, offset: 2226},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 72, col: 1, offset: 2226},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 3, offset: 2228},
									val:        "extern",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 12, offset: 2237},
									name: "__N",
								},
								&litMatcher{
									pos:        position{line: 72, col: 16, offset: 2241},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 23, offset: 2248},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 26, offset: 2251},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 72, col: 31, offset: 2256},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 42, offset: 2267},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 44, offset: 2269},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 48, offset: 2273},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 5, offset: 2279},
									label: "importName",
									expr: &ruleRefExpr{
										pos:  position{line: 73, col: 16, offset: 2290},
										name: "StringLiteral",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 30, offset: 2304},
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
			pos:  position{line: 81, col: 1, offset: 2499},
			expr: &choiceExpr{
				pos: position{line: 81, col: 12, offset: 2510},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 81, col: 12, offset: 2510},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 81, col: 12, offset: 2510},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 81, col: 12, offset: 2510},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 81, col: 14, offset: 2512},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 81, col: 21, offset: 2519},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 81, col: 24, offset: 2522},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 81, col: 29, offset: 2527},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 81, col: 40, offset: 2538},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 81, col: 47, offset: 2545},
										expr: &seqExpr{
											pos: position{line: 81, col: 48, offset: 2546},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 81, col: 48, offset: 2546},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 81, col: 51, offset: 2549},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 81, col: 67, offset: 2565},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 81, col: 69, offset: 2567},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 81, col: 73, offset: 2571},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 81, col: 79, offset: 2577},
										expr: &seqExpr{
											pos: position{line: 81, col: 80, offset: 2578},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 81, col: 80, offset: 2578},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 81, col: 83, offset: 2581},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 81, col: 100, offset: 2598},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 100, col: 1, offset: 3094},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 100, col: 1, offset: 3094},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 100, col: 1, offset: 3094},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 3, offset: 3096},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 10, offset: 3103},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 100, col: 13, offset: 3106},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 18, offset: 3111},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 100, col: 29, offset: 3122},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 100, col: 36, offset: 3129},
										expr: &seqExpr{
											pos: position{line: 100, col: 37, offset: 3130},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 100, col: 37, offset: 3130},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 100, col: 40, offset: 3133},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 56, offset: 3149},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 58, offset: 3151},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 62, offset: 3155},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 5, offset: 3161},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 9, offset: 3165},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 11, offset: 3167},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 101, col: 17, offset: 3173},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 33, offset: 3189},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 101, col: 35, offset: 3191},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 101, col: 40, offset: 3196},
										expr: &seqExpr{
											pos: position{line: 101, col: 41, offset: 3197},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 101, col: 41, offset: 3197},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 45, offset: 3201},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 47, offset: 3203},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 101, col: 63, offset: 3219},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 101, col: 67, offset: 3223},
									expr: &litMatcher{
										pos:        position{line: 101, col: 67, offset: 3223},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 72, offset: 3228},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 101, col: 74, offset: 3230},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 101, col: 78, offset: 3234},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 119, col: 1, offset: 3719},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 119, col: 1, offset: 3719},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 119, col: 1, offset: 3719},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 3, offset: 3721},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 10, offset: 3728},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 13, offset: 3731},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 119, col: 18, offset: 3736},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 119, col: 29, offset: 3747},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 119, col: 36, offset: 3754},
										expr: &seqExpr{
											pos: position{line: 119, col: 37, offset: 3755},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 119, col: 37, offset: 3755},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 119, col: 40, offset: 3758},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 56, offset: 3774},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 119, col: 58, offset: 3776},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 119, col: 62, offset: 3780},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 119, col: 64, offset: 3782},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 119, col: 69, offset: 3787},
										expr: &ruleRefExpr{
											pos:  position{line: 119, col: 70, offset: 3788},
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
			pos:  position{line: 134, col: 1, offset: 4195},
			expr: &actionExpr{
				pos: position{line: 134, col: 19, offset: 4213},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 134, col: 19, offset: 4213},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 134, col: 19, offset: 4213},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 24, offset: 4218},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 134, col: 37, offset: 4231},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 134, col: 39, offset: 4233},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 134, col: 43, offset: 4237},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 134, col: 45, offset: 4239},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 48, offset: 4242},
								name: "TypeAnnotation",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 138, col: 1, offset: 4343},
			expr: &choiceExpr{
				pos: position{line: 138, col: 22, offset: 4364},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 138, col: 22, offset: 4364},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 138, col: 22, offset: 4364},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 138, col: 22, offset: 4364},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 26, offset: 4368},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 28, offset: 4370},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 33, offset: 4375},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 44, offset: 4386},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 46, offset: 4388},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 50, offset: 4392},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 52, offset: 4394},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 58, offset: 4400},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 74, offset: 4416},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 76, offset: 4418},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 138, col: 81, offset: 4423},
										expr: &seqExpr{
											pos: position{line: 138, col: 82, offset: 4424},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 138, col: 82, offset: 4424},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 138, col: 86, offset: 4428},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 138, col: 88, offset: 4430},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 138, col: 104, offset: 4446},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 138, col: 108, offset: 4450},
									expr: &litMatcher{
										pos:        position{line: 138, col: 108, offset: 4450},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 113, offset: 4455},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 115, offset: 4457},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 119, offset: 4461},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 157, col: 1, offset: 5066},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 157, col: 1, offset: 5066},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 157, col: 1, offset: 5066},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 5, offset: 5070},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 157, col: 7, offset: 5072},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 157, col: 12, offset: 5077},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 157, col: 23, offset: 5088},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 157, col: 28, offset: 5093},
										expr: &seqExpr{
											pos: position{line: 157, col: 29, offset: 5094},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 157, col: 29, offset: 5094},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 157, col: 32, offset: 5097},
													name: "TypeAnnotation",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 49, offset: 5114},
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
			pos:  position{line: 174, col: 1, offset: 5551},
			expr: &choiceExpr{
				pos: position{line: 174, col: 14, offset: 5564},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 174, col: 14, offset: 5564},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 174, col: 14, offset: 5564},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 174, col: 14, offset: 5564},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 174, col: 16, offset: 5566},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 174, col: 22, offset: 5572},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 174, col: 26, offset: 5576},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 174, col: 28, offset: 5578},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 174, col: 39, offset: 5589},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 174, col: 42, offset: 5592},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 174, col: 46, offset: 5596},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 174, col: 49, offset: 5599},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 174, col: 54, offset: 5604},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 174, col: 59, offset: 5609},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 180, col: 1, offset: 5728},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 180, col: 1, offset: 5728},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 180, col: 1, offset: 5728},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 180, col: 3, offset: 5730},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 9, offset: 5736},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 180, col: 13, offset: 5740},
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 14, offset: 5741},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 184, col: 1, offset: 5839},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 184, col: 1, offset: 5839},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 184, col: 1, offset: 5839},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 184, col: 3, offset: 5841},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 9, offset: 5847},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 184, col: 13, offset: 5851},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 184, col: 15, offset: 5853},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 26, offset: 5864},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 184, col: 29, offset: 5867},
									expr: &litMatcher{
										pos:        position{line: 184, col: 30, offset: 5868},
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
			pos:  position{line: 188, col: 1, offset: 5962},
			expr: &actionExpr{
				pos: position{line: 188, col: 12, offset: 5973},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 188, col: 12, offset: 5973},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 188, col: 12, offset: 5973},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 188, col: 14, offset: 5975},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 20, offset: 5981},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 188, col: 24, offset: 5985},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 26, offset: 5987},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 39, offset: 6000},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 188, col: 42, offset: 6003},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 46, offset: 6007},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 188, col: 49, offset: 6010},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 188, col: 53, offset: 6014},
								expr: &seqExpr{
									pos: position{line: 188, col: 54, offset: 6015},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 188, col: 54, offset: 6015},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 188, col: 63, offset: 6024},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 67, offset: 6028},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 188, col: 71, offset: 6032},
								expr: &seqExpr{
									pos: position{line: 188, col: 72, offset: 6033},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 188, col: 72, offset: 6033},
											name: "TypeAnnotation",
										},
										&ruleRefExpr{
											pos:  position{line: 188, col: 87, offset: 6048},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 188, col: 91, offset: 6052},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 95, offset: 6056},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 188, col: 98, offset: 6059},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 188, col: 109, offset: 6070},
								expr: &ruleRefExpr{
									pos:  position{line: 188, col: 110, offset: 6071},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 122, offset: 6083},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 188, col: 124, offset: 6085},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 128, offset: 6089},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 214, col: 1, offset: 6699},
			expr: &actionExpr{
				pos: position{line: 214, col: 8, offset: 6706},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 214, col: 8, offset: 6706},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 214, col: 12, offset: 6710},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 214, col: 12, offset: 6710},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 214, col: 21, offset: 6719},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 220, col: 1, offset: 6836},
			expr: &choiceExpr{
				pos: position{line: 220, col: 10, offset: 6845},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 220, col: 10, offset: 6845},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 220, col: 10, offset: 6845},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 220, col: 10, offset: 6845},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 220, col: 12, offset: 6847},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 17, offset: 6852},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 220, col: 21, offset: 6856},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 220, col: 26, offset: 6861},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 36, offset: 6871},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 220, col: 39, offset: 6874},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 43, offset: 6878},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 220, col: 45, offset: 6880},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 220, col: 51, offset: 6886},
										expr: &ruleRefExpr{
											pos:  position{line: 220, col: 52, offset: 6887},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 64, offset: 6899},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 220, col: 67, offset: 6902},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 71, offset: 6906},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 220, col: 74, offset: 6909},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 81, offset: 6916},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 220, col: 84, offset: 6919},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 88, offset: 6923},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 220, col: 90, offset: 6925},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 220, col: 96, offset: 6931},
										expr: &ruleRefExpr{
											pos:  position{line: 220, col: 97, offset: 6932},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 109, offset: 6944},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 220, col: 112, offset: 6947},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 239, col: 1, offset: 7450},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 239, col: 1, offset: 7450},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 239, col: 1, offset: 7450},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 3, offset: 7452},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 8, offset: 7457},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 239, col: 12, offset: 7461},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 239, col: 17, offset: 7466},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 27, offset: 7476},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 239, col: 30, offset: 7479},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 34, offset: 7483},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 239, col: 36, offset: 7485},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 239, col: 42, offset: 7491},
										expr: &ruleRefExpr{
											pos:  position{line: 239, col: 43, offset: 7492},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 55, offset: 7504},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 239, col: 57, offset: 7506},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 61, offset: 7510},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 239, col: 64, offset: 7513},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 239, col: 71, offset: 7520},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 239, col: 79, offset: 7528},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 251, col: 1, offset: 7858},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 251, col: 1, offset: 7858},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 251, col: 1, offset: 7858},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 251, col: 3, offset: 7860},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 251, col: 8, offset: 7865},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 251, col: 12, offset: 7869},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 251, col: 17, offset: 7874},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 251, col: 27, offset: 7884},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 251, col: 30, offset: 7887},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 251, col: 34, offset: 7891},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 251, col: 36, offset: 7893},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 251, col: 42, offset: 7899},
										expr: &ruleRefExpr{
											pos:  position{line: 251, col: 43, offset: 7900},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 251, col: 55, offset: 7912},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 251, col: 58, offset: 7915},
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
			pos:  position{line: 263, col: 1, offset: 8213},
			expr: &choiceExpr{
				pos: position{line: 263, col: 8, offset: 8220},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 263, col: 8, offset: 8220},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 263, col: 8, offset: 8220},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 263, col: 8, offset: 8220},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 263, col: 10, offset: 8222},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 263, col: 17, offset: 8229},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 263, col: 28, offset: 8240},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 263, col: 32, offset: 8244},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 263, col: 35, offset: 8247},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 263, col: 48, offset: 8260},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 263, col: 53, offset: 8265},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 277, col: 1, offset: 8589},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 277, col: 1, offset: 8589},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 277, col: 1, offset: 8589},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 277, col: 3, offset: 8591},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 277, col: 6, offset: 8594},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 277, col: 19, offset: 8607},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 277, col: 24, offset: 8612},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 291, col: 1, offset: 8929},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 291, col: 1, offset: 8929},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 291, col: 1, offset: 8929},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 291, col: 3, offset: 8931},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 291, col: 6, offset: 8934},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 291, col: 19, offset: 8947},
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
			pos:  position{line: 298, col: 1, offset: 9118},
			expr: &actionExpr{
				pos: position{line: 298, col: 16, offset: 9133},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 298, col: 16, offset: 9133},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 298, col: 16, offset: 9133},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 23, offset: 9140},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 298, col: 36, offset: 9153},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 298, col: 41, offset: 9158},
								expr: &seqExpr{
									pos: position{line: 298, col: 42, offset: 9159},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 298, col: 42, offset: 9159},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 298, col: 46, offset: 9163},
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
			pos:  position{line: 315, col: 1, offset: 9600},
			expr: &actionExpr{
				pos: position{line: 315, col: 12, offset: 9611},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 315, col: 12, offset: 9611},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 315, col: 12, offset: 9611},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 315, col: 16, offset: 9615},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 315, col: 18, offset: 9617},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 315, col: 27, offset: 9626},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 315, col: 35, offset: 9634},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 315, col: 37, offset: 9636},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 315, col: 42, offset: 9641},
								expr: &seqExpr{
									pos: position{line: 315, col: 43, offset: 9642},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 315, col: 43, offset: 9642},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 47, offset: 9646},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 49, offset: 9648},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 315, col: 59, offset: 9658},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 315, col: 61, offset: 9660},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 333, col: 1, offset: 10082},
			expr: &actionExpr{
				pos: position{line: 333, col: 11, offset: 10092},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 333, col: 11, offset: 10092},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 333, col: 11, offset: 10092},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 333, col: 16, offset: 10097},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 333, col: 27, offset: 10108},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 333, col: 29, offset: 10110},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 333, col: 34, offset: 10115},
								expr: &seqExpr{
									pos: position{line: 333, col: 35, offset: 10116},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 333, col: 35, offset: 10116},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 333, col: 39, offset: 10120},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 333, col: 41, offset: 10122},
											name: "TypeAnnotation",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 333, col: 59, offset: 10140},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "TypeAnnotation",
			pos:  position{line: 353, col: 1, offset: 10628},
			expr: &choiceExpr{
				pos: position{line: 353, col: 18, offset: 10645},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 353, col: 18, offset: 10645},
						name: "AnyType",
					},
					&ruleRefExpr{
						pos:  position{line: 353, col: 28, offset: 10655},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 355, col: 1, offset: 10667},
			expr: &choiceExpr{
				pos: position{line: 355, col: 11, offset: 10677},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 355, col: 11, offset: 10677},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 355, col: 22, offset: 10688},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 357, col: 1, offset: 10703},
			expr: &choiceExpr{
				pos: position{line: 357, col: 13, offset: 10715},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 357, col: 13, offset: 10715},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 357, col: 13, offset: 10715},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 357, col: 13, offset: 10715},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 357, col: 17, offset: 10719},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 357, col: 19, offset: 10721},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 28, offset: 10730},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 357, col: 40, offset: 10742},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 357, col: 42, offset: 10744},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 357, col: 47, offset: 10749},
										expr: &seqExpr{
											pos: position{line: 357, col: 48, offset: 10750},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 357, col: 48, offset: 10750},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 357, col: 52, offset: 10754},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 357, col: 54, offset: 10756},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 357, col: 68, offset: 10770},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 357, col: 70, offset: 10772},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 374, col: 1, offset: 11194},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 374, col: 1, offset: 11194},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 374, col: 1, offset: 11194},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 374, col: 5, offset: 11198},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 374, col: 7, offset: 11200},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 374, col: 16, offset: 11209},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 374, col: 21, offset: 11214},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 374, col: 23, offset: 11216},
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
			pos:  position{line: 379, col: 1, offset: 11321},
			expr: &actionExpr{
				pos: position{line: 379, col: 16, offset: 11336},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 379, col: 16, offset: 11336},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 379, col: 16, offset: 11336},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 379, col: 18, offset: 11338},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 379, col: 21, offset: 11341},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 379, col: 27, offset: 11347},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 379, col: 32, offset: 11352},
								expr: &seqExpr{
									pos: position{line: 379, col: 33, offset: 11353},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 379, col: 33, offset: 11353},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 379, col: 37, offset: 11357},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 379, col: 46, offset: 11366},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 379, col: 50, offset: 11370},
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
			pos:  position{line: 399, col: 1, offset: 11976},
			expr: &choiceExpr{
				pos: position{line: 399, col: 9, offset: 11984},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 399, col: 9, offset: 11984},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 399, col: 21, offset: 11996},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 399, col: 37, offset: 12012},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 399, col: 48, offset: 12023},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 399, col: 60, offset: 12035},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 401, col: 1, offset: 12048},
			expr: &actionExpr{
				pos: position{line: 401, col: 13, offset: 12060},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 401, col: 13, offset: 12060},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 401, col: 13, offset: 12060},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 401, col: 15, offset: 12062},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 401, col: 21, offset: 12068},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 401, col: 35, offset: 12082},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 401, col: 40, offset: 12087},
								expr: &seqExpr{
									pos: position{line: 401, col: 41, offset: 12088},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 401, col: 41, offset: 12088},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 401, col: 45, offset: 12092},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 401, col: 61, offset: 12108},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 401, col: 65, offset: 12112},
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
			pos:  position{line: 434, col: 1, offset: 13005},
			expr: &actionExpr{
				pos: position{line: 434, col: 17, offset: 13021},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 434, col: 17, offset: 13021},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 434, col: 17, offset: 13021},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 434, col: 19, offset: 13023},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 434, col: 25, offset: 13029},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 434, col: 34, offset: 13038},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 434, col: 39, offset: 13043},
								expr: &seqExpr{
									pos: position{line: 434, col: 40, offset: 13044},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 434, col: 40, offset: 13044},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 434, col: 44, offset: 13048},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 434, col: 61, offset: 13065},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 434, col: 65, offset: 13069},
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
			pos:  position{line: 466, col: 1, offset: 13956},
			expr: &actionExpr{
				pos: position{line: 466, col: 12, offset: 13967},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 466, col: 12, offset: 13967},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 466, col: 12, offset: 13967},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 466, col: 14, offset: 13969},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 466, col: 20, offset: 13975},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 466, col: 30, offset: 13985},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 466, col: 35, offset: 13990},
								expr: &seqExpr{
									pos: position{line: 466, col: 36, offset: 13991},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 466, col: 36, offset: 13991},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 466, col: 40, offset: 13995},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 466, col: 52, offset: 14007},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 466, col: 56, offset: 14011},
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
			pos:  position{line: 498, col: 1, offset: 14899},
			expr: &actionExpr{
				pos: position{line: 498, col: 13, offset: 14911},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 498, col: 13, offset: 14911},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 498, col: 13, offset: 14911},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 498, col: 15, offset: 14913},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 498, col: 21, offset: 14919},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 498, col: 33, offset: 14931},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 498, col: 38, offset: 14936},
								expr: &seqExpr{
									pos: position{line: 498, col: 39, offset: 14937},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 498, col: 39, offset: 14937},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 498, col: 43, offset: 14941},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 498, col: 56, offset: 14954},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 498, col: 60, offset: 14958},
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
			pos:  position{line: 529, col: 1, offset: 15847},
			expr: &choiceExpr{
				pos: position{line: 529, col: 15, offset: 15861},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 529, col: 15, offset: 15861},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 529, col: 15, offset: 15861},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 529, col: 15, offset: 15861},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 529, col: 17, offset: 15863},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 529, col: 21, offset: 15867},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 529, col: 24, offset: 15870},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 529, col: 30, offset: 15876},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 529, col: 36, offset: 15882},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 529, col: 39, offset: 15885},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 532, col: 5, offset: 16008},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 534, col: 1, offset: 16015},
			expr: &choiceExpr{
				pos: position{line: 534, col: 12, offset: 16026},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 534, col: 12, offset: 16026},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 534, col: 30, offset: 16044},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 534, col: 49, offset: 16063},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 534, col: 64, offset: 16078},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 536, col: 1, offset: 16091},
			expr: &actionExpr{
				pos: position{line: 536, col: 19, offset: 16109},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 536, col: 21, offset: 16111},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 536, col: 21, offset: 16111},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 536, col: 28, offset: 16118},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 540, col: 1, offset: 16200},
			expr: &actionExpr{
				pos: position{line: 540, col: 20, offset: 16219},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 540, col: 22, offset: 16221},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 540, col: 22, offset: 16221},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 540, col: 29, offset: 16228},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 540, col: 36, offset: 16235},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 540, col: 42, offset: 16241},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 540, col: 48, offset: 16247},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 540, col: 56, offset: 16255},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 544, col: 1, offset: 16334},
			expr: &choiceExpr{
				pos: position{line: 544, col: 16, offset: 16349},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 544, col: 16, offset: 16349},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 544, col: 18, offset: 16351},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 544, col: 18, offset: 16351},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 544, col: 24, offset: 16357},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 547, col: 3, offset: 16440},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 547, col: 5, offset: 16442},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 550, col: 3, offset: 16522},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 550, col: 3, offset: 16522},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 554, col: 1, offset: 16603},
			expr: &actionExpr{
				pos: position{line: 554, col: 15, offset: 16617},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 554, col: 17, offset: 16619},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 554, col: 17, offset: 16619},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 554, col: 23, offset: 16625},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 558, col: 1, offset: 16707},
			expr: &choiceExpr{
				pos: position{line: 558, col: 9, offset: 16715},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 558, col: 9, offset: 16715},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 558, col: 16, offset: 16722},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 558, col: 31, offset: 16737},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 558, col: 46, offset: 16752},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 560, col: 1, offset: 16759},
			expr: &choiceExpr{
				pos: position{line: 560, col: 14, offset: 16772},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 560, col: 14, offset: 16772},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 560, col: 29, offset: 16787},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 562, col: 1, offset: 16795},
			expr: &choiceExpr{
				pos: position{line: 562, col: 14, offset: 16808},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 562, col: 14, offset: 16808},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 562, col: 29, offset: 16823},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 564, col: 1, offset: 16835},
			expr: &actionExpr{
				pos: position{line: 564, col: 16, offset: 16850},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 564, col: 16, offset: 16850},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 564, col: 16, offset: 16850},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 564, col: 20, offset: 16854},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 564, col: 22, offset: 16856},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 564, col: 28, offset: 16862},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 564, col: 33, offset: 16867},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 564, col: 35, offset: 16869},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 564, col: 40, offset: 16874},
								expr: &seqExpr{
									pos: position{line: 564, col: 41, offset: 16875},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 564, col: 41, offset: 16875},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 564, col: 45, offset: 16879},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 564, col: 47, offset: 16881},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 564, col: 52, offset: 16886},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 564, col: 56, offset: 16890},
							expr: &litMatcher{
								pos:        position{line: 564, col: 56, offset: 16890},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 564, col: 61, offset: 16895},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 564, col: 63, offset: 16897},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 580, col: 1, offset: 17342},
			expr: &actionExpr{
				pos: position{line: 580, col: 19, offset: 17360},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 580, col: 19, offset: 17360},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 580, col: 19, offset: 17360},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 580, col: 24, offset: 17365},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 580, col: 35, offset: 17376},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 580, col: 37, offset: 17378},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 580, col: 42, offset: 17383},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 593, col: 1, offset: 17653},
			expr: &actionExpr{
				pos: position{line: 593, col: 18, offset: 17670},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 593, col: 18, offset: 17670},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 593, col: 18, offset: 17670},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 593, col: 23, offset: 17675},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 593, col: 34, offset: 17686},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 593, col: 36, offset: 17688},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 593, col: 40, offset: 17692},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 593, col: 42, offset: 17694},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 593, col: 52, offset: 17704},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 593, col: 65, offset: 17717},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 593, col: 67, offset: 17719},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 593, col: 71, offset: 17723},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 593, col: 73, offset: 17725},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 593, col: 84, offset: 17736},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 593, col: 89, offset: 17741},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 593, col: 94, offset: 17746},
								expr: &seqExpr{
									pos: position{line: 593, col: 95, offset: 17747},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 593, col: 95, offset: 17747},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 593, col: 99, offset: 17751},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 593, col: 101, offset: 17753},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 593, col: 114, offset: 17766},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 593, col: 116, offset: 17768},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 593, col: 120, offset: 17772},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 593, col: 122, offset: 17774},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 593, col: 130, offset: 17782},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 613, col: 1, offset: 18366},
			expr: &actionExpr{
				pos: position{line: 613, col: 17, offset: 18382},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 613, col: 17, offset: 18382},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 613, col: 17, offset: 18382},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 613, col: 22, offset: 18387},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 617, col: 1, offset: 18460},
			expr: &actionExpr{
				pos: position{line: 617, col: 16, offset: 18475},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 617, col: 16, offset: 18475},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 617, col: 16, offset: 18475},
							expr: &ruleRefExpr{
								pos:  position{line: 617, col: 17, offset: 18476},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 617, col: 27, offset: 18486},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 617, col: 27, offset: 18486},
									expr: &charClassMatcher{
										pos:        position{line: 617, col: 27, offset: 18486},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 617, col: 34, offset: 18493},
									expr: &charClassMatcher{
										pos:        position{line: 617, col: 34, offset: 18493},
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
			pos:  position{line: 621, col: 1, offset: 18568},
			expr: &actionExpr{
				pos: position{line: 621, col: 14, offset: 18581},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 621, col: 15, offset: 18582},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 621, col: 15, offset: 18582},
							expr: &charClassMatcher{
								pos:        position{line: 621, col: 15, offset: 18582},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 621, col: 22, offset: 18589},
							expr: &charClassMatcher{
								pos:        position{line: 621, col: 22, offset: 18589},
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
			pos:  position{line: 625, col: 1, offset: 18664},
			expr: &choiceExpr{
				pos: position{line: 625, col: 9, offset: 18672},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 625, col: 9, offset: 18672},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 625, col: 9, offset: 18672},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 625, col: 9, offset: 18672},
									expr: &litMatcher{
										pos:        position{line: 625, col: 9, offset: 18672},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 625, col: 14, offset: 18677},
									expr: &charClassMatcher{
										pos:        position{line: 625, col: 14, offset: 18677},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 625, col: 21, offset: 18684},
									expr: &litMatcher{
										pos:        position{line: 625, col: 22, offset: 18685},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 632, col: 3, offset: 18860},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 632, col: 3, offset: 18860},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 632, col: 3, offset: 18860},
									expr: &litMatcher{
										pos:        position{line: 632, col: 3, offset: 18860},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 632, col: 8, offset: 18865},
									expr: &charClassMatcher{
										pos:        position{line: 632, col: 8, offset: 18865},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 632, col: 15, offset: 18872},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 632, col: 19, offset: 18876},
									expr: &charClassMatcher{
										pos:        position{line: 632, col: 19, offset: 18876},
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
						pos: position{line: 639, col: 3, offset: 19065},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 639, col: 3, offset: 19065},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 643, col: 3, offset: 19150},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 643, col: 3, offset: 19150},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 646, col: 3, offset: 19236},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 647, col: 3, offset: 19243},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 648, col: 3, offset: 19259},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 648, col: 3, offset: 19259},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 648, col: 3, offset: 19259},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 648, col: 7, offset: 19263},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 648, col: 12, offset: 19268},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 648, col: 12, offset: 19268},
												expr: &ruleRefExpr{
													pos:  position{line: 648, col: 13, offset: 19269},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 648, col: 25, offset: 19281,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 648, col: 28, offset: 19284},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 650, col: 5, offset: 19376},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 650, col: 20, offset: 19391},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 650, col: 37, offset: 19408},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 652, col: 1, offset: 19425},
			expr: &actionExpr{
				pos: position{line: 652, col: 8, offset: 19432},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 652, col: 8, offset: 19432},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 656, col: 1, offset: 19495},
			expr: &actionExpr{
				pos: position{line: 656, col: 17, offset: 19511},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 656, col: 17, offset: 19511},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 656, col: 17, offset: 19511},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 656, col: 21, offset: 19515},
							expr: &seqExpr{
								pos: position{line: 656, col: 22, offset: 19516},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 656, col: 22, offset: 19516},
										expr: &ruleRefExpr{
											pos:  position{line: 656, col: 23, offset: 19517},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 656, col: 35, offset: 19529,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 656, col: 39, offset: 19533},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 664, col: 1, offset: 19716},
			expr: &actionExpr{
				pos: position{line: 664, col: 10, offset: 19725},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 664, col: 11, offset: 19726},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 668, col: 1, offset: 19781},
			expr: &seqExpr{
				pos: position{line: 668, col: 12, offset: 19792},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 668, col: 13, offset: 19793},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 668, col: 13, offset: 19793},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 21, offset: 19801},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 28, offset: 19808},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 37, offset: 19817},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 48, offset: 19828},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 57, offset: 19837},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 66, offset: 19846},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 76, offset: 19856},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 668, col: 88, offset: 19868},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 668, col: 97, offset: 19877},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 668, col: 107, offset: 19887},
						expr: &oneOrMoreExpr{
							pos: position{line: 668, col: 108, offset: 19888},
							expr: &charClassMatcher{
								pos:        position{line: 668, col: 108, offset: 19888},
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
			pos:  position{line: 670, col: 1, offset: 19896},
			expr: &choiceExpr{
				pos: position{line: 670, col: 12, offset: 19907},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 670, col: 12, offset: 19907},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 670, col: 14, offset: 19909},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 670, col: 14, offset: 19909},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 24, offset: 19919},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 32, offset: 19927},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 41, offset: 19936},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 52, offset: 19947},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 61, offset: 19956},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 70, offset: 19965},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 80, offset: 19975},
									val:        "()",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 670, col: 87, offset: 19982},
									val:        "func",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 673, col: 3, offset: 20082},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 675, col: 1, offset: 20088},
			expr: &charClassMatcher{
				pos:        position{line: 675, col: 15, offset: 20102},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 677, col: 1, offset: 20118},
			expr: &choiceExpr{
				pos: position{line: 677, col: 18, offset: 20135},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 677, col: 18, offset: 20135},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 677, col: 37, offset: 20154},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 679, col: 1, offset: 20169},
			expr: &charClassMatcher{
				pos:        position{line: 679, col: 20, offset: 20188},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 681, col: 1, offset: 20201},
			expr: &charClassMatcher{
				pos:        position{line: 681, col: 16, offset: 20216},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 683, col: 1, offset: 20223},
			expr: &charClassMatcher{
				pos:        position{line: 683, col: 23, offset: 20245},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 685, col: 1, offset: 20252},
			expr: &charClassMatcher{
				pos:        position{line: 685, col: 12, offset: 20263},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 687, col: 1, offset: 20274},
			expr: &choiceExpr{
				pos: position{line: 687, col: 22, offset: 20295},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 687, col: 22, offset: 20295},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 687, col: 33, offset: 20306},
						expr: &charClassMatcher{
							pos:        position{line: 687, col: 33, offset: 20306},
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
			pos:         position{line: 689, col: 1, offset: 20318},
			expr: &choiceExpr{
				pos: position{line: 689, col: 21, offset: 20338},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 689, col: 21, offset: 20338},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 689, col: 32, offset: 20349},
						expr: &charClassMatcher{
							pos:        position{line: 689, col: 32, offset: 20349},
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
			pos:         position{line: 691, col: 1, offset: 20361},
			expr: &oneOrMoreExpr{
				pos: position{line: 691, col: 33, offset: 20393},
				expr: &charClassMatcher{
					pos:        position{line: 691, col: 33, offset: 20393},
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
			pos:         position{line: 693, col: 1, offset: 20401},
			expr: &zeroOrMoreExpr{
				pos: position{line: 693, col: 32, offset: 20432},
				expr: &charClassMatcher{
					pos:        position{line: 693, col: 32, offset: 20432},
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
			pos:         position{line: 695, col: 1, offset: 20440},
			expr: &choiceExpr{
				pos: position{line: 695, col: 15, offset: 20454},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 695, col: 15, offset: 20454},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 695, col: 26, offset: 20465},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 695, col: 26, offset: 20465},
								expr: &charClassMatcher{
									pos:        position{line: 695, col: 26, offset: 20465},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 695, col: 35, offset: 20474},
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
			pos:  position{line: 697, col: 1, offset: 20480},
			expr: &oneOrMoreExpr{
				pos: position{line: 697, col: 12, offset: 20491},
				expr: &ruleRefExpr{
					pos:  position{line: 697, col: 13, offset: 20492},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 699, col: 1, offset: 20503},
			expr: &choiceExpr{
				pos: position{line: 699, col: 11, offset: 20513},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 699, col: 11, offset: 20513},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 699, col: 11, offset: 20513},
								expr: &charClassMatcher{
									pos:        position{line: 699, col: 11, offset: 20513},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 699, col: 22, offset: 20524},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 699, col: 27, offset: 20529},
								expr: &seqExpr{
									pos: position{line: 699, col: 28, offset: 20530},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 699, col: 28, offset: 20530},
											expr: &charClassMatcher{
												pos:        position{line: 699, col: 29, offset: 20531},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 699, col: 34, offset: 20536,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 699, col: 38, offset: 20540},
								expr: &litMatcher{
									pos:        position{line: 699, col: 39, offset: 20541},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 699, col: 46, offset: 20548},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 699, col: 46, offset: 20548},
								expr: &charClassMatcher{
									pos:        position{line: 699, col: 46, offset: 20548},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 699, col: 57, offset: 20559},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 699, col: 62, offset: 20564},
								expr: &seqExpr{
									pos: position{line: 699, col: 63, offset: 20565},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 699, col: 63, offset: 20565},
											expr: &litMatcher{
												pos:        position{line: 699, col: 64, offset: 20566},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 699, col: 69, offset: 20571,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 699, col: 73, offset: 20575},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 699, col: 78, offset: 20580},
								expr: &charClassMatcher{
									pos:        position{line: 699, col: 78, offset: 20580},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 699, col: 84, offset: 20586},
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
			pos:  position{line: 701, col: 1, offset: 20592},
			expr: &notExpr{
				pos: position{line: 701, col: 7, offset: 20598},
				expr: &anyMatcher{
					line: 701, col: 8, offset: 20599,
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
		Arguments: args.(Container).Subvalues, ReturnAnnotation: ret.(BasicAst).StringValue}, nil
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
